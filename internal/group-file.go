package gdeyamloperator

import(
	//"io/ioutil"
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func GroupNexusOutput(input string , output string){

	var versionMap = make(map[string]string)
	fileContent , fileContentCount , _ := readLines(input)
	for i:=0 ; i<fileContentCount ; i++ {
		tempContentArray := strings.Split(fileContent[i],"/")
		value , ok := versionMap[tempContentArray[0]]
        if ok==true{
			newVersionArray := strings.Split(tempContentArray[1],".")
			oldVersionArray := strings.Split(value,".")
			latestVersion := NexusVersionCompare(newVersionArray,oldVersionArray)
			versionMap[tempContentArray[0]] = latestVersion
		}else if ok==false{
			versionMap[tempContentArray[0]] = tempContentArray[1]
		}
	}
	fmt.Println("------Map start -----")
	fmt.Println(versionMap)
	fmt.Println("------Map end-----")
	resultContent := putContentToFile(versionMap , fileContent)
	
	WriteWithIoutil(output, resultContent)
}

func readLines(path string) ([]string, int, error) {
	file, err := os.Open(path)
	if err != nil {
	  return nil,0, err
	}
	defer file.Close()
  
	var lines []string
	linecount :=0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	  linecount++
	}
	return lines,linecount,scanner.Err()
  }


func NexusVersionCompare(version1 []string , version2 []string)string{

	maxLength :=0
	var result string
	if len(version1) >= len(version2){
		maxLength = len(version1)
	}else if len(version2) >= len(version1){
		maxLength = len(version2)
	}
    for s:=0 ; s<maxLength ; s++ {
		
		if s > len(version1)-1{
			result = ComposeString(version2,".")
		}else if s > len(version2)-1{
			result = ComposeString(version1,".")
		}
		intValue1,_ := strconv.Atoi(version1[s])
		intValue2,_ := strconv.Atoi(version2[s])
		if intValue1 > intValue2{
			result = ComposeString(version1,".")
			break
		}else if intValue2 > intValue1 {
            result = ComposeString(version2,".")
			break
		}
	}
	return result
}

func ComposeString(array []string , insertChar string)string {

	var result string
	for i:=0 ; i<len(array) ; i++ {
		result = result + array[i]
		if i != len(array)-1 {
			result = result + insertChar
		}
	}
	return result
}

func putContentToFile( Map1 map[string]string , fileContent []string)string{

	var resultContent string
	for i:=0 ; i<len(fileContent) ; i++ {
		tempContentArray := strings.Split(fileContent[i],"/")
		if Map1[tempContentArray[0]]== tempContentArray[1]{
			fmt.Printf("Put this content to result : %s",fileContent[i])
			resultContent = resultContent + fileContent[i]
		}
	}
	return resultContent
}