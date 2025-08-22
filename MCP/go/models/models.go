package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// CertificateCreateOrUpdateParameters represents the CertificateCreateOrUpdateParameters schema from the OpenAPI specification
type CertificateCreateOrUpdateParameters struct {
	Properties CertificateCreateOrUpdateProperties `json:"properties,omitempty"` // Parameters supplied to the CreateOrUpdate certificate operation.
}

// CertificateCreateOrUpdateProperties represents the CertificateCreateOrUpdateProperties schema from the OpenAPI specification
type CertificateCreateOrUpdateProperties struct {
	Data string `json:"data"` // Base 64 encoded certificate using the application/x-pkcs12 representation.
	Password string `json:"password"` // Password for the Certificate
}

// CertificateCollection represents the CertificateCollection schema from the OpenAPI specification
type CertificateCollection struct {
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []CertificateContract `json:"value,omitempty"` // Page values.
}

// CertificateContract represents the CertificateContract schema from the OpenAPI specification
type CertificateContract struct {
	Id string `json:"id,omitempty"` // Resource ID.
	Name string `json:"name,omitempty"` // Resource name.
	TypeField string `json:"type,omitempty"` // Resource type for API Management resource.
}

// CertificateContractProperties represents the CertificateContractProperties schema from the OpenAPI specification
type CertificateContractProperties struct {
	Subject string `json:"subject"` // Subject attribute of the certificate.
	Thumbprint string `json:"thumbprint"` // Thumbprint of the certificate.
	Expirationdate string `json:"expirationDate"` // Expiration date of the certificate. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
}
