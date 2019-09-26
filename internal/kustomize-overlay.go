package gdeyamloperator

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

type KustomizeArgument struct {
	Outputdir   string
	Comparedata string
	Namespace   string
	RelPath     string
	K8sBaseloc  string
	OfBaseloc   string
	Kmodules    string
	UrlPattern  string
	EnvFile     string
}

type KuzImageConfig struct {
	Name    string `yaml:"name,omitempty"`
	NewName string `yaml:"newName,omitempty"`
	NewTag  string `yaml:"newTag,omitempty"`
	//tt KuzNameTag
}
type KuzConfigs struct {
	Bases     []string `yaml:"bases,omitempty"`
	Namespace string   `yaml:"namespace,omitempty"`
	//Images    []KuzImageConfig `yaml:"images,omitempty"`
	Images []KuzImageConfig `yaml:"images,omitempty"`
}

func (k *KustomizeArgument) UpdateKustomizeArgument(outputdir string, comparedata string, namespace string, relpath string, k8sbase string, ofbaseloc string, kmodule string, urlpattern string, envfilename string) {
	k.Outputdir = outputdir
	k.Comparedata = comparedata
	k.Namespace = namespace
	k.RelPath = relpath
	k.K8sBaseloc = k8sbase
	k.OfBaseloc = ofbaseloc
	k.Kmodules = kmodule
	k.UrlPattern = urlpattern
	k.EnvFile = envfilename
}

//func OutputOverlays(envfile string, deployfile string, namespace string, kmodules string, relPath string, k8sBaseloc string) {
func OutputOverlays(kus_argument *KustomizeArgument, deployfile string) {
	//if deployfile == "" {
	//	deployfile = "deploy.yml"
	//}
	//attrsns := GetNameSpaceMapping(envfile, namespace, kmodules)
	attrsns := GetNameSpaceMapping(kus_argument)
	//loop key
	for attr := range attrsns {
		fmt.Println("------attr (for)-----------")
		fmt.Println(attr)
		ImageList := GetYamlAttribute(deployfile, attr, kus_argument)
		fmt.Println("------ImageList-----------")
		fmt.Println(ImageList)
		OutYaml(ImageList, attrsns[attr], attr, kus_argument)

	}

}

//func GetNameSpaceMapping(envfile string, namespace string, kmodules string) map[string]string {
func GetNameSpaceMapping(kus_argument *KustomizeArgument) map[string]string {
	ret := make(map[string]string)
	if kus_argument.Kmodules != "" {
		ret["k8s"] = kus_argument.Namespace
		return ret
	}
	data, err := ioutil.ReadFile(kus_argument.EnvFile)
	if err != nil {
		if kus_argument.Namespace == "" {
			log.Panic("no env file, you need to enter namespace for deploy in GDE")
		}
		ret["k8s"] = kus_argument.Namespace
		nsopenfaas := fmt.Sprintf("%s-openfaas-fn", kus_argument.Namespace)
		ret["openfaas"] = nsopenfaas
		ret["monitor"] = kus_argument.Namespace
		ret["redis"] = kus_argument.Namespace
		return ret
	}
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	d := m["namespaces"].([]interface{})[0]
	fmt.Println(d)
	ll, _ := GetKeys(d)
	for _, l := range ll {
		key := l
		val := d.(map[string]interface{})[l].(string)
		fmt.Println("key:", l)
		fmt.Println("val:", d.(map[string]interface{})[l])
		ret[key] = val
	}

	//ret["openfaas"] = "openfaas-fn"
	fmt.Println(ret)
	//log.Panic("stop")
	return ret
}
func GetKeys(v interface{}) ([]string, error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Map {
		return nil, errors.New("not a map")
	}
	t := rv.Type()
	if t.Key().Kind() != reflect.String {
		return nil, errors.New("not string key")
	}
	var result []string
	for _, kv := range rv.MapKeys() {
		result = append(result, kv.String())
	}
	return result, nil
}

func GetYamlAttribute(filename string, key string, kus_argument *KustomizeArgument) []ImgInfo {
	//var ret []string
	var ret []ImgInfo
	if kus_argument.Kmodules != "" {
		im := GetAttrByCommand(kus_argument.Kmodules)
		ret = append(ret, im)
		return ret

	}
	fmt.Println("------------")
	//filename := "test.yaml"
	fmt.Println(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	d := m["deployment"].(map[interface{}]interface{})[key]
	fmt.Println(d)
	for _, dd := range d.([]interface{}) {
		var im ImgInfo
		if dd.(map[string]interface{})["module"] != nil {
			im.ModName = dd.(map[string]interface{})["module"].(string)
		} else {
			im.ModName = ""
		}

		//im.Module = dd.(map[string]interface{})["module"].(string)
		if dd.(map[string]interface{})["image"] != nil {
			im.ImgName = dd.(map[string]interface{})["image"].(string)
		} else {
			im.ImgName = ""
		}
		//im.Image = dd.(map[string]interface{})["image"].(string)
		if dd.(map[string]interface{})["tag"] != nil {
			im.Tag = dd.(map[string]interface{})["tag"].(string)
		} else {
			im.Tag = ""
		}
		//im.Tag = dd.(map[string]interface{})["tag"].(string)
		if dd.(map[string]interface{})["stage"] != nil {
			im.Stage = dd.(map[string]interface{})["stage"].(string)
		} else {
			im.Stage = ""
		}
		ret = append(ret, im)
		//keys := GetAllKeys(dd)
		//fmt.Println(keys)
	}
	//fmt.Println(ret)
	return ret
}

func GetAttrByCommand(kmodules string) ImgInfo {
	var nc ImgInfo
	kPods := strings.Split(kmodules, ",")
	if kmodules != "" {
		fmt.Println("command from k8s")
		for n := 0; n < len(kPods); n++ {
			fmt.Println(n)
			mm := strings.Split(kPods[n], ":")
			if len(mm) == 1 {
				//mc.Module = mm[0]
				nc.ModName = mm[0]
			} else {
				//	mc.Module = mm[0]
				//	mc.Image = mm[1]
				//	mc.Stage = mm[2]
				//	mc.Tag = mm[3]
				nc.ModName = mm[0]
				nc.ImgName = mm[1]
				nc.Stage = mm[2]
				nc.Tag = mm[3]
			}
		}
	}
	return nc
	//panic("stop")
}

func OutYaml(Mods []ImgInfo, namespace string, typeapp string, kus_argument *KustomizeArgument) {
	// create path
	var basepath []string
	for i := 0; i < len(Mods); i++ {
		// relative path with /
		realPath := fmt.Sprintf("%s%s/%s", kus_argument.RelPath, kus_argument.K8sBaseloc, Mods[i].ModName)
		basepath = append(basepath, realPath)
	}
	//Path := "./../base"
	//basepath = append(basepath, "111")

	// create images
	var Imgs []KuzImageConfig
	for i := 0; i < len(basepath); i++ {
		//ignore no tag and image, and deploy it directly
		if Mods[i].ImgName != "" {
			var Img KuzImageConfig
			sk8sBasesLoc := fmt.Sprintf("%s/%s", kus_argument.K8sBaseloc, Mods[i].ModName)
			Img.Name, _ = TryMatchingImage(sk8sBasesLoc, Mods[i].ImgName, "yml&yaml")
			//Img.Name = K8SMods[i].ImgName
			Img.NewTag = Mods[i].Tag
			//Img.NewName = changeStage fmt.Sprintf("%s/%s", K8SMods[i].Stage, K8SMods[i].ImgName)
			Img.NewName = ChangeStage(Mods[i], kus_argument)
			//fmt.Println("search:", sa)
			Imgs = append(Imgs, Img)
		}
	}
	fmt.Println("show list:", Imgs)
	data := KuzConfigs{Bases: basepath, Images: Imgs, Namespace: namespace}
	//data := HybridInfoData{OpenFaasModules: OpenFaasMods, K8SModules: K8SMods, Namespace: namespace, Path: Path}
	d, _ := yaml.Marshal(&data)
	fmt.Println(string(d))
	var output string
	output = fmt.Sprintf("%s/%s", kus_argument.Outputdir, typeapp)
	os.MkdirAll(output, os.ModePerm)
	output = fmt.Sprintf("%s/%s/%s", kus_argument.Outputdir, typeapp, "kustomization.yaml")

	//outputtest := "./test.txt"
	outfile, _ := os.Create(output)
	outfile.Write([]byte(d))

}

func ChangeStage(imginfo ImgInfo, kus_argument *KustomizeArgument) string {
	ret := strings.Replace(kus_argument.UrlPattern, "{{stage}}", imginfo.Stage, -1)
	retImg := strings.Replace(ret, "{{image}}", imginfo.ImgName, -1)
	retAll := strings.Replace(retImg, "{{tag}}", imginfo.Tag, -1)
	retCut := strings.Split(retAll, ":")[0]
	return retCut
}
