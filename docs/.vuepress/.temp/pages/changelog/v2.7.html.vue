<template><div><h1 id="what-s-new-in-murex-v2-7-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-7-change-log" aria-hidden="true">#</a> What's new in murex v2.7 - Change Log</h1>
<p>This update has introduced another potential breaking change for your safety: zero length arrays now fail by default. Also errors inside subshells will cause the parent command to fail if ran inside a <code v-pre>try</code> or <code v-pre>trypipe</code> block.</p>
<p>Breaking Changes:</p>
<ul>
<li>
<p>zero length arrays returned from subshells (eg <code v-pre>echo @{g this-file-does-not-exist}</code>) should fail by default, like unset variables. This is enabled by default but can be disabled via <code v-pre>config: set proc strict-arrays false</code></p>
</li>
<li>
<p>autoglob should fail if it doesn't match any results. eg <code v-pre>@g echo this-file-does-not-exist.*</code></p>
</li>
<li>
<p>Subshells should fail parent command when used inside <code v-pre>try</code> and <code v-pre>trypipe</code> blocks. eg <code v-pre>try { echo ${false} }</code></p>
</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p><code v-pre>function</code> now supports defining parameters</p>
</li>
<li>
<p>Added support fro <code v-pre>&amp;&amp;</code> eg <code v-pre>do-something &amp;&amp; do-something-else</code> for chaining successful commands</p>
</li>
<li>
<p>Added support for <code v-pre>||</code> eg <code v-pre>do-something || do-something-else</code> for chaining unsuccessful commands</p>
</li>
<li>
<p>Added support for writing to the terminal emulator's titlebar via <code v-pre>config: set shell titlebar-func { block }</code></p>
</li>
<li>
<p><code v-pre>titlebar-func</code> can also be written to your <code v-pre>tmux</code> window title via <code v-pre>config: set shell tmux-echo true</code>.</p>
</li>
<li>
<p>New reserved variable: <code v-pre>$HOSTNAME</code></p>
</li>
<li>
<p>New reserved variables: <code v-pre>$1</code> (and upwards) which correlates to the scope's parameter index. <code v-pre>$1</code> is the functions first parameter. <code v-pre>$2</code> is the second. <code v-pre>$13</code> is the thirteenth.</p>
</li>
<li>
<p>New reserved variable: <code v-pre>$0</code> which contains the function name</p>
</li>
<li>
<p>New event: <code v-pre>onCommandCompletion</code> (this is experimental and subject to change in the next release)</p>
</li>
<li>
<p>Macro variables. eg <code v-pre>echo Hello ^$name</code> will prompt the user to enter a name. Macro variables are only support in the REPL</p>
</li>
<li>
<p><code v-pre>read</code> now supports flags (eg default option, etc) to allow for a better experience in shell scripting</p>
</li>
</ul>
<p>Minor Changes:</p>
<ul>
<li>
<p>You can now overwrite <code v-pre>onKeyPress</code> events. This no longer produces an error forcing you to remove the old event before adding the new one</p>
</li>
<li>
<p>Autocompletion suggestions shouldn't be sorted is results include paths (improves the ordering of autocompletion suggestions)</p>
</li>
<li>
<p>Autocompletion suggestions for <code v-pre>openagent</code> builtin</p>
</li>
<li>
<p>Autocompletion suggestions for hashtags</p>
</li>
<li>
<p>Test counts re-added to website</p>
</li>
<li>
<p>Windows should show file extensions by default in autocompletion results</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li>
<p>Fix inverted logic on <code v-pre>forceTTY</code>: <code v-pre>config: get proc force-tty false</code> set by default, which then proxies STDERR and prints them in red</p>
</li>
<li>
<p>ctrl+c (^c) now currectly sends SIGTERM rather than just ending the child process</p>
</li>
<li>
<p>Better handling of SIGTERM</p>
</li>
<li>
<p>SIGTSTP isn't working. Switched to SIGSTOP when (^z) doesn't</p>
</li>
<li>
<p>Fix panic in event <code v-pre>onFilesystemChange</code> where fileRef is not getting passed correctly</p>
</li>
<li>
<p>Fix panic in event <code v-pre>onFilesystemChange</code> where path is zero length string</p>
</li>
<li>
<p>Some improvements to variable previews in the REPL</p>
</li>
<li>
<p><code v-pre>count</code> should check if it is a method</p>
</li>
<li>
<p>AST cache now checked more regukarly. This is to reduce the risk of memory leaks during fuzz or load testing</p>
</li>
<li>
<p><code v-pre>murex-docs</code> still referred to <code v-pre>len</code> builtin. That should be changed to <code v-pre>count</code></p>
</li>
<li>
<p>Lots of fuzzing added -- a few edge case bugs discovered</p>
</li>
</ul>
<hr>
<p>Published: 15.05.2022 at 22:49</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/logical-and.html">And (<code v-pre>&amp;&amp;</code>) Logical Operator</RouterLink>:
Continues next operation if previous operation passes</li>
<li><RouterLink to="/parser/logical-or.html">Or (<code v-pre>||</code>) Logical Operator</RouterLink>:
Continues next operation only if previous operation fails</li>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/commands/autoglob.html"><code v-pre>@g</code> (autoglob) </RouterLink>:
Command prefix to expand globbing (deprecated)</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/openagent.html"><code v-pre>openagent</code></RouterLink>:
Creates a handler function for `open</li>
<li><RouterLink to="/commands/read.html"><code v-pre>read</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a variable</li>
</ul>
</div></template>


