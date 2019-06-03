package main

import (
	"flag"
	"fmt"
	"regexp"
	//	"strconv"
	"strings"
	"time"
)

var hubSource string

func main() {

	Init()
	flag.Parse()
	fmt.Println(hubSource)
	raw_image_hub, raw_image_name := ImagenameSplit(hubSource)
	/*
		fmt.Println("------------------")
		fmt.Println(raw_image_hub)
		fmt.Println("------------------")
		fmt.Println(raw_image_name)
		fmt.Println("------------------")
	*/
	var tag_result string
	var time_latest = "2000-01-01T00:00:00.508640172Z"
	var tag_latest string
	var querylistcmd string

	querylistcmd = "curl -X GET https://" + raw_image_hub + "/v2/" + raw_image_name + "/tags/list -s| jq -r .tags"
	//fmt.Println(querylistcmd)
	//	fmt.Println("------------------")
	//tag_result, _ = exec_shell("curl -X GET https://dockerhub.pentium.network/v2/grafana/tags/list| jq -r .tags")
	tag_result, _ = exec_shell(querylistcmd)

	tag_result = strings.Replace(tag_result, "[", "", 1)
	tag_result = strings.Replace(tag_result, "]", "", 1)
	//fmt.Println(tag_result)

	tag_result = DeleteExtraSpace(tag_result)
	//fmt.Println(tag_result)
	tag_result = strings.Replace(tag_result, "\n", "", -1)
	//fmt.Println(tag_result)
	tagssplit := strings.Split(tag_result, ",")

	//	fmt.Println("Amount of image tag : " + strconv.Itoa(len(tagssplit)))
	imagemap := make(map[string]string, len(tagssplit))
	for i := range tagssplit {
		time := QueryLatestTag(tagssplit[i], raw_image_name, raw_image_hub)
		//		fmt.Println(tagssplit[i] + " :  " + time)
		time = strings.Replace(time, "\n", "", -1)
		imagemap[tagssplit[i]] = time
		time_latest = SelectLatestTime(time, time_latest)
		if time_latest == imagemap[tagssplit[i]] {
			tag_latest = tagssplit[i]
		}
	}
	//test := SelectLatestTime("2019-05-16T02:07:18.508640172Z", "2019-04-22T07:47:39.89748501Z")
	fmt.Println(tag_latest)
	tag_latest = strings.Trim(tag_latest, "\"")
	WriteWithIoutil("getImageLatestTag_result.txt", tag_latest)
	//fmt.Println(time_latest)
}

func Init() {
	flag.StringVar(&hubSource, "imagename", "dockerhub.pentium.network/grafana", "docker image , such as dockerhub.pentium.network/grafana")
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
	//curltagresult, _ := exec_shell("curl -X GET https://" + hub + "/v2/" + imgname + "/manifests/" + tag + " | jq -r '.history[].v1Compatibility' | jq '.created' | sort | sed 's/\"//g'|tail -n1 ")
	return curltagresult
}

/*
func trimQuotes(s string) string {
    if len(s) >= 2 {
	        if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			            return s[1 : len(s)-1]
					        }
							    }
								    return s
								}
*/
