package tools

// type of c2ffi output
type Ctype struct {
	Tag  string `json:"tag"`
	Type Type   `json:"type"`
}
type Cparam struct {
	Tag  string `json:"tag"`
	Name string `json:"name"`
	Type Ctype  `json:"type"`
}
type CFunction struct {
	Name       string   `json:"name"`
	Parameters []Cparam `json:"parameters"`
	Location   string   `json:"location"`
	ReturnType Type     `json:"return-type"`
	Tag        string   `json:"tag"`
}
type Type struct {
	Tag  string `json:"tag"`
	Type string `json:"type"`
}
