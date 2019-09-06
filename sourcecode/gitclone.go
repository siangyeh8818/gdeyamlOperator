package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

func GitClone(url string, branch string, directory string, name string, token string) {
	giterr := CloneYaml(url, branch, directory, name, token)
	if giterr != nil {
		tags := branch
		CloneYamlByTag(url, tags, directory, name, token)
	}
}

func CloneYaml(url string, branch string, directory string, name string, token string) error {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	//	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone -b %s --single-branch %s %s", branch, url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: name, // yes, this can be anything except an empty string
			Password: token,
		},
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		URL:           url,
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

func CloneYamlByTag(url string, tag string, directory string, name string, token string) error {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	//	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone -b %s --single-branch %s %s", tag, url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: name, // yes, this can be anything except an empty string
			Password: token,
		},
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/tags/%s", tag)),
		URL:           url,
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


