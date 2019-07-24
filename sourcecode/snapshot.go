package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	//appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type K8sClient kubernetes.Clientset

func snapshot(pattern string, outputfilename string, kustomyamlfolder string) {

	test := K8sYaml{}
	var kubeconfig *string
	var n int
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
		panic(err.Error())
	}
	namespace_array, pattern_array := SnapshotPatternParser(pattern)

	for n = 0; n < len(namespace_array); n++ {

		fmt.Printf("%d namespace: %s\n", n, namespace_array[n])

		//對deployment做處理
		deploy_array := ListDeployment(clientSet, namespace_array[n])
		fmt.Println(len(deploy_array))
		for i := range deploy_array {
			if deploy_array[i] != "" && deploy_array[i] != "NAME" {
				fmt.Println("deployment name : " + deploy_array[i])
				imagename := GetDeploymentImage(clientSet, namespace_array[n], deploy_array[i])
				fmt.Println("Get deployment image name : " + imagename)
				modulestage, modulename, moduletag := ImagenameSplitReturnTag(imagename)

				switch pattern_array[n] {
				case "k8s":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddK8sStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddK8sStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddK8sStruct(deploy_array[i], modulename, moduletag, modulestage)
					}
				case "openfaas":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddOpenfaasStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddOpenfaasStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddOpenfaasStruct(deploy_array[i], modulename, moduletag, modulestage)
					}
				case "monitor":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddMonitorStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddMonitorStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddMonitorStruct(deploy_array[i], modulename, moduletag, modulestage)
					}
				case "redis":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddRedisStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddRedisStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddRedisStruct(deploy_array[i], modulename, moduletag, modulestage)
					}
				}
			}
		}
		//對statefulset做處理
		statefulset_array := ListStatefulSet(clientSet, namespace_array[n])
		fmt.Println(len(statefulset_array))
		for i := range statefulset_array {
			if statefulset_array[i] != "" && statefulset_array[i] != "NAME" {
				fmt.Println("statefulset name : " + statefulset_array[i])
				imagename := GetStatefulSetsImage(clientSet, namespace_array[n], statefulset_array[i])
				fmt.Println("Get deployment image name : " + imagename)
				modulestage, modulename, moduletag := ImagenameSplitReturnTag(imagename)

				switch pattern_array[n] {
				case "k8s":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddK8sStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddK8sStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddK8sStruct(statefulset_array[i], modulename, moduletag, modulestage)
					}
				case "openfaas":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddOpenfaasStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddOpenfaasStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddOpenfaasStruct(statefulset_array[i], modulename, moduletag, modulestage)
					}
				case "monitor":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddMonitorStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddMonitorStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddMonitorStruct(statefulset_array[i], modulename, moduletag, modulestage)
					}

				case "redis":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddRedisStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddRedisStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddRedisStruct(statefulset_array[i], modulename, moduletag, modulestage)
					}
				}
			}
		}
		//對daemonset做處理
		daemonset_array := ListDaemonset(clientSet, namespace_array[n])
		fmt.Println(len(daemonset_array))
		for i := range daemonset_array {
			if daemonset_array[i] != "" && daemonset_array[i] != "NAME" {
				fmt.Println("daemonset name : " + daemonset_array[i])
				imagename := GetDaemonsetImage(clientSet, namespace_array[n], daemonset_array[i])
				fmt.Println("Get daemonset image name : " + imagename)
				modulestage, modulename, moduletag := ImagenameSplitReturnTag(imagename)

				switch pattern_array[n] {
				case "k8s":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddK8sStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddK8sStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddK8sStruct(daemonset_array[i], modulename, moduletag, modulestage)
					}
				case "openfaas":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddOpenfaasStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddOpenfaasStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddOpenfaasStruct(daemonset_array[i], modulename, moduletag, modulestage)
					}
				case "monitor":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddMonitorStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddMonitorStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddMonitorStruct(daemonset_array[i], modulename, moduletag, modulestage)
					}
				case "redis":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddRedisStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddRedisStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddRedisStruct(daemonset_array[i], modulename, moduletag, modulestage)
					}
				}
			}
		}

		//clientSet.BatchV1beta1().CronJobs(namespace).Get(cronjobName, metav1.GetOptions{})
		//clientSet.BatchV1beta1().CronJobs()
		//對cronjob做處理
		cronjob_array := ListCronjob(clientSet, namespace_array[n])
		fmt.Println(len(cronjob_array))
		for i := range cronjob_array {
			if cronjob_array[i] != "" && cronjob_array[i] != "NAME" {
				fmt.Println("cronjob name : " + cronjob_array[i])
				imagename := GetCronjobImage(clientSet, namespace_array[n], cronjob_array[i])
				fmt.Println("Get cronjob image name : " + imagename)
				modulestage, modulename, moduletag := ImagenameSplitReturnTag(imagename)

				switch pattern_array[n] {
				case "k8s":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddK8sStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddK8sStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddK8sStruct(cronjob_array[i], modulename, moduletag, modulestage)
					}
				case "openfaas":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddOpenfaasStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddOpenfaasStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddOpenfaasStruct(cronjob_array[i], modulename, moduletag, modulestage)
					}
				case "monitor":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddMonitorStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddMonitorStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddMonitorStruct(cronjob_array[i], modulename, moduletag, modulestage)
					}
				case "redis":
					if len(kustomyamlfolder) > 0 {
						base_folder := grepFolderName(modulename, kustomyamlfolder)
						if len(base_folder) > 0 {
							(&test.Deployment).AddRedisStruct(base_folder, modulename, moduletag, modulestage)
						} else {
							fmt.Println("folder name can't be space")
							(&test.Deployment).AddRedisStruct("You_have_to_fix_base_repo", modulename, moduletag, modulestage)
						}
					} else {
						(&test.Deployment).AddRedisStruct(cronjob_array[i], modulename, moduletag, modulestage)
					}
				}
			}
		}

	}
	d, err := yaml.Marshal(&test)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	WriteWithIoutil(outputfilename, string(d))
}

func IdentifyOpenfaas(i_namesapce string, i_deployment string) bool {
	var i_token bool
	i_cmd := "kubectl get deploy -l faas_function -n " + i_namesapce + " | grep " + i_deployment
	i_cmd_result := RunCommand(i_cmd)

	if strings.Contains(i_cmd_result, i_deployment) {
		i_token = true
	} else {
		i_token = false
	}

	return i_token
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func SnapshotPatternParser(pattern string) ([]string, []string) {
	temp_array := strings.Split(pattern, ",")
	fmt.Println(temp_array)
	var ns_array []string
	var yamlstruct_array []string

	for i := 0; i < len(temp_array); i++ {
		content_temp_array := strings.Split(temp_array[i], ":")
		ns_array = append(ns_array, content_temp_array[1])
		yamlstruct_array = append(yamlstruct_array, content_temp_array[0])
		//ns_array[i] = content_temp_array[1]
		//yamlstruct_array[i] = content_temp_array[0]
	}
	//	fmt.Println(ns_array)
	//	fmt.Println(yamlstruct_array)
	return ns_array, yamlstruct_array
}

func ListDeployment(clientSet *kubernetes.Clientset, namespace string) []string {

	var result []string
	deploymentsclient := clientSet.ExtensionsV1beta1().Deployments(namespace)
	deployments, err := deploymentsclient.List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	} else {
		for i, e := range deployments.Items {
			log.Printf("Deployment #%d\n", i)
			log.Printf("%s", e.Name)
			result = append(result, e.Name)
			//log.Printf("%s", e.ObjectMeta.SelfLink)
		}
	}
	return result
}

func GetDeploymentImage(clientSet *kubernetes.Clientset, namespace string, deploymentName string) string {
	deployment, err := clientSet.AppsV1beta1().Deployments(namespace).Get(deploymentName, metav1.GetOptions{})
	var getimage string
	if err != nil {
		panic(err.Error())
	}
	if errors.IsNotFound(err) {
		fmt.Printf("Deployment not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting deployment%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found deployment\n")
		name := deployment.GetName()
		fmt.Println("name ->", name)
		containers := &deployment.Spec.Template.Spec.Containers
		//	found := false
		for i := range *containers {
			c := *containers
			getimage = c[i].Image
		}
		/*
							fmt.Println("Old image ->", c[i].Image)
							if c[i].Name == *appName {
								found = true
								fmt.Println("Old image ->", c[i].Image)
								fmt.Println("New image ->", *imageName)
								c[i].Image = *imageName
				}
			}
					if found == false {
						fmt.Println("The application container not exist in the deployment pods.")
						os.Exit(0)
					}
					_, err := clientset.AppsV1beta1().Deployments("default").Update(deployment)
					if err != nil {
						panic(err.Error())
					}*/
	}
	return getimage
}

func ListStatefulSet(clientSet *kubernetes.Clientset, namespace string) []string {

	var result []string
	statefulsetclient := clientSet.AppsV1().StatefulSets(namespace)
	statefulset, err := statefulsetclient.List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	} else {
		for i, e := range statefulset.Items {
			log.Printf("WStatefulset #%d\n", i)
			log.Printf("%s", e.Name)
			result = append(result, e.Name)
			//log.Printf("%s", e.ObjectMeta.SelfLink)
		}
	}
	return result
}

func GetStatefulSetsImage(clientSet *kubernetes.Clientset, namespace string, statefulsetName string) string {
	statefulset, err := clientSet.AppsV1().StatefulSets(namespace).Get(statefulsetName, metav1.GetOptions{})
	var getimage string
	if err != nil {
		panic(err.Error())
	}
	if errors.IsNotFound(err) {
		fmt.Printf("Statefulset not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting statefulset%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found statefulset\n")
		name := statefulset.GetName()
		fmt.Println("name ->", name)
		containers := &statefulset.Spec.Template.Spec.Containers
		for i := range *containers {
			c := *containers
			getimage = c[i].Image

		}
	}
	return getimage
}

func ListDaemonset(clientSet *kubernetes.Clientset, namespace string) []string {

	var result []string
	daemonsetclient := clientSet.ExtensionsV1beta1().DaemonSets(namespace)
	daemonsets, err := daemonsetclient.List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	} else {
		for i, e := range daemonsets.Items {
			log.Printf("Daemonset #%d\n", i)
			log.Printf("%s", e.Name)
			result = append(result, e.Name)
			//log.Printf("%s", e.ObjectMeta.SelfLink)
		}
	}
	return result
}

func GetDaemonsetImage(clientSet *kubernetes.Clientset, namespace string, daemonsetName string) string {
	daemonset, err := clientSet.ExtensionsV1beta1().DaemonSets(namespace).Get(daemonsetName, metav1.GetOptions{})
	var getimage string
	if err != nil {
		panic(err.Error())
	}
	if errors.IsNotFound(err) {
		fmt.Printf("Daemonset not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting daemonset%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found daemonset\n")
		name := daemonset.GetName()
		fmt.Println("name ->", name)
		containers := &daemonset.Spec.Template.Spec.Containers
		for i := range *containers {
			c := *containers
			getimage = c[i].Image

		}
	}
	return getimage
}

func ListCronjob(clientSet *kubernetes.Clientset, namespace string) []string {

	var result []string
	cronjobclient := clientSet.BatchV1beta1().CronJobs(namespace)
	cronjobs, err := cronjobclient.List(metav1.ListOptions{})
	if err != nil {
		log.Println(err)
	} else {
		for i, e := range cronjobs.Items {
			log.Printf("Cronjob #%d\n", i)
			log.Printf("%s", e.Name)
			if IdentifyCrongob(e.Name) {
				result = append(result, e.Name)
			} else {
				log.Printf("%s is not nedd to snapshot", e.Name)
			}

			//log.Printf("%s", e.ObjectMeta.SelfLink)
		}
	}
	return result
}

func GetCronjobImage(clientSet *kubernetes.Clientset, namespace string, cronjobName string) string {
	cronjob, err := clientSet.BatchV1beta1().CronJobs(namespace).Get(cronjobName, metav1.GetOptions{})
	//	daemonset, err := clientSet.ExtensionsV1beta1().DaemonSets(namespace).Get(daemonsetName, metav1.GetOptions{})
	var getimage string
	if err != nil {
		panic(err.Error())
	}
	if errors.IsNotFound(err) {
		fmt.Printf("CronJob not found\n")
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting CronJob%v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found CronJob\n")
		name := cronjob.GetName()
		fmt.Println("name ->", name)
		containers := &cronjob.Spec.JobTemplate.Spec.Template.Spec.Containers
		for i := range *containers {
			c := *containers
			getimage = c[i].Image

		}
	}
	return getimage
}
