package main

import (
	"fmt"

	"github.com/cryptologicpsy/kubectl-audit/pkg/checks"
	"github.com/cryptologicpsy/kubectl-audit/pkg/output"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	fmt.Println("kubectl-audit starting...")

	pods := []v1.Pod{
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "prod", Name: "pod-1"},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{Name: "app", Resources: v1.ResourceRequirements{}}, // no limits
				},
			},
		},
		{
			ObjectMeta: metav1.ObjectMeta{Namespace: "prod", Name: "pod-2"},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name: "api",
						Resources: v1.ResourceRequirements{
							Limits: v1.ResourceList{
								v1.ResourceCPU:    resource.MustParse("500m"),
								v1.ResourceMemory: resource.MustParse("256Mi"),
							},
						},
					},
				},
			},
		},
	}

	// run checks
	findings := checks.RunAllPodsChecks(pods)

	// print results
	output.PrintTable(findings)
}
