// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/bennorris123/go-sdk-test/internal/apijson"
	"github.com/bennorris123/go-sdk-test/internal/requestconfig"
	"github.com/bennorris123/go-sdk-test/option"
	"github.com/bennorris123/go-sdk-test/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the relaxai API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// List all the available models
func (r *ModelService) ListModels(ctx context.Context, opts ...option.RequestOption) (res *ModelList, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the details of the given model
func (r *ModelService) GetModel(ctx context.Context, model string, opts ...option.RequestOption) (res *Model, err error) {
	opts = append(r.Options[:], opts...)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("v1/models/%s", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Model struct {
	ID         string              `json:"id,required"`
	Created    int64               `json:"created,required"`
	HTTPHeader map[string][]string `json:"httpHeader,required"`
	Object     string              `json:"object,required"`
	OwnedBy    string              `json:"owned_by,required"`
	Parent     string              `json:"parent,required"`
	Permission []ModelPermission   `json:"permission,required"`
	Root       string              `json:"root,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		HTTPHeader  respjson.Field
		Object      respjson.Field
		OwnedBy     respjson.Field
		Parent      respjson.Field
		Permission  respjson.Field
		Root        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelPermission struct {
	ID                 string `json:"id,required"`
	AllowCreateEngine  bool   `json:"allow_create_engine,required"`
	AllowFineTuning    bool   `json:"allow_fine_tuning,required"`
	AllowLogprobs      bool   `json:"allow_logprobs,required"`
	AllowSampling      bool   `json:"allow_sampling,required"`
	AllowSearchIndices bool   `json:"allow_search_indices,required"`
	AllowView          bool   `json:"allow_view,required"`
	Created            int64  `json:"created,required"`
	Group              any    `json:"group,required"`
	IsBlocking         bool   `json:"is_blocking,required"`
	Object             string `json:"object,required"`
	Organization       string `json:"organization,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AllowCreateEngine  respjson.Field
		AllowFineTuning    respjson.Field
		AllowLogprobs      respjson.Field
		AllowSampling      respjson.Field
		AllowSearchIndices respjson.Field
		AllowView          respjson.Field
		Created            respjson.Field
		Group              respjson.Field
		IsBlocking         respjson.Field
		Object             respjson.Field
		Organization       respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelPermission) RawJSON() string { return r.JSON.raw }
func (r *ModelPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelList struct {
	Data       []Model             `json:"data,required"`
	HTTPHeader map[string][]string `json:"httpHeader,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		HTTPHeader  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelList) RawJSON() string { return r.JSON.raw }
func (r *ModelList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
