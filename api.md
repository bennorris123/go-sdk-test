# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatCompletionMessageParam">ChatCompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#FunctionCallParam">FunctionCallParam</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#FunctionDefinitionParam">FunctionDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#StreamOptionsParam">StreamOptionsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatCompletionMessage">ChatCompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ContentFilterResults">ContentFilterResults</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#FunctionCall">FunctionCall</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#Usage">Usage</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatNewCompletionResponse">ChatNewCompletionResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatNewCompletionParams">ChatNewCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ChatNewCompletionResponse">ChatNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /v1/completions">client.Completions.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Methods:

- <code title="get /v1/health">client.Health.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models/{model}">client.Models.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ToolScrapResponse">ToolScrapResponse</a>

Methods:

- <code title="post /v1/tools/scrap">client.Tools.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ToolService.Scrap">Scrap</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ToolScrapParams">ToolScrapParams</a>) ([]<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/relax-ai/go-sdk#ToolScrapResponse">ToolScrapResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
