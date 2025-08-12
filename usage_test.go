// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest_test

import (
	"context"
	"os"
	"testing"

	"github.com/bennorris123/go-sdk-test"
	"github.com/bennorris123/go-sdk-test/internal/testutil"
	"github.com/bennorris123/go-sdk-test/option"
)

func TestUsage(t *testing.T) {
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
	response, err := client.Chat.NewCompletion(context.TODO(), relaxaitest.ChatNewCompletionParams{
		Messages: []relaxaitest.ChatCompletionMessageParam{{
			MultiContent: []relaxaitest.ChatCompletionMessageMultiContentParam{{}},
			Role:         "role",
		}},
		Model: "model",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ID)
}
