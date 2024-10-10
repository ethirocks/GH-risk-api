package v1

import (
	"net/http"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/ethirajmudhaliar/GH-risk-api/logger"
	"github.com/gorilla/mux"
)

// GetRiskByID to get a specific risk by ID
func GetRiskByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	riskID := vars["id"]

	// Get the risk by ID from the storage
	risk, err := common.Storage.GetRiskByID(riskID)
	if err != nil {
		logger.Info("Risk with ID not found" + riskID)
		logger.Info("Error: " + err.Error())
		common.RespondWithError(w, http.StatusNoContent, err.Error())
		return
	}

	// Log the successful retrieval
	logger.Info("Returning risk with ID: " + riskID)

	// Respond with the specific risk
	common.RespondWithJSON(w, http.StatusOK, risk)
}
