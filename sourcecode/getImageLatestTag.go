package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

var hubSource string
var completeimagename string
var list int
var inputfile string
var ouputfile string
var pushimage bool
var inputstage string
var query_mode string
var fqdn string
var new_hub string
var loginuser string
var loginpassword string
var version bool
var latest_mode string
var push_pattern string
var pull_pattern string

func main() {

	Init()
	flag.Parse()

	if version {
		fmt.Println("version : 1.3.2")
		os.Exit(0)
	}

	fmt.Printf("flag  imagename: %s\n", completeimagename)
	fmt.Printf("flag  user: %s\n", loginuser)
	fmt.Printf("flag  password: %s\n", loginpassword)
	fmt.Printf("flag  list: %d\n", list)
	fmt.Printf("flag  inputfile: %s\n", inputfile)
	fmt.Printf("flag  ouputfile: %s\n", ouputfile)
	//fmt.Printf("flag  hubSource: %s\n", hubSource)
	//fmt.Printf("flag  fqdn: %s\n", fqdn)
	//fmt.Printf("flag  query-mode: %s\n", query_mode)
	fmt.Printf("flag  stage: %s\n", inputstage)
	fmt.Printf("flag  push: %t\n", pushimage)
	fmt.Printf("flag  version: %t\n", version)
	fmt.Printf("flag latest-mode: %s\n", latest_mode)
	fmt.Printf("flag push-pattern: %s\n", push_pattern)
	fmt.Printf("flag pull-pattern: %s\n", pull_pattern)

	/*
		if query_mode == "fqdn" {
			new_hub = hubSource + fqdn
		} else if query_mode == "nexus" {
			new_hub = hubSource + inputstage + fqdn
		}
		fmt.Printf("complete pull/query hub-source url: %s\n", new_hub)

	*/

	if loginuser != "" && loginpassword != "" {
		LoginDockerHub(inputstage, loginuser, loginpassword)
	}

	//判斷
	if inputfile != "" && Exists(inputfile) {
		inyaml := K8sYaml{}
		inyaml.getConf(inputfile)
		//fmt.Printf("input_YAML:\n%v\n\n", inyaml)
		//fmt.Println(ComposeImageName(inyaml.Deployment.K8S[0].Stage, inyaml.Deployment.K8S[0].Image, inyaml.Deployment.K8S[0].Tag))
		for i := 0; i < len(inyaml.Deployment.K8S); i++ {
			if inyaml.Deployment.K8S[i].Image != "" {
				fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)
				var tmp_cpmplete_imagename string
				if inputstage != "" {
					tmp_cpmplete_imagename = PatternParse(pull_pattern, inputstage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
				} else if inputstage == "" {
					tmp_cpmplete_imagename = PatternParse(pull_pattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
				}

				//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
				fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
				if pushimage == true {
					cmd_1 := "docker pull " + tmp_cpmplete_imagename
					fmt.Println(cmd_1)
					RunCommand(cmd_1)
					push_cpmplete_imagename := PatternParse(push_pattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					//push_cpmplete_imagename := ComposeImageName(push_mode, new_push_hub, inputstage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					cmd_2 := "docker tag " + tmp_cpmplete_imagename + " " + push_cpmplete_imagename
					fmt.Println(cmd_2)
					RunCommand(cmd_2)
					cmd_3 := "docker push " + push_cpmplete_imagename
					fmt.Println(cmd_3)
					RunCommand(cmd_3)
					cmd_4 := "docker rmi " + tmp_cpmplete_imagename
					fmt.Println(cmd_4)
					RunCommand(cmd_4)
					cmd_5 := "docker rmi " + push_cpmplete_imagename
					fmt.Println(cmd_5)
					RunCommand(cmd_5)

				}
				new_tag_latest := GetTag(tmp_cpmplete_imagename, latest_mode)
				new_tag_latest = strings.Trim(new_tag_latest, "\"")
				(&inyaml.Deployment.K8S[i]).UpdateK8sTag(new_tag_latest)
				fmt.Printf("new_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)

			} else {
				continue
			}

		}
		for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
			if inyaml.Deployment.Openfaas[i].Image != "" {
				fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.Openfaas[i].Tag)
				var tmp_cpmplete_imagename string

				if inputstage != "" {
					tmp_cpmplete_imagename = PatternParse(pull_pattern, inputstage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
				} else if inputstage == "" {
					tmp_cpmplete_imagename = PatternParse(pull_pattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
				}
				//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
				fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
				if pushimage == true {
					cmd_1 := "docker pull " + tmp_cpmplete_imagename
					fmt.Println(cmd_1)
					RunCommand(cmd_1)
					push_cpmplete_imagename := PatternParse(push_pattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					cmd_2 := "docker tag " + tmp_cpmplete_imagename + " " + push_cpmplete_imagename
					fmt.Println(cmd_2)
					RunCommand(cmd_2)
					cmd_3 := "docker push " + push_cpmplete_imagename
					fmt.Println(cmd_3)
					RunCommand(cmd_3)
					cmd_4 := "docker rmi " + tmp_cpmplete_imagename
					fmt.Println(cmd_4)
					RunCommand(cmd_4)
					cmd_5 := "docker rmi " + push_cpmplete_imagename
					fmt.Println(cmd_5)
					RunCommand(cmd_5)

				}
				new_tag_latest := GetTag(tmp_cpmplete_imagename, latest_mode)
				new_tag_latest = strings.Trim(new_tag_latest, "\"")
				(&inyaml.Deployment.Openfaas[i]).UpdateOpenfaasTag(new_tag_latest)
				fmt.Printf("new_tag:\n%v\n\n", inyaml.Deployment.Openfaas[i].Tag)

			} else {
				continue
			}
		}
		d, err := yaml.Marshal(&inyaml)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		//	fmt.Printf("--- t dump:\n%s\n\n", string(d))

		WriteWithIoutil(ouputfile, string(d))

	} else {
		new_tag_latest := GetTag(completeimagename, latest_mode)
		fmt.Println(new_tag_latest)
		new_tag_latest = strings.Trim(new_tag_latest, "\"")
		WriteWithIoutil("getImageLatestTag_result.txt", new_tag_latest)
	}

}

func Init() {
	flag.StringVar(&completeimagename, "imagename", "dockerhub.pentium.network/grafana", "docker image , such as dockerhub.pentium.network/grafana")
	flag.StringVar(&loginuser, "user", "", "user for docker login")
	flag.StringVar(&loginpassword, "password", "", "password for docker login")
	flag.IntVar(&list, "list", 5, "After sort tag list , we only deal with these top'number tags ")
	flag.StringVar(&inputfile, "inputfile", "", "input file name , such as deploy.yml")
	flag.StringVar(&ouputfile, "ouputfile", "tmp_out.yml", "output file name , such as deploy-out.yml")
	//flag.StringVar(&hubSource, "hub", "dockerhub", "dockerhub url")
	//flag.StringVar(&fqdn, "fqdn", ".pentium.network", "setting FQDN")
	//flag.StringVar(&query_mode, "query-mode", "fqdn", "how to conpose hubname , you can choose 'fqdn' or 'nexus'")
	flag.StringVar(&inputstage, "stage", "", "replace stage , new stage content")
	flag.StringVar(&latest_mode, "latest-mode", "push", "push or build , choose one mode to identify latest tag to you")
	flag.StringVar(&push_pattern, "push-pattern", "", "(push)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.StringVar(&pull_pattern, "pull-pattern", "", "(pull)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.BoolVar(&pushimage, "push", false, "push this image , default is false")
	flag.BoolVar(&version, "v", false, "prints current binary version")

}

/*
func ComposeImageName(stage string, module string, tag string) string {

	var complete_image string

	complete_image = hubSource + "/" + stage + "/" + module + ":" + tag

	return complete_image
}
*/

func GetTag(name string, latestmode string) string {
	raw_image_hub, raw_image_name := ImagenameSplit(name)

	var tag_result string
	var time_latest = "2000-01-01T00:00:00.508640172Z"
	var tag_latest string
	var querylistcmd string
	var loop_break_count int

	querylistcmd = "curl -X GET https://" + raw_image_hub + "/v2/" + raw_image_name + "/tags/list -s| jq -r .tags"
	//fmt.Println(querylistcmd)
	//	fmt.Println("------------------")
	//tag_result, _ = exec_shell("curl -X GET https://dockerhub.pentium.network/v2/grafana/tags/list| jq -r .tags")

	tag_result = RunCommand(querylistcmd)
	tag_result = strings.Replace(tag_result, "[", "", 1)
	tag_result = strings.Replace(tag_result, "]", "", 1)
	tag_result = DeleteExtraSpace(tag_result)
	tag_result = strings.Replace(tag_result, "\n", "", -1)
	tagssplit := strings.Split(tag_result, ",")

	//fmt.Printf("Ints %v\n", tagssplit)
	reverse_tagssplit := reverseInts(tagssplit)
	//fmt.Printf("Reversed: %v\n", reverse_tagssplit)
	//	fmt.Println("Amount of image tag : " + strconv.Itoa(len(tagssplit)))
	imagemap := make(map[string]string, len(reverse_tagssplit))
	if latestmode == "build" {
		for i := range reverse_tagssplit {

			time := QueryLatestTag(reverse_tagssplit[i], raw_image_name, raw_image_hub)
			fmt.Println(reverse_tagssplit[i] + ":" + time)
			time = strings.Replace(time, "\n", "", -1)

			if strings.Compare(strings.Trim(reverse_tagssplit[i], "\""), "latest") == -1 {
				imagemap[reverse_tagssplit[i]] = time
				time_latest = SelectLatestTime(time, time_latest)
				if time_latest == imagemap[reverse_tagssplit[i]] {
					tag_latest = reverse_tagssplit[i]
				}
			}
			loop_break_count++
			if loop_break_count >= list {
				break
			}
		}
	} else if latestmode == "push" {
		tag_latest = reverse_tagssplit[0]
	}

	return tag_latest
}

func SelectLatestTime(t1 string, t2 string) string {
	var earlytime string
	time1, _ := time.Parse(time.RFC3339Nano, t1)
	//fmt.Println(time1)
	time2, _ := time.Parse(time.RFC3339Nano, t2)
	//fmt.Println(time2)
	if time2.After(time1) {
		earlytime = t2
		//	fmt.Println("time2 is winner")
	} else if time1.After(time2) {
		earlytime = t1
		//	fmt.Println("time1 is winner")
	}
	//fmt.Println("function SelectLatestTime result : " + earlytime)

	return earlytime
}

func DeleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "  ", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}

func QueryLatestTag(tag string, imgname string, hub string) string {

	curltagresult := RunCommand("curl -X GET https://" + hub + "/v2/" + imgname + "/manifests/" + tag + " | jq -r '.history[].v1Compatibility' | jq '.created' | sort | sed 's/\"//g'|tail -n1 ")

	return curltagresult
}

func reverseInts(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func PatternParse(patterns string, structstage string, structimage string, structtag string) string {

	patterns = strings.Replace(patterns, "{{stage}}", structstage, 1)
	patterns = strings.Replace(patterns, "{{image}}", structimage, 1)
	patterns = strings.Replace(patterns, "{{tag}}", structtag, 1)
	return patterns
}