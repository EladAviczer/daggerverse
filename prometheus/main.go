// A generated module for Prometheus functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/prometheus/internal/dagger"
	"fmt"
	"strings"
)

type Prometheus struct{}

// PromQl runs an *instant* PromQL query via /api/v1/query (JSON output).
func (m *Prometheus) PromQl(
	ctx context.Context,
	// +default="http://localhost:9090"
	server string,
	promQuery string,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", server).
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
func (m *Prometheus) FiringAlerts(
	ctx context.Context,
	// +default="http://localhost:9090"
	server string,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	alertsURL := strings.TrimRight(server, "/") + "/api/v1/alerts"

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
func (m *Prometheus) Targets(
	ctx context.Context,
	// +default="http://localhost:9090"
	server string,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", server)

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
func (m *Prometheus) Rules(
	ctx context.Context,
	// +default="http://localhost:9090"
	server string,
	// +optional
	bearer *dagger.Secret,
) (string, error) {
	c := dag.Container().
		From("curlimages/curl:8.9.1").
		WithEnvVariable("PROMETHEUS_URL", server)

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
