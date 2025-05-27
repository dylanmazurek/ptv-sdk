package models

type Disruption struct {
	ID             int    `json:"disruption_id"`
	Title          string `json:"title"`
	URL            string `json:"url"`
	Description    string `json:"description"`
	Status         string `json:"disruption_status"`
	DisruptionType string `json:"disruption_type"`
	PublishedOn    string `json:"published_on"`
	LastUpdated    string `json:"last_updated"`
	FromDate       string `json:"from_date"`
	ToDate         string `json:"to_date"`
	Colour         string `json:"colour"`
	DisplayOnBoard bool   `json:"display_on_board"`
	DisplayStatus  bool   `json:"display_status"`
}
