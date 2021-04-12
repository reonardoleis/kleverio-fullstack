package utxo_models

type UTXOBalance struct {
	Confirmed   int `json:"confirmed"`
	Unconfirmed int `json:"unconfirmed"`
}
