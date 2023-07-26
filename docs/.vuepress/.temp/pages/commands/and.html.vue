<template><div><h1 id="and" tabindex="-1"><a class="header-anchor" href="#and" aria-hidden="true">#</a> <code v-pre>and</code></h1>
<blockquote>
<p>Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Returns a boolean results (<code v-pre>true</code> or <code v-pre>false</code>) depending on whether all of the
code-blocks included as parameters are successful or not.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>and: { code-block } { code-block } -&gt; `&lt;stdout&gt;`

!and: { code-block } { code-block } -&gt; `&lt;stdout&gt;`
</code></pre>
<p><code v-pre>and</code> supports as many or as few code-blocks as you wish.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>if { and { = 1+1==2 } { = 2+2==4 } { = 3+3==6 } } then {
    out: The laws of mathematics still exist in this universe.
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>and</code> does not set the exit number on failure so it is safe to use inside a <code v-pre>try</code>
or <code v-pre>trypipe</code> block.</p>
<p>If <code v-pre>and</code> is prefixed by a bang then it returns <code v-pre>true</code> only when all code-blocks
are unsuccessful.</p>
<h3 id="code-block-testing" tabindex="-1"><a class="header-anchor" href="#code-block-testing" aria-hidden="true">#</a> Code-Block Testing</h3>
<ul>
<li>
<p><code v-pre>and</code> tests all code-blocks up until one of the code-blocks is unsuccessful,
then <code v-pre>and</code> exits and returns <code v-pre>false</code>.</p>
</li>
<li>
<p><code v-pre>!and</code> tests all code-blocks up until one of the code-blocks is successful,
then <code v-pre>!and</code> exits and returns <code v-pre>false</code> (ie <code v-pre>!and</code> is <code v-pre>not</code>ing every code-block).</p>
</li>
</ul>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>and</code></li>
<li><code v-pre>!and</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/not.html"><code v-pre>!</code> (not)</RouterLink>:
Reads the STDIN and exit number from previous process and not's it's condition</li>
<li><RouterLink to="/commands/catch.html"><code v-pre>catch</code></RouterLink>:
Handles the exception code raised by <code v-pre>try</code> or <code v-pre>trypipe</code></li>
<li><RouterLink to="/commands/false.html"><code v-pre>false</code></RouterLink>:
Returns a <code v-pre>false</code> value</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/or.html"><code v-pre>or</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</li>
<li><RouterLink to="/commands/true.html"><code v-pre>true</code></RouterLink>:
Returns a <code v-pre>true</code> value</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


