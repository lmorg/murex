// +build linux

package defaults

/*
   WARNING:
   --------

   This Go source file has been automatically generated from
   profile_linux.mx using docgen.

   Please do not manually edit this file because it will be automatically
   overwritten by the build pipeline. Instead please edit the aforementioned
   profile_linux.mx file located in the same directory.
*/

func init() {
	murexProfile = append(murexProfile, "alias ls=ls --color=auto\nalias grep=grep --color=auto\nalias egrep=egrep --color=auto\n\nautocomplete set systemctl { [\n\t{\n\t\t\"Dynamic\": ({ systemctl --help -> regexp (f/(--.*?[= ])/) }),\n\t\t\"Optional\": true,\n\t\t\"AllowMultiple\": true\n\t},\n\t{\n\t    \"Flags\": [ \"list-units\", \"list-sockets\", \"list-timers\", \"start\", \"stop\", \"reload\", \"restart\", \"try-restart\", \"reload-or-restart\", \"try-reload-or-restart\", \"isolate\", \"kill\", \"is-active\", \"is-failed\", \"status\", \"show\", \"cat\", \"set-property\", \"help\", \"reset-failed\", \"list-dependencies\", \"list-unit-files\", \"enable\", \"disable\", \"reenable\", \"preset\", \"preset-all\", \"is-enabled\", \"mask\", \"unmask\", \"link\", \"revert\", \"add-wants\", \"add-requires\", \"edit\", \"get-default\", \"set-default\", \"list-machines\", \"list-jobs\", \"cancel\", \"show-environment\", \"set-environment\", \"unset-environment\", \"import-environment\", \"daemon-reload\", \"daemon-reexec\", \"is-system-running\", \"default\", \"rescue\", \"emergency\", \"halt\", \"poweroff\", \"reboot\", \"kexec\", \"exit\", \"switch-root\", \"suspend\", \"hibernate\", \"hybrid-sleep\" ]\n\t},\n\t{\n\t\t\"Dynamic\": ({ systemctl list-units --plain -> @[UNIT..^\\$]re -> [ :0 ] }),\n\t\t#\"AllowMultiple\": true,\n\t\t\"Optional\": false\n\t}\n] }\n\n# I have a feeling this is old code that needs replacing with the OSX code \n#private kill-autocomplete {\n#    test define ps {\n#        \"ExitNum\": 0\n#    }\n#\n#    test define map {\n#        \"StdoutRegex\": (\\{(\".*?\":\".*?\",?)+\\})\n#    }\n#\n#    ps <test_ps> -A -o pid,cmd --no-headers -> set ps\n#    map <test_map> { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n#}\n#\n#test unit private kill-autocomplete {\n#    \"StdoutType\": \"json\",\n#    \"StdoutRegex\": \"\\\\{\\\\\\\"[0-9]+\\\\\\\":\\\\\\\".*?\\\\\\\"(,|)\\\\}\"\n#}\n\nfunction progress {\n    # Pulls the read progress of a Linux pid via /proc/$pid/fdinfo (only runs on Linux)\n\n    if { = `+\"`${os}`==`linux`\"+` } then {\n        params -> [ 1 ] -> set pid\n        \n        g <!null> /proc/$pid/fd/* -> regexp <!null> (f#/proc/[0-9]+/fd/([0-9]+)) -> foreach <!null> fd {\n            trypipe <!null> {\n                open /proc/$pid/fdinfo/$fd -> cast yaml -> [ pos ] -> set pos\n                readlink: /proc/$pid/fd/$fd -> set file\n                du -b $file -> [ :0 ] -> set int size\n                if { = size > 0 } then {\n                    = ($pos/$size)*100 -> set int percent\n                    out \"$percent% ($pos/$size) $file\"\n                }\n            }\n        }\n    }\n}\n\nautocomplete set progress {\n    [{\n        \"DynamicDesc\": ({\n            ps -A -o pid,cmd --no-headers -> set ps\n            map { $ps[:0] } { $ps -> regexp 'f/^[ 0-9]+ (.*)$' }\n        }),\n        \"ListView\": true\n    }]\n}\n\nconfig set shell stop-status-func {\n    progress $PID\n}\n")
}
