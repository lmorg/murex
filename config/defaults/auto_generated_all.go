package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_all.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_all.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "func h {\n    # Output the murex history log in a human readable format\n    history -> foreach { -> set json line; out \"$line[Index]: $line[Block]\" } -> cast *\n}\n\nfunc aliases {\n\t# Output the aliases in human readable format\n\truntime: --aliases -> formap name alias {\n        $name -> sprintf: \"%10s => ${esccli @alias}\\n\"\n\t} -> cast str\n}\n\ntest define-unit public aliases {\n    \"PreBlock\": ({\n        global ALIAS_UNIT_TEST=${rand: int 9999999999}\n        alias $ALIAS_UNIT_TEST=example param1 param2 param3\n    }),\n    \"StdoutRegex\": \"([- _0-9a-zA-Z]+ => .*?\\n)+\",\n    \"StdoutType\": \"str\",\n    \"PostBlock\": ({\n        !alias $ALIAS_UNIT_TEST\n        !global ALIAS_UNIT_TEST\n    })\n}\n\nautocomplete set cd { [{\n    \"IncDirs\": true\n}] }\n\nautocomplete set mkdir { [{\n    \"IncDirs\": true,\n    \"AllowMultiple\": true\n}] }\n\nautocomplete set rmdir { [{\n    \"IncDirs\": true,\n    \"AllowMultiple\": true\n}] }\n\nautocomplete set exec { [\n    {\n        \"IncFiles\": true,\n        \"IncDirs\": true,\n        \"IncExePath\": true\n    },\n    {\n        \"NestedCommand\": true\n    }\n] }\n\nprivate get.datatypes {\n    runtime: --readarray -> format: str\n    runtime: --writearray -> format: str\n    runtime: --readmap -> format: str\n    runtime: --marshallers -> format: str\n    runtime: --unmarshallers -> format: str\n}\n\nprivate datatypes {\n    get.datatypes -> sort -> uniq -> cast str\n}\n\ntest define-unit private datatypes {\n    \"StdoutRegex\": \"[a-z]+\",\n    \"StdoutType\":  \"str\",\n    \"StdoutBlock\": ({\n        -> len -> set len;\n        if { = len>0 } then {\n            out \"Len greater than 0\"\n        } else {\n            err \"No elements returned\"\n        }\n    })\n}\n\nautocomplete set cast { [{\n    \"Dynamic\": ({ datatypes })\n}] }\n\nautocomplete set tout { [{\n    \"Dynamic\": ({ datatypes })\n}] }\n\nautocomplete set set { [{\n    \"Dynamic\": ({ datatypes })\n}] }\n\nautocomplete set format { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }\n\nautocomplete set swivel-datatype { [{\n    \"Dynamic\": ({ runtime: --marshallers })\n}] }\n\nprivate config.get.apps {\n    config: -> formap k v { out $k } -> msort\n}\n\n#private config.get.keys {\n#    config: -> [ ${params->[2]} ] -> formap k v { out $k } -> msort\n#}\n#\n#private config.get.meta {\n#    config: -> [ ${params->[2]} ] -> [ ${params->[3]} ] -> [ ${params->[4]} ]\n#}\n\nautocomplete set config { [{\n    \"Flags\": [ \"get\", \"set\", \"eval\", \"define\", \"default\" ],\n    \"FlagValues\": {\n        \"get\": [\n            { \"Dynamic\": ({ config.get.apps }) },\n            { \"Dynamic\": ({ source { config } -> [ $ARGS[2] ] -> formap k v { out $k } -> msort }) }\n            #{ \"Dynamic\": ({ err $SELF\\n$ARGS }) }\n        ],               \n        \"set\": [\n            { \"Dynamic\": ({ config.get.apps }) },\n            { \"Dynamic\": ({ source { config } -> [ $ARGS[2] ] -> formap k v { out $k } -> msort }) },\n            { \"Dynamic\": ({\n\t\t\t\tswitch {\n\t\t\t\t\tcase { = `${source { config } -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Data-Type ]}`==`bool` } {\n\t\t\t\t\t\tja [true,false]\n\t\t\t\t\t}\n\n\t\t\t\t\tcase { source { config } -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ <!null> Options ] } {\n\t\t\t\t\t\tsource { config } -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Options ]\n\t\t\t\t\t}\n\t\t\t\t\t\n                \tcatch {\n\t\t\t\t\t\tout ${source { config } -> [ $ARGS[2] ] -> [ $ARGS[3] ] -> [ Default ]}\n\t\t\t\t\t}\n\t\t\t\t}\n\t\t\t}) }\n        ],\n        \"eval\": [\n            { \"Dynamic\": ({ config.get.apps }) },\n            { \"Dynamic\": ({ source { config } -> [ $ARGS[2] ] -> formap k v { out $k } -> msort }) }\n        ],\n        \"default\": [\n            { \"Dynamic\": ({ config.get.apps }) },\n            { \"Dynamic\": ({ source { config } -> [ $ARGS[2] ] -> formap k v { out $k } -> msort }) }\n        ]\n    }\n}] }\n\nautocomplete set !config { [\n    { \"Dynamic\": ({ config.get.apps }) },\n    { \"Dynamic\": ({ source { config } -> [ $ARGS[1] ] -> formap k v { out $k } -> msort }) }\n] }\n\nautocomplete set event { [\n    {\n        \"Dynamic\": ({ runtime: --events -> formap k v { out $k } })\n    }\n] }\n\nautocomplete set !event { [\n    {\n        \"Dynamic\": ({ runtime: --events -> formap k v { out $k } -> msort })\n    },\n    {\n        \"Dynamic\": ({ runtime: --events -> [ $ARGS[1] ] -> formap k v { out $k } -> msort })\n    }\n] }\n\nautocomplete set !alias { [\n    {\n        \"Dynamic\": ({ runtime: --aliases -> formap k v { out $k } })\n    }\n]}\n\nautocomplete set !func { [\n    {\n        \"Dynamic\": ({ runtime: --funcs -> formap k v { out $k } })\n    }\n]}\n\nautocomplete set autocomplete { [{\n    \"Flags\" : [ \"get\", \"set\" ]\n}] }\n\nprivate git-branch {\n    git branch -> [ :0 ] -> !match *\n}\n\nautocomplete set git { [{\n    \"Flags\": [ \"clone\", \"init\", \"add\", \"mv\", \"reset\", \"rm\", \"bisect\", \"grep\", \"log\", \"show\", \"status\", \"branch\", \"checkout\", \"commit\", \"diff\", \"merge\", \"rebase\", \"tag\", \"fetch\", \"pull\", \"push\", \"stash\" ],\n    \"FlagValues\": {\n        \"init\": [{\n            \"Flags\": [\"--bare\"]\n        }],\n        \"add\": [{\n            \"IncFiles\": true,\n            \"AllowMultiple\": true\n        }],\n        \"mv\": [\n            { \"IncFiles\": true }\n        ],\n        \"rm\": [{\n            \"IncFiles\": true,\n            \"AllowMultiple\": true\n        }],\n        \"checkout\": [{\n            \"Dynamic\": ({ git-branch }),\n            \"Flags\": [ \"-b\" ]\n        }],\n        \"merge\": [{\n            \"Dynamic\": ({ git-branch })\n        }],\n        \"commit\": [{\n            \"Flags\": [\"-a\", \"-m\", \"--amend\"],\n            \"FlagValues\": {\n                \"--amend\": [{ \"AnyValue\": true }]\n            },\n            \"AllowMultiple\": true\n        }]\n    }\n}] }\n\nautocomplete set docker { [\n    {\n        \"Flags\": [ \"config\", \"container\", \"image\", \"network\", \"node\", \"plugin\", \"secret\", \"service\", \"stack\", \"swarm\", \"system\", \"volume\", \"attach\", \"build\", \"commit\", \"cp\", \"create\", \"diff\", \"events\", \"exec\", \"export\", \"history\", \"images\", \"info\", \"inspect\", \"kill\", \"load\", \"login\", \"logout\", \"logs\", \"pause\", \"port\", \"ps\", \"pull\", \"push\", \"rename\", \"restart\", \"rm\", \"rmi\", \"run\", \"save\", \"search\", \"start\", \"stats\", \"stop\", \"tag\", \"top\", \"unpause\", \"update\", \"version\", \"wait\" ]\n    },\n    {\n        \"Flags\": [ \"-t\" ],\n        \"Optional\": true,\n        \"AllowMultiple\": true,\n        \"AnyValue\": true\n    },\n    {\n        \"IncFiles\": true\n    }\n] }\n\nautocomplete set terraform { [{\n    \"Flags\": [\"apply\",\"console\",\"destroy\",\"env\",\"fmt\",\"get\",\"graph\",\"import\",\"init\",\"output\",\"plan\",\"providers\",\"push\",\"refresh\",\"show\",\"taint\",\"untaint\",\"validate\",\"version\",\"workspace\"],\n    \"FlagValues\": {\n        \"workspace\": [\n            {\n                \"Flags\": [ \"new\", \"delete\", \"select\", \"list\", \"show\" ]\n            }\n        ]\n    }\n}] }\n\nautocomplete set gopass { [{\n    \"Flags\": [\"--yes\",\"--clip\",\"-c\",\"--help\",\"-h\",\"--version\",\"-v\"],\n    \"AllowMultiple\": true,\n    \"Dynamic\": \"{ exec: @{params} --generate-bash-completion }\",\n    \"AutoBranch\": true\n}] }\n\nautocomplete set debug { [{\n    \"Flags\": [\"on\", \"off\"]\n}] }\n\n#func progress {\n#    # Pulls the read progress of a Linux pid via /proc/$pid/fdinfo (only runs on Linux)\n#\n#    if { = `+\"`${os}`==`linux`\"+` } then {\n#        params -> [ 1 ] -> set pid\n#        \n#        g <!null> /proc/$pid/fd/* -> regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+)) -> foreach <!null> fd {\n#            trypipe <!null> {\n#                open /proc/$pid/fdinfo/$fd -> cast yaml -> [ pos ] -> set pos\n#                readlink: /proc/$pid/fd/$fd -> set file\n#                du -b $file -> [ :0 ] -> set int size\n#                if { = size > 0 } then {\n#                    = ($pos/$size)*100 -> set int percent\n#                    out \"$percent% ($pos/$size) $file\"\n#                }\n#            }\n#        }\n#    }\n#}\n\n#autocomplete set progress {\n#    [{\n#        \"DynamicDesc\": ({\n#            ps -A -o pid,cmd --no-headers -> set ps\n#            map { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n#        }),\n#        \"ListView\": true\n#    }]\n#}\n\nprivate fid.list.all {\n    set json fids=([])\n    set json procs=([])\n    fid-list -> foreach f {\n        if { = $f[Scope]!=$SELF[Scope] && $f[FID]!=$SELF[Parent] && $f[FID]!=0 } then {\n            $fids -> append $f[FID] -> set fids\n            $procs -> append \"$f[Command] $f[Parameters]\" -> set procs\n        }\n    }\n    map { $fids } { $procs }\n}\n\nprivate fid.list.stopped {\n    set json fids=([])\n    set json procs=([])\n    fid-list -> foreach f {\n        if { = `$f[State]`==`Stopped` && $f[Scope]!=$SELF[Scope] && $f[FID]!=$SELF[Parent] && $f[FID]!=0 } then {\n            $fids -> append $f[FID] -> set fids\n            $procs -> append \"$f[Command] $f[Parameters]\" -> set procs\n        }\n    }\n    map { $fids } { $procs }\n}\n\nfunc jobs {\n    # List any stopped jobs\n    out \"FID - Process\"\n    fid-list -> foreach f {\n        if { = `$f[State]`==`Stopped` } {\n            out \"$f[FID] - $f[Command] $f[Parameters]\"\n        }\n    }\n}\n\nautocomplete: set bg {\n    [{\n        \"DynamicDesc\": ({ fid.list.stopped }),\n        \"ListView\": true\n    }]\n}\n\nautocomplete: set fg {\n    [{\n        \"DynamicDesc\": ({ fid.list.stopped }),\n        \"ListView\": true\n    }]\n}\n\nautocomplete: set fid-kill {\n    [{\n        \"DynamicDesc\": ({ fid.list.all }),\n        \"ListView\": true,\n        \"AllowMultiple\": true\n    }]\n}\n\nautocomplete: set murex-package {\n    [{\n        \"Flags\": [ \"install\", \"update\", \"import\", \"enable\", \"disable\", \"reload\", \"status\" ],\n        \"FlagValues\": {\n            \"import\": [{\n                \"IncFiles\": true\n            }]\n        }\n    }]\n}")
}
