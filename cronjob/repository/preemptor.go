package repository

import (
	"ztalk/cronjob/domain"
	"ztalk/cronjob/repository/dao"
)

type PreemptRepository interface {
	AddJob(job domain.Job) error
	Preempt() domain.Job
	UpdateNextTime(job domain.Job) error
	Release(job domain.Job) error
}

type preemptRepository struct {
	dao dao.PreemptDao
}

func NewPreemptRepository(dao dao.PreemptDao) PreemptRepository {
	return &preemptRepository{
		dao: dao,
	}
}

func (s *preemptRepository) AddJob(job domain.Job) error {
	return s.dao.AddJob(job)
}

func (p *preemptRepository) Preempt() domain.Job {
	return p.dao.Preempt()
}

func (p *preemptRepository) UpdateNextTime(job domain.Job) error {
	err := p.dao.UpdateNextTime(job)
	if err != nil {
		return err
	}
	return nil
}

func (p *preemptRepository) Release(job domain.Job) error {
	return p.dao.Release(job)
}