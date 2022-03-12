package model

//Device model for devices
type Device struct {
	GUID   string `json:"guid"`
	UserID string `json:"user_id"`
	Model  string `json:"model"`
}
