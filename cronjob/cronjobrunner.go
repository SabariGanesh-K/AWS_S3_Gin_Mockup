package cron

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/robfig/cron/v3"
)

func DeleteExpiredFiles(ctx context.Context, s3Client *s3.Client, db *db.Queries) error {
	fmt.println("fetching and eleting old files...")
    return nil 
}

func StartBackgroundJob(ctx context.Context, s3Client *s3.Client, db *db.Queries) error {
    cfg, err := config.LoadDefaultConfig(ctx)
    if err != nil {
        return err
    }

    if s3Client == nil {
        s3Client = s3.NewFromConfig(cfg)
    }

    c := cron.New()
    c.AddFunc("@hourly", func() {
        err := DeleteExpiredFiles(ctx, s3Client, db)
        if err != nil {
            fmt.Println("Error deleting expired files:", err)
        }
    })

    go c.Start() 
    return nil
}