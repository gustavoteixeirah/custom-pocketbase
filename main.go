package main

import (
	"github.com/spf13/cast"
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// load app settings from env variables
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// Load SMTP settings
		app.Settings().Smtp.Enabled = cast.ToBool(os.Getenv("SMTP_ENABLED"))
		app.Settings().Smtp.Host = os.Getenv("SMTP_HOST")
		app.Settings().Smtp.Port = cast.ToInt(os.Getenv("SMTP_PORT"))
		app.Settings().Smtp.Username = os.Getenv("SMTP_USERNAME")
		app.Settings().Smtp.Password = os.Getenv("SMTP_PASSWORD")
		err := app.Settings().Smtp.Validate()
		if err != nil {
			log.Fatal("Error when trying to connect to Email Server through SMTP: ", err)
			return err
		}

		// Load S3 settings
		app.Settings().S3.Enabled = cast.ToBool(os.Getenv("S3_ENABLED"))
		app.Settings().S3.Endpoint = os.Getenv("S3_ENDPOINT")
		app.Settings().S3.Bucket = os.Getenv("S3_BUCKET")
		app.Settings().S3.AccessKey = os.Getenv("S3_ACCESS_KEY")
		app.Settings().S3.Secret = os.Getenv("S3_SECRET")
		app.Settings().S3.Region = os.Getenv("S3_REGION")
		app.Settings().S3.ForcePathStyle = cast.ToBool(os.Getenv("S3_FORCE_PATH_STYLE"))

		err = app.Settings().S3.Validate()
		if err != nil {
			log.Fatal("Error when trying to connect to S3: ", err)
			return err
		}

		return nil
	})

	// prevent settings change...
	// (remaining unchanged from your original code)

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
