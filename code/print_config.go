package main

import (
    "fmt"
    "os/exec"
	"bytes"
)

func ExecCmd(cmdName string, cmdStdIn string, args ...string) (string, error) {
	cmd := exec.Command(cmdName, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if cmdStdIn != "" {
		buffer := bytes.Buffer{}
		buffer.Write([]byte(cmdStdIn))
		cmd.Stdin = &buffer
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if len(errStr) > 1 {
		fmt.Println(outStr)
		fmt.Println(errStr)
	}
	return outStr, err
}


func main() {
    output, _ := ExecCmd("git", "", "config", "--list")
    println(output)
}
