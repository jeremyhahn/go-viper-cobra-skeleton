package main

import (
	"github.com/jeremyhahn/go-viper-cobra-skeleton/app"
	"github.com/jeremyhahn/go-viper-cobra-skeleton/cmd"
)

func main() {
	cmd.App = app.NewApp()
	cmd.Execute()
}
