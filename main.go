package main

import (
	"github.com/ResourceAPI/Flatfile-JSON/storage"
	"github.com/ResourceAPI/Interface/plugins"
)

type FlatfileJSONPlugin string

func (FlatfileJSONPlugin) Name() string {
	return "Flatfile-JSON"
}

func (FlatfileJSONPlugin) Entrypoint() {
	plugins.GetRegistry().RegisterStorage("Flatfile-JSON", &storage.FlatfileJSONStorage{})
}

var CorePlugin FlatfileJSONPlugin
