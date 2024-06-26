// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"ztalk/cronjob/ioc"
	"ztalk/cronjob/repository"
	"ztalk/cronjob/repository/dao"
)

// Injectors from wire.go:

func InitCronJob() *Scheduler {
	db := ioc.InitDB()
	preemptDao := dao.NewPreemptDao(db)
	preemptRepository := repository.NewPreemptRepository(preemptDao)
	mainPreemptor := NewPreemptor(preemptRepository)
	scheduler := NewScheduler(mainPreemptor)
	return scheduler
}
