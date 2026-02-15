package output

import (
	"fmt"

	"github.com/cryptologicpsy/kubectl-audit/pkg/checks"
)

// PrintTable τυπώνει τα findings σε απλό table
func PrintTable(findings []checks.Finding) {
	fmt.Printf("%-15s %-20s %-40s %-10s\n", "NAMESPACE", "RESOURCE", "ISSUE", "SEVERITY")
	for _, f := range findings {
		fmt.Printf("%-15s %-20s %-40s %-10s\n", f.Namespace, f.Resource, f.Issue, f.Severity)
	}
}
