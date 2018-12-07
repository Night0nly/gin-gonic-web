package service

import "github.com/mediadotech/FY2018_2H_fp_team_training_hung/gin-gonic-web/domain"

type Collector interface {
	GetSongList(query string) (domain.SongList, error)
	GetCollectorType() domain.CollectorType
}

type CollectorList []Collector

func EmptyCollectorList() *CollectorList {
	return &CollectorList{}
}

func (c *CollectorList) Add(colector Collector) {
	*c = append(*c, colector)
}
