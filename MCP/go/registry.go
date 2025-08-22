package main

import (
	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	tools_certificates "github.com/apimanagementclient/mcp-server/tools/certificates"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_certificates.CreateCertificate_listbyserviceTool(cfg),
		tools_certificates.CreateCertificate_deleteTool(cfg),
		tools_certificates.CreateCertificate_getTool(cfg),
		tools_certificates.CreateCertificate_getentitytagTool(cfg),
		tools_certificates.CreateCertificate_createorupdateTool(cfg),
	}
}
