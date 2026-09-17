# K8s Multi-Cluster Orchestrator

> **High-performance Kubernetes Operator built with Go, controller-runtime, and Kubebuilder for orchestrating multi-region cluster deployments, global service discovery, and zero-downtime traffic failover.**

[![Go Report Card](https://goreportcard.com/badge/github.com/rennanfsimoes/k8s-multicluster-orchestrator)](https://goreportcard.com)
[![Kubernetes](https://img.shields.io/badge/K8s-1.28+-326CE5.svg)](https://kubernetes.io)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-purple.svg)](LICENSE)

---

##  Overview

The **K8s Multi-Cluster Orchestrator** reconciles Custom Resource Definitions (`GlobalDeployment`, `ClusterGroup`) across a mesh of hybrid and multi-cloud Kubernetes clusters (EKS, GKE, AKS, On-Premises).

### Features
- **Custom Resource Definitions (CRDs)**: Declarative multi-cluster deployment definitions with canary rollout support.
- **Health-Aware Load Shedding**: Detects control plane degradation or regional outages and shifts ingress traffic dynamically.
- **GitOps Integration**: Works seamlessly with ArgoCD ApplicationSets and FluxCD.
- **Resource Synchronization**: Bi-directional secret and configmap replication encrypted via sealed-secrets.

---

##  Repository Structure

```
├── cmd/
│   └── manager/
│       └── main.go           # Operator entry point
├── pkg/
│   ├── apis/
│   │   └── orchestrator/v1alpha1/
│   └── controllers/
│       └── cluster_controller.go
├── config/
│   ├── crd/bases/
│   ├── rbac/
│   └── samples/
├── Makefile
├── go.mod
└── README.md
```

---

##  Development & Testing

```bash
# Install CRDs into current Kubernetes cluster
make install

# Run Controller locally against current kubeconfig context
make run
```

---

##  License
MIT License - Copyright (c) 2025-2026 Rennan Simões.
