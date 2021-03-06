package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	clusterop "github.com/siangyeh8818/gdeyamlOperator/internal/clusterop"
	"gopkg.in/yaml.v2"

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
	myTool "github.com/siangyeh8818/gdeyamlOperator/internal/utility"
)

var completeimagename string
var list2 int
var inputfile string
var ouputfile string
var pushimage bool
var inputstage string
var loginuser string
var loginpassword string
var version bool
var latest_mode string
var push_pattern string
var pull_pattern string
var action string
var namespace string
var promote_url string
var promote_source string
var kustom_base string
var environment_file string
var git_url string

//var clone_path string
var git_user string
var git_token string
var git_branch string
var git_tag string
var snapshot_pattern string
var docker_login string
var nexus_api_method string
var nexus_req_body string
var nexus_output_pattern string
var promote_type string
var promote_destination string
var git_action string
var GitCommitFile string
var git_repo_path string
var git_new_branch string
var replace_type string
var replace_value string
var replace_image string
var replace_pattern string
var ReplaceYamlType string
var kmodules string
var relPath string
var outputdir string
var comparedata string
var omodules string
var UrlPattern string
var Baseloc string
var optype string

func main() {

	Init()
	flag.Parse()

	if version {
		fmt.Println("version : 1.12.1")
		os.Exit(0)
	}
	newgit := mygit.GIT{}
	(&newgit).UpdateGit(git_url, git_branch, git_tag, git_repo_path, git_user, git_token, GitCommitFile)

	kustomize_argument := myKustomize.KustomizeArgument{}
	(&kustomize_argument).UpdateKustomizeArgument(outputdir, comparedata, namespace, relPath, Baseloc, Baseloc, kmodules, UrlPattern, environment_file)

	replace_struct := myKustomize.REPLACEYAML{}
	(&replace_struct).UpdateREPLACEYAML(replace_type, replace_pattern, replace_image, replace_value, ReplaceYamlType)

	fmt.Println("--------------Main Action -----------------")
	fmt.Printf("flag -action: %s\n", action)
	fmt.Println("--------------Git Related Flag -----------------")
	fmt.Printf("flag -git-url: %s\n", git_url)
	//fmt.Printf("flag -clone-path: %s\n", clone_path)
	fmt.Printf("flag -git-repo-path: %s\n", git_repo_path)
	fmt.Printf("flag -git-user: %s\n", git_user)
	fmt.Printf("flag -git-token: %s\n", git_token)
	fmt.Printf("flag -git-branch: %s\n", git_branch)
	fmt.Printf("flag -git-new-branch: %s\n", git_new_branch)
	fmt.Printf("flag -git-tag: %s\n", git_tag)
	fmt.Printf("flag -git-action: %s\n", git_action)
	fmt.Printf("flag -git-commit-file: %s\n", GitCommitFile)
	fmt.Println("--------------Docker Related Flag -----------------")
	fmt.Printf("flag -docker-login: %s\n", docker_login)
	fmt.Printf("flag -push: %t\n", pushimage)
	fmt.Printf("flag -push-pattern: %s\n", push_pattern)
	fmt.Printf("flag -pull-pattern: %s\n", pull_pattern)
	fmt.Printf("flag -imagename: %s\n", completeimagename)
	fmt.Printf("flag -list: %d\n", list2)
	fmt.Printf("flag -latest-mode: %s\n", latest_mode)
	fmt.Println("--------------Nexus-API Related Flag -----------------")
	fmt.Printf("flag -nexus-api-method: %s\n", nexus_api_method)
	fmt.Printf("flag -nexus-req-body: %s\n", nexus_req_body)
	fmt.Printf("flag -nexus-output-pattern: %s\n", nexus_output_pattern)
	fmt.Printf("flag -promote-type: %s\n", promote_type)
	fmt.Printf("flag -promote-destination: %s\n", promote_destination)
	fmt.Printf("flag -promote-url: %s\n", promote_url)
	fmt.Printf("flag -promote-source: %s\n", promote_source)
	fmt.Println("--------------GDEyaml/kustomize Related Flag -----------------")
	fmt.Printf("flag -environment-file: %s\n", environment_file)
	fmt.Printf("flag -snapshot-pattern: %s\n", snapshot_pattern)
	fmt.Printf("flag -kustom-base-path: %s\n", kustom_base)
	fmt.Printf("flag -stage: %s\n", inputstage)
	fmt.Printf("flag -replace-type: %s\n", replace_type)
	fmt.Printf("flag -replace-yaml-type: %s\n", ReplaceYamlType)
	fmt.Printf("flag -replace-image: %s\n", replace_image)
	fmt.Printf("flag -replace-pattern: %s\n", replace_pattern)
	fmt.Printf("flag -replace-value: %s\n", replace_value)
	fmt.Printf("flag -kustomize-outputdir: %s\n", outputdir)
	fmt.Printf("flag -kustomize-relpath: %s\n", relPath)
	fmt.Printf("flag -kustomize-urlpattern: %s\n", UrlPattern)
	fmt.Printf("flag -kustomize-module: %s\n", kmodules)
	fmt.Printf("flag -kustomize-openfaasmodule: %s\n", omodules)
	fmt.Printf("flag -kustomize-compare: %s\n", comparedata)
	fmt.Printf("flag -kustomize-basefolder: %s\n", Baseloc)
	fmt.Println("--------------Kubernetes Related Flag -----------------")
	fmt.Printf("flag -namespace: %s\n", namespace)
	fmt.Println("--------------General Related Flag -----------------")
	fmt.Printf("flag -user: %s\n", loginuser)
	fmt.Printf("flag -password: %s\n", loginpassword)
	fmt.Printf("flag -inputfile: %s\n", inputfile)
	fmt.Printf("flag -ouputfile: %s\n", ouputfile)
	fmt.Println("--------------Version Related Flag -----------------")
	fmt.Printf("flag -v: %t\n", version)

	if loginuser != "" && loginpassword != "" {
		myDocker.LoginDockerHub(inputstage, loginuser, loginpassword)
	}

	switch action {
	case "gettag":
		if inputfile != "" && IO.Exists(inputfile) {
			inyaml := CustomStruct.K8sYaml{}
			inyaml.GetConf(inputfile)
			//fmt.Printf("input_YAML:\n%v\n\n", inyaml)
			//fmt.Println(ComposeImageName(inyaml.Deployment.K8S[0].Stage, inyaml.Deployment.K8S[0].Image, inyaml.Deployment.K8S[0].Tag))
			for i := 0; i < len(inyaml.Deployment.K8S); i++ {
				if inyaml.Deployment.K8S[i].Image != "" {
					fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)
					var tmp_cpmplete_imagename string
					if inputstage != "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(pull_pattern, inputstage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					} else if inputstage == "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(pull_pattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					}

					//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
					if pushimage == true {
						cmd_1 := "docker pull " + tmp_cpmplete_imagename
						fmt.Println(cmd_1)
						ShellCommand.RunCommand(cmd_1)
						push_cpmplete_imagename := myDocker.PatternParse(push_pattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
						//push_cpmplete_imagename := ComposeImageName(push_mode, new_push_hub, inputstage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
						cmd_2 := "docker tag " + tmp_cpmplete_imagename + " " + push_cpmplete_imagename
						fmt.Println(cmd_2)
						ShellCommand.RunCommand(cmd_2)
						cmd_3 := "docker push " + push_cpmplete_imagename
						fmt.Println(cmd_3)
						ShellCommand.RunCommand(cmd_3)
						cmd_4 := "docker rmi " + tmp_cpmplete_imagename
						fmt.Println(cmd_4)
						ShellCommand.RunCommand(cmd_4)
						cmd_5 := "docker rmi " + push_cpmplete_imagename
						fmt.Println(cmd_5)
						ShellCommand.RunCommand(cmd_5)
					}
					new_tag_latest := GetTag(tmp_cpmplete_imagename, latest_mode)
					new_tag_latest = strings.Trim(new_tag_latest, "\"")
					(&inyaml.Deployment.K8S[i]).UpdateK8sTag(new_tag_latest)
					fmt.Printf("new_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)

				} else {
					continue
				}

			}
			for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
				if inyaml.Deployment.Openfaas[i].Image != "" {
					fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.Openfaas[i].Tag)
					var tmp_cpmplete_imagename string

					if inputstage != "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(pull_pattern, inputstage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					} else if inputstage == "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(pull_pattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					}
					//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
					if pushimage == true {
						cmd_1 := "docker pull " + tmp_cpmplete_imagename
						fmt.Println(cmd_1)
						ShellCommand.RunCommand(cmd_1)
						push_cpmplete_imagename := myDocker.PatternParse(push_pattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
						cmd_2 := "docker tag " + tmp_cpmplete_imagename + " " + push_cpmplete_imagename
						fmt.Println(cmd_2)
						ShellCommand.RunCommand(cmd_2)
						cmd_3 := "docker push " + push_cpmplete_imagename
						fmt.Println(cmd_3)
						ShellCommand.RunCommand(cmd_3)
						cmd_4 := "docker rmi " + tmp_cpmplete_imagename
						fmt.Println(cmd_4)
						ShellCommand.RunCommand(cmd_4)
						cmd_5 := "docker rmi " + push_cpmplete_imagename
						fmt.Println(cmd_5)
						ShellCommand.RunCommand(cmd_5)

					}
					new_tag_latest := GetTag(tmp_cpmplete_imagename, latest_mode)
					new_tag_latest = strings.Trim(new_tag_latest, "\"")
					(&inyaml.Deployment.Openfaas[i]).UpdateOpenfaasTag(new_tag_latest)
					fmt.Printf("new_tag:\n%v\n\n", inyaml.Deployment.Openfaas[i].Tag)

				} else {
					continue
				}
			}
			d, err := yaml.Marshal(&inyaml)
			if err != nil {
				log.Fatalf("error: %v", err)
			}
			//	fmt.Printf("--- t dump:\n%s\n\n", string(d))

			IO.WriteWithIoutil(ouputfile, string(d))

		} else {
			new_tag_latest := GetTag(completeimagename, latest_mode)
			fmt.Println(new_tag_latest)
			new_tag_latest = strings.Trim(new_tag_latest, "\"")
			IO.WriteWithIoutil("getImageLatestTag_result.txt", new_tag_latest)
		}

	case "snapshot":
		myK8s.Snapshot(snapshot_pattern, ouputfile, kustom_base, git_branch)
	case "nexus_api":
		var output myJson.OutputContent
		switch nexus_api_method {
		case "GET":
			myNexus.GET_NesusAPI(promote_url, loginuser, loginpassword, ouputfile, nexus_output_pattern, &output)
		case "Get":
			myNexus.GET_NesusAPI(promote_url, loginuser, loginpassword, ouputfile, nexus_output_pattern, &output)
		case "get":
			myNexus.GET_NesusAPI(promote_url, loginuser, loginpassword, ouputfile, nexus_output_pattern, &output)
		case "POST":
			myNexus.POST_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "Post":
			myNexus.POST_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "post":
			myNexus.POST_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "PUT":
			myNexus.PUT_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "Put":
			myNexus.PUT_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "put":
			myNexus.PUT_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "DELETE":
			myNexus.DELETE_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "Delete":
			myNexus.DELETE_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		case "delete":
			myNexus.DELETE_NesusAPI(promote_url, loginuser, loginpassword, nexus_req_body)
		}

	case "promote":
		switch promote_type {
		case "move":
			if inputfile != "" && IO.Exists(inputfile) {
				inyaml := CustomStruct.K8sYaml{}
				inyaml.GetConf(inputfile)
				for i := 0; i < len(inyaml.Deployment.K8S); i++ {
					if inyaml.Deployment.K8S[i].Image != "" && inyaml.Deployment.K8S[i].Tag != "" {
						myNexus.Promoteimage(promote_url, promote_source, loginuser, loginpassword, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					}
				}
				for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
					if inyaml.Deployment.Openfaas[i].Image != "" && inyaml.Deployment.Openfaas[i].Tag != "" {
						myNexus.Promoteimage(promote_url, promote_source, loginuser, loginpassword, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					}
				}
			} else {
				fmt.Println("Yoy have to setting -inputfile <filename>")
			}
		case "cp":
			myNexus.Cpcomponetname(promote_url, loginuser, loginpassword, promote_destination)
		}

	case "gitclone":
		deployyamlgit := mygit.GIT{}
		configurationgit := mygit.GIT{}
		baseyamlgit := mygit.GIT{}
		if git_url != "" && environment_file == "" {
			if git_branch != "" && git_tag == "" {
				//GitClone(git_url, git_branch, git_repo_path, git_user, git_token)
				mygit.GitClone(&newgit)
			} else if git_branch == "" && git_tag != "" {
				mygit.GitClone(&newgit)
			} else if git_branch != "" && git_tag != "" {
				fmt.Println("Only one flag that you have to setting (git-branch or git-tag)")
				fmt.Println("While you setting git-branch , you can't set git-tag")
				os.Exit(0)
			}

		} else if git_url == "" && environment_file != "" {
			envir_yaml := CustomStruct.Environmentyaml{}
			envir_yaml.GetConf(environment_file)
			log.Println("Used environment file to collect information of git")

			if len(envir_yaml.Configuration) > 0 {
				(&configurationgit).UpdateGitUrl(envir_yaml.Configuration[0].Git)
				(&configurationgit).UpdateGitBranch(envir_yaml.Configuration[0].Branch)
				(&configurationgit).UpdateGitPath("configuration")
				(&configurationgit).UpdateGitAccessUser(git_user)
				(&configurationgit).UpdateGitAccessToken(git_token)
				mygit.GitClone(&configurationgit)
				//CloneYaml(envir_yaml.Configuration[0].Git, envir_yaml.Configuration[0].Branch, "configuration", git_user, git_token)
			}
			if len(envir_yaml.Deploymentfile) > 0 {
				(&deployyamlgit).UpdateGitUrl(envir_yaml.Deploymentfile[0].Git)
				(&deployyamlgit).UpdateGitBranch(envir_yaml.Deploymentfile[0].Branch)
				(&deployyamlgit).UpdateGitPath("deploymentfile")
				(&deployyamlgit).UpdateGitAccessUser(git_user)
				(&deployyamlgit).UpdateGitAccessToken(git_token)
				mygit.GitClone(&deployyamlgit)
				//CloneYaml(envir_yaml.Deploymentfile[0].Git, envir_yaml.Deploymentfile[0].Branch, "deploymentfile", git_user, git_token)
			}

		} else if git_url == "" && environment_file == "" && inputfile != "" {
			if inputfile != "" && IO.Exists(inputfile) {
				inyaml := CustomStruct.K8sYaml{}
				inyaml.GetConf(inputfile)
				if len(inyaml.Deployment.BASE) > 0 {
					(&baseyamlgit).UpdateGitUrl(inyaml.Deployment.BASE[0].Git)
					(&baseyamlgit).UpdateGitBranch(inyaml.Deployment.BASE[0].Branch)
					(&baseyamlgit).UpdateGitPath("base")
					(&baseyamlgit).UpdateGitAccessUser(git_user)
					(&baseyamlgit).UpdateGitAccessToken(git_token)
					mygit.GitClone(&baseyamlgit)
					//CloneYaml(inyaml.Deployment.BASE[0].Git, inyaml.Deployment.BASE[0].Branch, "base", git_user, git_token)
				}

			} else {
				fmt.Printf("%s is not exists !!!!", inputfile)
			}
		} else if git_url != "" && environment_file != "" {
			fmt.Println("only one flag you can setting , 'git-url' or 'environment-file'")
			os.Exit(0)
		}
	case "git":
		switch git_action {
		case "clone":
			mygit.CloneRepo(git_url, git_branch, git_repo_path, git_user, git_token)
		case "branch":
			mygit.CreateBranch(git_url, git_branch, git_repo_path)
		case "checkout":
			mygit.CheckoutBranch(git_url, git_branch, git_repo_path)
		case "push":
			mygit.PushGit(git_repo_path, git_user, git_token, git_new_branch, git_url)
		case "clonepush_new-branch":
			mygit.ClonePushNewBranch(git_url, git_branch, git_new_branch, git_repo_path, git_user, git_token)
		}
	case "replace":
		switch replace_type {
		case "local":
			if environment_file != "" && IO.Exists(environment_file) {
				if inputfile != "" && IO.Exists(inputfile) {
					if ouputfile != "" {
						fmt.Println("success to enter func Replacedeploymentfile")
						myKustomize.Replacedeploymentfile(environment_file, inputfile, ouputfile)
					} else if ouputfile == "" {
						fmt.Println("you have to  setting  flag (ouputfile)")
						os.Exit(0)
					}
				} else if inputfile == "" {
					fmt.Println("you have to  setting  flag (inputfile)")
					os.Exit(0)
				}
			} else if environment_file == "" {
				fmt.Println("you have to  setting  flag (environment_file)")
				os.Exit(0)
			}
		case "git":
			if git_url == "" {
				log.Println("you have to  setting  flag (git-url)")
				os.Exit(0)
			}
			if git_branch == "" {
				log.Println("you have to  setting  flag (git-branch)")
				os.Exit(0)
			}
			log.Println("-----action >> git CloneRepo----")
			mygit.GitClone(&newgit)
			log.Println("-----action >> Update Image-Tag to deploy.yml----")
			myKustomize.ReplacedeByPattern(&replace_struct, inputfile, ouputfile)
			//Replacedeploymentfile_Image_Tag(&replace_struct, inputfile, ouputfile)
			log.Println("-----action >> git CommitRepo----")
			mygit.CommitRepo(&newgit, inputfile)
			log.Println("-----action >> git PushRepo----")
			mygit.PushGit(git_repo_path, git_user, git_token, git_branch, git_url)
			log.Println("-----action finishing----")
		case "git-patch":
			if git_url == "" {
				log.Println("you have to  setting  flag (git-url)")
				os.Exit(0)
			}
			if git_branch == "" {
				log.Println("you have to  setting  flag (git-branch)")
				os.Exit(0)
			}
			log.Println("-----action >> git CloneRepo----")
			mygit.GitClone(&newgit)
			log.Println("-----action >> Update Image Info to deploy.yml----")
			myKustomize.PatchDeployFile(&replace_struct, inputfile, ouputfile, &kustomize_argument)
			log.Println("-----action >> git CommitRepo----")
			mygit.CommitRepo(&newgit, inputfile)
			log.Println("-----action >> git PushRepo----")
			mygit.PushGit(git_repo_path, git_user, git_token, git_branch, git_url)
			log.Println("-----action finishing----")

		}

	case "new-release":
		mygit.NewRelease(git_url, git_branch, git_new_branch, git_repo_path, git_user, git_token, ouputfile, &newgit)
	case "imagedump":
		myDocker.LoginDockerHubNew(docker_login, loginuser, loginpassword)
		myK8s.DumpImage(push_pattern, snapshot_pattern, pushimage)
	case "kustomize":
		myKustomize.OutputOverlays(&kustomize_argument, inputfile)
		//OutputOverlays(environment_file, inputfile, namespace, kmodules, relPath, k8sBaseloc)
	case "argu-dump":
		IO.DumpArguments(inputfile, environment_file, ouputfile)
	case "jenkins":
		myJenkins.INit_Jenkins()
	case "group-file":
		mygit.GroupNexusOutput(inputfile, ouputfile, &newgit)
	case "cluster-op":
		switch optype {
		case "add":
			clusterop.AddResource()
			break
		case "patch":
			clusterop.PatchResource()
			break
		case "delete":
			clusterop.DeleteResources(&newgit)
			break
		case "deploy-scripts":
			myK8s.CeateScriptsJob(inputfile, environment_file)

		}
	}

}

func Init() {
	flag.StringVar(&completeimagename, "imagename", "dockerhub.pentium.network/grafana", "docker image , such as dockerhub.pentium.network/grafana")
	flag.StringVar(&namespace, "namespace", "default", "k8s namesapce , such as default")
	flag.StringVar(&loginuser, "user", "", "user for docker login")
	flag.StringVar(&loginpassword, "password", "", "password for docker login")
	flag.IntVar(&list2, "list", 5, "After sort tag list , we only deal with these top'number tags ")
	flag.StringVar(&inputfile, "inputfile", "", "input file name , such as deploy.yml")
	flag.StringVar(&ouputfile, "ouputfile", "tmp_out.yml", "output file name , such as deploy-out.yml")
	flag.StringVar(&inputstage, "stage", "", "replace stage , new stage content")
	flag.StringVar(&latest_mode, "latest-mode", "push", "push or build , choose one mode to identify latest tag to you")
	flag.StringVar(&push_pattern, "push-pattern", "", "(push)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.StringVar(&pull_pattern, "pull-pattern", "", "(pull)pattern for imagename , ex: cr-{{stage}}.pentium.network/{{image}}:{{tag}}")
	flag.StringVar(&action, "action", "gettag", "choose 'gettag' or 'snapshot' or 'promote' or 'gitclone' or 'replace' or 'imagedump' or 'nexus_api' or 'new-release' or 'kustomize' or 'argu-dump' or 'jenkins'")
	flag.StringVar(&git_url, "git-url", "", "url for git repo")
	flag.StringVar(&git_branch, "git-branch", "", "branch for git repo")
	flag.StringVar(&git_new_branch, "git-new-branch", "", "New branch for git repo, this branch will be created")
	flag.StringVar(&git_action, "git-action", "", "git related operation , such as 'branch','push'")
	flag.StringVar(&git_tag, "git-tag", "", "Tag for git repo")
	flag.StringVar(&git_user, "git-user", "", "user for git clone")
	flag.StringVar(&git_token, "git-token", "", "token for git clone")
	flag.StringVar(&GitCommitFile, "git-commit-file", "deploy.yml", "File name that you want to commit , default value is 'deploy.yml'")
	flag.StringVar(&environment_file, "environment-file", "", "file path of environment.yml")
	flag.StringVar(&promote_url, "promote-url", "", "destination for you promoting image url (nexus)'")
	flag.StringVar(&promote_source, "promote-source", "", "sourece(Repository name) for you promoting image url (nexus)'")
	flag.StringVar(&kustom_base, "kustom-base-path", "", "folder path for your base yaml of kustomization'")
	flag.BoolVar(&pushimage, "push", false, "push this image , default is false")
	flag.BoolVar(&version, "v", false, "prints current binary version")
	flag.StringVar(&snapshot_pattern, "snapshot-pattern", "", "pattern fot output , such as : k8s:default,openfaas:openfaas-fn,monitor:monitor,redis:redis")
	flag.StringVar(&docker_login, "docker-login", "", "DockerHub url/IP for docekr login")
	flag.StringVar(&nexus_api_method, "nexus-api-method", "", "Http method for NexusAPI Request, such as 'GET','POST','PUT','DELETE'")
	flag.StringVar(&nexus_req_body, "nexus-req-body", "", "Requets body for NexusAPI Request")
	flag.StringVar(&nexus_output_pattern, "nexus-output-pattern", "", "Pattern for output by requesting Nexus-API")
	flag.StringVar(&promote_type, "promote-type", "move", "Different model , 'move' or 'cp' ")
	flag.StringVar(&promote_destination, "promote-destination", "", "Destination for repository name ")
	flag.StringVar(&git_repo_path, "git-repo-path", "", "directory for git-repo")
	flag.StringVar(&replace_type, "replace-type", "local", "you can choose 'local' or 'git'")
	flag.StringVar(&ReplaceYamlType, "replace-yaml-type", "deployyaml", "whinh yaml-type you want to deal with , 'deployyaml' or 'environmentyaml'")
	flag.StringVar(&replace_image, "replace-image", "", "which one image-name you want to br replace")
	flag.StringVar(&kmodules, "kustomize-module", "", "k8s modules from command: module:image:stage:tag,module1,image1,stage1,tag1")
	flag.StringVar(&omodules, "kustomize-openfaasmodule", "", "openfaas modules from command: module:image:stage:tag,module1,image1,stage1,tag1")
	flag.StringVar(&UrlPattern, "kustomize-urlpattern", "cr.pentium.network/{{image}}:{{tag}}", "define url pattern by {{stage}}, {{image}}, and {{tag}}")
	flag.StringVar(&outputdir, "kustomize-outputdir", "./overlays", "output data of")
	flag.StringVar(&comparedata, "kustomize-compare", "./deploy.yml", "deploy data")
	flag.StringVar(&relPath, "kustomize-relpath", "../../", "relative path of current execution path and kustomize path")
	flag.StringVar(&Baseloc, "kustomize-basefolder", "base", "could be {relPath}/{Baseloc}, default is ../../{Baseloc}")
	flag.StringVar(&replace_pattern, "replace-pattern", "", "pattern for release , for example : blcks:version")
	flag.StringVar(&replace_value, "replace-value", "", "value for pattern tou want to update")
	flag.StringVar(&optype, "optype", "delete", "please choose a type of operation to perform")
}

func GetTag(name string, latestmode string) string {
	raw_image_hub, raw_image_name := myTool.ImagenameSplit(name)

	var tag_result string
	var time_latest = "2000-01-01T00:00:00.508640172Z"
	var tag_latest string
	var querylistcmd string
	var loop_break_count int

	querylistcmd = "curl -X GET https://" + raw_image_hub + "/v2/" + raw_image_name + "/tags/list -s| jq -r .tags"
	//fmt.Println(querylistcmd)
	//	fmt.Println("------------------")
	//tag_result, _ = exec_shell("curl -X GET https://dockerhub.pentium.network/v2/grafana/tags/list| jq -r .tags")

	tag_result = ShellCommand.RunCommand(querylistcmd)
	tag_result = strings.Replace(tag_result, "[", "", 1)
	tag_result = strings.Replace(tag_result, "]", "", 1)
	tag_result = DeleteExtraSpace(tag_result)
	tag_result = strings.Replace(tag_result, "\n", "", -1)
	tagssplit := strings.Split(tag_result, ",")

	//fmt.Printf("Ints %v\n", tagssplit)
	reverse_tagssplit := reverseInts(tagssplit)
	//fmt.Printf("Reversed: %v\n", reverse_tagssplit)
	//	fmt.Println("Amount of image tag : " + strconv.Itoa(len(tagssplit)))
	imagemap := make(map[string]string, len(reverse_tagssplit))
	if latestmode == "build" {
		for i := range reverse_tagssplit {

			time := QueryLatestTag(reverse_tagssplit[i], raw_image_name, raw_image_hub)
			fmt.Println(reverse_tagssplit[i] + ":" + time)
			time = strings.Replace(time, "\n", "", -1)

			if strings.Compare(strings.Trim(reverse_tagssplit[i], "\""), "latest") == -1 {
				imagemap[reverse_tagssplit[i]] = time
				time_latest = SelectLatestTime(time, time_latest)
				if time_latest == imagemap[reverse_tagssplit[i]] {
					tag_latest = reverse_tagssplit[i]
				}
			}
			loop_break_count++
			if loop_break_count >= list2 {
				break
			}
		}
	} else if latestmode == "push" {
		tag_latest = reverse_tagssplit[0]
	}

	return tag_latest
}

func SelectLatestTime(t1 string, t2 string) string {
	var earlytime string
	time1, _ := time.Parse(time.RFC3339Nano, t1)
	//fmt.Println(time1)
	time2, _ := time.Parse(time.RFC3339Nano, t2)
	//fmt.Println(time2)
	if time2.After(time1) {
		earlytime = t2
		//	fmt.Println("time2 is winner")
	} else if time1.After(time2) {
		earlytime = t1
		//	fmt.Println("time1 is winner")
	}
	//fmt.Println("function SelectLatestTime result : " + earlytime)

	return earlytime
}

func DeleteExtraSpace(s string) string {
	s1 := strings.Replace(s, "  ", " ", -1)
	regstr := "\\s{2,}"
	reg, _ := regexp.Compile(regstr)
	s2 := make([]byte, len(s1))
	copy(s2, s1)
	spc_index := reg.FindStringIndex(string(s2))
	for len(spc_index) > 0 {
		s2 = append(s2[:spc_index[0]+1], s2[spc_index[1]:]...)
		spc_index = reg.FindStringIndex(string(s2))
	}
	return string(s2)
}

func QueryLatestTag(tag string, imgname string, hub string) string {

	curltagresult := ShellCommand.RunCommand("curl -X GET https://" + hub + "/v2/" + imgname + "/manifests/" + tag + " | jq -r '.history[].v1Compatibility' | jq '.created' | sort | sed 's/\"//g'|tail -n1 ")

	return curltagresult
}

func reverseInts(input []string) []string {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}
