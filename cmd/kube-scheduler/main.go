package main

import (
	"github.com/dimitrisdol/socketquadScheduler/socketquad"

	"k8s.io/klog/v2"
	sched "k8s.io/kubernetes/cmd/kube-scheduler/app"
)

func main() {
	cmd := sched.NewSchedulerCommand(
		sched.WithPlugin(socketquad.Name, socketquad.New),
	)
	if err := cmd.Execute(); err != nil {
		klog.Fatalf("failed to execute %q: %v", socketquad.Name, err)
	}
}
