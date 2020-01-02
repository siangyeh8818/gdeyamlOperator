package config

import (
	//. "github.com/siangyeh8818/gdeyamlOperator/internal"
	myDocker "github.com/siangyeh8818/gdeyamlOperator/internal/docker"
	kustomize "github.com/siangyeh8818/gdeyamlOperator/internal/kustomize"
	myNexus "github.com/siangyeh8818/gdeyamlOperator/internal/nexus"
)

type ACTION interface {
	InitConfig()
	RunAction() bool
}

type BINARYCONFIG struct {
	Action string
	//-----git flag -----------
	/*
		GitUrl string                               //GIT{}
		GitRepoPath  string                         //GIT{}
		GitUser  string                             //GIT{}
		GitToken string                             //GIT{}
		GitBranch string                            //GIT{}
		GitCommitFile string                        //GIT{}
		GitTag string                               //GIT{}
	*/
	GIT          GIT
	GitAction    string
	GitNewBranch string
	//-----docker flag ----------
	/*
		DockerLogin string
		Push bool
		PushPattern string
		PullPattern string
		Imagename string
		List int
		LatestMode string
		Stage string
	*/
	Docker myDocker.Docker
	//-------nexus flag ----------
	/*
		NexusApiMethod string
		NexusReqBody string
		NexusOutputPattern string
		NexusPromoteType string
		NexusPromoteDestination string
		NexusPromoteUrl string
		NexusPromoteSource string
	*/
	Nexus myNexus.Nexus
	//-------replace flag ----------
	/*
		ReplaceType string                         //REPLACEYAML{}
		ReplaceImage string                        //REPLACEYAML{}
		ReplacePattern string                      //REPLACEYAML{}
		ReplaceYamlType string                     //REPLACEYAML{}
		ReplaceValue string                        //REPLACEYAML{}
	*/
	REPLACEYAML REPLACEYAML
	//--------- kustomize flag ------------
	/*
		KustomBasePath string                      //KustomizeArgument{}
		KustomizeOitputDir string                  //KustomizeArgument{}
		KustomizeRelPath string                    //KustomizeArgument{}
		KustomizeUrlPattern string                 //KustomizeArgument{}
		KustomizeModule string                     //KustomizeArgument{}
		KustomizeOpenfaasModule string             //KustomizeArgument{}
		KustomizeCompare string                    //KustomizeArgument{}
		KustomizeBaseFolder string                 //KustomizeArgument{}

		EnvironmentFile string                     //KustomizeArgument{}
	*/
	KustomizeArgument kustomize.KustomizeArgument
	//----------kubernetes flag -------------
	Namespace       string
	SnapshotPattern string
	//-------------- Account flag -----------
	User     string
	Password string
	//--------------I/O flag --------------
	InputFile  string
	OutputFile string
	//--------------version flag -------------
	Version bool
}
