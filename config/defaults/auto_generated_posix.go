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
	murexProfile = append(murexProfile, "config set shell prompt           { out \"${pwd -> egrep -o '[^/]+$'} » \" }\nconfig set shell prompt-multiline { $linenum -> sprintf \"%${eval ${pwd -> egrep -o '[^/]+$' -> wc -c}-1}s » \" }\n\n#autocomplete set man { [{\n#    \"IncExePath\": true\n#}] }\n\nautocomplete set man-summary { [{\n    \"IncExePath\": true,\n    \"AllowMultiple\": true\n}] }\n\n\nautocomplete set which { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set whereis { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set sudo { [\n    {\n        \"IncFiles\": true,\n        \"IncDirs\": true,\n        \"IncExePath\": true\n    },\n    {\n        \"NestedCommand\": true\n    }\n] }\n\nautocomplete set dd { [{\n    \"Flags\": [ \"if=\", \"of=\", \"bs=\", \"iflag=\", \"oflag=\", \"count=\", \"status=\" ]\n}] }\n\nprivate getHostsFile {\n    # Parse the hosts file and return uniq host names and IPs\n    egrep -v '^(#.*|\\s*)$' /etc/hosts -> sed -r 's/\\s+/\\n/g' -> sort -> uniq\n}\n\nautocomplete set ssh { [{\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set ping { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set rsync { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set sftp { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set go { [\n    { \"Flags\": [ \"build\", \"clean\", \"doc\", \"env\", \"bug\", \"fix\", \"fmt\", \"generate\", \"get\", \"install\", \"list\", \"run\", \"test\", \"tool\", \"version\", \"vet\", \"help\" ] },\n    {\n        \"Dynamic\": ({ find <!null> $GOPATH/src/ -type d -not -path */.* -> sed -r s:$GOPATH/src/:: }),\n        \"AutoBranch\": true\n    }\n] }\n\nprivate get.datatypes {\n    runtime: --readarray -> format: str\n    runtime: --writearray -> format: str\n    runtime: --readmap -> format: str\n    runtime: --marshallers -> format: str\n    runtime: --unmarshallers -> format: str\n}\n\nprivate datatypes {\n    get.datatypes -> sort -> uniq -> cast str\n}\n\nautocomplete set cast { [{\n    \"Dynamic\": ({ datatypes })\n}] }\n\nautocomplete set tout { [{\n    \"Dynamic\": ({ datatypes })\n}] }\n\nautocomplete set set { [{\n    \"Dynamic\": ({ datatypes })\n}] }")
}
