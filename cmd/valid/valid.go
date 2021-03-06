package main

import (
	"fmt"

	valid "github.com/siangyeh8818/gdeyamlOperator/internal/validation"
)

func main() {
	// branch := "0.145"

	// branch := "0.145.22"

	// branch := "p/123/ricks_sta-ss"

	// branch := "rick-0.167.2-releasev2"
	branch := "wayne002"

	// branch := "fasdfjl%&3"
	pattern, err := valid.Validate(branch)

	if err != nil {
		fmt.Printf("err: %v", err)
	}
	fmt.Printf("Convention: %s\n", pattern)
}
