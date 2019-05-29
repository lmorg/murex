// +build darwin

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_osx.mx using docgen.

   Please do not manually edit this file because it will be overwritten.
*/

func init() {
	murexProfile = append(murexProfile, "autocomplete set kill {\n    [{\n        \"DynamicDesc\": ({\n            test define ps {\n                \"ExitNum\": 0\n            }\n            test define map {\n                \"OutRegexp\": (\\{(\".*?\":\".*?\",?)+\\})\n            }\n\n            ps <test_ps> -A -o pid -o command -> sed 1d -> set ps\n            map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n        }),\n        \"ListView\": true,\n        \"AllowMultiple\": true\n    }]\n}")
}
