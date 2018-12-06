package domain

type CollectorType int

const (
	CHIASENHAC CollectorType = iota
	MP3ZING
)

func (s CollectorType) String() string {
	switch s {
	case CHIASENHAC:
		return "CHIASENHAC"
	case MP3ZING:
		return "MP3ZING"
	}
	return "Unknown"
}
