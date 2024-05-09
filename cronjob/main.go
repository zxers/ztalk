package main

import (
	"fmt"
	"ztalk/cronjob/domain"
)

func main() {
	s := InitCronJob()
	err := s.AddJob(domain.Job{
		Name:"ranking",
		Exec: func() {
			fmt.Println("hello")
		},
		Spec: "*/10 * * * * *",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Add Succeed")
	s.Start()
}