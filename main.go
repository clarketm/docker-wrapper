package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func main() {
	_, callPath, _, _ := runtime.Caller(0)
	scriptPath := fmt.Sprintf("%s/%s-%s/docker-wrapper", filepath.Dir(callPath), runtime.GOOS, runtime.GOARCH)
	cmd := exec.Command(scriptPath, os.Args[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	cmd.Start()
	ch := make(chan string, 100)
	stdoutScan := bufio.NewScanner(stdout)
	stderrScan := bufio.NewScanner(stderr)

	go func() {
		for stdoutScan.Scan() {
			line := stdoutScan.Text()
			ch <- line
		}
	}()
	go func() {
		for stderrScan.Scan() {
			line := stderrScan.Text()
			ch <- line
		}
	}()
	go func() {
		cmd.Wait()
		close(ch)
	}()
	for line := range ch {
		fmt.Println(line)
	}
}
