package goform

type Definition struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Fields []Field `json:"fields"`
}

type Field struct {
	ID                      string `json:"id"`
	Title                   string `json:"title"`
	Type                    string `json:"type"`
	Ref                     string `json:"ref"`
	AllowMultipleSelections bool   `json:"allow_multiple_selections"`
	AllowOtherChoice        bool   `json:"allow_other_choice"`
}

func (d Definition) GetFieldByName(name string) string {
	targetID := ""
	for _, field := range d.Fields {
		if field.Title == name {
			targetID = field.ID
		}
	}

	return targetID
}
