package main

import (
	"github.com/spf13/cobra"
	"log"
)

func init() {
	syncCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to commit in")
	syncCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	root.AddCommand(syncCommand)
}

var syncCommand = &cobra.Command{
	Use:   "sync",
	Short: "sync changes",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("syncing changes...")
		config, err := GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		for _, group := range config.Packages {
			log.Printf("syncing group: %v", group.Group)
		}
	},
}
