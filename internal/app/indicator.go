package app

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
)

func (a *App) CreateIndicator(args domain.CreateIndicatorArgs) error {
	user, err := a.repo.GetUserByUsername(args.Username)
	if err != nil {
		return err
	}

	scale, err := a.GetScaleByType(args.ScaleType)
	if err != nil {
		return err
	}

	_, err = a.repo.CreateIndicator(&domain.Indicator{
		Code:        makeIndicatorCode(args.Username, args.ScaleType, args.Title),
		Title:       args.Title,
		Description: args.Description,
		Scale:       scale,
		Author:      user,
		Active:      true,
	})
	if err != nil {
		return err
	}

	return nil
}

func (a *App) CreateIndicatorFromPreset(ind domain.Indicator) (*domain.Indicator, error) {
	scale, err := a.GetScaleByType(ind.Scale.Type)
	if err != nil {
		return nil, err
	}

	indicator, err := a.repo.CreateIndicator(&domain.Indicator{
		Code:        ind.Code,
		Title:       ind.Title,
		Description: ind.Description,
		Scale:       scale,
		BuiltIn:     true,
		Active:      true,
	})
	if err != nil {
		return nil, err
	}

	return indicator, nil
}

func (a *App) GetIndicators(args domain.GetIndicatorsArgs) ([]*domain.Indicator, error) {
	// if user wants to get built-in indicators author is obsolete
	// otherwise user gets only indicators belonging to them
	if args.Filter.BuiltIn != nil && *args.Filter.BuiltIn {
		args.Filter.AuthorID = nil
	}
	if args.ObservationLimit > a.cfg.App.DefaultObservationLimit {
		args.ObservationLimit = a.cfg.App.DefaultObservationLimit
	}

	inds, err := a.repo.GetIndicators(args)
	if err != nil {
		return nil, err
	}

	for _, ind := range inds {
		ind.UserDataset.Observations = orderObservationsAsc(ind.UserDataset.Observations)
	}

	return inds, nil
}

func makeIndicatorCode(username, scaleType, title string) string {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	reg, _ := regexp.Compile(`[^\d\w]+`)
	code := strings.Join([]string{
		strings.ToLower(username),
		scaleType,
		reg.ReplaceAllString(strings.ToLower(title), ""), // TODO: what about cyrillic?
		ts,
	}, "-")
	return code
}

func (a *App) initBuiltinIndicators() error {
	data, err := ioutil.ReadFile(a.cfg.Data.Presets.IndicatorPresetsPath)
	if err != nil {
		return err
	}

	var indicators []domain.Indicator
	err = json.Unmarshal(data, &indicators)
	if err != nil {
		return err
	}

	for _, indicator := range indicators {
		ind, err := a.repo.GetIndicatorByCode(indicator.Code)
		if err == nil && ind.ID != 0 {
			a.logger.Info("indicator already exists", ind.JSONString())
			continue
		}

		ind, err = a.CreateIndicatorFromPreset(indicator)
		if err != nil {
			return err
		}

		a.logger.Info("created indicator", ind.JSONString())
	}

	return nil
}
