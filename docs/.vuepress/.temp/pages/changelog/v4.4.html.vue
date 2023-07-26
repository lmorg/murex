<template><div><h1 id="v4-4-change-log" tabindex="-1"><a class="header-anchor" href="#v4-4-change-log" aria-hidden="true">#</a> v4.4 - Change Log</h1>
<p>v4.4 features two new builtins, improvements in testing, and automatic generation of autocompletion suggestions backed by man page parsing. Plus there has been a lot of focus on improving <em>readline</em> responsiveness</p>
<h2 id="v4-4-9000" tabindex="-1"><a class="header-anchor" href="#v4-4-9000" aria-hidden="true">#</a> v4.4.9000</h2>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes</h3>
<p>None</p>
<h3 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>readline: <code v-pre>f1</code> should work even outside of autocomplete</p>
</li>
<li>
<p>readline: wider preview pane</p>
</li>
</ul>
<h3 id="bug-fixes" tabindex="-1"><a class="header-anchor" href="#bug-fixes" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>readline: fix preview wrapping on long lines</p>
</li>
<li>
<p>readline: image previews should scale to preview height</p>
</li>
<li>
<p>readline: better handling of terminal resizing</p>
</li>
</ul>
<h2 id="v4-4-8000" tabindex="-1"><a class="header-anchor" href="#v4-4-8000" aria-hidden="true">#</a> v4.4.8000</h2>
<h3 id="breaking-changes-1" tabindex="-1"><a class="header-anchor" href="#breaking-changes-1" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li><code v-pre>alt</code>+<code v-pre>1</code>..<code v-pre>9</code> hotkeys replaced with <code v-pre>shift</code>+<code v-pre>f1</code>..<code v-pre>f12</code>. Changed because <code v-pre>alt</code>+numeric rarely worked (<RouterLink to="/user-guide/terminal-keys.html#recalling-previous-words">read more</RouterLink>)</li>
</ul>
<h3 id="features-1" tabindex="-1"><a class="header-anchor" href="#features-1" aria-hidden="true">#</a> Features</h3>
<p>None</p>
<h3 id="bug-fixes-1" tabindex="-1"><a class="header-anchor" href="#bug-fixes-1" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p><code v-pre>@IncManPages</code> <code v-pre>autocomplete</code> value for <strong>Dynamic</strong> directive renamed to <code v-pre>@IncManPage</code>, like the <strong>IncManPage</strong> directive</p>
</li>
<li>
<p>readline: render glitch fixed with delayed completions racing against hint text updates</p>
</li>
<li>
<p>readline: render glitches fixed when buffers are not being reset correctly after preview box has, or should have, closed</p>
</li>
<li>
<p>readline: man page preview now removes backspace characters correctly from UNIX docs</p>
</li>
<li>
<p>readline: man page preview now more reliably scrolls to the right line when a flag is highlighted in autocompletion</p>
</li>
<li>
<p>readline: preview now works for list views as well as grid views</p>
</li>
</ul>
<h2 id="v4-4-7000" tabindex="-1"><a class="header-anchor" href="#v4-4-7000" aria-hidden="true">#</a> v4.4.7000</h2>
<h3 id="breaking-changes-2" tabindex="-1"><a class="header-anchor" href="#breaking-changes-2" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li><strong>preview-enabled</strong> <code v-pre>config</code> option removed. This actually doesn't break anything, just produces a warning. The <code v-pre>config</code> option was also undocumented and experimental</li>
</ul>
<h3 id="features-2" tabindex="-1"><a class="header-anchor" href="#features-2" aria-hidden="true">#</a> Features</h3>
<ul>
<li>Preview is now considered stable. Press <code v-pre>f1</code> while autocomplete is open to use (<RouterLink to="/user-guide/terminal-keys.html#preview-box">read more</RouterLink>)</li>
</ul>
<h3 id="bug-fixes-2" tabindex="-1"><a class="header-anchor" href="#bug-fixes-2" aria-hidden="true">#</a> Bug Fixes</h3>
<ul>
<li>
<p>readline: render glitch fixed with delayed path completions</p>
</li>
<li>
<p>readline: excess white spaces removed from man page descriptions</p>
</li>
</ul>
<h2 id="v4-4-x" tabindex="-1"><a class="header-anchor" href="#v4-4-x" aria-hidden="true">#</a> v4.4.x</h2>
<h3 id="breaking-changes-3" tabindex="-1"><a class="header-anchor" href="#breaking-changes-3" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li>'Name' field dropped from <code v-pre>onCommandCompletion</code> interrupt (this field was never documented)</li>
</ul>
<h3 id="features-3" tabindex="-1"><a class="header-anchor" href="#features-3" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>New builtin, <code v-pre>round</code>, which can perform rounding operations on numerical inputs (<a href="https://github.com/lmorg/murex/issues/630" target="_blank" rel="noopener noreferrer">#630<ExternalLinkIcon/></a>, <RouterLink to="/commands/round.html">read more</RouterLink>)</p>
</li>
<li>
<p>Vastly improved automatic man page parsing. Now descriptions are pulled alongside the flags and results are cached</p>
</li>
<li>
<p><code v-pre>man-get-flags</code> now includes a <code v-pre>-d</code> / <code v-pre>--description</code> flag to expose the improved man page parser</p>
</li>
<li>
<p>New builtin, <code v-pre>return</code>, which exits a given function scope (eg <code v-pre>function</code>, <code v-pre>private</code>, <strong>Dynamic</strong> block in <code v-pre>autocomplete</code>, etc)</p>
</li>
<li>
<p>Improved <code v-pre>git</code> autocompletions</p>
</li>
<li>
<p><code v-pre>find</code> autocompletions added</p>
</li>
<li>
<p>builtin profiles are now imported as separate modules. This makes debugging easier</p>
</li>
<li>
<p><code v-pre>/integrations</code> directory added to Murex source, the aim of which is to make it easier for people to contribute autocompletions and other integrations with common command line tools (<a href="https://github.com/lmorg/murex/tree/master/intergrations" target="_blank" rel="noopener noreferrer">Github<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>readline: new word jump hotkeys <code v-pre>ctrl</code>+<code v-pre>left</code> / <code v-pre>ctrl</code>+<code v-pre>right</code> (PC)</p>
</li>
<li>
<p>readline: new word jump hotkeys <code v-pre>option</code>+<code v-pre>left</code> / <code v-pre>option</code>+<code v-pre>right</code> (Mac)</p>
</li>
</ul>
<h3 id="bug-fixes-3" tabindex="-1"><a class="header-anchor" href="#bug-fixes-3" aria-hidden="true">#</a> Bug Fixes:</h3>
<ul>
<li>
<p><code v-pre>test</code> would always run first in any block, regardless of the preferred order of execution. This was because <code v-pre>test (unit|state|config)</code> required altering the execution state of the shell. The drawback was that <code v-pre>test (run|define|report)</code> would also run unexpectedly. This lead to hacks like <code v-pre>try { test run * }</code> to force the correct order of operations. Now the parameters of <code v-pre>test</code> are checked to determine when to execute the builtin.</p>
</li>
<li>
<p><code v-pre>int</code> types couldn't be compared against <code v-pre>num</code> types with <code v-pre>&gt;</code>, <code v-pre>&gt;=</code>, <code v-pre>&lt;</code>, <code v-pre>&lt;=</code> operators</p>
</li>
<li>
<p>readline: <code v-pre>^d</code> will not send EOF if line is not empty. The original behaviour was by design however after pressing <code v-pre>^d</code> a few too many times when I intended to press <code v-pre>^f</code> or <code v-pre>^r</code>, I decided the original design was wrong</p>
</li>
<li>
<p><code v-pre>FileRef</code> wasn't being set in <code v-pre>test</code>. This caused some tests to fail if they called private functions</p>
</li>
<li>
<p>Check <code v-pre>~/.ssh/config</code> exists before trying to parse it in <code v-pre>getHostsFile</code> private (used for SSH and similar tools autocompletions)</p>
</li>
<li>
<p>readline: lots of work done on speeding up redraws and overall responsiveness. eg buffered autocompletion menus</p>
</li>
<li>
<p>readline: hint text should <em>never</em> be displayed if disabled via <code v-pre>config</code></p>
</li>
<li>
<p>readline: soft timeout halved (this can be overridden via <code v-pre>config</code>)</p>
</li>
<li>
<p>readline: cropped autocompletion suggestions in gridded layout were one character too short (off by one error)</p>
</li>
</ul>
<hr>
<p>Published: 18.07.2023 at 22:58</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/fileref.html">FileRef</RouterLink>:
How to track what code was loaded and from where</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/types/int.html"><code v-pre>int</code> </RouterLink>:
Whole number (primitive)</li>
<li><RouterLink to="/commands/man-get-flags.html"><code v-pre>man-get-flags</code> </RouterLink>:
Parses man page files for command line flags</li>
<li><RouterLink to="/types/num.html"><code v-pre>num</code> (number)</RouterLink>:
Floating point number (primitive)</li>
<li><RouterLink to="/events/oncommandcompletion.html"><code v-pre>onCommandCompletion</code></RouterLink>:
Trigger an event upon a command's completion</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/return.html"><code v-pre>return</code></RouterLink>:
Exits current function scope</li>
<li><RouterLink to="/commands/round.html"><code v-pre>round</code> </RouterLink>:
Round a number by a user defined precision</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
</ul>
</div></template>


