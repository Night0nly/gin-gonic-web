package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/repository"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/usecase/interactor"
	"net/http"
)

type SongController struct {
	interactor *interactor.SongSearch
}

func NewSongController(songResultPool *repository.SongResultPool) *SongController {
	return &SongController{
		interactor: interactor.NewSongSearch(songResultPool),
	}
}

func (s *SongController) SearchSongByName(context *gin.Context) {
	songMap, err := s.interactor.SearchByName(context.Query("songname"))
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, songMap)
}

func (s *SongController) GetAllSong(context *gin.Context) {
	songMap, err := s.interactor.SearchByName("")
	if err != nil {
		context.JSON(500, gin.H{
			"connectedToGG": false,
		})
		return
	}
	context.JSON(http.StatusOK, songMap)
}
