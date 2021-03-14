package job

import (
	"fmt"

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
	}
}
