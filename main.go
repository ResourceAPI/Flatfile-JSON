package main

import (
	"github.com/StratoAPI/Flatfile-JSON/config"
	"github.com/StratoAPI/Flatfile-JSON/storage"
	"github.com/StratoAPI/Interface/plugins"
)

type FlatfileJSONPlugin string

func (FlatfileJSONPlugin) Name() string {
	return "Flatfile-JSON"
}

func (FlatfileJSONPlugin) Entrypoint() {
	plugins.GetRegistry().RegisterStorage("Flatfile-JSON", &storage.FlatfileJSONStorage{})
	plugins.GetRegistry().RegisterConfig("flatfile-json", config.Get())
	plugins.GetRegistry().AssociateFilter("simple", "Flatfile-JSON")
}

var CorePlugin FlatfileJSONPlugin
