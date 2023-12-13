package env

import (
	"log"
	"os"
	"path/filepath"
)

var _appPath string

func init() {
	appPath, err := os.Getwd()
	if err != nil {
		log.Panicln(err)
	}
	_appPath = appPath
}

type ENV string

const (
	Local ENV = "local"
	Dev   ENV = "dev"
	Test  ENV = "test"
	Prod  ENV = "prod"
)

func GetEnv() ENV {
	env := os.Getenv("APP_ENV")
	switch ENV(env) {
	case Test:
		return Test
	case Prod:
		return Prod
	case Dev:
		return Dev
	default:
		return Local
	}
}

func IsProd() bool {
	return GetEnv() == Prod
}

func IsTest() bool {
	return GetEnv() == Test
}

func IsDev() bool {
	return GetEnv() == Dev
}

func IsLocal() bool {
	return GetEnv() == Local
}

func GetAppName() string {
	app := filepath.Base(os.Args[0])
	return app
}

func GetAppPath() string {
	return _appPath
}
