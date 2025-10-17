package pokeapi

// RespShallowLocations -
type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}


type ExploreLocation struct {
	GameIndex int `json:"game_index"`
	ID int `json:"id"`
	Location Location `json:"location"`
	Name string `json:"name"`
	Names []Names `json:"names"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
}
type Location struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Language struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Names struct {
	Name string `json:"name"`
	Language Language `json:"language"`
}
type Pokemon struct {
	URL string `json:"url"`
	Name string `json:"name"`
}
type Version struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type Method struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type EncounterDetails struct {
	MinLevel int `json:"min_level"`
	Chance int `json:"chance"`
	ConditionValues []any `json:"condition_values"`
	MaxLevel int `json:"max_level"`
	Method Method `json:"method"`
}
type VersionDetails struct {
	MaxChance int `json:"max_chance"`
	Version Version `json:"version"`
	EncounterDetails []EncounterDetails `json:"encounter_details"`
}
type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}
type EncounterMethodRatesVersionDetailsVersion struct {
	Name string `json:"name"`
	URL string `json:"url"`
}
type EncounterMethodRatesVersionDetails struct {
	EncounterMethodRatesVersionDetailsVersion EncounterMethodRatesVersionDetailsVersion `json:"version"`
	Rate int `json:"rate"`
}
type EncounterMethod struct {
	URL string `json:"url"`
	Name string `json:"name"`
}
type EncounterMethodRates struct {
	EncounterMethodRatesVersionDetails []EncounterMethodRatesVersionDetails `json:"version_details"`
	EncounterMethod EncounterMethod `json:"encounter_method"`
}

// PokemonDetails represents a Pokemon from the API
type PokemonDetails struct {
	Name           string      `json:"name"`
	BaseExperience int         `json:"base_experience"`
	Height         int         `json:"height"`
	Weight         int         `json:"weight"`
	ID             int         `json:"id"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type PokemonType struct {
	Slot int `json:"slot"`
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}