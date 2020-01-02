package docker

import (
	"fmt"
	"strings"

	ShellCommand "github.com/siangyeh8818/gdeyamlOperator/internal/shellcommand"
)

type Docker struct {
	DockerLogin string
	Push        bool
	PushPattern string
	PullPattern string
	Imagename   string
	List        int
	LatestMode  string
	Stage       string
}

func LoginDockerHub(stage string, user string, password string) {
	var login_cmd string
	login_cmd = "docker login cr-" + stage + ".pentium.network -u " + user + " -p " + password
	ShellCommand.RunCommand(login_cmd)
}
func LoginDockerHubNew(url string, user string, password string) {
	var login_cmd string
	login_cmd = "docker login " + url + " -u " + user + " -p " + password
	ShellCommand.RunCommand(login_cmd)
}

func PushTagimage(imagename string, push_pattern string, modulename string, moduletag string) {
	cmd_1 := "docker pull " + imagename
	fmt.Println(cmd_1)
	_, err := ShellCommand.ExecShell(cmd_1)
	if err != "" {
		fmt.Println(err)
	}
	push_cpmplete_imagename := PatternParse(push_pattern, "preview", modulename, moduletag)
	cmd_2 := "docker tag " + imagename + " " + push_cpmplete_imagename
	fmt.Println(cmd_2)
	_, err = ShellCommand.ExecShell(cmd_2)
	if err != "" {
		fmt.Println(err)
	}
	cmd_3 := "docker push " + push_cpmplete_imagename
	fmt.Println(cmd_3)
	_, err = ShellCommand.ExecShell(cmd_3)
	if err != "" {
		fmt.Println(err)
	}
	cmd_4 := "docker rmi " + imagename
	fmt.Println(cmd_4)
	_, err = ShellCommand.ExecShell(cmd_4)
	if err != "" {
		fmt.Println(err)
	}

	if imagename == push_cpmplete_imagename {
		fmt.Println("we don't need to docker rmi this image")
	} else if imagename != push_cpmplete_imagename {
		cmd_5 := "docker rmi " + push_cpmplete_imagename
		fmt.Println(cmd_5)
		_, err = ShellCommand.ExecShell(cmd_5)
		if err != "" {
			fmt.Println(err)
		}
	}

}

func PatternParse(patterns string, structstage string, structimage string, structtag string) string {

	patterns = strings.Replace(patterns, "{{stage}}", structstage, 1)
	patterns = strings.Replace(patterns, "{{image}}", structimage, 1)
	patterns = strings.Replace(patterns, "{{tag}}", structtag, 1)
	return patterns
}

func ComposeImageName(mode string, hubdomain string, stage string, module string, tag string) string {

	var complete_image string
	if mode == "fqdn" {
		complete_image = hubdomain + "/" + stage + "/" + module + ":" + tag
	} else if mode == "nexus" {
		complete_image = hubdomain + "/" + module + ":" + tag
	}
	return complete_image

}
