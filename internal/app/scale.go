package app

import (
	"fmt"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

func (a *App) GetScaleByType(scaleType string) (*domain.Scale, error) {
	scale, ok := a.cache.scales[scaleType]
	if !ok {
		return nil, fmt.Errorf("got unknown scale type '%s'", scaleType)
	}

	return scale, nil
}
