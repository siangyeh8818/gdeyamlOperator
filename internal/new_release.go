package  gdeyamloperator

import (
	"log"

	"gopkg.in/yaml.v2"
)

func NewRelease(git_url string, old_branch string, new_branch string, git_repo_path string, git_user string, git_token string, outputfilename string , g *GIT) {
	log.Println("-----action >> cloneRepo----")
	CloneRepo(git_url, old_branch, git_repo_path, git_user, git_token)
	log.Println("-----action >> CreateBranch----")
	CreateBranch(git_url, new_branch, git_repo_path)
	log.Println("-----action >> CheckoutBranch----")
	CheckoutBranch(git_url, new_branch, git_repo_path)

	log.Println("-----action >> Parser deploy.yml----")
	inputfile := git_repo_path + "/deploy.yml"
	deployyaml := K8sYaml{}
	deployyaml.GetConf(inputfile)
	log.Println("-----action >> Replace content about deploy.yml----")
	deployyaml.Deployment.UpdateBaseStructBranch(deployyaml.Deployment.BASE[0].Git, trimQuotes(new_branch))
	outputcontent, err := yaml.Marshal(&deployyaml)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Println("-----action >> Write output file----")
	WriteWithIoutil(inputfile, string(outputcontent))
	log.Println("-----action >> CommitRepo----")
	CommitRepo(g, "deploy.yml")
	log.Println("-----action >> PushGit----")
	PushGit(git_repo_path, git_user, git_token,new_branch,git_url)
	log.Println("-----action finishing----")
}

func trimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
