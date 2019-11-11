package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/buger/jsonparser"
)

// JSONType a structure contains Items and a continuation token
type JSONType struct {
	Items             []Item `json:"items"`
	ContinuationToken string `json:"continuationToken"`
}

// Item a structure represent an item
type Item struct {
	ID         string  `json:"id"`
	Repository string  `json:"repository"`
	Format     string  `json:"format"`
	Group      string  `json:"group"`
	Name       string  `json:"name"`
	Version    string  `json:"version"`
	Assets     []Asset `json:"assets"`
	Tags       []Tag   `json:"tags"`
}

// Tag structure represents a tag
type Tag struct {
	TagName string
}

// Asset structure represents an asset
type Asset struct {
	DownloadURL string `json:"downloadUrl"`
	Path        string `json:"path"`
	ID          string `json:"id"`
	Repository  string `json:"repository"`
	Format      string `json:"format"`
}

// OutputContent collect the content strings
type OutputContent struct {
	Content []string
}

func (s *OutputContent) Addcontent(tempcontent string) {

	s.Content = append(s.Content, tempcontent)
}

func continueTokenParse(jsondata string) string {
	var data = []byte(jsondata)
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
	continueToken, err := jsonparser.GetString(data, "continuationToken")
	if err != nil {
		log.Println(err)
	}
	log.Println(continueToken)
	return continueToken
}

func JsonParse2(jsondata string, out_pattern string, output *OutputContent) {
	var nexusresponse JSONType

	json.Unmarshal([]byte(jsondata), &nexusresponse)

	log.Println("---------JsonParse2--------------")
	var insertContent string

	if len(nexusresponse.Items) > 0 {
		for i := 0; i < len(nexusresponse.Items); i++ {
			tempInsertContent := ""

			if strings.Contains(out_pattern, "id") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].ID)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].ID)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Id) + ","
			}
			if strings.Contains(out_pattern, "repository") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Repository)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Repository)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Repository) + ","
			}
			if strings.Contains(out_pattern, "format") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Format)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Format)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Format) + ","
			}
			if strings.Contains(out_pattern, "group") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Group)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Group)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Group) + ","
			}
			if strings.Contains(out_pattern, "name") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Name)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Name)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Name) + ","
			}
			if strings.Contains(out_pattern, "version") {
				if len(tempInsertContent) > 0 {
					tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Version)
				} else {
					tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Version)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Version) + ","
			}
			for k := 0; k < len(nexusresponse.Items[i].Assets); k++ {
				if strings.Contains(out_pattern, "assets.downloadUrl") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Assets[k].DownloadURL)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Assets[k].DownloadURL)
					}
					//log.Println(string(nexusresponse.Items[i].Assets[k].DownloadUrl))
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].DownloadUrl) + ","
				}
				if strings.Contains(out_pattern, "assets.path") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Assets[k].Path)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Assets[k].Path)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Path) + ","
				}
				if strings.Contains(out_pattern, "assets.id") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Assets[k].ID)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Assets[k].ID)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Id) + ","
				}
				if strings.Contains(out_pattern, "assets.repository") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Assets[k].Repository)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Assets[k].Repository)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Repository) + ","
				}
				if strings.Contains(out_pattern, "assets.format") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Assets[k].Format)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Assets[k].Format)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Format) + ","
				}
			}

			for z := 0; z < len(nexusresponse.Items[i].Tags); z++ {
				if strings.Contains(out_pattern, "tags") {
					if len(tempInsertContent) > 0 {
						tempInsertContent = tempInsertContent + "," + string(nexusresponse.Items[i].Tags[z].TagName)
					} else {
						tempInsertContent = tempInsertContent + string(nexusresponse.Items[i].Tags[z].TagName)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Tags[z].Tagname) + ","
				}
			}
			tempInsertContent = tempInsertContent + "\n"
			insertContent = insertContent + insertContent
		}
	}

	log.Println(insertContent)
	output.Addcontent(insertContent)

}

/*
func toStringArray(data []byte) (result []string) {
	jsonparser.ArrayEach(data, func(value []byte, dataType ValueType, offset int, err error) {

	}, "assets")

	return
}
*/
