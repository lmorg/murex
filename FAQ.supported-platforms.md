# Supported Platforms

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux (AMD64)

* ArchLinux

Ths is the primary dev platform so all features will work.

* Debian

There is a supported Docker container included in this project. All
regression tests pass.

## OS X / Darwin

Untested but should compile and have similar support to FreeBSD, which
is a tested platform. I don't currently have access to an Mac and don't
have any plans on building a Darwin VM any time soon so I would be
looking towards the community for any testing and bug fixes on OS X.

## FreeBSD (AMD64)

_murex_ has been tested inside a 10.3-RELEASE jail. All regression tests
pass however there is a bug where some forked processes don't use the
entire TTY screen while most others do (specifically `top`).

Aside that, FreeBSD support is considered very good.

## OpenBSD (AMD64)

Tested. Regression tests cannot be run because of `timeout` dependency
however _murex_ does compile and run well - seems fully functional from
a functional test.

## NetBSD

Untested but should compile. I will be creating a test environment for
NetBSD soon.

## Windows (AMD64)

_murex_ does compile but a few features don't work which should:
* forked processes don't grab user input
* autocompleter doesn't pick up EXEs in %PATH%

...and a few features are unsupported due to a lack of support on the OS
platform itself:
* Windows doesn't have `man` pages so _murex_ cannot automatically offer
autocompletion suggestions for a commands supported flags

For these reasons, Windows support will be classed as experimental until
either I get POSIX support to a level that I am happy with, or other
developers are happy to submit pull requests.

## Plan 9

Not currently supported. There are a few differences in the `syscall`
package which would lead me to believe that _murex_ will not compile. I
do eventually plan on supporting Plan 9 so if you do use that OS then
keep checking back as support may be added soon.

## Other CPU architectures

While there isn't any CPU specific code in _murex_ below is a breakdown
of state of support and testing on alternative architectures:

* 386

Untested but should function the same as AMD64. 64 bit data types are
rarely used so performance should be roughly equal. If you do still use
32 bit platforms and notice any issues then please raise an issue ticket
and I will investigate.

* ARM

Untested currently but Linux ARM testing is expected to be undertaken
soon

* PPC

Untested however the Go compiler does support various PPC architectures.

* MIPS

Untested however the Go compiler does support various MIPS architectures.

* SPARC

Unsupported.