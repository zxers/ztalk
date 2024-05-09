package main

import (
	"fmt"
	"ztalk/cronjob/domain"
)

type Scheduler struct {
	Preemptor Preemptor
	Exector map[string]domain.Exector
}

func NewScheduler(preemptor Preemptor) *Scheduler {
	return &Scheduler{
		Preemptor: preemptor,
		Exector: make(map[string]domain.Exector),
	}
}

func (s *Scheduler) AddJob(job domain.Job) error {
	err := s.Preemptor.AddJob(job)
	if err != nil {
		return err
	}
	s.Exector[job.Name] = job.Exec
	return nil
}

func (s *Scheduler) UpdateNextTime(job domain.Job) error {
	err := s.Preemptor.UpdateNextTime(job)
	if err != nil {
		return err
	}
	return nil
}

func (s *Scheduler) Start() {
	for {
		job := s.Preemptor.Preempt()
		if job.Name == "" {
			continue
		}
		fmt.Println("Preempt Succeed")
		go func (job domain.Job) {
			s.Exector[job.Name]()
			s.Preemptor.Release(job)
		}(job)
		s.UpdateNextTime(job)
	}
}