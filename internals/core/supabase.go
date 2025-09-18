package core

import (
    "os"
	storage_go "github.com/supabase-community/storage-go"
)

func SupabaseClient() *storage_go.Client {
    url := os.Getenv("SUPABASE_URL")
    key := os.Getenv("SUPABASE_ANON_KEY")
    client := storage_go.NewClient(url+"/storage/v1", key, nil)
    return client
}