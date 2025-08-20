# Shared Response Types

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared">shared</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared#OpenAICompletionTokensDetails">OpenAICompletionTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared">shared</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared#OpenAIPromptTokensDetails">OpenAIPromptTokensDetails</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared">shared</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test/shared#OpenAIUsage">OpenAIUsage</a>

# relaxaitest

Methods:

- <code title="get /v1/health">client.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#RelaxaitestService.Health">Health</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

Params Types:

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatCompletionMessageParam">ChatCompletionMessageParam</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatCompletionRequestParam">ChatCompletionRequestParam</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#FunctionCallParam">FunctionCallParam</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#FunctionDefinitionParam">FunctionDefinitionParam</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#StreamOptionsParam">StreamOptionsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatCompletionMessage">ChatCompletionMessage</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatCompletionResponse">ChatCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ContentFilterResults">ContentFilterResults</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#FunctionCall">FunctionCall</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatNewCompletionParams">ChatNewCompletionParams</a>) (<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ChatCompletionResponse">ChatCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Params Types:

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#EmbeddingRequestParam">EmbeddingRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#EmbeddingResponse">EmbeddingResponse</a>

Methods:

- <code title="post /v1/embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#EmbeddingService.NewEmbedding">NewEmbedding</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#EmbeddingNewEmbeddingParams">EmbeddingNewEmbeddingParams</a>) (<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#EmbeddingResponse">EmbeddingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ModelList">ModelList</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ModelService.ListModels">ListModels</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ModelList">ModelList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/models/{model}">client.Models.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#ModelService.GetModel">GetModel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test">relaxaitest</a>.<a href="https://pkg.go.dev/github.com/bennorris123/go-sdk-test#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
