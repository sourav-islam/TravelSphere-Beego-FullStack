package models

// API Response Wrappers for v5 format
type APIResponse struct {
	Data APIData `json:"data"`
}

type APIData struct {
	Objects []Country `json:"objects"`
	Meta    APIMeta   `json:"meta"`
}

type APIMeta struct {
	Total  int  `json:"total"`
	Count  int  `json:"count"`
	Limit  int  `json:"limit"`
	Offset int  `json:"offset"`
	More   bool `json:"more"`
}

// Country model matches RestCountries v5 API response
type Country struct {
	Names      Names       `json:"names"`
	Capitals   []Capital   `json:"capitals"`
	Population int64       `json:"population"`
	Region     string      `json:"region"`
	Subregion  string      `json:"subregion"`
	Flag       Flag        `json:"flag"`
	Languages  []Language  `json:"languages"`
	Currencies []Currency  `json:"currencies"`
	Codes      Codes       `json:"codes"`
}

// Names struct for nested name fields
type Names struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

// Capital struct with location info
type Capital struct {
	Name string `json:"name"`
}

// Flag struct for v5 format (multiple formats available)
type Flag struct {
	Emoji       string `json:"emoji"`
	URLPng      string `json:"url_png"`
	URLSvg      string `json:"url_svg"`
	Description string `json:"description"`
}

// Language struct with more details
type Language struct {
	Name string `json:"name"`
}

// Currency struct
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Codes struct for ISO codes
type Codes struct {
	Alpha2 string `json:"alpha_2"`
	Alpha3 string `json:"alpha_3"`
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
