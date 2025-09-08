package main

import (
	"context"
)

type PrometheusAgnet struct{}

// Ask queries the Prometheus server with a natural language question.
func (m *PrometheusAgnet) Ask(
	ctx context.Context,
	// the prometheus server URL to use
	server string,
	// the natural language question to ask about the prometheus server
	question string,
	// // +optional
	// bearer *dagger.Secret,
) (string, error) {
	prom := dag.Prometheus(server)

	env := dag.Env().
		WithStringInput("question", question, "The question about the prometheus server").
		WithPrometheusInput("prometheus", prom, "The prometheus module to use for inspecting the promethues server")

	return dag.LLM().
		WithEnv(env).
		WithPrompt(`You are an expert on prometheus and TSDB . You have been given
a prometheus module that already has tools and the ability to connect to the database to run PromQL queries and more.
please use the server argument to connect to the server.
The question is: $question

DO NOT STOP UNTIL YOU HAVE ANSWERED THE QUESTION COMPLETELY.`).
		LastReply(ctx)

}
