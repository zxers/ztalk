// +build wireinject

package main

import (
	"ztalk/cronjob/ioc"
	"ztalk/cronjob/repository"
	"ztalk/cronjob/repository/dao"

	"github.com/google/wire"
)

func InitCronJob() *Scheduler {
	wire.Build(
		NewScheduler,
		NewPreemptor,
		repository.NewPreemptRepository,
		dao.NewPreemptDao,
		ioc.InitDB,
	)
	return &Scheduler{}
}