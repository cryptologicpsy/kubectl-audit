# kubectl-audit

A simple kubectl plugin written in Go to audit Kubernetes pods for missing resource limits.

## Description

This project is a hands-on DevOps exercise to create a kubectl plugin that:

- Connects to Kubernetes (or uses mock pods for testing)
- Checks pods for containers without CPU or memory resource limits
- Prints the findings in a simple table format

Currently, it uses **mock pods** so you can run it without having Kubernetes installed.

## Usage

Run the plugin locally:

\`\`\`bash
go run main.go
\`\`\`

Example output:

\`\`\`
kubectl-audit starting...
NAMESPACE       RESOURCE             ISSUE                                    SEVERITY
prod            pod-1                Container app has no resource limits     HIGH
\`\`\`

## Future Improvements

- Connect to a real Kubernetes cluster using kubeconfig
- Add more checks, e.g., privileged containers, missing requests, etc.
- Improve output formatting and CLI options

## Tech Stack

- Go
- Kubernetes client-go API
- CLI / Table output
