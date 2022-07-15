# _murex_ Shell Docs

## Change Log

Track new features, any breaking changes, and the release history here.

## Articles

### 15.07.2022 - [What's new in murex v2.9](changelog/v2.9.md)

This release sees a step up  


### 23.05.2022 - [What's new in murex v2.8](changelog/v2.8.md)

This release comes with a number of experimental but stable features that might eventually become standard practice. The features are there to use if you with but adjacent from the older code so there is zero risk in updating to this version.


### 15.05.2022 - [What's new in murex v2.7](changelog/v2.7.md)

This update has introduced another potential breaking change for your safety: zero length arrays now fail by default. Also errors inside subshells will cause the parent command to fail if ran inside a `try` or `trypipe` block.


### 26.02.2022 - [What's new in murex v2.6](changelog/v2.6.md)

This update has introduced a potential breaking change: variables now need to be defined before usage otherwise the commandline will fail. Read notes to learn how to disable this feature where needed. Also included in this release is the `select` command as part of the standard build.


### 12.02.2022 - [What's new in murex v2.5](changelog/v2.5.md)

This release introduces a number of new builtins, fixes some regression bugs and supercharges the `select` optional builtin (which I plan to include into the core builtins for non-Windows users in the next release).


### 09.12.2021 - [What's new in murex v2.4](changelog/v2.4.md)

This release introduces a strict mode for variables, new builtin, performance improvements, and better error messages; plus a potential breaking change


### 26.09.2021 - [What's new in murex v2.3](changelog/v2.3.md)

This release includes significant changes to the interactive terminal


### 21.06.2021 - [What's new in murex v2.2](changelog/v2.2.md)

This is mainly a bug fix release but it does include one breaking change for `config`. Please read for details.


### 30.04.2021 - [What's new in murex v2.1](changelog/v2.1.md)

This release comes with support for inlining SQL and some major bug fixes plus a breaking change for `config`. Please read for details.


### 17.04.2021 - [What's new in murex v2.0](changelog/v2.0.md)

This release comes with spellchecking, inlined images, smarter syntax completion and more

