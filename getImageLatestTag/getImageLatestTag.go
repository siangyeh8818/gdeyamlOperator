package main

import (
	"flag"
	"fmt"
	"log"
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
var inputformat string

func main() {

	Init()
	flag.Parse()
	fmt.Println(completeimagename)
	fmt.Println(list)
	fmt.Println(inputfile)
	fmt.Println(inputformat)

	//判斷
	if inputfile != "" && Exists(inputfile) {
		inyaml := K8sYaml{}
		inyaml.getConf(inputfile)
		//fmt.Printf("input_YAML:\n%v\n\n", inyaml)

		//fmt.Println(ComposeImageName(inyaml.Deployment.K8S[0].Stage, inyaml.Deployment.K8S[0].Image, inyaml.Deployment.K8S[0].Tag))

		for i := 0; i < len(inyaml.Deployment.K8S); i++ {
			if inyaml.Deployment.K8S[i].Image != "" {
				fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)
				tmp_cpmplete_imagename := ComposeImageName(inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
				new_tag_latest := GetTag(tmp_cpmplete_imagename)
				(&inyaml.Deployment.K8S[i]).UpdateK8sTag(new_tag_latest)
				fmt.Printf("new_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)
			} else {
				continue
			}

		}
		for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
			if inyaml.Deployment.Openfaas[i].Image != "" {
				fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.Openfaas[i].Tag)
				tmp_cpmplete_imagename := ComposeImageName(inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
				new_tag_latest := GetTag(tmp_cpmplete_imagename)
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
		new_tag_latest := GetTag(completeimagename)
		fmt.Println(new_tag_latest)
		new_tag_latest = strings.Trim(new_tag_latest, "\"")
		WriteWithIoutil("getImageLatestTag_result.txt", new_tag_latest)
	}

}

func Init() {
	flag.StringVar(&completeimagename, "imagename", "dockerhub.pentium.network/grafana", "docker image , such as dockerhub.pentium.network/grafana")
	flag.IntVar(&list, "list", 5, "After sort tag list , we only deal with these top'number tags ")
	flag.StringVar(&inputfile, "inputfile", "", "input file name , such as deploy.yml")
	flag.StringVar(&ouputfile, "ouputfile", "tmp_out.yml", "output file name , such as deploy-out.yml")
	flag.StringVar(&hubSource, "hub", "dockerhub.pentium.network", "dockerhub url")
}

func ComposeImageName(stage string, module string, tag string) string {

	var complete_image string

	complete_image = hubSource + "/" + stage + "/" + module + ":" + tag

	return complete_image
}

func GetTag(name string) string {
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
