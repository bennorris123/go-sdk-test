// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/bennorris123/go-sdk-test/internal/apijson"
	"github.com/bennorris123/go-sdk-test/packages/param"
	"github.com/bennorris123/go-sdk-test/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type OpenAICompletionTokensDetails struct {
	AcceptedPredictionTokens int64 `json:"accepted_prediction_tokens,required"`
	AudioTokens              int64 `json:"audio_tokens,required"`
	ReasoningTokens          int64 `json:"reasoning_tokens,required"`
	RejectedPredictionTokens int64 `json:"rejected_prediction_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AcceptedPredictionTokens respjson.Field
		AudioTokens              respjson.Field
		ReasoningTokens          respjson.Field
		RejectedPredictionTokens respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAICompletionTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *OpenAICompletionTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIPromptTokensDetails struct {
	AudioTokens  int64 `json:"audio_tokens,required"`
	CachedTokens int64 `json:"cached_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioTokens  respjson.Field
		CachedTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIPromptTokensDetails) RawJSON() string { return r.JSON.raw }
func (r *OpenAIPromptTokensDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OpenAIUsage struct {
	CompletionTokens        int64                         `json:"completion_tokens,required"`
	CompletionTokensDetails OpenAICompletionTokensDetails `json:"completion_tokens_details,required"`
	PromptTokens            int64                         `json:"prompt_tokens,required"`
	PromptTokensDetails     OpenAIPromptTokensDetails     `json:"prompt_tokens_details,required"`
	TotalTokens             int64                         `json:"total_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletionTokens        respjson.Field
		CompletionTokensDetails respjson.Field
		PromptTokens            respjson.Field
		PromptTokensDetails     respjson.Field
		TotalTokens             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OpenAIUsage) RawJSON() string { return r.JSON.raw }
func (r *OpenAIUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
