package gdeyamloperator

import (
	//"io/ioutil"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	valid "github.com/siangyeh8818/gdeyamlOperator/internal/validation"
	yaml "gopkg.in/yaml.v3"
)

// GroupNexusOutput TODO: explanation here
func GroupNexusOutput(input string, output string, git *GIT) {

	var versionMap = make(map[string]string)
	fileContent, fileContentCount, _ := readLines(input)
	fmt.Printf("File line number : %d\n", fileContentCount)
	for i := 0; i < fileContentCount; i++ {
		fmt.Println(fileContent[i])
		fmt.Println("----------")
		tempContentArray := strings.Split(fileContent[i], "/")
		value, ok := versionMap[tempContentArray[0]]
		if ok == true {
			newVersionArray := strings.Split(tempContentArray[1], ".")
			oldVersionArray := strings.Split(value, ".")
			latestVersion := NexusVersionCompare(newVersionArray, oldVersionArray)
			versionMap[tempContentArray[0]] = latestVersion
		} else if ok == false {
			versionMap[tempContentArray[0]] = tempContentArray[1]
		}
	}
	fmt.Println("------Map start -----")
	fmt.Println(versionMap)
	fmt.Println("------Map end-----")
	/*
		resultContent := putContentToFile(versionMap , fileContent)
		WriteWithIoutil(output, resultContent)
	*/
	putContentToGityaml(versionMap, fileContent, git)

}

func readLines(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer file.Close()

	var lines []string
	linecount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		linecount++
	}
	return lines, linecount, scanner.Err()
}

// NexusVersionCompare TODO: explanation here
func NexusVersionCompare(version1 []string, version2 []string) string {

	maxLength := 0
	var result string
	if len(version1) >= len(version2) {
		maxLength = len(version1)
	} else if len(version2) >= len(version1) {
		maxLength = len(version2)
	}
	for s := 0; s < maxLength; s++ {

		if s > len(version1)-1 {
			result = ComposeString(version2, ".")
			break
		} else if s > len(version2)-1 {
			result = ComposeString(version1, ".")
			break
		}
		intValue1, _ := strconv.Atoi(version1[s])
		intValue2, _ := strconv.Atoi(version2[s])
		if intValue1 > intValue2 {
			result = ComposeString(version1, ".")
			break
		} else if intValue2 > intValue1 {
			result = ComposeString(version2, ".")
			break
		}
	}
	return result
}

// ComposeString TODO: explanation here
func ComposeString(array []string, insertChar string) string {

	var result string
	for i := 0; i < len(array); i++ {
		result = result + array[i]
		if i != len(array)-1 {
			result = result + insertChar
		}
	}
	return result
}

func putContentToFile(Map1 map[string]string, fileContent []string) string {

	var resultContent string
	for i := 0; i < len(fileContent); i++ {
		tempContentArray := strings.Split(fileContent[i], "/")
		if Map1[tempContentArray[0]] == tempContentArray[1] {
			fmt.Printf("Put this content to result : %s", fileContent[i])
			resultContent = resultContent + fileContent[i]
		}
	}
	return resultContent
}

func putContentToGityaml(Map1 map[string]string, fileContent []string, git *GIT) {

	log.Println("-----action >> cloneRepo----")
	pattern, err := valid.Validate(git.Branch)
	if err != nil {
		log.Printf("Validate gitbranch error :%v\n", err)
	}
	switch pattern {
	case "release":
		log.Printf("Gitbranch: %s ,ValidReturn: %s", git.Branch, pattern)
		CloneRepo(git.Url, git.Branch, git.Path, git.AccessUser, git.AccessToken)
	case "patch":
		patternArray := strings.Split(git.Branch, ".")
		log.Printf("len(patternArray): %d", len(patternArray))
		var tempGitBranch string
		// for i := 0; i < len(patternArray); i++ {
		// 	log.Printf("tempGitBranch: %s", tempGitBranch)
		// 	if i == len(patternArray)-1 {
		// 	} else if i == 0 {
		// 		tempGitBranch = patternArray[i]
		// 	} else {
		// 		tempGitBranch = tempGitBranch + "." + patternArray[i]
		// 	}
		// }

		// shouldn't be just that simple?
		tempGitBranch = patternArray[0] + "." + patternArray[1]

		log.Printf("Gitbranch: %s ,ValidReturn: %s , Newbranch: %s", git.Branch, pattern, tempGitBranch)
		CloneRepo(git.Url, tempGitBranch, git.Path, git.AccessUser, git.AccessToken)
	case "feature":
		log.Printf("Gitbranch: %s ,ValidReturn: %s", git.Branch, pattern)
		CloneRepo(git.Url, git.Branch, git.Path, git.AccessUser, git.AccessToken)
	case "misc":
		log.Printf("Gitbranch: %s ,ValidReturn: %s", git.Branch, pattern)
		CloneRepo(git.Url, git.Branch, git.Path, git.AccessUser, git.AccessToken)
	}

	log.Println("-----action >> add urls to deploy.yml----")
	deployYaml := K8sYaml{}
	deployYaml.GetConf(git.Path + "/" + git.CommitFIle)

	var urlArray []string
	for i := 0; i < len(fileContent); i++ {
		tempContentArray := strings.Split(fileContent[i], "/")
		if Map1[tempContentArray[0]] == tempContentArray[1] {
			fmt.Printf("Put this content to deploy.yml : %s", fileContent[i])
			tempInsert := fileContent[i]
			//(&deployYaml.Deployment.SCRIPTS).addUrl(tempInsert)
			urlArray = append(urlArray, tempInsert)
		}
	}
	(&deployYaml.Deployment.SCRIPTS).URLS = (&urlArray)
	outputcontent, err := yaml.Marshal(&deployYaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	WriteWithIoutil(git.Path+"/"+git.CommitFIle, string(outputcontent))
	log.Println("-----action >> CommitRepo----")
	CommitRepo(git, git.CommitFIle)
	log.Println("-----action >> PushGit----")
	PushGit(git.Path, git.AccessUser, git.AccessToken, git.Branch, git.Url)
	log.Println("-----action finishing----")

}
