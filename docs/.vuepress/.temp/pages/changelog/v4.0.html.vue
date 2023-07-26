<template><div><h1 id="what-s-new-in-murex-v4-0-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v4-0-change-log" aria-hidden="true">#</a> What's new in murex v4.0 - Change Log</h1>
<p>This release sees significant improvements for use with non-latin characters in both the interactive prompt and shell scripting. It introduces new syntax to make working with structured data even easier than before. As well as new data types and smoother user experience.</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes</h3>
<ul>
<li>
<p><code v-pre>$PWD</code> is no longer a <code v-pre>str</code>. If you need to parse <code v-pre>$PWD</code> as a string then you should <code v-pre>out: $PWD</code> or use <code v-pre>${pwd}</code></p>
</li>
<li>
<p><code v-pre>.</code> is now a valid bareword character for variables. This will break any instances in code that contain something like <code v-pre>out $hello.</code> however you can now use <code v-pre>out $(hello).</code> syntax for safer string infixing</p>
</li>
</ul>
<h3 id="features" tabindex="-1"><a class="header-anchor" href="#features" aria-hidden="true">#</a> Features</h3>
<ul>
<li>
<p>dot notation added for variables which allow accessing and writing to object properties. eg <code v-pre>$json.array.index</code></p>
</li>
<li>
<p>added support variable encapsulation via <code v-pre>$()</code> to avoid variable barewords from leaking into other values when infixed. eg <code v-pre>$(foobar)</code> is the same as <code v-pre>$foobar</code> (<RouterLink to="/parser/string.html">read more</RouterLink>)</p>
</li>
<li>
<p>variable names can now be Unicode, eg <code v-pre>$(比如) = &quot;举手之劳就可以使办公室更加环保，比如，使用再生纸。&quot;</code> (<RouterLink to="/parser/string.html">read more</RouterLink>)</p>
</li>
<li>
<p>Unicode wide characters are properly handled in readline</p>
</li>
<li>
<p>new data type <code v-pre>path</code> that is an object</p>
</li>
<li>
<p>new data type <code v-pre>paths</code></p>
</li>
<li>
<p>support for lambdas added to variables: eg <code v-pre>@array[{ function }]</code></p>
</li>
<li>
<p>readline <code v-pre>[ctrl]</code>+<code v-pre>[f]</code> now supports word matching as well as regexp (<RouterLink to="/user-guide/terminal-keys.html">read more</RouterLink>)</p>
</li>
<li>
<p>readline <code v-pre>[ctrl]</code>+<code v-pre>[f]</code> now supports glob filters (eg <code v-pre>*.txt</code>) too (<RouterLink to="/user-guide/terminal-keys.html">read more</RouterLink>)</p>
</li>
<li>
<p><code v-pre>$ENV</code> and <code v-pre>$GLOBAL</code> objects (deprecates the need for <code v-pre>global</code> and <code v-pre>export</code>)</p>
</li>
<li>
<p>improved error messages in expressions</p>
</li>
<li>
<p>scalars can be used as assignees in expressions, eg <code v-pre>$foo = &quot;bar&quot;</code></p>
</li>
<li>
<p>added <code v-pre>-i</code> flag to ssh autocompletion</p>
</li>
<li>
<p>readline prompts are now buffered to reduce cursor flickering</p>
</li>
</ul>
<h3 id="non-user-facing-maintenance-changes" tabindex="-1"><a class="header-anchor" href="#non-user-facing-maintenance-changes" aria-hidden="true">#</a> Non-User Facing / Maintenance Changes:</h3>
<ul>
<li>
<p>backend support added for data types to be dynamic objects</p>
</li>
<li>
<p>updated autocompletion code for <code v-pre>docker</code> and <code v-pre>git</code></p>
</li>
</ul>
<h3 id="bug-fixes" tabindex="-1"><a class="header-anchor" href="#bug-fixes" aria-hidden="true">#</a> Bug Fixes:</h3>
<ul>
<li>
<p>better wide character support in auto completion suggestions</p>
</li>
<li>
<p>function parameters default and comment fields are no longer order dependent</p>
</li>
<li>
<p>better string parsing for barewords inside arrays and objects</p>
</li>
<li>
<p>ahead of time hint caching now done on demand per hint page rather than against every executable upon Murex start up</p>
</li>
<li>
<p>new lines should be ignored in <code v-pre>%{}</code></p>
</li>
<li>
<p>bareword <code v-pre>true</code> and <code v-pre>false</code> should be converted to boolean in <code v-pre>%{}</code></p>
</li>
<li>
<p>autocomplete suggestions are now dedup'ed when delayed results return similar suggestions to fast results</p>
</li>
<li>
<p>various preview pane bug fixes</p>
</li>
<li>
<p>support for ANSI escape sequences in experimental (disabled by default) TTY buffer</p>
</li>
<li>
<p>correction to <code v-pre>count</code> docs</p>
</li>
<li>
<p>error not handled in regexp expression matches, `=~</p>
</li>
</ul>
<hr>
<p>Published: 26.03.2023 at 17:50</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/parser/string.html">String (<code v-pre>$</code>) Token</RouterLink>:
Expand values as a string</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/parser/create-object.html">create-object</RouterLink>:</li>
<li><RouterLink to="/types/path.html">path</RouterLink>:</li>
</ul>
</div></template>


