// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"context"
	"net/http"

	"github.com/bennorris123/go-sdk-test/internal/apijson"
	"github.com/bennorris123/go-sdk-test/internal/requestconfig"
	"github.com/bennorris123/go-sdk-test/option"
	"github.com/bennorris123/go-sdk-test/packages/param"
	"github.com/bennorris123/go-sdk-test/packages/respjson"
)

// CompletionService contains methods and other services that help with interacting
// with the relaxai-test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCompletionService] method instead.
type CompletionService struct {
	Options []option.RequestOption
}

// NewCompletionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCompletionService(opts ...option.RequestOption) (r CompletionService) {
	r = CompletionService{}
	r.Options = opts
	return
}

// Creates a completion for the given model
//
// Deprecated: deprecated
func (r *CompletionService) New(ctx context.Context, body CompletionNewParams, opts ...option.RequestOption) (res *CompletionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CompletionNewResponse struct {
	ID         string                        `json:"id,required"`
	Choices    []CompletionNewResponseChoice `json:"choices,required"`
	Created    int64                         `json:"created,required"`
	HTTPHeader map[string][]string           `json:"httpHeader,required"`
	Model      string                        `json:"model,required"`
	Object     string                        `json:"object,required"`
	Usage      Usage                         `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Choices     respjson.Field
		Created     respjson.Field
		HTTPHeader  respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewResponseChoice struct {
	FinishReason string                              `json:"finish_reason,required"`
	Index        int64                               `json:"index,required"`
	Logprobs     CompletionNewResponseChoiceLogprobs `json:"logprobs,required"`
	Text         string                              `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FinishReason respjson.Field
		Index        respjson.Field
		Logprobs     respjson.Field
		Text         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionNewResponseChoice) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewResponseChoiceLogprobs struct {
	TextOffset    []int64              `json:"text_offset,required"`
	TokenLogprobs []float64            `json:"token_logprobs,required"`
	Tokens        []string             `json:"tokens,required"`
	TopLogprobs   []map[string]float64 `json:"top_logprobs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TextOffset    respjson.Field
		TokenLogprobs respjson.Field
		Tokens        respjson.Field
		TopLogprobs   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CompletionNewResponseChoiceLogprobs) RawJSON() string { return r.JSON.raw }
func (r *CompletionNewResponseChoiceLogprobs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompletionNewParams struct {
	Model            string             `json:"model,required"`
	BestOf           param.Opt[int64]   `json:"best_of,omitzero"`
	Echo             param.Opt[bool]    `json:"echo,omitzero"`
	FrequencyPenalty param.Opt[float64] `json:"frequency_penalty,omitzero"`
	Logprobs         param.Opt[int64]   `json:"logprobs,omitzero"`
	MaxTokens        param.Opt[int64]   `json:"max_tokens,omitzero"`
	N                param.Opt[int64]   `json:"n,omitzero"`
	PresencePenalty  param.Opt[float64] `json:"presence_penalty,omitzero"`
	Seed             param.Opt[int64]   `json:"seed,omitzero"`
	Store            param.Opt[bool]    `json:"store,omitzero"`
	Stream           param.Opt[bool]    `json:"stream,omitzero"`
	Suffix           param.Opt[string]  `json:"suffix,omitzero"`
	Temperature      param.Opt[float64] `json:"temperature,omitzero"`
	TopP             param.Opt[float64] `json:"top_p,omitzero"`
	User             param.Opt[string]  `json:"user,omitzero"`
	LogitBias        map[string]int64   `json:"logit_bias,omitzero"`
	Metadata         map[string]string  `json:"metadata,omitzero"`
	Prompt           any                `json:"prompt,omitzero"`
	Stop             []string           `json:"stop,omitzero"`
	StreamOptions    StreamOptionsParam `json:"stream_options,omitzero"`
	paramObj
}

func (r CompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
