package gdeyamloperator

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func ZipHpptDownload(url string) {
	/****Http下载方式*****/

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	log.Println(resp.Header.Get("content-type"))
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	CopyFile(b, UrlParserZipname(url))
}
func CopyFile(byte []byte, dst string) (w int64, err error) {
	dstFile, err := os.Create(dst)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer dstFile.Close()
	return io.Copy(dstFile, bytes.NewReader(byte))
}
func UrlParserZipname(url string) string {
	var zip_name string
	url_componet_name := strings.Split(url, "/")
	log.Printf("output zip name should be &s", url_componet_name[len(url_componet_name)-1])
	zip_name = url_componet_name[len(url_componet_name)-1]
	return zip_name

}
