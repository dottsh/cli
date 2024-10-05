package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	syncCommand.Flags().StringSliceP("workspace", "w", []string{}, "the name of the workspace(s) to commit in")
	syncCommand.Flags().StringSliceP("tag", "t", []string{}, "the tags to filter the repositories by")
	root.AddCommand(syncCommand)
}
func isInDir(name string, files []os.DirEntry) bool {
	for _, file := range files {
		if file.IsDir() {
			log.Printf("reading dir: %v", filepath.Join(file.Name(), name))
			repoFiles, err := os.ReadDir(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			absPath, err := filepath.Abs(file.Name())
			if err != nil {
				log.Fatal(err)
			}
			return isInDir(filepath.Join(absPath, name), repoFiles)
		} else if file.Name() == name {
			return true
		}
	}
	return false
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
			if group.Repo != "" {
				log.Printf("syncing repo: %v", group.Repo)
				Clone(CloneArgs{
					URL:  group.Repo,
					Path: group.Dest,
				})
				if err != nil {
					log.Fatal(err)
				}
				for _, item := range group.Items {
					log.Printf("syncing item: %v", item.Name)
					repoFiles, err := os.ReadDir(group.Dest)
					if err != nil {
						log.Fatal(err)
					}
					if isInDir(group.Dest, repoFiles) {
						log.Printf("repo already exists: %v", group.Dest)
					}
					if item.Type == File {
						log.Printf("syncing file: %v", item.Name)
					} else if item.Type == Brew {
						log.Printf("syncing brew: %v", item.Name)
					}
				}
			}
		}
	},
}
