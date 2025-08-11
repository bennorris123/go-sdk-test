// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"context"
	"net/http"

	"github.com/relax-ai/go-sdk/internal/apijson"
	"github.com/relax-ai/go-sdk/internal/requestconfig"
	"github.com/relax-ai/go-sdk/option"
	"github.com/relax-ai/go-sdk/packages/param"
	"github.com/relax-ai/go-sdk/packages/respjson"
)

// ToolService contains methods and other services that help with interacting with
// the relaxai-test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewToolService] method instead.
type ToolService struct {
	Options []option.RequestOption
}

// NewToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewToolService(opts ...option.RequestOption) (r ToolService) {
	r = ToolService{}
	r.Options = opts
	return
}

// Scrap the given URL and return the text content.
func (r *ToolService) Scrap(ctx context.Context, body ToolScrapParams, opts ...option.RequestOption) (res *[]ToolScrapResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/tools/scrap"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ToolScrapResponse struct {
	// Error message if scraping failed
	Error     string `json:"error,required"`
	Markdown  string `json:"markdown,required"`
	Summarize string `json:"summarize,required"`
	Title     string `json:"title,required"`
	URL       string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Markdown    respjson.Field
		Summarize   respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolScrapResponse) RawJSON() string { return r.JSON.raw }
func (r *ToolScrapResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolScrapParams struct {
	Query string   `json:"query,required"`
	URLs  []string `json:"urls,omitzero,required"`
	paramObj
}

func (r ToolScrapParams) MarshalJSON() (data []byte, err error) {
	type shadow ToolScrapParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolScrapParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
