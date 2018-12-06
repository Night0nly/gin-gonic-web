package domain

type SongList []*Song

func EmptySongList() *SongList {
	return &SongList{}
}

func (s *SongList) Add(song *Song) {
	if song == nil {
		return
	}
	*s = append(*s, song)
}

func (s *SongList) AddAll(songList *SongList) {
	for _, song := range *songList {
		s.Add(song)
	}
}
