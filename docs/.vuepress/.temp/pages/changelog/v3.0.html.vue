<template><div><h1 id="what-s-new-in-murex-v3-0-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v3-0-change-log" aria-hidden="true">#</a> What's new in murex v3.0 - Change Log</h1>
<p>This is a major release that brings a significant number of changes and improvements, including a complete overhaul of the parser. Backwards compatibility is a high priority however these new features bring greater readability and consistency to shell scripting. So while the older syntax remains for compatibility, it is worth migrating over to the newer syntax for all new code being written</p>
<p>Breaking Changes:</p>
<ul>
<li>
<p>Optional builtin removed: <code v-pre>bson</code>. This was disabled by default and likely never used. So it has been removed to reduce Murex's testing footprint. It can easily be re-added if anyone does actually use it</p>
</li>
<li>
<p>Optional builtin removed: <code v-pre>coreutils</code>. This was actually enabled by default for Windows builds. However rewriting Linux/UNIX coreutils for Windows support is a massive project in its own right and with the maturity of WSL there's almost no reason to run Murex on &quot;native Windows&quot;. So the <code v-pre>coreutils</code> builtin has been dropped to allow us to focus on the responsibilities of the shell</p>
</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p>Support for expressions, eg <code v-pre>5 * 5</code> or <code v-pre>foobar = $foo + &quot;bar&quot;</code>, etc. This syntax can be used directly or specified specifically via the <code v-pre>expr</code> builtin</p>
</li>
<li>
<p>New syntax sugar for creating JSON objects: <code v-pre>%{ foo: bar }</code></p>
</li>
<li>
<p>New syntax sugar for creating JSON arrays: <code v-pre>%[ foo bar ]</code></p>
</li>
<li>
<p>New syntax sugar for creating strings: <code v-pre>%()</code> (this is preferred over the, now deprecated, feature of parenthesis quotes)</p>
</li>
<li>
<p>Ranges supported in <code v-pre>[]</code> (<code v-pre>@[</code> is now deprecated)</p>
</li>
<li>
<p>Support for multiline comments: <code v-pre>/# comment #/</code>. It is unfortunate this differs from C-style comments (<code v-pre>/* comment */</code>) but this has to be the case because <code v-pre>/*</code> is ambiguous for barewords in shells: is it a path and glob or a comment? Where as <code v-pre>/#</code> isn't a common term due to <code v-pre>#</code> denoting a comment</p>
</li>
<li>
<p>If any processes pass <code v-pre>null</code> as a data type across the pipe, it will be ignored. This solves the problem where functions that don't write to STDOUT would still define the data type</p>
</li>
<li>
<p>Config option <strong>auto-glob</strong> renamed to <strong>expand-glob</strong>, and now enabled by default</p>
</li>
<li>
<p>Globbing exclusion list. This allows you to avoid annoying prompts when parameters shouldn't be expanded as globs by the shell (eg when using regular expressions). This can be managed via <strong>shell expand-glob-unsafe-commands</strong> option in <code v-pre>config</code></p>
</li>
<li>
<p><code v-pre>@g</code> removed. It is no longer needed with <strong>expand-glob</strong> enabled by default</p>
</li>
<li>
<p>New builtin: <code v-pre>continue</code>: skip subsequent processes in an iteration block and continue to next iteration</p>
</li>
<li>
<p>New builtin: <code v-pre>break</code>: exit out of a block of code (eg in an iteration loop)</p>
</li>
<li>
<p>Additional syntax for <em>index</em> (<code v-pre>[</code>): <code v-pre>*1</code>: 1st row, <code v-pre>*A</code>: 1st column</p>
</li>
<li>
<p>New alias: <code v-pre>help</code> -&gt; <code v-pre>murex-docs</code>. This brings Murex a little more inline with Bash et al</p>
</li>
<li>
<p><strong>pre-cache-hint-summaries</strong> now enabled by default after testing has demonstrated it doesn't have nearly as expensive footprint as first assumed</p>
</li>
<li>
<p>Hitting <strong>TAB</strong> when nothing has been typed in the REPL will suggest past command lines</p>
</li>
<li>
<p><code v-pre>^</code> autocompletion added</p>
</li>
<li>
<p><code v-pre>getfile</code> writes to disk if STDOUT is a TTY</p>
</li>
<li>
<p><strong>mkarray</strong> (eg <code v-pre>ja</code>) now writes an integer array if range is integers. eg <code v-pre>ja: [1..3]</code>. This change wouldn't affect <code v-pre>a</code> since that outputs as list of strings (for streaming performance reasons) rather than a data type aware document</p>
</li>
<li>
<p><code v-pre>debug</code> (method) output tweaked</p>
</li>
<li>
<p>Improved error messages in a number places</p>
</li>
<li>
<p>Revamped README / website landing page</p>
</li>
</ul>
<p>Non-User Facing / Maintenance Changes:</p>
<ul>
<li>
<p>Minimum Go version supported is now 1.17.x</p>
</li>
<li>
<p>Main parser completely rewritten</p>
</li>
<li>
<p><code v-pre>ReadArray</code> API now requires a <code v-pre>context.Context</code></p>
</li>
<li>
<p><code v-pre>egrep</code> references changed to <code v-pre>grep -E</code> to work around GNU grep deprecating support for <em>egrep</em></p>
</li>
<li>
<p>Added marshallers for <code v-pre>boolean</code>, <code v-pre>null</code></p>
</li>
<li>
<p><code v-pre>Variables.GetValue()</code> now errors instead of returns <code v-pre>nil</code> when no variable set</p>
</li>
<li>
<p>Additional tests. So many new tests added</p>
</li>
<li>
<p>Lots of code refactoring</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>
<p><code v-pre>regexp</code> wasn't erroring if nothing was matched</p>
</li>
<li>
<p>readline: fixed deadlock</p>
</li>
<li>
<p><code v-pre>append</code> and <code v-pre>prepend</code> now type aware (no longer converts all arrays into string arrays)</p>
</li>
<li>
<p><code v-pre>foreach</code> was setting variables as strings rather than honoring their original data type</p>
</li>
<li>
<p><code v-pre>yarn</code> autocompletion errors should be suppressed</p>
</li>
<li>
<p>spellcheck missing <code v-pre>break</code> causing more occasionally incorrect instances of underlined words</p>
</li>
<li>
<p><code v-pre>config</code> wasn't passing data type when executing blocks via <strong>eval</strong></p>
</li>
<li>
<p><code v-pre>debug</code> wasn't setting data type when used as a function</p>
</li>
<li>
<p>macro variables don't re-prompt when the same variable is used multiple times</p>
</li>
</ul>
<hr>
<p>Published: 31.12.2022 at 08:10</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/continue.html"><code v-pre>continue</code></RouterLink>:
terminate process of a block within a caller function</li>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/getfile.html"><code v-pre>getfile</code></RouterLink>:
Makes a standard HTTP request and return the contents as Murex-aware data type for passing along Murex pipelines.</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/murex-docs.html"><code v-pre>murex-docs</code></RouterLink>:
Displays the man pages for Murex builtins</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
</ul>
</div></template>


