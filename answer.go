package goform

type Answer struct {
	Type  string `json:"type"`
	Text  string `json:"text,omitempty"`
	Field struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"field"`
	Email   string `json:"email,omitempty"`
	Date    string `json:"date,omitempty"`
	Choices struct {
		Labels []string `json:"labels"`
	} `json:"choices,omitempty"`
	Number  int  `json:"number,omitempty"`
	Boolean bool `json:"boolean,omitempty"`
	Choice  struct {
		Label string `json:"label"`
	} `json:"choice,omitempty"`
	FileURL string `json:"file_url"`
}

func (a Answer) MatchesID(needle string) bool {
	return needle == a.Field.ID
}

func (a Answer) Value() interface{} {
	switch a.Type {
	case "text":
		return a.Text
	case "email":
		return a.Email
	case "date":
		return a.Date
	case "choices":
		return a.Choices.Labels
	case "boolean":
		return a.Boolean
	case "number":
		return a.Number
	case "choice":
		return a.Choice.Label
	case "file_url":
		return a.FileURL
	}
	return nil
}
