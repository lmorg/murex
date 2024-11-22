package lang_test

import (
	"strings"
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

func convertProcessSliceToStringSlice(processes []*lang.Process) []string {
	s := make([]string, len(processes))

	for i, p := range processes {
		switch {
		case p == nil:
			s[i] = "nil"

		case p.HasTerminated():
			s[i] = "terminated"

		default:
			s[i] = "active"
		}
	}

	return s
}

func TestJobsGarbageCollect(t *testing.T) {
	tests := []struct {
		PreState  []*lang.Process
		PostState []*lang.Process
		Kill      []bool
	}{
		{
			PreState:  []*lang.Process{},
			PostState: []*lang.Process{},
		},
		{
			PreState:  []*lang.Process{nil},
			PostState: []*lang.Process{},
		},
		{
			PreState:  []*lang.Process{nil, nil, nil},
			PostState: []*lang.Process{},
		},
		/////
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				nil,
			},
			PostState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				nil,
				lang.NewTestProcess(),
			},
			PostState: []*lang.Process{
				lang.NewTestProcess(),
				nil,
				lang.NewTestProcess(),
			},
		},
		{
			PreState: []*lang.Process{
				nil,
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			PostState: []*lang.Process{
				nil,
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
		},
		/////
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				nil,
				lang.NewTestProcess(),
			},
			PostState: []*lang.Process{
				lang.NewTestProcess(),
			},
			Kill: []bool{false, false, true},
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			PostState: []*lang.Process{},
			Kill:      []bool{true, true, true},
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			PostState: []*lang.Process{
				nil,
				nil,
				lang.NewTestProcess(),
			},
			Kill: []bool{true, true, false},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jobs := lang.NewJobs()
		for pi, p := range test.PreState {
			jobs.Add(p)
			if p != nil && test.Kill != nil {
				p.SetTerminatedState(test.Kill[pi])
			}
		}

		jobs.GarbageCollect()

		expState := convertProcessSliceToStringSlice(test.PostState)
		actState := convertProcessSliceToStringSlice(jobs.Raw())

		expJson := json.LazyLogging(expState)
		actJson := json.LazyLogging(actState)

		if expJson != actJson {

			t.Errorf("Unexpected state in test %d:", i)
			t.Logf("PreState: %s", json.LazyLogging(convertProcessSliceToStringSlice(test.PreState)))
			t.Logf("Expected: %s", expJson)
			t.Logf("Actual:   %s", actJson)
		}
	}
}

func TestJobsList(t *testing.T) {
	reformat := func(s string) string {
		s = strings.ReplaceAll(s, `\u003c`, "<")
		return strings.ReplaceAll(s, `\u003e`, ">")
	}

	tests := []struct {
		PreState []*lang.Process
		Count    int
		Kill     []bool
	}{
		{
			PreState: []*lang.Process{},
			Count:    0,
		},
		{
			PreState: []*lang.Process{nil},
			Count:    0,
		},
		{
			PreState: []*lang.Process{nil, nil, nil},
			Count:    0,
		},
		/////
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				nil,
			},
			Count: 2,
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				nil,
				lang.NewTestProcess(),
			},
			Count: 2,
		},
		{
			PreState: []*lang.Process{
				nil,
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			Count: 2,
		},
		/////
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				nil,
				lang.NewTestProcess(),
			},
			Kill:  []bool{false, false, true},
			Count: 1,
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			Kill:  []bool{true, true, true},
			Count: 0,
		},
		{
			PreState: []*lang.Process{
				lang.NewTestProcess(),
				lang.NewTestProcess(),
				lang.NewTestProcess(),
			},
			Kill:  []bool{true, true, false},
			Count: 1,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		jobs := lang.NewJobs()
		for pi, p := range test.PreState {
			jobs.Add(p)
			if p != nil && test.Kill != nil {
				p.SetTerminatedState(test.Kill[pi])
			}
		}

		list := jobs.List()

		actJson := json.LazyLogging(list)

		if len(list) != test.Count {
			t.Errorf("Unexpected state in test %d:", i)
			t.Logf("PreState: %s", json.LazyLogging(convertProcessSliceToStringSlice(test.PreState)))
			t.Logf("Expected: %d", test.Count)
			t.Logf("Actual:   %d", len(list))
			t.Logf("act json: %s", reformat(actJson))
		}
	}
}
