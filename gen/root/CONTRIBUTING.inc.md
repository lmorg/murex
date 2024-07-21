{{ if env "DOCGEN_TARGET=vuepress" }}
{{ if env "DOCGEN_TARGET=ignore-prefix" }}
### {{ end }}icon: handshake-angle

---
{{ end }}<h1>Contributing to Murex</h1>

Murex is community project. We gratefully accept contributions.

{{ if env "DOCGEN_TARGET=" }}<h2>Table of Contents</h2>

<div id="toc">

- [Ways To Contribute](#ways-to-contribute)
  - [Writing or Updating Documentation](#writing-or-updating-documentation)
  - [Writing Integrations](#writing-integrations)
  - [Raising Bug Reports or Feature Requests](#raising-bug-reports-or-feature-requests)
  - [Committing Code](#committing-code)
  - [Blogging](#blogging)
- [Raising Pull Requests](#raising-pull-requests)
- [Etiquette](#etiquette)
- [Licensing](#licensing)

</div>
{{ end }}
## Ways To Contribute

You don't have to be a software developer to support this project, there are
multiple ways you can contribute to Murex. Listed below are some examples of
areas we are looking for support:

### Writing or Updating Documentation

Writing documentation is probably the dullest part of any project yet arguably
one of the most important. The vast majority of the documentation for Murex
has been knocked out in a hurry by one guy - a guy who's core weaknesses might
include "writing". So there is a considerable room for improvement to be made.

You don't even have to write any documentation from scratch. If you notice a
grammatical error, spelling mistakes or even just a confusing sentence, then
please do raise a pull request.

All documentation is written in markdown. Even the website is HTML generated
from the original markdown documents. And all markdown documents are themselves
generated from files with the extension `_doc.yaml`. These template files are
the backbone of the documentation's pseudo-CMS.

If you are unsure where to find a document, you can either `grep` the source
repository for a known phrase, or ask in the [Github discussions](https://github.com/lmorg/murex/discussions)
group.

### Writing Integrations

Murex is a smart shell - it parses man pages for command line flags. This helps
to reduce the impact of missing custom autocompletions. However sometimes it is
in escapable to need a custom completion. Maybe that is because the options are
atypical (like how `kill` should list PIDs with their application adjacent).
Sometimes it might be because other integrations are required, such as events,
aliases or functions defined.

One of the core tenets of Murex is that it's out-of-the-box experience should
already be excellent and any customization that happens after is for
personalization rather than because the default experience is lacking in some
way. So Murex has made it very easy for you to contribute.

In the root of the project resides the [integrations directory](https://github.com/lmorg/murex/tree/master/integrations).
Each file in there is compiled into Murex, assuming the following condition is
met.

The file must be named `xxx_platform.mx` where:

* `xxx` refers to the name of the integration
* `_platform` refers to either: `any` (runs on every platform), `posix` (Linux or UNIX only) or `linux`, `darwin` (macOS), `freebsd`, `openbsd`, `netbsd`, `dragonfly` (DragonflyBSD), `solaris`, `plan9`, `windows`.
* The file extension most be `.mx`.

Example files also exist in `/integrations` to help you get started.

### Raising Bug Reports or Feature Requests

It might seem counterintuitive that raising issues is a form of contribution
but without feedback Murex cannot provide the out-of-box experience it aims to.
So bug reports and feature requests do help.

These can be raised on our [Github issue tracker](github.com/lmorg/murex/issues).

### Committing Code

Murex is written in a language called Go. Not a lot of Murex's code is well
documented however that was is document can be found in the [API section]({{if env "DOCGEN_TARGET="}}/docs{{end}}/apis) of the
user guide.

### Blogging

If you want to share some shell tips, be it for Murex or any of the more
traditional shells like Bash, then we welcome them as short articles for the
[blog section]({{if env "DOCGEN_TARGET="}}/docs{{end}}/blog). We want the website to be a
valuable resource for shells and scripting regardless for the platform and
language.

Articles must be in markdown format and credit will be attributed with links to
your Github account, Twitter or other social platforms and/or promotions.

## Raising Pull Requests

Pull requests should be raised against the `develop` branch. This allows us to
stage and test changes before releasing them to everyone.

## Etiquette

Murex is a community project and as such, everyone is entitled to an opinion
and opinions might differ. With that in mind, please be patient if discussions
happen regarding your contributions. All contributions are welcome however we
do also need to ensure that Murex has focus and a consistent design. This means
sometimes a conversation might be needed to work that contribution into the
wider, holistic, design of the shell.

This should not put anyone off contributing. However if you are unsure about
whether a contribution fits, then you're welcome to [start a discussion ](https://github.com/lmorg/murex/discussions) first.

## Licensing

By contributing, you agree to license your code under the same license as the
existing source code (see the [LICENSE](https://github.com/lmorg/murex/blob/master/LICENSE) file) and that @lmorg has the
right to relicense Murex under an alternative _open source_ license in the,
future should the need arise.

Murex will always be open source software. It wouldn't exist without open
source tooling and therefore it will always serve to enhance the open source
community.
