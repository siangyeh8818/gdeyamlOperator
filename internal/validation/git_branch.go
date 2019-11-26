package validation

import (
	"fmt"
	"regexp"
)

type BranchPattern struct {
}

// 0.155: continue
const release = `^v?([0-9]+)(\.[0-9]+)$`

// 0.155.2: need to remove last '.2' and use '0.155' to continue
const patch = `^v?(\d+)(\.(\d+))(\.(\d+))$`

// p/rick/add_release-version: continue
const feature = `^[a-zA-Z][/_][0-9]+[/_]{1,1}[a-zA-Z0-9_-]+$`

// wayne002: continue
const misc = `^[a-z0-9_-]+$`

// Validate check the input git branch is matched a predefined pattern
func Validate(branch string) string {

	match, err := regexp.MatchString(feature, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}
	if match {
		return "feature"
	}

	match, err = regexp.MatchString(misc, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}
	if match {
		return "misc"
	}

	match, err = regexp.MatchString(release, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}
	if match {
		return "release"
	}

	match, err = regexp.MatchString(patch, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}
	if match {
		return "patch"
	}

	return "N/A"
}
