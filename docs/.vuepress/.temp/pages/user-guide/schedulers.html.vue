<template><div><h1 id="schedulers" tabindex="-1"><a class="header-anchor" href="#schedulers" aria-hidden="true">#</a> Schedulers</h1>
<blockquote>
<p>Overview of the different schedulers (or 'run modes') in Murex</p>
</blockquote>
<p>There are a few distinct schedulers (or run modes) in Murex which are invoked
by builtin commands. This means you can alter the way commands are executed
dynamically within Murex shell scripts.</p>
<h2 id="normal" tabindex="-1"><a class="header-anchor" href="#normal" aria-hidden="true">#</a> Normal</h2>
<p>This is a traditional shell where anything in a pipeline (eg <code v-pre>cmd1 -&gt; cmd2 -&gt; cmd3</code>)
is executed in parallel. The scheduler only pauses launching new commands when
the last command in any pipeline is still executing. A pipeline could be multiple
commands (like above) or a single command (eg <code v-pre>top</code>).</p>
<h2 id="try" tabindex="-1"><a class="header-anchor" href="#try" aria-hidden="true">#</a> Try</h2>
<p>This is similar to normal where commands in a pipeline are run in parallel except
Murex validates the stderr and exit status of the last command in any pipeline.</p>
<p>If stderr is greater than stdout (per bytes written) <strong>OR</strong> the exit status is
non-zero then the scheduler exits that entire block.</p>
<h2 id="try-pipe" tabindex="-1"><a class="header-anchor" href="#try-pipe" aria-hidden="true">#</a> Try Pipe</h2>
<p>This runs the commands sequentially because the stderr and the exit status of
each command is checked irrespective of whether that command is at the start of
the pipeline (eg <code v-pre>start -&gt; middle -&gt; end</code>), or anywhere else.</p>
<p>Like with <code v-pre>try</code>, if stderr is greater than stdout (per bytes written) <strong>OR</strong>
the exit status is non-zero then the scheduler exits that entire block. Unlike
with <code v-pre>try</code>, this check happens on every command rather than the last command in
the pipeline.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-generic.html">Generic Pipe (<code v-pre>=&gt;</code>) Token</RouterLink>:
Pipes a reformatted STDOUT stream from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/parser/pipe-posix.html">POSIX Pipe (<code v-pre>|</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/parser/pipe-err.html">STDERR Pipe (<code v-pre>?</code>) Token</RouterLink>:
Pipes STDERR from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/commands/runmode.html"><code v-pre>runmode</code></RouterLink>:
Alter the scheduler's behaviour at higher scoping level</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


