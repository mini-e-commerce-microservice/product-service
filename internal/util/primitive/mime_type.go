package primitive

type MimeType string

const (
	MimeTypePng  MimeType = "image/png"
	MimeTypeJpeg MimeType = "image/jpeg"
	MimeTypeJpg  MimeType = "image/jpg"
	MimeTypeGif  MimeType = "image/gif"
	MimeTypeMP4  MimeType = "video/mp4"
)

func (m MimeType) Extension() string {
	switch m {
	case MimeTypePng:
		return ".png"
	case MimeTypeJpeg:
		return "/jpeg"
	case MimeTypeJpg:
		return "/jpg"
	case MimeTypeGif:
		return ".gif"
	case MimeTypeMP4:
		return ".mp4"
	default:
		return ""
	}
}
func (m MimeType) IsValid() bool {
	switch m {
	case MimeTypePng, MimeTypeJpeg, MimeTypeGif, MimeTypeMP4, MimeTypeJpg:
		return true
	default:
		return false
	}
}

func (m MimeType) MediaType() string {
	switch m {
	case MimeTypePng, MimeTypeJpeg, MimeTypeGif, MimeTypeJpg:
		return "image"
	case MimeTypeMP4:
		return "video/mp4"
	default:
		return ""
	}
}

func (m MimeType) MediaMaxSize() float64 {
	switch m {
	case MimeTypePng, MimeTypeJpeg, MimeTypeGif, MimeTypeJpg:
		return 1024
	case MimeTypeMP4:
		return 4092
	default:
		return 0
	}
}
