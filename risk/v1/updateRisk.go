package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/ethirajmudhaliar/GH-risk-api/logger"
	"github.com/ethirajmudhaliar/GH-risk-api/validation"
	"github.com/gorilla/mux"
)

// UpdateRisk handles updating an existing risk
func UpdateRisk(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	riskID := vars["id"]

	// Retrieve the existing risk
	existingRisk, err := common.Storage.GetRiskByID(riskID)
	if err != nil {
		logger.Error("Risk with ID " + riskID + " not found")
		common.RespondWithError(w, http.StatusNoContent, "Risk not found")
		return
	}

	// Parse the JSON body for the update
	var updatedRisk common.Risk
	err = json.NewDecoder(r.Body).Decode(&updatedRisk)
	if err != nil {
		logger.Error("Error decoding request body: " + err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update fields if they are provided
	if updatedRisk.State != "" {
		if err := validation.ValidateState(updatedRisk.State); err != nil {
			logger.Error(err.Error())
			common.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		existingRisk.State = updatedRisk.State
	}

	if updatedRisk.Title != "" {
		existingRisk.Title = updatedRisk.Title
	}

	if updatedRisk.Description != "" {
		existingRisk.Description = updatedRisk.Description
	}

	common.Storage.UpdateRisk(riskID, existingRisk)
	logger.Info("Updated risk with ID: " + riskID)
	common.RespondWithJSON(w, http.StatusOK, existingRisk)
}
