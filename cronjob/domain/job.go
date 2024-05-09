package domain

import (
	"time"

	"github.com/robfig/cron/v3"
)

type Exector func()

type Job struct {
	Name string
	Exec Exector
	Spec string
}

func (j *Job) Next(time time.Time) (int64, error) {
	parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	sche, err := parser.Parse(j.Spec)
	if err != nil {
		return 0, err
	}
	return sche.Next(time).UnixMilli(), nil
}