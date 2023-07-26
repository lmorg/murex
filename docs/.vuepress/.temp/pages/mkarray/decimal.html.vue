<template><div><h1 id="decimal-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#decimal-ranges-mkarray" aria-hidden="true">#</a> Decimal Ranges - mkarray</h1>
<blockquote>
<p>Create arrays of decimal integers</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This document describes how to create arrays of decimals using mkarray (<code v-pre>a</code> et
al).</p>
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
<pre><code>» a: [1..3]
1
2
3

» a: [3..1]
3
2
1

» a: [01..03]
01
02
03
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="floating-point-numbers" tabindex="-1"><a class="header-anchor" href="#floating-point-numbers" aria-hidden="true">#</a> Floating Point Numbers</h3>
<p>If you do need a range of fixed floating point numbers generated then you can
do so by merging two decimal integer ranges together. For example</p>
<pre><code>» a [0..5].[0..9]
0.0
0.1
0.2
0.3
0.4
0.5
0.6
0.7
0.8
0.9
1.0
1.1
1.2
1.3
...
4.8
4.9
5.0
5.1
5.2
5.3
5.4
5.5
5.6
5.7
5.8
5.9
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
<li><RouterLink to="/mkarray/character.html">Character arrays</RouterLink>:
Making character arrays (a to z)</li>
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


