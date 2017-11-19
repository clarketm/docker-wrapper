package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	script, _ := exec.Command("command", "-v", "docker-wrapper").Output()
	scriptPath := string(script[:len(script)-1])
	// scriptPath := "./docker-wrapper"
	// scriptPath := strings.Trim(string(script), "\n")
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
