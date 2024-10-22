package lang

import (
	"fmt"
	"sync"
)

type jobs struct {
	mutex sync.Mutex
	jobs  []*Process
}

func NewJobs() *jobs {
	return &jobs{}
}

func (j *jobs) Add(p *Process) {
	j.mutex.Lock()
	j.jobs = append(j.jobs, p)
	j.mutex.Unlock()
}

func (j *jobs) GarbageCollect() {
	j.mutex.Lock()

	var (
		last    = -1
		running bool
	)

	for i := len(j.jobs) - 1; i >= 0; i-- {
		if j.jobs[i] != nil {
			if j.jobs[i].HasTerminated() {
				j.jobs[i] = nil
			} else {
				running = true
			}
		}

		if j.jobs[i] == nil && !running {
			last = i
		}
	}

	switch last {
	case -1:
		break
	/*case 0:
	j.jobs = nil*/
	default:
		j.jobs = j.jobs[:last]
	}

	j.mutex.Unlock()
}

func (j *jobs) Get(jobId int) (*Process, error) {
	j.mutex.Lock()
	defer j.mutex.Unlock()

	if jobId < 1 {
		return nil, fmt.Errorf("invalid job ID '%d': job ID cannot be less than 1", jobId)
	}

	if jobId > len(j.jobs) {
		return nil, fmt.Errorf("invalid job ID '%d': job ID was greater than number of jobs (%d)", jobId, len(j.jobs))
	}

	i := jobId - 1

	if j._hasTerminated(i) {
		return nil, fmt.Errorf("job '%d' has already terminated", jobId)
	}

	return j.jobs[i], nil
}

type JobT struct {
	JobId   string
	Process *Process
}

func (j *jobs) List() []*JobT {
	var s []*JobT

	j.mutex.Lock()

	for i := range j.jobs {
		if j._hasTerminated(i) {
			continue
		}
		s = append(s, &JobT{
			JobId:   fmt.Sprintf("%%%d", i+1),
			Process: j.jobs[i],
		})
	}

	j.mutex.Unlock()

	return s
}

func (j *jobs) _hasTerminated(i int) bool {
	return j.jobs[i] == nil || j.jobs[i].HasTerminated()
}
