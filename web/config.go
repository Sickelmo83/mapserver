package web

import (
	"encoding/json"
	"mapserver/app"
	"mapserver/layer"
	"net/http"
)

//Public facing config
type PublicConfig struct {
	Version         string               `json:"version"`
	Layers          []*layer.Layer       `json:"layers"`
	MapObjects      *app.MapObjectConfig `json:"mapobjects"`
	DefaultOverlays []string             `json:"defaultoverlays"`
	EnableSearch    bool                 `json:"enablesearch"`
}

type ConfigHandler struct {
	ctx *app.App
}

func (h *ConfigHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("content-type", "application/json")

	webcfg := PublicConfig{}
	webcfg.Layers = h.ctx.Config.Layers
	webcfg.MapObjects = h.ctx.Config.MapObjects
	webcfg.Version = app.Version
	webcfg.DefaultOverlays = h.ctx.Config.DefaultOverlays
	webcfg.EnableSearch = h.ctx.Config.EnableSearch

	json.NewEncoder(resp).Encode(webcfg)
}
