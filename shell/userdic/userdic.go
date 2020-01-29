package userdic

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

var dictionary = []string{
	config.AppName,
}

// IsInUserDic checks if word is in user dictionary
func IsInUserDic(word string) bool {
	s := strings.ToLower(word)
	for i := range dictionary {
		if s == strings.ToLower(dictionary[i]) {
			return true
		}
	}

	return false
}

// GetSpellcheckUserDic returns a slice of the spellcheckUserDic
func GetSpellcheckUserDic() []string {
	a := make([]string, len(dictionary))
	copy(a, dictionary)
	return a
}

// ReadSpellcheckUserDic returns an interface{} of the user dictionary.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func ReadSpellcheckUserDic() (interface{}, error) {
	return GetSpellcheckUserDic(), nil
}

// WriteSpellcheckUserDic takes a JSON-encoded string and writes it to the
// spellcheckUserDic slice.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func WriteSpellcheckUserDic(v interface{}) error {
	switch v.(type) {
	case string:
		return json.Unmarshal([]byte(v.(string)), &dictionary)

	default:
		return fmt.Errorf("Invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
