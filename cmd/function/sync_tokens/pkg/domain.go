package pkg

type CsvRow struct {
	CID string
}

type IPFSGatewayResponse struct {
	Image       string `json:"image"`
	Name        string `json:"name"`
	Description string `json:"description"`
	YearCreated int    `json:"yearCreated"`
	CreatedBy   string `json:"createdBy"`
	Artist      string `json:"artist"`
	Edition     int    `json:"edition"`
	Media       struct {
		URI        string `json:"uri"`
		Dimensions string `json:"dimensions"`
		Size       string `json:"size"`
		MimeType   string `json:"mimeType"`
	} `json:"media"`
}
