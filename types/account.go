package types

type LinkedAccountsResponse struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      []UserAccountData `json:"data"`
	Meta      PaginationMeta    `json:"meta"`
}

type UserAccountData struct {
	ID            string       `json:"id"`
	Name          string       `json:"name"`
	AccountNumber string       `json:"account_number"`
	Currency      string       `json:"currency"`
	Balance       int          `json:"balance"`
	AuthMethod    string       `json:"auth_method"`
	Status        string       `json:"status"`
	BVN           string       `json:"bvn"`
	Type          string       `json:"type"`
	Institution   Institution  `json:"institution"`
	Customer      CustomerData `json:"customer"`
}

type Institution struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	BankCode string `json:"bank_code"`
	Type     string `json:"type"`
}

type CustomerData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AccountDataResponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Timestamp string      `json:"timestamp"`
	Data      AccountData `json:"data"`
}

type AccountData struct {
	Account Account     `json:"account"`
	Meta    AccountMeta `json:"meta"`
}

type Account struct {
	ID            string      `json:"id"`
	Name          string      `json:"name"`
	AccountNumber string      `json:"account_number"`
	Currency      string      `json:"currency"`
	Balance       int         `json:"balance"`
	AuthMethod    string      `json:"auth_method"`
	Status        string      `json:"status"`
	BVN           string      `json:"bvn"`
	Type          string      `json:"type"`
	Institution   Institution `json:"institution"`
}

type AccountMeta struct {
	DataStatus string `json:"data_status"`
	AuthMethod string `json:"auth_method"`
}

type IdentityResponse struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Data    IdentityResponseData `json:"data"`
}

type IdentityResponseData struct {
	FullName      string `json:"full_name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	BVN           string `json:"bvn"`
	MaritalStatus string `json:"marital_status"`
	AddressLine1  string `json:"address_line_1"`
	AddressLine2  string `json:"address_line_2"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type AccountBalanceResponse struct {
	Status  string                     `json:"successful"`
	Message string                     `json:"message"`
	Data    AccountBalanceResponseData `json:"data"`
}

type AccountBalanceResponseData struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	AccountNumber string `json:"account_number"`
	Balance       int    `json:"balance"`
	Currency      string `json:"currency"`
}

// TODO: Review this struct
type AccountBalanceErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription error  `json:"error_description"`
}

type UnlinkAccountResponse struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

type CustomerAccountsResponse struct {
	Status    string            `json:"status"`
	Message   string            `json:"message"`
	Timestamp string            `json:"timestamp"`
	Data      []UserAccountData `json:"data"`
	Meta      PaginationMeta    `json:"meta"`
}

// TODO: Review this struct
type CustomerAccountsErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription error  `json:"error_description"`
}

type AccountStatementsResponse struct {
	Status  string                `json:"status"`
	Message string                `json:"message"`
	Data    AccountStatementsData `json:"data"`
}

type AccountStatementsData struct {
	Count           int                `json:"count"`
	RequestedLength int                `json:"requested_length"`
	AvailableLength int                `json:"available_length"`
	Statement       []AccountStatement `json:"statement"`
}

type AccountStatement struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	Balance   int    `json:"balance"`
	Date      string `json:"date"`
	Narration string `json:"narration"`
	Currency  string `json:"currency"`
}

type AccountStatementPDFResponse struct {
	Status  string                  `json:"status"`
	Message string                  `json:"message"`
	Data    AccountStatementPDFData `json:"data"`
}

type AccountStatementPDFData struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Path   string `json:"path"`
}

type AccountTransactionParams struct {
	Start     string `json:"start"`
	End       string `json:"end"`
	Narration string `json:"narration"`
	Type      string `json:"type"`
	Paginate  bool   `json:"paginate"`
	Limits    int32  `json:"limits"`
}

type AccountTransactionResponse struct {
	Status    string         `json:"status"`
	Message   string         `json:"message"`
	Timestamp string         `json:"timestamp"`
	Data      []Transaction  `json:"data"`
	Meta      PaginationMeta `json:"meta"`
}

type AssetsResponse struct {
	Status    string             `json:"status"`
	Message   string             `json:"message"`
	Timestamp string             `json:"timestamp"`
	Data      AssetsResponseData `json:"data"`
}

type AssetsResponseData struct {
	ID       string                 `json:"id"`
	Balances AssetsResponseBalances `json:"balances"`
	Assets   []Asset                `json:"asset"`
}

type AssetsResponseBalances struct {
	USD int `json:"_u_s_d"`
}

type Asset struct {
	Name     string       `json:"name"`
	Type     string       `json:"type"`
	Cost     int          `json:"cost"`
	Return   int          `json:"return"`
	Quantity int          `json:"quantity"`
	Currency string       `json:"currency"`
	Details  AssetDetails `json:"details"`
}

type AssetDetails struct {
	Symbol         string `json:"symbol"`
	Price          int    `json:"price"`
	CurrentBalance int    `json:"current_balance"`
}

type EarningsResponse struct {
	Status    string                 `json:"status"`
	Message   string                 `json:"message"`
	Timestamp string                 `json:"timestamp"`
	Data      []EarningsResponseData `json:"data"`
}

type EarningsResponseData struct {
	ID        string `json:"id"`
	Amount    int    `json:"amount"`
	Narration string `json:"narration"`
	Date      string `json:"date"`
	Asset     Asset  `json:"asset"`
}

type EarningsAsset struct {
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	SalePrice    string `json:"sale_price"`
	QuantitySold string `json:"quantity_sold"`
}

type FlowResponse struct {
	Total            int           `json:"total"`
	TransactionCount int           `json:"transaction_count"`
	History          []FlowHistory `json:"history"`
}

type FlowHistory struct {
	Period           string `json:"period"`
	Amount           int    `json:"amount"`
	TransactionCount int    `json:"transaction_count"`
}

type DataSyncResponse struct {
	Status     string `json:"status"`
	HasNewData bool   `json:"hasNewData"`
	Code       string `json:"code"`
}

type DataSyncErrorResponse struct {
	Status string `json:"status"`
	Code   string `json:"code"`
}

type ReauthCodeResponse struct {
	Token string `json:"token"`
}

type AuthLoginRequest struct {
	PhoneNumber string `json:"phone" validate:"request"`
	Provider    string `json:"provider" validate:"request"`
}

type AuthLoginResponse struct {
	Status    string          `json:"status"`
	SessionID string          `json:"session_id"`
	Result    AuthLoginResult `json:"result"`
}

type AuthLoginResult struct {
	Title string          `json:"title"`
	Form  []AuthLoginForm `json:"form"`
}

type AuthLoginForm struct {
	Type        string `json:"type"`        //elements.input
	Name        string `json:"name"`        //otp
	Hint        string `json:"hint"`        //Enter OTP
	ContentType string `json:"contentType"` //password
	MinLength   string `json:"minLength"`
	MaxLength   string `json:"maxLength"`
}

type VerifyOTPRequest struct {
	OTP string `json:"otp" validate:"required"`
}

type VerifyOTPResponse struct {
	Status string                `json:"status"`
	Data   VerifyOTPResponseData `json:"data"`
}

type VerifyOTPResponseData struct {
	Code string `json:"code"`
}

type AuthTokenExchangeRequest struct {
	Code string `json:"code"`
}

type AuthTokenExchangeResponse struct {
	ID string `json:"id"`
}

type TelcoAccountDetailsResponse struct {
	Meta    TelcoAccountMeta `json:"meta"`
	Account TelcoAccount     `json:"account"`
}

type TelcoAccountMeta struct {
	DataStatus string `json:"data_status"`
	AuthMethod string `json:"auth_method"`
}

type TelcoAccount struct {
	ID            string           `json:"_id"`
	Institution   TelcoInstitution `json:"institution"`
	Name          string           `json:"name"`
	AccountNumber string           `json:"accountNumber"`
	Type          string           `json:"type"`
	Balance       int              `json:"balance"`
	Currency      string           `json:"currency"`
	BVN           string           `json:"bvn"`
}

type TelcoInstitution struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TelcoAccountBalanceResponse struct {
	ID    string `json:"_id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value int    `json:"value"`
	Unit  string `json:"unit"`
}

type TelcoTransactionsResponse struct {
	Paging PaginationMeta         `json:"paging"`
	Data   []TelcoTransactionData `json:"data"`
}

type TelcoTransactionData struct {
	ID        string `json:"_id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	Narration string `json:"narration"`
	Balance   int    `json:"balance"`
	Date      string `json:"date"`
	Currency  string `json:"currency"`
}

type TelcoTransactionsErrorResponse struct {
	Message string `json:"message"`
}

type TelcoIdentityResponse struct {
	FullName      string `json:"fullName"`
	Phone         string `json:"phone"`
	Gender        string `json:"gender"`
	BVN           int    `json:"bvn"`
	DOB           string `json:"dob"`
	MaritalStatus string `json:"maritalStatus"`
	AddressLine1  string `json:"addressLine1"`
	AddressLine2  string `json:"addressLine2"`
}

type TelcoIdentityErrorResponse struct {
	Message string `json:"message"`
}

type BankCoverageParams struct {
	Type       string `json:"type"`
	Country    string `json:"country"`
	Scope      string `json:"scope"`
	AuthMethod string `json:"auth_method"`
}

type BankCoverageResponse struct {
	Status    string             `json:"status"`
	Message   string             `json:"message"`
	Timestamp string             `json:"timestamp"`
	Data      []BankCoverageData `json:"data"`
}

type BankCoverageData struct {
	ID          string                   `json:"id"`
	Institution string                   `json:"institution"`
	Type        string                   `json:"type"`
	NipCode     string                   `json:"nip_code"`
	BankCode    string                   `json:"bank_code"`
	Country     string                   `json:"country"`
	AuthMethods []BankCoverageAuthMethod `json:"auth_methods"`
	Scope       []BankCoverageScope      `json:"scope"`
}

type BankCoverageAuthMethod struct {
	ID         string `json:"_id"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	Identifier string `json:"identifier"`
}

type BankCoverageScope struct {
	Name string `json:"name"`
	Type string `json:"type"`
}
