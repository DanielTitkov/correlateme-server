package app

import (
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

func makeIndicatorCode(username, scaleType, title string) string {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	code := strings.Join([]string{
		strings.ToLower(username),
		scaleType,
		strings.ToLower(title), // TODO: strip non-letter
		ts,
	}, "-")
	return code
}
