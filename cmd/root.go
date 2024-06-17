package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/jeremyhahn/go-viper-cobra-skeleton/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var App *app.App
var DebugFlag bool
var ConfigDir string
var DataDir string
var LogDir string
var HomeDir string

var rootCmd = &cobra.Command{
	Use:   app.Name,
	Short: "Golang Viper / Cobra trusted-platform project",
	Long: `
 A trusted-platform project to quickly get started building a Golang Viper / Cobra based
 command line interace that includes support for logging, versioning, and dynamic
 configuration.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// app.NewApp().Init()
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
	TraverseChildren: true,
}

func init() {

	wd, _ := os.Getwd()

	rootCmd.PersistentFlags().BoolVarP(&DebugFlag, "debug", "", false, "Enable debug mode")
	rootCmd.PersistentFlags().StringVarP(&HomeDir, "home", "", wd, "Program home directory") // doesnt work as system daemon if not wd (/)
	rootCmd.PersistentFlags().StringVarP(&DataDir, "data-dir", "", fmt.Sprintf("%s/db", wd), "Directory where database files are stored")
	rootCmd.PersistentFlags().StringVarP(&ConfigDir, "config-dir", "", fmt.Sprintf("/etc/%s", app.Name), "Directory where configuration files are stored")
	rootCmd.PersistentFlags().StringVarP(&LogDir, "log-dir", "", "./logs", "Logging directory")

	viper.BindPFlags(rootCmd.PersistentFlags())

	if runtime.GOOS == "darwin" {
		signal.Ignore(syscall.Signal(0xd))
	}

	cobra.OnInitialize(func() {
		app.NewApp().Init()
	})
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	return nil
}
