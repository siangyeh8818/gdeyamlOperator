package gdeyamloperator

import (
	"fmt"

	"github.com/bndr/gojenkins"
)

func INit_Jenkins() {
	/*
		httpclient := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}}
	*/
	jenkins := gojenkins.CreateJenkins("https://ci.pentium.network/", "", "").Init()
	fmt.Printf("jenkins: %+v\n", jenkins)
	// Provide CA certificate if server is using self-signed certificate
	// caCert, _ := ioutil.ReadFile("/tmp/ca.crt")
	// jenkins.Requester.CACert = caCert
	nodes := jenkins.GetAllNodes()
	for _, node := range nodes {

		// Fetch Node Data
		node.Poll()
		if node.IsOnline() {
			fmt.Println("Node is Online")
		}
		fmt.Printf("jenkins node: %+v\n", nodes)
	}

	job := jenkins.GetAllJobs(false)
	//jobparameters := job.GetParameters()
	//jobconfig := job.GetConfig()

	//fmt.Printf("Config of job deploy.v2/: %+v\n", jobconfig)
	//fmt.Printf("Parameters of job (deploy.v2/: %+v\n", jobparameters)
	for i := 0; i < len(job); i++ {
		fmt.Println(job[i].GetName())
		//job[i].GetUpstreamJobs()
		//job[i].GetDetails()
		//job[i].GetConfig()
		//	fmt.Println(job[i].Raw)
		//	fmt.Println(job[i].Jenkins)
		if job[i].GetName() == "deploy.v2" {
			//jenkins.GetSubJob(job[i].GetName(), childId string)
		}

	}
	//fmt.Println("-------GetAllJobNames-------")
	//innerjob, _ := jenkins.GetAllJobNames()
	//for i := 0; i < len(innerjob); i++ {
	//	fmt.Println(innerjob[i].GetName())
	//}

}
