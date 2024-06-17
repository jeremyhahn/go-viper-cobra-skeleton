package app

import (
	"fmt"
	"log"
	"os"

	logging "github.com/op/go-logging"
	"github.com/spf13/viper"
)

type App struct {
	DebugFlag bool            `yaml:"debug" json:"debug" mapstructure:"debug"`
	ConfigDir string          `yaml:"config-dir" json:"config_dir" mapstructure:"config-dir"`
	DataDir   string          `yaml:"data-dir" json:"data_dir" mapstructure:"data-dir"`
	LogDir    string          `yaml:"log-dir" json:"log_dir" mapstructure:"log-dir"`
	Logger    *logging.Logger `yaml:"-" json:"-" mapstructure:"-"`
}

func NewApp() *App {
	return new(App)
}

func (app *App) Init() {
	app.initConfig()
	app.initLogger()
	if app.DebugFlag {
		logging.SetLevel(logging.DEBUG, "")
		app.Logger.Debug("Starting logger in debug mode...")
		for k, v := range viper.AllSettings() {
			app.Logger.Debugf("%s: %+v", k, v)
		}
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

func (app *App) initConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(app.ConfigDir)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", Name))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", Name))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&app); err != nil {
		log.Fatal(err)
	}

	if app.DebugFlag {
		log.Println(viper.AllSettings())
	}
}

func (app *App) initLogger() {
	logFormat := logging.MustStringFormatter(
		`%{color}%{time:15:04:05.000} %{shortpkg}.%{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
	)
	f := app.InitLogFile(os.Getuid(), os.Getgid())
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	logfile := logging.NewLogBackend(f, "", log.Lshortfile)
	logFormatter := logging.NewBackendFormatter(logfile, logFormat)
	//syslog, _ := logging.NewSyslogBackend(appName)
	backends := logging.MultiLogger(stdout, logFormatter)
	logging.SetBackend(backends)
	if app.DebugFlag {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.ERROR, "")
	}
	app.Logger = logging.MustGetLogger(Name)
}

func (app *App) InitLogFile(uid, gid int) *os.File {
	logFile := fmt.Sprintf("%s/%s.log", app.LogDir, Name)
	if err := os.MkdirAll(app.LogDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	var f *os.File
	var err error
	f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
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
		if err = os.Chown(logFile, uid, gid); err != nil {
			log.Fatal(err)
		}
		if app.DebugFlag {
			if err = os.Chmod(logFile, os.ModePerm); err != nil {
				log.Fatal(err)
			}
		} else {
			if err = os.Chmod(logFile, 0644); err != nil {
				log.Fatal(err)
			}
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	return f
}
