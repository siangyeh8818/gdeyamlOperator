package gdeyamloperator

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
	var insert_content string

	if len(nexusresponse.Items) > 0 {
		for i := 0; i < len(nexusresponse.Items); i++ {
			temp_insert_content := ""

			if strings.Contains(out_pattern, "id") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Id)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Id)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Id) + ","
			}
			if strings.Contains(out_pattern, "repository") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Repository)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Repository)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Repository) + ","
			}
			if strings.Contains(out_pattern, "format") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Format)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Format)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Format) + ","
			}
			if strings.Contains(out_pattern, "group") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Group)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Group)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Group) + ","
			}
			if strings.Contains(out_pattern, "name") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Name)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Name)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Name) + ","
			}
			if strings.Contains(out_pattern, "version") {
				if len(temp_insert_content) > 0 {
					temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Version)
				} else {
					temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Version)
				}
				//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Version) + ","
			}
			for k := 0; k < len(nexusresponse.Items[i].Assets); k++ {
				if strings.Contains(out_pattern, "assets.downloadUrl") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Assets[k].DownloadUrl)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].DownloadUrl)
					}
					//log.Println(string(nexusresponse.Items[i].Assets[k].DownloadUrl))
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].DownloadUrl) + ","
				}
				if strings.Contains(out_pattern, "assets.path") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Assets[k].Path)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Path)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Path) + ","
				}
				if strings.Contains(out_pattern, "assets.id") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Assets[k].Id)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Id)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Id) + ","
				}
				if strings.Contains(out_pattern, "assets.repository") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Assets[k].Repository)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Repository)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Repository) + ","
				}
				if strings.Contains(out_pattern, "assets.format") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Assets[k].Format)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Format)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Assets[k].Format) + ","
				}
			}

			for z := 0; z < len(nexusresponse.Items[i].Tags); z++ {
				if strings.Contains(out_pattern, "tags") {
					if len(temp_insert_content) > 0 {
						temp_insert_content = temp_insert_content + "," + string(nexusresponse.Items[i].Tags[z].Tagname)
					} else {
						temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Tags[z].Tagname)
					}
					//temp_insert_content = temp_insert_content + string(nexusresponse.Items[i].Tags[z].Tagname) + ","
				}
			}
			if len(nexusresponse.Items)-1 == i {
				temp_insert_content = temp_insert_content
			}else {
				temp_insert_content = temp_insert_content + "\n"
			}
			
			insert_content = insert_content + temp_insert_content
		}
	}

	log.Println(insert_content)
	output.Addcontent(insert_content)

}

/*
func toStringArray(data []byte) (result []string) {
	jsonparser.ArrayEach(data, func(value []byte, dataType ValueType, offset int, err error) {

	}, "assets")

	return
}
*/
