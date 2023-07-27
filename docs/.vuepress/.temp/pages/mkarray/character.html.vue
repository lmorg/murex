<template><div><h1 id="character-arrays-mkarray" tabindex="-1"><a class="header-anchor" href="#character-arrays-mkarray" aria-hidden="true">#</a> Character arrays - mkarray</h1>
<blockquote>
<p>Making character arrays (a to z)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>You can create arrays from a range of letters (a to z):</p>
<pre><code>» a: [a..z]
» a: [z..a]
» a: [A..Z]
» a: [Z..A]
</code></pre>
<p>...or any characters within that range.</p>
<p>Please refer to <RouterLink to="/commands/a.html">a (mkarray)</RouterLink> for more detailed usage of mkarray.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>a: [start..end] -&gt; `&lt;stdout&gt;`
a: [start..end,start..end] -&gt; `&lt;stdout&gt;`
a: [start..end][start..end] -&gt; `&lt;stdout&gt;`
</code></pre>
<p>All usages also work with <code v-pre>ja</code> and <code v-pre>ta</code> as well, eg:</p>
<pre><code>ja: [start..end] -&gt; `&lt;stdout&gt;`
ta: data-type [start..end] -&gt; `&lt;stdout&gt;`
</code></pre>
<p>You can also inline arrays with the <code v-pre>%[]</code> syntax, eg:</p>
<pre><code>%[start..end]
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» a: [a..c]
a
b
c

» a: [c..a]
c
b
a
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/mkarray/decimal.html">Decimal Ranges</RouterLink>:
Create arrays of decimal integers</li>
<li><RouterLink to="/mkarray/non-decimal.html">Non-Decimal Ranges</RouterLink>:
Create arrays of integers from non-decimal number bases</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


