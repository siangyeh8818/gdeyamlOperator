package myIo

import (
	"log"

	CustomStruct "github.com/siangyeh8818/gdeyamlOperator/internal/structs"
)

func DumpArguments(inputfile string, environment_file string, ouputfile string) {

	var write_content string
	if inputfile != "" && Exists(inputfile) {
		deploy_yaml := CustomStruct.K8sYaml{}
		deploy_yaml.GetConf(inputfile)
		log.Println(deploy_yaml.Deployment)
		blcks_Branch := "BLCKS_BRANCH=" + deploy_yaml.Deployment.BLCKS.Branch + "\n"
		write_content = write_content + blcks_Branch
		blcks_version := "BLCKS_VERSION=" + deploy_yaml.Deployment.BLCKS.Version + "\n"
		write_content = write_content + blcks_version
		blcks_tool_tag := "BLCKS_TOOL_TAG=" + deploy_yaml.Deployment.BLCKS.TOOL.Tag + "\n"
		write_content = write_content + blcks_tool_tag
		playbook_Branch := "PLAYBOOK_BRANCH=" + deploy_yaml.Deployment.PLAYBOOKS.Branch + "\n"
		write_content = write_content + playbook_Branch
		playbook_version := "PLAYBOOK_VERSION=" + deploy_yaml.Deployment.PLAYBOOKS.Version + "\n"
		write_content = write_content + playbook_version
		playbook_tool_tag := "PLAYBOOK_TOOL_TAG=" + deploy_yaml.Deployment.PLAYBOOKS.TOOL.Tag + "\n"
		write_content = write_content + playbook_tool_tag
		kubestomize_base_Branch := "BASE_BRANCH=" + deploy_yaml.Deployment.BASE[0].Branch + "\n"
		write_content = write_content + kubestomize_base_Branch
	}
	if environment_file != "" && Exists(inputfile) {
		envir_yaml := CustomStruct.Environmentyaml{}
		envir_yaml.GetConf(environment_file)
		Portalv2Domain := "DomainPortalV2=" + envir_yaml.Domain.DomainPortalV2 + "\n"
		write_content = write_content + Portalv2Domain
		Portalv3Domain := "DomainPortalV3=" + envir_yaml.Domain.DomainPortalV3 + "\n"
		write_content = write_content + Portalv3Domain
		deploy_Branch := "DEPLOY_BRANCH=" + envir_yaml.Deploymentfile[0].Branch + "\n"
		write_content = write_content + deploy_Branch
		config_Branch := "CONFIGURATION_BRANCH=" + envir_yaml.Configuration[0].Branch + "\n"
		write_content = write_content + config_Branch
		default_namesapce := "DEFAULT_NAMESPACES=" + envir_yaml.Namespaces[0].K8S + "\n"
		write_content = write_content + default_namesapce
		openfaas_fn_namesapce := "OPENFAAS_FN_NAMESPACES=" + envir_yaml.Namespaces[0].Openfaas + "\n"
		write_content = write_content + openfaas_fn_namesapce

		openfaas_namesapce := "FAAS_NETES_NAMESPACES=" + envir_yaml.Namespaces[0].FaasNets + "\n"
		write_content = write_content + openfaas_namesapce

		monitor_namesapce := "MONITOR_NAMESPACES=" + envir_yaml.Namespaces[0].Monitor + "\n"
		write_content = write_content + monitor_namesapce
		redis_namesapce := "REDIS_NAMESPACES=" + envir_yaml.Namespaces[0].Redis + "\n"
		write_content = write_content + redis_namesapce
	}
	WriteWithIoutil(ouputfile, write_content)
}
