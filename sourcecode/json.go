package main

import (
	"log"

	"github.com/buger/jsonparser"
)

type JsonType struct {
	Array []string
}

func JsonParse(jsondata string) {
	var data []byte = []byte(jsondata)
	/*items, err := jsonparser.GetString(data, "items")
	log.Println(items)
	items_arr := JsonType{}
	_ = json.Unmarshal([]byte(items), &items_arr)
	log.Printf("Unmarshaled: %v", items_arr)
	*/
	scripts_name, err := jsonparser.GetString(data, "items", "[0]", "assets", "[0]", "path")
	if err != nil {
		log.Println(err)
	}
	log.Println(scripts_name)
	download_url, err := jsonparser.GetString(data, "items", "[0]", "assets", "[0]", "downloadUrl")
	if err != nil {
		log.Println(err)
	}
	log.Println(download_url)

	continue_token, err := jsonparser.GetString(data, "continuationToken")
	if err != nil {
		log.Println(err)
	}
	log.Println(continue_token)
	//if continue_token != "" {
	//	func JsonParse()
	//}
}

/*
type NexusResponse struct {
	Items []Items
	continuationToken string
}

func JsonParse(jsondata string) {

}
*/
