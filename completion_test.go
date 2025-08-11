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

func TestCompletionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Completions.New(context.TODO(), relaxaitest.CompletionNewParams{
		Model:            "model",
		BestOf:           relaxaitest.Int(0),
		Echo:             relaxaitest.Bool(true),
		FrequencyPenalty: relaxaitest.Float(0),
		LogitBias: map[string]int64{
			"foo": 0,
		},
		Logprobs:  relaxaitest.Int(0),
		MaxTokens: relaxaitest.Int(0),
		Metadata: map[string]string{
			"foo": "string",
		},
		N:               relaxaitest.Int(0),
		PresencePenalty: relaxaitest.Float(0),
		Prompt:          map[string]interface{}{},
		Seed:            relaxaitest.Int(0),
		Stop:            []string{"string"},
		Store:           relaxaitest.Bool(true),
		Stream:          relaxaitest.Bool(true),
		StreamOptions: relaxaitest.StreamOptionsParam{
			IncludeUsage: relaxaitest.Bool(true),
		},
		Suffix:      relaxaitest.String("suffix"),
		Temperature: relaxaitest.Float(0),
		TopP:        relaxaitest.Float(0),
		User:        relaxaitest.String("user"),
	})
	if err != nil {
		var apierr *relaxaitest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
