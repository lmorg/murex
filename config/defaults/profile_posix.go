// +build !windows

package defaults

func init() {
	murexProfile = append(murexProfile, `

config set shell prompt           { out "${pwd -> egrep -o '[^/]+$'} » " }
config set shell prompt-multiline { $linenum -> sprintf "%${eval ${pwd -> egrep -o '[^/]+$' -> wc -c}-1}s » " }

autocomplete set man { [{
    "IncExePath": true
}] }

autocomplete set which { [{
    "IncExePath": true
}] }

autocomplete set whereis { [{
    "IncExePath": true
}] }

autocomplete set sudo { [
    {
        "IncFiles": true,
        "IncDirs": true,
        "IncExePath": true
    },
    {
        "NestedCommand": true
    }
] }

autocomplete set dd { [{
    "Flags": [ "if=", "of=", "bs=", "iflag=", "oflag=", "count=", "status=" ]
}] }

autocomplete set ssh { [{
    "Dynamic": "{ getHostsFile }"
}] }

autocomplete set ping { [ {
    "Dynamic": "{ getHostsFile }"
}] }

autocomplete set rsync { [ {
    "Dynamic": "{ getHostsFile }"
}] }

autocomplete set sftp { [ {
    "Dynamic": "{ getHostsFile }"
}] }

func getHostsFile {
    # Parse the hosts file and return uniq host names and IPs
    egrep -v '^(#.*|\s*)$' /etc/hosts -> sed -r 's/\s+/\n/g' -> sort -> uniq
}

autocomplete set go { [
    { "Flags": [ "build", "clean", "doc", "env", "bug", "fix", "fmt", "generate", "get", "install", "list", "run", "test", "tool", "version", "vet", "help" ] },
    {
        "Dynamic": ({ find <!null> $GOPATH/src/ -type d -not -path */.* -> sed -r s:$GOPATH/src/:: }),
        "AutoBranch": true
    }
] }

autocomplete set kill {
    [{
        "DynamicDesc": ({
            test define ps {
                "ExitNum": 0
            }
            test define map {
                "OutRegexp": (\{(".*?":".*?",?)+\})
            }

            ps <test_ps> -A -o pid,cmd --no-headers -> set ps
            map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }
        }),
        "ListView": true,
        "AllowMultiple": true
    }]
}
`)
}
