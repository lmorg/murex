package streams

import (
	"bufio"
	"encoding/json"
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/debug"
	"regexp"
	"strconv"
	"strings"
)

func readMapJson(read Io, _ *config.Config, callback func(key, value string, last bool)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	var jObj interface{}
	err = json.Unmarshal(b, &jObj)
	if err == nil {

		switch v := jObj.(type) {
		case []interface{}:
			for i := range jObj.([]interface{}) {
				j, err := json.Marshal(jObj.([]interface{})[i])
				if err != nil {
					return err
				}
				callback(strconv.Itoa(i), string(j), i != len(jObj.([]interface{}))-1)
			}

		case map[string]interface{}, map[interface{}]interface{}:
			i := 1
			for key := range jObj.(map[string]interface{}) {
				j, err := json.Marshal(jObj.(map[string]interface{})[key])
				if err != nil {
					return err
				}
				callback(key, string(j), i != len(jObj.(map[string]interface{})))
				i++
			}
			return nil

		default:
			if debug.Enable {
				panic(v)
			}
		}
		return nil
	}
	return err
}

func readMapString(read Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	i := -1
	for scanner.Scan() {
		i++
		callback(strconv.Itoa(i), strings.TrimSpace(string(scanner.Bytes())), false)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

var rxWhitespace *regexp.Regexp = regexp.MustCompile(`\s+`)

func readMapGeneric(read Io, _ *config.Config, callback func(key, value string, last bool)) error {
	scanner := bufio.NewScanner(read)
	for scanner.Scan() {
		recs := rxWhitespace.Split(scanner.Text(), -1)
		for i := range recs {
			callback(strconv.Itoa(i), string(recs[i]), i+1 == len(recs))
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
