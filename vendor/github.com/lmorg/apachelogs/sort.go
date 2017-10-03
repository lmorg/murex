package apachelogs

import (
	"fmt"
	"sort"
)

// Structure for sorting slices of `AccessLine`. Currently experimental.
type Sort struct {
	AccessLog *AccessLog
	Key       AccessFieldId
}

// The following methods are untested. Example usage of experimental sort:
//
//	a := new(apachelogs.AccessLog)
//	sort := apachelogs.Sort{
//		AccessLogs: a,
//		Key:        apachelogs.AccFieldDate,
//	}
//	sort.Sort()
func (sal Sort) Remove(index int) { sal.AccessLog.Remove(index) }
func (sal Sort) Len() int         { return sal.AccessLog.Len() }
func (sal Sort) Swap(i, j int)    { sal.AccessLog.Swap(i, j) }
func (sal Sort) Sort()            { sort.Sort(sal) }
func (sal Sort) Less(i, j int) bool {
	switch sal.Key {
	case AccFieldIp:
		return (*sal.AccessLog)[i].IP < (*sal.AccessLog)[j].IP

	case AccFieldUserId:
		return (*sal.AccessLog)[i].UserID < (*sal.AccessLog)[j].UserID

	case AccFieldDateTime, AccFieldDate, AccFieldTime:
		return (*sal.AccessLog)[i].DateTime.Before((*sal.AccessLog)[j].DateTime)

	case AccFieldMethod:
		return (*sal.AccessLog)[i].Method < (*sal.AccessLog)[j].Method

	case AccFieldUri:
		return (*sal.AccessLog)[i].URI < (*sal.AccessLog)[j].URI

	case AccFieldQueryString:
		return (*sal.AccessLog)[i].QueryString < (*sal.AccessLog)[j].QueryString

	case AccFieldProtocol:
		return (*sal.AccessLog)[i].Protocol < (*sal.AccessLog)[j].Protocol

	case AccFieldStatus:
		return (*sal.AccessLog)[i].Status.I < (*sal.AccessLog)[j].Status.I

	case AccFieldSize:
		return (*sal.AccessLog)[i].Size < (*sal.AccessLog)[j].Size

	case AccFieldReferrer:
		return (*sal.AccessLog)[i].Referrer < (*sal.AccessLog)[j].Referrer

	case AccFieldUserAgent:
		return (*sal.AccessLog)[i].UserAgent < (*sal.AccessLog)[j].UserAgent

	case AccFieldProcTime:
		return (*sal.AccessLog)[i].ProcTime < (*sal.AccessLog)[j].ProcTime

	case AccFieldFileName:
		return (*sal.AccessLog)[i].FileName < (*sal.AccessLog)[j].FileName

	case 0:
		panic("Key unset on sort")
	}

	panic(fmt.Sprintf("%s is not a valid sort key", sal.Key))
}
