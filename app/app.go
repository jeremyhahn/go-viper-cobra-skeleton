package app

import (
	"fmt"
	"log"
	"os"

	logging "github.com/op/go-logging"
)

var Name string

type App struct {
	DebugFlag bool            `yaml:"debug" json:"debug" mapstructure:"debug"`
	DataDir   string          `yaml:"datadir" json:"datadir" mapstructure:"datadir"`
	HomeDir   string          `yaml:"homedir" json:"homedir" mapstructure:"homedir"`
	LogDir    string          `yaml:"logdir" json:"logdir" mapstructure:"logdir"`
	LogFile   string          `yaml:"logfile" json:"logfile" mapstructure:"logfile"`
	Logger    *logging.Logger `yaml:"-" json:"-" mapstructure:"-"`
}

func NewApp() *App {
	return new(App)
}

func (_app *App) InitLogFile(uid, gid int) *os.File {
	var logFile string
	if _app.LogFile == "" {
		logFile = fmt.Sprintf("%s/%s.log", _app.LogDir, Name)
	} else {
		logFile = _app.LogFile
	}
	var f *os.File
	var err error
	f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stat(logFile)
	if err != nil {
		if os.IsNotExist(err) {
			_, err2 := os.Create(logFile)
			if err2 != nil {
				log.Fatal(err2)
			}
		}
		log.Fatal(err)
	}
	if uid == 0 {
		err = os.Chown(logFile, uid, gid)
		if err != nil {
			log.Fatal(err)
		}
		if _app.DebugFlag {
			err = os.Chmod(logFile, 0777)
		} else {
			err = os.Chmod(logFile, 0644)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return f
}
