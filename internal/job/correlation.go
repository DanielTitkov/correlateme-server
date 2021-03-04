package job

import "fmt"

func (j *Job) ListenUpdateUserCorrelationsChannel() {
	for args := range j.app.Channels.UpdateUserCorrelationsChan {
		msg := fmt.Sprintf("user correlation with args: %+v", args)
		j.logger.Info("got request to update correlations", msg)
		err := j.app.UpdateCorrelations(args)
		if err != nil {
			j.logger.Error("failed to update "+msg, err)
			continue
		}
		j.logger.Info("updated correlations", msg)
	}
}
