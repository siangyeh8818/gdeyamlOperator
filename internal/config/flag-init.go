package config

import (
	"flag"
)

func Init(binaryConfig *BINARYCONFIG) {

	//binaryConfig := BINARYCONFIG{}

	// init Action
	flag.StringVar(&binaryConfig.Action, "action", "gettag", "choose 'gettag' or 'snapshot' or 'promote' or 'gitclone' or 'replace' or 'imagedump' or 'nexus_api' or 'new-release' or 'kustomize' or 'argu-dump' or 'jenkins'")

	// init Git
	flag.StringVar(&binaryConfig.GIT.Url, "git-url", "", "url for git repo")
	flag.StringVar(&binaryConfig.GIT.Branch, "git-branch", "", "branch for git repo")
	flag.StringVar(&binaryConfig.GIT.Tag, "git-tag", "", "Tag for git repo")
	flag.StringVar(&binaryConfig.GIT.Path, "git-repo-path", "", "directory for git-repo")
	flag.StringVar(&binaryConfig.GIT.AccessUser, "git-user", "", "user for git clone")
	flag.StringVar(&binaryConfig.GIT.AccessToken, "git-token", "", "token for git clone")
	flag.StringVar(&binaryConfig.GIT.CommitFIle, "git-commit-file", "deploy.yml", "File name that you want to commit , default value is 'deploy.yml'")

	// Init GitAction
	flag.StringVar(&binaryConfig.GitAction, "git-action", "", "git related operation , such as 'branch','push'")

	// Init GitNewBranch
	flag.StringVar(&binaryConfig.GitNewBranch, "git-new-branch", "", "New branch for git repo, this branch will be created")

	// Init Docker
	flag.StringVar(&binaryConfig.Docker.DockerLogin, "docker-login", "", "DockerHub url/IP for docekr login")
	flag.BoolVar(&binaryConfig.Docker.Push, "push", false, "push this image , default is false")
	flag.StringVar(&binaryConfig.Docker.PushPattern, "push-pattern", "", "(push)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.StringVar(&binaryConfig.Docker.PullPattern, "pull-pattern", "", "(pull)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.StringVar(&binaryConfig.Docker.Imagename, "imagename", "dockerhub.pentium.network/grafana", "docker image , such as dockerhub.pentium.network/grafana")
	flag.IntVar(&binaryConfig.Docker.List, "list", 5, "After sort tag list , we only deal with these top'number tags ")
	flag.StringVar(&binaryConfig.Docker.LatestMode, "latest-mode", "push", "push or build , choose one mode to identify latest tag to you")
	flag.StringVar(&binaryConfig.Docker.Stage, "stage", "", "replace stage , new stage content")

	// Init Nexus
	flag.StringVar(&binaryConfig.Nexus.NexusApiMethod, "nexus-api-method", "", "Http method for NexusAPI Request, such as 'GET','POST','PUT','DELETE'")
	flag.StringVar(&binaryConfig.Nexus.NexusReqBody, "nexus-req-body", "", "Requets body for NexusAPI Request")
	flag.StringVar(&binaryConfig.Nexus.NexusOutputPattern, "nexus-output-pattern", "", "Pattern for output by requesting Nexus-API")
	flag.StringVar(&binaryConfig.Nexus.NexusPromoteType, "promote-type", "move", "Different model , 'move' or 'cp' ")
	flag.StringVar(&binaryConfig.Nexus.NexusPromoteDestination, "promote-destination", "", "Destination for repository name ")
	flag.StringVar(&binaryConfig.Nexus.NexusPromoteUrl, "promote-url", "", "destination for you promoting image url (nexus)'")
	flag.StringVar(&binaryConfig.Nexus.NexusPromoteSource, "promote-source", "", "sourece(Repository name) for you promoting image url (nexus)'")

	// Init REPLACEYAML
	flag.StringVar(&binaryConfig.REPLACEYAML.Type, "replace-type", "local", "you can choose 'local' or 'git'")
	flag.StringVar(&binaryConfig.REPLACEYAML.Image, "replace-image", "", "which one image-name you want to br replace")
	flag.StringVar(&binaryConfig.REPLACEYAML.Pattern, "replace-pattern", "", "pattern for release , for example : blcks:version")
	flag.StringVar(&binaryConfig.REPLACEYAML.YamlType, "replace-yaml-type", "deployyaml", "whinh yaml-type you want to deal with , 'deployyaml' or 'environmentyaml'")
	flag.StringVar(&binaryConfig.REPLACEYAML.NewValue, "replace-value", "", "value for pattern tou want to update")

	// Init KustomizeArgument

	flag.StringVar(&binaryConfig.KustomizeArgument.Outputdir, "kustomize-outputdir", "./overlays", "output data of")
	flag.StringVar(&binaryConfig.KustomizeArgument.Comparedata, "kustomize-compare", "./deploy.yml", "deploy data")
	flag.StringVar(&binaryConfig.KustomizeArgument.Namespace, "namespace", "default", "k8s namesapce , such as default")
	flag.StringVar(&binaryConfig.KustomizeArgument.RelPath, "kustomize-relpath", "../../", "relative path of current execution path and kustomize path")
	flag.StringVar(&binaryConfig.KustomizeArgument.K8sBaseloc, "kustom-base-path", "base", "could be {relPath}/{Baseloc}, default is ../../{Baseloc}")
	flag.StringVar(&binaryConfig.KustomizeArgument.OfBaseloc, "kustomize-basefolder", "base", "could be {relPath}/{Baseloc}, default is ../../{Baseloc}")
	flag.StringVar(&binaryConfig.KustomizeArgument.Kmodules, "kustomize-module", "", "k8s modules from command: module:image:stage:tag,module1,image1,stage1,tag1")
	flag.StringVar(&binaryConfig.KustomizeArgument.UrlPattern, "kustomize-urlpattern", "cr.pentium.network/{{image}}:{{tag}}", "define url pattern by {{stage}}, {{image}}, and {{tag}}")
	flag.StringVar(&binaryConfig.KustomizeArgument.EnvFile, "environment-file", "", "file path of environment.yml")
	//flag.StringVar(&omodules, "kustomize-openfaasmodule", "", "openfaas modules from command: module:image:stage:tag,module1,image1,stage1,tag1")

	// Init SnapshotPattern
	flag.StringVar(&binaryConfig.SnapshotPattern, "snapshot-pattern", "", "pattern fot output , such as : k8s:default,openfaas:openfaas-fn,monitor:monitor,redis:redis")

	// Init K8s-Action
	flag.StringVar(&binaryConfig.K8sClusterAction, "optype", "delete", "please choose a type of operation to perform")

	// Init User
	flag.StringVar(&binaryConfig.User, "user", "", "user for docker login")

	// Init Password
	flag.StringVar(&binaryConfig.Password, "password", "", "password for docker login")

	// Init InputFile
	flag.StringVar(&binaryConfig.InputFile, "inputfile", "", "input file name , such as deploy.yml")

	// Init OutputFile
	flag.StringVar(&binaryConfig.OutputFile, "ouputfile", "tmp_out.yml", "output file name , such as deploy-out.yml")

	// Init Version
	flag.BoolVar(&binaryConfig.Version, "v", false, "prints current binary version")

}

func Parse() {

	flag.Parse()
}
