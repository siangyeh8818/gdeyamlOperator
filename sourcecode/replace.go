package main

import (
	"fmt"
	"log"

	yaml "gopkg.in/yaml.v3"
)

func Replacedeploymentfile(environment string, deployfile string, outputfile string) {
	envir_yaml := Environmentyaml{}
	envir_yaml.getConf(environment)

	inyaml := K8sYaml{}
	inyaml.getConf(deployfile)

	Replace_total := len(envir_yaml.Deploymentfile[0].Replace.K8S) + len(envir_yaml.Deploymentfile[0].Replace.Openfaas) + len(envir_yaml.Deploymentfile[0].Replace.Monitor) + len(envir_yaml.Deploymentfile[0].Replace.Redis)
	fmt.Printf("Replace_total : %d", Replace_total)
	if Replace_total > 0 {
		if len(envir_yaml.Deploymentfile[0].Replace.K8S) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.K8S); i++ {
				current_index := SearchReplace(&envir_yaml, &inyaml, envir_yaml.Deploymentfile[0].Replace.K8S[i].Image, "k8s")
				fmt.Printf("current_index : %d", current_index)
				(&inyaml.Deployment.K8S[current_index]).UpdateK8sTag(envir_yaml.Deploymentfile[0].Replace.K8S[i].Tag)
				(&inyaml.Deployment.K8S[current_index]).UpdateK8sModule(envir_yaml.Deploymentfile[0].Replace.K8S[i].Module)
				(&inyaml.Deployment.K8S[current_index]).UpdateK8sImage(envir_yaml.Deploymentfile[0].Replace.K8S[i].Image)
				(&inyaml.Deployment.K8S[current_index]).UpdateK8sStage(envir_yaml.Deploymentfile[0].Replace.K8S[i].Stage)
			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Openfaas) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Openfaas); i++ {
				current_index := SearchReplace(&envir_yaml, &inyaml, envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Image, "openfaas")
				fmt.Printf("current_index : %d", current_index)
				(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasTag(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Tag)
				(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasModule(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Module)
				(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasImage(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Image)
				(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasStage(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Stage)
			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Monitor) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Monitor); i++ {
				current_index := SearchReplace(&envir_yaml, &inyaml, envir_yaml.Deploymentfile[0].Replace.Monitor[i].Image, "monitor")
				fmt.Printf("current_index : %d", current_index)
				(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorTag(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Tag)
				(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorModule(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Module)
				(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorImage(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Image)
				(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorStage(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Stage)
			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Redis) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Redis); i++ {
				current_index := SearchReplace(&envir_yaml, &inyaml, envir_yaml.Deploymentfile[0].Replace.Redis[i].Image, "redis")
				fmt.Printf("current_index : %d", current_index)
				(&inyaml.Deployment.Redis[current_index]).UpdateRedisTag(envir_yaml.Deploymentfile[0].Replace.Redis[i].Tag)
				(&inyaml.Deployment.Redis[current_index]).UpdateRedisModule(envir_yaml.Deploymentfile[0].Replace.Redis[i].Module)
				(&inyaml.Deployment.Redis[current_index]).UpdateRedisImage(envir_yaml.Deploymentfile[0].Replace.Redis[i].Image)
				(&inyaml.Deployment.Redis[current_index]).UpdateRedisStage(envir_yaml.Deploymentfile[0].Replace.Redis[i].Stage)
			}
		}
	}
	yamlcontent, err := yaml.Marshal(&inyaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//	fmt.Printf("--- t dump:\n%s\n\n", string(yamlcontent))

	WriteWithIoutil(outputfile, string(yamlcontent))
}

func SearchReplace(envir_yaml *Environmentyaml, inyaml *K8sYaml, imagesname string, rangestr string) int {
	var resultindex int
	switch rangestr {
	case "k8s":
		for i := 0; i < len(inyaml.Deployment.K8S); i++ {
			if inyaml.Deployment.K8S[i].Image == imagesname {
				resultindex = i
			}
		}
	case "openfaas":
		for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
			if inyaml.Deployment.Openfaas[i].Image == imagesname {
				resultindex = i
			}
		}
	case "monitor":
		for i := 0; i < len(inyaml.Deployment.Monitor); i++ {
			if inyaml.Deployment.Monitor[i].Image == imagesname {
				resultindex = i
			}
		}
	case "redis":
		for i := 0; i < len(inyaml.Deployment.Redis); i++ {
			if inyaml.Deployment.Redis[i].Image == imagesname {
				resultindex = i
			}
		}
	}
	return resultindex
}
