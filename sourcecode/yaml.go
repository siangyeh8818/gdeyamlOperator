package main

type Deployment struct {
	K8S      []K8S      `yaml:"k8s"`
	Openfaas []Openfaas `yaml:"openfaas"`
}
type K8S struct {
	Module string `yaml:"module"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	Stage  string `yaml:"stage"`
}

type Openfaas struct {
	Module string `yaml:"module"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	Stage  string `yaml:"stage"`
}

type K8sYaml struct {
	Deployment Deployment `yaml:"deployment"`
}

func (s *Deployment) AddK8sStruct(module string, image string, tag string, stage string) {
	var a K8S = K8S{
		Module: module,
		Image:  image,
		Tag:    tag,
		Stage:  stage,
	}
	s.K8S = append(s.K8S, a)
	/*
		append(s.K8S, K8S{})
		length := len(s.K8S)
		s.K8S[length-1].Module = module
		s.K8S[length-1].Image = image
		s.K8S[length-1].Tag = tag
		s.K8S[length-1].Stage = stage
	*/
}

func (s *Deployment) AddOpenfaasStruct(module string, image string, tag string, stage string) {
	var a Openfaas = Openfaas{
		Module: module,
		Image:  image,
		Tag:    tag,
		Stage:  stage,
	}
	s.Openfaas = append(s.Openfaas, a)

}
func (s *K8S) UpdateK8sTag(newtag string) {

	s.Tag = newtag
}

func (s *Openfaas) UpdateOpenfaasTag(newtag string) {

	s.Tag = newtag
}
