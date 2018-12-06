package repository

import (
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/adapter/persistent/service"
	"github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"
	"github.com/pkg/errors"
)

type SongResultPool struct {
	collectorList service.CollectorList
}

func NewSongResultPool(
	collectorList service.CollectorList,
) *SongResultPool {
	return &SongResultPool{
		collectorList: collectorList,
	}
}

func (s *SongResultPool) GetSongListByName(songName string) (*domain.SongCollectorMap, error) {
	resultSongMap := domain.EmptySongCollectorMap()

	collectorChan := make(chan domain.CollectorType, len(s.collectorList))
	errorChan := make(chan error, len(s.collectorList))

	for _, collector := range s.collectorList {
		go func(collector service.Collector) {
			songList, err := collector.Search(songName)
			if err != nil {
				errorChan <- err
			}
			resultSongMap.Add(collector.GetCollectorType(), &songList)
			collectorChan <- collector.GetCollectorType()
		}(collector)
	}

	for {
		select {
		case err := <-errorChan:
			return nil, errors.Wrap(err, "Error fetching song from collector")
		case <-collectorChan:
			if len(*resultSongMap) == len(s.collectorList) {
				return resultSongMap, nil
			}
		}
	}
}