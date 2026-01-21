package ui_embed

import "embed"

var (
	//go:embed all:*
	EmbeddedUI embed.FS
)
