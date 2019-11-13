package gdeyamloperator

import (
	"fmt"
	"log"
	"os"
	"strings"

	yaml "gopkg.in/yaml.v3"
)

type REPLACEYAML struct {
	Type     string
	Pattern  string
	Image    string
	NewValue string
	YamlType string
}

func (rep *REPLACEYAML) UpdateREPLACEYAML(types string, pattern string, image string, newvalue string , yamltype string) {
	rep.Type = types
	rep.Pattern = pattern
	rep.Image = image
	rep.NewValue = newvalue
	rep.YamlType = yamltype
}

func Replacedeploymentfile(environment string, deployfile string, outputfile string) {
	envir_yaml := Environmentyaml{}
	envir_yaml.GetConf(environment)

	inyaml := K8sYaml{}
	inyaml.GetConf(deployfile)

	Replace_total := len(envir_yaml.Deploymentfile[0].Replace.K8S) + len(envir_yaml.Deploymentfile[0].Replace.Openfaas) + len(envir_yaml.Deploymentfile[0].Replace.Monitor) + len(envir_yaml.Deploymentfile[0].Replace.Redis)
	fmt.Printf("Replace_total : %d", Replace_total)
	if Replace_total > 0 {
		if len(envir_yaml.Deploymentfile[0].Replace.K8S) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.K8S); i++ {
				current_index := SearchReplace(&inyaml, envir_yaml.Deploymentfile[0].Replace.K8S[i].Image, "k8s")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Replace,this module is not exist in k8s")
				}else {
					(&inyaml.Deployment.K8S[current_index]).UpdateK8sTag(envir_yaml.Deploymentfile[0].Replace.K8S[i].Tag)
					(&inyaml.Deployment.K8S[current_index]).UpdateK8sModule(envir_yaml.Deploymentfile[0].Replace.K8S[i].Module)
					(&inyaml.Deployment.K8S[current_index]).UpdateK8sImage(envir_yaml.Deploymentfile[0].Replace.K8S[i].Image)
					(&inyaml.Deployment.K8S[current_index]).UpdateK8sStage(envir_yaml.Deploymentfile[0].Replace.K8S[i].Stage)
				}
			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Openfaas) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Openfaas); i++ {
				current_index := SearchReplace(&inyaml, envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Image, "openfaas")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Replace,this module is not exist in openfaas")
				}else {
					(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasTag(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Tag)
					(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasModule(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Module)
					(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasImage(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Image)
					(&inyaml.Deployment.Openfaas[current_index]).UpdateOpenfaasStage(envir_yaml.Deploymentfile[0].Replace.Openfaas[i].Stage)
				}

			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Monitor) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Monitor); i++ {
				current_index := SearchReplace(&inyaml, envir_yaml.Deploymentfile[0].Replace.Monitor[i].Image, "monitor")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Replace,this module is not exist in monitor")
				}else {
					(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorTag(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Tag)
					(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorModule(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Module)
					(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorImage(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Image)
					(&inyaml.Deployment.Monitor[current_index]).UpdateMonitorStage(envir_yaml.Deploymentfile[0].Replace.Monitor[i].Stage)
				}

			}
		}
		if len(envir_yaml.Deploymentfile[0].Replace.Redis) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Replace.Redis); i++ {
				current_index := SearchReplace(&inyaml, envir_yaml.Deploymentfile[0].Replace.Redis[i].Image, "redis")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Replace,this module is not exist in redis")
				}else {
					(&inyaml.Deployment.Redis[current_index]).UpdateRedisTag(envir_yaml.Deploymentfile[0].Replace.Redis[i].Tag)
					(&inyaml.Deployment.Redis[current_index]).UpdateRedisModule(envir_yaml.Deploymentfile[0].Replace.Redis[i].Module)
					(&inyaml.Deployment.Redis[current_index]).UpdateRedisImage(envir_yaml.Deploymentfile[0].Replace.Redis[i].Image)
					(&inyaml.Deployment.Redis[current_index]).UpdateRedisStage(envir_yaml.Deploymentfile[0].Replace.Redis[i].Stage)
				}

			}
		}
	}

	Ignore_total := len(envir_yaml.Deploymentfile[0].Ignore.K8S) + len(envir_yaml.Deploymentfile[0].Ignore.Openfaas) + len(envir_yaml.Deploymentfile[0].Ignore.Monitor) + len(envir_yaml.Deploymentfile[0].Ignore.Redis)
	fmt.Printf("Ignore_total : %d", Ignore_total)
	if Ignore_total > 0 {
		if len(envir_yaml.Deploymentfile[0].Ignore.K8S) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Ignore.K8S); i++ {
				current_index := SearchYamlModuleIndex(&inyaml, envir_yaml.Deploymentfile[0].Ignore.K8S[i].Module, "k8s")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Ignore,this module is not exist in k8s")
				}else {
					(&inyaml.Deployment).RemoveK8sStruct(current_index)
				}
				
			}
		}
		if len(envir_yaml.Deploymentfile[0].Ignore.Openfaas) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Ignore.Openfaas); i++ {
				current_index := SearchYamlModuleIndex(&inyaml, envir_yaml.Deploymentfile[0].Ignore.Openfaas[i].Module, "openfaas")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Ignore,this module is not exist in openfaas")
				}else {
					(&inyaml.Deployment).RemoveOpenfaasStruct(current_index)
				}
				
			}

		}
		if len(envir_yaml.Deploymentfile[0].Ignore.Monitor) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Ignore.Monitor); i++ {
				current_index := SearchYamlModuleIndex(&inyaml, envir_yaml.Deploymentfile[0].Ignore.Monitor[i].Module, "monitor")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Ignore,this module is not exist in monitor")
				}else {
					(&inyaml.Deployment).RemoveMonitorStruct(current_index)
				}
			}
		}
		if len(envir_yaml.Deploymentfile[0].Ignore.Redis) > 0 {
			for i := 0; i < len(envir_yaml.Deploymentfile[0].Ignore.Redis); i++ {
				current_index := SearchYamlModuleIndex(&inyaml, envir_yaml.Deploymentfile[0].Ignore.Redis[i].Module, "redis")
				fmt.Printf("current_index : %d", current_index)
				if current_index ==-1 {
                    fmt.Println("Action Ignore,this module is not exist in redis")
				}else {
					(&inyaml.Deployment).RemoveRedisStruct(current_index)
				}
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

//func SearchReplace(envir_yaml *Environmentyaml, inyaml *K8sYaml, imagesname string, rangestr string) int {
func SearchReplace(inyaml *K8sYaml, imagesname string, rangestr string) int {
	var resultindex int
	changotoken := false
	switch rangestr {
	case "k8s":
		for i := 0; i < len(inyaml.Deployment.K8S); i++ {
			if inyaml.Deployment.K8S[i].Image == imagesname {
				resultindex = i
				changotoken = true
			}
		}
	case "openfaas":
		for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
			if inyaml.Deployment.Openfaas[i].Image == imagesname {
				resultindex = i
				changotoken = true
			}
		}
	case "monitor":
		for i := 0; i < len(inyaml.Deployment.Monitor); i++ {
			if inyaml.Deployment.Monitor[i].Image == imagesname {
				resultindex = i
				changotoken = true
			}
		}
	case "redis":
		for i := 0; i < len(inyaml.Deployment.Redis); i++ {
			if inyaml.Deployment.Redis[i].Image == imagesname {
				resultindex = i
				changotoken = true
			}
		}
	}
	if changotoken == false {
		resultindex = -1
	}

	return resultindex
}

func SearchYamlModuleIndex(inyaml *K8sYaml, modulename string, rangestr string) int {
	var resultindex int
	changotoken := false
	switch rangestr {
	case "k8s":
		for i := 0; i < len(inyaml.Deployment.K8S); i++ {
			if inyaml.Deployment.K8S[i].Module == modulename {
				resultindex = i
				changotoken = true
			}
		}
	case "openfaas":
		for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
			if inyaml.Deployment.Openfaas[i].Module == modulename {
				resultindex = i
				changotoken = true
			}
		}
	case "monitor":
		for i := 0; i < len(inyaml.Deployment.Monitor); i++ {
			if inyaml.Deployment.Monitor[i].Module == modulename {
				resultindex = i
				changotoken = true
			}
		}
	case "redis":
		for i := 0; i < len(inyaml.Deployment.Redis); i++ {
			if inyaml.Deployment.Redis[i].Module == modulename {
				resultindex = i
				changotoken = true
			}
		}
	}
	if changotoken == false {
		resultindex = -1
	}
	return resultindex
}

func Replacedeploymentfile_Image_Tag(rep *REPLACEYAML, inputfile string, outputfile string) {

	deployyaml := K8sYaml{}
	deployyaml.GetConf(inputfile)

	current_index1 := SearchReplace(&deployyaml, rep.Image, "k8s")
	if current_index1 != -1 {
		(&deployyaml.Deployment.K8S[current_index1]).UpdateK8sTag(rep.NewValue)
	}
	current_index2 := SearchReplace(&deployyaml, rep.Image, "openfaas")
	if current_index2 != -1 {
		(&deployyaml.Deployment.Openfaas[current_index2]).UpdateOpenfaasTag(rep.NewValue)
	}
	current_index3 := SearchReplace(&deployyaml, rep.Image, "monitor")
	if current_index3 != -1 {
		(&deployyaml.Deployment.Monitor[current_index3]).UpdateMonitorTag(rep.NewValue)
	}
	current_index4 := SearchReplace(&deployyaml, rep.Image, "redis")
	if current_index4 != -1 {
		(&deployyaml.Deployment.Redis[current_index4]).UpdateRedisTag(rep.NewValue)
	}

	yamlcontent, err := yaml.Marshal(&deployyaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	WriteWithIoutil(outputfile, string(yamlcontent))
}

func ReplacedeByPattern(rep *REPLACEYAML, inputfile string, outputfile string) {

	if rep.YamlType == "deployyaml" {
		UpdateDeployFile(rep , inputfile , outputfile)

	}else if rep.YamlType == "environmentyaml"{
		UpdateEnvironmentFile(rep , inputfile , outputfile)
	}
}

func UpdateEnvironmentFile(rep *REPLACEYAML, inputfile string, outputfile string) {

	environmentfile := Environmentyaml{}
	fmt.Println("284")
	fmt.Println(inputfile)
	environmentfile.GetConf(inputfile)
	fmt.Println("286")
	pattern := strings.Split(rep.Pattern, ":")
	switch pattern[0] {
	case "configuration":
		switch pattern[1] {
		case "branch":
			//temp_branch := (&environmentfile.Configuration[0]).Branch
			(&environmentfile.Configuration[0]).UpdateBranch(rep.NewValue)
		}
	case "deploymentfile":	
		switch pattern[1] {
		case "branch":
			//temp_branch := (&environmentfile.Deploymentfile[0]).Branch
			fmt.Println(environmentfile.Deploymentfile[0])
			(&environmentfile.Deploymentfile[0]).UpdateBranch(rep.NewValue)
		}	
	}

	yamlcontent, err := yaml.Marshal(&environmentfile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	WriteWithIoutil(outputfile, string(yamlcontent))

}


func UpdateDeployFile(rep *REPLACEYAML, inputfile string, outputfile string) {

	deployyaml := K8sYaml{}
	deployyaml.GetConf(inputfile)

	pattern := strings.Split(rep.Pattern, ":")

	switch pattern[0] {
	case "base":
		switch pattern[1] {
		case "git":
			temp_branch := (&deployyaml.Deployment).BASE[0].Branch
			(&deployyaml.Deployment).UpdateBaseStructBranch(rep.NewValue, temp_branch)
		case "branch":
			temp_git := (&deployyaml.Deployment).BASE[0].Git
			(&deployyaml.Deployment).UpdateBaseStructBranch(temp_git, rep.NewValue)
		}
	case "blcks":
		switch pattern[1] {
		case "git":
			temp_version := (&deployyaml.Deployment).BLCKS.Version
			temp_branch := (&deployyaml.Deployment).BLCKS.Branch
			(&deployyaml.Deployment).UpdateBlcksStructBranch(rep.NewValue, temp_branch, temp_version)
		case "branch":
			temp_git := (&deployyaml.Deployment).BLCKS.Git
			temp_version := (&deployyaml.Deployment).BLCKS.Version
			(&deployyaml.Deployment).UpdateBlcksStructBranch(temp_git, rep.NewValue, temp_version)
		case "version":
			temp_git := (&deployyaml.Deployment).BLCKS.Git
			temp_branch := (&deployyaml.Deployment).BLCKS.Branch
			(&deployyaml.Deployment).UpdateBlcksStructBranch(temp_git, temp_branch, rep.NewValue)

		}

	case "playbooks":
		switch pattern[1] {
		case "git":
			temp_version := (&deployyaml.Deployment).PLAYBOOKS.Version
			temp_branch := (&deployyaml.Deployment).PLAYBOOKS.Branch
			(&deployyaml.Deployment).UpdatePLAYBOOKStructBranch(rep.NewValue, temp_branch, temp_version)
		case "branch":
			temp_git := (&deployyaml.Deployment).PLAYBOOKS.Git
			temp_version := (&deployyaml.Deployment).PLAYBOOKS.Version
			(&deployyaml.Deployment).UpdatePLAYBOOKStructBranch(temp_git, rep.NewValue, temp_version)
		case "version":
			temp_git := (&deployyaml.Deployment).PLAYBOOKS.Git
			temp_branch := (&deployyaml.Deployment).PLAYBOOKS.Branch
			(&deployyaml.Deployment).UpdatePLAYBOOKStructBranch(temp_git, temp_branch, rep.NewValue)
		}
	case "all":
		switch pattern[1] {
		case "module":
			fmt.Println("This pattern have't not supported~~~~~")
			os.Exit(0)
		case "image":
			fmt.Println("This pattern have't not supported~~~~~")
			os.Exit(0)
		case "tag":
			current_index1 := SearchReplace(&deployyaml, rep.Image, "k8s")
			fmt.Printf("current_index1 : %d\n", current_index1)
			if current_index1 != -1 {
				(&deployyaml.Deployment.K8S[current_index1]).UpdateK8sTag(rep.NewValue)
			}
			current_index2 := SearchReplace(&deployyaml, rep.Image, "openfaas")
			fmt.Printf("current_index2 : %d\n", current_index2)
			if current_index2 != -1 {
				(&deployyaml.Deployment.Openfaas[current_index2]).UpdateOpenfaasTag(rep.NewValue)
			}
			current_index3 := SearchReplace(&deployyaml, rep.Image, "monitor")
			fmt.Printf("current_index3 : %d\n", current_index3)
			if current_index3 != -1 {
				(&deployyaml.Deployment.Monitor[current_index3]).UpdateMonitorTag(rep.NewValue)
			}
			current_index4 := SearchReplace(&deployyaml, rep.Image, "redis")
			fmt.Printf("current_index4 : %d\n", current_index4)
			if current_index4 != -1 {
				(&deployyaml.Deployment.Redis[current_index4]).UpdateRedisTag(rep.NewValue)
			}
			if rep.Image == deployyaml.Deployment.BLCKS.TOOL.Image {
				(&deployyaml.Deployment.BLCKS.TOOL).UpdateToolTag(rep.NewValue)
			}
			if rep.Image == deployyaml.Deployment.PLAYBOOKS.TOOL.Image {
				(&deployyaml.Deployment.PLAYBOOKS.TOOL).UpdateToolTag(rep.NewValue)
			}
		}
	}
	yamlcontent, err := yaml.Marshal(&deployyaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	WriteWithIoutil(outputfile, string(yamlcontent))
}


func PatchDeployFile(rep *REPLACEYAML, inputfile string, outputfile string , kust *KustomizeArgument) {

	deployyaml := K8sYaml{}
	deployyaml.GetConf(inputfile)
	var ss = make(map[string]int)

	base_folder := grepFolderName(rep.Image, kust.K8sBaseloc, ss)

	if len(base_folder) > 0 {
		
		switch rep.Pattern {
		case "k8s":
			current_index1 := SearchYamlModuleIndex(&deployyaml,base_folder , "k8s")
			if current_index1 == -1 {

				ss[base_folder] = 1
				(&deployyaml.Deployment).AddK8sStruct(base_folder, rep.Image, rep.NewValue, "")
			}else {
				(&deployyaml.Deployment.K8S[current_index1]).UpdateK8sImage(rep.Image)
				(&deployyaml.Deployment.K8S[current_index1]).UpdateK8sTag(rep.NewValue)
			}
		case "openfaas":
			current_index1 := SearchYamlModuleIndex(&deployyaml,base_folder , "openfaas")
			if current_index1 == -1 {
				ss[base_folder] = 1
				(&deployyaml.Deployment).AddOpenfaasStruct(base_folder, rep.Image, rep.NewValue, "")
			}else {
				(&deployyaml.Deployment.Openfaas[current_index1]).UpdateOpenfaasImage(rep.Image)
				(&deployyaml.Deployment.Openfaas[current_index1]).UpdateOpenfaasTag(rep.NewValue)
			}
		case "monitor":
			current_index1 := SearchYamlModuleIndex(&deployyaml,base_folder , "monitor")
			if current_index1 == -1 {
				ss[base_folder] = 1
				(&deployyaml.Deployment).AddMonitorStruct(base_folder, rep.Image, rep.NewValue, "")
			}else {
				(&deployyaml.Deployment.Monitor[current_index1]).UpdateMonitorImage(rep.Image)
				(&deployyaml.Deployment.Monitor[current_index1]).UpdateMonitorTag(rep.NewValue)
			}
		case "redis":
			current_index1 := SearchYamlModuleIndex(&deployyaml,base_folder , "redis")
			if current_index1 == -1 {
				ss[base_folder] = 1
				(&deployyaml.Deployment).AddRedisStruct(base_folder, rep.Image, rep.NewValue, "")
			}else {
				(&deployyaml.Deployment.Redis[current_index1]).UpdateRedisImage(rep.Image)
				(&deployyaml.Deployment.Redis[current_index1]).UpdateRedisTag(rep.NewValue)
			}
		case "blcks/ansibleDeployTool":
			(&deployyaml.Deployment.PLAYBOOKS.TOOL).UpdateToolImage(rep.Image)
			(&deployyaml.Deployment.PLAYBOOKS.TOOL).UpdateToolTag(rep.NewValue)
			(&deployyaml.Deployment.BLCKS.TOOL).UpdateToolImage(rep.Image)
			(&deployyaml.Deployment.BLCKS.TOOL).UpdateToolTag(rep.NewValue)
		}

	} else {
		fmt.Println("folder name can't be space")
		(&deployyaml.Deployment).AddK8sStruct("You_have_to_fix_base_repo", rep.Image, rep.NewValue, "")
	}

	yamlcontent, err := yaml.Marshal(&deployyaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	WriteWithIoutil(outputfile, string(yamlcontent))
}