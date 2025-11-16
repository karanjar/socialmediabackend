package friendship

import (
	"socialmediabackend/models/users"
	"time"

	"github.com/google/uuid"
)

type Friendship struct {
	ID           uuid.UUID `json:"id"`
	UserID       uuid.UUID ` json:"user_id"`
	FriendshipID uuid.UUID ` json:"friendship_id"`

	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User       users.Users `json:"user"`
	Friendship users.Users ` json:"friend"`
}
