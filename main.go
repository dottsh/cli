package main

import (
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "dott",
	Short: "Replicate environment files, packages, and more ♻️.",
}

func main() {
	root.PersistentFlags().StringP("config", "c", "", "the path to the polyrepo config file")
	root.PersistentFlags().BoolP("verbose", "v", false, "output detailed logs")

	logLevel := multilog.INFO

	if verbose, _ := root.PersistentFlags().GetBool("verbose"); verbose {
		logLevel = multilog.DEBUG
	}

	multilog.RegisterLogger(multilog.LogMethod("console"), multilog.NewConsoleLogger(&multilog.NewConsoleLoggerArgs{
		Level:  logLevel,
		Format: multilog.FormatText,
	}))

	root.CompletionOptions.DisableDefaultCmd = true
	root.Execute()
}
