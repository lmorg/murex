<template><div><h1 id="what-s-new-in-murex-v2-11-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-11-change-log" aria-hidden="true">#</a> What's new in murex v2.11 - Change Log</h1>
<p>This release mainly focuses on refinements in performance and usability, rather than introducing new features</p>
<p>Breaking Changes:</p>
<ul>
<li>None</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p>File index caching for faster completions (eg on mechanical drives). This happens upon shell start up (except on Darwin for security reasons) and after every <code v-pre>cd</code></p>
</li>
<li>
<p>Relaxed <code v-pre>alias</code> syntax. It now supports whitespace between the <code v-pre>=</code> statement (like <code v-pre>set</code> and <code v-pre>let</code> do)</p>
</li>
<li>
<p>New aliases added for common methods which work against lists/arrays. Type <code v-pre>list.{TAB}</code> to see the complete list of methods</p>
</li>
<li>
<p>Added support for alternative separator values in <code v-pre>struct-keys</code></p>
</li>
<li>
<p>Hint summaries are now cached, to improve performance on slower machines</p>
</li>
<li>
<p>Opt in support to pre-cache hint summaries. This will also auto-disable appending colon to command suggestions in autocomplete. This is not enabled by default because it could negatively affect performance on some slower machines and is more of an aesthetic improvement. You can enable it by adding <code v-pre>config: set shell pre-cache-hint-summaries true</code> to your Murex profile</p>
</li>
<li>
<p>New autocompletions for <code v-pre>yay</code> (ArchLinux)</p>
</li>
<li>
<p>Some minor performance improvements</p>
</li>
<li>
<p>website: updates to docs</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>
<p>readline: max field size adjusted for autocompletes so more completions are visible at one time</p>
</li>
<li>
<p>readline: all cropped autocompletions now provide an uncropped version in the hint text field</p>
</li>
<li>
<p>readline: fix bug with history not getting written</p>
</li>
<li>
<p>OS signals (^c et al) are now re-registered after each exec (fixes bug where some programs would &quot;steal&quot; ^c, breaking support for those signals upon subsequent execs)</p>
</li>
<li>
<p>Fixed regression bug in man page parsing</p>
</li>
<li>
<p>Added check for zero-length string used as a process name (this shouldn't ever happen under normal circumstances however closing this bug allows for a managed error rather than the code panicking)</p>
</li>
<li>
<p>readline: Prompt environment reset on each loop (this change is to future-proof against potential bugs)</p>
</li>
<li>
<p>Fixed race condition bug when some of the tests are run concurrently (this bug only existed in the unit tests but sometimes caused those tests to fail)</p>
</li>
<li>
<p>Added CI tests for Go v19.x</p>
</li>
</ul>
<hr>
<p>Published: 12.09.2022 at 08:10</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/cd.html"><code v-pre>cd</code></RouterLink>:
Change (working) directory</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/struct-keys.html"><code v-pre>struct-keys</code></RouterLink>:
Outputs all the keys in a structure as a file path</li>
</ul>
</div></template>


