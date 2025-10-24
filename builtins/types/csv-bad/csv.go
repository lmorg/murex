package csvbad

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

// Parser is the CSV parser settings
type Parser struct {
	reader    io.Reader
	Separator byte
	Quote     byte
	Comment   byte
	Headings  bool
}

// NewParser creates a new CSV reader and writer - albeit it doesn't conform to Go's io.Reader / io.Writer any.
// The sensible thing might have been to create this as a marshaller but it's written now and works so little point
// breaking it at this point in time.
func NewParser(reader io.Reader, config *config.Config) (parser *Parser, err error) {
	parser = new(Parser)
	parser.reader = reader

	parser.Separator = ','
	parser.Quote = '"'
	parser.Comment = '#'

	v, err := config.Get("csv", "separator", types.String)
	if err != nil {
		return nil, err
	}
	if len(v.(string)) > 0 {
		parser.Separator = v.(string)[0]
	}

	v, err = config.Get("csv", "comment", types.String)
	if err != nil {
		return nil, err
	}
	if len(v.(string)) > 0 {
		parser.Comment = v.(string)[0]
	}

	v, err = config.Get("csv", "headings", types.Boolean)
	if err != nil {
		return nil, err
	}
	parser.Headings = v.(bool)
	return
}

// ReadLine - read a line from a CSV file
func (parser *Parser) ReadLine(callback func(records []string, headings []string)) (err error) {
	scanner := bufio.NewScanner(parser.reader)

	var headings []string

	for scanner.Scan() {
		var (
			records     []string
			current     []byte
			escape      bool
			quoted      bool
			relativePos int
		)

		for _, b := range scanner.Bytes() {
			switch b {
			case parser.Comment:
				switch {
				case relativePos == 0:
					break
				default:
					current = append(current, b)
					relativePos++
				}

			case parser.Quote:
				switch {
				case escape:
					current = append(current, b)
					escape = false
					relativePos++
				case quoted:
					quoted = false
				case relativePos == 0:
					quoted = true
				default:
					current = append(current, b)
					relativePos++
				}

			case '\\':
				switch {
				case escape:
					current = append(current, b)
					escape = false
					relativePos++
				default:
					escape = true
				}

			case parser.Separator:
				switch {
				case escape, quoted:
					current = append(current, b)
					escape = false
					relativePos++
				default:
					records = append(records, string(current))
					current = make([]byte, 0)
					relativePos = 0
				}

			default:
				current = append(current, b)
				escape = false
				relativePos++
			}
		}

		if len(current) > 0 {
			records = append(records, string(current))
		}

		if len(headings) == 0 {
			if parser.Headings {
				headings = records
			} else {
				for i := range records {
					headings = append(headings, strconv.Itoa(i))
				}
			}
			continue
		}

		if len(records) > len(headings) {
			for i := len(headings); i < len(records); i++ {
				headings = append(headings, strconv.Itoa(i))
			}
		}

		callback(records, headings)
	}

	err = scanner.Err()
	return
}

// ArrayToCsv marshals a list into a CSV line
func (parser *Parser) ArrayToCsv(array []string) (csv []byte) {
	quote := string(parser.Quote)
	escapedQuote := `\` + quote
	separator := quote + string(parser.Separator) + quote

	for i := range array {
		array[i] = strings.Replace(array[i], `\`, `\\`, -1)
		array[i] = strings.Replace(array[i], "\n", `\n`, -1)
		array[i] = strings.Replace(array[i], "\r", `\r`, -1)
		array[i] = strings.Replace(array[i], quote, escapedQuote, -1)
	}
	csv = []byte(quote + strings.Join(array, separator) + quote)
	return
}
