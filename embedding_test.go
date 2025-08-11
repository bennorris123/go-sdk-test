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

func TestEmbeddingNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Embeddings.New(context.TODO(), relaxaitest.EmbeddingNewParams{
		Input:          map[string]interface{}{},
		Model:          "model",
		Dimensions:     relaxaitest.Int(0),
		EncodingFormat: relaxaitest.String("encoding_format"),
		User:           relaxaitest.String("user"),
	})
	if err != nil {
		var apierr *relaxaitest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
