package domain

type SongCollectorMap map[CollectorType]*SongList

func EmptySongCollectorMap() *SongCollectorMap {
	return &SongCollectorMap{}
}

func (s *SongCollectorMap) Add(colectorType CollectorType, songList *SongList) {
	(*s)[colectorType] = songList
}
