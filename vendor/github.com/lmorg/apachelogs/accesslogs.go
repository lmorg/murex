//go:generate stringer -type=AccessFieldId

package apachelogs

import "time"

// Timestamp formatting in Apache access logs.
// You shouldn't need to change this but it's provided as a variable just in case.
var DateTimeAccessFormat string = "02/Jan/2006:15:04:05"

// A broken down record of each field in an access log.
type AccessLine struct {
	IP          string
	UserID      string
	DateTime    time.Time
	Method      string
	URI         string
	QueryString string
	Protocol    string
	Status      Status
	Size        int
	Referrer    string
	UserAgent   string
	ProcTime    int
	FileName    string
}

// Field ID when returning an `AccessLine` field by the `ByFieldId` method.
type AccessFieldId byte

// Constants assignable to `AccessFieldId`
const (
	AccFieldIp AccessFieldId = iota + 1
	AccFieldUserId
	AccFieldDateTime
	AccFieldDate
	AccFieldTime
	AccFieldMethod
	AccFieldUri
	AccFieldQueryString
	AccFieldProtocol
	AccFieldStatus
	AccFieldSize
	AccFieldReferrer
	AccFieldUserAgent
	AccFieldProcTime
	AccFieldFileName
)

// This method allows you to recall an `AccessLine` field dynamically at runtime
// ie for instances where you want the field to be user selectable rather than hardcoded into the compiled binary.
func (a AccessLine) ByFieldId(id AccessFieldId) interface{} {
	switch id {
	case AccFieldIp:
		return a.IP
	case AccFieldUserId:
		return a.UserID
	case AccFieldDateTime, AccFieldDate, AccFieldTime:
		return a.DateTime
	case AccFieldMethod:
		return a.Method
	case AccFieldUri:
		return a.URI
	case AccFieldQueryString:
		return a.QueryString
	case AccFieldProtocol:
		return a.Protocol
	case AccFieldStatus:
		return a.Status.A
	case AccFieldSize:
		return a.Size
	case AccFieldReferrer:
		return a.Referrer
	case AccFieldUserAgent:
		return a.UserAgent
	case AccFieldProcTime:
		return a.ProcTime
	case AccFieldFileName:
		return a.FileName
	default:
		return nil
	}
}

// This method allows you to set an `AccessLine` field dynamically at runtime
// ie for instances where you want the field to be user selectable rather than hardcoded into the compiled binary.
func (a *AccessLine) SetFieldID(id AccessFieldId, val interface{}) {
	switch id {
	case AccFieldIp:
		a.IP = val.(string)
	case AccFieldUserId:
		a.UserID = val.(string)
	case AccFieldDateTime, AccFieldDate, AccFieldTime:
		a.DateTime = val.(time.Time)
	case AccFieldMethod:
		a.Method = val.(string)
	case AccFieldUri:
		a.URI = val.(string)
	case AccFieldQueryString:
		a.QueryString = val.(string)
	case AccFieldProtocol:
		a.Protocol = val.(string)
	case AccFieldStatus:
		a.Status = NewStatus(val.(string))
	case AccFieldSize:
		a.Size = val.(int)
	case AccFieldReferrer:
		a.Referrer = val.(string)
	case AccFieldUserAgent:
		a.UserAgent = val.(string)
	case AccFieldProcTime:
		a.ProcTime = val.(int)
	case AccFieldFileName:
		a.FileName = val.(string)
	}
}

// This type is provided to offer some easier tools for working with slices of `AccessLine`.
type AccessLog []*AccessLine

func (al AccessLog) Remove(index int) { al = append(al[:index], al[index+1:]...) }
func (al AccessLog) Len() int         { return len(al) }
func (al AccessLog) Swap(i, j int)    { al[i], al[j] = al[j], al[i] }
