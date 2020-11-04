package handler

import (
	"app/infra"
	"app/infra/rest/converter"
	"fmt"
	"net/http"
	"time"
)

//About get app info
func About(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	success := true
	statusMsg := ""
	alog, _ := LogInit(r, "Info About my app")
	defer alog.LogMessageEnd(&success, &statusMsg)
	alog.LogMessageReceived()
	dbAppInfo, err := infra.GetUseCaseInteractor().Db.GetApp("gormv2-poc")
	if err != nil {
		success = false
		statusMsg = fmt.Sprintf("Error in getting app info, reason %s", err.Error())
		HandleErrorResponse(w, statusMsg, http.StatusNotFound)
		return
	}
	fitnessApp, _ := converter.GetAppSwaggerResponse(&dbAppInfo)
	HandleSuccessResponse(w, fitnessApp)
	statusMsg = "Get successful"
}

//BulkInsert bulk insert
func BulkInsert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	success := true
	statusMsg := ""
	alog, _ := LogInit(r, "Bulk Insert app")
	defer alog.LogMessageEnd(&success, &statusMsg)
	alog.LogMessageReceived()
	startTime := time.Now()
	infra.GetUseCaseInteractor().Db.BulkInsertFabrics()
	elapsedTime := time.Since(startTime)
	HandleGenericSuccess(w, "Bulk insert elapsed time:"+elapsedTime.String())
	statusMsg = "Get successful"
}
