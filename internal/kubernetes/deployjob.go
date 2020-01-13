package kubernetes

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	IO "github.com/siangyeh8818/gdeyamlOperator/internal/myIo"
	CustomStruct "github.com/siangyeh8818/gdeyamlOperator/internal/structs"
)

func CeateScriptsJob(yamlfile string, environment string) {
	var jobname string
	var completeImageName string
	var jobNs string

	if yamlfile != "" && IO.Exists(yamlfile) {
		deployYaml := CustomStruct.K8sYaml{}
		deployYaml.GetConf(yamlfile)
		jobname = deployYaml.Deployment.SCRIPTS.TOOL.Image
		completeImageName = deployYaml.Deployment.SCRIPTS.TOOL.Image + ":" + deployYaml.Deployment.SCRIPTS.TOOL.Tag
	}
	if environment != "" && IO.Exists(environment) {
		envirYaml := CustomStruct.Environmentyaml{}
		envirYaml.GetConf(environment)
		jobNs = envirYaml.Namespaces[0].K8S
	}

	clientSets, err := getClientSet()
	if err != nil {
		panic(err)
	}
	completeImageName = "cr.pentium.network/" + completeImageName

	fmt.Println("--------check argv of func createJob(clientSets, jobname, jobNs, completeImageName)-------")
	fmt.Printf("jobname: %s\n", jobname)
	fmt.Printf("jobNs: %s\n", jobNs)
	fmt.Printf("completeImageName: %s\n", completeImageName)
	fmt.Println("--------------------------------------")
	createJob(clientSets, jobNs, jobname, completeImageName)

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
			kubeConfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
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

func createJob(cs *kubernetes.Clientset, ns string, jobname string, imagename string) {

	jobsClient := cs.BatchV1().Jobs(ns)

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			Name: jobname,
		},
		Spec: batchv1.JobSpec{
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  jobname,
							Image: imagename,
							Env: []apiv1.EnvVar{
								{
									Name:  "PN_GLOBAL_SCRIPT_URLS",
									Value: "",
								}, {
									Name: "FAAS_ID",
									ValueFrom: &apiv1.EnvVarSource{
										SecretKeyRef: &apiv1.SecretKeySelector{
											Key: "basic-auth-password",
										},
									},
								}, {
									Name: "FAAS_PASSWORD",
									ValueFrom: &apiv1.EnvVarSource{
										SecretKeyRef: &apiv1.SecretKeySelector{
											Key: "basic-auth-password",
										},
									},
								},
							},
						},
					},
					RestartPolicy: "Never",
					Volumes: []apiv1.Volume{
						{
							Name: "pn-config",
							VolumeSource: apiv1.VolumeSource{
								Secret: &apiv1.SecretVolumeSource{
									SecretName: "pn-config",
								},
							},
						},
					},
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
		fmt.Printf("Delete Job Error: %v\n", err)
	} else {
		fmt.Printf("Successfully deleted job %v at %v namespace\n", name, namespace)
	}
}
