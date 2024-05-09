package dao

import (
	"fmt"
	"time"
	"ztalk/cronjob/domain"

	"gorm.io/gorm"
)

type PreemptDao interface {
	AddJob(job domain.Job) error
	Preempt() domain.Job
	UpdateNextTime(job domain.Job) error
	Release(job domain.Job) error
}

type preemptDao struct {
	db *gorm.DB
}

func NewPreemptDao(db *gorm.DB) PreemptDao {
	return &preemptDao{
		db: db,
	}
}

func (s *preemptDao) AddJob(domainJob domain.Job) error {
	job := &Job{
		Name: domainJob.Name,
		Spec: domainJob.Spec,
		Status: JobWaiting,
	}
	now := time.Now()
	fmt.Println("start Time:", now)
	var err error
	job.NextTime, err = domainJob.Next(now)
	if err != nil {
		return nil
	}
	result := s.db.Create(job)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *preemptDao) Preempt() domain.Job {
	for {
		now := time.Now()
		job := Job{}
		result := p.db.Where("next_time < ?", now.UnixMilli()).First(&job)
		if result.Error != nil {
			return domain.Job{}
		}
		p.db.Model(&job).Updates(Job{Status: JobRunning})
		return domain.Job{
			Name: job.Name,
			Spec: job.Spec,
		}
	}
}

func (p *preemptDao) UpdateNextTime(job domain.Job) error {
	now := time.Now()
	
	nextTime, err := job.Next(now)
	fmt.Println("nextTime:", nextTime, job)
	if err != nil {
		return err
	}
	result := p.db.Model(&Job{}).Where("name = ?", job.Name).Updates(Job{NextTime: nextTime})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *preemptDao) Release(job domain.Job) error {
	result := p.db.Model(&Job{Name: job.Name}).Updates(Job{Status: JobWaiting})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type Job struct {
	Id       int64  `gorm:"primaryKey"`
	Name     string `gorm:"index:idx_name,unique"`
	Status   int
	Spec     string
	NextTime int64
}

const (
	JobWaiting = iota
	JobRunning
)
