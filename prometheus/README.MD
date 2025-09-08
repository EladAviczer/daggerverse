# Prometheus Dagger Module

This module provides a simple Dagger interface for querying a Prometheus server.  
It wraps the Prometheus HTTP API (`/api/v1/...`) using `curl` in ephemeral containers.

```bash
dagger install github.com/EladAviczer/daggerverse/prometheus
```