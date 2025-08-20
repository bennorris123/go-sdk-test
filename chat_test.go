// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package relaxaitest_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/bennorris123/go-sdk-test"
	"github.com/bennorris123/go-sdk-test/internal/testutil"
	"github.com/bennorris123/go-sdk-test/option"
)

func TestChatNewCompletionWithOptionalParams(t *testing.T) {
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
	_, err := client.Chat.NewCompletion(context.TODO(), relaxaitest.ChatNewCompletionParams{
		ChatCompletionRequest: relaxaitest.ChatCompletionRequestParam{
			Messages: []relaxaitest.ChatCompletionMessageParam{{
				Content: "content",
				Role:    "role",
				FunctionCall: relaxaitest.FunctionCallParam{
					Arguments: relaxaitest.String("arguments"),
					Name:      relaxaitest.String("name"),
				},
				MultiContent: []relaxaitest.ChatCompletionMessageMultiContentParam{{
					ImageURL: relaxaitest.ChatCompletionMessageMultiContentImageURLParam{
						Detail: relaxaitest.String("detail"),
						URL:    relaxaitest.String("url"),
					},
					Text: relaxaitest.String("text"),
					Type: relaxaitest.String("type"),
				}},
				Name:             relaxaitest.String("name"),
				ReasoningContent: relaxaitest.String("reasoning_content"),
				Refusal:          relaxaitest.String("refusal"),
				ToolCallID:       relaxaitest.String("tool_call_id"),
				ToolCalls: []relaxaitest.ChatCompletionMessageToolCallParam{{
					Function: relaxaitest.FunctionCallParam{
						Arguments: relaxaitest.String("arguments"),
						Name:      relaxaitest.String("name"),
					},
					Type:  "type",
					ID:    relaxaitest.String("id"),
					Index: relaxaitest.Int(0),
				}},
			}},
			Model:              "model",
			ChatTemplateKwargs: map[string]interface{}{},
			FrequencyPenalty:   relaxaitest.Float(0),
			FunctionCall:       map[string]interface{}{},
			Functions: []relaxaitest.FunctionDefinitionParam{{
				Name:        "name",
				Parameters:  map[string]interface{}{},
				Description: relaxaitest.String("description"),
				Strict:      relaxaitest.Bool(true),
			}},
			LogitBias: map[string]int64{
				"foo": 0,
			},
			Logprobs:            relaxaitest.Bool(true),
			MaxCompletionTokens: relaxaitest.Int(0),
			MaxTokens:           relaxaitest.Int(0),
			Metadata: map[string]string{
				"foo": "string",
			},
			N:                 relaxaitest.Int(0),
			ParallelToolCalls: map[string]interface{}{},
			Prediction: relaxaitest.ChatCompletionRequestPredictionParam{
				Content: "content",
				Type:    "type",
			},
			PresencePenalty: relaxaitest.Float(0),
			ReasoningEffort: relaxaitest.String("reasoning_effort"),
			ResponseFormat: relaxaitest.ChatCompletionRequestResponseFormatParam{
				JsonSchema: relaxaitest.ChatCompletionRequestResponseFormatJsonSchemaParam{
					Name:        "name",
					Strict:      true,
					Description: relaxaitest.String("description"),
				},
				Type: relaxaitest.String("type"),
			},
			Seed:   relaxaitest.Int(0),
			Stop:   []string{"string"},
			Store:  relaxaitest.Bool(true),
			Stream: relaxaitest.Bool(true),
			StreamOptions: relaxaitest.StreamOptionsParam{
				IncludeUsage: relaxaitest.Bool(true),
			},
			Temperature: relaxaitest.Float(0),
			ToolChoice:  map[string]interface{}{},
			Tools: []relaxaitest.ChatCompletionRequestToolParam{{
				Type: "type",
				Function: relaxaitest.FunctionDefinitionParam{
					Name:        "name",
					Parameters:  map[string]interface{}{},
					Description: relaxaitest.String("description"),
					Strict:      relaxaitest.Bool(true),
				},
			}},
			TopLogprobs: relaxaitest.Int(0),
			TopP:        relaxaitest.Float(0),
			User:        relaxaitest.String("user"),
		},
	})
	if err != nil {
		var apierr *relaxaitest.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
