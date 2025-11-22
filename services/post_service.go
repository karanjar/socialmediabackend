package services

import (
	"context"
	"encoding/json"
	"fmt"
	"socialmediabackend/internals/database"
	"socialmediabackend/internals/dto"
	"socialmediabackend/models/posts"
	"time"

	"github.com/supabase-community/postgrest-go"

	"github.com/google/uuid"
)

type PostsService struct{}

func NewPostsService() *PostsService {
	return &PostsService{}
}

// Create Posts
func (*PostsService) CreatePost(ctx context.Context, input dto.CreatePost) (*posts.Post, error) {
	PostID := uuid.New()
	now := time.Now()

	p := posts.Post{
		ID:        PostID,
		Content:   input.Content,
		ImageUrl:  input.ImageURL,
		UserID:    input.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	records := map[string]interface{}{
		"id":         p.ID,
		"user_id":    p.UserID,
		"content":    p.Content,
		"image_url":  p.ImageUrl,
		"created_at": p.CreatedAt,
		"updated_at": p.UpdatedAt,
	}

	data, _, err := database.SupabaseClient.
		From("posts").
		Insert(records, false, "", "representation", "").
		Execute()
	if err != nil {
		return nil, fmt.Errorf("superbase create post error: %w", err)
	}

	var result []posts.Post

	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("post id %s is empty", PostID)
	}

	return &result[0], nil
}

// Get ALl Post
func (*PostsService) GetAllPosts(ctx context.Context) ([]posts.Post, error) {

	data, _, err := database.SupabaseClient.
		From("posts").
		Select("* user:user_id(*)", "exact", false).
		Order("created_at", &postgrest.OrderOpts{Ascending: false}).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("superbase get all posts error: %w", err)
	}
	var result []posts.Post
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	return result, nil
}

// Get post by Id
func (*PostsService) GetPostByID(ctx context.Context, postID uuid.UUID) (*posts.Post, error) {
	data, _, err := database.SupabaseClient.
		From("posts").
		Select("* user:user_id(*)", "exact", false).
		Eq("id", postID.String()).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("superbase get post by id error: %w", err)
	}
	var result []posts.Post
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("post id %s is empty", postID)
	}
	return &result[0], nil
}

// Update Posts

func (*PostsService) UpdatePost(ctx context.Context, postID uuid.UUID, input dto.CreatePost) (*posts.Post, error) {
	// map records
	update := map[string]interface{}{
		"content":    input.Content,
		"image_url":  input.ImageURL,
		"updated_at": time.Now(),
	}

	data, _, err := database.SupabaseClient.
		From("posts").
		Update(update, "representation", "").
		Eq("id", postID.String()).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("superbase create post error: %w", err)
	}
	var result []posts.Post
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("json decode error: %w", err)
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("post id %s is empty", postID)
	}
	return &result[0], nil

}

// Deleteing post
func (*PostsService) DeletePost(ctx context.Context, postID string) error {
	_, _, err := database.SupabaseClient.
		From("posts").
		Delete("", "").
		Eq("id", postID).
		Execute()
	if err != nil {
		return fmt.Errorf("superbase delete post error: %w", err)
	}
	return nil
}
