package goform

import "time"

type Form struct {
	FormID      string     `json:"form_id"`
	Token       string     `json:"token"`
	SubmittedAt time.Time  `json:"submitted_at"`
	LandedAt    time.Time  `json:"landed_at"`
	Definition  Definition `json:"definition"`
	Answers     []Answer   `json:"answers"`
}

func (f Form) MapAnswersByName() map[string]Answer {
	idMap := make(map[string]string)
	output := make(map[string]Answer)

	for _, dfx := range f.Definition.Fields {
		idMap[dfx.ID] = dfx.Title
	}
	for _, ax := range f.Answers {
		output[idMap[ax.Field.ID]] = ax
	}

	return output
}
