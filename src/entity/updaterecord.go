package entity

type UpdateRecord struct {
	UpdateID int64 `json:"last_update_id"`
	UpdateAT int64 `json:"last_update_at"`
}
