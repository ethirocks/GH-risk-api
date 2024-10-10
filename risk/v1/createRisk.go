package v1

import (
	"encoding/json"
	"net/http"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/ethirajmudhaliar/GH-risk-api/logger"
	"github.com/ethirajmudhaliar/GH-risk-api/validation"
	"github.com/google/uuid"
)

// CreateRisk handles the creation of a new risk
func CreateRisk(w http.ResponseWriter, r *http.Request) {
	var newRisk common.Risk

	// Parse the JSON body
	err := json.NewDecoder(r.Body).Decode(&newRisk)
	if err != nil {
		logger.Error("Error decoding request body: " + err.Error())
		common.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Validate required fields (Title, Description)
	if newRisk.Title == "" || newRisk.Description == "" {
		logger.Error("Missing required fields in request")
		common.RespondWithError(w, http.StatusBadRequest, "Missing required fields")
		return
	}

	// Validate the state value using the validation package
	if err := validation.ValidateState(newRisk.State); err != nil {
		logger.Error(err.Error())
		common.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Generate a new UUID for the risk ID
	newRisk.ID = uuid.New().String()

	// Add the new risk to the storage
	if err := common.Storage.AddRisk(newRisk); err != nil {
		logger.Error("Error adding risk to storage: " + err.Error())
		common.RespondWithError(w, http.StatusInternalServerError, "Could not store the risk")
		return
	}

	logger.Info("Created new risk with ID: " + newRisk.ID)

	common.RespondWithJSON(w, http.StatusCreated, newRisk)
}
