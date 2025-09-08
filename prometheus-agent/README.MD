# Prometheus Agent
Ask questions regarding a prometheus instance.

## Overview
- The agent connects to the provided prometheus instance.
- The agent uses prometheus dagger module with predefined dagger functions
- The agent invoke the relevant function/tool if necessary and provides clear answer.

### Works with multiple LLMs:
- OpenAI
- Anthropic
- Gemini
- Amazon Bedrock
- Azure OpenAI
- Docker Model Runner
- Ollama


#### Prerequisites
```
cp .env.example .env
```
update .env file according to this example
https://docs.dagger.io/reference/configuration/llm


