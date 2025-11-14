package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/supabase-community/supabase-go"
)

var SupabaseClient *supabase.Client

func Connect() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
		return err
	}

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL or SUPABASE_KEY missing in .env")
	}

	client, err := supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("Failed to initialize Supabase client:", err)
		return err
	}

	SupabaseClient = client // assign before using

	// Connectivity test with proper Limit args
	_, _, err = SupabaseClient.From("users").Select("*", "exact", false).Limit(1, "").Execute()
	if err != nil {
		log.Println("Supabase connectivity test failed:", err)
		return err
	}

	fmt.Println("Connected to Supabase client")
	return nil
}
