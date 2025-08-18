// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/bennorris123/go-sdk-test/internal/apijson"
	shimjson "github.com/bennorris123/go-sdk-test/internal/encoding/json"
	"github.com/bennorris123/go-sdk-test/internal/requestconfig"
	"github.com/bennorris123/go-sdk-test/option"
	"github.com/bennorris123/go-sdk-test/packages/param"
	"github.com/bennorris123/go-sdk-test/packages/respjson"
	"github.com/bennorris123/go-sdk-test/shared"
)

// ChatService contains methods and other services that help with interacting with
// the relaxai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewChatService] method instead.
type ChatService struct {
	Options []option.RequestOption
}

// NewChatService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewChatService(opts ...option.RequestOption) (r ChatService) {
	r = ChatService{}
	r.Options = opts
	return
}

// Creates a chat completion for the given model
func (r *ChatService) NewCompletion(ctx context.Context, body ChatNewCompletionParams, opts ...option.RequestOption) (res *ChatCompletionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/chat/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ChatCompletionMessage struct {
	Content          string                              `json:"content,required"`
	Role             string                              `json:"role,required"`
	FunctionCall     FunctionCall                        `json:"function_call"`
	MultiContent     []ChatCompletionMessageMultiContent `json:"MultiContent"`
	Name             string                              `json:"name"`
	ReasoningContent string                              `json:"reasoning_content"`
	Refusal          string                              `json:"refusal"`
	ToolCallID       string                              `json:"tool_call_id"`
	ToolCalls        []ChatCompletionMessageToolCall     `json:"tool_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content          respjson.Field
		Role             respjson.Field
		FunctionCall     respjson.Field
		MultiContent     respjson.Field
		Name             respjson.Field
		ReasoningContent respjson.Field
		Refusal          respjson.Field
		ToolCallID       respjson.Field
		ToolCalls        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessage) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ChatCompletionMessage to a ChatCompletionMessageParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ChatCompletionMessageParam.Overrides()
func (r ChatCompletionMessage) ToParam() ChatCompletionMessageParam {
	return param.Override[ChatCompletionMessageParam](json.RawMessage(r.RawJSON()))
}

type ChatCompletionMessageMultiContent struct {
	ImageURL ChatCompletionMessageMultiContentImageURL `json:"image_url"`
	Text     string                                    `json:"text"`
	Type     string                                    `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ImageURL    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageMultiContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageMultiContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageMultiContentImageURL struct {
	Detail string `json:"detail"`
	URL    string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detail      respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageMultiContentImageURL) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageMultiContentImageURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageToolCall struct {
	Function FunctionCall `json:"function,required"`
	Type     string       `json:"type,required"`
	ID       string       `json:"id"`
	Index    int64        `json:"index"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Function    respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		Index       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionMessageToolCall) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionMessageToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Role are required.
type ChatCompletionMessageParam struct {
	Content          string                                   `json:"content,required"`
	Role             string                                   `json:"role,required"`
	Name             param.Opt[string]                        `json:"name,omitzero"`
	ReasoningContent param.Opt[string]                        `json:"reasoning_content,omitzero"`
	Refusal          param.Opt[string]                        `json:"refusal,omitzero"`
	ToolCallID       param.Opt[string]                        `json:"tool_call_id,omitzero"`
	FunctionCall     FunctionCallParam                        `json:"function_call,omitzero"`
	MultiContent     []ChatCompletionMessageMultiContentParam `json:"MultiContent,omitzero"`
	ToolCalls        []ChatCompletionMessageToolCallParam     `json:"tool_calls,omitzero"`
	paramObj
}

func (r ChatCompletionMessageParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageMultiContentParam struct {
	Text     param.Opt[string]                              `json:"text,omitzero"`
	Type     param.Opt[string]                              `json:"type,omitzero"`
	ImageURL ChatCompletionMessageMultiContentImageURLParam `json:"image_url,omitzero"`
	paramObj
}

func (r ChatCompletionMessageMultiContentParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageMultiContentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageMultiContentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionMessageMultiContentImageURLParam struct {
	Detail param.Opt[string] `json:"detail,omitzero"`
	URL    param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r ChatCompletionMessageMultiContentImageURLParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageMultiContentImageURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageMultiContentImageURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Function, Type are required.
type ChatCompletionMessageToolCallParam struct {
	Function FunctionCallParam `json:"function,omitzero,required"`
	Type     string            `json:"type,required"`
	ID       param.Opt[string] `json:"id,omitzero"`
	Index    param.Opt[int64]  `json:"index,omitzero"`
	paramObj
}

func (r ChatCompletionMessageToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionMessageToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionMessageToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Messages, Model are required.
type ChatCompletionRequestParam struct {
	Messages            []ChatCompletionMessageParam             `json:"messages,omitzero,required"`
	Model               string                                   `json:"model,required"`
	FrequencyPenalty    param.Opt[float64]                       `json:"frequency_penalty,omitzero"`
	Logprobs            param.Opt[bool]                          `json:"logprobs,omitzero"`
	MaxCompletionTokens param.Opt[int64]                         `json:"max_completion_tokens,omitzero"`
	MaxTokens           param.Opt[int64]                         `json:"max_tokens,omitzero"`
	N                   param.Opt[int64]                         `json:"n,omitzero"`
	PresencePenalty     param.Opt[float64]                       `json:"presence_penalty,omitzero"`
	ReasoningEffort     param.Opt[string]                        `json:"reasoning_effort,omitzero"`
	Seed                param.Opt[int64]                         `json:"seed,omitzero"`
	Store               param.Opt[bool]                          `json:"store,omitzero"`
	Stream              param.Opt[bool]                          `json:"stream,omitzero"`
	Temperature         param.Opt[float64]                       `json:"temperature,omitzero"`
	TopLogprobs         param.Opt[int64]                         `json:"top_logprobs,omitzero"`
	TopP                param.Opt[float64]                       `json:"top_p,omitzero"`
	User                param.Opt[string]                        `json:"user,omitzero"`
	ChatTemplateKwargs  any                                      `json:"chat_template_kwargs,omitzero"`
	FunctionCall        any                                      `json:"function_call,omitzero"`
	Functions           []FunctionDefinitionParam                `json:"functions,omitzero"`
	LogitBias           map[string]int64                         `json:"logit_bias,omitzero"`
	Metadata            map[string]string                        `json:"metadata,omitzero"`
	ParallelToolCalls   any                                      `json:"parallel_tool_calls,omitzero"`
	Prediction          ChatCompletionRequestPredictionParam     `json:"prediction,omitzero"`
	ResponseFormat      ChatCompletionRequestResponseFormatParam `json:"response_format,omitzero"`
	Stop                []string                                 `json:"stop,omitzero"`
	StreamOptions       StreamOptionsParam                       `json:"stream_options,omitzero"`
	ToolChoice          any                                      `json:"tool_choice,omitzero"`
	Tools               []ChatCompletionRequestToolParam         `json:"tools,omitzero"`
	paramObj
}

func (r ChatCompletionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Content, Type are required.
type ChatCompletionRequestPredictionParam struct {
	Content string `json:"content,required"`
	Type    string `json:"type,required"`
	paramObj
}

func (r ChatCompletionRequestPredictionParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestPredictionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestPredictionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionRequestResponseFormatParam struct {
	Type       param.Opt[string]                                  `json:"type,omitzero"`
	JsonSchema ChatCompletionRequestResponseFormatJsonSchemaParam `json:"json_schema,omitzero"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Strict are required.
type ChatCompletionRequestResponseFormatJsonSchemaParam struct {
	Name        string            `json:"name,required"`
	Strict      bool              `json:"strict,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ChatCompletionRequestResponseFormatJsonSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestResponseFormatJsonSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestResponseFormatJsonSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Type is required.
type ChatCompletionRequestToolParam struct {
	Type     string                  `json:"type,required"`
	Function FunctionDefinitionParam `json:"function,omitzero"`
	paramObj
}

func (r ChatCompletionRequestToolParam) MarshalJSON() (data []byte, err error) {
	type shadow ChatCompletionRequestToolParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ChatCompletionRequestToolParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponse struct {
	ID                  string                                     `json:"id,required"`
	Choices             []ChatCompletionResponseChoice             `json:"choices,required"`
	Created             int64                                      `json:"created,required"`
	HTTPHeader          map[string][]string                        `json:"httpHeader,required"`
	Model               string                                     `json:"model,required"`
	Object              string                                     `json:"object,required"`
	SystemFingerprint   string                                     `json:"system_fingerprint,required"`
	Usage               shared.OpenAIUsage                         `json:"usage,required"`
	PromptFilterResults []ChatCompletionResponsePromptFilterResult `json:"prompt_filter_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Choices             respjson.Field
		Created             respjson.Field
		HTTPHeader          respjson.Field
		Model               respjson.Field
		Object              respjson.Field
		SystemFingerprint   respjson.Field
		Usage               respjson.Field
		PromptFilterResults respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponseChoice struct {
	FinishReason         string                               `json:"finish_reason,required"`
	Index                int64                                `json:"index,required"`
	Message              ChatCompletionMessage                `json:"message,required"`
	ContentFilterResults ContentFilterResults                 `json:"content_filter_results"`
	Logprobs             ChatCompletionResponseChoiceLogprobs `json:"logprobs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason         respjson.Field
		Index                respjson.Field
		Message              respjson.Field
		ContentFilterResults respjson.Field
		Logprobs             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponseChoiceLogprobs struct {
	Content []ChatCompletionResponseChoiceLogprobsContent `json:"content,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponseChoiceLogprobsContent struct {
	Token       string                                                  `json:"token,required"`
	Logprob     float64                                                 `json:"logprob,required"`
	TopLogprobs []ChatCompletionResponseChoiceLogprobsContentTopLogprob `json:"top_logprobs,required"`
	Bytes       string                                                  `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		TopLogprobs respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseChoiceLogprobsContent) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseChoiceLogprobsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponseChoiceLogprobsContentTopLogprob struct {
	Token   string  `json:"token,required"`
	Logprob float64 `json:"logprob,required"`
	Bytes   string  `json:"bytes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Logprob     respjson.Field
		Bytes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponseChoiceLogprobsContentTopLogprob) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponseChoiceLogprobsContentTopLogprob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionResponsePromptFilterResult struct {
	Index                int64                `json:"index,required"`
	ContentFilterResults ContentFilterResults `json:"content_filter_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Index                respjson.Field
		ContentFilterResults respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponsePromptFilterResult) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponsePromptFilterResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResults struct {
	Hate      ContentFilterResultsHate      `json:"hate"`
	Jailbreak ContentFilterResultsJailbreak `json:"jailbreak"`
	Profanity ContentFilterResultsProfanity `json:"profanity"`
	SelfHarm  ContentFilterResultsSelfHarm  `json:"self_harm"`
	Sexual    ContentFilterResultsSexual    `json:"sexual"`
	Violence  ContentFilterResultsViolence  `json:"violence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Hate        respjson.Field
		Jailbreak   respjson.Field
		Profanity   respjson.Field
		SelfHarm    respjson.Field
		Sexual      respjson.Field
		Violence    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResults) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResults) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsHate struct {
	Filtered bool   `json:"filtered,required"`
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filtered    respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsHate) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsHate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsJailbreak struct {
	Detected bool `json:"detected,required"`
	Filtered bool `json:"filtered,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detected    respjson.Field
		Filtered    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsJailbreak) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsJailbreak) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsProfanity struct {
	Detected bool `json:"detected,required"`
	Filtered bool `json:"filtered,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Detected    respjson.Field
		Filtered    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsProfanity) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsProfanity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsSelfHarm struct {
	Filtered bool   `json:"filtered,required"`
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filtered    respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsSelfHarm) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsSelfHarm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsSexual struct {
	Filtered bool   `json:"filtered,required"`
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filtered    respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsSexual) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsSexual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentFilterResultsViolence struct {
	Filtered bool   `json:"filtered,required"`
	Severity string `json:"severity"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filtered    respjson.Field
		Severity    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentFilterResultsViolence) RawJSON() string { return r.JSON.raw }
func (r *ContentFilterResultsViolence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionCall struct {
	Arguments string `json:"arguments"`
	Name      string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arguments   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionCall) RawJSON() string { return r.JSON.raw }
func (r *FunctionCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionCall to a FunctionCallParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionCallParam.Overrides()
func (r FunctionCall) ToParam() FunctionCallParam {
	return param.Override[FunctionCallParam](json.RawMessage(r.RawJSON()))
}

type FunctionCallParam struct {
	Arguments param.Opt[string] `json:"arguments,omitzero"`
	Name      param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r FunctionCallParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Name, Parameters are required.
type FunctionDefinitionParam struct {
	Name        string            `json:"name,required"`
	Parameters  any               `json:"parameters,omitzero,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Strict      param.Opt[bool]   `json:"strict,omitzero"`
	paramObj
}

func (r FunctionDefinitionParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionDefinitionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionDefinitionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StreamOptionsParam struct {
	IncludeUsage param.Opt[bool] `json:"include_usage,omitzero"`
	paramObj
}

func (r StreamOptionsParam) MarshalJSON() (data []byte, err error) {
	type shadow StreamOptionsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StreamOptionsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatNewCompletionParams struct {
	ChatCompletionRequest ChatCompletionRequestParam
	paramObj
}

func (r ChatNewCompletionParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ChatCompletionRequest)
}
func (r *ChatNewCompletionParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ChatCompletionRequest)
}
