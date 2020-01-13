package action

import (
	myConfig "github.com/siangyeh8818/gdeyamlOperator/internal/config"
	mygit "github.com/siangyeh8818/gdeyamlOperator/internal/git"
)

type GitActions struct {
}

func (action GitActions) DefineAction(config myConfig.BINARYCONFIG) {
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
}
