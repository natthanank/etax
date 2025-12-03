package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Document represents the document structure in the request body.
type Document struct {
	DocumentType string `json:"document_type"`
	Data         string `json:"data"`
}

// EtaxRequestBody represents the overall structure of the request body.
type EtaxRequestBody struct {
	Success       bool     `json:"success"`
	ResponseCode  string   `json:"response_code"`
	TaxID         string   `json:"tax_id"`
	InvoiceNumber string   `json:"invoice_number"`
	State         string   `json:"state"`
	StatusFile    string   `json:"status_file"`
	Document      Document `json:"document"`
}

// sendToNetSuite mocks sending data to NetSuite.
func sendToNetSuite(data EtaxRequestBody) error {
	// In a real application, you would implement the logic to send the data to NetSuite.
	// For this example, we'll just print a message.
	println("Sending data to NetSuite for invoice:", data.InvoiceNumber)
	return nil
}

// etaxCallbackHandler creates a gin.HandlerFunc for processing e-tax callbacks.
func etaxCallbackHandler(successMessage string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody EtaxRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := sendToNetSuite(requestBody); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send data to NetSuite"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": successMessage})
	}
}

func main() {
	r := gin.Default()

	// Endpoint for /etax/invoice_callback
	r.POST("/etax/invoice_callback", etaxCallbackHandler("Invoice callback processed successfully"))

	// Endpoint for /etax/cn_callback
	r.POST("/etax/cn_callback", etaxCallbackHandler("CN callback processed successfully"))

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
