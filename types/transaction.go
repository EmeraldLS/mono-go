package types

type CustomerTransactionsResponse struct {
	Status      string                `json:"status"`
	Message     string                `json:"message"`
	AccountData []AccountTransactions `json:"data"`
}

type AccountTransactions struct {
	Account                string                 `json:"account"`
	AccountName            string                 `json:"account_name"`
	Bank                   string                 `json:"bank"`
	AccountTransactionData AccountTransactionData `json:"account_transaction_data"`
}

type AccountTransactionData struct {
	Errored bool                    `json:"errored"`
	Data    TransactionDataResponse `json:"data"`
	Error   string                  `json:"error"`
}

type TransactionDataResponse struct {
	Message string           `json:"message"`
	Status  string           `json:"status"`
	Data    TransactionsData `json:"data"`
}

type TransactionsData struct {
	Transactions []Transaction  `json:"transactions"`
	Meta         PaginationMeta `json:"meta"`
}

type Transaction struct {
	ID       string `json:"id"`
	Type     string `json:"type" validate:"eq=credit|eq=debit"`
	Amount   int    `json:"amount"`
	Date     string `json:"date"`
	Balance  string `json:"balance"`
	Currency string `json:"currency"`
	Category string `json:"category"`
}

type PaginationMeta struct {
	Total    int    `json:"total"`
	Pages    int    `json:"pages"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
}
