package checks

import (
	"fmt"

	v1 "k8s.io/api/core/v1"
)

type Finding struct {
	Namespace string
	Resource  string
	Issue     string
	Severity  string
}

// Πρώτο check: pods χωρίς resource limits
func RunAllPodsChecks(pods []v1.Pod) []Finding {
	var findings []Finding
	for _, pod := range pods {
		for _, c := range pod.Spec.Containers {
			if c.Resources.Limits == nil || len(c.Resources.Limits) == 0 {
				findings = append(findings, Finding{
					Namespace: pod.Namespace,
					Resource:  pod.Name,
					Issue:     fmt.Sprintf("Container %s has no resource limits", c.Name),
					Severity:  "HIGH",
				})
			}
		}
	}
	return findings
}
