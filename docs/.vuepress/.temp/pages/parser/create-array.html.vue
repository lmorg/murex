<template><div><h1 id="create-array-constructor" tabindex="-1"><a class="header-anchor" href="#create-array-constructor" aria-hidden="true">#</a> Create array (<code v-pre>%[]</code>) constructor</h1>
<blockquote>
<p>Quickly generate arrays</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>%[]</code> is a way of defining arrays in expressions and statements. Whenever a
<code v-pre>%[]</code> array is outputted as a string, it will be converted to minified JSON.</p>
<p>Array elements inside <code v-pre>%[]</code> can be whitespace and/or comma delimited. This
allows for compatibility with both Bash muscle memory, and people more
familiar with JSON.</p>
<p>Additionally you can also embed <code v-pre>a</code> style parameters inside <code v-pre>%[]</code> arrays too.</p>
<p>Like with YAML, strings in <code v-pre>%[]</code> do not need to be quoted unless you need to
force numeric or boolean looking values to be stored as strings.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<h3 id="arrays-passed-as-a-json-string" tabindex="-1"><a class="header-anchor" href="#arrays-passed-as-a-json-string" aria-hidden="true">#</a> Arrays passed as a JSON string:</h3>
<pre><code>» echo %[1..3]
[1,2,3]

» %[1..3] -&gt; cat
[1,2,3]
</code></pre>
<h3 id="different-supported-syntax-for-creating-a-numeric-array" tabindex="-1"><a class="header-anchor" href="#different-supported-syntax-for-creating-a-numeric-array" aria-hidden="true">#</a> Different supported syntax for creating a numeric array:</h3>
<h4 id="as-a-range" tabindex="-1"><a class="header-anchor" href="#as-a-range" aria-hidden="true">#</a> As a range</h4>
<pre><code>» %[1..3]
[
    1,
    2,
    3
]
</code></pre>
<h4 id="json-formatted" tabindex="-1"><a class="header-anchor" href="#json-formatted" aria-hidden="true">#</a> JSON formatted</h4>
<pre><code>» %[1,2,3]
[
    1,
    2,
    3
]
</code></pre>
<h4 id="whitespace-separated" tabindex="-1"><a class="header-anchor" href="#whitespace-separated" aria-hidden="true">#</a> Whitespace separated</h4>
<pre><code>» %[1 2 3]
[
    1,
    2,
    3
]
</code></pre>
<h4 id="values-and-ranges" tabindex="-1"><a class="header-anchor" href="#values-and-ranges" aria-hidden="true">#</a> Values and ranges</h4>
<pre><code>» %[1,2..3]
[
    1,
    2,
    3
]
</code></pre>
<h3 id="strings" tabindex="-1"><a class="header-anchor" href="#strings" aria-hidden="true">#</a> Strings:</h3>
<h4 id="barewords-and-whitespace-separated" tabindex="-1"><a class="header-anchor" href="#barewords-and-whitespace-separated" aria-hidden="true">#</a> barewords and whitespace separated</h4>
<p>This will allow you to copy/paste lists from traditional shells like Bash</p>
<pre><code>» %[foo bar]
[
    &quot;foo&quot;,
    &quot;bar&quot;
]
</code></pre>
<h4 id="json-formatted-1" tabindex="-1"><a class="header-anchor" href="#json-formatted-1" aria-hidden="true">#</a> JSON formatted</h4>
<pre><code>» %[&quot;foo&quot;, &quot;bar&quot;]
[
    &quot;foo&quot;,
    &quot;bar&quot;
]
</code></pre>
<h3 id="special-ranges" tabindex="-1"><a class="header-anchor" href="#special-ranges" aria-hidden="true">#</a> Special ranges</h3>
<pre><code>» %[June..August]
[
    &quot;June&quot;,
    &quot;July&quot;,
    &quot;August&quot;
]
</code></pre>
<p>A full list of special ranges are available at <RouterLink to="/mkarray/special.html">docs/mkarray/special</RouterLink></p>
<h3 id="multiple-expansion-blocks" tabindex="-1"><a class="header-anchor" href="#multiple-expansion-blocks" aria-hidden="true">#</a> Multiple expansion blocks:</h3>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» %[[A,B]:[1..4]]
[
    "A:1",
    "A:2",
    "A:3",
    "A:4",
    "B:1",
    "B:2",
    "B:3",
    "B:4"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="nested-arrays" tabindex="-1"><a class="header-anchor" href="#nested-arrays" aria-hidden="true">#</a> Nested arrays:</h3>
<pre><code>» %[foo [bar]]
[
    &quot;foo&quot;,
    [
        &quot;bar&quot;
    ]
]
</code></pre>
<p>The <code v-pre>%</code> prefix for the nested array is optional.</p>
<h3 id="json-objects-within-arrays" tabindex="-1"><a class="header-anchor" href="#json-objects-within-arrays" aria-hidden="true">#</a> JSON objects within arrays</h3>
<pre><code>» %[foo {bar: baz}]
[
    &quot;foo&quot;,
    {
        &quot;bar&quot;: &quot;baz&quot;
    }
]
</code></pre>
<p>The <code v-pre>%</code> prefix for the nested object is optional.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
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
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/create-object.html">Create object (<code v-pre>%{}</code>) constructor</RouterLink>:
Quickly generate objects and maps</li>
<li><RouterLink to="/parser/double-quote.html">Double Quote (<code v-pre>&quot;</code>) Token</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/parser/single-quote.html">Single Quote (<code v-pre>'</code>) Token</RouterLink>:
Initiates or terminates a string (variables not expanded)</li>
<li><RouterLink to="/mkarray/special.html">Special Ranges</RouterLink>:
Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


