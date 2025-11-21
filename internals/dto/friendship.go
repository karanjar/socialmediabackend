package dto

import "github.com/google/uuid"

type FriendsCreate struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	FriendID uuid.UUID `json:"friend_id" validate:"required"`
}

type FriendsUpdate struct {
	FriendID uuid.UUID `json:"friend_id" validate:"required"`
	UserID   uuid.UUID `json:"user_id" validate:"required"`
}
