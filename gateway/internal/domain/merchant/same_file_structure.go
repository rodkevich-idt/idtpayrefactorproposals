package card

import "gateway/internal/models"

type Res struct {
}

func Resource(card *models.Card) (Res, error) {
	var resource Res

	return resource, nil
}

func Resources(cards []*models.Card) ([]Res, error) {
	if len(cards) == 0 {
		return make([]Res, 0), nil
	}

	var resources []Res
	for _, card := range cards {
		res, _ := Resource(card)
		resources = append(resources, res)
	}
	return resources, nil
}
