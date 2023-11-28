package main

import (
	"github.com/mheers/k8s-secret-ui/cmd"
	"k8s.io/klog/v2"
)

func main() {
	// execute the command
	err := cmd.Execute()
	if err != nil {
		klog.Fatalf("Execute failed: %+v", err)
	}
}
