package mkarray

var mapRanges = []map[string]int{
	rangeWeekdayLong,
	rangeWeekdayShort,
	rangeMonthLong,
	rangeMonthShort,
	rangeSeason,
	rangeMoon,
}

var rangeWeekdayLong = map[string]int{
	"monday":    1,
	"tuesday":   2,
	"wednesday": 3,
	"thursday":  4,
	"friday":    5,
	"saturday":  6,
	"sunday":    7,
}

var rangeWeekdayShort = map[string]int{
	"mon": 1,
	"tue": 2,
	"wed": 3,
	"thu": 4,
	"fri": 5,
	"sat": 6,
	"sun": 7,
}

var rangeMonthLong = map[string]int{
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

var rangeMonthShort = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

var rangeSeason = map[string]int{
	"spring": 1,
	"summer": 2,
	"autumn": 3,
	"winter": 4,
}

var rangeMoon = map[string]int{
	"new moon":        1,
	"waxing crescent": 2,
	"first quarter":   3,
	"waxing gibbous":  4,
	"full moon":       5,
	"waning gibbous":  6,
	"third quarter":   7,
	"waning crescent": 8,
}

var dateFormat = []string{
	// dd mm yy

	"02-Jan-06",
	"02-January-06",
	"02-Jan-2006",
	"02-January-2006",

	"02 Jan 06",
	"02 January 06",
	"02 Jan 2006",
	"02 January 2006",

	"02/Jan/06",
	"02/January/06",
	"02/Jan/2006",
	"02/January/2006",

	// mm dd yy

	"Jan-02-06",
	"January-02-06",
	"Jan-02-2006",
	"January-02-2006",

	"Jan 02 06",
	"January 02 06",
	"Jan 02 2006",
	"January 02 2006",

	"Jan/02/06",
	"January/02/06",
	"Jan/02/2006",
	"January/02/2006",

	// dd mm

	"02-Jan",
	"02-January",

	"02 Jan",
	"02 January",

	"02/Jan",
	"02/January",
}
