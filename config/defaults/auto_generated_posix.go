// +build !windows

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_posix.mx using docgen.

   Please do not manually edit this file because it will be overwritten.
*/

func init() {
	murexProfile = append(murexProfile, "config set shell prompt           { out \"${pwd -> egrep -o '[^/]+$'} » \" }\nconfig set shell prompt-multiline { $linenum -> sprintf \"%${eval ${pwd -> egrep -o '[^/]+$' -> wc -c}-1}s » \" }\n\n#autocomplete set man { [{\n#    \"IncExePath\": true\n#}] }\n\nautocomplete set man-summary { [{\n    \"IncExePath\": true,\n    \"AllowMultiple\": true\n}] }\n\n\nautocomplete set which { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set whereis { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set sudo { [\n    {\n        \"IncFiles\": true,\n        \"IncDirs\": true,\n        \"IncExePath\": true\n    },\n    {\n        \"NestedCommand\": true\n    }\n] }\n\nautocomplete set dd { [{\n    \"Flags\": [ \"if=\", \"of=\", \"bs=\", \"iflag=\", \"oflag=\", \"count=\", \"status=\" ]\n}] }\n\nprivate getHostsFile {\n    # Parse the hosts file and return uniq host names and IPs\n    egrep -v '^(#.*|\\s*)$' /etc/hosts -> sed -r 's/\\s+/\\n/g' -> sort -> uniq\n}\n\nautocomplete set ssh { [{\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set ping { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set rsync { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set sftp { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set go { [\n    { \"Flags\": [ \"build\", \"clean\", \"doc\", \"env\", \"bug\", \"fix\", \"fmt\", \"generate\", \"get\", \"install\", \"list\", \"run\", \"test\", \"tool\", \"version\", \"vet\", \"help\" ] },\n    {\n        \"Dynamic\": ({ find <!null> $GOPATH/src/ -type d -not -path */.* -> sed -r s:$GOPATH/src/:: }),\n        \"AutoBranch\": true\n    }\n] }\n\nautocomplete set kill {\n    [{\n        \"DynamicDesc\": ({\n            test define ps {\n                \"ExitNum\": 0\n            }\n            test define map {\n                \"OutRegexp\": (\\{(\".*?\":\".*?\",?)+\\})\n            }\n\n            ps <test_ps> -A -o pid,cmd --no-headers -> set ps\n            map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n        }),\n        \"ListView\": true,\n        \"AllowMultiple\": true\n    }]\n}\n\n\nprivate get.datatypes {\n\tpipe privategetdatatypes\n\truntime: --readarray -> format: <privategetdatatypes> str\n\truntime: --writearray -> format: <privategetdatatypes> str\n\truntime: --readmap -> format: <privategetdatatypes> str\n\truntime: --marshallers -> format: <privategetdatatypes> str\n\truntime: --unmarshallers -> format: <privategetdatatypes> str\n\n\t!pipe privategetdatatypes\n\t<privategetdatatypes> -> set datatypes\n\t$datatypes -> sort -> uniq\n}\n\nautocomplete set cast { [{\n    \"Dynamic\": ({ get.datatypes })\n}] }\n\nautocomplete set tout { [{\n    \"Dynamic\": ({ get.datatypes })\n}] }")
}
