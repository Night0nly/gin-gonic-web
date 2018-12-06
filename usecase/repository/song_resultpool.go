package repository

import "github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"

type SongResultPool interface {
	GetSongListByName(songName string) (*domain.SongCollectorMap, error)
}
