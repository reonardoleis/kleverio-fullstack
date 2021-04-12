package utxo_controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	utxo_models "github.com/reonardoleis/klever.io-backend/src/utxo/models"
)

const BLOCKBOOK_BASE_URL = "https://blockbook-bitcoin.tronwallet.me/api/v2"

func GetBalance(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	desiredAddress := mux.Vars(r)["address"]
	log.Println("New request for address " + desiredAddress + " [" + r.RemoteAddr + "]")
	resp, err := http.Get(BLOCKBOOK_BASE_URL + "/utxo/" + desiredAddress)
	if err != nil {
		currentError := utxo_models.UTXOError{Message: "API_FETCHING_ERROR"}
		outputError, _ := json.Marshal(currentError)
		w.WriteHeader(400)
		w.Write(outputError)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		currentError := utxo_models.UTXOError{Message: "API_PARSING_FETCHING_ERROR"}
		outputError, _ := json.Marshal(currentError)
		w.WriteHeader(400)
		w.Write(outputError)
		return
	}
	var currentUTXOList utxo_models.UTXOList
	json.Unmarshal(body, &currentUTXOList.Entries)
	if err != nil {
		currentError := utxo_models.UTXOError{Message: "DECODING_ERROR"}
		outputError, _ := json.Marshal(currentError)
		w.WriteHeader(400)
		w.Write(outputError)
		return
	}

	if currentUTXOList.Entries == nil {
		currentUTXOList.Entries = []utxo_models.UTXOEntry{}
	}

	outputBalance := utxo_models.EvaluateList(currentUTXOList)

	jsonOutputBalance, _ := json.Marshal(outputBalance)
	w.WriteHeader(200)
	w.Write(jsonOutputBalance)
}
