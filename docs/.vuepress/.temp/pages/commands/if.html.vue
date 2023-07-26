<template><div><h1 id="if" tabindex="-1"><a class="header-anchor" href="#if" aria-hidden="true">#</a> <code v-pre>if</code></h1>
<blockquote>
<p>Conditional statement to execute different blocks of code depending on the result of the condition</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Conditional control flow</p>
<p><code v-pre>if</code> can be utilized both as a method as well as a standalone function. As a
method, the conditional state is derived from the calling function (eg if the
previous function succeeds then the condition is <code v-pre>true</code>).</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<h3 id="function-if" tabindex="-1"><a class="header-anchor" href="#function-if" aria-hidden="true">#</a> Function <code v-pre>if</code>:</h3>
<pre><code>if { code-block } then {
    # true
} else {
    # false
}
</code></pre>
<h3 id="method-if" tabindex="-1"><a class="header-anchor" href="#method-if" aria-hidden="true">#</a> Method <code v-pre>if</code>:</h3>
<pre><code>command -&gt; if {
    # true
} else {
    # false
}
</code></pre>
<h3 id="negative-function-if" tabindex="-1"><a class="header-anchor" href="#negative-function-if" aria-hidden="true">#</a> Negative Function <code v-pre>if</code>:</h3>
<pre><code>!if { code-block } then {
    # false
}
</code></pre>
<h3 id="negative-method-if" tabindex="-1"><a class="header-anchor" href="#negative-method-if" aria-hidden="true">#</a> Negative Method <code v-pre>if</code>:</h3>
<pre><code>command -&gt; !if {
    # false
}
</code></pre>
<h3 id="please-note" tabindex="-1"><a class="header-anchor" href="#please-note" aria-hidden="true">#</a> Please Note:</h3>
<p>the <code v-pre>then</code> and <code v-pre>else</code> statements are optional. So the first usage could
also be written as:</p>
<pre><code>if { code-block } {
    # true
} {
    # false
}
</code></pre>
<p>However the practice of omitting those statements isn't recommended beyond
writing short one liners in the interactive command prompt.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Check if a file exists:</p>
<pre><code>if { g somefile.txt } then {
    out &quot;File exists&quot;
}
</code></pre>
<p>...or does not exist (both ways are valid):</p>
<pre><code>!if { g somefile.txt } then {
    out &quot;File does not exist&quot;
}

if { g somefile.txt } else {
    out &quot;File does not exist&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>The conditional block can contain entire pipelines - even multiple lines of code
let alone a single pipeline - as well as solitary commands as demonstrated in
the examples above. However the conditional block does not output STDOUT nor
STDERR to the rest of the pipeline so you don't have to worry about redirecting
the output streams to <code v-pre>null</code>.</p>
<p>If you require output from the conditional blocks STDOUT then you will need to
use either a Murex named pipe to redirect the output, or test or debug flags
(depending on your use case) if you only need to occasionally inspect the
conditionals output.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>if</code></li>
<li><code v-pre>!if</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/not.html"><code v-pre>!</code> (not)</RouterLink>:
Reads the STDIN and exit number from previous process and not's it's condition</li>
<li><RouterLink to="/commands/and.html"><code v-pre>and</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/debug.html"><code v-pre>debug</code></RouterLink>:
Debugging information</li>
<li><RouterLink to="/commands/false.html"><code v-pre>false</code></RouterLink>:
Returns a <code v-pre>false</code> value</li>
<li><RouterLink to="/commands/or.html"><code v-pre>or</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
<li><RouterLink to="/commands/test.html"><code v-pre>test</code></RouterLink>:
Murex's test framework - define tests, run tests and debug shell scripts</li>
<li><RouterLink to="/commands/true.html"><code v-pre>true</code></RouterLink>:
Returns a <code v-pre>true</code> value</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


