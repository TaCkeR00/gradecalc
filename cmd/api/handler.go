package main

import (
	"encoding/json"
	"log"
	"net/http"

	"simple-api/calculator"
)

func calculateGradeHandler(w http.ResponseWriter, r *http.Request) {
	modules, err := calculator.FromJSON(r.Body)
	if err != nil {
		log.Println("error reading json:", err)
		http.Error(w, "error reading json file", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	modules.CalculateGrade()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(calculateResponse{
		Grade:  modules.Grade(),
		Credit: modules.Credit(),
	})
}
