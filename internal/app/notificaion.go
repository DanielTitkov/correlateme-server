package app

func (a *App) Notify() error {
	// get all unprocess notifications
	// process them with notifier
	// depending on config do or do not set processed = true
	// on relevant anomalies
	return nil
}

func (a *App) ScheduleNotifications() error {
	err := a.cron.AddFunc(a.cfg.Notification.Schedule, func() {
		err := a.Notify()
		if err != nil {
			a.logger.Error("failed sending notifications", err)
		}
		a.logger.Info("notifications sent", "")
	})
	return err
}
