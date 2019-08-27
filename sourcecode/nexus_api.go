package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	JsonParse(string(responseData))
	os.Exit(0)
}

func GET_NesusAPI(nexusurl string, nexus_user string, nexus_password string) {

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
	log.Println(string(responseData))
	log.Println(resp.Status)
	log.Println(resp)
	//JsonParse(string(responseData))
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
