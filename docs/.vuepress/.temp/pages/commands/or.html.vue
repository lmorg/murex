<template><div><h1 id="or" tabindex="-1"><a class="header-anchor" href="#or" aria-hidden="true">#</a> <code v-pre>or</code></h1>
<blockquote>
<p>Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Returns a boolean results (<code v-pre>true</code> or <code v-pre>false</code>) depending on whether any of the
code-blocks included as parameters are successful or not.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>or: { code-block } { code-block } -&gt; `&lt;stdout&gt;`

!or: { code-block } { code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>or</code> supports as many or as few code-blocks as you wish.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>if { or { = 1+1==2 } { = 2+2==5 } { = 3+3==6 } } then {
    out: At least one of those equations are correct
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>or</code> does not set the exit number on failure so it is safe to use inside a <code v-pre>try</code>
or <code v-pre>trypipe</code> block.</p>
<p>If <code v-pre>or</code> is prefixed by a bang (<code v-pre>!or</code>) then it returns <code v-pre>true</code> when one or more
code-blocks are unsuccessful (ie the opposite of <code v-pre>or</code>).</p>
<h3 id="code-block-testing" tabindex="-1"><a class="header-anchor" href="#code-block-testing" aria-hidden="true">#</a> Code-Block Testing</h3>
<ul>
<li>
<p><code v-pre>or</code> only executes code-blocks up until one of the code-blocks is successful
then it exits the function and returns <code v-pre>true</code>.</p>
</li>
<li>
<p><code v-pre>!or</code> only executes code-blocks while the code-blocks are successful. Once one
is unsuccessful <code v-pre>!or</code> exits and returns <code v-pre>true</code> (ie it <code v-pre>not</code>s every code-block).</p>
</li>
</ul>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>or</code></li>
<li><code v-pre>!or</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/not.html"><code v-pre>!</code> (not)</RouterLink>:
Reads the STDIN and exit number from previous process and not's it's condition</li>
<li><RouterLink to="/commands/and.html"><code v-pre>and</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/false.html"><code v-pre>false</code></RouterLink>:
Returns a <code v-pre>false</code> value</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/true.html"><code v-pre>true</code></RouterLink>:
Returns a <code v-pre>true</code> value</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


