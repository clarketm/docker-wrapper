// +build linux darwin

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"syscall"
)

func main() {
	_, callPath, _, _ := runtime.Caller(0)
	binary := fmt.Sprintf("%s/%s-%s/docker-wrapper", filepath.Dir(callPath), runtime.GOOS, runtime.GOARCH)
	env := os.Environ()
	syscall.Exec(binary, os.Args, env)
}
