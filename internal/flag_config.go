package gdeyamloperator

type ACTION interface {
	InitConfig()
    RunAction() bool
}

type BINARYCONFIG struct {
	Action string
	//-----git flag -----------
	GitUrl string                               //GIT{}
	GitRepoPath  string                         //GIT{}
	GitUser  string                             //GIT{}
	GitToken string                             //GIT{}
	GitBranch string                            //GIT{}
	GitNewBranch string
	GitCommitFile string                        //GIT{}
	GitTag string                               //GIT{}
	GitAction string
	//-----docker flag ----------
	DockerLogin string
	Push bool
	PushPattern string
	PullPattern string
	Imagename string
	List int
	LatestMode string

	//-------nexus flag ----------
	NexusApiMethod string
	NexusReqBody string
	NexusOutputPattern string
	NexusPromoteType string
	NexusPromoteDestination string
	NexusPromoteUrl string
	NexusPromoteSource string

	//-------
	EnvironmentFile string
	Stage string

	//-------replace flag ----------
	ReplaceType string                         //REPLACEYAML{}
	ReplaceImage string                        //REPLACEYAML{}
	ReplacePattern string                      //REPLACEYAML{}
	ReplaceYamlType string                     //REPLACEYAML{}
	ReplaceValue string                        //REPLACEYAML{}

	//--------- kustomize flag ------------
	KustomBasePath string
	KustomizeOitputDir string
	KustomizeRelPath string
	KustomizeUrlPattern string
	KustomizeModule string
	KustomizeOpenfaasModule string
	KustomizeCompare string
	KustomizeBaseFolder string
	//----------kubernetes flag -------------
	Namespace string
	SnapshotPattern string

	//------- Account flag -----------
	User string
	Password string

	//---------I/O flag --------------
	InputFile string
	OutputFile string
	//--------------version flag -------------
    Version bool
}
