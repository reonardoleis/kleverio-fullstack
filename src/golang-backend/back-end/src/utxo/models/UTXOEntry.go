package utxo_models

type UTXOEntry struct {
	Txid          string `json:"txid"`
	Vout          int    `json:"vout"`
	Value         string `json:"value"`
	Height        int    `json:"height"`
	Confirmations int    `json:"confirmations"`
}
