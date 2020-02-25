package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	appName        = "tinygo-docker"
	appDescription = ""
)

type cli struct {
	outStream io.Writer
	errStream io.Writer
}

func (c *cli) Run(args []string) error {
	if len(os.Args) >= 2 {
		switch os.Args[1] {
		case `--version`:
			fmt.Fprintf(c.outStream, "%s version %s build %s", appName, VERSION, BUILDDATE)
			return nil
		}
	}

	dockerImage := `tinygo/tinygo`
	if len(os.Args) >= 3 {
		if os.Args[1] == `--docker-image` {
			dockerImage = os.Args[2]
			os.Args = append([]string{os.Args[0]}, os.Args[3:]...)
		}
	}

	k := ""
	switch k {
	default:
		currentDir, err := os.Getwd()
		if err != nil {
			return err
		}
		//targetPath := `github.com/sago35/pyportal-private/touch`
		gopath, err := getGopath()
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(gopath, currentDir)
		if err != nil {
			return err
		}
		targetPath := filepath.ToSlash(rel)

		currentDirForDocker, err := modDir(currentDir)
		if err != nil {
			return err
		}

		//args := []string{`build`, `-o`, `app.uf2`, `-target`, `pyportal`, `.`}
		err = runTinyGo(dockerImage, currentDirForDocker, targetPath, os.Args[1:])
		if err != nil {
			return err
		}
	}

	return nil
}

func getGopath() (string, error) {
	out, err := exec.Command(`go`, `env`, `GOPATH`).Output()
	if err != nil {
		return "", err
	}

	gopath := strings.TrimRight(string(out), "\r\n")
	return gopath, nil
}
