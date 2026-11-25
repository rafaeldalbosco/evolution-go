package utils

import (
	"testing"

	"go.mau.fi/whatsmeow"
)

func TestDocumentFallbackMediaType(t *testing.T) {
	cases := map[string]whatsmeow.MediaType{
		"image/jpeg":      whatsmeow.MediaImage,
		"video/mp4":       whatsmeow.MediaVideo,
		"audio/ogg":       whatsmeow.MediaAudio,
		"application/pdf": "",
		"":                "",
	}
	for mime, want := range cases {
		if got := documentFallbackMediaType(mime); got != want {
			t.Errorf("%q: got %q, want %q", mime, got, want)
		}
	}
}
