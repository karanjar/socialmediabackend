package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"socialmediabackend/internals/database"
	"socialmediabackend/internals/dto"
	"socialmediabackend/models/friendship"
	"time"

	"github.com/google/uuid"
)

type FriendshipService struct{}

func NewFriendshipService() *FriendshipService {
	return &FriendshipService{}
}

func (s *FriendshipService) sendFriendrequest(ctx context.Context, input dto.FriendsCreate) (*friendship.Friendship, error) {

	if input.UserID == input.FriendID {
		return nil, errors.New("you cannot send friendship to your self ")
	}
	//generate uuid
	id := uuid.New()

	now := time.Now()

	f := friendship.Friendship{
		ID:           id,
		UserID:       input.UserID,
		FriendshipID: input.FriendID,
		Status:       "pending",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	//mapping the friendship for superbase

	records := map[string]interface{}{
		"id":            f.ID,
		"user_id":       f.UserID,
		"friendship_id": f.FriendshipID,
		"status":        f.Status,
		"created_at":    f.CreatedAt,
		"updated_at":    f.UpdatedAt,
	}

	data, _, err := database.SupabaseClient.
		From("friendship").
		Insert(records, false, "", "representation", "").
		Execute()
	if err != nil {
		return nil, fmt.Errorf("superbase create friendship: %w", err)
	}
	var result []friendship.Friendship
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decoded error: %w", err)
	}

	if len(result) == 0 {
		return nil, errors.New("no friendship found/inserted returned no friendship")
	}

	return &result[0], nil

}
