package domain

type DownloadURLList map[Quality]string

func EmptyDownloadURLList() *DownloadURLList {
	return &DownloadURLList{}
}

func (d *DownloadURLList) Add(quality Quality, url string) {
	(*d)[quality] = url
}