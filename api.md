# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionMessageParam">ChatCompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionCallParam">FunctionCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionDefinitionParam">FunctionDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#StreamOptionsParam">StreamOptionsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionMessage">ChatCompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ContentFilterResults">ContentFilterResults</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionCall">FunctionCall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#Usage">Usage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatNewCompletionResponse">ChatNewCompletionResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatNewCompletionParams">ChatNewCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatNewCompletionResponse">ChatNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /v1/completions">client.Completions.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingNewResponse">EmbeddingNewResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingNewResponse">EmbeddingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models/{model}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ToolScrapResponse">ToolScrapResponse</a>

Methods:

- <code title="post /v1/tools/scrap">client.Tools.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ToolService.Scrap">Scrap</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ToolScrapParams">ToolScrapParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ToolScrapResponse">ToolScrapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
