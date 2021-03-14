package job

import (
	"fmt"

	"github.com/DanielTitkov/correlateme-server/internal/domain"
	"github.com/DanielTitkov/correlateme-server/internal/service/metrics"
)

func (j *Job) ListenUpdateDatasetAggregationsChannel() {
	for args := range j.app.Channels.UpdateDatasetAggregationsChan {
		msg := fmt.Sprintf("dataset aggragations with args: %+v", args)
		j.logger.Info("got request to update aggregations", msg)
		err := j.app.UpdateAggregations(args)
		metrics.UnprocessedUpdateAggregationsRequests.Add(-1)
		if err != nil {
			j.logger.Error("failed to update "+msg, err)
			continue
		}
		j.logger.Info("updated aggregations", msg)
		go func(args domain.UpdateAggregationsArgs) {
			// update correlations for all user datasets with granularity == week
			metrics.UnprocessedUpdateCorrelationsRequests.Add(1)
			j.app.Channels.UpdateUserCorrelationsChan <- domain.UpdateCorrelationsArgs{
				UserID:      args.UserID,
				WithShared:  true,
				Method:      "auto",
				Granularity: domain.GranularityWeek,
			}
			// update correlations for all user datasets with granularity == month
			metrics.UnprocessedUpdateCorrelationsRequests.Add(1)
			j.app.Channels.UpdateUserCorrelationsChan <- domain.UpdateCorrelationsArgs{
				UserID:      args.UserID,
				WithShared:  true,
				Method:      "auto",
				Granularity: domain.GranularityMonth,
			}
		}(args)
	}
}
