package model

type Pvc struct {
	Namespace string
	Name      string
	Capacity  int
}

type ServiceParam struct {
	Name      string
	Namespace string
}
