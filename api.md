# Shared Response Types

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared#OpenAICompletionTokensDetails">OpenAICompletionTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared#OpenAIPromptTokensDetails">OpenAIPromptTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go/shared#OpenAIUsage">OpenAIUsage</a>

# relaxaitest

Methods:

- <code title="get /v1/health">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#RelaxaitestService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionMessageParam">ChatCompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionRequestParam">ChatCompletionRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionCallParam">FunctionCallParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionDefinitionParam">FunctionDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#StreamOptionsParam">StreamOptionsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionMessage">ChatCompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ContentFilterResults">ContentFilterResults</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#FunctionCall">FunctionCall</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatNewCompletionParams">ChatNewCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ChatCompletionResponse">ChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingRequestParam">EmbeddingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingResponse">EmbeddingResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingService.NewEmbedding">NewEmbedding</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingNewEmbeddingParams">EmbeddingNewEmbeddingParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#EmbeddingResponse">EmbeddingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelList">ModelList</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelService.ListModels">ListModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelList">ModelList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#ModelService.GetModel">GetModel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/relaxai-test-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
