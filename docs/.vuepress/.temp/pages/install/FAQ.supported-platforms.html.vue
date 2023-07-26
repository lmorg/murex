<template><div><h1 id="supported-operating-systems" tabindex="-1"><a class="header-anchor" href="#supported-operating-systems" aria-hidden="true">#</a> Supported Operating Systems</h1>
<p>The following is a list of platforms Murex has been tested on and the
level of support it has:</p>
<h2 id="linux" tabindex="-1"><a class="header-anchor" href="#linux" aria-hidden="true">#</a> Linux</h2>
<p>The shell has been extensively tested across a number of distributions. There
are no known distribution specific issues.</p>
<h2 id="macos-darwin" tabindex="-1"><a class="header-anchor" href="#macos-darwin" aria-hidden="true">#</a> macOS (Darwin)</h2>
<p>All features work aside alt-hotkeys.</p>
<p>Both x86 (Intel) and AMD (M2 etc) architectures are supported.</p>
<h2 id="windows" tabindex="-1"><a class="header-anchor" href="#windows" aria-hidden="true">#</a> Windows</h2>
<p>Windows is supported there are a few known issue with the way how Windows
internals are built. These cannot be easily worked around:</p>
<ul>
<li>
<p>Windows doesn't decouple the terminal emulator and the shell Which means you
cannot rely upon STDIN working as expected (eg some commands don't read input
from STDIN but instead poll the terminal emulator directly)</p>
</li>
<li>
<p>Windows sends parameters as a single string rather than an array of string.
This is to retain backwards compatibility with DOS but it breaks the way how
quotation marks and escaping works. Murex will compile an array of
parameters based on the quotation strings (there are 3 different types of
quotations in Murex), infixed variables, subshells, etc. These would not be
honoured by any Windows commands because every Windows application then has
to handle how the one long string of parameters is chopped up into different
arguments; how quotation marks are handles, spaces, escaping, etc. This means
there is no standard so one command might handle spaces correctly but another
wouldn't.</p>
</li>
<li>
<p>Job control (<code v-pre>bg</code>, <code v-pre>^z</code>, <code v-pre>fg</code>, etc) isn't supported because Windows doesn't
have an equivalent of the SIGSTSP (etc) POSIX signals.</p>
</li>
<li>
<p>There is also the caveat that without a broad range of command line utilities
(eg GNU coreutils) the usefulness of Murex is seriously diminished. You can
mitigate this by installing <a href="https://www.msys2.org/" target="_blank" rel="noopener noreferrer">MSYS2<ExternalLinkIcon/></a> or <a href="https://cygwin.com/" target="_blank" rel="noopener noreferrer">Cygwin<ExternalLinkIcon/></a>.</p>
</li>
</ul>
<h2 id="freebsd" tabindex="-1"><a class="header-anchor" href="#freebsd" aria-hidden="true">#</a> FreeBSD</h2>
<p>FreeBSD is officially supported and tested by the community.</p>
<h2 id="openbsd" tabindex="-1"><a class="header-anchor" href="#openbsd" aria-hidden="true">#</a> OpenBSD</h2>
<p>FreeBSD is officially supported and tested by the community.</p>
<h2 id="netbsd" tabindex="-1"><a class="header-anchor" href="#netbsd" aria-hidden="true">#</a> NetBSD</h2>
<p>FreeBSD is officially supported and tested by the community.</p>
<h2 id="dragonflybsd" tabindex="-1"><a class="header-anchor" href="#dragonflybsd" aria-hidden="true">#</a> DragonflyBSD</h2>
<p>FreeBSD is officially supported and tested by the community.</p>
<h2 id="plan-9" tabindex="-1"><a class="header-anchor" href="#plan-9" aria-hidden="true">#</a> Plan 9</h2>
<p>Plan 9 is included as part of the automated built tests however no functional
tests have been run.</p>
<p>If you do happen to run into any such bugs then I do welcome pull requests.</p>
<p>Feature wise, job control isn't supported in Plan 9 because Plan 9 doesn't
support all of the required signals. All other functions are expected to work.</p>
<h2 id="other-cpu-architectures" tabindex="-1"><a class="header-anchor" href="#other-cpu-architectures" aria-hidden="true">#</a> Other CPU architectures</h2>
<p>Several CPU architectures are supported:</p>
<ul>
<li>386 (x86 32bit)</li>
<li>AMD64 (x86 64bit)</li>
<li>ARMv7 (32bit)</li>
<li>ARMv8 (64bit)</li>
</ul>
</div></template>


