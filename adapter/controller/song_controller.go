package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/repository"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
	service2 "github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/infrastructure/service"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/usecase/interactor"
)

type SongController struct {
	interactor *interactor.SongSearch
	context    *gin.Context
}

func NewSongController(context *gin.Context) *SongController {
	collectorList := service.EmptyColectorList()
	collectorList.Add(service2.NewChiasenhacCollector())
	return &SongController{
		interactor: interactor.NewSongSearch(repository.NewSongResultPool(*collectorList)),
		context:    context,
	}
}

func (s *SongController) SearchSongByName() {
	songMap, err := s.interactor.SearchByName(s.context.Query("songname"))
	if err != nil {
		s.context.JSON(500, gin.H{
			"message": err,
		})
		return
	}
	s.context.JSON(http.StatusOK, songMap)
}
