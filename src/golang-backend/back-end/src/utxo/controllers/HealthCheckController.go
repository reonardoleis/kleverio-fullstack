package utxo_controllers

import (
	"encoding/json"
	"net/http"

	utxo_models "github.com/reonardoleis/klever.io-backend/src/utxo/models"
)

func CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//checks if the blockbook API is up
	_, err := http.Get(BLOCKBOOK_BASE_URL + "/utxo/")
	if err != nil {
		currentStatus := utxo_models.HealthStatus{Status: "down"}
		outputStatus, _ := json.Marshal(currentStatus)
		w.WriteHeader(400)
		w.Write(outputStatus)
		return
	}
	currentStatus := utxo_models.HealthStatus{Status: "up"}
	outputStatus, _ := json.Marshal(currentStatus)
	w.Write(outputStatus)
}
