package main

import (
	"github.com/ResourceAPI/Core/plugins"
	"github.com/ResourceAPI/Flatfile-JSON/storage"
)

type FlatfileJSONPlugin string

func (FlatfileJSONPlugin) Name() string {
	return "Flatfile-JSON"
}

func (FlatfileJSONPlugin) Entrypoint() {
	plugins.RegisterStorage("Flatfile-JSON", &storage.FlatfileJSONStorage{})
}

var CorePlugin FlatfileJSONPlugin
