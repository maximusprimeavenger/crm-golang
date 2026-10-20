package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty"`
	UserID       string             `json:"user_id"`
	Name         *string            `json:"name" validate:"required"`
	Email        *string            `json:"email" validate:"required"`
	Phone        *string            `json:"phone" validate:"required"`
	Password     *string            `json:"password" validate:"required"`
	Type         *string            `json:"type" validate:"eq=ADMIN|eq=USER"`
	Token        string             `json:"token"`
	RefreshToken string             `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}

type PasswordResetToken struct {
	ID        primitive.ObjectID `json:"id,omitempty"`
	UserID    string             `json:"user_id"`
	TokenHash string             `json:"token_hash"`
	ExpiresAt time.Time          `json:"expires_at"`
	Used      bool               `json:"used"`
}
