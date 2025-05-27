package models

type Status struct {
	Version string `json:"version"`
	Health  int    `json:"health"`
}
