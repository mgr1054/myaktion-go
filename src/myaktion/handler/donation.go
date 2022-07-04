package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/mgr1054/myaktion-go/src/myaktion/model"
	"github.com/mgr1054/myaktion-go/src/myaktion/service"
)

func AddDonation(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	donation, err := getDonation(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: if the campaign doesn't exist, return 404 - don't show FK error
	err = service.AddDonation(id, donation)
	if err != nil {
		log.Errorf("Failure adding donation to campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, donation)
}

func getDonation(r *http.Request) (*model.Donation, error) {
	var donation model.Donation
	err := json.NewDecoder(r.Body).Decode(&donation)
	if err != nil {
		log.Errorf("Can't serialize request body to donation struct: %v", err)
		return nil, err
	}
	return &donation, nil
}
