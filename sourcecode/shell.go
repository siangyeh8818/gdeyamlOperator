package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func exec_shell(s_command string) (string, string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", s_command)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		fmt.Printf("stdout: %v, stderr: %v\n", errStdout, errStderr)
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	return outStr, errStr
}
func RunCommand(commandStr string) string {
	cmdstr := commandStr
	out, _ := exec.Command("sh", "-c", cmdstr).Output()
	strout := string(out)
	return strout
}

func KubectlGetDeployment(namespace string) []string {

	cmd := "kubectl get deploy -n " + namespace + "| awk '{print $1}'"
	result, _ := exec_shell(cmd)
	//fmt.Println(result)
	totaldeploy := strings.Split(result, "\n")
	return totaldeploy
}

func KubectlGetStefulset(namespace string) []string {
	cmd := "kubectl get statefulset -n " + namespace + "| awk '{print $1}'"
	result, _ := exec_shell(cmd)
	//fmt.Println(result)
	totalststefulset := strings.Split(result, "\n")
	return totalststefulset
}

func KubectlGetDaemonset(namespace string) []string {
	cmd := "kubectl get daemonset -n " + namespace + "| awk '{print $1}'"
	result, _ := exec_shell(cmd)
	//fmt.Println(result)
	totaldaemonset := strings.Split(result, "\n")
	return totaldaemonset
}

func KubectlGetCronJob(namespace string) []string {
	cmd := "kubectl get cronJob -n " + namespace + "| awk '{print $1}'"
	result, _ := exec_shell(cmd)
	//fmt.Println(result)
	totaldaemonset := strings.Split(result, "\n")
	return totaldaemonset
}
