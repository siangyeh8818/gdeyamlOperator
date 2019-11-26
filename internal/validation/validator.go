package validation

import (
	"fmt"
	"regexp"
)

// 0.155: continue
const release = `^v?([0-9]+)(\.[0-9]+)$`

// 0.155.2: need to remove last '.2' and use '0.155' to continue
const patch = `^v?(\d+)(\.(\d+))(\.(\d+))$`

// p/rick/add_release-version: continue
const feature = `^[a-zA-Z]\/[0-9]+\/{1,1}[a-zA-Z0-9_-]+$`

// wayne002: continue
const misc = `^[a-z0-9_-]+$`

// Check fff
func Check() {
	var branch = "wayne0_-02"
	fmt.Println("start to validate")
	matched := validateBranch(branch)
	fmt.Printf("matched: %v\n", matched)
}

func validateBranch(branch string) bool {
	match, err := regexp.MatchString(misc, branch)
	if err != nil {
		fmt.Printf("MatchString err: %v", err)
	}
	fmt.Printf("match: %v\n", match)
	return match
}
