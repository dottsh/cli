package main

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mateothegreat/go-multilog/multilog"
	"github.com/mateothegreat/go-util/files"
	"github.com/mateothegreat/go-util/urls"
	"github.com/polyrepopro/api/config"
)

func GetAuth(url string, auth *config.Auth) transport.AuthMethod {
	if auth == nil && urls.GetProtocol(url) == "ssh" {
		// Check if the default SSH key exists
		defaultSSHKey := files.ExpandPath("~/.ssh/id_rsa")
		if _, err := os.Stat(defaultSSHKey); err == nil {
			// Default SSH key exists, use it
			sshAuth, err := ssh.NewPublicKeysFromFile("git", defaultSSHKey, "")
			if err != nil {
				multilog.Fatal("git.clone", "failed to create SSH auth with default key", map[string]interface{}{
					"error": err,
				})
				return nil
			}
			return sshAuth
		} else {
			// Default SSH key doesn't exist, proceed without auth
			multilog.Warn("git.clone", "no auth provided and default SSH key not found", map[string]interface{}{
				"path": defaultSSHKey,
			})
		}
	} else if auth != nil && auth.Key != "" {
		// Use SSH key
		sshAuth, err := ssh.NewPublicKeysFromFile("git", auth.Key, "")
		if err != nil {
			multilog.Fatal("git.clone", "failed to create SSH auth with provided key", map[string]interface{}{
				"error": err,
			})
			return nil
		}
		return sshAuth
	} else if auth != nil && auth.Env.Username != "" && auth.Env.Password != "" {
		// Use HTTP auth
		return &http.BasicAuth{
			Username: os.Getenv(auth.Env.Username),
			Password: os.Getenv(auth.Env.Password),
		}
	}
	return nil
}

type CloneArgs struct {
	URL  string
	Path string
	Auth *config.Auth
}

type progress struct{}

func (h *progress) Write(p []byte) (n int, err error) {
	multilog.Info("git.clone", "cloning progress", map[string]interface{}{
		"message": string(p),
	})
	return len(p), nil
}

func Clone(args CloneArgs) error {
	var err error

	multilog.Info("git.clone", "cloning repository", map[string]interface{}{
		"url":  args.URL,
		"path": args.Path,
	})

	opts := &git.CloneOptions{
		URL:               args.URL,
		Progress:          &progress{},
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	}

	auth := GetAuth(args.URL, args.Auth)
	if auth.Name() != "" {
		opts.Auth = auth
	}

	_, err = git.PlainClone(args.Path, false, opts)
	if err != nil {
		multilog.Fatal("git.clone", "failed to clone repository", map[string]interface{}{
			"error": err,
		})
		return err
	}

	return nil
}
