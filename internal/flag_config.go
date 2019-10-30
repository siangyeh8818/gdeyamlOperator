package gdeyamloperator

//import (
//	"flag"
//)

type BINARYCONFIG struct {
	ACTION string
	GIT_URL  string
	GIT_REPO_PATH  string
	GIT_USER  string
	GIT_TOKEN string
	GIT_BRANCH string
	GIT_NEW_BRANCH string
	GIT_TAG string
	GIT_ACTION string
	DOCKER_LOGIN string
	PUSH bool
	PUSH_PATTERN string
	PULL_PATTERN string
	IMAGENAME string
	LIST int
	LATEST_MODE string
	NEXUS_API_METHOD string
	NEXUS_REQ_BODY string
	NEXUS_OUTPUT_PATTERN string
	NEXUS_PROMOTE_TYPE string
	NEXUS_PROMOTE_DESTINATION string
	NEXUS_PROMOTE_URL string
	NEXUS_PROMOTE_SOURCE string
	ENVIRONMENT_FILE string
	SNAPSHOT_PATTERN string
	STAGE string
	REPLACE_TYPE string
	REPLACE_IMAGE string
	REPLACE_PATTERN string
	REPLACE_VALUE string
	KUSTOM_BASE_PATH string
	KUSTOMIZE_OUTPAT_DIR string
	KUSTOMIZE_RELPATH string
	KUSTOMIZE_URLPATTERN string
	KUSTOMIZE_MODULE string
	KUSTOMIZE_OPENFAAS_MODULE string
	KUSTOMIZE_COMPARE string
	KUSTOMIZE_BASE_FOLDER string
	NAMESPACE string
	USER string
	PASSWORD string
	INPUTFILE string
	OUTPUTFILE string
    VERSION bool
}

//func InitBinaryConfig () {

//}