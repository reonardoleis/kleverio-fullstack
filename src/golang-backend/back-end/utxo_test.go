package main

import (
	"testing"

	utxo_models "github.com/reonardoleis/klever.io-backend/src/utxo/models"
)

func TestEmptyEntryList(t *testing.T) {
	correctBalance := utxo_models.UTXOBalance{
		Confirmed:   0,
		Unconfirmed: 0,
	}
	emptyUTXOList := utxo_models.UTXOList{
		Entries: []utxo_models.UTXOEntry{},
	}
	evaluatedBalance := utxo_models.EvaluateList(emptyUTXOList)

	if evaluatedBalance.Confirmed != correctBalance.Confirmed || evaluatedBalance.Unconfirmed != correctBalance.Unconfirmed {
		t.Errorf("List evaluating failed, expected %v, got %v", correctBalance, evaluatedBalance)
	}

}

func TestFirstNonEmptyEntryList(t *testing.T) {

	entry1 := utxo_models.UTXOEntry{
		Txid:          "01e2c94ec7b36241543c019686b8cc8727ca34a7970a9771bb202180381946d4",
		Vout:          0,
		Value:         "2396",
		Height:        52309,
		Confirmations: 1,
	}
	entry2 := utxo_models.UTXOEntry{
		Txid:          "8958472f189e153fb5d7e085fb586cda29e88245495a7d010a71348ad10196ec",
		Vout:          13,
		Value:         "2425",
		Height:        26316,
		Confirmations: 10,
	}
	entry3 := utxo_models.UTXOEntry{
		Txid:          "fc3d5031e4589cc99ad69ee908854a99c689342c4c9923f9a28d5d81142fa4f0",
		Vout:          1,
		Value:         "239623",
		Height:        23214,
		Confirmations: 19874,
	}

	correctBalance := utxo_models.UTXOBalance{
		Confirmed:   242048,
		Unconfirmed: 2396,
	}

	nonEmptyUTXOList := utxo_models.UTXOList{
		Entries: []utxo_models.UTXOEntry{entry1, entry2, entry3},
	}
	evaluatedBalance := utxo_models.EvaluateList(nonEmptyUTXOList)

	if evaluatedBalance.Confirmed != correctBalance.Confirmed || evaluatedBalance.Unconfirmed != correctBalance.Unconfirmed {
		t.Errorf("List evaluating failed, expected %v, got %v", correctBalance, evaluatedBalance)
	}

}

func TestSecondNonEmptyEntryList(t *testing.T) {

	entry1 := utxo_models.UTXOEntry{
		Txid:          "01e2c94ec7b36241543c019686b8cc8727ca34a7970a9771bb202180381946d4",
		Vout:          0,
		Value:         "2396",
		Height:        52309,
		Confirmations: 15,
	}
	entry2 := utxo_models.UTXOEntry{
		Txid:          "8958472f189e153fb5d7e085fb586cda29e88245495a7d010a71348ad10196ec",
		Vout:          13,
		Value:         "2425",
		Height:        26316,
		Confirmations: 235,
	}
	entry3 := utxo_models.UTXOEntry{
		Txid:          "fc3d5031e4589cc99ad69ee908854a99c689342c4c9923f9a28d5d81142fa4f0",
		Vout:          1,
		Value:         "239623",
		Height:        23214,
		Confirmations: 19874,
	}

	correctBalance := utxo_models.UTXOBalance{
		Confirmed:   244444,
		Unconfirmed: 0,
	}

	nonEmptyUTXOList := utxo_models.UTXOList{
		Entries: []utxo_models.UTXOEntry{entry1, entry2, entry3},
	}
	evaluatedBalance := utxo_models.EvaluateList(nonEmptyUTXOList)

	if evaluatedBalance.Confirmed != correctBalance.Confirmed || evaluatedBalance.Unconfirmed != correctBalance.Unconfirmed {
		t.Errorf("List evaluating failed, expected %v, got %v", correctBalance, evaluatedBalance)
	}

}
