package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GetNesuxCpmponet(nexusurl string, nexus_user string, nexus_password string) {

	var url string
	// curl -X GET "https://package.pentium.network/service/rest/v1/searchsort=group&repository=scripts&format=raw" -H "accept: application/json"
	// curl -X GET "https://package.pentium.network/service/rest/v1/search?continuationToken=35303a6562313438303661303938346263663537613436613861663432663439353266&sort=group&repository=scripts&format=raw" -H "accept: application/json"

	url = nexusurl

	fmt.Printf("your request url : %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// handle err
	}
	//req.SetBasicAuth(nexus_user, nexus_password)
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	log.Println("----------srart of responseData-----------")
	log.Println(string(responseData))
	log.Println("----------end of responseData-----------")
	//JsonParse(string(responseData))
	//os.Exit(0)
}

func GET_NesusAPI(nexusurl string, nexus_user string, nexus_password string, outfile string, out_pattern string, output *OutputContent) {
	var token string
	// curl -X GET "https://package.pentium.network/service/rest/v1/searchsort=group&repository=scripts&format=raw" -H "accept: application/json"
	// curl -X GET "https://package.pentium.network/service/rest/v1/search?continuationToken=35303a6562313438303661303938346263663537613436613861663432663439353266&sort=group&repository=scripts&format=raw" -H "accept: application/json"

	fmt.Printf("your request url : %s\n", nexusurl)
	req, err := http.NewRequest("GET", nexusurl, nil)
	if err != nil {
		// handle err
	}
	if nexus_user != "" && nexus_password != "" {
		req.SetBasicAuth(nexus_user, nexus_password)
	}

	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	log.Println("---------srart of responseData---------")
	log.Println(string(responseData))
	log.Println("---------end of responseData-----------")
	log.Println(resp.Status)
	log.Println(resp)

	//if out_pattern != "" {
	log.Println("---------srart of jaonparse------------")
	JsonParse2(string(responseData), out_pattern, output)
	log.Println("---------end of jaonparse----------------")
	//}
	token = continueTokenParse(string(responseData))
	if token == "null" || token == "" {
		log.Printf("output.content 數量 : %d\n", len(output.Content))
		var temp_out string
		for i := 0; i < len(output.Content); i++ {
			temp_out = temp_out + output.Content[i] + "\n"
			//log.Println(temp_out)
		}
		WriteWithIoutil(outfile, string(temp_out))
	} else {
		origin_request := strings.Split(nexusurl, "?")
		new_request_url := origin_request[0] + "?" + "continuationToken=" + token + origin_request[1]
		GET_NesusAPI(new_request_url, nexus_user, nexus_password, outfile, out_pattern, output)
	}

}

func POST_NesusAPI(nexusurl string, nexus_user string, nexus_password string, request_body string) {

	// curl -X POST "https://package.pentium.network/service/rest/v1/staging/move/events-preview?repository=events&name=siang-test%2F01%2Fevent.yml" -H "accept: application/json"

	fmt.Printf("your request url : %s\n", nexusurl)
	req, err := http.NewRequest("POST", nexusurl, strings.NewReader(request_body))
	if err != nil {
		// handle err
	}
	if nexus_user != "" && nexus_password != "" {
		req.SetBasicAuth(nexus_user, nexus_password)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	log.Println(string(responseData))
	log.Println(resp.Status)
	log.Println(resp)

}

func PUT_NesusAPI(nexusurl string, nexus_user string, nexus_password string, request_body string) {

	// curl -X POST "https://package.pentium.network/service/rest/v1/staging/move/events-preview?repository=events&name=siang-test%2F01%2Fevent.yml" -H "accept: application/json"

	fmt.Printf("your request url : %s\n", nexusurl)
	req, err := http.NewRequest("PUT", nexusurl, strings.NewReader(request_body))
	if err != nil {
		// handle err
	}
	if nexus_user != "" && nexus_password != "" {
		req.SetBasicAuth(nexus_user, nexus_password)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	log.Println(string(responseData))
	log.Println(resp.Status)
	log.Println(resp)
}

func DELETE_NesusAPI(nexusurl string, nexus_user string, nexus_password string, request_body string) {

	// curl -X POST "https://package.pentium.network/service/rest/v1/staging/move/events-preview?repository=events&name=siang-test%2F01%2Fevent.yml" -H "accept: application/json"

	fmt.Printf("your request url : %s\n", nexusurl)
	req, err := http.NewRequest("DELETE", nexusurl, strings.NewReader(request_body))
	if err != nil {
		// handle err
	}
	if nexus_user != "" && nexus_password != "" {
		req.SetBasicAuth(nexus_user, nexus_password)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	log.Println(string(responseData))
	log.Println(resp.Status)
	log.Println(resp)
}

func POSTForm_NesusAPI(nexusurl string, nexus_user string, nexus_password string, filename string, dest string) {

	log.Println("POSTForm_NesusAPI")
	log.Println(filename)
	//curl -X POST "https://package.pentium.network/service/rest/v1/components?repository=scripts-qa"
	//-H "accept: application/json" -H "Content-Type: multipart/form-data"
	//-F "raw.directory=/byos-host-bootstrap-uninstall/0.3/dist" -F "raw.asset1=@script.zip;type=application/zip" -F "raw.asset1.filename=script.zip"
	path, _ := os.Getwd()
	path += "/" + filename
	//fmt.Printf("your request url : %s\n", nexusurl)
	extraParams := map[string]string{
		"raw.directory":       "/tmp",
		"raw.asset1":          "@" + filename + ";type=application/zip",
		"raw.asset1.filename": filename,
	}

	req, err := newfileUploadRequest("https://package.pentium.network/service/rest/v1/components?repository="+dest, extraParams, filename, filename)
	//req, err := http.NewRequest("POST", nexusurl, strings.NewReader(request_body))
	if err != nil {
		// handle err
	}
	if nexus_user != "" && nexus_password != "" {
		req.SetBasicAuth(nexus_user, nexus_password)
	}
	req.Header.Set("Accept", "application/json")

	req.Header.Set("Content-Type", "multipart/form-data")
	log.Println(req.Header)

	log.Println("------------------")
	log.Println(req)
	resp, err := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	if err != nil {
		log.Println("request failed")
	}
	//responseData, err := ioutil.ReadAll(resp.Body)
	//log.Println(string(responseData))
	log.Println(resp.Status)
	log.Println(resp)

}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	log.Println("newfileUploadRequest")
	log.Println(path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	//log.Println(body)
	req, err := http.NewRequest("POST", uri, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, err
}
