<template><div><h1 id="what-s-new-in-murex-v2-3-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-3-change-log" aria-hidden="true">#</a> What's new in murex v2.3 - Change Log</h1>
<p>This release includes significant changes to the interactive terminal</p>
<p>This release brings some major improvements to the discoverability of commands
and their parameters in the interactive terminal.</p>
<h3 id="user-facing-changes" tabindex="-1"><a class="header-anchor" href="#user-facing-changes" aria-hidden="true">#</a> User Facing Changes:</h3>
<ul>
<li>
<p>New builtin: <code v-pre>method</code>. This is used to help turn some commands into a
&quot;method&quot;, similar in concept to object orientated development where the tab-
completer will return a list of other commands that are compatible with the
formers output data types. It's purely an autocompletion tool and has no
effect in shell scripts.</p>
</li>
<li>
<p>Autocompletion suggestions will differ based on what pipes are used. The idea
behind this is to improve the discoverability of commands. All tokens can
still be used to pipeline to any command type (like before). This rule only
affects autocomplete suggestions in the interactive prompt.</p>
<pre><code>- `|` will show all available aliases, functions, builtins and external
  commands, as usual

- `-&gt;` will default to showing only compatible commands defined in `method`
  (see above), however if no methods are found then `-&gt;` will fall back to
  suggesting all available aliases, functions, builtins and external commands,
  like `|`.

The idea here is for `-&gt;` to surface commands that are commonly
expected to be pipelined together but without forcing the user to pause and
remember which token exposes what commands. While `|` is still available for
familiarity
</code></pre>
</li>
<li>
<p>Autocompletion de-duplication now places hyphen-prefixed flags at the end of
the results when sorting. This helps bring more relevant flags forward, with
hyphen-prefixed flags still visible when users press <code v-pre>-</code></p>
</li>
<li>
<p>External commands return a error message if the exit status is non-zero
(<a href="https://github.com/lmorg/murex/issues/299" target="_blank" rel="noopener noreferrer">issue #299<ExternalLinkIcon/></a>)</p>
<pre><code>&lt;pre&gt;&lt;code&gt;» exec: false
Error in `exec` (1,1): exit status 1&lt;/code&gt;&lt;/pre&gt;
</code></pre>
</li>
<li>
<p><code v-pre>Index()</code> (<code v-pre>[</code>) now returns a different data type from the root object if the
element is a string, integer, float64, boolean or nil (<a href="https://github.com/lmorg/murex/issues/316" target="_blank" rel="noopener noreferrer">issue #316<ExternalLinkIcon/></a>)</p>
</li>
<li>
<p><code v-pre>runtime</code> was missing a few flags for reporting supported data types against
the newer APIs</p>
</li>
<li>
<p>Generic types (<code v-pre>*</code> / <code v-pre>generic</code>) to include two spaces between columns when
marshalled</p>
</li>
<li>
<p>Syntax highlighter colour tweaks -- no longer using background highlighting
because that looks weird and unreadable on some terminal colour schemes</p>
</li>
<li>
<p>Bug fixes to the fast parser</p>
</li>
<li>
<p>Bug fixes to readline merged from <code v-pre>v2.2.1200</code> (<a href="https://github.com/lmorg/murex/issues/362" target="_blank" rel="noopener noreferrer">issue #362<ExternalLinkIcon/></a>)</p>
</li>
</ul>
<h3 id="non-user-facing-maintenance-changes" tabindex="-1"><a class="header-anchor" href="#non-user-facing-maintenance-changes" aria-hidden="true">#</a> Non-User Facing / Maintenance Changes</h3>
<ul>
<li>
<p><code v-pre>hcl</code> and <code v-pre>bson</code> types now use <code v-pre>ReadArray</code> templates so their behavior should
mimic JSON and YAML</p>
</li>
<li>
<p>Fast parser is now aware of which pipe token (eg <code v-pre>|</code>, <code v-pre>-&gt;</code>, <code v-pre>=&gt;</code>, <code v-pre>?</code>) is used</p>
</li>
<li>
<p>Unit test concurrency fix</p>
</li>
<li>
<p>New tests written for the fast parser</p>
</li>
<li>
<p>Refactored code around hint text and completions to allow for a potential new
UI feature to be built in the future</p>
</li>
<li>
<p>Updated documentation</p>
</li>
</ul>
<hr>
<p>Published: 26.09.2021 at 11:33</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/commands/method.html"><code v-pre>method</code></RouterLink>:
Define a methods supported data-types</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


