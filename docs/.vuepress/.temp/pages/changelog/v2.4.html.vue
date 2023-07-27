<template><div><h1 id="what-s-new-in-murex-v2-4-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-4-change-log" aria-hidden="true">#</a> What's new in murex v2.4 - Change Log</h1>
<p>This release introduces a strict mode for variables, new builtin, performance improvements, and better error messages; plus a potential breaking change</p>
<p>There are a number of new features in this release</p>
<h3 id="breaking-changes" tabindex="-1"><a class="header-anchor" href="#breaking-changes" aria-hidden="true">#</a> Breaking Changes:</h3>
<ul>
<li>mkarray (<code v-pre>a</code> et al) no longer returns an error if the start and end ranges
are the same. Instead it returns an array of 1 value.</li>
</ul>
<h3 id="user-facing-changes" tabindex="-1"><a class="header-anchor" href="#user-facing-changes" aria-hidden="true">#</a> User Facing Changes:</h3>
<ul>
<li>
<p>Strict variables now supported (like <code v-pre>set -u</code> in Bash). This will mean any
variables passed that haven't been initialized beforehand will cause that
pipeline to fail. Currently this is opt in, a future release of Murex will
flip that and make it opt out. So take this opportunity to enable it in your
<code v-pre>~/.murex_profile</code> and test your scripts. Enable this via <code v-pre>config</code>:</p>
<pre><code>&lt;pre&gt;&lt;code&gt;config: set proc strict-vars true&lt;/code&gt;&lt;/pre&gt;

This feature was requested in [issue #380](https://github.com/lmorg/murex/issues/380).
</code></pre>
</li>
<li>
<p>New builtin: <code v-pre>datetime</code>. This builtin allows you to convert date and/or time
strings of one format into strings of another format. <code v-pre>datetime</code> is a
supercharged alternative to the UNIX command <code v-pre>date</code> aimed at making scripting
easier.</p>
</li>
<li>
<p>mkarray (<code v-pre>a</code> et al) now supports dates. eg <code v-pre>[01-Jan-20..05-May-21]</code>. If no
start nor end date appears then mkarray assumes range starts or ends at
current date.</p>
</li>
<li>
<p><code v-pre>openagent</code> profile for <code v-pre>image</code> data types has been improved. Murex now
better supports tmux and iTerm2.</p>
</li>
<li>
<p><code v-pre>runtime --config</code> now displays <code v-pre>FileRef</code> for every <code v-pre>set</code> as well as <code v-pre>define</code>,
named <code v-pre>FileRefSet</code> and <code v-pre>FileRefDefine</code> respectively. So you can now easily
trace where global config is being set and defined.</p>
</li>
<li>
<p>Better error messages in the interactive terminal.</p>
</li>
<li>
<p>Prompt now defaults to only displaying current directory rather than the full
path. You can revert this change by adding your own prompt in <code v-pre>config</code>. eg:</p>
<pre><code>&lt;pre&gt;&lt;code&gt;config: set shell prompt {
    out &quot;{RESET}{YELLOW}${pwd_short} {RESET}» &quot;
}

config: set shell prompt-multiline {
    let len = ${pwd_short -&gt; wc -c} - 1
    printf &quot;%${$len}s » &quot; $linenum
}&lt;/code&gt;&lt;/pre&gt;
</code></pre>
</li>
<li>
<p>Parser updated to better support multiline pipelines where the newline is
escaped and a comment exists after <a href="https://github.com/lmorg/murex/issues/379" target="_blank" rel="noopener noreferrer">issue #379<ExternalLinkIcon/></a>.
This only applies to shell scripts, the interactive terminal hasn't yet been
updated to reflect this change.</p>
</li>
<li>
<p>Fixed regression bugs with autocomplete parameters that affected some dynamic
blocks.</p>
</li>
<li>
<p><code v-pre>readline</code> now caches syntax highlighting and hint text to improve the
responsiveness of the interactive terminal. This is set to a hard limit of
200 cached items however that will be a configurable metric in a future
release. Also planned for the future is caching autocompletion suggestions.</p>
</li>
<li>
<p>Loading message added for the default profile, ie the one that is compiled
into and thus shipped with Murex.</p>
</li>
<li>
<p>Fixed bug with <code v-pre>fid-list</code> and <code v-pre>jobs</code> where they were outputting the <code v-pre>p.Name</code>
struct rather than <code v-pre>p.Name.String()</code>. This lead to the process name appearing
garbled under some circumstances.</p>
</li>
<li>
<p><code v-pre>{BG-BLUE}</code> emitted the wrong ANSI escape code, this has been corrected.</p>
</li>
<li>
<p>Several <code v-pre>readline</code> bug fixes.</p>
</li>
</ul>
<h3 id="non-user-facing-maintenance-changes" tabindex="-1"><a class="header-anchor" href="#non-user-facing-maintenance-changes" aria-hidden="true">#</a> Non-User Facing / Maintenance Changes</h3>
<ul>
<li>
<p>Thread safe copying of parameters upon fork. The previous code never actually
generated any race conditions and I don't think ever could. However it was
ambiguous. This new code makes the copy more explicit and appears to have
also brought some minor performance improvements in benchmarks too.</p>
</li>
<li>
<p>Behavioural test framework has been refactored to make it easier to add new
behavioural tests.</p>
</li>
<li>
<p>Lots of new tests added.</p>
</li>
<li>
<p>Updated documentation.</p>
</li>
</ul>
<hr>
<p>Published: 09.12.2021 at 08:00</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/ansi.html">ANSI Constants</RouterLink>:
Infixed constants that return ANSI escape sequences</li>
<li><RouterLink to="/user-guide/fileref.html">FileRef</RouterLink>:
How to track what code was loaded and from where</li>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/user-guide/profile.html">Murex Profile Files</RouterLink>:
A breakdown of the different files loaded on start up</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/datetime.html"><code v-pre>datetime</code> </RouterLink>:
A date and/or time conversion tool (like <code v-pre>printf</code> but for date and time values)</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/openagent.html"><code v-pre>openagent</code></RouterLink>:
Creates a handler function for `open</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


