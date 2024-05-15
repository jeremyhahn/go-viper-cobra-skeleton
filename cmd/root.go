package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jeremyhahn/go-viper-cobra-skeleton/app"

	logging "github.com/op/go-logging"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logFormat = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortpkg}.%{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

var App *app.App
var DebugFlag bool
var ConfigDir string
var DataDir string
var LogDir string
var LogFile string
var HomeDir string

var rootCmd = &cobra.Command{
	Use:   app.Name,
	Short: "Golang Viper / Cobra skeleton project",
	Long: `
 A skeleton project to quickly get started building a Golang Viper / Cobra based
 command line interace that includes support for logging, versioning, and dynamic
 configuration.`,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//initApp()
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	TraverseChildren: true,
}

func init() {
	cobra.OnInitialize(initApp)

	wd, _ := os.Getwd()

	rootCmd.PersistentFlags().BoolVarP(&DebugFlag, "debug", "", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&HomeDir, "home", "", wd, "Program home directory") // doesnt work as system daemon if not wd (/)
	rootCmd.PersistentFlags().StringVarP(&DataDir, "data-dir", "", fmt.Sprintf("%s/db", wd), "Directory where database files are stored")
	rootCmd.PersistentFlags().StringVarP(&ConfigDir, "config-dir", "", fmt.Sprintf("/etc/%s", app.Name), "Directory where configuration files are stored")
	rootCmd.PersistentFlags().StringVarP(&LogDir, "log-dir", "", "/var/log", "Logging directory")
	rootCmd.PersistentFlags().StringVarP(&LogFile, "log-file", "", fmt.Sprintf("/var/log/%s.log", app.Name), "Application log file")

	viper.BindPFlags(rootCmd.PersistentFlags())

	if runtime.GOOS == "darwin" {
		signal.Ignore(syscall.Signal(0xd))
	}
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	return nil
}

func initApp() {
	App.DebugFlag = viper.GetBool("debug")
	App.HomeDir = viper.GetString("home")
	initLogger()
	initConfig()
	if App.DebugFlag {
		logging.SetLevel(logging.DEBUG, "")
		App.Logger.Debug("Starting logger in debug mode...")
		for k, v := range viper.AllSettings() {
			App.Logger.Debugf("%s: %+v", k, v)
		}
	} else {
		logging.SetLevel(logging.INFO, "")
	}
}

func initLogger() {
	App.LogDir = LogDir
	App.LogFile = LogFile
	f := App.InitLogFile(os.Getuid(), os.Getgid())
	stdout := logging.NewLogBackend(os.Stdout, "", 0)
	logfile := logging.NewLogBackend(f, "", log.Lshortfile)
	logFormatter := logging.NewBackendFormatter(logfile, logFormat)
	//syslog, _ := logging.NewSyslogBackend(appName)
	backends := logging.MultiLogger(stdout, logFormatter)
	logging.SetBackend(backends)
	if App.DebugFlag {
		logging.SetLevel(logging.DEBUG, "")
	} else {
		logging.SetLevel(logging.ERROR, "")
	}
	App.Logger = logging.MustGetLogger(app.Name)
}

func initConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(ConfigDir)
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", app.Name))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.%s/", app.Name))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		App.Logger.Errorf("%s", err)
	}

	viper.Unmarshal(&App)

	App.DataDir = viper.GetString("data-dir")

	App.Logger.Debugf("%+v", App)
}
