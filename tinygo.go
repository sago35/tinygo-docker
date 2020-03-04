package main

import (
	"fmt"
	"os/exec"

	"github.com/mattn/go-tty"
)

func runTinyGo(dockerImage, currentDir, targetPath string, args []string, verbose, cmdMode bool) error {
	cmd := exec.Command(
		`docker`, `run`, `-it`, `--rm`,
		`-v`, fmt.Sprintf(`%s:/go/%s`, currentDir, targetPath),
		`-w`, fmt.Sprintf(`/go/%s`, targetPath),
		`-e`, `GOPATH=/go`,
		dockerImage,
		`tinygo`)
	if cmdMode {
		cmd = exec.Command(
			`docker`, `run`, `-it`, `--rm`,
			`-v`, fmt.Sprintf(`%s:/go/%s`, currentDir, targetPath),
			`-w`, fmt.Sprintf(`/go/%s`, targetPath),
			`-e`, `GOPATH=/go`,
			dockerImage,
		)
	}
	cmd.Args = append(cmd.Args, args...)
	if verbose {
		fmt.Println(cmd)
	}

	tty, err := tty.Open()
	if err != nil {
		return err
	}
	defer tty.Close()

	cmd.Stdin = tty.Input()
	cmd.Stdout = tty.Output()
	cmd.Stderr = tty.Output()
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
