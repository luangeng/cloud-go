package model

type Node struct {
	Name             string `json:"name"`
	Status           string
	Roles            string
	Age              string
	Version          string
	INTERNALIP       string
	EXTERNALIP       string
	OsImage          string
	KERNELVERSION    string
	CONTAINERRUNTIME string
}

type Pv struct {
	Name         string `json:"name"`
	StorageClass string `json:"storageClass"`
	Capacity     string `json:"capacity"`
	Status       string `json:"status"`
	AccessModes  string `json:"accessModes"`
}

type Deploy struct {
	Name     string `json:"name"`
	Replicas int    `json:"replicas"`
}
