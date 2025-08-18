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

// EmbeddingService contains methods and other services that help with interacting
// with the relaxai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r EmbeddingService) {
	r = EmbeddingService{}
	r.Options = opts
	return
}

// Creates an embedding vector representing the input text.
func (r *EmbeddingService) NewEmbedding(ctx context.Context, body EmbeddingNewEmbeddingParams, opts ...option.RequestOption) (res *EmbeddingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// The properties Input, Model are required.
type EmbeddingRequestParam struct {
	Input          any               `json:"input,omitzero,required"`
	Model          string            `json:"model,required"`
	Dimensions     param.Opt[int64]  `json:"dimensions,omitzero"`
	EncodingFormat param.Opt[string] `json:"encoding_format,omitzero"`
	User           param.Opt[string] `json:"user,omitzero"`
	paramObj
}

func (r EmbeddingRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow EmbeddingRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmbeddingRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingResponse struct {
	Data       []EmbeddingResponseData `json:"data,required"`
	HTTPHeader map[string][]string     `json:"httpHeader,required"`
	Model      string                  `json:"model,required"`
	Object     string                  `json:"object,required"`
	Usage      shared.OpenAIUsage      `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HTTPHeader  respjson.Field
		Model       respjson.Field
		Object      respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingResponse) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingResponseData struct {
	Embedding []float64 `json:"embedding,required"`
	Index     int64     `json:"index,required"`
	Object    string    `json:"object,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Embedding   respjson.Field
		Index       respjson.Field
		Object      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmbeddingResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmbeddingResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmbeddingNewEmbeddingParams struct {
	EmbeddingRequest EmbeddingRequestParam
	paramObj
}

func (r EmbeddingNewEmbeddingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EmbeddingRequest)
}
func (r *EmbeddingNewEmbeddingParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.EmbeddingRequest)
}
