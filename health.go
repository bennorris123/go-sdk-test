// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest

import (
	"context"
	"net/http"

	"github.com/relax-ai/go-sdk/internal/requestconfig"
	"github.com/relax-ai/go-sdk/option"
)

// HealthService contains methods and other services that help with interacting
// with the relaxai-test API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Check the health of the service.
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
