package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"socialmediabackend/internals/database"
	"socialmediabackend/internals/dto"
	"socialmediabackend/models/users"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct{}

// Constructor
func NewUserService() *UserService {
	return &UserService{}
}

// CREATE USER
func (s *UserService) CreateUser(ctx context.Context, input dto.Usercreate) (*users.Users, error) {
	// Generate UUID
	id := uuid.New()

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("password hash error: %w", err)
	}

	now := time.Now()

	u := users.Users{
		ID:        id,
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Prepare map for Supabase
	record := map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"password":   u.Password,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}

	data, _, err := database.SupabaseClient.
		From("users").
		Insert(record, false, "", "representation", "").
		Execute()
	if err != nil {
		return nil, fmt.Errorf("supabase insert error: %w", err)
	}

	var result []users.Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("insert returned no user")
	}

	return &result[0], nil
}

// ===== GET USER BY ID =====
func (s *UserService) GetUserByID(ctx context.Context, id string) (*users.Users, error) {
	data, _, err := database.SupabaseClient.
		From("users").
		Select("*", "exact", false). // columns, count, head
		Eq("id", id).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("supabase get error: %w", err)
	}

	var result []users.Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

// ===== UPDATE USER =====
func (s *UserService) UpdateUser(ctx context.Context, id string, input dto.Userupdate) (*users.Users, error) {
	update := map[string]interface{}{
		"name":       input.Name,
		"email":      input.Email,
		"updated_at": time.Now(),
	}

	if input.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, fmt.Errorf("password hash error: %w", err)
		}
		update["password"] = string(hashedPassword)
	}

	data, _, err := database.SupabaseClient.
		From("users").
		Update(update, "representation", "").
		Eq("id", id).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("supabase update error: %w", err)
	}

	var result []users.Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

// ===== DELETE USER =====
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	_, _, err := database.SupabaseClient.
		From("users").
		Delete("", ""). // returning, count
		Eq("id", id).
		Execute()
	if err != nil {
		return fmt.Errorf("supabase delete error: %w", err)
	}

	return nil
}
