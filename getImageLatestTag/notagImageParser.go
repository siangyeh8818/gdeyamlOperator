package main

import (
	"strings"
)

func ImagenameSplit(rawimage string) (string, string) {
	var parsername []string
	var tag []string

	var return_image string

	parsername = strings.Split(rawimage, "/")
	if len(parsername) == 2 {
		for i := range parsername {
			//	fmt.Println(parsername[i])
			if i == 1 {
				tag = strings.Split(parsername[i], ":")
				/*      for j := range tag {
				                fmt.Println(tag[j])
							              }
				*/
			}
		}
		return_image = tag[0]
	} else if len(parsername) == 3 {
		for i := range parsername {
			//	fmt.Println(parsername[i])
			if i == 2 {
				tag = strings.Split(parsername[i], ":")
			}
		}
		return_image = parsername[1] + "/" + tag[0]

	}
	return parsername[0], return_image
}
