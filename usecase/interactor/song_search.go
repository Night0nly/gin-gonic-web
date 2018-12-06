package interactor

import (
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/usecase/repository"
	"github.com/pkg/errors"
)

type SongSearch struct {
	songPool repository.SongResultPool
}

func NewSongSearch(
	pool repository.SongResultPool,
) *SongSearch {
	return &SongSearch{
		songPool: pool,
	}
}

func (s *SongSearch) SearchByName(songName string) (*domain.SongCollectorMap, error) {
	songCollectorMap, err := s.songPool.GetSongListByName(songName)

	if err != nil {
		return nil, errors.Wrap(err, "Error on fetch data from pool!\n")
	}

	return songCollectorMap, nil
}
