package app

import (
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

func (a *App) FilterIndicators(filter domain.FilterIndicatorsArgs) ([]*domain.Indicator, error) {
	// if user want to get built-in indicators author is obsolete
	// otherwise user gets only indicators belonging to them
	if filter.BuiltIn != nil && *filter.BuiltIn == true {
		filter.AuthorUsername = nil
	}

	return a.repo.FilterIndicators(filter)
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
