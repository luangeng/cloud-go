package model

type Env struct {
	Key   string
	Value string
}

type Label struct {
	Key   string
	Value string
}

type Port struct {
	Name          string
	Port          int
	NodePort      int
	ContainerPort int
	TargetPort    int
}

type Quota struct {
	Memory string
	Cpu    string
}

type VolumeMount struct {
	Name string
	Path string
}

type Container struct {
	Name            string
	Image           string
	ImagePullPolicy string
	Envs            []Env
	Labels          []Label
	Ports           []Port
	Limit           Quota
	Request         Quota
	VolumeMounts    []VolumeMount
}

type Pod struct {
	Name       string
	Labels     []Label
	Containers []Container
}

type Volume struct {
	Name      string
	ClaimName string
}

type Deploy1 struct {
	Repilica   int
	Containers []Container
	Volumes    []Volume
}

type Pvc struct {
	Namespace string
	Name      string
	Capacity  int
}

type ServiceParam struct {
	Name      string
	Namespace string
}

type Statefulset struct {
	Name       string
	Repilica   int
	Containers []Container
}
