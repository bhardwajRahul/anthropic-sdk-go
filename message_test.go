// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package anthropic_test

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/internal/testutil"
	"github.com/anthropics/anthropic-sdk-go/option"
)

func TestMessageNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anthropic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-anthropic-api-key"),
	)
	_, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: anthropic.F(int64(1024)),
		Messages: anthropic.F([]anthropic.MessageParam{{
			Content: anthropic.F([]anthropic.ContentBlockParamUnion{anthropic.TextBlockParam{Text: anthropic.F("What is a quaternion?"), Type: anthropic.F(anthropic.TextBlockParamTypeText), CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral)}), Citations: anthropic.F([]anthropic.TextCitationParamUnion{anthropic.CitationCharLocationParam{CitedText: anthropic.F("cited_text"), DocumentIndex: anthropic.F(int64(0)), DocumentTitle: anthropic.F("x"), EndCharIndex: anthropic.F(int64(0)), StartCharIndex: anthropic.F(int64(0)), Type: anthropic.F(anthropic.CitationCharLocationParamTypeCharLocation)}})}}),
			Role:    anthropic.F(anthropic.MessageParamRoleUser),
		}}),
		Model: anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
		Metadata: anthropic.F(anthropic.MetadataParam{
			UserID: anthropic.F("13803d75-b4b5-4c3e-b2a2-6f21399b021b"),
		}),
		StopSequences: anthropic.F([]string{"string"}),
		System:        anthropic.F([]anthropic.TextBlockParam{{Text: anthropic.F("x"), Type: anthropic.F(anthropic.TextBlockParamTypeText), CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral)}), Citations: anthropic.F([]anthropic.TextCitationParamUnion{anthropic.CitationCharLocationParam{CitedText: anthropic.F("cited_text"), DocumentIndex: anthropic.F(int64(0)), DocumentTitle: anthropic.F("x"), EndCharIndex: anthropic.F(int64(0)), StartCharIndex: anthropic.F(int64(0)), Type: anthropic.F(anthropic.CitationCharLocationParamTypeCharLocation)}})}}),
		Temperature:   anthropic.F(1.000000),
		Thinking: anthropic.F[anthropic.ThinkingConfigParamUnion](anthropic.ThinkingConfigEnabledParam{
			BudgetTokens: anthropic.F(int64(1024)),
			Type:         anthropic.F(anthropic.ThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: anthropic.F[anthropic.ToolChoiceUnionParam](anthropic.ToolChoiceAutoParam{
			Type:                   anthropic.F(anthropic.ToolChoiceAutoTypeAuto),
			DisableParallelToolUse: anthropic.F(true),
		}),
		Tools: anthropic.F([]anthropic.ToolUnionUnionParam{anthropic.ToolParam{
			Description: anthropic.F("Get the current weather in a given location"),
			Name:        anthropic.F("x"),
			InputSchema: anthropic.F[interface{}](map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"location": map[string]interface{}{
						"description": "The city and state, e.g. San Francisco, CA",
						"type":        "string",
					},
					"unit": map[string]interface{}{
						"description": "Unit for the output - one of (celsius, fahrenheit)",
						"type":        "string",
					},
				},
			}),
			CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{
				Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral),
			}),
		}}),
		TopK: anthropic.F(int64(5)),
		TopP: anthropic.F(0.700000),
	})
	if err != nil {
		var apierr *anthropic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageCountTokensWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anthropic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-anthropic-api-key"),
	)
	_, err := client.Messages.CountTokens(context.TODO(), anthropic.MessageCountTokensParams{
		Messages: anthropic.F([]anthropic.MessageParam{{
			Content: anthropic.F([]anthropic.ContentBlockParamUnion{anthropic.TextBlockParam{Text: anthropic.F("What is a quaternion?"), Type: anthropic.F(anthropic.TextBlockParamTypeText), CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral)}), Citations: anthropic.F([]anthropic.TextCitationParamUnion{anthropic.CitationCharLocationParam{CitedText: anthropic.F("cited_text"), DocumentIndex: anthropic.F(int64(0)), DocumentTitle: anthropic.F("x"), EndCharIndex: anthropic.F(int64(0)), StartCharIndex: anthropic.F(int64(0)), Type: anthropic.F(anthropic.CitationCharLocationParamTypeCharLocation)}})}}),
			Role:    anthropic.F(anthropic.MessageParamRoleUser),
		}}),
		Model: anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
		System: anthropic.F[anthropic.MessageCountTokensParamsSystemUnion](anthropic.MessageCountTokensParamsSystemArray([]anthropic.TextBlockParam{{
			Text: anthropic.F("Today's date is 2024-06-01."),
			Type: anthropic.F(anthropic.TextBlockParamTypeText),
			CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{
				Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral),
			}),
			Citations: anthropic.F([]anthropic.TextCitationParamUnion{anthropic.CitationCharLocationParam{
				CitedText:      anthropic.F("cited_text"),
				DocumentIndex:  anthropic.F(int64(0)),
				DocumentTitle:  anthropic.F("x"),
				EndCharIndex:   anthropic.F(int64(0)),
				StartCharIndex: anthropic.F(int64(0)),
				Type:           anthropic.F(anthropic.CitationCharLocationParamTypeCharLocation),
			}}),
		}})),
		Thinking: anthropic.F[anthropic.ThinkingConfigParamUnion](anthropic.ThinkingConfigEnabledParam{
			BudgetTokens: anthropic.F(int64(1024)),
			Type:         anthropic.F(anthropic.ThinkingConfigEnabledTypeEnabled),
		}),
		ToolChoice: anthropic.F[anthropic.ToolChoiceUnionParam](anthropic.ToolChoiceAutoParam{
			Type:                   anthropic.F(anthropic.ToolChoiceAutoTypeAuto),
			DisableParallelToolUse: anthropic.F(true),
		}),
		Tools: anthropic.F([]anthropic.MessageCountTokensToolUnionParam{anthropic.ToolParam{
			InputSchema: anthropic.F[interface{}](map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"location": map[string]interface{}{
						"description": "The city and state, e.g. San Francisco, CA",
						"type":        "string",
					},
					"unit": map[string]interface{}{
						"description": "Unit for the output - one of (celsius, fahrenheit)",
						"type":        "string",
					},
				},
			}),
			Name: anthropic.F("name"),
			CacheControl: anthropic.F(anthropic.CacheControlEphemeralParam{
				Type: anthropic.F(anthropic.CacheControlEphemeralTypeEphemeral),
			}),
			Description: anthropic.F("Get the current weather in a given location"),
		}}),
	})
	if err != nil {
		var apierr *anthropic.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMessageNewWithLargeMaxTokens(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := anthropic.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("my-anthropic-api-key"),
	)
	_, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: anthropic.F(int64(23000)),
		Messages:  anthropic.F([]anthropic.MessageParam{{}}),
		Model:     anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
	})
	if err == nil {
		t.Fatalf("err should be set")
	}
	if !strings.Contains(err.Error(), "Streaming is strongly recommended") {
		t.Log(err.Error())
		t.Fatalf("err should be our streaming warning")
	}

	// with custom request timeout
	_, err = client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: anthropic.F(int64(23000)),
		Messages:  anthropic.F([]anthropic.MessageParam{{}}),
		Model:     anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
	}, option.WithRequestTimeout(9999999999))
	if err != nil {
		t.Fatalf("err should be nil")
	}

	// with custom context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	_, err = client.Messages.New(ctx, anthropic.MessageNewParams{
		MaxTokens: anthropic.F(int64(23000)),
		Messages:  anthropic.F([]anthropic.MessageParam{{}}),
		Model:     anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
	})
	if err != nil {
		t.Fatalf("err should be nil")
	}

	// streaming does not error
	stream := client.Messages.NewStreaming(context.TODO(), anthropic.MessageNewParams{
		MaxTokens: anthropic.F(int64(23000)),
		Messages:  anthropic.F([]anthropic.MessageParam{{}}),
		Model:     anthropic.F(anthropic.ModelClaude3_7SonnetLatest),
	})
	if stream.Err() != nil {
		t.Log(stream.Err())
		t.Fatalf("err should be nil")
	}
}
