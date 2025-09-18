package repositories

import (
    "os"
    "mime/multipart"
    "keuangan/backend/internals/core"
)

type PostImageRepository struct{}

func NewPostImageRepository() *PostImageRepository {
    return &PostImageRepository{}
}

func (r *PostImageRepository) UploadImage(file multipart.File, fileName string) (string, error) {
    bucket := os.Getenv("SUPABASE_BUCKET")
    client := core.SupabaseClient()

    // Upload file to Supabase Storage
    _, err := client.UploadFile(bucket, fileName, file)
    if err != nil {
        return "", err
    }

    // Construct public URL
    url := os.Getenv("SUPABASE_URL")
    publicURL := url + "/storage/v1/object/public/" + bucket + "/" + fileName
    return publicURL, nil
}