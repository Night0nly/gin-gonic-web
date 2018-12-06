package domain

type Song struct {
	Name        string
	Author      string
	Length      string
	Url         string
	DownloadUrl DownloadURLList
}

func NewSong(name, author, length, url string, downloadUrl DownloadURLList) *Song {
	return &Song{
		Name:        name,
		Author:      author,
		Url:         url,
		Length:      length,
		DownloadUrl: downloadUrl,
	}
}

//
//func (s *Song) fmtDuration(d time.Duration) string {
//	durations := regexp.MustCompile(`\D`).Split(d.String(), 4)
//	for i := 0; i < len(durations)-1; i++ {
//		if time, _ := strconv.ParseInt(durations[i], 10, 64); time < 10 {
//			durations[i] = "0" + durations[i]
//		}
//	}
//	return strings.Join(durations[:len(durations)-1], ":")
//}
