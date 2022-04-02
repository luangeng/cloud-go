package k8s

import (
	"bufio"
	"bytes"
	model "cloud/model"
	"context"
	"fmt"
	"io"
	"strings"

	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

func ListPod(ns string) ([]v1.Pod, error) {
	pods, err := GetClient().CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	return pods.Items, nil

	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	// namespace := "default"
	// pod := "me"
	// _, err = GetClient().CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	// if errors.IsNotFound(err) {
	// 	fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	// } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	// 	fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	// 		pod, namespace, statusError.ErrStatus.Message)
	// } else if err != nil {
	// 	panic(err.Error())
	// } else {
	// 	fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	// }
}

func CreatePod(pod0 model.Pod) error {
	var containers []v1.Container
	for _, c := range pod0.Containers {
		var envs []apiv1.EnvVar
		for _, env := range c.Envs {
			envs = append(envs, apiv1.EnvVar{Name: env.Key, Value: env.Value})
		}

		var ports []apiv1.ContainerPort
		for _, p := range c.Ports {
			ports = append(ports, apiv1.ContainerPort{Name: p.Name, ContainerPort: int32(p.ContainerPort)})
		}

		var volumeMounts []apiv1.VolumeMount
		for _, v := range c.VolumeMounts {
			volumeMounts = append(volumeMounts, apiv1.VolumeMount{Name: v.Name, MountPath: v.Path})
		}

		containers = append(containers, v1.Container{
			Name:            c.Name,
			Image:           c.Image,
			ImagePullPolicy: v1.PullIfNotPresent,
			Env:             envs,
			Ports:           ports,
			VolumeMounts:    volumeMounts,
			Resources: apiv1.ResourceRequirements{
				Requests: apiv1.ResourceList{
					apiv1.ResourceName(apiv1.ResourceCPU):    resource.MustParse("200m"),
					apiv1.ResourceName(apiv1.ResourceMemory): resource.MustParse("100Mi"),
				},
			},
		})
	}

	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:   pod0.Name,
			Labels: model.Labels2map(pod0.Labels),
		},
		Spec: v1.PodSpec{
			Containers: containers,
		},
	}
	_, err := GetClient().CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeletePod(ns string, name string) error {
	err := GetClient().CoreV1().Pods(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	return err
}

func LogPod() string {
	req := GetClient().CoreV1().Pods("default").GetLogs("demo-0", &v1.PodLogOptions{Follow: true})
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "error in opening stream"
	}
	defer podLogs.Close()

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		return "error in copy information from podLogs to buf"
	}
	str := buf.String()

	return str
}

func LogPodFollow(ctx context.Context, ch chan string) {
	defer close(ch)
	podLogOpts := v1.PodLogOptions{}
	podLogOpts.Follow = true
	podLogOpts.TailLines = &[]int64{int64(2000)}[0]
	podLogOpts.Container = "web"
	req := GetClient().CoreV1().Pods("default").GetLogs("demo-0", &podLogOpts)
	podLogs, err := req.Stream(ctx)
	if err != nil {
		fmt.Printf(err.Error())
		ch <- err.Error()
		return
	}
	defer podLogs.Close()

	reader := bufio.NewScanner(podLogs)
	for reader.Scan() {
		line := reader.Text()
		ch <- line
	}
}

func ExecPodOnce(stdin io.Reader) (string, error) {
	cmd := "curl localhost"
	req := GetClient().CoreV1().RESTClient().Post().Resource("pods").Name("demo-0").
		Namespace("default").SubResource("exec")
	req.VersionedParams(
		&v1.PodExecOptions{
			Container: "web",
			Command:   strings.Fields(cmd),
			Stdin:     stdin != nil,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		},
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return "", err
	}
	var stdout, stderr bytes.Buffer
	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})
	if err != nil {
		return "", err
	}
	return stdout.String() + "\n" + stderr.String(), nil
}

func ExecPod(stdin io.Reader, stdout io.Writer, stderr io.Writer) error {
	cmd := "bash"
	req := GetClient().CoreV1().RESTClient().Post().Resource("pods").Name("demo-0").
		Namespace("default").SubResource("exec")
	req.VersionedParams(
		&v1.PodExecOptions{
			Command: strings.Fields(cmd),
			Stdin:   true,
			Stdout:  true,
			Stderr:  true,
			TTY:     true,
		},
		scheme.ParameterCodec,
	)
	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}

	err = exec.Stream(remotecommand.StreamOptions{
		Stdin:  stdin,
		Stdout: stdout,
		Stderr: stderr,
		Tty:    true,
	})
	if err != nil {
		return err
	}

	return nil
}
