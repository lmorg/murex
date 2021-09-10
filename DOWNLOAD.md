# Download _murex_

Below are the instructions to download a pre-compiled binary via HTTP. If you
wish to install from source or use your preferred package manager, then please
refer to the [INSTALL](INSTALL.md) page for further instructions.

## HTTP Downloads

Below are the latest builds from the stable (`master`) branch.

All files are approximately 6 MB in size (aside the Plan 9 builds which are 3 MB).

* [murex-darwin-amd64.gz](https://murex.rocks/bin/latest/murex-darwin-amd64.gz)
* [murex-darwin-arm64.gz](https://murex.rocks/bin/latest/murex-darwin-amd64.gz)
* [murex-dragonfly-amd64.gz](https://murex.rocks/bin/latest/murex-dragonfly-amd64.gz)
* [murex-freebsd-386.gz](https://murex.rocks/bin/latest/murex-freebsd-386.gz)
* [murex-freebsd-amd64.gz](https://murex.rocks/bin/latest/murex-freebsd-amd64.gz)
* [murex-freebsd-arm.gz](https://murex.rocks/bin/latest/murex-freebsd-arm.gz)
* [murex-linux-386.gz](https://murex.rocks/bin/latest/murex-linux-386.gz)
* [murex-linux-amd64.gz](https://murex.rocks/bin/latest/murex-linux-amd64.gz)
* [murex-linux-arm.gz](https://murex.rocks/bin/latest/murex-linux-arm.gz)
* [murex-linux-arm64.gz](https://murex.rocks/bin/latest/murex-linux-arm64.gz)
* [murex-netbsd-386.gz](https://murex.rocks/bin/latest/murex-netbsd-386.gz)
* [murex-netbsd-amd64.gz](https://murex.rocks/bin/latest/murex-netbsd-amd64.gz)
* [murex-netbsd-arm.gz](https://murex.rocks/bin/latest/murex-netbsd-arm.gz)
* [murex-openbsd-386.gz](https://murex.rocks/bin/latest/murex-openbsd-386.gz)
* [murex-openbsd-amd64.gz](https://murex.rocks/bin/latest/murex-openbsd-amd64.gz)
* [murex-openbsd-arm.gz](https://murex.rocks/bin/latest/murex-openbsd-arm.gz)
* [murex-plan9-386.gz](https://murex.rocks/bin/latest/murex-plan9-386.gz)
* [murex-plan9-amd64.gz](https://murex.rocks/bin/latest/murex-plan9-amd64.gz)
* [murex-plan9-arm.gz](https://murex.rocks/bin/latest/murex-plan9-arm.gz)
* [murex-solaris-amd64.gz](https://murex.rocks/bin/latest/murex-solaris-amd64.gz) 
* [murex-windows-386.exe.zip](https://murex.rocks/bin/latest/murex-windows-386.exe.zip)
* [murex-windows-amd64.exe.zip](https://murex.rocks/bin/latest/murex-windows-amd64.exe.zip)
* [murex-windows-arm.exe.zip](https://murex.rocks/bin/latest/murex-windows-arm.exe.zip)

> macOS (Darwin) builds for arm64 should support the ARM-based M1 processor.
> However please treat these builds as experimental. If you do run into any
> issues then log them at (https://github.com/lmorg/murex/issues) and use the
> amd64 builds in the meantime.

### Linux / UNIX / macOS Instructions

Please download the appropriate `.gz` file from the list above, one that
matches both your OS and CPU architecture.

For example, to download a 64bit version for Linux:

```
wget https://murex.rocks/bin/latest/murex-linux-amd64.gz
gunzip murex-linux-amd64.gz
chmod +x murex-linux-amd64
./murex-linux-amd64
```

> macOS builds are listed as [darwin](https://en.wikipedia.org/wiki/Darwin_(operating_system))
> as per the name of Apple have given to their [underlying OS](https://en.wikipedia.org/wiki/MacOS#Architecture).

Most of these builds have received _some_ level of user acceptance testing with
Linux and macOS builds receiving the most attention, because that's what we use
ourselves.

### Windows Instructions

Click the Windows link that matches your CPU architecture. Unzip using your
preferred too then launch using your preferred console. _murex_ cannot be
started via double clicking the executable -- it requires a starting from
within an existing console session.

Please also note that Windows support is also considered experimental. In part
due to the lack of **coreutils** (as seen on Linux and UNIX) and in part due to
the different underpinning technologies behind consoles / terminal emulators.
If you do experience some wonky behavior then our recommendation is to run the
`linux-amd64` build for Linux on top of WSL. The instructions above will guide
you through installing on Linux, WSL install instructions can be found at the
following site: (https://docs.microsoft.com/en-us/windows/wsl/install-win10)

### Plan9 Instructions

Plan9 is untested. The code compiles and it is syscall compatible with Plan9
operating systems, however you may experience bugs using _murex_ on Plan9. If
you do encounter any issues then please raise them at:
(https://github.com/lmorg/murex/issues)
