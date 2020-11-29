// +build ignore

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

	jsonMake = `{"--abort-on-container-exit":"Stops all containers if any container was stopped. Incompatible with -d.","--build":"Build images before starting containers.","--exit-code-from SERVICE":"Return the exit code of the selected service container. Implies --abort-on-container-exit.","--force-recreate":"Recreate containers even if their configuration and image haven't changed. Incompatible with --no-recreate.","--no-build":"Don't build an image, even if it's missing.","--no-color":"Produce monochrome output.","--no-deps":"Don't start linked services.","--no-recreate":"If containers already exist, don't recreate them. Incompatible with --force-recreate.","--no-start":"Don't start the services after creating them.","--remove-orphans":"Remove containers for services not defined in the Compose file","--scale SERVICE=NUM":"Scale SERVICE to NUM instances. Overrides the ` + "`" + `scale` + "`" + ` setting in the Compose file if present.","--timeout TIMEOUT":"Use this timeout in seconds for container shutdown when attached or when containers are already running. (default: 10)","-d":"Detached mode: Run containers in the background, print new container names. Incompatible with --abort-on-container-exit.","-t":"Use this timeout in seconds for container shutdown when attached or when containers are already running. (default: 10)"}`
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
