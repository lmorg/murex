<template><div><h1 id="msort" tabindex="-1"><a class="header-anchor" href="#msort" aria-hidden="true">#</a> <code v-pre>msort</code></h1>
<blockquote>
<p>Sorts an array - data type agnostic</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This builtin takes input from stdin, sorts it and the outputs it to stdout.</p>
<p>The code behind <code v-pre>msort</code> is significantly more lightweight than UNIX sort.
It doesn't work with numeric types (eg sorting floating point numbers),
reversed order nor multi-column data. It is specifically designed to work
with lists of data. For example arrays in data formats like JSON (<code v-pre>json</code>),
YAML (<code v-pre>yaml</code>) or S-Expressions (<code v-pre>sexp</code>); or lists of strings (<code v-pre>str</code>). The
intention is to cover use cases not already covered by UNIX sort while also
providing something rudimentary for Murex scripts to function on Windows
without having to write lots of ugly platform-specific code. This is also
the reason this builtin is called <code v-pre>msort</code> rather than conflicting with the
existing UNIX name, <code v-pre>sort</code>.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; msort -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» tout: json ([&quot;c&quot;, &quot;b&quot;, &quot;a&quot;]) -&gt; msort
[
    &quot;a&quot;,
    &quot;b&quot;,
    &quot;c&quot;
]
</code></pre>
<p>Since <code v-pre>msort</code> does not support reversed order, you will need to pipe the
output of <code v-pre>msort</code> into another builtin:</p>
<pre><code>» tout: json ([&quot;c&quot;, &quot;b&quot;, &quot;a&quot;]) -&gt; msort -&gt; mtac
[
    &quot;c&quot;,
    &quot;b&quot;,
    &quot;a&quot;
]
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>msort</code></li>
<li><code v-pre>list.sort</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/alter.html"><code v-pre>alter</code></RouterLink>:
Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/jsplit.html"><code v-pre>jsplit</code> </RouterLink>:
Splits STDIN into a JSON array based on a regex parameter</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
</ul>
</div></template>


