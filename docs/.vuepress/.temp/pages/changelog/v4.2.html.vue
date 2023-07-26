<template><div><h1 id="what-s-new-in-murex-v4-2-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v4-2-change-log" aria-hidden="true">#</a> What's new in murex v4.2 - Change Log</h1>
<p>Murex usage has raised considerably in recent weeks. This release addresses a number of feature requests and bugs raised on Github.</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes</h3>
<p>none</p>
<h3 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>new flag in <code v-pre>foreach</code>: <code v-pre>--step &lt;int&gt;</code>. This allows the <code v-pre>foreach</code> to jump indexes at a time. When used, all the lines jumped are merged into a JSON array and that is passed to the <code v-pre>foreach</code> nested code block</p>
</li>
<li>
<p>new alias <code v-pre>builtin</code> =&gt; <code v-pre>fexec builtin</code>. This was added to bring more familiarity to those coming from Bash et al (<a href="https://github.com/lmorg/murex/issues/608" target="_blank" rel="noopener noreferrer">#608<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>on Darwin, Homebrew defaults are imported, just like you'd expect them to be from any other shell</p>
</li>
<li>
<p>new hotkeys (<a href="https://github.com/lmorg/murex/issues/611" target="_blank" rel="noopener noreferrer">#611<ExternalLinkIcon/></a>) (<a href="https://github.com/lmorg/murex/issues/613" target="_blank" rel="noopener noreferrer">#613<ExternalLinkIcon/></a>):</p>
<ol>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>a</code>: jump to beginning of line</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>e</code>: jump to end of line</p>
</li>
<li>
<p><code v-pre>alt</code>+<code v-pre>f</code>: jump forwards a word at a time</p>
</li>
<li>
<p><code v-pre>alt</code>+<code v-pre>b</code>: jump backwards a word at a time</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>k</code>: clear line after cursor position</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>l</code>: clear entire screen</p>
</li>
<li>
<p><code v-pre>ctrl</code>+<code v-pre>g</code>: same as <code v-pre>esc</code></p>
</li>
</ol>
</li>
<li>
<p><code v-pre>open</code> builtin should fallback to system default if type unknown (<a href="https://github.com/lmorg/murex/issues/620" target="_blank" rel="noopener noreferrer">#620<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p><code v-pre>pre-prompt-func</code> added to <code v-pre>config</code>. Block runs before the interactive prompt is displayed</p>
</li>
<li>
<p><code v-pre>post-prompt-func</code> added to <code v-pre>config</code>. Block runs after the interactive prompt has been exited but before any command lines have been executed</p>
</li>
<li>
<p><strong>auto-cd</strong> option added to <code v-pre>config</code>: <code v-pre>config: set shell auto-cd true</code>. This allows you to traverse directories without having to prefix <code v-pre>cd</code>. By default it is disabled</p>
</li>
<li>
<p>Meta values added to <code v-pre>foreach</code>, <code v-pre>formap</code>, and <code v-pre>while</code>. Meta values are a <code v-pre>$.</code> variable that holds meta-information about the running block, such as the number of iterations in a loop</p>
</li>
<li>
<p><code v-pre>command</code> alias for <code v-pre>exec</code> (for familiarity with Bash)</p>
</li>
<li>
<p><code v-pre>builtin</code> alias for <code v-pre>fexec builtin</code> (for familiarity with Bash)</p>
</li>
</ul>
<h3 id="bug-fixes" tabindex="-1"><a class="header-anchor" href="#bug-fixes" aria-hidden="true">#</a> Bug Fixes:</h3>
<ul>
<li>
<p>brace encapsulated variables, eg <code v-pre>$(foobar)</code>, now autocomplete as expected</p>
</li>
<li>
<p>improvements to inlining images in iTerm2 when shell running on top of <code v-pre>tmux</code>. It still doesn't work reliably but this is an issue with <code v-pre>tmux</code> rather than Murex</p>
</li>
<li>
<p><code v-pre>method</code> forwards exceptions rather than errors. This behaviour has been changed to surface issues as errors</p>
</li>
<li>
<p><code v-pre>profile_preload.mx</code> added to profile defaults. This will be loaded before any of the OS-specific profiles</p>
</li>
<li>
<p>missing rune length check causing the shell to panic (<a href="https://github.com/lmorg/murex/issues/617" target="_blank" rel="noopener noreferrer">#617<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>fixed deadlock in flag parsing</p>
</li>
<li>
<p>zero length string title bars should be ignored (<a href="https://github.com/lmorg/murex/issues/614" target="_blank" rel="noopener noreferrer">#614<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p>minor improvements to documentation</p>
</li>
</ul>
<hr>
<p>Published: 19.06.2023 at 23:15</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/exec.html"><code v-pre>exec</code></RouterLink>:
Runs an executable</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/method.html"><code v-pre>method</code></RouterLink>:
Define a methods supported data-types</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
</ul>
</div></template>


