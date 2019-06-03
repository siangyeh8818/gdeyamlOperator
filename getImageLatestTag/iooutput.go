package main

import (
  "fmt"
    "io/ioutil"
)

func WriteWithIoutil(name, content string) {
  data := []byte(content)
    if ioutil.WriteFile(name, data, 0644) == nil {
	    fmt.Println("Success to export to file", content)
	  }
  
  }
