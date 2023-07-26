<template><div><h1 id="a-mkarray" tabindex="-1"><a class="header-anchor" href="#a-mkarray" aria-hidden="true">#</a> <code v-pre>a</code> (mkarray)</h1>
<blockquote>
<p>A sophisticated yet simple way to build an array or list</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Pronounced &quot;make array&quot;, like <code v-pre>mkdir</code> (etc), Murex has a pretty sophisticated
builtin for generating arrays. Think like bash's <code v-pre>{1..9}</code> syntax:</p>
<pre><code>a: [1..9]
</code></pre>
<p>Except Murex also supports other sets of ranges like dates, days of the week,
and alternative number bases.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>a: [start..end] -> &lt;stdout>
a: [start..end,start..end] -> &lt;stdout>
a: [start..end][start..end] -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>All usages also work with <code v-pre>ja</code> and <code v-pre>ta</code> as well, eg:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>ja: [start..end] -> &lt;stdout>
ta: data-type [start..end] -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>You can also inline arrays with the <code v-pre>%[]</code> syntax, eg:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[start..end]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: [1..3]
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
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="advanced-array-syntax" tabindex="-1"><a class="header-anchor" href="#advanced-array-syntax" aria-hidden="true">#</a> Advanced Array Syntax</h3>
<p>The syntax for <code v-pre>a</code> is a comma separated list of parameters with expansions
stored in square brackets. You can have an expansion embedded inside a
parameter or as it's own parameter. Expansions can also have multiple
parameters.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: 01,02,03,05,06,07
01
02
03
05
06
07

» a: 0[1..3],0[5..7]
01
02
03
05
06
07

» a: 0[1..3,5..7]
01
02
03
05
06
07

» a: b[o,i]b
bob
bib
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>You can also have multiple expansion blocks in a single parameter:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: a[1..3]b[5..7]
a1b5
a1b6
a1b7
a2b5
a2b6
a2b7
a3b5
a3b6
a3b7
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>a</code> will cycle through each iteration of the last expansion, moving itself
backwards through the string; behaving like an normal counter.</p>
<h3 id="creating-json-arrays-with-ja" tabindex="-1"><a class="header-anchor" href="#creating-json-arrays-with-ja" aria-hidden="true">#</a> Creating JSON arrays with <code v-pre>ja</code></h3>
<p>As you can see from the previous examples, <code v-pre>a</code> returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: <code v-pre>a: [0..254].[0..254].[0..254].[0..254]</code>
(this kind of array expansion would hang bash).</p>
<p>However if you needed a JSON string then you can use all the same syntax
as <code v-pre>a</code> but forgo the streaming capability:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ja: [Monday..Sunday]
[
    "Monday",
    "Tuesday",
    "Wednesday",
    "Thursday",
    "Friday",
    "Saturday",
    "Sunday"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>This is particularly useful if you are adding formatting that might break
under <code v-pre>a</code>'s formatting (which uses the <code v-pre>str</code> data type).</p>
<h3 id="smart-arrays" tabindex="-1"><a class="header-anchor" href="#smart-arrays" aria-hidden="true">#</a> Smart arrays</h3>
<p>Murex supports a number of different formats that can be used to generate
arrays. For more details on these please refer to the documents for each format</p>
<ul>
<li><RouterLink to="/mkarray/date.html">Calendar Date Ranges</RouterLink>:
Create arrays of dates</li>
<li><RouterLink to="/mkarray/character.html">Character arrays</RouterLink>:
Making character arrays (a to z)</li>
<li><RouterLink to="/mkarray/decimal.html">Decimal Ranges</RouterLink>:
Create arrays of decimal integers</li>
<li><RouterLink to="/mkarray/non-decimal.html">Non-Decimal Ranges</RouterLink>:
Create arrays of integers from non-decimal number bases</li>
<li><RouterLink to="/mkarray/special.html">Special Ranges</RouterLink>:
Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/types/str.html"><code v-pre>str</code> (string) </RouterLink>:
string (primitive)</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


