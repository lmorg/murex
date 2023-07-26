<template><div><h1 id="what-s-new-in-murex-v2-9-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-9-change-log" aria-hidden="true">#</a> What's new in murex v2.9 - Change Log</h1>
<p>This release focuses on testing and REPL usability improvements but also includes updates several new run modes to make error handling easier in larger scripts.</p>
<p>Breaking Changes:</p>
<ul>
<li>None</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p><code v-pre>&gt;&gt;</code> redirect (append to file) operator added. This is just syntactic sugar for <code v-pre>-&gt; &gt;&gt;</code>. To redirect and overwrite/truncate a file use <code v-pre>|&gt;</code> rather than <code v-pre>&gt;</code></p>
</li>
<li>
<p>Support for automatic globbing in the commandline. This has to be enabled via <code v-pre>config: set shell auto-glob true</code> and you will get prompted for each glob. Any globs inside blocks will be ignored and this only works on the interactive prompt (ie not in scripts or background processes) -- these limitations are not a bug but by design</p>
</li>
<li>
<p>Dynamic autocomplete suggestions are now cached for 5 seconds by default. This cache can be changed via the <strong>CacheTTL</strong> directive in <code v-pre>autocomplete</code></p>
</li>
<li>
<p>Additional scoping for <code v-pre>try</code> and <code v-pre>trypipe</code>. You can define the scope as being an entire function or a module. If set for an entire module then any functions, autocompletions, etc inside that module will inherit that <code v-pre>try</code> or <code v-pre>trypipe</code> mode automatically. These are set via the new builtin <code v-pre>runmode</code></p>
</li>
<li>
<p>New reserved variable <code v-pre>$PWD</code> (this was just a global variable previously)</p>
</li>
<li>
<p>Additional flags added to range, <code v-pre>@[</code></p>
</li>
<li>
<p>New builtin, <code v-pre>man-get-flags</code>. This exposes Murex's internal man page parser to the command line scripts</p>
</li>
<li>
<p>Binaries now produced for ARM64 Windows</p>
</li>
</ul>
<p>Minor Changes:</p>
<ul>
<li>
<p>The order in which namespaces are looked up has been altered since <code v-pre>private</code>s should be resolved before <code v-pre>function</code>s</p>
</li>
<li>
<p>Updated autocomplete for homebrew, <code v-pre>brew</code></p>
</li>
<li>
<p>New autocomplete for AWS CLI, <code v-pre>aws</code></p>
</li>
<li>
<p>New autocomplete for yarn, <code v-pre>yarn</code></p>
</li>
<li>
<p>Murex unit tests should report progressively rather than everything upon completion</p>
</li>
<li>
<p>Minimum recommended version of Go has been raised to 1.15. This is due to some bugs in dependencies when running tests against v1.13 and v1.14. Murex will technically compile and work on them but it not worth my time debugging why the CircleCI tests fail for those versions given their age</p>
</li>
<li>
<p>Vendor directory removed. This isn't really needed now support is focused on newer versions of Go</p>
</li>
<li>
<p>Website: New 404 error</p>
</li>
<li>
<p>Website: search bar added</p>
</li>
<li>
<p>The start of a massive push towards upping the test coverage</p>
</li>
<li>
<p>Lots of improvements to the docs</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>FileRef module string for functions forked from the REPL changed to <code v-pre>murex/shell</code> to bring it inline with the naming convention of other modules</li>
</ul>
<hr>
<p>Published: 15.07.2022 at 20:35</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-append.html">Append Pipe (<code v-pre>&gt;&gt;</code>) Token</RouterLink>:
Redirects STDOUT to a file and append its contents</li>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runmode.html"><code v-pre>runmode</code></RouterLink>:
Alter the scheduler's behaviour at higher scoping level</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


