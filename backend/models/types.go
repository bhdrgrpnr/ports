package models

type Port struct {
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Province    string    `json:"province"`
	Country     string    `json:"country"`
	Alias       []string  `json:"alias"`
	Regions     []string  `json:"regions"`
	Coordinates []float64 `json:"coordinates"`
	TimeZone    string    `json:"timezone"`
	Unlocks     []string  `json:"Unlocks"`
	Code        string    `json:"code"`
}

type PortRequest struct {
	Index string `json:"index"`
	Port  Port   `json:"port"`
}
