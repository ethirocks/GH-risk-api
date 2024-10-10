package v1

import (
	"net/http"
	"strconv"

	"github.com/ethirajmudhaliar/GH-risk-api/common"
	"github.com/ethirajmudhaliar/GH-risk-api/logger"
)

// GetRisks to get all risks
func GetRisks(w http.ResponseWriter, r *http.Request) {
	// Retrieve all risks
	riskList, err := common.Storage.GetAllRisks()

	// If there are no risks, return 204 No Content
	if err != nil {
		logger.Error("Error: " + err.Error())
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Log the number of risks returned
	logger.Info("Returning " + strconv.Itoa(len(riskList)) + " risks")

	// If risks exist, respond with the list of risks
	common.RespondWithJSON(w, http.StatusOK, riskList)
}
