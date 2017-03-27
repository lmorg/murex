# Language Guide: `Builtins` (functions and methods)

I haven't written much documentation on this at the moment, however
there are three places you can look:

1. [GUIDE.control-structures.md], which contains builtins required for
building logic.

2. Go source code: under the `lang/builtins` path of the shell source
source code project files is several packages, each hosting a few
grouped builtin functions and methods. Each package will include a
README.md file with a basic summary of what the package is used for and
the importance of including it should you decide to recompile the shell.

3. From the shell itself: run `builtins` to list the builtin functions
and methods.

Eventually I will better document this, but please bare in mind this
this shell is still currently in ALPHA (essentially an MVP) so
documentation will follow as and when features are unlikely to change.

The eventual plan is to write the function creator APIs in such a way
that the language to become self-documenting (the config settings APIs
are already built this way)
