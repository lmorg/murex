# Supported Operating Systems

The following is a list of platforms _murex_ has been tested on and the
level of support it had:

## Linux

**ArchLinux:**

Ths is my primary dev platform so all features will work.

**Debian:**

There is a supported Docker container included in this project. All
regression tests pass.

**Ubuntu:**

_murex_ has been tested and works well

## OS X / Darwin

Untested but I expect _murex_ to function. If you do encounter any
problems then please raise an issue.

## FreeBSD

_murex_ has been tested inside a 10.3-RELEASE AMD64 jail. All regression tests
pass however there is a bug where some forked processes don't use the
entire TTY screen while most others do (specifically `top`).

Aside that, FreeBSD support is considered very good.

## OpenBSD

Tested. Regression tests cannot be run because of `timeout` dependency
however _murex_ does compile and run well - seems fully functional from
a functional test.

## NetBSD

Untested but should compile.

## Windows

Windows support should be there however I have not had access to a
Windows environment for some time so there maybe some new bugs
introduced since. The project may even fail to compile due to code
refactoring that hasn't been correctly ported to the Windows _murex_
source. If you do have any issues then please raise an issue and I will
investigate.

Personally though, I would recommend running inside WSL anyway if just
for the GNU Coreutils support.

## Plan 9

Not currently supported. There are a few differences in the `syscall`
package which would lead me to believe that _murex_ will not compile.
That's not to say that I wont ever support Plan 9 but it's not a feature
I'm giving any priority to at present.

# Other CPU architectures

While there isn't any CPU specific code in _murex_ below is a breakdown
of state of support and testing on alternative architectures:

## 386

Untested but should function the same as AMD64. 64 bit data types are
rarely used so performance should be roughly equal. If you do still use
32 bit platforms and notice any issues then please raise an issue ticket
and I will investigate.

## ARM

Only tested on Linux but full compatibility was present.

## PPC

Untested however the Go compiler does support various PPC architectures
so _murex_ might work.

## MIPS

Untested however the Go compiler does support various MIPS architectures
so _murex_ might work.

## SPARC

Unsupported.