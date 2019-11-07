package gdeyamloperator

//import (
//	"flag"
//)

type BINARYCONFIG struct {
	Action string
	GitUrl string
	GitRepoPath  string
	GitUser  string
	GitToken string
	GitBranch string
	GitNewBranch string
	GitCommitFile string
	GitTag string
	GitAction string
	DockerLogin string
	Push bool
	PushPattern string
	PullPattern string
	Imagename string
	List int
	LatestMode string
	NexusApiMethod string
	NexusReqBody string
	NexusOutputPattern string
	NexusPromoteType string
	NexusPromoteDestination string
	NexusPromoteUrl string
	NexusPromoteSource string
	EnvironmentFile string
	SnapshotPattern string
	Stage string
	ReplaceType string
	ReplaceImage string
	ReplacePattern string
	ReplaceYamlType string
	ReplaceValue string
	KustomBasePath string
	KustomizeOitputDir string
	KustomizeRelPath string
	KustomizeUrlPattern string
	KustomizeModule string
	KustomizeOpenfaasModule string
	KustomizeCompare string
	KustomizeBaseFolder string
	Namespace string
	User string
	Password string
	InputFile string
	OutputFile string
    Version bool
}
