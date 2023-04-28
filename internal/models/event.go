package models

type Event struct {
	ID         int    `json:"id"`
	Name       string `json:"nameParty"`
	Date       string `json:"dateStart"`
	Type       string `json:"nameType"`
	Country    string `json:"nameCountry"`
	City       string `json:"nameTown"`
	Organiser  string `json:"nameOrganizer"`
	URL        string `json:"urlOrganizer"`
	GoabaseURL string `json:"urlPartyHtml"`
}

type PartyList struct {
	PartyList []Event `json:"partylist"`
}
