package main

import (
	"fmt"
	"log"
	"net/http"
)

func promoteimage(nexusurl string, nexus_user string, nexus_password string, imagename string, imagetag string) {

	//url := "https://package.pentium.network/service/rest/v1/staging/move/docker-qa" + "?docker.imageName=" + imagename + "&docker.imageTag=" + imagetag
	url := nexusurl + "?docker.imageName=" + imagename + "&docker.imageTag=" + imagetag
	fmt.Println(url)
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
