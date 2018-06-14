// +build linux

package defaults

func init() {
	murexProfile = append(murexProfile, `

alias ls=ls --color=auto
alias grep=grep --color=auto
alias egrep=egrep --color=auto

autocomplete set systemctl { [
	{
		"Dynamic": ({ systemctl --help -> regexp (f/(--.*?[= ])/) }),
		"Optional": true,
		"AllowMultiple": true
	},
	{
	    "Flags": [ "list-units", "list-sockets", "list-timers", "start", "stop", "reload", "restart", "try-restart", "reload-or-restart", "try-reload-or-restart", "isolate", "kill", "is-active", "is-failed", "status", "show", "cat", "set-property", "help", "reset-failed", "list-dependencies", "list-unit-files", "enable", "disable", "reenable", "preset", "preset-all", "is-enabled", "mask", "unmask", "link", "revert", "add-wants", "add-requires", "edit", "get-default", "set-default", "list-machines", "list-jobs", "cancel", "show-environment", "set-environment", "unset-environment", "import-environment", "daemon-reload", "daemon-reexec", "is-system-running", "default", "rescue", "emergency", "halt", "poweroff", "reboot", "kexec", "exit", "switch-root", "suspend", "hibernate", "hybrid-sleep" ]
	},
	{
		"Dynamic": ({ systemctl list-units --plain -> @[UNIT..^\$]re -> [ :0 ] }),
		#"AllowMultiple": true,
		"Optional": false
	}
] }
`)
}
