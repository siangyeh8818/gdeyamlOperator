package gdeyamloperator

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func WriteWithIoutil(name, content string) {
	data := []byte(content)
	if ioutil.WriteFile(name, data, 0644) == nil {
		fmt.Println("Success to export to file\n", content)
	}

}

//getConf
func (c *K8sYaml) GetConf(f string) *K8sYaml {
	//应该是 绝对地址
	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return c
}

func (c *Environmentyaml) GetConf(f string) *Environmentyaml {

	yamlFile, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	return c
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
