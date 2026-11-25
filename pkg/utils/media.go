package utils

import (
	"context"
	"errors"
	"strings"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/proto/waE2E"
)

// DownloadDocument baixa um documento. Alguns clientes enviam foto/vídeo/áudio "como documento"
// (documentWithCaptionMessage) cifrados com as chaves do tipo real: o hash do blob confere, mas o
// HMAC com "WhatsApp Document Keys" falha. Nesse caso tenta de novo com o tipo do mimetype.
func DownloadDocument(ctx context.Context, cli *whatsmeow.Client, doc *waE2E.DocumentMessage) ([]byte, error) {
	data, err := cli.Download(ctx, doc)
	if !errors.Is(err, whatsmeow.ErrInvalidMediaHMAC) {
		return data, err
	}
	mediaType := documentFallbackMediaType(doc.GetMimetype())
	if mediaType == "" {
		return nil, err
	}
	return cli.DownloadMediaWithPath(ctx, doc.GetDirectPath(), doc.GetFileEncSHA256(), doc.GetFileSHA256(),
		doc.GetMediaKey(), mediaType, "document", false)
}

func documentFallbackMediaType(mimetype string) whatsmeow.MediaType {
	switch strings.SplitN(mimetype, "/", 2)[0] {
	case "image":
		return whatsmeow.MediaImage
	case "video":
		return whatsmeow.MediaVideo
	case "audio":
		return whatsmeow.MediaAudio
	}
	return ""
}
