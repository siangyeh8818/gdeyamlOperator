package clusterop

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	gdeyamloperator "github.com/siangyeh8818/gdeyamlOperator/internal"
)

// AddResource is a method to create/apply k8s resources in the cluster
func AddResource() {

}

// PatchResource is a method to update those k8s resources in the cluster
func PatchResource() {

}

// DeleteResources is a method to delete those k8s resources in the cluster
func DeleteResources(git *gdeyamloperator.GIT) {
	fmt.Printf("DeleteResource: %v\n", git)

	// clone environment file
	gdeyamloperator.GitClone(git)

	// read and parse "deletion" section
	environmentFile, err := ioutil.ReadFile("clone/pn.environment/environment.yml")
	if err != nil {
		fmt.Printf("read environment file error: %v\n", err)
	}

	fmt.Printf("%v\n", string(environmentFile))

	envYaml := &gdeyamloperator.Environmentyaml{}
	err = yaml.Unmarshal(environmentFile, envYaml)
	if err != nil {
		log.Printf("Unmarshal error: %v\n", err)
	}
	fmt.Printf("envYaml: %v", envYaml)

	// get k8s client
	clientSet, err := getClientSet()
	if err != nil {
		panic(err)
	}

	// clone prune file
	newGit := &gdeyamloperator.GIT{
		Branch:      envYaml.Prune.Branch,
		Url:         envYaml.Prune.Git,
		AccessUser:  git.AccessUser,
		AccessToken: git.AccessToken,
	}
	gdeyamloperator.GitClone(newGit)

	pruneFile, err := ioutil.ReadFile("clone/pn.prune/prune.yml")
	if err != nil {
		fmt.Printf("read environment file error: %v\n", err)
	}

	fmt.Printf("%v\n", string(pruneFile))
	pruneYaml := &gdeyamloperator.PruneYaml{}
	if err := yaml.Unmarshal(pruneFile, pruneYaml); err != nil {
		log.Printf("Unmarshal error: %v\n", err)
	}

	fmt.Printf("%v\n", pruneYaml)
	fmt.Printf("%v\n", pruneYaml.Targets[0].Namespace)

	// Debug
	// createJob(clientSet, "workflow-stable", "demo-job")
	// time.Sleep(2 * time.Second)

	// dict := makeResourceDict(envYaml)
	// fmt.Printf("dict: %v\n", dict)

	// Delete
	for i := 0; i < len(pruneYaml.Targets); i++ {
		deleteResource(clientSet, pruneYaml.Targets[i])
	}

}

func deleteResource(cs *kubernetes.Clientset, deletion gdeyamloperator.PruneTarget) {
	switch deletion.Kind {
	case "namespace", "ns":
		deleteNamesapce(cs, deletion.Namespace, deletion.Name)
	case "job":
		deleteJob(cs, deletion.Namespace, deletion.Name)
	case "cronjob":
		deleteCronjob(cs, deletion.Namespace, deletion.Name)
	case "deploy", "deployment":
		deleteDeployment(cs, deletion.Namespace, deletion.Name)
	case "svc", "service":
		deleteService(cs, deletion.Namespace, deletion.Name)
	case "secret":
		deleteSecret(cs, deletion.Namespace, deletion.Name)
	case "cm", "configmap":
		deleteConfigMap(cs, deletion.Namespace, deletion.Name)
	case "statefulset":
		deleteStatefulSet(cs, deletion.Namespace, deletion.Name)
	case "ds", "daemonset":
		deleteDaemonSet(cs, deletion.Namespace, deletion.Name)
	case "ingress", "ing":
		deleteIngress(cs, deletion.Namespace, deletion.Name)
	case "pod", "po":
		deletePod(cs, deletion.Namespace, deletion.Name)
	default:
		fmt.Println("Notthing to delete")
		break
	}
}

// func makeResourceDict(envYaml *Environmentyaml) map[string][]string {
// 	// make resource dictionary
// 	resourceDict := make(map[string][]string)
// 	for i := 0; i < len(envYaml.Deletions); i++ {
// 		key := envYaml.Deletions[i].Namespace + "#" + envYaml.Deletions[i].Kind
// 		if val, ok := resourceDict[key]; ok {
// 			resourceDict[key] = append(val, envYaml.Deletions[i].Name)
// 		} else {
// 			resourceDict[key] = []string{envYaml.Deletions[i].Name}
// 		}
// 	}

// 	for k, v := range resourceDict {
// 		fmt.Printf("key[%s], value[%s]\n", k, v)
// 		namespace := strings.Split(k, "#")[0]
// 		kind := strings.Split(k, "#")[1]

// 		for _, name := range v {
// 			fmt.Printf("namespace: %s\n", namespace)

// 			fmt.Printf("kind: %s\n", kind)
// 			fmt.Printf("name: %v\n", name)
// 			fmt.Println("-------------")
// 		}
// 	}

// 	return resourceDict
// }

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func getClientSet() (*kubernetes.Clientset, error) {
	// actually perform k8s cluster operations
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientSet, err := kubernetes.NewForConfig(config)
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
	client := clientSet.BatchV1().Jobs(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		// panic(err)
		fmt.Printf("deleteJob Error: %v\n", err)
	} else {
		fmt.Printf("Successfully deleted job %v at %v namespace\n", name, namespace)
	}
}

func deleteCronjob(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.BatchV1beta1().CronJobs(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteCronjob Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted cronjob %v at %v namespace\n", name, namespace)
	}

}

func deleteNamesapce(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Namespaces()
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteNamesapce Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted namespace %v\n", name)
	}

}

func deletePod(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Pods(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deletePod Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted pod %v at %v namespace\n", name, namespace)
	}

}

func deleteDeployment(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().Deployments(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteDeployment Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted deployment %v at %v namespace\n", name, namespace)
	}

}

func deleteDaemonSet(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().DaemonSets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteDaemonSet Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted daemonset %v at %v namespace\n", name, namespace)
	}

}

func deleteStatefulSet(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.AppsV1().StatefulSets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteStatefulSet Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted statefulset %v at %v namespace\n", name, namespace)
	}

}

func deleteService(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Services(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteService Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted service %v at %v namespace\n", name, namespace)
	}

}

func deleteSecret(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().Secrets(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteSecret Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted secrets %v at %v namespace\n", name, namespace)
	}

}

func deleteConfigMap(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.CoreV1().ConfigMaps(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteConfigMap Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted configmaps %v at %v namespace\n", name, namespace)
	}
}

func deleteIngress(clientSet *kubernetes.Clientset, namespace string, name string) {
	client := clientSet.ExtensionsV1beta1().Ingresses(namespace)
	if err := client.Delete(name, &metav1.DeleteOptions{}); err != nil {
		fmt.Printf("deleteIngress Error: %v\n", err.Error())
	} else {
		fmt.Printf("Successfully deleted ingress %v at %v namespace\n", name, namespace)
	}
}

// ApplyByYaml creates or configures k8s resources via a yaml file
func ApplyByYaml() {

}

// DeleteByYaml remove related k8s resources via a yaml file
func DeleteByYaml() {

}
