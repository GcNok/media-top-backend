package logger

import (
	"context"
	"encoding/json"
	"fmt"
	golog "log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/pkg/errors"
)

const (
	escape = "\x1b"
	reset  = escape + "[0m"
)

const (
	StdoutDump = 1 << iota
	StdoutDebug
	StdoutInfo
	StdoutWarn
	StdoutCrit
	StdoutTrace
	FileDump
	FileDebug
	FileInfo
	FileWarn
	FileCrit
	FileTrace
)

type Logger struct {
	GoLogger  *golog.Logger
	WriteFlag int
	Path      string
	LastIno   uint64
	File      *os.File
}

type Level struct {
	Label     string // INFO|WARN|CRIT|DEBUG|TRACE
	TermBegin string // start color code
	TermEnd   string // reset color code
}

var (
	dumpLevel = Level{
		Label:     "DUMP",
		TermBegin: escape + "[34m" + escape + "[1m",
		TermEnd:   reset,
	}
	debugLevel = Level{
		Label:     "DEBUG",
		TermBegin: escape + "[36m" + escape + "[1m",
		TermEnd:   reset,
	}
	infoLevel = Level{
		Label:     "INFO",
		TermBegin: escape + "[32m" + escape + "[1m",
		TermEnd:   reset,
	}
	warnLevel = Level{
		Label:     "WARN",
		TermBegin: escape + "[33m" + escape + "[1m",
		TermEnd:   reset,
	}
	critLevel = Level{
		Label:     "CRIT",
		TermBegin: escape + "[31m" + escape + "[1m",
		TermEnd:   reset,
	}
	traceLevel = Level{
		Label:     "TRACE",
		TermBegin: escape + "[93m" + escape + "[1m",
		TermEnd:   reset,
	}
	loggerContainer = map[string]Logger{}
)

func getContextValue(ctx context.Context, key string) (interface{}, bool) {
	val := ctx.Value(key)
	return val, val != nil
}

func Add(alias string, path string, writeMode int) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		d := filepath.Dir(path)
		if _, err := os.Stat(d); os.IsNotExist(err) {
			err = os.Mkdir(d, 0755)
			return errors.Wrap(err, fmt.Sprintf("failed to create dir. %s", d))
		}

		return errors.Wrap(err, "failed to open logfile. "+path)
	}
	goLogger := golog.New(file, "", 0)
	var stat syscall.Stat_t
	err = syscall.Stat(path, &stat)
	if err != nil {
		return errors.Wrap(err, "failed to get stat:"+path)
	}
	logger := Logger{GoLogger: goLogger, WriteFlag: writeMode, Path: path, File: file, LastIno: stat.Ino}
	loggerContainer[alias] = logger
	return nil
}

func (l *Logger) printf(ctx context.Context, lv *Level, enableStdout bool, enableWrite bool, msg interface{}) {
	if enableStdout == false && enableWrite == false {
		return
	}

	if enableWrite {
		err := l.reopenIfFileIsMissed()
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to reopen if file is missed"))
			return
		}
	}

	var (
		now          = time.Now()
		requestID    = "none"
		userID       = "unknown"
		customerID   = "unknown"
		servicerID   = "unknown"
		method, path string
	)

	r, ok := getContextValue(ctx, "RequestID")
	if ok {
		requestID = r.(string)
	}
	md, ok := getContextValue(ctx, "RequestMethod")
	if ok {
		method = md.(string)
	}
	pt, ok := getContextValue(ctx, "RequestPath")
	if ok {
		path = pt.(string)
	}

	_, file, line, ok := runtime.Caller(2)
	if !ok {
		fmt.Println("failed to get caller")
		file = "???"
		line = 0
	}

	caller := file + ":" + strconv.Itoa(line)

	if enableWrite {
		bytes, err := json.Marshal(map[string]interface{}{
			"time":        now,
			"user_id":     userID,
			"customer_id": customerID,
			"servicer_id": servicerID,
			"request_id":  requestID,
			"level":       lv.Label,
			"caller":      caller,
			"message":     fmt.Sprintf("%s", msg),
			"method":      method,
			"path":        path,
		})
		if err != nil {
			fmt.Println(errors.Wrap(err, "failed to encode to json"))
			return
		}
		l.GoLogger.Print(string(bytes))
	}
	if enableStdout {
		var (
			mt = reflect.TypeOf(msg)
			s  interface{}
		)
		if mt.Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			s = fmt.Sprintf("%+v", msg)
		} else {
			s = fmt.Sprintf("%+v", msg)
		}

		shortNow := now.Format("2006/01/02 15:04:05.000000")
		if lv.Label == "DUMP" {
			fmt.Println(strings.Join([]string{lv.TermBegin + "[" + lv.Label + "]" + lv.TermEnd, shortNow, userID, customerID, servicerID, requestID}, "\t"))
			fmt.Println(caller)
			fmt.Print(s)
			fmt.Println()
		} else if lv.Label == "TRACE" {
			fmt.Println(strings.Join([]string{lv.TermBegin + "[" + lv.Label + "]" + lv.TermEnd, shortNow, userID, customerID, servicerID, requestID, s.(string)}, "\t"))
		} else {
			fmt.Println(strings.Join([]string{lv.TermBegin + "[" + lv.Label + "]" + lv.TermEnd, shortNow, userID, customerID, servicerID, requestID, s.(string), caller}, "\t"))
		}
	}
}

func (l *Logger) reopenIfFileIsMissed() error {
	// for td-agent rotation
	var stat syscall.Stat_t
	err := syscall.Stat(l.Path, &stat)
	if err != nil && err.Error() != "no such file or directory" {
		return errors.Wrap(err, "failed to stat:"+l.Path)
	}
	if stat.Ino != l.LastIno {
		err = l.File.Close()
		if err != nil {
			return errors.Wrap(err, "failed to close:"+l.Path)
		}
		l.File, err = os.OpenFile(l.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return errors.Wrap(err, "failed to open:"+l.Path)
		}
		l.GoLogger.SetOutput(l.File)
		l.LastIno = stat.Ino
	}
	return nil
}

func Dump(ctx context.Context, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutDump == 0 && l.WriteFlag&FileDump == 0 {
		return
	}
	l.printf(ctx, &dumpLevel, l.WriteFlag&StdoutDump != 0, l.WriteFlag&FileDump != 0, spew.Sdump(v...))
}
func Debug(ctx context.Context, msg interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutDebug == 0 && l.WriteFlag&FileDebug == 0 {
		return
	}
	l.printf(ctx, &debugLevel, l.WriteFlag&StdoutDebug != 0, l.WriteFlag&FileDebug != 0, msg)
}
func Debugf(ctx context.Context, format string, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutDebug == 0 && l.WriteFlag&FileDebug == 0 {
		return
	}
	l.printf(ctx, &debugLevel, l.WriteFlag&StdoutDebug != 0, l.WriteFlag&FileDebug != 0, fmt.Sprintf(format, v...))
}
func Info(ctx context.Context, msg interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutInfo == 0 && l.WriteFlag&FileInfo == 0 {
		return
	}
	l.printf(ctx, &infoLevel, l.WriteFlag&StdoutInfo != 0, l.WriteFlag&FileInfo != 0, msg)
}
func Infof(ctx context.Context, format string, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutInfo == 0 && l.WriteFlag&FileInfo == 0 {
		return
	}
	l.printf(ctx, &infoLevel, l.WriteFlag&StdoutInfo != 0, l.WriteFlag&FileInfo != 0, fmt.Sprintf(format, v...))
}
func Warn(ctx context.Context, msg interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutWarn == 0 && l.WriteFlag&FileWarn == 0 {
		return
	}
	l.printf(ctx, &warnLevel, l.WriteFlag&StdoutWarn != 0, l.WriteFlag&FileWarn != 0, msg)
}
func Warnf(ctx context.Context, format string, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutWarn == 0 && l.WriteFlag&FileWarn == 0 {
		return
	}
	l.printf(ctx, &warnLevel, l.WriteFlag&StdoutWarn != 0, l.WriteFlag&FileWarn != 0, fmt.Sprintf(format, v...))
}
func Crit(ctx context.Context, msg interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutCrit == 0 && l.WriteFlag&FileCrit == 0 {
		return
	}
	l.printf(ctx, &critLevel, l.WriteFlag&StdoutCrit != 0, l.WriteFlag&FileCrit != 0, msg)
}
func Critf(ctx context.Context, format string, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutCrit == 0 && l.WriteFlag&FileCrit == 0 {
		return
	}
	l.printf(ctx, &critLevel, l.WriteFlag&StdoutCrit != 0, l.WriteFlag&FileCrit != 0, fmt.Sprintf(format, v...))
}
func Trace(ctx context.Context, msg interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutTrace == 0 && l.WriteFlag&FileTrace == 0 {
		return
	}
	l.printf(ctx, &critLevel, l.WriteFlag&StdoutTrace != 0, l.WriteFlag&FileTrace != 0, msg)
}
func Tracef(ctx context.Context, format string, v ...interface{}) {
	l := loggerContainer["default"]
	if l.WriteFlag&StdoutTrace == 0 && l.WriteFlag&FileTrace == 0 {
		return
	}
	l.printf(ctx, &traceLevel, l.WriteFlag&StdoutTrace != 0, l.WriteFlag&FileTrace != 0, fmt.Sprintf(format, v...))
}

func ADump(ctx context.Context, alias string, v ...interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutDump == 0 && l.WriteFlag&FileDump == 0 {
		return
	}
	l.printf(ctx, &dumpLevel, l.WriteFlag&StdoutDump != 0, l.WriteFlag&FileDump != 0, spew.Sdump(v...))
}
func ADebug(ctx context.Context, alias string, msg interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutDebug == 0 && l.WriteFlag&FileDebug == 0 {
		return
	}
	l.printf(ctx, &debugLevel, l.WriteFlag&StdoutDebug != 0, l.WriteFlag&FileDebug != 0, msg)
}
func ADebugf(ctx context.Context, alias string, format string, v ...interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutDebug == 0 && l.WriteFlag&FileDebug == 0 {
		return
	}
	l.printf(ctx, &debugLevel, l.WriteFlag&StdoutDebug != 0, l.WriteFlag&FileDebug != 0, fmt.Sprintf(format, v...))
}
func AInfo(ctx context.Context, alias string, msg interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutInfo == 0 && l.WriteFlag&FileInfo == 0 {
		return
	}
	l.printf(ctx, &infoLevel, l.WriteFlag&StdoutInfo != 0, l.WriteFlag&FileInfo != 0, msg)
}
func AInfof(ctx context.Context, alias string, format string, v ...interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutInfo == 0 && l.WriteFlag&FileInfo == 0 {
		return
	}
	l.printf(ctx, &infoLevel, l.WriteFlag&StdoutInfo != 0, l.WriteFlag&FileInfo != 0, fmt.Sprintf(format, v...))
}
func AWarn(ctx context.Context, alias string, msg interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutWarn == 0 && l.WriteFlag&FileWarn == 0 {
		return
	}
	l.printf(ctx, &warnLevel, l.WriteFlag&StdoutWarn != 0, l.WriteFlag&FileWarn != 0, msg)
}
func AWarnf(ctx context.Context, alias string, format string, v ...interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutWarn == 0 && l.WriteFlag&FileWarn == 0 {
		return
	}
	l.printf(ctx, &warnLevel, l.WriteFlag&StdoutWarn != 0, l.WriteFlag&FileWarn != 0, fmt.Sprintf(format, v...))
}
func ACrit(ctx context.Context, alias string, msg interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutCrit == 0 && l.WriteFlag&FileCrit == 0 {
		return
	}
	l.printf(ctx, &critLevel, l.WriteFlag&StdoutCrit != 0, l.WriteFlag&FileCrit != 0, msg)
}
func ACritf(ctx context.Context, alias string, format string, v ...interface{}) {
	l := loggerContainer[alias]
	if l.WriteFlag&StdoutCrit == 0 && l.WriteFlag&FileCrit == 0 {
		return
	}
	l.printf(ctx, &critLevel, l.WriteFlag&StdoutCrit != 0, l.WriteFlag&FileCrit != 0, fmt.Sprintf(format, v...))
}
