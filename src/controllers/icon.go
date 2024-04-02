package controllers

import (
	"log"
	"net/http"

	"github.com/Locally-build/acme-weather-backend/services"
	"github.com/gorilla/mux"
)

func IconHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	osSystem := params["os"]
	icon := params["icon"]

	log.Printf("Getting icon %s for os %s", icon, osSystem)
	svc := services.IconService{}
	iconPath, err := svc.GetIconURL(osSystem, icon)
	if err != nil {
		log.Print(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(iconPath))
	if err != nil {
		log.Print(err)
		return
	}
}
