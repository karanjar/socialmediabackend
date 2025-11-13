package friendship

import (
	"socialmediabackend/models/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendship struct {
	gorm.Model
	UserID       uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendshipID uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"friendship_id"`

	User       users.Users `gorm:"foreignkey:UserID;references:ID" json:"user"`
	Friendship users.Users `gorm:"foreignKey:FriendshipID;references:ID" json:"-"`
}
