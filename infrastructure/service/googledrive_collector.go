package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jteeuwen/go-bindata"
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

	driveService, err := g.newDrive()
	if err != nil{
		return nil, errors.Wrap(err, "Unable to connect to Goolge Drive")
	}

	musicFolder, err := driveService.Files.List().
		Q("mimeType = 'application/vnd.google-apps.folder' and name = 'Music'").Do()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get file list")
	}

	if len(musicFolder.Files) == 0 {
		testFile := drive.File{
			MimeType: "application/vnd.google-apps.folder",
			Name:     "Music",
		}
		_, err := driveService.Files.Create(&testFile).Fields("id").Do()
		if err != nil {
			fmt.Println("Unable to create folder: %v", err)
		}
		return *songList, nil
	}

	driveQuery := ""
	if songName == "" {
		driveQuery = "'" + musicFolder.Files[0].Id + "' in parents and trashed = false "
	} else {
		driveQuery = "'" + musicFolder.Files[0].Id + "' in parents and trashed = false " +
			"and name contains " + "'" + songName + "'"
	}

	musicFiles, err := driveService.Files.List().
		Q(driveQuery).Do()

	if err != nil {
		return nil, errors.Wrap(err, "Unable to get song list")
	}

	for _, file := range musicFiles.Files {
		songList.Add(domain.NewSongWithName(file.Name))
	}

	return *songList, nil
}

func (g *GoogleDriveCollector) newDrive() (*drive.Service, error) {
	config, err := g.getConfig()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to parse client secret file to Config")
	}

	token, err := g.getToken()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to get token")
	}

	client := config.Client(context.Background(), token)
	service, err := drive.New(client)
	if err != nil {
		return nil, errors.Wrap(err, "Unable to retrieve Drive client")
	}
	return service, nil
}


func (g *GoogleDriveCollector) chich(config *oauth2.Config) string{
	config.RedirectURL = "https://caidmbanbeonhe.com"

}



// parse credential file and return oauth2 Config
func (g *GoogleDriveCollector) getConfig() (*oauth2.Config, error) {
	credentials, err := ioutil.ReadFile("google-credentials.json")
	if err != nil {
		return nil, err
	}
	return google.ConfigFromJSON(credentials, drive.DriveScope)
}

func (g *GoogleDriveCollector) getToken() (*oauth2.Token, error) {
	f, err := os.Open("token.json")
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}
