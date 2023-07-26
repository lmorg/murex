<template><div><h1 id="what-s-new-in-murex-v2-1-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-1-change-log" aria-hidden="true">#</a> What's new in murex v2.1 - Change Log</h1>
<p>This release comes with support for inlining SQL and some major bug fixes plus a breaking change for <code v-pre>config</code>. Please read for details.</p>
<p>This release sees new optional features plus major bug fixes to the existing
code base:</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes:</h3>
<p>Two <code v-pre>config</code> <strong>shell</strong> keys have changed names:</p>
<ul>
<li>recursive-soft-timeout -&gt; <code v-pre>autocomplete-soft-timeout</code></li>
<li>recursive-hard-timeout -&gt; <code v-pre>autocomplete-hard-timeout</code></li>
</ul>
<p>This is to better describe their functionality now that those values are
also used for <code v-pre>Dynamic</code> and <code v-pre>DynamicDesc</code> autocompletions as well as
recursive directory lookups.</p>
<p><strong>This change might break some of your existing profile config!</strong></p>
<h3 id="user-facing-changes" tabindex="-1"><a class="header-anchor" href="#user-facing-changes" aria-hidden="true">#</a> User Facing Changes:</h3>
<ul>
<li>
<p><code v-pre>config</code> <strong>shell</strong> <strong>max-suggestions</strong> now defaults at <code v-pre>12</code> rather than 6</p>
</li>
<li>
<p>New optional builtin, <code v-pre>select</code>, allows you to inline SQL queries against
any tabulated output (eg <code v-pre>ps -fe</code>, jsonlines arrays, CSV files, etc). This
works via importing output into an in memory sqlite3 database. However this
also breaks cross compiling due to the C includes with sqlite3. Thus this
builtin will remain optional for now.</p>
</li>
<li>
<p>Rethink of how optionals are imported. Rather than modifying <code v-pre>// +build</code>
headers in <code v-pre>.go</code> files, optionals can be copied (or symlinked) from
<code v-pre>builtins/imports_src</code> -&gt; <code v-pre>builtins/imports_build</code>. This enables us to
write a user friendly pre-compiling build script to enable users to easily
select which optional builtins to include.</p>
</li>
<li>
<p>Stopping jobs via <code v-pre>^z</code> has been fixed in UNIX. This was a regression bug
introduced a while back however no tests were in place to catch it.
Unfortunately this kind of testing would fall outside of unit testing each
function so I'll need to add another layer of testing against the compiled
executable to verify any future regressions like these: <a href="https://github.com/lmorg/murex/issues/318" target="_blank" rel="noopener noreferrer">discussion<ExternalLinkIcon/></a>
To use this feature, run a command and then press <code v-pre>^z</code> (ctrl+z) to pause
the process. You can check which jobs have been paused via <code v-pre>jobs</code> and/or
modify processes to run in the background/foreground via <code v-pre>bg</code> and <code v-pre>fg</code>.</p>
</li>
<li>
<p>Added new API endpoints: <code v-pre>ReadArrayWithType()</code>. This solves some edge cases
in <code v-pre>foreach</code> where elements might not match the same data type as the parent
object (eg a <code v-pre>json</code> object might have <code v-pre>int</code> or <code v-pre>str</code> elements in an array)</p>
</li>
<li>
<p>Rewritten how <code v-pre>Dynamic</code> autocompletions are executed to fall in line with
<code v-pre>DynamicDesc</code>. This should bring improvements to running autocompletions
in the background and thus improve the user experience with regards to the
shell's responsiveness. The next step would be to have a lower soft-timeout</p>
</li>
<li>
<p>Improvements to the context completions</p>
</li>
<li>
<p>Default lengths for autocompletions where all results are deferred to the
background have been tweaked slightly to give some extra length</p>
</li>
<li>
<p>Minor website tweaks</p>
</li>
</ul>
<h3 id="non-user-facing-maintenance-changes" tabindex="-1"><a class="header-anchor" href="#non-user-facing-maintenance-changes" aria-hidden="true">#</a> Non-User Facing / Maintenance Changes</h3>
<ul>
<li>
<p>All dependencies have been updated, pinned and the <code v-pre>vendors</code> directory
rebuilt</p>
</li>
<li>
<p>Fixed some issues flagged up in <a href="https://goreportcard.com/report/github.com/lmorg/murex" target="_blank" rel="noopener noreferrer">goreportcard.com<ExternalLinkIcon/></a></p>
</li>
<li>
<p>Some internal API changes that have no UI/UX ramifications but makes the
code more maintainable</p>
</li>
<li>
<p>Lots more unit tests added</p>
</li>
</ul>
<hr>
<p>Published: 30.04.2021 at 10:00</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/apis/ReadArrayWithType.html"><code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
<li><RouterLink to="/commands/bg.html"><code v-pre>bg</code></RouterLink>:
Run processes in the background</li>
<li><RouterLink to="/types/csv.html"><code v-pre>csv</code> </RouterLink>:
CSV files (and other character delimited tables)</li>
<li><RouterLink to="/commands/fg.html"><code v-pre>fg</code></RouterLink>:
Sends a background process into the foreground</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>jobs</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/optional/select.html"><code v-pre>select</code> </RouterLink>:
Inlining SQL into shell pipelines</li>
</ul>
</div></template>


