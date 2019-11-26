package main

import (
	"fmt"

	valid "github.com/siangyeh8818/gdeyamlOperator/internal/validation"
)

func main() {
	// branch := "0.145"

	// branch := "0.145.22"

	// branch := "p/123/ricks_sta-ss"

	branch := "wayn002"
	pattern := valid.Validate(branch)

	fmt.Printf("pattern: %v\n", pattern)
}
