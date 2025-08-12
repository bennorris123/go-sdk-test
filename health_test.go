// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/relax-ai/go-sdk"
	"github.com/relax-ai/go-sdk/internal/testutil"
	"github.com/relax-ai/go-sdk/option"
)

func TestHealthCheck(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := relaxaitest.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Health.Check(context.TODO())
	if err != nil {
		var apierr *relaxaitest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
