package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func promoteimage(nexusurl string, repository string, nexus_user string, nexus_password string, imagename string, imagetag string) {

	var url string
	if repository != "" {
		url = nexusurl + "?repository=" + repository + "&docker.imageName=" + imagename + "&docker.imageTag=" + imagetag
	} else if repository == "" {
		url = nexusurl + "?docker.imageName=" + imagename + "&docker.imageTag=" + imagetag
	}
	fmt.Printf("your promote request url : %s", url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// handle err
	}
	req.SetBasicAuth(nexus_user, nexus_password)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("move image failed")
	}
	log.Println(resp)

}

func cpcomponetname(nexusurl string, nexus_user string, nexus_password string, dest string) {

	var output OutputContent

	GET_NesusAPI(nexusurl, nexus_user, nexus_password, "", "name,assets.downloadUrl", &output)

	output_array := strings.Split(string(output.Content[0]), ",")

	downloadComponetFile(output_array[0], output_array[1])

	temp_filename, temp_path := ParserRealFilename(output_array[0])

	//POSTForm_NesusAPI(nexusurl, nexus_user, nexus_password, temp_path, dest)
	cmd := "curl -u " + nexus_user + ":" + nexus_password + " -X POST \"https://package.pentium.network/service/rest/v1/components?repository=" + dest + "\" -H \"accept: application/json\"" + " -H \"Content-Type: multipart/form-data\"" +
		" -F " + "\"raw.directory=" + temp_path + "\"" + " -F " + "\"raw.asset1=@" + temp_filename + ";type=application/zip\"" + " -F " + "\"raw.asset1.filename=" + temp_filename + "\""
	log.Println(cmd)
	exec_shell(cmd)
}

func downloadComponetFile(fileFullPath string, downloadurl string) {

	filename, _ := ParserRealFilename(fileFullPath)

	resp, err := http.Get(downloadurl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}

func ParserRealFilename(path string) (string, string) {

	path_array := strings.Split(path, "/")
	temp_path := ""
	for i := 0; i < len(path_array)-1; i++ {
		temp_path = temp_path + "/" + path_array[i]

	}
	return path_array[len(path_array)-1], temp_path

}

/*
func downloadFile(fileFullPath string, res *restful.Response) {
	file, err := os.Open(fileFullPath)

	if err != nil {
		res.WriteEntity(_dto.ErrorDto{Err: err})
		return
	}

	defer file.Close()
	fileName := path.Base(fileFullPath)
	fileName = url.QueryEscape(fileName) // 防止中文乱码
	res.AddHeader("Content-Type", "application/octet-stream")
	res.AddHeader("content-disposition", "attachment; filename=\""+fileName+"\"")
	_, error := io.Copy(res.ResponseWriter, file)
	if error != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

*/
