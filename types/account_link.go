package types

type AccountLinkRequest struct {
	Customer    CustomerMininalData `json:"customer"`
	Meta        AccountLinkMeta     `json:"meta"`
	Scope       string              `json:"scope"`
	RedirectURL string              `json:"redirect_url"`
}

type AccountLinkMeta struct {
	Ref string `json:"ref"`
}

type AccountLinkResponse struct {
	Status    string          `json:"status"`
	Message   string          `json:"message"`
	Timestamp string          `json:"timestamp"`
	Data      AccountLinkData `json:"data"`
}

type AccountLinkData struct {
	MonoURL     string          `json:"mono_url"`
	Customer    string          `json:"customer"`
	Meta        AccountLinkMeta `json:"meta"`
	Scope       string          `json:"scope"`
	RedirectURL string          `json:"redirect_url"`
	CreatedAt   string          `json:"created_at"`
}

type ReauthLinkRequest struct {
	Account     string          `json:"account"`
	Meta        AccountLinkMeta `json:"meta"`
	Scope       string          `json:"scope"`
	RedirectURL string          `json:"redirect_url"`
}

type AccountExchangeTokenRequest struct {
	Code string `json:"code"`
}

type AccountExchangeTokenResponse struct {
	Status    string                   `json:"status"`
	Message   string                   `json:"message"`
	Timestamp string                   `json:"timestamp"`
	Data      AccountExchangeTokenData `json:"data"`
}

type AccountExchangeTokenData struct {
	ID string `json:"id"`
}
