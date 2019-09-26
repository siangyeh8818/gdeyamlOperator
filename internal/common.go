package gdeyamloperator

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ImgInfo struct {
	ModName  string
	ImgName  string
	Tag      string
	Stage    string
	ShowInfo bool
}

func GetAllFiles(loc string, suffixFilter string) []string {
	var allFiles []string
	err := filepath.Walk(loc,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			if suffixFilter == "" {
				allFiles = append(allFiles, path)
			} else {
				sf := strings.Split(suffixFilter, "&")
				for i := 0; i < len(sf); i++ {
					if strings.HasSuffix(path, sf[i]) {
						allFiles = append(allFiles, path)
					}
				}

			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return allFiles

}

func tryMatching(inPath string, searchPattern string, suffixPattern string) (string, error) {
	allFiles := GetAllFiles(inPath, suffixPattern)
	for i := 0; i < len(allFiles); i++ {
		if strings.Contains(allFiles[i], searchPattern) {

			return allFiles[i], nil
		}
	}
	return "", errors.New("can not match directory files")
}

func LastStringList(img string) string {
	return strings.Split(img, ":")[len(strings.Split(img, ":"))-1]
}
func LastSecondSlash(img string) string {
	return strings.Split(img, "/")[len(strings.Split(img, "/"))-2]
}

func DetectImageLine(img string) bool {
	if strings.Contains(img, "image: ") {
		if len(strings.Split(img, ":")) > 2 {
			//fmt.Println("got image")
			return true
		} else {
			return false
		}
		return false
	}
	return false
}

func DetectGroup(img string) (string, string) {
	var rTag string
	var rStage string
	//fmt.Println(strings.Split(img, ":")[len(strings.Split(img, ":"))-1])
	rTag = LastStringList(img)
	if len(strings.Split(img, "/")) == 1 {
		rStage = ""
		return rTag, rStage
	}
	if len(strings.Split(img, "/")) < 3 {
		rStage = ""
		rStage = fmt.Sprintf("/%s/", rStage)
		return rTag, rStage
	}
	if len(strings.Split(img, "/")) > 3 {
		fmt.Println("large size of image name")
		segs := strings.Split(img, "/")
		rStage = segs[1]
		for i := 2; i < len(segs)-1; i++ {
			rStage = fmt.Sprintf("%s/%s", rStage, segs[i])
		}
		rStage = fmt.Sprintf("/%s/", rStage)
		return rTag, rStage
	}

	rStage = LastSecondSlash(img)
	rStage = fmt.Sprintf("/%s/", rStage)

	//replace1:=
	//fmt.Println(rTag, rStage)
	return rTag, rStage
}

func TryMatchingContextByDir(inPath string, searchPattern string, suffixPattern string) (ImgInfo, error) {
	allFiles := GetAllFiles(inPath, suffixPattern)
	//fmt.Println(allFiles)
	var imginfo ImgInfo
	for _, input := range allFiles {
		//fmt.Println("start detect file:", input)
		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return ImgInfo{}, errors.New("err")
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//time.Sleep(1 * time.Second)
			img := scanner.Text()
			//fmt.Println(img)
			if DetectImageLine(img) {
				if strings.Contains(img, searchPattern) {
					rTag, rStage := DetectGroup(img)
					fmt.Println(input, rTag, rStage)
					// here ImgName is directory name, however
					//imginfo.ImgName = input
					imginfo.ImgName = searchPattern
					imginfo.Stage = strings.Replace(rStage, "/", "", -1)
					imginfo.Stage = rStage[1 : len(rStage)-1]
					imginfo.Tag = rTag
					imginfo.ShowInfo = true
					return imginfo, nil
				}
			}
		}
		file.Close()
	}
	//log.Fatalf("can not find pattern:", searchPattern, "in this directory:", inPath)
	fmt.Println("can not find pattern:", searchPattern, "in this directory:", inPath)
	imginfo.ImgName = searchPattern
	imginfo.Stage = ""
	imginfo.Tag = ""
	imginfo.ShowInfo = false

	//time.Sleep(10 * time.Second)
	return imginfo, errors.New("can not find img pattern in this directory")
}
func TryMatchingImage(inPath string, searchPattern string, suffixPattern string) (string, error) {
	allFiles := GetAllFiles(inPath, suffixPattern)
	//fmt.Println(allFiles)
	var imginfo ImgInfo
	for _, input := range allFiles {
		//fmt.Println("start detect file:", input)
		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("err")
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//time.Sleep(1 * time.Second)
			img := scanner.Text()
			//fmt.Println(img)
			if DetectImageLine(img) {
				if strings.Contains(img, "image: ") {
					csearchPattern := fmt.Sprintf("%s%s%s", "/", searchPattern, ":")
					if strings.Contains(img, csearchPattern) {
						retImg := strings.Split(img, "image: ")[1]
						//retImg = strings.Split(retImg, ":")[0]
						retImg = strings.TrimSpace(retImg)
						return retImg, nil
					}
				}
			}
		}
		file.Close()
	}
	//log.Fatalf("can not find pattern:", searchPattern, "in this directory:", inPath)
	fmt.Println("can not find pattern:", searchPattern, "in this directory:", inPath)
	panic("can not find pattern; adding more resources in base")
	imginfo.ImgName = searchPattern
	imginfo.Stage = ""
	imginfo.Tag = ""
	imginfo.ShowInfo = false

	//time.Sleep(10 * time.Second)
	return "", errors.New("can not find img pattern in this directory")
}
func TryMatchingDir(inPath string, searchPattern string, suffixPattern string) (string, error) {
	allFiles := GetAllFiles(inPath, suffixPattern)
	//fmt.Println(allFiles)
	for _, input := range allFiles {
		//fmt.Println("start detect file:", input)
		file, err := os.Open(input)
		if err != nil {
			fmt.Println(err)
			return "", errors.New("err")
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			//time.Sleep(1 * time.Second)
			img := scanner.Text()
			//fmt.Println(img)
			if DetectImageLine(img) {
				csearchPattern := fmt.Sprintf("%s%s%s", "/", searchPattern, ":")
				if strings.Contains(img, csearchPattern) {
					//fmt.Println(input)
					ret := strings.Split(input, "/")[len(strings.Split(input, "/"))-2]
					ret = strings.Replace(ret, ".yaml", "", -1)
					ret = strings.Replace(ret, ".yml", "", -1)
					return ret, nil
				}
			}
		}
		file.Close()
	}
	//log.Fatalf("can not find pattern:", searchPattern, "in this directory:", inPath)

	//time.Sleep(10 * time.Second)
	return "", errors.New("can not find img pattern in this directory")
}