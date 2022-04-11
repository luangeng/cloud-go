package k8s

import (
	"fmt"
	"time"

	apiCoreV1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

func podNumOfSpecifyNode() {
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
	for range time.Tick(time.Millisecond * 1000) {
		podInformer.GetIndexer().ListKeys()
		nodeName := "u1"
		podList, err := podInformer.GetIndexer().ByIndex("nodeName", nodeName)
		mustSuccess(err)
		fmt.Printf("%s 上面有 %v 个pod处于Running或Pending中:\n", nodeName, len(podList))
	}

}

func mustSuccess(err error) {
	if err != nil {
		panic(err)
	}
}
