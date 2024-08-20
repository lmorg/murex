<h1>Compatibility Commitment</h1>

Murex is committed to backwards compatibility. While we do want to continue to
grow and improve the shell, this will not come at the expense of long term
usability. 

<h2>Table of Contents</h2>

<div id="toc">

- [Our compatibility commitment](#our-compatibility-commitment)
  - [Features](#features)
  - [Breaking changes](#breaking-changes)
  - [Experimental features](#experimental-features)
  - [Development releases](#development-releases)
  - [Versioning](#versioning)

</div>



## Our compatibility commitment

You can consider Murex as stable. Many of us are using Murex as our primary
shell, some for years. There is already a non-trivial amount of code written
for Murex and that code will remain compatible for many years to come.

The following is a breakdown of Murex's development and backwards compatibility
commitment, in the hope it brings confidence to new users.

### Features

Any feature in the `master` branch (ie in a stable build) and thus published on
https://murex.rocks is considered stable.

Stable features are seldom removed (seriously, there are still parser rules for
undocumented but deprecated features from five years ago!).

If a feature is to be deprecated, the following steps are followed:
* first a deprecation notice is served in these docs
* after the next new [major](https://semver.org/) update, a warning will then
  be issued with the feature itself. When that feature is invoked, the warning
  will give notice of the deprecation
* after the following new major update, that feature will then be removed

This process is expected to take around two years. You do not need to regularly
follow the Github discussions to keep track of changes to the shell.

Features are only likely to be deprecated if they are unpopular.

### Breaking changes

A **breaking change** is considered to be any change that could affect any
Murex shell script already written.

Breaking changes _might_ happen outside of the feature deprecation life cycle
(described above) if:
* it is adding a new syntax rather than deprecating something (such as a new
  operator)
* and the breakages are edge cases as opposed to common (eg a bareword string
  that solely consists of the new operator is now parsed as an operator rather
  than a string)

Breaking changes will be published in the [changelog](https://murex.rocks/changelog/).

### Experimental features

Any feature marked as **EXPERIMENTAL** is subject to change at short notice.
Very few features end up as experimental and those that do might be because
they either introduce weird syntax that needs using in real situations to
determine their value, or might have some unresolved bugs and/or edge cases
that harm the overall UX. Generally features do not remain experimental for
long.

### Development releases

The `develop` branch is considered unstable. It is a place for contributors to
write and test code. This means all new features added to `develop` that hasn't
yet been released to `master` is considered experimental.

### Versioning

Murex releases roughly follows [semantic versioning](https://semver.org/).

## See Also

* [Contributing](/CONTRIBUTING.md):
  Guide to contributing to Murex
* [Install](/INSTALL.md):
  Installation instructions

<hr/>

This document was generated from [gen/root/compatibility_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/root/compatibility_doc.yaml).