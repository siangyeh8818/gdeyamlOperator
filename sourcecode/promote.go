package main

import (
	"fmt"
	"log"
	"net/http"
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
