package main

import (
	"regexp"
)

func IdentifyCrongob(test_string string) bool {
	var idf_token bool
	idf_token = true
	if ok, _ := regexp.Match(`wf-\S{9}-\d{10}`, []byte(test_string)); ok {
		idf_token = false
	}
	return idf_token
}
