package model

type (
	FindUserCorrelationsRequest struct {
		UserID     int  `json:"userID"`
		WithShared bool `json:"withShared"`
	}
)
