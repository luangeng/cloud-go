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
	for {
		fmt.Printf("I am the leader with lock\n")
		time.Sleep(4 * time.Second)
	}
}

func runLeaderElection(lock *resourcelock.LeaseLock, ctx context.Context, cancel context.CancelFunc, id string) {
	fmt.Printf("runLeaderElection\n")
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:            lock,
		ReleaseOnCancel: true,
		LeaseDuration:   15 * time.Second,
		RenewDeadline:   10 * time.Second,
		RetryPeriod:     2 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(c context.Context) {
				fmt.Printf("leader gets lock\n")
				doStuff()
				cancel()
			},
			OnStoppedLeading: func() {
				klog.Info("no longer the leader, staying inactive.")
			},
			OnNewLeader: func(current_id string) {
				if current_id == id {
					klog.Info("still the leader!")
					return
				}
				klog.Info("leader changed, new leader is %s", current_id)
			},
		},
	})
}

func TestLock() {
	podName := os.Getenv("POD_NAME")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lock := getNewLock("mylock", podName, "default")
	runLeaderElection(lock, ctx, cancel, podName)
}
