package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/infrastructure/service"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
	"net/http"
)

type GoogleLoginController struct {
	googleDriveService *service.GoogleDrive
}

func NewGoogleLoginController() (*GoogleLoginController, error) {
	googleDriveService, err := service.NewGoogleDrive()
	if err != nil {
		return nil, errors.Wrap(err, "Fail to connect to Google Drive")
	}
	return &GoogleLoginController{
		googleDriveService: googleDriveService,
	}, nil
}

func (g *GoogleLoginController) RedirectToGoogleProvider(ctx *gin.Context) {
	authURL := g.googleDriveService.Config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	ctx.Redirect(http.StatusMovedPermanently, authURL)
}

func (g *GoogleLoginController) HandleProviderCallBack(ctx *gin.Context) {
	err := g.googleDriveService.NewTokenFile(ctx.Query("code"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"connectedToGG": false,
		})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8088")
}
