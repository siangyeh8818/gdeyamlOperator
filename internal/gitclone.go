package test

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func GitClone(g *GIT) {
	giterr := CloneYaml(g)
	if giterr != nil {
		tags := g.Branch
		CloneYamlByTag(g, tags)
	}
}

func CloneYaml(g *GIT) error {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	//	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone -b %s --single-branch %s %s", g.Branch, g.Url, g.Path)

	_, err := git.PlainClone(g.Path, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: g.AccessUser, // yes, this can be anything except an empty string
			Password: g.AccessToken,
		},
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", g.Branch)),
		URL:           g.Url,
		Progress:      os.Stdout,
		SingleBranch:  true,
	})
	//CheckIfError(err)
	// ... retrieving the branch being pointed by HEAD
	//ref, err := r.Head()
	//CheckIfError(err)
	// ... retrieving the commit object
	//commit, err := r.CommitObject(ref.Hash())
	//CheckIfError(err)

	//fmt.Println(commit)
	return err
}

func CloneYamlByTag(g *GIT, tag string) error {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	//	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone -b %s --single-branch %s %s", tag, g.Url, g.Path)

	r, err := git.PlainClone(g.Path, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: g.AccessUser, // yes, this can be anything except an empty string
			Password: g.AccessToken,
		},
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tag)),
		URL:           g.Url,
		Progress:      os.Stdout,
		SingleBranch:  true,
	})

	CheckIfError(err)
	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
	return err
}
