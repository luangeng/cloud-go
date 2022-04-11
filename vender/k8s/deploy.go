package k8s

import (
	"cloud/model"
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func CreateDeploy(d model.Deploy1) error {
	var containers []apiv1.Container
	for _, c := range d.Containers {

		var envs []apiv1.EnvVar
		for _, env := range c.Envs {
			envs = append(envs, apiv1.EnvVar{Name: env.Key, Value: env.Value})
		}

		var ports []apiv1.ContainerPort
		for _, port := range c.Ports {
			ports = append(ports, apiv1.ContainerPort{
				Name:          port.Name,
				Protocol:      apiv1.ProtocolTCP,
				ContainerPort: int32(port.ContainerPort),
			})
		}

		var volumeMounts []apiv1.VolumeMount
		for _, v := range c.VolumeMounts {
			volumeMounts = append(volumeMounts, apiv1.VolumeMount{
				Name:      v.Name,
				MountPath: v.Path,
			})
		}

		containers = append(containers, apiv1.Container{
			Name:            c.Name,
			Image:           c.Image,
			ImagePullPolicy: "IfNotPresent",
			Ports:           ports,
			Env:             envs,
			VolumeMounts:    volumeMounts,
		})
	}

	var volumes []apiv1.Volume
	for _, v := range d.Volumes {
		volumes = append(volumes, apiv1.Volume{
			Name: v.Name,
			VolumeSource: apiv1.VolumeSource{
				PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
					ClaimName: v.ClaimName,
				},
			},
		})
	}

	deploymentsClient := GetClient().AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: containers,
					Volumes:    volumes,
				},
			},
		},
	}

	_, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	return err
}

func ListDeploy(ns string) ([]appsv1.Deployment, error) {
	deploymentsClient := GetClient().AppsV1().Deployments(ns)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return list.Items, nil
}

func DeleteDeploy(ns string, name string) error {
	deploymentsClient := GetClient().AppsV1().Deployments(ns)
	deletePolicy := metav1.DeletePropagationForeground
	err := deploymentsClient.Delete(context.TODO(), name, metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	return err
}

func UpdateDeploy() error {
	deploymentsClient := GetClient().AppsV1().Deployments(apiv1.NamespaceDefault)
	fmt.Println("Updating deployment...")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(context.TODO(), "demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			fmt.Errorf("Failed to get latest version of Deployment: %v", getErr)
		}

		result.Spec.Replicas = int32Ptr(1)
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	return retryErr
}

func int32Ptr(i int32) *int32 { return &i }
