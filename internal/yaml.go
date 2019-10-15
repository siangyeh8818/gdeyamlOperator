package gdeyamloperator

type Deployment struct {
	BASE      []BASE     `yaml:"base"`
	BLCKS     BLCKS      `yaml:"blcks"`
	PLAYBOOKS PLAYBOOKS  `yaml:"playbooks"`
	K8S       []K8S      `yaml:"k8s"`
	Openfaas  []Openfaas `yaml:"openfaas"`
	Monitor   []Monitor  `yaml:"monitor" `
	Redis     []Redis    `yaml:"redis"`
}
type BASE struct {
	Git    string `yaml:"git"`
	Branch string `yaml:"branch"`
}

type TOOL struct {
	Module string `yaml:"module"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
}

type BLCKS struct {
	Git     string `yaml:"git"`
	Branch  string `yaml:"branch"`
	Version string `yaml:"version"`
	TOOL    TOOL   `yaml:"tool"`
}

type PLAYBOOKS struct {
	Git     string `yaml:"git"`
	Branch  string `yaml:"branch"`
	Version string `yaml:"version"`
	TOOL    TOOL   `yaml:"tool"`
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

type Environmentyaml struct {
	Namespaces []struct {
		K8S      string `yaml:"k8s"`
		Openfaas string `yaml:"openfaas"`
		Monitor  string `yaml:"monitor"`
		Redis    string `yaml:"redis"`
	} `yaml:"namespaces"`
	Configuration []struct {
		Git    string `yaml:"git"`
		Branch string `yaml:"branch"`
	} `yaml:"configuration"`
	Deploymentfile []struct {
		Git     string `yaml:"git"`
		Branch  string `yaml:"branch"`
		Replace struct {
			K8S      []K8S      `yaml:"k8s"`
			Openfaas []Openfaas `yaml:"openfaas"`
			Monitor  []Monitor  `yaml:"monitor"`
			Redis    []Redis    `yaml:"redis"`
		} `yaml:"replace"`
		Ignore struct {
			K8S      []K8S      `yaml:"k8s"`
			Openfaas []Openfaas `yaml:"openfaas"`
			Monitor  []Monitor  `yaml:"monitor"`
			Redis    []Redis    `yaml:"redis"`
		} `yaml:"ignore"`
	} `yaml:"deploymentfile"`
}

type Monitor struct {
	Module string `yaml:"module"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	Stage  string `yaml:"stage"`
}
type Redis struct {
	Module string `yaml:"module"`
	Image  string `yaml:"image"`
	Tag    string `yaml:"tag"`
	Stage  string `yaml:"stage"`
}

func (s *Deployment) AddBaseStruct(git string, gitbranch string) {
	var a BASE = BASE{
		Git:    git,
		Branch: gitbranch,
	}
	s.BASE = append(s.BASE, a)
}
func (s *Deployment) UpdateBaseStructBranch(git string, gitbranch string) {
	s.BASE[0].Git = git
	s.BASE[0].Branch = gitbranch
}
func (s *Deployment) UpdateBlcksStructBranch(git string, gitbranch string, ver string) {
	s.BLCKS.Git = git
	s.BLCKS.Branch = gitbranch
	s.BLCKS.Version = ver
}

func (s *Deployment) UpdatePLAYBOOKStructBranch(git string, gitbranch string, ver string) {
	s.PLAYBOOKS.Git = git
	s.PLAYBOOKS.Branch = gitbranch
	s.PLAYBOOKS.Version = ver
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

func (s *Deployment) AddMonitorStruct(module string, image string, tag string, stage string) {
	var a Monitor = Monitor{
		Module: module,
		Image:  image,
		Tag:    tag,
		Stage:  stage,
	}
	s.Monitor = append(s.Monitor, a)
}

func (s *Deployment) AddRedisStruct(module string, image string, tag string, stage string) {
	var a Redis = Redis{
		Module: module,
		Image:  image,
		Tag:    tag,
		Stage:  stage,
	}
	s.Redis = append(s.Redis, a)
}
func (s *Deployment) RemoveK8sStruct(index int) {
	s.K8S = append(s.K8S[:index], s.K8S[index+1:]...)
}
func (s *Deployment) RemoveOpenfaasStruct(index int) {
	s.Openfaas = append(s.Openfaas[:index], s.Openfaas[index+1:]...)
}
func (s *Deployment) RemoveMonitorStruct(index int) {

	s.Monitor = append(s.Monitor[:index], s.Monitor[index+1:]...)
}
func (s *Deployment) RemoveRedisStruct(index int) {

	s.Redis = append(s.Redis[:index], s.Redis[index+1:]...)
}

func (s *K8S) UpdateK8sTag(newtag string) {

	s.Tag = newtag
}

func (s *K8S) UpdateK8sModule(newmodule string) {

	s.Module = newmodule
}

func (s *K8S) UpdateK8sImage(newimage string) {

	s.Image = newimage
}

func (s *K8S) UpdateK8sStage(newstage string) {

	s.Stage = newstage
}

func (s *Openfaas) UpdateOpenfaasModule(newmodule string) {

	s.Module = newmodule
}

func (s *Openfaas) UpdateOpenfaasImage(newimage string) {

	s.Image = newimage
}

func (s *Openfaas) UpdateOpenfaasTag(newtag string) {

	s.Tag = newtag
}

func (s *Openfaas) UpdateOpenfaasStage(newstage string) {

	s.Stage = newstage
}

func (s *Monitor) UpdateMonitorModule(newmodule string) {

	s.Module = newmodule
}

func (s *Monitor) UpdateMonitorImage(newimage string) {

	s.Image = newimage
}

func (s *Monitor) UpdateMonitorTag(newtag string) {

	s.Tag = newtag
}

func (s *Monitor) UpdateMonitorStage(newstage string) {

	s.Stage = newstage
}

func (s *Redis) UpdateRedisModule(newmodule string) {

	s.Module = newmodule
}

func (s *Redis) UpdateRedisImage(newimage string) {

	s.Image = newimage
}

func (s *Redis) UpdateRedisTag(newtag string) {

	s.Tag = newtag
}

func (s *Redis) UpdateRedisStage(newstage string) {

	s.Stage = newstage
}

func (s *TOOL) UpdateToolTag(newtag string) {

	s.Tag = newtag
}
