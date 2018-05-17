package defaults

func init() {
	murexProfile = append(murexProfile, `
func h {
    # Output the murex history log in a human readable format
    history -> foreach { -> [ Index Block ] -> sprintf: "%6s => %s\n" }
}

func aliases {
    # Output the aliases in human readable format
    alias -> formap: key val {
        $key -> sprintf: "%-10s { ${ $val -> sprintf: %s %s } }\n"
    }
}

autocomplete set cd { [{
    "IncDirs": true
}] }

autocomplete set mkdir { [{
    "IncDirs": true
}] }

autocomplete set rmdir { [{
    "IncDirs": true
}] }

autocomplete set exec { [
    {
        "IncFiles": true,
        "IncDirs": true,
        "IncExePath": true
    },
    {
        "NestedCommand": true
    }
] }

autocomplete set pty { [
    {
        "IncFiles": true,
        "IncDirs": true,
        "IncExePath": true
    },
    {
        "NestedCommand": true
    }
] }

autocomplete set format { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set cast { [{
    "Dynamic": ({ runtime: --unmarshallers })
}] }

autocomplete set tout { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set swivel-datatype { [{
    "Dynamic": ({ runtime: --marshallers })
}] }

autocomplete set config { [{
    "Flags": [ "get", "set" ],
    "FlagValues": {
        "get": [
            { "Dynamic": "{ config: -> formap k v { out $k } -> sort }" },
            { "Dynamic": "{ config: -> [ ${params->[2]} ] -> formap k v { out $k } -> sort }" }
        ],
        "set": [
            { "Dynamic": "{ config: -> formap k v { out $k } -> sort }" },
            { "Dynamic": "{ config: -> [ ${params->[2]} ] -> formap k v { out $k } -> sort }" },
            { "Dynamic": "{ config: -> [ ${params->[2]} ] -> [ ${params->[3]} ] -> [ Data-Type ] -> set dt; if { = dt==`+"`bool`"+` } { a [true,false] } { a ${config -> [ ${params->[2]} ] -> [ ${params->[3]} ] -> [ Default ]} } }" }
        ]
    }
}] }

autocomplete set runtime { [{
    "Flags": ["--vars", "--aliases" ,"--config" ,"--pipes" ,"--funcs" ,"--fids" ,"--arrays" ,"--maps" ,"--indexes" ,"--marshallers" ,"--unmarshallers" ,"--events" ,"--flags" ,"--memstats" ],
    "AllowMultiple": true
}] }

autocomplete set event { [
    {
        "Dynamic": "{ runtime: --events -> formap k v { out $k } }"
    }
] }

autocomplete set !event { [
    {
        "Dynamic": "{ runtime: --events -> formap k v { out $k } -> sort }"
    },
    {
        "Dynamic": "{ runtime: --events -> [ ${ params->[1] } ] -> formap k v { out $k } -> sort }"
    }
] }

autocomplete set autocomplete { [{
    "Flags" : [ "get", "set" ]
}] }

autocomplete set pipe { [
    {
        "Flags": [ "--create", "-c", "--close", "-x" ],
        "FlagValues": {
            "--close": [{
                "Dynamic": "{ runtime: --pipes -> formap k v { if { = k!=`+"`null`"+` } { $k } } }"
            }],
            "--create": [
                {
                    "AnyValue": true
                },
                {
                    "Flags": [ "--file", "--udp-dial", "--tcp-dial", "--udp-listen", "--tcp-listen" ],
                    "FlagValues": {
                        "--file": [{
                            "IncFiles": true
                        }]
                    }
                }
            ],
            "-x": [{ "Alias": "--close" }],
            "-c": [{ "Alias": "--create" }]
        }
    }
] }

autocomplete set git { [{
    "Flags": [ "clone", "init", "add", "mv", "reset", "rm", "bisect", "grep", "log", "show", "status", "branch", "checkout", "commit", "diff", "merge", "rebase", "tag", "fetch", "pull", "push" ],
    "FlagValues": {
        "init": [{ "Flags": ["--bare"] }],
        "add": [{ "IncFiles": true }],
        "mv": [{ "IncFiles": true }],
        "rm": [{ "IncFiles": true }],
        "checkout": [{
            "Dynamic": ({ git branch -> [ :1 ] }),
            "Flags": [ "-b" ]
        }]
    }
}] }

autocomplete set docker { [
    {
        "Flags": [ "config", "container", "image", "network", "node", "plugin", "secret", "service", "stack", "swarm", "system", "volume", "attach", "build", "commit", "cp", "create", "diff", "events", "exec", "export", "history", "images", "info", "inspect", "kill", "load", "login", "logout", "logs", "pause", "port", "ps", "pull", "push", "rename", "restart", "rm", "rmi", "run", "save", "search", "start", "stats", "stop", "tag", "top", "unpause", "update", "version", "wait" ]
    },
    {
        "Flags": [ "-t" ],
        "Optional": true,
        "AllowMultiple": true,
        "AnyValue": true
    },
    {
        "IncFiles": true
    }
] }

autocomplete set terraform { [{
    "Flags": ["apply","console","destroy","env","fmt","get","graph","import","init","output","plan","providers","push","refresh","show","taint","untaint","validate","version","workspace"],
    "FlagValues": {
        "workspace": [
            {
                "Flags": [ "new", "delete", "select", "list", "show" ]
            }
        ]
    }
}] }

autocomplete set gopass { [
    {
        "Flags": ["--yes","--clip","-c","--help","-h","--version","-v"],
        "AllowMultiple": true,
        "Dynamic": "{ exec: @{params} --generate-bash-completion }",
        "AutoBranch": true
    }
] }

autocomplete set debug { [{
    "Flags": ["on", "off"]
}] }

#tout: qs KB=1024&MB=${= 1024*1024}&GB=${= 1024*1024*1024}&TB=${= 1024*1024*1024*1024}&PB=${= 1024*1024*1024*1024*1024}&EB=${= 1024*1024*1024*1024*1024*1024}&min=60&hour=${= 60*60}&day=${= 60*60*24}&week=${= 60*60*24*7} -> format json -> set C
`)
}
