package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

func (a *App) GetScaleByType(scaleType string) (*domain.Scale, error) {
	scale, ok := a.cache.scales[scaleType]
	if !ok {
		return nil, fmt.Errorf("got unknown scale type '%s'", scaleType)
	}

	return scale, nil
}

func (a *App) initScales() error {
	var scales []domain.Scale

	data, err := ioutil.ReadFile(a.cfg.Data.Presets.ScalePresetsPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &scales)
	if err != nil {
		return err
	}

	for _, scale := range scales {
		s, err := a.repo.GetScaleByType(scale.Type)
		if err == nil && s.ID != 0 {
			a.logger.Info("scale already exists", fmt.Sprintf("%+v", s))
			continue
		}

		s, err = a.repo.CreateScale(scale)
		if err != nil {
			return err
		}
		a.logger.Info("created scale", fmt.Sprintf("%+v", s))
	}

	return nil
}
