package main

import (
	"flag"
	"fmt"
	"path/filepath"

	//appsv1 "k8s.io/api/apps/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//type K8sClient kubernetes.Clientset

func DumpImage(push_pattern string, snapshot_pattern string, pushimage bool) {

	//test := K8sYaml{}
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

	namespace_array, _ := SnapshotPatternParser(snapshot_pattern)

	for n = 0; n < len(namespace_array); n++ {

		fmt.Printf("%d namespace: %s\n", n, namespace_array[n])

		deploy_array := ListDeployment(clientSet, namespace_array[n])
		fmt.Println(len(deploy_array))
		for i := range deploy_array {
			if deploy_array[i] != "" && deploy_array[i] != "NAME" {
				fmt.Println("deployment name : " + deploy_array[i])
				imagename := GetDeploymentImage(clientSet, namespace_array[n], deploy_array[i])
				fmt.Println("Get deployment image name : " + imagename)
				_, modulename, moduletag := ImagenameSplitReturnTag(imagename)
				if pushimage == true {
					PushTagimage(imagename, push_pattern, modulename, moduletag)
				}
			}
		}
		statefulset_array := ListStatefulSet(clientSet, namespace_array[n])
		fmt.Println(len(statefulset_array))
		for i := range statefulset_array {
			if statefulset_array[i] != "" && statefulset_array[i] != "NAME" {
				fmt.Println("statefulset name : " + statefulset_array[i])
				imagename := GetStatefulSetsImage(clientSet, namespace_array[n], statefulset_array[i])
				fmt.Println("Get deployment image name : " + imagename)
				_, modulename, moduletag := ImagenameSplitReturnTag(imagename)
				if pushimage == true {
					PushTagimage(imagename, push_pattern, modulename, moduletag)
				}
			}
		}
		daemonset_array := ListDaemonset(clientSet, namespace_array[n])
		fmt.Println(len(daemonset_array))
		for i := range daemonset_array {
			if daemonset_array[i] != "" && daemonset_array[i] != "NAME" {
				fmt.Println("daemonset name : " + daemonset_array[i])
				imagename := GetDaemonsetImage(clientSet, namespace_array[n], daemonset_array[i])
				fmt.Println("Get daemonset image name : " + imagename)
				_, modulename, moduletag := ImagenameSplitReturnTag(imagename)
				if pushimage == true {
					PushTagimage(imagename, push_pattern, modulename, moduletag)
				}
			}
		}
	}
}
