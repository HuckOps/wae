# WAE — Web Application Engine

A Kubernetes-native Web Application Engine for building, deploying, and managing cloud-native microservices with integrated CI/CD and service state management.

---

<div align="center">

[![CI](https://github.com/HuckOps/wae/actions/workflows/go.yml/badge.svg)](https://github.com/HuckOps/wae/actions/workflows/go.yml)

</div>

---

## 🚀 Overview

**WAE (Web Application Engine)** is a modern deployment platform built on Kubernetes, designed to unify the entire application lifecycle:

- Continuous Integration (CI)
- Automated Deployment (CD)
- Service State Management
- Multi-environment delivery
- Extensible execution backend (custom runner/controller)

It provides a single control plane for managing web applications in cloud-native infrastructure.

---

## ✨ Key Features

- **CI/CD Pipeline**

  - Git-triggered build & deploy workflows
  - Dockerfile-based image building
  - Pluggable runner architecture

- **Kubernetes Native Deployment**

  - Automated deployment to Kubernetes clusters
  - Rolling updates & rollback support
  - Namespace / environment isolation

- **Service State Management**

  - Real-time deployment status tracking
  - Version history & audit logs
  - Health monitoring & lifecycle visibility

- **Extensible Architecture**
  - Custom CI runners
  - Custom deployment strategies
  - Event-driven orchestration

---

## 🧱 Architecture

```text
Git Push
   ↓
CI Engine
   ↓
Build Runner ─────→ Container Registry
   ↓
Deployment Controller
   ↓
Kubernetes Cluster
   ↓
Service Runtime
   ↓
State Management DB
```
