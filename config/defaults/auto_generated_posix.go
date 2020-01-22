// +build !windows,!plan9

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_posix.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_posix.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "config set shell prompt {\n    out \"${pwd -> egrep -o '[^/]+$'} » \"\n}\n\nconfig set shell prompt-multiline {\n    let len = ${pwd_short -> wc -c} - 1\n    printf \"%${$len}s » \" $linenum\n}\n\n#autocomplete set man { [{\n#    \"IncExePath\": true\n#}] }\n\nautocomplete set man-summary { [{\n    \"IncExePath\": true,\n    \"AllowMultiple\": true\n}] }\n\nautocomplete set which { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set whereis { [{\n    \"IncExePath\": true\n}] }\n\nautocomplete set sudo { [\n    {\n        \"IncFiles\": true,\n        \"IncDirs\": true,\n        \"IncExePath\": true\n    },\n    {\n        \"NestedCommand\": true\n    }\n] }\n\nautocomplete set dd { [{\n    \"Flags\": [ \"if=\", \"of=\", \"bs=\", \"iflag=\", \"oflag=\", \"count=\", \"status=\" ]\n}] }\n\nprivate getHostsFile {\n    # Parse the hosts file and return uniq host names and IPs\n    egrep -v '^(#.*|\\s*)$' /etc/hosts -> sed -r 's/\\s+/\\n/g' -> sort -> uniq\n}\n\nautocomplete set ssh { [{\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set ping { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set rsync { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nautocomplete set sftp { [ {\n    \"Dynamic\": \"{ getHostsFile }\"\n}] }\n\nprivate kill-autocomplete {\n    # Autocomplete suggestion for `kill`\n    test define ps {\n        \"ExitNum\": 0\n    }\n    \n    test define map {\n        \"StdoutRegex\": (\\{(\".*?\":\".*?\",?)+\\})\n    }\n\n    ps <test_ps> -A -o pid -o command -> sed 1d -> set ps\n    map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n}\n\ntest define-unit private kill-autocomplete {\n    \"StdoutType\": \"json\",\n    \"StdoutRegex\": \"\\\\{\\\\\\\"[0-9]+\\\\\\\":\\\\\\\".*?\\\\\\\"(,|)\\\\}\"\n}\n\nautocomplete set kill {\n    [{\n        \"DynamicDesc\": ({ kill-autocomplete }),\n        \"ListView\": true,\n        \"AllowMultiple\": true\n    }]\n}\n\nautocomplete: set bg {\n    [{\n        \"DynamicDesc\": ({ fid-list --stopped }),\n        \"ListView\": true\n    }]\n}\n\nautocomplete: set fg {\n    [{\n        \"DynamicDesc\": ({ fid-list --stopped }),\n        \"ListView\": true\n    }]\n}\n\nprivate go-package {\n    # returns all the packages in $GOPATH\n    find <!null> $GOPATH/src/ -type d -not -path */.* -> sed -r s:$GOPATH/src/::\n}\n\nautocomplete set go { [{\n    \"Flags\": [ \"build\", \"clean\", \"doc\", \"env\", \"bug\", \"fix\", \"fmt\", \"generate\", \"get\", \"install\", \"list\", \"run\", \"test\", \"tool\", \"version\", \"vet\", \"help\" ],\n    \"FlagValues\": {\n        \"run\": [{\n            \"Dynamic\": ({ go-package }),\n            \"AutoBranch\": true\n        }],\n        \"test\": [{\n            \"Dynamic\": ({ go-package }),\n            \"AutoBranch\": true,\n            \"Flags\": [ \"./...\" ]\n        }],\n        \"build\": [{ \"Dynamic\": ({ go-package }), \"AutoBranch\": true }],\n        \"install\": [{ \"Dynamic\": ({ go-package }), \"AutoBranch\": true }],\n        \"fmt\": [{ \"IncFiles\": true }],\n        \"vet\": [{ \"IncFiles\": true }],\n        \"generate\": [{ \"IncFiles\": true }]\n    }\n}] }\n\nconfig eval shell safe-commands {\n    -> alter --merge / ([\n        \"cat\", \"ps\", \"grep\", \"egrep\", \"ls\", \"head\", \"tail\", \"printf\", \"awk\", \"sed\", \"td\", \"cut\"\n    ])\n}")
}
