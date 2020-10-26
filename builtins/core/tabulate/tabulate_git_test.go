package tabulate

import (
	"testing"

	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/test"
)

var (
	inGit = `These are common Git commands used in various situations:

start a working area (see also: git help tutorial)
   clone      Clone a repository into a new directory
   init       Create an empty Git repository or reinitialize an existing one

work on the current change (see also: git help everyday)
   add        Add file contents to the index
   mv         Move or rename a file, a directory, or a symlink
   reset      Reset current HEAD to the specified state
   rm         Remove files from the working tree and from the index

examine the history and state (see also: git help revisions)
   bisect     Use binary search to find the commit that introduced a bug
   grep       Print lines matching a pattern
   log        Show commit logs
   show       Show various types of objects
   status     Show the working tree status

grow, mark and tweak your common history
   branch     List, create, or delete branches
   checkout   Switch branches or restore working tree files
   commit     Record changes to the repository
   diff       Show changes between commits, commit and working tree, etc
   merge      Join two or more development histories together
   rebase     Reapply commits on top of another base tip
   tag        Create, list, delete or verify a tag object signed with GPG

collaborate (see also: git help workflows)
   fetch      Download objects and refs from another repository
   pull       Fetch from and integrate with another repository or a local branch
   push       Update remote refs along with associated objects

'git help -a' and 'git help -g' list available subcommands and some
concept guides. See 'git help <command>' or 'git help <concept>'
to read about a specific subcommand or concept.`

	csvGit = `clone,Clone a repository into a new directory
init,Create an empty Git repository or reinitialize an existing one
add,Add file contents to the index
mv,"Move or rename a file, a directory, or a symlink"
reset,Reset current HEAD to the specified state
rm,Remove files from the working tree and from the index
bisect,Use binary search to find the commit that introduced a bug
grep,Print lines matching a pattern
log,Show commit logs
show,Show various types of objects
status,Show the working tree status
branch,"List, create, or delete branches"
checkout,Switch branches or restore working tree files
commit,Record changes to the repository
diff,"Show changes between commits, commit and working tree, etc"
merge,Join two or more development histories together
rebase,Reapply commits on top of another base tip
tag,"Create, list, delete or verify a tag object signed with GPG"
fetch,Download objects and refs from another repository
pull,Fetch from and integrate with another repository or a local branch
push,Update remote refs along with associated objects
`

	jsonGit = `{"add":"Add file contents to the index","bisect":"Use binary search to find the commit that introduced a bug","branch":"List, create, or delete branches","checkout":"Switch branches or restore working tree files","clone":"Clone a repository into a new directory","commit":"Record changes to the repository","diff":"Show changes between commits, commit and working tree, etc","fetch":"Download objects and refs from another repository","grep":"Print lines matching a pattern","init":"Create an empty Git repository or reinitialize an existing one","log":"Show commit logs","merge":"Join two or more development histories together","mv":"Move or rename a file, a directory, or a symlink","pull":"Fetch from and integrate with another repository or a local branch","push":"Update remote refs along with associated objects","rebase":"Reapply commits on top of another base tip","reset":"Reset current HEAD to the specified state","rm":"Remove files from the working tree and from the index","show":"Show various types of objects","status":"Show the working tree status","tag":"Create, list, delete or verify a tag object signed with GPG"}`
)

func TestTabulateGit(t *testing.T) {
	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inGit,
		types.Generic,
		[]string{},
		csvGit,
		nil,
	)

	test.RunMethodTest(t,
		cmdTabulate, "tabulate",
		inGit,
		types.Generic,
		[]string{fMap},
		jsonGit,
		nil,
	)
}
