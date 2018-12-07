package service

import (
	"context"
	"encoding/json"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"io/ioutil"
	"os"
)

type GoogleDriveCollector struct {
	service.Collector
}

func NewGoogleDriveCollector() *GoogleDriveCollector {
	return &GoogleDriveCollector{}
}

func (g *GoogleDriveCollector) GetCollectorType() domain.CollectorType {
	return domain.GOOGLEDRIVE
}

func (g *GoogleDriveCollector) GetSongList(songName string) (domain.SongList, error) {
	songList := domain.EmptySongList()

	config, err := getConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse client secret file to config")
	}

	token, err := getToken()
	if err != nil {
		return nil, errors.Wrap(err,"Unable to get token")
	}

	client := config.Client(context.Background(), token)
	srv, err := drive.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to retrieve Drive client")
	}

	musicFolder, err := srv.Files.List().
		Q("mimeType = 'application/vnd.google-apps.folder' and name = 'Music'").Do()
	if err != nil {
		return nil, errors.Wrap(err, "Music folder not found")
	}

	driveQuery := ""
	if songName == "" {
		driveQuery = "'" + musicFolder.Files[0].Id +"' in parents and trashed = false "
	} else {
		driveQuery = "'" + musicFolder.Files[0].Id +"' in parents and trashed = false " +
			"and name contains " + "'" + songName + "'"
	}

	musicFiles, err := srv.Files.List().
		Q(driveQuery).Do()

	if err != nil {
		return nil, nil
	}

	for _, file := range musicFiles.Files{
		songList.Add(domain.NewSongWithName(file.Name))
	}

	return *songList, nil
}

func getConfig() (*oauth2.Config, error) {
	credentials, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(credentials, drive.DriveMetadataReadonlyScope)
}

func getToken() (*oauth2.Token, error) {
	f, err := os.Open("token.json")
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
