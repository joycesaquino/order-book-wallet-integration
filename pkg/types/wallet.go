package types

const (
	credit = "CREDIT"
	debit  = "DEBIT"
)

type Orders []Order
type Order struct {
	Value         float64 `json:"value"`
	Quantity      int     `json:"quantity"`
	OperationType string  `json:"operationType"`
	UserId        int     `json:"userId"`
	RequestId     string  `json:"RequestId"`
}

type Body struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
