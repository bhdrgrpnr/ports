package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"saasteamtest/backend/domain"
	"saasteamtest/backend/models"
)

func CreatePort(portService domain.PortServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}

		var request models.PortRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			return RespondBadRequest(w, nil)
		}

		savedPort, err := portService.Save(request)

		if err != nil {
			return fmt.Errorf("save: %w", err)
		}
		return RespondCreated(w, savedPort)
	}
}

func UpdatePort(portService domain.PortServiceInterface) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("readAll: %w", err)
		}

		var request models.PortRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			return RespondBadRequest(w, nil)
		}

		updatedPort, err := portService.Update(request)

		if err != nil {
			if errors.Is(err, domain.ErrIndexNotFound) {
				return RespondBadRequest(w, updatedPort)
			}
			return Respond500(w, updatedPort)
		}
		return RespondCreated(w, updatedPort)
	}
}
