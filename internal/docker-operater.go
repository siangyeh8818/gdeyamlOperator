package gdeyamloperator

import (
	"fmt"
	"strings"
)

type Docker struct {
	DockerLogin string
	Push bool
	PushPattern string
	PullPattern string
	Imagename string
	List int
	LatestMode string
	Stage string
}

func LoginDockerHub(stage string, user string, password string) {
	var login_cmd string
	login_cmd = "docker login cr-" + stage + ".pentium.network -u " + user + " -p " + password
	RunCommand(login_cmd)
}
func LoginDockerHubNew(url string, user string, password string) {
	var login_cmd string
	login_cmd = "docker login " + url + " -u " + user + " -p " + password
	RunCommand(login_cmd)
}

func PushTagimage(imagename string, push_pattern string, modulename string, moduletag string) {
	cmd_1 := "docker pull " + imagename
	fmt.Println(cmd_1)
	_, err := exec_shell(cmd_1)
	if err != "" {
		fmt.Println(err)
	}
	push_cpmplete_imagename := PatternParse(push_pattern, "preview", modulename, moduletag)
	cmd_2 := "docker tag " + imagename + " " + push_cpmplete_imagename
	fmt.Println(cmd_2)
	_, err = exec_shell(cmd_2)
	if err != "" {
		fmt.Println(err)
	}
	cmd_3 := "docker push " + push_cpmplete_imagename
	fmt.Println(cmd_3)
	_, err = exec_shell(cmd_3)
	if err != "" {
		fmt.Println(err)
	}
	cmd_4 := "docker rmi " + imagename
	fmt.Println(cmd_4)
	_, err = exec_shell(cmd_4)
	if err != "" {
		fmt.Println(err)
	}

	if imagename == push_cpmplete_imagename {
		fmt.Println("we don't need to docker rmi this image")
	} else if imagename != push_cpmplete_imagename {
		cmd_5 := "docker rmi " + push_cpmplete_imagename
		fmt.Println(cmd_5)
		_, err = exec_shell(cmd_5)
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
