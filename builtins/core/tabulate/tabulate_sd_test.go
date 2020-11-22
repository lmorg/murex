package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inSd = `Query or send control commands to the systemd manager.

  -h --help           Show this help
     --version        Show package version
     --system         Connect to system manager
     --user           Connect to user service manager
  -H --host=[USER@]HOST
                      Operate on remote host
  -M --machine=CONTAINER
                      Operate on local container
  -t --type=TYPE      List units of a particular type
     --state=STATE    List units with particular LOAD or SUB or ACTIVE state
  -p --property=NAME  Show only properties by this name
  -a --all            Show all properties/all units currently in memory,
                      including dead/empty ones. To list all units installed on
                      the system, use the 'list-unit-files' command instead.
     --failed         Same as --state=failed
  -l --full           Don't ellipsize unit names on output
  -r --recursive      Show unit list of host and local containers
     --reverse        Show reverse dependencies with 'list-dependencies'
     --job-mode=MODE  Specify how to deal with already queued jobs, when
                      queueing a new job
     --show-types     When showing sockets, explicitly show their type
     --value          When showing properties, only print the value
  -i --ignore-inhibitors
                      When shutting down or sleeping, ignore inhibitors
     --kill-who=WHO   Who to send signal to
  -s --signal=SIGNAL  Which signal to send
     --now            Start or stop unit in addition to enabling or disabling it
     --dry-run        Only print what would be done
  -q --quiet          Suppress output
     --wait           For (re)start, wait until service stopped again
     --no-block       Do not wait until operation finished
     --no-wall        Don't send wall message before halt/power-off/reboot
     --no-reload      Don't reload daemon after en-/dis-abling unit files
     --no-legend      Do not print a legend (column headers and hints)
     --no-pager       Do not pipe output into a pager
     --no-ask-password
                      Do not ask for system passwords
     --global         Enable/disable/mask unit files globally
     --runtime        Enable/disable/mask unit files temporarily until next
                      reboot
  -f --force          When enabling unit files, override existing symlinks
                      When shutting down, execute action immediately
     --preset-mode=   Apply only enable, only disable, or all presets
     --root=PATH      Enable/disable/mask unit files in the specified root
                      directory
  -n --lines=INTEGER  Number of journal entries to show
  -o --output=STRING  Change journal output mode (short, short-precise,
                             short-iso, short-iso-precise, short-full,
                             short-monotonic, short-unix,
                             verbose, export, json, json-pretty, json-sse, cat)
     --firmware-setup Tell the firmware to show the setup menu on next boot
     --plain          Print unit dependencies as a list instead of a tree`

	jsonSd = `{"--all":"Show all properties/all units currently in memory, including dead/empty ones. To list all units installed on the system, use the 'list-unit-files' command instead.","--dry-run":"Only print what would be done","--failed":"Same as --state=failed","--force":"When enabling unit files, override existing symlinks When shutting down, execute action immediately","--full":"Don't ellipsize unit names on output","--global":"Enable/disable/mask unit files globally","--help":"Show this help","--job-mode=":"(args: MODE) Specify how to deal with already queued jobs, when queueing a new job","--kill-who=":"(args: WHO) Who to send signal to","--lines=":"(args: INTEGER) Number of journal entries to show","--no-block":"Do not wait until operation finished","--no-legend":"Do not print a legend (column headers and hints)","--no-pager":"Do not pipe output into a pager --no-ask-password Do not ask for system passwords","--no-reload":"Don't reload daemon after en-/dis-abling unit files","--no-wall":"Don't send wall message before halt/power-off/reboot","--now":"Start or stop unit in addition to enabling or disabling it","--output=":"(args: STRING) Change journal output mode (short, short-precise, short-iso, short-iso-precise, short-full, short-monotonic, short-unix, verbose, export, json, json-pretty, json-sse, cat) --firmware-setup Tell the firmware to show the setup menu on next boot","--plain":"Print unit dependencies as a list instead of a tree","--preset-mode=":"Apply only enable, only disable, or all presets","--property=":"(args: NAME) Show only properties by this name","--quiet":"Suppress output","--recursive":"Show unit list of host and local containers","--reverse":"Show reverse dependencies with 'list-dependencies'","--root=":"(args: PATH) Enable/disable/mask unit files in the specified root directory","--runtime":"Enable/disable/mask unit files temporarily until next reboot","--show-types":"When showing sockets, explicitly show their type","--signal=":"(args: SIGNAL) Which signal to send","--state=":"(args: STATE) List units with particular LOAD or SUB or ACTIVE state","--system":"Connect to system manager","--type=":"(args: TYPE) List units of a particular type","--user":"Connect to user service manager -H --host=[USER@]HOST Operate on remote host -M --machine=CONTAINER Operate on local container","--value":"When showing properties, only print the value -i --ignore-inhibitors When shutting down or sleeping, ignore inhibitors","--version":"Show package version","--wait":"For (re)start, wait until service stopped again","-a":"Show all properties/all units currently in memory, including dead/empty ones. To list all units installed on the system, use the 'list-unit-files' command instead.","-f":"When enabling unit files, override existing symlinks When shutting down, execute action immediately","-h":"Show this help","-l":"Don't ellipsize unit names on output","-n":"(args: INTEGER) Number of journal entries to show","-o":"(args: STRING) Change journal output mode (short, short-precise, short-iso, short-iso-precise, short-full, short-monotonic, short-unix, verbose, export, json, json-pretty, json-sse, cat) --firmware-setup Tell the firmware to show the setup menu on next boot","-p":"(args: NAME) Show only properties by this name","-q":"Suppress output","-r":"Show unit list of host and local containers","-s":"(args: SIGNAL) Which signal to send","-t":"(args: TYPE) List units of a particular type"}`
)

func TestTabulateSystemD(t *testing.T) {
	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inSd,
		types.Generic,
		[]string{fMap, fKeyIncHint, fSplitSpace, fColumnWraps},
		jsonSd,
		nil,
	)
}
