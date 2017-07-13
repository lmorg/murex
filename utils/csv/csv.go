package csv

import (
	"bufio"
	"io"
)

type Parser struct {
	reader    io.Reader
	Separator byte
	Quote     byte
	Comment   byte
}

func NewParser(reader io.Reader) (parser *Parser) {
	parser = new(Parser)
	parser.reader = reader
	parser.Separator = ','
	parser.Quote = '"'
	parser.Comment = '#'
	return
}

func (csv *Parser) ReadLine(callback func([]string)) (err error) {
	scanner := bufio.NewScanner(csv.reader)
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
			case csv.Comment:
				switch {
				case relativePos == 0:
					break
				default:
					current = append(current, b)
					relativePos++
				}

			case csv.Quote:
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

			case csv.Separator:
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

		callback(records)
	}

	err = scanner.Err()
	return
}
