package types

type WebhookPayload struct {
	Event string             `json:"event"`
	Data  WebhookPayloadData `json:"data"`
}

type WebhookPayloadData struct {
	Meta    WebhookPayloadMetaData `json:"meta"`
	Account UserAccountData        `json:"account_data"`
}

type WebhookPayloadAccount struct {
	ID            string      `json:"_id"`
	Institution   Institution `json:"institution"`
	AccountNumber string      `json:"accountNumber"`
	Name          string      `json:"name"`
	Type          string      `json:"type"`
	Currency      string      `json:"currency"`
	BVN           string      `json:"bvn"`
	Balance       int         `json:"balance"`
	CreatedAt     string      `json:"created_at"`
	UpdatedAt     string      `json:"updated_at"`
	V__           int         `json:"__v"`
}

type WebhookPayloadMetaData struct {
	DataStatus string `json:"data_status"`
}
