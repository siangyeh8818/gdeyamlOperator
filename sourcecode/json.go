package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/buger/jsonparser"
)

type JsonType struct {
	Items             []Items `json:"items"`
	ContinuationToken string  `json:"continuationToken"`
}
type Items struct {
	Id         string   `json:"id"`
	Repository string   `json:"repository"`
	Format     string   `json:"format"`
	Group      string   `json:"group"`
	Name       string   `json:"name"`
	Version    string   `json:"version"`
	Assets     []Assets `json:"assets"`
	Tags       []Tags   `json:"tags"`
}
type Tags struct {
	Tagname string
}
type Assets struct {
	DownloadUrl string `json:"downloadUrl"`
	Path        string `json:"path"`
	Id          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
}
type OutputContent struct {
	Content []string
}

func (s *OutputContent) Addcontent(tempcontent string) {

	s.Content = append(s.Content, tempcontent)
}

func continueTokenParse(jsondata string) string {
	var data []byte = []byte(jsondata)
	/*
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
	*/
	continue_token, err := jsonparser.GetString(data, "continuationToken")
	if err != nil {
		log.Println(err)
	}
	log.Println(continue_token)
	return continue_token
}

func JsonParse2(jsondata string, out_pattern string, output *OutputContent) {
	var nexusresponse JsonType

	json.Unmarshal([]byte(jsondata), &nexusresponse)

	log.Println("---------JsonParse2--------------")
	var temp_insert_content string

	if len(nexusresponse.Items) > 0 {
		for i := 0; i < len(nexusresponse.Items); i++ {
			temp_insert_content = ""
			if strings.Contains(out_pattern, "id") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Id) + ","
			}
			if strings.Contains(out_pattern, "repository") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Repository) + ","
			}
			if strings.Contains(out_pattern, "format") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Format) + ","
			}
			if strings.Contains(out_pattern, "group") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Group) + ","
			}
			if strings.Contains(out_pattern, "name") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Name) + ","
			}
			if strings.Contains(out_pattern, "version") {
				temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Version) + ","
			}
			for k := 0; k < len(nexusresponse.Items[i].Assets); k++ {
				if strings.Contains(out_pattern, "assets.downloadUrl") {
					log.Println(string(nexusresponse.Items[i].Assets[k].DownloadUrl))
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].DownloadUrl) + ","
				}
				if strings.Contains(out_pattern, "assets.path") {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Path) + ","
				}
				if strings.Contains(out_pattern, "assets.id") {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Id) + ","
				}
				if strings.Contains(out_pattern, "assets.repository") {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Repository) + ","
				}
				if strings.Contains(out_pattern, "assets.format") {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Format) + ","
				}
			}

			for z := 0; z < len(nexusresponse.Items[i].Tags); z++ {
				if strings.Contains(out_pattern, "tags") {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Tags[z].Tagname) + ","
				}
			}

		}
	}

	log.Println(temp_insert_content)
	output.Addcontent(temp_insert_content)

}

/*
func toStringArray(data []byte) (result []string) {
	jsonparser.ArrayEach(data, func(value []byte, dataType ValueType, offset int, err error) {

	}, "assets")

	return
}
*/
