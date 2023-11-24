package job

import (
	"errors"

	"github.com/arybolovlev/orkestrator/api/task"
)

type Job struct {
	ID    string      `hcl:"id,optional"`
	Name  string      `hcl:"name,label"`
	Tasks []task.Task `hcl:"task,block"`
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
