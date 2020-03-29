package svcdef

import (
	"plugin"

	log "github.com/sirupsen/logrus"
)

// TemplateFile represents a file of the template
// pointed by T
type TemplateFile struct {
	Path string
	T    *Template
}

// Template is a template descriptor holding
// its representation as a Go Plugin
type Template struct {
	Plugin *plugin.Plugin
}

// Load loads the template under p as a Go Plugin
func (t *Template) Load(p string) {
	var err error

	t.Plugin, err = plugin.Open(p)
	if err != nil {
		log.Fatal(err)
	}
}

// Asset Asset() wrapper for configured Go Plugin
func (t *Template) Asset(n string) ([]byte, error) {
	assetSym, err := t.Plugin.Lookup("Asset")
	if err != nil {
		log.Fatal(err)
	}

	asset, ok := assetSym.(func(string) ([]byte, error))
	if !ok {
		log.Fatal(ok)
	}

	b, err := asset(n)

	return b, err
}

// AssetNames AssetNames() wrapper for configured Go Plugin
func (t *Template) AssetNames() []string {
	assetNamesSym, err := t.Plugin.Lookup("AssetNames")
	if err != nil {
		log.Fatal(err)
	}

	assetNames, ok := assetNamesSym.(func() []string)
	if !ok {
		log.Fatal(assetNamesSym)
	}

	return assetNames()
}
