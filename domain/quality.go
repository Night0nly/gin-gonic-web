package domain

type Quality int

const (
	LOSSLESS Quality = iota
	HD1080P
	HD720P
	KBPS320
	KBPS192
	KBPS128
)

func (q Quality) String() string {
	switch q {
	case LOSSLESS:
		return "Lossless"
	case HD1080P:
		return "HD1080p"
	case HD720P:
		return "HD720p"
	case KBPS320:
		return "320kbps"
	case KBPS192:
		return "192kbps"
	case KBPS128:
		return "128kbps"
	}
	return "Unknown"
}