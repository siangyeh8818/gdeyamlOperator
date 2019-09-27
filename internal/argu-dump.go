package gdeyamloperator

import (
	"log"
)

func DumpArguments(inputfile string, environment_file string, ouputfile string) {

	var write_content string
	if inputfile != "" && Exists(inputfile) {
		deploy_yaml := K8sYaml{}
		deploy_yaml.GetConf(inputfile)
		log.Println(deploy_yaml.Deployment)
		blcks_Branch := "BLCKS_BRANCH=" + deploy_yaml.Deployment.BLCKS.Branch + "\n"
		write_content = write_content + blcks_Branch
		blcks_version := "BLCKS_VERSION=" + deploy_yaml.Deployment.BLCKS.Version + "\n"
		write_content = write_content + blcks_version
		playbook_Branch := "PLAYBOOK_BRANCH=" + deploy_yaml.Deployment.PLAYBOOKS.Branch + "\n"
		write_content = write_content + playbook_Branch
		playbook_version := "PLAYBOOK_VERSION=" + deploy_yaml.Deployment.PLAYBOOKS.Version + "\n"
		write_content = write_content + playbook_version
		kubestomize_base_Branch := "BASE_BRANCH=" + deploy_yaml.Deployment.BASE[0].Branch + "\n"
		write_content = write_content + kubestomize_base_Branch
	}
	if environment_file != "" && Exists(inputfile) {
		envir_yaml := Environmentyaml{}
		envir_yaml.GetConf(environment_file)
		deploy_Branch := "DEPLOY_BRANCH=" + envir_yaml.Deploymentfile[0].Branch + "\n"
		write_content = write_content + deploy_Branch
		config_Branch := "CONFIGURATION_BRANCH=" + envir_yaml.Configuration[0].Branch + "\n"
		write_content = write_content + config_Branch
		default_namesapce := "DEFAULT_NAMESPACES=" + envir_yaml.Namespaces[0].K8S + "\n"
		write_content = write_content + default_namesapce
		openfaas_fn_namesapce := "OPENFAAS_FN_NAMESPACES=" + envir_yaml.Namespaces[0].Openfaas + "\n"
		write_content = write_content + openfaas_fn_namesapce
		monitor_namesapce := "MONITOR_NAMESPACES=" + envir_yaml.Namespaces[0].Monitor + "\n"
		write_content = write_content + monitor_namesapce
		redis_namesapce := "REDIS_NAMESPACES=" + envir_yaml.Namespaces[0].Redis + "\n"
		write_content = write_content + redis_namesapce
	}
	WriteWithIoutil(ouputfile, write_content)
}
