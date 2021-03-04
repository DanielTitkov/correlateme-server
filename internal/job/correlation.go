package job

import (
	"fmt"

	"github.com/DanielTitkov/correlateme-server/internal/service/metrics"
)

func (j *Job) ListenUpdateUserCorrelationsChannel() {
	for args := range j.app.Channels.UpdateUserCorrelationsChan {
		msg := fmt.Sprintf("user correlation with args: %+v", args)
		j.logger.Info("got request to update correlations", msg)
		err := j.app.UpdateCorrelations(args)
		metrics.UnprocessedUpdateCorrelationsRequests.Add(-1)
		if err != nil {
			j.logger.Error("failed to update "+msg, err)
			continue
		}
		j.logger.Info("updated correlations", msg)
	}
}
