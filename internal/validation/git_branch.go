package validation

import (
	"fmt"
	"regexp"
)

// BranchConvention An new int type
type BranchConvention int

const (
	// Invalid not able to recognized
	Invalid BranchConvention = iota - 1
	// Release representing a release branch
	Release
	// Patch a pattern represents a patch
	Patch
	// Feature representing a feature branch
	Feature
	// Custom representing a custom named branch
	Custom
)

// 0.155: continue
const release = `^v?([0-9]+)(\.[0-9]+)$`

// 0.155.2: need to remove last '.2' and use '0.155' to continue
const patch = `^v?(\d+)(\.(\d+))(\.(\d+))$`

// p/rick/add_release-version: continue
const feature = `^[a-zA-Z][/_][0-9]+[/_]{1,1}[a-zA-Z0-9_-]+$`

// wayne002: continue
const custom = `^[a-z0-9._-]+$`

func (bc BranchConvention) String() string {
	switch bc {
	case Release:
		return "release"
	case Patch:
		return "patch"
	case Feature:
		return "feature"
	case Custom:
		return "custom"
	case Invalid:
		return "invalid"
	}
	return "unknown"
}

// Validate check the input git branch is matched a predefined pattern
func Validate(branch string) (BranchConvention, error) {

	match, err := regexp.MatchString(release, branch)
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

	match, err = regexp.MatchString(feature, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Feature, Feature)
		return Feature, err
	}

	match, err = regexp.MatchString(custom, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}

	if match {
		fmt.Printf("%s %d\n", Custom, Custom)
		return Custom, err
	}

	fmt.Printf("%s %d\n", Invalid, Invalid)
	return Invalid, err
}
