package test

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
)

type GIT struct {
	Url         string
	Branch      string
	Tag         string
	Path        string
	AccessUser  string
	AccessToken string
}

func (g *GIT) UpdateGit(url string, branch string, tag string, path string, user string, token string) {
	g.Url = url
	g.Branch = branch
	g.Tag = tag
	g.Path = path
	g.AccessUser = user
	g.AccessToken = token
}
func (g *GIT) UpdateGitUrl(url string) {
	g.Url = url
}
func (g *GIT) UpdateGitBranch(branch string) {
	g.Branch = branch
}
func (g *GIT) UpdateGitTag(tag string) {
	g.Tag = tag
}
func (g *GIT) UpdateGitPath(path string) {
	g.Path = path
}
func (g *GIT) UpdateGitAccessUser(user string) {
	g.AccessUser = user
}
func (g *GIT) UpdateGitAccessToken(token string) {
	g.AccessToken = token
}

func cloneRepo(url string, branch string, directory string, name string, token string) error {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	//	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone -b %s --single-branch %s %s", branch, url, directory)

	_, err := git.PlainClone(directory, false, &git.CloneOptions{

		Auth: &http.BasicAuth{
			Username: name, // yes, this can be anything except an empty string
			Password: token,
		},
		ReferenceName: plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
		URL:           url,
		Progress:      os.Stdout,
		SingleBranch:  true,
	})
	CheckIfError(err)
	return err
}

func CreateBranch(url string, newbranch string, directory string) error {
	/*
		CheckArgs("<url>", "<directory>")
		//url, directory := os.Args[1], os.Args[2]
		// Clone the given repository to the given directory
		Info("git clone %s %s", url, directory)
		r, err := git.PlainClone(directory, false, &git.CloneOptions{
			URL: url,
		})
		CheckIfError(err)
	*/
	r, err := git.PlainOpen(directory)
	// Create a new branch to the current HEAD
	Info("git branch my-branch")

	headRef, err := r.Head()
	CheckIfError(err)

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.

	// For tags we should use `refs/tags/%s` instead of `refs/heads/%s` used
	// for branches.
	n := plumbing.ReferenceName("refs/heads/" + newbranch)

	ref := plumbing.NewHashReference(n, headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
	CheckIfError(err)
	return err
}

func PushGit(path string, name string, token string, newbranch string, url string) error {
	CheckArgs("<repository-path>")
	//path := os.Args[1]
	log.Println("-------74------")
	r, err := git.PlainOpen(path)
	//r, err := git.PlainOpen(path)
	CheckIfError(err)

	Info("git push")
	// push using default options
	log.Println("---------88---------")
	err = r.Push(&git.PushOptions{
		Auth: &http.BasicAuth{
			Username: name,
			Password: token,
		},
	})
	CheckIfError(err)
	return err
}

func CheckoutBranch(url string, newbranch string, directory string) {
	//	CheckArgs("<url>", "<directory>", "<commit>")
	//	url, directory, commit := os.Args[1], os.Args[2], os.Args[3]
	/*
		// Clone the given repository to the given directory
		Info("git clone %s %s", url, directory)
		r, err := git.PlainClone(directory, false, &git.CloneOptions{
			URL: url,
		})
	*/
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD
	Info("git show-ref --head HEAD")
	ref, err := r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())

	w, err := r.Worktree()
	CheckIfError(err)

	// ... checking out to commit
	Info("git checkout %s", newbranch)
	err = w.Checkout(&git.CheckoutOptions{
		//	Hash: plumbing.NewHash(newbranch),
		Branch: plumbing.ReferenceName("refs/heads/" + newbranch),
	})
	CheckIfError(err)

	// ... retrieving the commit being pointed by HEAD, it shows that the
	// repository is pointing to the giving commit in detached mode
	Info("git show-ref --head HEAD")
	ref, err = r.Head()
	CheckIfError(err)
	fmt.Println(ref.Hash())
}

func CommitRepo(directory string, filename string) {
	CheckArgs("<directory>")
	//directory := os.Args[1]

	// Opens an already existing repository.
	r, err := git.PlainOpen(directory)
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.

	/*
		Info("echo \"hello world!\" > example-git-file")
		filename := filepath.Join(directory, "example-git-file")
		err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
		CheckIfError(err)
	*/
	// Adds the new file to the staging area.
	fmt.Printf("git add %s\n", filename)
	_, err = w.Add(filename)
	CheckIfError(err)

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	Info("git commit -m \"misc: gdeyamlOperator auto commit to update base.branch\"")
	commit, err := w.Commit("misc: gdeyamlOperator auto commit to update base.branch", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "siangyeh8818",
			Email: "siangyeh8818@gmail.com",
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)
}

func ClonePushNewBranch(git_url string, old_branch string, new_branch string, git_repo_path string, git_user string, git_token string) {
	log.Println("-----action >> cloneRepo----")
	cloneRepo(git_url, old_branch, git_repo_path, git_user, git_token)
	log.Println("-----action >> CreateBranch----")
	CreateBranch(git_url, new_branch, git_repo_path)
	log.Println("-----action >> CheckoutBranch----")
	CheckoutBranch(git_url, new_branch, git_repo_path)
	log.Println("-----action >> PushGit----")
	PushGit(git_repo_path, git_user, git_token, new_branch, git_url)
	log.Println("-----action finishing----")
}
