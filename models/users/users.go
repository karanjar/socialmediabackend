package users

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/supabase-community/supabase-go"
)

type Users struct {
	ID        uuid.UUID `json:"id"`
	Name      string    ` json:"name"`
	Email     string    ` json:"email"`
	Password  string    ` json:"password"`
	CreatedAt time.Time ` json:"created_at"`
	UpdatedAt time.Time ` json:"updated_at"`
}

func CreateUser(client *supabase.Client, u Users) (*Users, error) {
	// Prepare data to insert (use map for supabase-go)
	record := map[string]interface{}{
		"id":         u.ID,
		"name":       u.Name,
		"email":      u.Email,
		"password":   u.Password,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
	data, _, err := client.From("users").Insert(record, false, "", "representation", "").Execute()
	if err != nil {
		return nil, fmt.Errorf("supabase insert error: %w", err)
	}
	var result []Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("insert returned no user")
	}
	return &result[0], nil
}

func GetUserByID(client *supabase.Client, id string) (*Users, error) {
	data, _, err := client.
		From("users").
		Select("*", "exact", false). // ✔️ correct order
		Eq("id", id).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("supabase get error: %w", err)
	}

	var result []Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if len(result) == 0 {
		return nil, nil // not found
	}

	return &result[0], nil
}

func UpdateUser(client *supabase.Client, id string, u Users) (*Users, error) {
	update := map[string]interface{}{
		"name":       u.Name,
		"email":      u.Email,
		"password":   u.Password,
		"updated_at": u.UpdatedAt,
	}

	data, _, err := client.
		From("users").
		Update(update, "representation", "").
		Eq("id", id).
		Execute()

	if err != nil {
		return nil, fmt.Errorf("supabase update error: %w", err)
	}

	var result []Users
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}

	if len(result) == 0 {
		return nil, nil // not updated
	}

	return &result[0], nil
}

func DeleteUser(client *supabase.Client, id string) error {
	_, _, err := client.
		From("users").
		Delete("", ""). // count, returning, head
		Eq("id", id).
		Execute()

	if err != nil {
		return fmt.Errorf("supabase delete error: %w", err)
	}

	return nil
}
