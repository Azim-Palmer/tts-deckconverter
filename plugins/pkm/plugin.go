package pkm

import (
	"github.com/jeandeaual/tts-deckconverter/plugins"
)

type imageQuality string

const (
	normal imageQuality = "normal"
	hires  imageQuality = "hires"
)

type pokemonPlugin struct {
	id   string
	name string
}

func (p pokemonPlugin) PluginID() string {
	return p.id
}

func (p pokemonPlugin) PluginName() string {
	return p.name
}

func (p pokemonPlugin) AvailableOptions() plugins.Options {
	return plugins.Options{
		"quality": plugins.Option{
			Type:        plugins.OptionTypeEnum,
			Description: "image quality",
			AllowedValues: []string{
				string(normal),
				string(hires),
			},
			DefaultValue: string(hires),
		},
	}
}

func (p pokemonPlugin) URLHandlers() []plugins.URLHandler {
	return []plugins.URLHandler{}
}

func (p pokemonPlugin) FileExtHandlers() map[string]plugins.FileHandler {
	return map[string]plugins.FileHandler{
		".ptcgo": fromDeckFile,
	}
}

func (p pokemonPlugin) DeckTypeHandlers() map[string]plugins.FileHandler {
	return map[string]plugins.FileHandler{}
}

func (p pokemonPlugin) GenericFileHandler() plugins.FileHandler {
	return fromDeckFile
}

func (p pokemonPlugin) AvailableBacks() map[string]plugins.Back {
	return map[string]plugins.Back{
		plugins.DefaultBackKey: {
			URL:         defaultBackURL,
			Description: "standard paper card back",
		},
		"japanese": {
			URL:         japaneseBackURL,
			Description: "Japanese card back",
		},
		"japanese_old": {
			URL:         japaneseOldBackURL,
			Description: "Old Japanese card back",
		},
	}
}

// PokemonPlugin is the exported plugin for this package
var PokemonPlugin = pokemonPlugin{
	id:   "pkm",
	name: "Pokémon",
}
