package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inMake = `Usage: make [options] [target] ...
Options:
  -b, -m                      Ignored for compatibility.
  -B, --always-make           Unconditionally make all targets.
  -C DIRECTORY, --directory=DIRECTORY
                              Change to DIRECTORY before doing anything.
  -d                          Print lots of debugging information.
  --debug[=FLAGS]             Print various types of debugging information.
  -e, --environment-overrides
                              Environment variables override makefiles.
  -E STRING, --eval=STRING    Evaluate STRING as a makefile statement.
  -f FILE, --file=FILE, --makefile=FILE
                              Read FILE as a makefile.
  -h, --help                  Print this message and exit.
  -i, --ignore-errors         Ignore errors from recipes.
  -I DIRECTORY, --include-dir=DIRECTORY
                              Search DIRECTORY for included makefiles.
  -j [N], --jobs[=N]          Allow N jobs at once; infinite jobs with no arg.
  -k, --keep-going            Keep going when some targets can't be made.
  -l [N], --load-average[=N], --max-load[=N]
                              Don't start multiple jobs unless load is below N.
  -L, --check-symlink-times   Use the latest mtime between symlinks and target.
  -n, --just-print, --dry-run, --recon
                              Don't actually run any recipe; just print them.
  -o FILE, --old-file=FILE, --assume-old=FILE
                              Consider FILE to be very old and don't remake it.
  -O[TYPE], --output-sync[=TYPE]
                              Synchronize output of parallel jobs by TYPE.
  -p, --print-data-base       Print make's internal database.
  -q, --question              Run no recipe; exit status says if up to date.
  -r, --no-builtin-rules      Disable the built-in implicit rules.
  -R, --no-builtin-variables  Disable the built-in variable settings.
  -s, --silent, --quiet       Don't echo recipes.
  --no-silent                 Echo recipes (disable --silent mode).
  -S, --no-keep-going, --stop
                              Turns off -k.
  -t, --touch                 Touch targets instead of remaking them.
  --trace                     Print tracing information.
  -v, --version               Print the version number of make and exit.
  -w, --print-directory       Print the current directory.
  --no-print-directory        Turn off -w, even if it was turned on implicitly.
  -W FILE, --what-if=FILE, --new-file=FILE, --assume-new=FILE
                              Consider FILE to be infinitely new.
  --warn-undefined-variables  Warn when an undefined variable is referenced.

This program built for x86_64-pc-linux-gnu
Report bugs to <bug-make@gnu.org>
`

	jsonMake = `{"--always-make":"Unconditionally make all targets.","--assume-new=":"(args: FILE) Consider FILE to be infinitely new.","--assume-old=":"(args: FILE) Consider FILE to be very old and don't remake it.","--check-symlink-times":"Use the latest mtime between symlinks and target.","--debug":"(args: [=FLAGS]) Print various types of debugging information.","--directory=":"(args: DIRECTORY) Change to DIRECTORY before doing anything.","--dry-run":"Don't actually run any recipe; just print them.","--environment-overrides":"Environment variables override makefiles.","--eval=":"(args: STRING) Evaluate STRING as a makefile statement.","--file=":"(args: FILE) Read FILE as a makefile.","--help":"Print this message and exit.","--ignore-errors":"Ignore errors from recipes.","--include-dir=":"(args: DIRECTORY) Search DIRECTORY for included makefiles.","--jobs":"(args: [=N]) Allow N jobs at once; infinite jobs with no arg.","--just-print":"Don't actually run any recipe; just print them.","--keep-going":"Keep going when some targets can't be made.","--load-average":"(args: [=N]) Don't start multiple jobs unless load is below N.","--makefile=":"(args: FILE) Read FILE as a makefile.","--max-load":"(args: [=N]) Don't start multiple jobs unless load is below N.","--new-file=":"(args: FILE) Consider FILE to be infinitely new.","--no-builtin-rules":"Disable the built-in implicit rules.","--no-builtin-variables":"Disable the built-in variable settings.","--no-keep-going":"Turns off -k.","--no-print-directory":"Turn off -w, even if it was turned on implicitly.","--no-silent":"Echo recipes (disable --silent mode).","--old-file=":"(args: FILE) Consider FILE to be very old and don't remake it.","--output-sync":"(args: [=TYPE]) Synchronize output of parallel jobs by TYPE.","--print-data-base":"Print make's internal database.","--print-directory":"Print the current directory.","--question":"Run no recipe; exit status says if up to date.","--quiet":"Don't echo recipes.","--recon":"Don't actually run any recipe; just print them.","--silent":"Don't echo recipes.","--stop":"Turns off -k.","--touch":"Touch targets instead of remaking them.","--trace":"Print tracing information.","--version":"Print the version number of make and exit.","--warn-undefined-variables":"Warn when an undefined variable is referenced.","--what-if=":"(args: FILE) Consider FILE to be infinitely new.","-B":"Unconditionally make all targets.","-C":"(args: DIRECTORY) Change to DIRECTORY before doing anything.","-E":"(args: STRING) Evaluate STRING as a makefile statement.","-I":"(args: DIRECTORY) Search DIRECTORY for included makefiles.","-L":"Use the latest mtime between symlinks and target.","-O":"(args: [=TYPE]) Synchronize output of parallel jobs by TYPE.","-R":"Disable the built-in variable settings.","-S":"Turns off -k.","-W":"(args: FILE) Consider FILE to be infinitely new.","-b":"Ignored for compatibility.","-d":"Print lots of debugging information.","-e":"Environment variables override makefiles.","-f":"(args: FILE) Read FILE as a makefile.","-h":"Print this message and exit.","-i":"Ignore errors from recipes.","-j":"(args: [=N]) Allow N jobs at once; infinite jobs with no arg.","-k":"Keep going when some targets can't be made.","-l":"(args: [=N]) Don't start multiple jobs unless load is below N.","-m":"Ignored for compatibility.","-n":"Don't actually run any recipe; just print them.","-o":"(args: FILE) Consider FILE to be very old and don't remake it.","-p":"Print make's internal database.","-q":"Run no recipe; exit status says if up to date.","-r":"Disable the built-in implicit rules.","-s":"Don't echo recipes.","-t":"Touch targets instead of remaking them.","-v":"Print the version number of make and exit.","-w":"Print the current directory."}`
)

func TestTabulateMake(t *testing.T) {
	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inMake,
		types.Generic,
		[]string{fMap, fColumnWraps, fSplitComma, fKeyIncHint},
		jsonMake,
		nil,
	)
}
