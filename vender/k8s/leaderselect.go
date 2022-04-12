package k8s

import (
	"context"
	"fmt"
	"os"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog"
)

func getNewLock(lockname, podname, namespace string) *resourcelock.LeaseLock {
	return &resourcelock.LeaseLock{
		LeaseMeta: metav1.ObjectMeta{
			Name:      lockname,
			Namespace: namespace,
		},
		Client: clientset.CoordinationV1(),
		LockConfig: resourcelock.ResourceLockConfig{
			Identity: podname,
		},
	}
}

func doStuff() {
	var i = 0
	for {
		fmt.Printf("doing stuff...%d\n", i)
		i++
		time.Sleep(3 * time.Second)
		if i > 10 {
			return
		}
	}
}

func runLeaderElection(lock *resourcelock.LeaseLock, ctx context.Context, id string) {
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(c context.Context) {
				doStuff()
			},
			OnStoppedLeading: func() {
				klog.Info("no longer the leader, staying inactive.")
			},
			OnNewLeader: func(current_id string) {
				if current_id == id {
					klog.Info("still the leader!")
					return
				}
				klog.Info("new leader is %s", current_id)
			},
		},
	})
}

func TestLock() {
	podName := os.Getenv("POD_NAME")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lock := getNewLock("mylock", podName, "default")
	runLeaderElection(lock, ctx, podName)
}
