package models

// Error - Object representing an error
type Error struct {

	// Error code that identify of the error
	Code int32 `json:"code,omitempty"`

	// Short description of the error
	Message string `json:"message,omitempty"`
}
