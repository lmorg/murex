package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inSd = `systemctl [OPTIONS...] COMMAND ...

Query or send control commands to the system manager.

Unit Commands:
  list-units [PATTERN...]             List units currently in memory
  list-sockets [PATTERN...]           List socket units currently in memory,
                                      ordered by address
  list-timers [PATTERN...]            List timer units currently in memory,
                                      ordered by next elapse
  start UNIT...                       Start (activate) one or more units
  stop UNIT...                        Stop (deactivate) one or more units
  reload UNIT...                      Reload one or more units
  restart UNIT...                     Start or restart one or more units
  try-restart UNIT...                 Restart one or more units if active
  reload-or-restart UNIT...           Reload one or more units if possible,
                                      otherwise start or restart
  try-reload-or-restart UNIT...       If active, reload one or more units,
                                      if supported, otherwise restart
  isolate UNIT                        Start one unit and stop all others
  kill UNIT...                        Send signal to processes of a unit
  clean UNIT...                       Clean runtime, cache, state, logs or
                                      configuration of unit
  freeze PATTERN...                   Freeze execution of unit processes
  thaw PATTERN...                     Resume execution of a frozen unit
  is-active PATTERN...                Check whether units are active
  is-failed PATTERN...                Check whether units are failed
  status [PATTERN...|PID...]          Show runtime status of one or more units
  show [PATTERN...|JOB...]            Show properties of one or more
                                      units/jobs or the manager
  cat PATTERN...                      Show files and drop-ins of specified units
  set-property UNIT PROPERTY=VALUE... Sets one or more properties of a unit
  help PATTERN...|PID...              Show manual for one or more units
  reset-failed [PATTERN...]           Reset failed state for all, one, or more
                                      units
  list-dependencies [UNIT...]         Recursively show units which are required
                                      or wanted by the units or by which those
                                      units are required or wanted
Unit File Commands:
  list-unit-files [PATTERN...]        List installed unit files
  enable [UNIT...|PATH...]            Enable one or more unit files
  disable UNIT...                     Disable one or more unit files
  reenable UNIT...                    Reenable one or more unit files
  preset UNIT...                      Enable/disable one or more unit files
                                      based on preset configuration
  preset-all                          Enable/disable all unit files based on
                                      preset configuration
  is-enabled UNIT...                  Check whether unit files are enabled
  mask UNIT...                        Mask one or more units
  unmask UNIT...                      Unmask one or more units
  link PATH...                        Link one or more units files into
                                      the search path
  revert UNIT...                      Revert one or more unit files to vendor
                                      version
  add-wants TARGET UNIT...            Add 'Wants' dependency for the target
                                      on specified one or more units
  add-requires TARGET UNIT...         Add 'Requires' dependency for the target
                                      on specified one or more units
  edit UNIT...                        Edit one or more unit files
  get-default                         Get the name of the default target
  set-default TARGET                  Set the default target

Machine Commands:
  list-machines [PATTERN...]          List local containers and host

Job Commands:
  list-jobs [PATTERN...]              List jobs
  cancel [JOB...]                     Cancel all, one, or more jobs

Environment Commands:
  show-environment                    Dump environment
  set-environment VARIABLE=VALUE...   Set one or more environment variables
  unset-environment VARIABLE...       Unset one or more environment variables
  import-environment [VARIABLE...]    Import all or some environment variables

Manager State Commands:
  daemon-reload                       Reload systemd manager configuration
  daemon-reexec                       Reexecute systemd manager
  log-level [LEVEL]                   Get/set logging threshold for manager
  log-target [TARGET]                 Get/set logging target for manager
  service-watchdogs [BOOL]            Get/set service watchdog state

System Commands:
  is-system-running                   Check whether system is fully running
  default                             Enter system default mode
  rescue                              Enter system rescue mode
  emergency                           Enter system emergency mode
  halt                                Shut down and halt the system
  poweroff                            Shut down and power-off the system
  reboot                              Shut down and reboot the system
  kexec                               Shut down and reboot the system with kexec
  exit [EXIT_CODE]                    Request user instance or container exit
  switch-root ROOT [INIT]             Change to a different root file system
  suspend                             Suspend the system
  hibernate                           Hibernate the system
  hybrid-sleep                        Hibernate and suspend the system
  suspend-then-hibernate              Suspend the system, wake after a period of
                                      time, and hibernate
Options:
  -h --help              Show this help
     --version           Show package version
     --system            Connect to system manager
     --user              Connect to user service manager
  -H --host=[USER@]HOST  Operate on remote host
  -M --machine=CONTAINER Operate on a local container
  -t --type=TYPE         List units of a particular type
     --state=STATE       List units with particular LOAD or SUB or ACTIVE state
     --failed            Shortcut for --state=failed
  -p --property=NAME     Show only properties by this name
  -P NAME                Equivalent to --value --property=NAME
  -a --all               Show all properties/all units currently in memory,
                         including dead/empty ones. To list all units installed
                         on the system, use 'list-unit-files' instead.
  -l --full              Don't ellipsize unit names on output
  -r --recursive         Show unit list of host and local containers
     --reverse           Show reverse dependencies with 'list-dependencies'
     --with-dependencies Show unit dependencies with 'status', 'cat',
                         'list-units', and 'list-unit-files'.
     --job-mode=MODE     Specify how to deal with already queued jobs, when
                         queueing a new job
  -T --show-transaction  When enqueuing a unit job, show full transaction
     --show-types        When showing sockets, explicitly show their type
     --value             When showing properties, only print the value
  -i --ignore-inhibitors When shutting down or sleeping, ignore inhibitors
     --kill-who=WHO      Whom to send signal to
  -s --signal=SIGNAL     Which signal to send
     --what=RESOURCES    Which types of resources to remove
     --now               Start or stop unit after enabling or disabling it
     --dry-run           Only print what would be done
                         Currently supported by verbs: halt, poweroff, reboot,
                             kexec, suspend, hibernate, suspend-then-hibernate,
                             hybrid-sleep, default, rescue, emergency, and exit.
  -q --quiet             Suppress output
     --wait              For (re)start, wait until service stopped again
                         For is-system-running, wait until startup is completed
     --no-block          Do not wait until operation finished
     --no-wall           Don't send wall message before halt/power-off/reboot
     --no-reload         Don't reload daemon after en-/dis-abling unit files
     --no-legend         Do not print a legend (column headers and hints)
     --no-pager          Do not pipe output into a pager
     --no-ask-password   Do not ask for system passwords
     --global            Enable/disable/mask unit files globally
     --runtime           Enable/disable/mask unit files temporarily until next
                         reboot
  -f --force             When enabling unit files, override existing symlinks
                         When shutting down, execute action immediately
     --preset-mode=      Apply only enable, only disable, or all presets
     --root=PATH         Enable/disable/mask unit files in the specified root
                         directory
  -n --lines=INTEGER     Number of journal entries to show
  -o --output=STRING     Change journal output mode (short, short-precise,
                             short-iso, short-iso-precise, short-full,
                             short-monotonic, short-unix,
                             verbose, export, json, json-pretty, json-sse, cat)
     --firmware-setup    Tell the firmware to show the setup menu on next boot
     --boot-loader-menu=TIME
                         Boot into boot loader menu on next boot
     --boot-loader-entry=NAME
                         Boot into a specific boot loader entry on next boot
     --plain             Print unit dependencies as a list instead of a tree

See the systemctl(1) man page for details.
`

	jsonSd = `{"--all":"Show all properties/all units currently in memory, including dead/empty ones. To list all units installed on the system, use 'list-unit-files' instead.","--boot-loader-entry=":"(args: NAME) Boot into a specific boot loader entry on next boot","--boot-loader-menu=":"(args: TIME) Boot into boot loader menu on next boot","--dry-run":"Only print what would be done Currently supported by verbs: halt, poweroff, reboot, kexec, suspend, hibernate, suspend-then-hibernate, hybrid-sleep, default, rescue, emergency, and exit.","--failed":"Shortcut for --state=failed","--firmware-setup":"Tell the firmware to show the setup menu on next boot","--force":"When enabling unit files, override existing symlinks When shutting down, execute action immediately","--full":"Don't ellipsize unit names on output","--global":"Enable/disable/mask unit files globally","--help":"Show this help","--host=":"(args: [USER@]HOST) Operate on remote host","--ignore-inhibitors":"(args: When shutting down or sleeping, ignore inhibitors) ","--job-mode=":"(args: MODE) Specify how to deal with already queued jobs, when queueing a new job","--kill-who=":"(args: WHO) Whom to send signal to","--lines=":"(args: INTEGER) Number of journal entries to show","--machine=":"(args: Operate on a local container) ","--no-ask-password":"Do not ask for system passwords","--no-block":"Do not wait until operation finished","--no-legend":"Do not print a legend (column headers and hints)","--no-pager":"Do not pipe output into a pager","--no-reload":"Don't reload daemon after en-/dis-abling unit files","--no-wall":"Don't send wall message before halt/power-off/reboot","--now":"Start or stop unit after enabling or disabling it","--output=":"(args: STRING) Change journal output mode (short, short-precise, short-iso, short-iso-precise, short-full, short-monotonic, short-unix, verbose, export, json, json-pretty, json-sse, cat)","--plain":"Print unit dependencies as a list instead of a tree","--preset-mode=":"Apply only enable, only disable, or all presets","--property=":"(args: NAME) Show only properties by this name","--quiet":"Suppress output","--recursive":"Show unit list of host and local containers","--reverse":"Show reverse dependencies with 'list-dependencies'","--root=":"(args: PATH) Enable/disable/mask unit files in the specified root directory","--runtime":"Enable/disable/mask unit files temporarily until next reboot","--show-transaction":"When enqueuing a unit job, show full transaction","--show-types":"When showing sockets, explicitly show their type","--signal=":"(args: SIGNAL) Which signal to send","--state=":"(args: STATE) List units with particular LOAD or SUB or ACTIVE state","--system":"Connect to system manager","--type=":"(args: TYPE) List units of a particular type","--user":"Connect to user service manager","--value":"When showing properties, only print the value","--version":"Show package version","--wait":"For (re)start, wait until service stopped again For is-system-running, wait until startup is completed","--what=":"(args: RESOURCES) Which types of resources to remove","--with-dependencies":"(args: Show unit dependencies with 'status', 'cat',) 'list-units', and 'list-unit-files'.","-H":"(args: [USER@]HOST) Operate on remote host","-M":"(args: Operate on a local container) ","-P":"(args: NAME) Equivalent to --value --property=NAME","-T":"When enqueuing a unit job, show full transaction","-a":"Show all properties/all units currently in memory, including dead/empty ones. To list all units installed on the system, use 'list-unit-files' instead.","-f":"When enabling unit files, override existing symlinks When shutting down, execute action immediately","-h":"Show this help","-i":"(args: When shutting down or sleeping, ignore inhibitors) ","-l":"Don't ellipsize unit names on output","-n":"(args: INTEGER) Number of journal entries to show","-o":"(args: STRING) Change journal output mode (short, short-precise, short-iso, short-iso-precise, short-full, short-monotonic, short-unix, verbose, export, json, json-pretty, json-sse, cat)","-p":"(args: NAME) Show only properties by this name","-q":"Suppress output","-r":"Show unit list of host and local containers","-s":"(args: SIGNAL) Which signal to send","-t":"(args: TYPE) List units of a particular type","add-requires":"(args: TARGET UNIT...) Add 'Requires' dependency for the target on specified one or more units","add-wants":"(args: TARGET UNIT...) Add 'Wants' dependency for the target on specified one or more units","cancel":"(args: [JOB...]) Cancel all, one, or more jobs","cat":"(args: PATTERN...) Show files and drop-ins of specified units set-property UNIT PROPERTY=VALUE... Sets one or more properties of a unit","clean":"(args: UNIT...) Clean runtime, cache, state, logs or configuration of unit","daemon-reexec":"Reexecute systemd manager","daemon-reload":"Reload systemd manager configuration","default":"Enter system default mode","disable":"(args: UNIT...) Disable one or more unit files","edit":"(args: UNIT...) Edit one or more unit files","emergency":"Enter system emergency mode","enable":"(args: [UNIT...|PATH...]) Enable one or more unit files","exit":"(args: [EXIT_CODE]) Request user instance or container exit","freeze":"(args: PATTERN...) Freeze execution of unit processes","get-default":"Get the name of the default target","halt":"Shut down and halt the system","help":"(args: PATTERN...|PID...) Show manual for one or more units","hibernate":"Hibernate the system","hybrid-sleep":"Hibernate and suspend the system","import-environment":"(args: [VARIABLE...]) Import all or some environment variables","is-active":"(args: PATTERN...) Check whether units are active","is-enabled":"(args: UNIT...) Check whether unit files are enabled","is-failed":"(args: PATTERN...) Check whether units are failed","is-system-running":"Check whether system is fully running","isolate":"(args: UNIT) Start one unit and stop all others","kexec":"Shut down and reboot the system with kexec","kill":"(args: UNIT...) Send signal to processes of a unit","link":"(args: PATH...) Link one or more units files into the search path","list-dependencies":"(args: [UNIT...]) Recursively show units which are required or wanted by the units or by which those units are required or wanted","list-jobs":"(args: [PATTERN...]) List jobs","list-machines":"(args: [PATTERN...]) List local containers and host","list-sockets":"(args: [PATTERN...]) List socket units currently in memory, ordered by address","list-timers":"(args: [PATTERN...]) List timer units currently in memory, ordered by next elapse","list-unit-files":"(args: [PATTERN...]) List installed unit files","list-units":"(args: [PATTERN...]) List units currently in memory","log-level":"(args: [LEVEL]) Get/set logging threshold for manager","log-target":"(args: [TARGET]) Get/set logging target for manager","mask":"(args: UNIT...) Mask one or more units","poweroff":"Shut down and power-off the system","preset":"(args: UNIT...) Enable/disable one or more unit files based on preset configuration","preset-all":"Enable/disable all unit files based on preset configuration","reboot":"Shut down and reboot the system","reenable":"(args: UNIT...) Reenable one or more unit files","reload":"(args: UNIT...) Reload one or more units","reload-or-restart":"(args: UNIT...) Reload one or more units if possible, otherwise start or restart","rescue":"Enter system rescue mode","reset-failed":"(args: [PATTERN...]) Reset failed state for all, one, or more units","restart":"(args: UNIT...) Start or restart one or more units","revert":"(args: UNIT...) Revert one or more unit files to vendor version","service-watchdogs":"(args: [BOOL]) Get/set service watchdog state","set-default":"(args: TARGET) Set the default target","set-environment":"(args: VARIABLE=VALUE...) Set one or more environment variables","show":"(args: [PATTERN...|JOB...]) Show properties of one or more units/jobs or the manager","show-environment":"Dump environment","start":"(args: UNIT...) Start (activate) one or more units","status":"(args: [PATTERN...|PID...]) Show runtime status of one or more units","stop":"(args: UNIT...) Stop (deactivate) one or more units","suspend":"Suspend the system","suspend-then-hibernate":"Suspend the system, wake after a period of time, and hibernate","switch-root":"(args: [INIT]) Change to a different root file system","thaw":"(args: PATTERN...) Resume execution of a frozen unit","try-reload-or-restart":"(args: UNIT...) If active, reload one or more units, if supported, otherwise restart","try-restart":"(args: UNIT...) Restart one or more units if active","unmask":"(args: UNIT...) Unmask one or more units","unset-environment":"(args: VARIABLE...) Unset one or more environment variables"}`
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
