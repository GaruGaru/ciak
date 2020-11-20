package models

import (
	"errors"
	"fmt"
)

var (
	ErrMediaFormatNotSupported = errors.New("media format not supported")
)

type MediaFormat int

const (
	MediaFormatAvi MediaFormat = iota
	MediaFormatMkv
	MediaFormatFlac
	MediaFormatMp4
	MediaFormatMp4a
	MediaFormatMp3
	MediaFormatOgv
	MediaFormatOgm
	MediaFormatOgg
	MediaFormatOga
	MediaFormatOpus
	MediaFormatWebm
	MediaFormatWav
)

var SupportedMediaFormat = []MediaFormat{
	MediaFormatAvi,
	MediaFormatMkv,
	MediaFormatFlac,
	MediaFormatMp4,
	MediaFormatMp4a,
	MediaFormatMp3,
	MediaFormatOgv,
	MediaFormatOgm,
	MediaFormatOgg,
	MediaFormatOga,
	MediaFormatOpus,
	MediaFormatWebm,
	MediaFormatWav,
}

func MediaFormatFrom(raw string) (MediaFormat, error) {
	for _, format := range SupportedMediaFormat {
		if format.Name() == raw || format.Extension() == raw {
			return format, nil
		}
	}

	return MediaFormat(-1), ErrMediaFormatNotSupported
}

func (d MediaFormat) Extension() string {
	return fmt.Sprintf(".%s", d.Name())
}

func (d MediaFormat) Name() string {
	return [...]string{
		"avi", "mkv", "flac", "mp4", "m4a", "mp3", "ogv",
		"ogm", "ogg", "oga", "opus", "webm", "wav",
	}[d]
}

