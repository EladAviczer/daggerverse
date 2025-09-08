package main

import (
	"context"
	"dagger/prometheus/internal/dagger"
	"fmt"
	"strings"
)

func New(
	// prometheus server URL
	prometheusurl string,
) *Prometheus {
	return &Prometheus{
		BaseURL: prometheusurl,
	}
}

type Prometheus struct {
	BaseURL string
}

// PromQl runs an *instant* PromQL query via /api/v1/query (JSON output).
func (p *Prometheus) PromQl(
	ctx context.Context,
	// query in PromQL format
	promQuery string,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", p.BaseURL).
		WithEnvVariable("PROM_QUERY", promQuery)

	if bearer != nil {
		c = c.WithSecretVariable("BEARER", bearer)
	}

	return c.WithExec([]string{
		"sh", "-c",
		`H=""; [ -n "$BEARER" ] && H='-H "Authorization: Bearer $BEARER"'; \
		 curl -fsSL -G \
		      -H "Accept: application/json" $H \
		      --data-urlencode "query=$PROM_QUERY" \
		      "${PROMETHEUS_URL%/}/api/v1/query"`,
	}).Stdout(ctx)
}

// FiringAlerts queries the /api/v1/alerts endpoint to list all firing alerts.
func (p *Prometheus) FiringAlerts(
	ctx context.Context,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	fmt.Println(p.BaseURL)
	alertsURL := strings.TrimRight(p.BaseURL, "/") + "/api/v1/alerts"
	fmt.Println(alertsURL)
	c := dag.Container().From("curlimages/curl:8.9.1")

	// add bearer only when provided
	if bearer != nil {
		c = c.WithSecretVariable("BEARER", bearer)
		return c.WithExec([]string{
			"sh", "-c",
			fmt.Sprintf(`curl -fsSL -H "Accept: application/json" -H "Authorization: Bearer $BEARER" %q`, alertsURL),
		}).Stdout(ctx)
	}

	return c.WithExec([]string{
		"curl", "-fsSL",
		"-H", "Accept: application/json",
		alertsURL,
	}).Stdout(ctx)
}

// Targets queries the /api/v1/targets endpoint to list all targets.
func (p *Prometheus) Targets(
	ctx context.Context,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", p.BaseURL)

	if bearer != nil {
		c = c.WithSecretVariable("BEARER", bearer)
	}
	return c.WithExec([]string{
		"sh", "-c",
		`H=""; [ -n "$BEARER" ] && H='-H "Authorization: Bearer $BEARER"'; \
		 curl -fsSL -H "Accept: application/json" $H \
		      "${PROMETHEUS_URL%/}/api/v1/targets"`,
	}).Stdout(ctx)
}

// Rules queries the /api/v1/rules endpoint to list all alerting and recording rules.
func (p *Prometheus) Rules(
	ctx context.Context,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", p.BaseURL)

	if bearer != nil {
		c = c.WithSecretVariable("BEARER", bearer)
	}
	return c.WithExec([]string{
		"sh", "-c",
		`H=""; [ -n "$BEARER" ] && H='-H "Authorization: Bearer $BEARER"'; \
		 curl -fsSL -H "Accept: application/json" $H \
		      "${PROMETHEUS_URL%/}/api/v1/rules"`,
	}).Stdout(ctx)
}
