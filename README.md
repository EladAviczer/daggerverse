# Daggerverse Modules

This repo contains my Dagger modules, currently:

* **prometheus** – call Prometheus APIs (PromQL, alerts, targets, rules)
* **prometheus-agent** – ask Prometheus questions in natural language

---

## Install

```bash
dagger install github.com/EladAviczer/daggerverse/prometheus@<version>
dagger install github.com/EladAviczer/daggerverse/prometheus-agent@<version>
```

---

## Usage

PromQL query:

```bash
dagger -m github.com/EladAviczer/daggerverse/prometheus@<version> call \
  --server https://prom.example.com \
  prom-ql --prom-query 'up'
```

Ask in natural language:

```bash
dagger -m github.com/EladAviczer/daggerverse/prometheus-agent@<version> call \
  ask --server https://prom.example.com \
  --question "Which services had errors in last 5m?"
```

---

## Auth

Pass bearer tokens via secrets:

```bash
dagger call --server https://prom.example.com prom-ql \
  --prom-query 'up' --bearer env:BEARER_TOKEN
```

---

## Dev

```bash
git clone https://github.com/EladAviczer/daggerverse
cd prometheus
dagger call --server http://localhost:9090 prom-ql --prom-query 'up'
```
