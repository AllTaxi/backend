package api

import (
	"encoding/json"
	"net/http"
)

func (api *api) registerDeviceV2(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Middleware: Parse JWT and Inject UserID into Ctx (request)
	// userID := auth.GetUserIDFromCtx(ctx)
	userID := "12345"

	type payload struct {
		Model      string `json:"model"`
		Token      string `json:"token"`
		AppVersion string `json:"app_version"`
		Language   string `json:"language"`
	}

	var body payload

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		writeError(w, "error parsing request body", http.StatusBadRequest)
		return
	}

	device, err := api.deviceService.CreateDevice(ctx, userID, body.Model)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"user_guid": userID,
		"device":    device,
	}

	writeJSON(w, response)
}
