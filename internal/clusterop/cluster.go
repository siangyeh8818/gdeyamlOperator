package clusterop

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v3"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	mygit "github.com/siangyeh8818/gdeyamlOperator/internal/git"
	IO "github.com/siangyeh8818/gdeyamlOperator/internal/myIo"
	CustomStruct "github.com/siangyeh8818/gdeyamlOperator/internal/structs"
)

// AddResource is a method to create/apply k8s resources in the cluster
func AddResource() {

}

// PatchResource is a method to update those k8s resources in the cluster
func PatchResource() {

}

// DeleteResources is a method to delete those k8s resources in the cluster
func DeleteResources(git *mygit.GIT) {
	fmt.Printf("DeleteResource: %v\n", git)

	// clone environment file
	mygit.GitClone(git)

	// read and parse "deletion" section
	environmentFile, err := ioutil.ReadFile("clone/pn.environment/environment.yml")
	if err != nil {
		fmt.Printf("read environment file error: %v\n", err)
	}

	fmt.Printf("%v\n", string(environmentFile))

	envYaml := &CustomStruct.Environmentyaml{}
	err = yaml.Unmarshal(environmentFile, envYaml)
	if err != nil {
		log.Printf("Unmarshal error: %v\n", err)
	}
	fmt.Printf("envYaml: %v\n", envYaml)
	if (CustomStruct.Prune{}) == envYaml.Prune {
		// if we don't have a prune section
		fmt.Printf("Prune section is empty !!!!!!")
		return
	} else if envYaml.Prune.Git == "" || envYaml.Prune.Branch == "" {
		fmt.Printf("Prune.Git/Prune.Branch is empty !!!!!!")
		return
	}
	// clone prune file
	newGit := &mygit.GIT{
		Branch:      envYaml.Prune.Branch,
		Url:         envYaml.Prune.Git,
		AccessUser:  git.AccessUser,
		AccessToken: git.AccessToken,
	}
	mygit.GitClone(newGit)

	pruneFile, err := ioutil.ReadFile("clone/pn.prune/prune.yml")
	if err != nil {
		fmt.Printf("read clone/pn.prune/prune.yml error: %v\n", err)
	}

	fmt.Printf("%v\n", string(pruneFile))
	pruneYaml := &CustomStruct.PruneYaml{}
	if err := yaml.Unmarshal(pruneFile, pruneYaml); err != nil {
		log.Printf("Unmarshal error: %v\n", err)
	}

	fmt.Printf("%v\n", pruneYaml)

	// get k8s client
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	// Debug
	// createJob(clientSet, "workflow-stable", "demo-job")
	// time.Sleep(2 * time.Second)
	// Delete
	for i := 0; i < len(pruneYaml.Targets); i++ {
		deleteResource(clientSet, pruneYaml.Targets[i], envYaml)
	}
}

//在查找字段時使用Value.FieldByNameFunc和strings.ToLower忽略大小寫
func caseInsenstiveFieldByName(v reflect.Value, name string) reflect.Value {
	name = strings.ToLower(name)
	return v.FieldByNameFunc(func(n string) bool { return strings.ToLower(n) == name })
}

func deleteResource(cs *kubernetes.Clientset, deletion CustomStruct.PruneTarget, environment *CustomStruct.Environmentyaml) {
	var tempnamespace string

	switch strings.ToLower(deletion.Namespace.Type) {
	case "local":
		tempnamespace = deletion.Namespace.MappingNs
	case "environment":

		reflectStruct := reflect.TypeOf(environment)
		fmt.Println("-----reflect-------")
		fmt.Println(reflectStruct.Kind())
		if reflectStruct.Kind() == reflect.Struct {
			fmt.Println("返回是個結構(Struct)")
			fmt.Println("reflectStruct.NumField() : %d", reflectStruct.NumField())
			for i := 0; i < reflectStruct.NumField(); i++ {
				fmt.Println(reflectStruct.Field(i))
			}
		} else if reflectStruct.Kind() == reflect.Ptr {
			fmt.Println("返回是個指針(Ptr) , 使用Elem()獲取其指向的元素")
			reflectStruct = reflectStruct.Elem()
			fmt.Printf("再次測試kind : %s \n", reflectStruct.Kind().String())
			fmt.Printf("該struct的Field個數 : %d \n", reflectStruct.NumField())
			for i := 0; i < reflectStruct.NumField(); i++ {
				//fmt.Println(reflectStruct.Field(i))
				fmt.Printf("reflectStruct 欄位名稱 : %s\n", reflectStruct.Field(i).Name)
				//fmt.Println(reflectStruct.Field(i).Type)
				if reflectStruct.Field(i).Name == "Namespaces" {
					fmt.Println("取得環境檔中Namespaces的struct")
					NsStruct := reflect.TypeOf(environment.Namespaces[0])
					//fmt.Println(NsStruct.Kind())
					fmt.Println("NsStruct.NumField() : %d", NsStruct.NumField())
					for k := 0; k < NsStruct.NumField(); k++ {
						//fmt.Println(NsStruct.Field(k))
						if strings.ToLower(NsStruct.Field(k).Name) == deletion.Namespace.MappingNs {
							fmt.Printf("prune.yml中的MappingNs 符合返回的struct內的欄位名稱 , 欄位: %s \n", strings.ToLower(NsStruct.Field(k).Name))
							NsValue := reflect.ValueOf(environment.Namespaces[0])
							fmt.Printf("取得上述欄位名稱在environment.yml中的值 , 值為: %s  \n", NsValue.Field(k))
							//fmt.Println(NsValue.Field(k))
							tempnamespace = NsValue.Field(k).String()

						}
					}
				}
			}
		}
	}

	fmt.Printf("tempnamespace : %s \n", tempnamespace)

	switch strings.ToLower(deletion.Kind) {
	case "namespace", "ns":
		deleteNamespace(cs, deletion.Name)
	case "job":
		deleteJob(cs, tempnamespace, deletion.Name)
	case "cronjob":
		deleteCronjob(cs, tempnamespace, deletion.Name)
	case "deploy", "deployment":
		deleteDeployment(cs, tempnamespace, deletion.Name)
	case "svc", "service":
		deleteService(cs, tempnamespace, deletion.Name)
	case "secret":
		deleteSecret(cs, tempnamespace, deletion.Name)
	case "cm", "configmap":
		deleteConfigMap(cs, tempnamespace, deletion.Name)
	case "statefulset":
		deleteStatefulSet(cs, tempnamespace, deletion.Name)
	case "ds", "daemonset":
		deleteDaemonSet(cs, tempnamespace, deletion.Name)
	case "ingress", "ing":
		deleteIngress(cs, tempnamespace, deletion.Name)
	case "pod", "po":
		deletePod(cs, tempnamespace, deletion.Name)
	case "clusterrole":
		deleteClusterRole(cs, deletion.Name)
	case "clusterrolebinding":
		deleteClusterRoleBinding(cs, deletion.Name)
	case "sa", "serviceaccount":
		deleteServiceAccount(cs, tempnamespace, deletion.Name)
	case "hpa", "horizontalpodautoscaler":
		deleteHorizontalPodAutoscaler(cs, tempnamespace, deletion.Name)
	default:
		fmt.Println("Notthing to delete")
		break
	}

}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getKubeConfig() (*rest.Config, error) {
	// actually perform k8s cluster operations
	// try to read environment variable KUBECONFIG
	var kubeConfig *string
	if os.Getenv("KUBECONFIG") != "" {
		fmt.Printf("KUBECONFIG: %v\n", os.Getenv("KUBECONFIG"))
		kubeConfig = flag.String("kubeconfig", os.Getenv("KUBECONFIG"), "")
	} else {
		if home := homeDir(); home != "" {
			fmt.Printf("if homeDir: %v\n", home)

			if IO.Exists(filepath.Join(home, ".kube", "config")) {
				kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
				fmt.Println("K8s Config is exist")
			} else {
				kubeConfig = flag.String("kubeconfig", " ", "absolute path to the kubeconfig file")
			}

		} else {
			fmt.Printf("else homeDir: %v\n", home)
			kubeConfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
		}
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	return config, nil
}

func getClientSet() (*kubernetes.Clientset, error) {
	kubeConfig, err := getKubeConfig()
	if err != nil {
		panic(err)
	}

	// fmt.Printf("kubeConfig: %v\n", kubeConfig.GoString())
	// create the clientset
	clientSet, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return clientSet, nil
}

func createJob(cs *kubernetes.Clientset, ns string, name string) {

	jobsClient := cs.BatchV1().Jobs(ns)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-job",
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "demo-job",
							Image: "busybox",
						},
					},
					RestartPolicy: "OnFailure",
				},
			},
		},
	}

	// Create
	if job, err := jobsClient.Create(job); err != nil {
		// panic(err)
		fmt.Printf("job: %v\n", job)
		fmt.Printf("create err: %v\n", err)
	}

	fmt.Printf("job: %v\n", job)
}

func deleteJob(clientSet *kubernetes.Clientset, namespace string, name string) {
	Policyconst := metav1.DeletePropagationForeground
	client := clientSet.BatchV1().Jobs(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &Policyconst,
	}); err != nil {
		fmt.Printf("Delete Job Error: %v\n", err)
	} else {
		fmt.Printf("Successfully deleted job %v at %v namespace\n", name, namespace)
	}
}

func deleteCronjob(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.BatchV1beta1().CronJobs(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete CronJob Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted cronjob %v at %v namespace\n", name, namespace)
	}

}

func deleteNamespace(clientSet *kubernetes.Clientset, name string) {
	client := clientSet.CoreV1().Namespaces()
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Namespace Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted namespace %v\n", name)
	}

}

func deletePod(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Pods(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Pod Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted pod %v at %v namespace\n", name, namespace)
	}

}

func deleteDeployment(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().Deployments(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Deployment Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted deployment %v at %v namespace\n", name, namespace)
	}

}

func deleteDaemonSet(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().DaemonSets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete DaemonSet Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted daemonset %v at %v namespace\n", name, namespace)
	}

}

func deleteStatefulSet(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().StatefulSets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete StatefulSet Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted statefulset %v at %v namespace\n", name, namespace)
	}

}

func deleteService(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Services(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Service Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted service %v at %v namespace\n", name, namespace)
	}

}

func deleteSecret(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Secrets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Secret Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted secrets %v at %v namespace\n", name, namespace)
	}

}

func deleteConfigMap(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().ConfigMaps(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete ConfigMap Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted configmaps %v at %v namespace\n", name, namespace)
	}
}

func deleteIngress(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.ExtensionsV1beta1().Ingresses(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete Ingress Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted ingress %v at %v namespace\n", name, namespace)
	}
}

func deleteClusterRole(clientSet *kubernetes.Clientset, name string) {
	client := clientSet.RbacV1().ClusterRoles()
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete ClusterRole Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted clusterrole %v\n", name)
	}
}

func deleteClusterRoleBinding(clientSet *kubernetes.Clientset, name string) {
	client := clientSet.RbacV1().ClusterRoleBindings()
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete ClusterRoleBinding Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted clusterrolebinding %v\n", name)
	}
}

func deleteServiceAccount(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().ServiceAccounts(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete ServiceAccounts Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted serviceaccount %v at %v namespace\n", name, namespace)
	}
}

func deleteHorizontalPodAutoscaler(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AutoscalingV1().HorizontalPodAutoscalers(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("Delete HorizontalPodAutoscaler Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted hpa %v at %v namespace\n", name, namespace)
	}
}

// ApplyByYaml creates or configures k8s resources via a yaml file
func ApplyByYaml() {

}

// DeleteByYaml remove related k8s resources via a yaml file
func DeleteByYaml() {

}
