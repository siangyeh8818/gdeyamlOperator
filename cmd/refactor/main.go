package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	clusterop "github.com/siangyeh8818/gdeyamlOperator/internal/clusterop"
	"gopkg.in/yaml.v2"

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
)

func main() {

	config := myConfig.BINARYCONFIG{}
	myConfig.Init(&config)
	myConfig.Parse()

	if config.Version {
		fmt.Println("version : 1.12.5")
		os.Exit(0)
	}

	if config.User != "" && config.Password != "" {
		myDocker.LoginDockerHub(config.Docker.Stage, config.User, config.Password)
	}

	switch config.Action {
	case "gettag":
		if config.InputFile != "" && IO.Exists(config.InputFile) {
			inyaml := CustomStruct.K8sYaml{}
			inyaml.GetConf(config.InputFile)
			//fmt.Printf("input_YAML:\n%v\n\n", inyaml)
			//fmt.Println(ComposeImageName(inyaml.Deployment.K8S[0].Stage, inyaml.Deployment.K8S[0].Image, inyaml.Deployment.K8S[0].Tag))
			for i := 0; i < len(inyaml.Deployment.K8S); i++ {
				if inyaml.Deployment.K8S[i].Image != "" {
					fmt.Printf("old_tag:\n%v\n\n", inyaml.Deployment.K8S[i].Tag)
					var tmp_cpmplete_imagename string
					if config.Docker.Stage != "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(config.Docker.PullPattern, config.Docker.Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					} else if config.Docker.Stage == "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(config.Docker.PullPattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					}

					//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
					if config.Docker.Push == true {
						cmd_1 := "docker pull " + tmp_cpmplete_imagename
						fmt.Println(cmd_1)
						ShellCommand.RunCommand(cmd_1)
						push_cpmplete_imagename := myDocker.PatternParse(config.Docker.PushPattern, inyaml.Deployment.K8S[i].Stage, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
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
					new_tag_latest := myDocker.GetTag(tmp_cpmplete_imagename, config.Docker.LatestMode, config.Docker.List)
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

					if config.Docker.Stage != "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(config.Docker.PullPattern, config.Docker.Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					} else if config.Docker.Stage == "" {
						tmp_cpmplete_imagename = myDocker.PatternParse(config.Docker.PullPattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					}
					//tmp_cpmplete_imagename := ComposeImageName(query_mode, new_hub, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					fmt.Printf("complete image name : %s\n", tmp_cpmplete_imagename)
					if config.Docker.Push == true {
						cmd_1 := "docker pull " + tmp_cpmplete_imagename
						fmt.Println(cmd_1)
						ShellCommand.RunCommand(cmd_1)
						push_cpmplete_imagename := myDocker.PatternParse(config.Docker.PushPattern, inyaml.Deployment.Openfaas[i].Stage, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
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
					new_tag_latest := myDocker.GetTag(tmp_cpmplete_imagename, config.Docker.LatestMode, config.Docker.List)
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

			IO.WriteWithIoutil(config.OutputFile, string(d))

		} else {
			new_tag_latest := myDocker.GetTag(config.Docker.Imagename, config.Docker.LatestMode, config.Docker.List)
			fmt.Println(new_tag_latest)
			new_tag_latest = strings.Trim(new_tag_latest, "\"")
			IO.WriteWithIoutil("getImageLatestTag_result.txt", new_tag_latest)
		}

	case "snapshot":
		myK8s.Snapshot(config.SnapshotPattern, config.OutputFile, config.KustomizeArgument.K8sBaseloc, config.GIT.Branch)
	case "nexus_api":
		var output myJson.OutputContent

		apiMethod := strings.ToLower(config.Nexus.NexusApiMethod)
		switch apiMethod {
		case "get":
			myNexus.GET_NesusAPI(config.Nexus.NexusPromoteUrl, config.User, config.Password, config.OutputFile, config.Nexus.NexusOutputPattern, &output)
		case "post":
			myNexus.POST_NesusAPI(config.Nexus.NexusPromoteUrl, config.User, config.Password, config.Nexus.NexusReqBody)
		case "put":
			myNexus.PUT_NesusAPI(config.Nexus.NexusPromoteUrl, config.User, config.Password, config.Nexus.NexusReqBody)
		case "delete":
			myNexus.DELETE_NesusAPI(config.Nexus.NexusPromoteUrl, config.User, config.Password, config.Nexus.NexusReqBody)
		}

	case "promote":
		switch config.Nexus.NexusPromoteType {
		case "move":
			if config.InputFile != "" && IO.Exists(config.InputFile) {
				inyaml := CustomStruct.K8sYaml{}
				inyaml.GetConf(config.InputFile)
				for i := 0; i < len(inyaml.Deployment.K8S); i++ {
					if inyaml.Deployment.K8S[i].Image != "" && inyaml.Deployment.K8S[i].Tag != "" {
						myNexus.Promoteimage(config.Nexus.NexusPromoteUrl, config.Nexus.NexusPromoteSource, config.User, config.Password, inyaml.Deployment.K8S[i].Image, inyaml.Deployment.K8S[i].Tag)
					}
				}
				for i := 0; i < len(inyaml.Deployment.Openfaas); i++ {
					if inyaml.Deployment.Openfaas[i].Image != "" && inyaml.Deployment.Openfaas[i].Tag != "" {
						myNexus.Promoteimage(config.Nexus.NexusPromoteUrl, config.Nexus.NexusPromoteSource, config.User, config.Password, inyaml.Deployment.Openfaas[i].Image, inyaml.Deployment.Openfaas[i].Tag)
					}
				}
			} else {
				fmt.Println("Yoy have to setting -inputfile <filename>")
			}
		case "cp":
			myNexus.Cpcomponetname(config.Nexus.NexusPromoteUrl, config.User, config.Password, config.Nexus.NexusPromoteDestination)
		}

	case "gitclone":
		deployyamlgit := mygit.GIT{}
		configurationgit := mygit.GIT{}
		baseyamlgit := mygit.GIT{}
		if config.GIT.Url != "" && config.KustomizeArgument.EnvFile == "" {
			if config.GIT.Branch != "" && config.GIT.Tag == "" {
				//GitClone(git_url, git_branch, git_repo_path, git_user, git_token)
				mygit.GitClone(&config.GIT)
			} else if config.GIT.Branch == "" && config.GIT.Tag != "" {
				mygit.GitClone(&config.GIT)
			} else if config.GIT.Branch != "" && config.GIT.Tag != "" {
				fmt.Println("Only one flag that you have to setting (git-branch or git-tag)")
				fmt.Println("While you setting git-branch , you can't set git-tag")
				os.Exit(0)
			}

		} else if config.GIT.Url == "" && config.KustomizeArgument.EnvFile != "" {
			envir_yaml := CustomStruct.Environmentyaml{}
			envir_yaml.GetConf(config.KustomizeArgument.EnvFile)
			log.Println("Used environment file to collect information of git")

			if len(envir_yaml.Configuration) > 0 {
				(&configurationgit).UpdateGitUrl(envir_yaml.Configuration[0].Git)
				(&configurationgit).UpdateGitBranch(envir_yaml.Configuration[0].Branch)
				(&configurationgit).UpdateGitPath("configuration")
				(&configurationgit).UpdateGitAccessUser(config.GIT.AccessUser)
				(&configurationgit).UpdateGitAccessToken(config.GIT.AccessToken)
				mygit.GitClone(&configurationgit)
				//CloneYaml(envir_yaml.Configuration[0].Git, envir_yaml.Configuration[0].Branch, "configuration", git_user, git_token)
			}
			if len(envir_yaml.Deploymentfile) > 0 {
				(&deployyamlgit).UpdateGitUrl(envir_yaml.Deploymentfile[0].Git)
				(&deployyamlgit).UpdateGitBranch(envir_yaml.Deploymentfile[0].Branch)
				(&deployyamlgit).UpdateGitPath("deploymentfile")
				(&deployyamlgit).UpdateGitAccessUser(config.GIT.AccessUser)
				(&deployyamlgit).UpdateGitAccessToken(config.GIT.AccessToken)
				mygit.GitClone(&deployyamlgit)
				//CloneYaml(envir_yaml.Deploymentfile[0].Git, envir_yaml.Deploymentfile[0].Branch, "deploymentfile", git_user, git_token)
			}

		} else if config.GIT.Url == "" && config.KustomizeArgument.EnvFile == "" && config.InputFile != "" {
			if config.InputFile != "" && IO.Exists(config.InputFile) {
				inyaml := CustomStruct.K8sYaml{}
				inyaml.GetConf(config.InputFile)
				if len(inyaml.Deployment.BASE) > 0 {
					(&baseyamlgit).UpdateGitUrl(inyaml.Deployment.BASE[0].Git)
					(&baseyamlgit).UpdateGitBranch(inyaml.Deployment.BASE[0].Branch)
					(&baseyamlgit).UpdateGitPath("base")
					(&baseyamlgit).UpdateGitAccessUser(config.GIT.AccessUser)
					(&baseyamlgit).UpdateGitAccessToken(config.GIT.AccessToken)
					mygit.GitClone(&baseyamlgit)
					//CloneYaml(inyaml.Deployment.BASE[0].Git, inyaml.Deployment.BASE[0].Branch, "base", git_user, git_token)
				}

			} else {
				fmt.Printf("%s is not exists !!!!", config.InputFile)
			}
		} else if config.GIT.Url != "" && config.KustomizeArgument.EnvFile != "" {
			fmt.Println("only one flag you can setting , 'git-url' or 'environment-file'")
			os.Exit(0)
		}
	case "git":
		switch config.GitAction {
		case "clone":
			mygit.CloneRepo(config.GIT.Url, config.GIT.Branch, config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken)
		case "branch":
			mygit.CreateBranch(config.GIT.Url, config.GIT.Branch, config.GIT.Path)
		case "checkout":
			mygit.CheckoutBranch(config.GIT.Url, config.GIT.Branch, config.GIT.Path)
		case "push":
			mygit.PushGit(config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken, config.GitNewBranch, config.GIT.Url)
		case "clonepush_new-branch":
			mygit.ClonePushNewBranch(config.GIT.Url, config.GIT.Branch, config.GitNewBranch, config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken)
		}
	case "replace":
		switch config.REPLACEYAML.Type {
		case "local":
			if config.KustomizeArgument.EnvFile != "" && IO.Exists(config.KustomizeArgument.EnvFile) {
				if config.InputFile != "" && IO.Exists(config.InputFile) {
					if config.OutputFile != "" {
						fmt.Println("success to enter func Replacedeploymentfile")
						myKustomize.Replacedeploymentfile(config.KustomizeArgument.EnvFile, config.InputFile, config.OutputFile)
					} else if config.OutputFile == "" {
						fmt.Println("you have to  setting  flag (ouputfile)")
						os.Exit(0)
					}
				} else if config.InputFile == "" {
					fmt.Println("you have to  setting  flag (inputfile)")
					os.Exit(0)
				}
			} else if config.KustomizeArgument.EnvFile == "" {
				fmt.Println("you have to  setting  flag (environment_file)")
				os.Exit(0)
			}
		case "git":
			if config.GIT.Url == "" {
				log.Println("you have to  setting  flag (git-url)")
				os.Exit(0)
			}
			if config.GIT.Branch == "" {
				log.Println("you have to  setting  flag (git-branch)")
				os.Exit(0)
			}
			log.Println("-----action >> git CloneRepo----")
			mygit.GitClone(&config.GIT)
			log.Println("-----action >> Update Image-Tag to deploy.yml----")
			myKustomize.ReplacedeByPattern(&config.REPLACEYAML, config.InputFile, config.OutputFile)
			//Replacedeploymentfile_Image_Tag(&replace_struct, inputfile, ouputfile)
			log.Println("-----action >> git CommitRepo----")
			mygit.CommitRepo(&config.GIT, config.InputFile)
			log.Println("-----action >> git PushRepo----")
			mygit.PushGit(config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken, config.GIT.Branch, config.GIT.Url)
			log.Println("-----action finishing----")
		case "git-patch":
			if config.GIT.Url == "" {
				log.Println("you have to  setting  flag (git-url)")
				os.Exit(0)
			}
			if config.GIT.Branch == "" {
				log.Println("you have to  setting  flag (git-branch)")
				os.Exit(0)
			}
			log.Println("-----action >> git CloneRepo----")
			mygit.GitClone(&config.GIT)
			log.Println("-----action >> Update Image Info to deploy.yml----")
			myKustomize.PatchDeployFile(&config.REPLACEYAML, config.InputFile, config.OutputFile, &config.KustomizeArgument)
			log.Println("-----action >> git CommitRepo----")
			mygit.CommitRepo(&config.GIT, config.InputFile)
			log.Println("-----action >> git PushRepo----")
			mygit.PushGit(config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken, config.GIT.Branch, config.GIT.Url)
			log.Println("-----action finishing----")

		}

	case "new-release":
		mygit.NewRelease(config.GIT.Url, config.GIT.Branch, config.GitNewBranch, config.GIT.Path, config.GIT.AccessUser, config.GIT.AccessToken, config.OutputFile, &config.GIT)
	case "imagedump":
		myDocker.LoginDockerHubNew(config.Docker.DockerLogin, config.User, config.Password)
		myK8s.DumpImage(config.Docker.PushPattern, config.SnapshotPattern, config.Docker.Push)
	case "kustomize":
		myKustomize.OutputOverlays(&config.KustomizeArgument, config.InputFile)
		//OutputOverlays(environment_file, inputfile, namespace, kmodules, relPath, k8sBaseloc)
	case "argu-dump":
		IO.DumpArguments(config.InputFile, config.KustomizeArgument.EnvFile, config.OutputFile)
	case "jenkins":
		myJenkins.INit_Jenkins()
	case "group-file":
		mygit.GroupNexusOutput(config.InputFile, config.OutputFile, &config.GIT)
	case "cluster-op":
		switch config.K8sClusterAction {
		case "add":
			clusterop.AddResource()
			break
		case "patch":
			clusterop.PatchResource()
			break
		case "delete":
			clusterop.DeleteResources(&config.GIT)
			break
		case "deploy-scripts":
			myK8s.CeateScriptsJob(config.InputFile, config.KustomizeArgument.EnvFile)
		}
	}

}
