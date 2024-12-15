package boot

import (
	"embed"

	"github.com/kolukattai/kurl/models"
)

var (
	Config         *models.Config
	StaticFolder   embed.FS
	TemplateFolder embed.FS
)
