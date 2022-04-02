package k8s

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func Cmd(c string) (string, error) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("", strings.Fields(c)...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		return "", err
	}
	err = cmd.Wait()
	if err != nil {
		return "", err
	}
	out := append(stdout.Bytes(), stderr.Bytes()...)
	fmt.Printf("Output:\n%s\n", string(out))
	return string(out), nil
}
