package dto

import "time"

type AuditDto struct {
	RequestID   string    `json:"request_id"`
	ProductID   string    `json:"product_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
