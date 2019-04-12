package service

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"os"
)

type GoogleDrive struct {
	Config       *oauth2.Config
}

func NewGoogleDrive() (*GoogleDrive, error) {
	config, err := getConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse client secret file to Config")
	}

	return &GoogleDrive{
		Config:       config,
	}, nil
}

func getConfig() (*oauth2.Config, error) {
	credentials, err := ioutil.ReadFile("google-credentials.json")
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(credentials, drive.DriveScope)
}

func (g *GoogleDrive) NewTokenFile(tokenCode string) error {
	token, err := g.Config.Exchange(context.TODO(), tokenCode)
	if err != nil {
		return errors.Wrap(err, "Unable to retrieve token from web")
	}
	file, err := os.OpenFile("token.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		return errors.Wrap(err, "Unable to cache oauth token")
	}
	json.NewEncoder(file).Encode(token)
	return nil
}
