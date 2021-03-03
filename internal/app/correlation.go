package app

import (
	"fmt"
	"math"
	"time"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/dgryski/go-onlinestats"
	combinations "github.com/mxschmitt/golang-combinations"
)

func (a *App) FindCorrelations(args domain.FindCorrelationsArgs) error {
	fmt.Println("ARGS", args) // FIXME

	datasets, err := a.repo.GetUserDatasets(args.UserID, true, args.WithShared)
	if err != nil {
		return err
	}

	// map datasets
	// TODO: create combination method for ints to remove conversion
	datasetMap := mapDatasets(datasets)

	// get dataset ids
	var datasetIDs []string
	for k := range datasetMap {
		datasetIDs = append(datasetIDs, k)
	}

	// make combinations
	subsets := combinations.Combinations(datasetIDs, 2)
	fmt.Println("SUBSETS", subsets) // FIXME

	// make data slices
	for _, subset := range subsets {
		left := datasetMap[subset[0]]
		right := datasetMap[subset[1]]
		leftObs := mapObservations(left.Observations)
		rightObs := mapObservations(right.Observations)

		var leftData, rightData []float64
		for date, leftOb := range leftObs {
			rightOb, ok := rightObs[date]
			if !ok {
				continue
			}
			leftData = append(leftData, leftOb.Value)
			rightData = append(rightData, rightOb.Value)
		}
		if len(leftData) == 0 { // TODO: check minimal observation count
			continue
		}

		fmt.Println("DATA", leftData, rightData) // FIXME
		// find coef and p
		coef, p := onlinestats.Spearman(leftData, rightData)
		if math.IsNaN(coef) {
			continue
		}
		correlation := domain.Correlation{
			Left:  left,
			Right: right,
			Coef:  coef,
			P:     p,
			Type:  domain.SpearmanCorrelationType,
			R2:    math.Pow(coef, 2),
		}
		fmt.Println("CORRELATION", correlation) // FIXME
		// save correlations to db
		_, err = a.repo.CreateOrUpdateCorrelation(&correlation)
		if err != nil { // TODO: maybe save error, not exit right away
			return err
		}
	}

	return nil
}

func mapObservations(obs []*domain.Observation) map[time.Time]*domain.Observation {
	res := make(map[time.Time]*domain.Observation)
	for _, ob := range obs {
		res[*ob.Date] = ob
	}
	return res
}

func mapDatasets(datasets []*domain.Dataset) map[string]*domain.Dataset {
	datasetMap := make(map[string]*domain.Dataset)
	for _, dataset := range datasets {
		if len(dataset.Observations) == 0 {
			continue
		}
		datasetMap[fmt.Sprintf("%d", dataset.ID)] = dataset
	}
	return datasetMap
}
