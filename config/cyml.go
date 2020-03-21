package config

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets387c087837f1b2d86297f1ebaf919c4fe180845a = "database:\n  name: smartmat\n  host: 127.0.0.1\n  port: 3306 # default port for database\n  user: \"smashop\"\n  password: \"smashop\"\n  maxidleconnections: 10\n  maxopenconnections: 20\n  logmode: true\nserver:\n  endpoint: \"http://localhost\"\n  host: localhost\n  port: 5000\n "

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/env": []string{"local.yml"}, "/": []string{"env"}}, map[string]*assets.File{
	"/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1570765142, 1570765142646346058),
		Data:     nil,
	}, "/env": &assets.File{
		Path:     "/env",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1570774302, 1570774302695808338),
		Data:     nil,
	}, "/env/local.yml": &assets.File{
		Path:     "/env/local.yml",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1570774409, 1570774409911781875),
		Data:     []byte(_Assets387c087837f1b2d86297f1ebaf919c4fe180845a),
	}}, "")
