package model

type Character struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Summary string  `json:"summary"`
	Skills  []Skill `json:"skills"`
}

type Skill struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}
