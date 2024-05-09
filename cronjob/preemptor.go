package main

import (
	"ztalk/cronjob/domain"
	"ztalk/cronjob/repository"
)

type Preemptor interface {
	AddJob(job domain.Job) error
	Preempt() domain.Job
	UpdateNextTime(job domain.Job) error
	Release(job domain.Job) error
}

type preemptor struct {
	repo repository.PreemptRepository
}

func NewPreemptor(repo repository.PreemptRepository) Preemptor {
	return &preemptor{
		repo: repo,
	}
}

func (p *preemptor) AddJob(job domain.Job)error {
	return p.repo.AddJob(job)
}

func (p *preemptor) Preempt() domain.Job {
	return p.repo.Preempt()
}

func (p *preemptor) UpdateNextTime(job domain.Job) error {
	err := p.repo.UpdateNextTime(job)
	if err != nil {
		return err
	}
	return nil
}

func (p *preemptor) Release(job domain.Job) error {
	return p.repo.Release(job)
}