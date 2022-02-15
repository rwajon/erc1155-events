package models

type Transaction struct {
	BlockNumber     string `json:"blockNumber"`
	BlockHash       string `json:"blockHash"`
	Timestamp       string `json:"timestamp"`
	Date            string `json:"date"`
	From            string `json:"from"`
	Gas             string `json:"gas"`
	GasPrice        string `json:"gasPrice"`
	Hash            string `json:"hash"`
	To              string `json:"to"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	SenderBalance   string `json:"senderBalance"`
	ReceiverBalance string `json:"receiverBalance"`
	ContractAddress string `json:"contractAddress"`
}
