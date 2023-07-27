<template><div><h1 id="modules-and-packages" tabindex="-1"><a class="header-anchor" href="#modules-and-packages" aria-hidden="true">#</a> Modules and Packages</h1>
<blockquote>
<p>An introduction to Murex modules and packages</p>
</blockquote>
<p>Murex has it's own module system with namespacing and a package manager. But
why should a shell need all this?</p>
<p>The answer comes from years of me using Bash and wishing my Bash environment
could be consistent across multiple machines. So this document is authored from
the perspective of my personal usage (&quot;me&quot; being Laurence Morgan, the original
author of Murex).</p>
<p>What Murex's package system provides is:</p>
<ol>
<li>A way to ensure consistency across multiple platforms</li>
<li>An easy way to extend Murex</li>
<li>An easy way to share what you've extended with others</li>
<li>An easy way to ensure your extensions are kept up-to-date</li>
<li>An easy way to track what code is running in your shell and from where it
was loaded</li>
</ol>
<p>Before I address those points in more detail, a bit of background into what
modules and packages are:</p>
<h3 id="what-are-packages-and-modules" tabindex="-1"><a class="header-anchor" href="#what-are-packages-and-modules" aria-hidden="true">#</a> What Are Packages And Modules?</h3>
<p>Murex comes with it's own package manager to make managing plugins easier.</p>
<p>The format of the packages is a directory, typically located at <code v-pre>~/.murex_modules</code>,
which contains one or more murex scripts. Each script can be it's own module.
ie there are multiple modules that can be grouped together and distributed as a
single package.</p>
<p>The way packages and modules are represented is as a path:</p>
<pre><code>package/module
</code></pre>
<p><code v-pre>murex-package</code> is a package management tool for administrating murex modules
and packages.</p>
<h2 id="using-packages-and-modules" tabindex="-1"><a class="header-anchor" href="#using-packages-and-modules" aria-hidden="true">#</a> Using Packages And Modules</h2>
<h3 id="consistency" tabindex="-1"><a class="header-anchor" href="#consistency" aria-hidden="true">#</a> Consistency</h3>
<p>Package database are stored locally at <code v-pre>~/.murex_modules/packages.json</code>. This
file is portable so any new machine can have <code v-pre>packages.json</code> imported. The
easiest way of doing this is using <code v-pre>murex-package</code> which can import from a
local path or HTTP(S) URI and automatically download any packages described in
the database.</p>
<p>For example the command I run on any new dev machine to import all of my DevOps
tools and terminal preferences is the following:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>murex-package: import https://gist.githubusercontent.com/lmorg/770c71786935b44ba6667eaa9d470888/raw/fb7b79d592672d90ecb733944e144d722f77fdee/packages.json
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="extendability" tabindex="-1"><a class="header-anchor" href="#extendability" aria-hidden="true">#</a> Extendability</h3>
<p>Namespacing allows for <code v-pre>private</code> functions which allows you to write smaller
functions. Smaller functions are easier to write tests against (Murex also
has an inbuilt testing and debugging tools).</p>
<h3 id="sharing-code" tabindex="-1"><a class="header-anchor" href="#sharing-code" aria-hidden="true">#</a> Sharing Code</h3>
<p>Packages can be hosted via HTTP(S) or git. Anyone can import anyone elses
packages using <code v-pre>murex-package</code>.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>murex-package: install https://github.com/lmorg/murex-module-murex-dev.git
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="updating-packages" tabindex="-1"><a class="header-anchor" href="#updating-packages" aria-hidden="true">#</a> Updating Packages</h3>
<p>Updating packages is easy:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>murex-package: update
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h3 id="tracking-code" tabindex="-1"><a class="header-anchor" href="#tracking-code" aria-hidden="true">#</a> Tracking Code</h3>
<p>All code loaded in Murex, every function, variable and event (etc) is stored
in memory with metadata about where it was loaded from; which package, file and
at what time. This is called <code v-pre>FileRef</code>.</p>
<p>For more information on <code v-pre>FileRef</code> see the link below.</p>
<h3 id="module-strings-for-non-module-code" tabindex="-1"><a class="header-anchor" href="#module-strings-for-non-module-code" aria-hidden="true">#</a> Module Strings For Non-Module Code</h3>
<h4 id="source" tabindex="-1"><a class="header-anchor" href="#source" aria-hidden="true">#</a> Source</h4>
<p>A common shell idiom is to load shell script files via <code v-pre>source</code> / <code v-pre>.</code>. When
this is done the module string (as seen in the <code v-pre>FileRef</code> structures described
above) will be <code v-pre>source/hash</code> where <strong>hash</strong> will be a unique hash of the file
path and load time.</p>
<p>Thus no two sourced files will share the same module string. Even the same file
but modified and sourced twice (before and after the edit) will have different
module strings due to the load time being part of the hashed data.</p>
<h4 id="repl" tabindex="-1"><a class="header-anchor" href="#repl" aria-hidden="true">#</a> REPL</h4>
<p>Any functions, variables, events, auto-completions, etc created manually,
directly, in the interactive shell will have a module string of <code v-pre>murex</code> and an
empty Filename string.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/fileref.html">FileRef</RouterLink>:
How to track what code was loaded and from where</li>
<li><RouterLink to="/commands/murex-package.html"><code v-pre>murex-package</code></RouterLink>:
Murex's package manager</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
</ul>
</div></template>


