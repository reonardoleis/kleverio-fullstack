package utxo_models

import (
	"strconv"
)

type UTXOList struct {
	Entries []UTXOEntry `json:"entries"`
}

func EvaluateList(l UTXOList) UTXOBalance {
	outputBalance := UTXOBalance{
		Confirmed:   0,
		Unconfirmed: 0,
	}
	for _, element := range l.Entries {
		intValue, _ := strconv.Atoi(element.Value)
		if element.Confirmations < 2 {
			outputBalance.Unconfirmed += intValue
		} else {
			outputBalance.Confirmed += intValue
		}
	}
	return outputBalance
}
