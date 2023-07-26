<template><div><h1 id="non-decimal-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#non-decimal-ranges-mkarray" aria-hidden="true">#</a> Non-Decimal Ranges - mkarray</h1>
<blockquote>
<p>Create arrays of integers from non-decimal number bases</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>When making arrays you can specify ranges of an alternative number base by
using an <code v-pre>x</code> or <code v-pre>.</code> in the end range:</p>
<pre><code>a: [00..ffx16]
a: [00..ff.16]
</code></pre>
<p>All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: <code v-pre>[hex..hex.dec]</code></p>
<p>Also note that the additional zeros denotes padding (ie the results will
start at <code v-pre>00</code>, <code v-pre>01</code>, etc rather than <code v-pre>0</code>, <code v-pre>1</code>...)</p>
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
<pre><code>» a: [08..10x16]
08
09
0a
0b
0c
0d
0e
0f
10

» a: [10..08x16]
10
f
e
d
c
b
a
9
8
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="floating-point-numbers" tabindex="-1"><a class="header-anchor" href="#floating-point-numbers" aria-hidden="true">#</a> Floating Point Numbers</h3>
<p>If you do need a range of fixed floating point numbers generated then you can
do so by merging two decimal integer ranges together. For example</p>
<pre><code>» a: [05..10x8].[0..7]
05.0
05.1
05.2
05.3
05.4
05.5
05.6
05.7
06.0
06.1
06.2
...
07.5
07.6
07.7
10.0
10.1
10.2
10.3
10.4
10.5
10.6
10.7
</code></pre>
<h3 id="everything-is-a-string" tabindex="-1"><a class="header-anchor" href="#everything-is-a-string" aria-hidden="true">#</a> Everything Is A String</h3>
<p>Please note that all arrays are created as strings. Even when using typed
arrays such as JSON (<code v-pre>ja</code>).</p>
<pre><code>» ja [0..5]
[
    &quot;0&quot;,
    &quot;1&quot;,
    &quot;2&quot;,
    &quot;3&quot;,
    &quot;4&quot;,
    &quot;5&quot;
]
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/mkarray/character.html">Character arrays</RouterLink>:
Making character arrays (a to z)</li>
<li><RouterLink to="/commands/mkarray/decimal.html">Decimal Ranges</RouterLink>:
Create arrays of decimal integers</li>
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


