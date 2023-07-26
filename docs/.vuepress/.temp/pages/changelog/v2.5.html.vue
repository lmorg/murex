<template><div><h1 id="what-s-new-in-murex-v2-5-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-5-change-log" aria-hidden="true">#</a> What's new in murex v2.5 - Change Log</h1>
<p>This release introduces a number of new builtins, fixes some regression bugs and supercharges the <code v-pre>select</code> optional builtin (which I plan to include into the core builtins for non-Windows users in the next release).</p>
<p>Features:</p>
<ul>
<li>
<p><code v-pre>alter</code> now supports <code v-pre>--sum</code> where structures are merged and numeric values are added together</p>
</li>
<li>
<p>New builtin <code v-pre>count</code>. This has deprecated <code v-pre>len</code> however <code v-pre>len</code> will stick around as an alias for backwards compatibility</p>
</li>
<li>
<p>New operators added to <code v-pre>let</code>: <code v-pre>+=</code>, <code v-pre>-=</code>, <code v-pre>/=</code>, <code v-pre>*=</code></p>
</li>
<li>
<p>New builtin <code v-pre>addheading</code> for adding headings to lists</p>
</li>
<li>
<p>Compiled profile will now always execute even if Murex flags set to ignore the modules/user profile. This is so that aliases like <code v-pre>len</code> get set and thus Murex can still move forward with changes to builtins but without breaking backwards compatibility</p>
</li>
<li>
<p><code v-pre>autocomplete</code> now passes <code v-pre>ISMETHOD</code> variable to dynamic completions so those dynamic completions are aware if a command requesting auto-completion suggestions is being invoked as a method (mid-pipeline) or function (start of a pipeline)</p>
</li>
<li>
<p>Index, <code v-pre>[</code>, now supports inlining element, <code v-pre>[[</code>, lookups. eg <code v-pre>[ a b [/map/c] ]</code></p>
</li>
<li>
<p>Dynamic auto-completions that include <code v-pre>@IncFiles</code> or <code v-pre>@IncDirs</code> will now automatically append files and/or directories to their auto-completion suggestions</p>
</li>
<li>
<p>New <code v-pre>autocomplete</code> directives: <code v-pre>IncExeAll</code> (like <code v-pre>IncExePath</code> but includes builtins, functions, aliases), <code v-pre>IncManPage</code> (include results from the <code v-pre>man</code> page parser -- usually suppressed when <code v-pre>autocomplete</code> config is set)</p>
</li>
<li>
<p>Disabled 'Loading default profile' message -- this was always pretty redundant but now that the compiled profile is being loaded all the time (eg <code v-pre>murex -c 'command'</code> or when called in a shebang), it's also now ugly too</p>
</li>
<li>
<p><code v-pre>select</code> now supports passing a file in the <strong>FROM</strong> syntax. eg <code v-pre>select * FROM ./example.csv</code>. The caveat here is this breaks currently auto-complete on column names</p>
</li>
<li>
<p><code v-pre>select</code> now supports multiple tables using either named pipes (eg <code v-pre>select * FROM &lt;table1&gt;, &lt;table2&gt;</code>) or variables (eg <code v-pre>select * FROM \$table1, \$table2</code>) passed in the <strong>FROM</strong> syntax. Variables should be escaped and you cannot mix and match between named pipes, file names nor variables. You can use any number of tables from 1 to 2^63-1 (64bit systems) or 1 to 2^31-1 (32bit systems). Which should be more than enough 😉</p>
</li>
<li>
<p><code v-pre>config</code> option for <code v-pre>select</code> to define default output data type where multiple tables are imported</p>
</li>
<li>
<p>Lots of new and updated documentation!</p>
</li>
</ul>
<p>Non-user facing changes (internal changes to the Murex code base):</p>
<ul>
<li>
<p><code v-pre>open</code> functions can now be called by other functions to take advantage of auto-typing and auto gunzip etc.</p>
</li>
<li>
<p><code v-pre>tmp.Close()</code> should return <code v-pre>err</code>. This isn't a bug but it might catch future bugs</p>
</li>
<li>
<p><code v-pre>LazyLogging</code> created to speed up writing tests against data structures</p>
</li>
<li>
<p><code v-pre>utils/List</code> package created to handle list / array / map functions. Also makes testing more complex routines easier</p>
</li>
</ul>
<p>Bug fixes:</p>
<ul>
<li>
<p>Regression bug fixed where <code v-pre>prepend</code> was invoking <code v-pre>append</code></p>
</li>
<li>
<p><code v-pre>streams.ReadCloser</code> not setting context</p>
</li>
<li>
<p><code v-pre>parameters.StringArray()</code> should copy values instead of a pointer to ensure the underlying parameters are immutable</p>
</li>
</ul>
<hr>
<p>Published: 12.02.2022 at 16:16</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/commands/namedpipe.html"><code v-pre>&lt;&gt;</code> / <code v-pre>read-named-pipe</code></RouterLink>:
Reads from a Murex named pipe</li>
<li><RouterLink to="/commands/alter.html"><code v-pre>alter</code></RouterLink>:
Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/pipe.html"><code v-pre>pipe</code></RouterLink>:
Manage Murex named pipes</li>
<li><RouterLink to="/optional/select.html"><code v-pre>select</code> </RouterLink>:
Inlining SQL into shell pipelines</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


