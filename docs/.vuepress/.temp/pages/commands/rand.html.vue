<template><div><h1 id="rand" tabindex="-1"><a class="header-anchor" href="#rand" aria-hidden="true">#</a> <code v-pre>rand</code></h1>
<blockquote>
<p>Random field generator</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>rand</code> can generate random numbers, strings and other data types.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>rand data-type [ max-value ]
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Random integer: 64-bit on 64-bit machines</p>
<pre><code>rand int
</code></pre>
<p>Random integer between 0 and 9</p>
<pre><code>rand int 9
</code></pre>
<p>Random floating point number between 0 and 1</p>
<pre><code>rand float
</code></pre>
<p>Random string, fixed to 5 characters long</p>
<pre><code>rand str 5
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>*</code>
Alias for <strong>str</strong> (string)</li>
<li><code v-pre>float</code>
Random floating point number between 0 and 1. <strong>max-value</strong> not supported</li>
<li><code v-pre>int</code>
Random integer, <strong>max-value</strong> specifies highest value for integer</li>
<li><code v-pre>num</code>
Alias for <strong>int</strong> in this builtin</li>
<li><code v-pre>str</code>
Random string of printable ASCII characters. <strong>max-value</strong> specifies the fixed length of the string</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="security" tabindex="-1"><a class="header-anchor" href="#security" aria-hidden="true">#</a> Security</h3>
<p>WARNING: is should be noted that while <code v-pre>rand</code> can produce random numbers and
strings which might be useful for password generation, neither the RNG nor the
the random string generator (which is ostensibly the same RNG but applied to an
array of bytes within the range of printable ASCII characters) are considered
cryptographically secure.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


