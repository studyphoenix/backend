package controller

type MediaTypeExtensionMap map[string]string

// TODO: support more media types
var MEDIA_TYPE_MAP = map[string]string{
	"application/pdf":  "pdf",
	"application/json": "json",
	"image/jpg":        "jpg",
	"image/jpeg":       "jpg",
	"image/png":        "png",
}
