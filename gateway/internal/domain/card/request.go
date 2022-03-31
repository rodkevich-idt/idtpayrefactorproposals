package card

import (
	"encoding/json"
	"io"

	"gateway/internal/models"
)

type Request struct {
	Filter
	// CardID     		string `json:"-"`
	// FieldName        string `json:"field_name" validate:"required"`
}

func (r *Request) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func ToCard(req *Request) *models.Card {
	return &models.Card{
		// {{ FieldName }}:         req.{{ req.fieldName }},
	}
}
