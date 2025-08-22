package main

import (
	"github.com/payment-method-tokens/mcp-server/config"
	"github.com/payment-method-tokens/mcp-server/models"
	tools_payment_tokens "github.com/payment-method-tokens/mcp-server/tools/payment_tokens"
	tools_setup_tokens "github.com/payment-method-tokens/mcp-server/tools/setup_tokens"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_payment_tokens.CreateCustomer_payment_tokens_getTool(cfg),
		tools_payment_tokens.CreatePayment_tokens_createTool(cfg),
		tools_payment_tokens.CreatePayment_tokens_deleteTool(cfg),
		tools_payment_tokens.CreatePayment_tokens_getTool(cfg),
		tools_setup_tokens.CreateSetup_tokens_createTool(cfg),
		tools_setup_tokens.CreateSetup_tokens_getTool(cfg),
	}
}
