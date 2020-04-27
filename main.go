package main

import (
	"flag"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shiji-naoki/media-top-backend/config"
	log "github.com/shiji-naoki/media-top-backend/infrastructure/logger"
	"github.com/shiji-naoki/media-top-backend/interfaces/server"
)

var (
	env   string
	mode  string
	name  string
	debug bool
)

func main() {
	flag.StringVar(&env, "env", "local", "application runtime environment")
	flag.StringVar(&mode, "mode", "server", "application runtime mode [server|worker]")
	flag.StringVar(&name, "name", "*", "execution worker name")
	flag.BoolVar(&debug, "debug", false, "logger debug option")
	flag.Parse()

	if err := config.SetEnv(env); err != nil {
		panic(errors.New(err.Error()))
	}

	if err := config.Loads(); err != nil {
		panic(errors.New(err.Error()))
	}

	if err := setupLogger(); err != nil {
		panic(err)
	}

	// ginのインスタンスを取得
	router := gin.Default()
	// ルーティングを設定
	server.V1(router)
	// サーバー実行
	router.Run()
}

func setupLogger() error {
	lc := config.GetEnvValue().Logger
	spew.Dump(lc)
	var configMapping = map[int]bool{
		log.FileCrit:    lc.File.Crit,
		log.FileWarn:    lc.File.Warn,
		log.FileInfo:    lc.File.Info,
		log.FileDebug:   lc.File.Debug,
		log.FileDump:    lc.File.Dump,
		log.StdoutCrit:  lc.Stdout.Crit,
		log.StdoutWarn:  lc.Stdout.Warn,
		log.StdoutInfo:  lc.Stdout.Info,
		log.StdoutDebug: lc.Stdout.Debug,
		log.StdoutDump:  lc.Stdout.Dump,
	}
	var lFlag int
	for f, lcFlag := range configMapping {
		if lcFlag {
			lFlag = lFlag | f
		}
	}
	for t, p := range lc.File.Path {
		err := log.Add(t, p, lFlag)
		if err != nil {
			return errors.Wrap(err, "failed to create log instance")
		}
	}
	return nil
}
