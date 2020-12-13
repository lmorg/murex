package datatools_test

import (
	"strings"
	"testing"

	_ "github.com/lmorg/murex/builtins"
	"github.com/lmorg/murex/test"
)

func TestStructKeysNoParams(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
		/address/city
		/address/postalCode
		/address/state
		/address/streetAddress
		/age
		/children
		/firstName
		/isAlive
		/lastName
		/phoneNumbers
		/phoneNumbers/0
		/phoneNumbers/0/number
		/phoneNumbers/0/type
		/phoneNumbers/1
		/phoneNumbers/1/number
		/phoneNumbers/1/type
		/phoneNumbers/2
		/phoneNumbers/2/number
		/phoneNumbers/2/type
		/spouse
	`

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParamStr(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/0/number
	  /phoneNumbers/0/type
	  /phoneNumbers/1
	  /phoneNumbers/1/number
	  /phoneNumbers/1/type
	  /phoneNumbers/2
	  /phoneNumbers/2/number
	  /phoneNumbers/2/type
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys foobar -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParam0(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/0/number
	  /phoneNumbers/0/type
	  /phoneNumbers/1
	  /phoneNumbers/1/number
	  /phoneNumbers/1/type
	  /phoneNumbers/2
	  /phoneNumbers/2/number
	  /phoneNumbers/2/type
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys 0 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParamMinus10(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/0/number
	  /phoneNumbers/0/type
	  /phoneNumbers/1
	  /phoneNumbers/1/number
	  /phoneNumbers/1/type
	  /phoneNumbers/2
	  /phoneNumbers/2/number
	  /phoneNumbers/2/type
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys -10 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParam1(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys 1 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParam2(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/1
	  /phoneNumbers/2
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys 2 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParam3(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/0/number
	  /phoneNumbers/0/type
	  /phoneNumbers/1
	  /phoneNumbers/1/number
	  /phoneNumbers/1/type
	  /phoneNumbers/2
	  /phoneNumbers/2/number
	  /phoneNumbers/2/type
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys 3 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}

func TestStructKeysParam200(t *testing.T) {
	var input = `{
		"firstName": "John",
		"lastName": "Smith",
		"isAlive": true,
		"age": 27,
		"address": {
		  "streetAddress": "21 2nd Street",
		  "city": "New York",
		  "state": "NY",
		  "postalCode": "10021-3100"
		},
		"phoneNumbers": [
		  {
			"type": "home",
			"number": "212 555-1234"
		  },
		  {
			"type": "office",
			"number": "646 555-4567"
		  },
		  {
			"type": "mobile",
			"number": "123 456-7890"
		  }
		],
		"children": [],
		"spouse": null
	  }`

	var expected = `/address
	  /address/city
	  /address/postalCode
	  /address/state
	  /address/streetAddress
	  /age
	  /children
	  /firstName
	  /isAlive
	  /lastName
	  /phoneNumbers
	  /phoneNumbers/0
	  /phoneNumbers/0/number
	  /phoneNumbers/0/type
	  /phoneNumbers/1
	  /phoneNumbers/1/number
	  /phoneNumbers/1/type
	  /phoneNumbers/2
	  /phoneNumbers/2/number
	  /phoneNumbers/2/type
	  /spouse
  `

	expected = strings.ReplaceAll(expected, "\t", "")
	expected = strings.ReplaceAll(expected, " ", "")

	tests := []test.MurexTest{{
		Block: `
			tout json (` + input + `) -> struct-keys 200 -> format str -> msort
		`,
		ExitNum: 0,
		Stdout:  expected,
		Stderr:  ``,
	}}

	test.RunMurexTests(tests, t)
}
