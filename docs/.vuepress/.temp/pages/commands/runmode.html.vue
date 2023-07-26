<template><div><h1 id="runmode" tabindex="-1"><a class="header-anchor" href="#runmode" aria-hidden="true">#</a> <code v-pre>runmode</code></h1>
<blockquote>
<p>Alter the scheduler's behaviour at higher scoping level</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Due to dynamic nature in which blocks are compiled on demand, traditional <code v-pre>try</code>
and <code v-pre>trypipe</code> blocks cannot affect the runtime behaviour of schedulers already
invoked (eg for function blocks and modules which <code v-pre>try</code> et al would sit inside).
To solve this we need an additional command that is executed by the compiler
prior to the block being executed which can define the runmode of the scheduler.
This is the purpose of <code v-pre>runmode</code>.</p>
<p>The caveat of being a compiler command rather than a builtin is that <code v-pre>runmode</code>
needs be the first command in a block.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>runmode try|trypipe function|module
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>function hello {
    # Short conversation, exit on error

    runmode: try function

    read: name &quot;What is your name? &quot;
    out: &quot;Hello $name, pleased to meet you&quot;

    read: mood &quot;How are you feeling? &quot;
    out: &quot;I'm feeling $mood too&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>runmode</code>'s parameters are ordered:</p>
<h3 id="_1st-parameter" tabindex="-1"><a class="header-anchor" href="#_1st-parameter" aria-hidden="true">#</a> 1st parameter</h3>
<h4 id="try" tabindex="-1"><a class="header-anchor" href="#try" aria-hidden="true">#</a> try</h4>
<p>Checks only the last command in the pipeline for errors. However still allows
commands in a pipeline to run in parallel.</p>
<h4 id="trypipe" tabindex="-1"><a class="header-anchor" href="#trypipe" aria-hidden="true">#</a> trypipe</h4>
<p>Checks every command in the pipeline before executing the next. However this
blocks pipelines from running every command in parallel.</p>
<h3 id="_2nd-parameter" tabindex="-1"><a class="header-anchor" href="#_2nd-parameter" aria-hidden="true">#</a> 2nd parameter</h3>
<h4 id="function" tabindex="-1"><a class="header-anchor" href="#function" aria-hidden="true">#</a> function</h4>
<p>Sets the runmode for all blocks within the function when <code v-pre>runmode</code> is placed at
the start of the function. This includes privates, autocompletes, events, etc.</p>
<h4 id="module" tabindex="-1"><a class="header-anchor" href="#module" aria-hidden="true">#</a> module</h4>
<p>Sets the runmode for all blocks within that module when placed at the start of
the module. This include any functions, privates, autocompletes, events, etc
that are inside that module. The do not need a separate <code v-pre>runmode ... function</code>
if <code v-pre>runmode ... module</code> is set.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/user-guide/schedulers.html">Schedulers</RouterLink>:
Overview of the different schedulers (or 'run modes') in Murex</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/fid-list.html"><code v-pre>fid-list</code></RouterLink>:
Lists all running functions within the current Murex session</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/read.html"><code v-pre>read</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a variable</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


