package job

import (
	"errors"

	"github.com/arybolovlev/orkestrator/task"
)

type Job struct {
	ID    string
	Name  string
	Tasks []task.Task
}

func NewJob(n string, t []task.Task) *Job {
	return &Job{
		Name:  n,
		Tasks: t,
	}
}

func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("Job Name must be set")
	}

	if len(j.Tasks) == 0 {
		return errors.New("At least one Task must be set")
	}

	for _, t := range j.Tasks {
		if err := t.Validate(); err != nil {
			return err
		}
	}

	return nil
}
