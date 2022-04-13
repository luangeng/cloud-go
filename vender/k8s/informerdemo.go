package k8s

import (
	"encoding/json"
	"fmt"
	"time"

	apiCoreV1 "k8s.io/api/core/v1"
	"k8s.io/api/events/v1beta1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func PodNumOfSpecifyNode() {
	indexByPodNodeName := func(obj interface{}) ([]string, error) {
		pod, ok := obj.(*apiCoreV1.Pod)
		if !ok {
			return []string{}, nil
		}
		if len(pod.Spec.NodeName) == 0 || pod.Status.Phase == apiCoreV1.PodSucceeded || pod.Status.Phase == apiCoreV1.PodFailed {
			return []string{}, nil
		}
		return []string{pod.Spec.NodeName}, nil
	}
	// config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	// mustSuccess(err)

	// clientset, err := kubernetes.NewForConfig(config)
	// mustSuccess(err)

	sharedInformers := informers.NewSharedInformerFactory(clientset, 0)
	podInformer := sharedInformers.Core().V1().Pods().Informer()
	podInformer.GetIndexer().AddIndexers(cache.Indexers{
		"nodeName": indexByPodNodeName,
	})
	stopChan := make(chan struct{})
	defer close(stopChan)
	go podInformer.Run(stopChan)
	for range time.Tick(time.Millisecond * 5000) {
		podInformer.GetIndexer().ListKeys()
		nodeName := "u1"
		podList, err := podInformer.GetIndexer().ByIndex("nodeName", nodeName)
		mustSuccess(err)
		fmt.Printf("%s 上面有 %v 个pod处于Running或Pending中:\n", nodeName, len(podList))
	}

}

func WatchEvents() {
	sharedInformers := informers.NewSharedInformerFactory(clientset, 0)
	stopChan := make(chan struct{})
	defer close(stopChan)

	eventInformer := sharedInformers.Events().V1beta1().Events().Informer()
	// addChan := make(chan v1beta1.Event)

	eventInformer.AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			unstructObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
			mustSuccess(err)
			event := &v1beta1.Event{}
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructObj, event)
			mustSuccess(err)
			str, err := json.Marshal(event)
			fmt.Println(string(str))
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
		},
		DeleteFunc: func(obj interface{}) {
		},
	}, 0)

	eventInformer.Run(stopChan)
}

func mustSuccess(err error) {
	if err != nil {
		panic(err)
	}
}
