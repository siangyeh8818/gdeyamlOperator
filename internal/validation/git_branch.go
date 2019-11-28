package validation

import (
	"fmt"
	"regexp"
)

// BranchConvention An new int type
type BranchConvention int

const (
	// Release representing a release branch
	Release BranchConvention = iota
	// Patch a pattern represents a patch
	Patch
	// Feature representing a feature branch
	Feature
	// Misc representing a custom named branch
	Misc
	// Invalid not able to recognized
	Invalid
)

// 0.155: continue
const release = `^v?([0-9]+)(\.[0-9]+)$`

// 0.155.2: need to remove last '.2' and use '0.155' to continue
const patch = `^v?(\d+)(\.(\d+))(\.(\d+))$`

// p/rick/add_release-version: continue
const feature = `^[a-zA-Z][/_][0-9]+[/_]{1,1}[a-zA-Z0-9_-]+$`

// wayne002: continue
const misc = `^[a-z0-9_-]+$`

func (bc BranchConvention) String() string {
	switch bc {
	case Release:
		return "release"
	case Patch:
		return "patch"
	case Feature:
		return "feature"
	case Misc:
		return "misc"
	case Invalid:
		return "invalid"
	}
	return "unknown"
}

// Validate check the input git branch is matched a predefined pattern
func Validate(branch string) (BranchConvention, error) {

	match, err := regexp.MatchString(feature, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Feature, Feature)
		return Feature, err
	}

	match, err = regexp.MatchString(misc, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Misc, Misc)
		return Misc, err
	}

	match, err = regexp.MatchString(release, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Release, Release)
		return Release, err
	}

	match, err = regexp.MatchString(patch, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Patch, Patch)
		return Patch, err
	}

	fmt.Printf("%s %d\n", Invalid, Invalid)
	return Invalid, err
}
