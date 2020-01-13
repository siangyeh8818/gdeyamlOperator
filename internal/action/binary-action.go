package action

import (
	"fmt"
	"os"

	myConfig "github.com/siangyeh8818/gdeyamlOperator/internal/config"
	/*
		clusterop "github.com/siangyeh8818/gdeyamlOperator/internal/clusterop"
		myConfig "github.com/siangyeh8818/gdeyamlOperator/internal/config"
		myDocker "github.com/siangyeh8818/gdeyamlOperator/internal/docker"
		mygit "github.com/siangyeh8818/gdeyamlOperator/internal/git"
		myJenkins "github.com/siangyeh8818/gdeyamlOperator/internal/jenkins"
		myJson "github.com/siangyeh8818/gdeyamlOperator/internal/json"
		myK8s "github.com/siangyeh8818/gdeyamlOperator/internal/kubernetes"
		myKustomize "github.com/siangyeh8818/gdeyamlOperator/internal/kustomize"
		IO "github.com/siangyeh8818/gdeyamlOperator/internal/myIo"
		myNexus "github.com/siangyeh8818/gdeyamlOperator/internal/nexus"
		ShellCommand "github.com/siangyeh8818/gdeyamlOperator/internal/shellcommand"
		CustomStruct "github.com/siangyeh8818/gdeyamlOperator/internal/structs"
		"gopkg.in/yaml.v2"
	*/)

type BinaryAction struct {
}

func (action BinaryAction) InitConfig() interface{} {

	config := myConfig.BINARYCONFIG{}
	myConfig.Init(&config)
	myConfig.Parse()

	return config
}

func (action BinaryAction) DefineAction(config myConfig.BINARYCONFIG) bool {

	if config.Version {
		fmt.Println("version : 1.12.0")
		os.Exit(0)
	}

	return true
}
