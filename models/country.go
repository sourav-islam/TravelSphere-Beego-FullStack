package models

type Country struct {
	Name       CommonName          `json:"name"`
	Capital    []string            `json:"capital"`
	Population int64               `json:"population"`
	Region     string              `json:"region"`
	Subregion  string              `json:"subregion"`
	Flags      Flags               `json:"flags"`
	Languages  map[string]string   `json:"languages"`
	Currencies map[string]Currency `json:"currencies"`
	Cca2       string              `json:"cca2"`
	Cca3       string              `json:"cca3"`
}

type CommonName struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type Flags struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
	Alt string `json:"alt"`
}

type Currency struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// DTO for templates and API responses
type CountryDTO struct {
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Capital    string `json:"capital"`
	Population string `json:"population"`
	Region     string `json:"region"`
	Subregion  string `json:"subregion"`
	FlagURL    string `json:"flag_url"`
	FlagAlt    string `json:"flag_alt"`
	Languages  string `json:"languages"`
	Currencies string `json:"currencies"`
	Cca2       string `json:"cca2"`
}
