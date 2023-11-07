package model

import "time"

type EntityID string

type EntityStatus int

const (
	StatusPending   EntityStatus = 0
	StatusAvailable EntityStatus = 1
	StatusError     EntityStatus = 100
	StatusExpired   EntityStatus = 101
	StatusLocked    EntityStatus = 102
	StatusCancelled EntityStatus = 103
	StatusDeleted   EntityStatus = 999
)

type User struct {
	ID            EntityID     `json:"ID"`
	CreatedAt     time.Time    `json:"createdAt"`
	UpdatedAt     *time.Time   `json:"updatedAt,omitempty"`
	LastLoginAt   *time.Time   `json:"lastLoginAt"`
	Status        EntityStatus `json:"status"`
	LoginAttempts int          `json:"loginAttempts"`
	Name          string       `json:"name"`
	Email         string       `json:"email"`
	Password      string       `json:"password"`
}
