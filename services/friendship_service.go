package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"socialmediabackend/internals/database"
	"socialmediabackend/internals/dto"
	"socialmediabackend/models/friendship"
	"socialmediabackend/models/users"
	"time"

	"github.com/google/uuid"
)

type FriendshipService struct{}

func NewFriendshipService() *FriendshipService {
	return &FriendshipService{}
}

//SEND REQUEST / UPDATE FRIENDSHIP

func (s *FriendshipService) SendFriendrequest(ctx context.Context, input dto.FriendsCreate) (*friendship.Friendship, error) {

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
		From("friendships").
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
	name := users.Users{}
	fmt.Printf("you have friend request from :%v", name.Name)
	return &result[0], nil

}

// GET ALL FRIENDS
func (s FriendshipService) Getfriends(ctx context.Context, userId uuid.UUID) ([]friendship.Friendship, error) {
	idStr := userId.String()

	orFilter := fmt.Sprintf("user_id.eq.%s,friendship_id.eq.%s", idStr, idStr)

	data, _, err := database.SupabaseClient.
		From("friendships").
		Select("*,user:user_id(*),friend:friendship_id(*)", "exact", false).
		Or(orFilter, "").
		Eq("status", "accepted").
		Execute()

	if err != nil {
		return nil, fmt.Errorf("superbase get friends: %w", err)
	}
	var result []friendship.Friendship
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decoded error: %w", err)
	}
	return result, nil
}

//GETTING FRIEND BT ID

func (s *FriendshipService) GetfriendById(ctx context.Context, userID uuid.UUID) (*friendship.Friendship, error) {
	useridstr := userID.String()

	//filtering users
	selectQuery := "*, user:user_id(*), friend:friendship_id(*)"

	//combining filtered users with accepted status
	orFilter := fmt.Sprintf("user_id.eq.%s,friendship_id.eq.%s", useridstr, useridstr)
	data, _, err := database.SupabaseClient.
		From("friendships").
		Select(selectQuery, "exact", false).
		Eq("status", "accepted").
		Or(orFilter, "").
		Execute()

	if err != nil {
		return nil, fmt.Errorf("error superbase get friends: %w", err)
	}
	var result []friendship.Friendship

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decoded error: %w", err)
	}
	if len(result) == 0 {
		return nil, nil
	}
	for i := range result {
		result[i].User.Password = ""
		result[i].Friendship.Password = ""
	}
	return &result[0], nil
}

// Accepting friendrequest/Update friendship
func (s *FriendshipService) Updatefriends(ctx context.Context, userID, friendID uuid.UUID) (*friendship.Friendship, error) {

	useridstr := userID.String()
	friendidstr := friendID.String()

	records := map[string]interface{}{
		"status":     "accepted",
		"updated_at": time.Now(),
	}

	Filter := fmt.Sprintf("and(user_id.eq.%s,friendship_id.eq.%s),and(user_id.eq.%s,friendship_id.eq.%s)", useridstr, friendidstr, useridstr, friendidstr)

	data, _, err := database.SupabaseClient.
		From("friendships").
		Update(records, "representation", "").
		Or(Filter, "").
		Execute()

	if err != nil {
		return nil, fmt.Errorf("error superbase update friends: %w", err)
	}

	var result []friendship.Friendship

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decoded error: %w", err)
	}
	if len(result) == 0 {
		return nil, nil
	}
	return &result[0], nil

}
func (s *FriendshipService) DeleteFriendship(ctx context.Context, userID, friendID uuid.UUID) error {
	useridstr := userID.String()
	friendidstr := friendID.String()

	//filtered := fmt.Sprintf("and(user_id.eq.%s,friend_id.eq.%s)", useridstr, friendidstr)

	filter := fmt.Sprintf("or(and(user_id.eq.%s,friendship_id.eq.%s),and(user_id.eq.%s,friendship_id.eq.%s))", useridstr, friendidstr, friendidstr, useridstr)

	_, count, err := database.SupabaseClient.
		From("friendships").
		Delete("", "exact").
		Or(filter, "").
		Execute()

	if err != nil {
		return fmt.Errorf("error superbase delete friendship: %w", err)
	}

	if count == 0 {
		return errors.New("no friendship found")
	}
	return nil
}
