package gdeyamloperator

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	//	"k8s.io/apimachinery/pkg/api/errors"
)

func exec_shell(s_command string) (string, string) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", s_command)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	var errStdout, errStderr error
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	/*
		go func() {
			_, errStdout = io.Copy(stdout, stdoutIn)
		}()
		go func() {
			_, errStderr = io.Copy(stderr, stderrIn)
		}()
	*/
	_, errStdout = io.Copy(stdout, stdoutIn)
	_, errStderr = io.Copy(stderr, stderrIn)
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

func grepFolderName(module string, base_path string, ModuleMap map[string]int) string {
	var token int
	var current_module_pattern string
	current_module_pattern = "/" + module + ":"
	cmd := "grep -Rn " + current_module_pattern + " " + base_path + " | grep image | awk '{print $1}'"
	result, err := exec_shell(cmd)
	log.Println(result)
	if err != "" {
		log.Println("Find image base-folder failed")
	}
	result_slice := strings.Split(result, "/")

	for x := 0; x < len(result_slice); x++ {
		//log.Println(result_slice[x])
		if strings.Contains(result_slice[x], ".yml") || strings.Contains(result_slice[x], ".yaml") {
			//existtoken := Contain(result_slice[x-1], ModuleMap)
			existtoken := ModuleMap[result_slice[x-1]]
			if existtoken == 1 {
				log.Println("pn-base module can't repeat")
			} else if existtoken == 0 {
				if strings.Contains(result_slice[x-1], "gde") {
					log.Printf("pass folder %s , because it is the gde-folder\n ", result_slice[x-1])
				} else {
					token = x - 1
				}
				//token = x - 1
			}
		}
	}
	return result_slice[token]
}

func getBaseModuleNamespace(path string, doc string)string{

    f := path + "/"+doc
	namespaceFileContent, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Contents of file:", string(namespaceFileContent))
	return string(namespaceFileContent)
}

/*
func Contain(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}
*/
