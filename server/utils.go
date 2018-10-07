package server

var SupportedVideoFormats = []string{
	"flac",
	"mp4",
	"m4a",
	"mp3",
	"ogv",
	"ogm",
	"ogg",
	"oga",
	"opus",
	"webm",
	"wav",
}

func IsExtensionPlayable(ext string) bool {
	for _, format := range SupportedVideoFormats {
		if format == ext {
			return true
		}
	}
	return false
}
