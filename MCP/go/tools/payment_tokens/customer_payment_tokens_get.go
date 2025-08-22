package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/payment-method-tokens/mcp-server/config"
	"github.com/payment-method-tokens/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Customer_payment_tokens_getHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["customer_id"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("customer_id=%v", val))
		}
		if val, ok := args["page_size"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page_size=%v", val))
		}
		if val, ok := args["page"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("page=%v", val))
		}
		if val, ok := args["total_required"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("total_required=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/v3/vault/payment-tokens%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.Customervaultpaymenttokensresponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCustomer_payment_tokens_getTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_v3_vault_payment-tokens",
		mcp.WithDescription("List all payment tokens"),
		mcp.WithString("customer_id", mcp.Required(), mcp.Description("A unique identifier representing a specific customer in merchant's/partner's system or records.")),
		mcp.WithNumber("page_size", mcp.Description("A non-negative, non-zero integer indicating the maximum number of results to return at one time.")),
		mcp.WithNumber("page", mcp.Description("A non-negative, non-zero integer representing the page of the results.")),
		mcp.WithBoolean("total_required", mcp.Description("A boolean indicating total number of items (total_items) and pages (total_pages) are expected to be returned in the response.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Customer_payment_tokens_getHandler(cfg),
	}
}
