<template><div><h1 id="jsplit" tabindex="-1"><a class="header-anchor" href="#jsplit" aria-hidden="true">#</a> <code v-pre>jsplit</code></h1>
<blockquote>
<p>Splits STDIN into a JSON array based on a regex parameter</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>jsplit</code> will read from STDIN and split it based on a regex parameter. It outputs a JSON array.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; jsplit: regex -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» (hello, world) -&gt; jsplit: l+
[
    &quot;he&quot;,
    &quot;o, wor&quot;,
    &quot;d&quot;
]
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>jsplit</code> will trim trailing carriage returns and line feeds from each element
as well as any trailing empty elements (zero length strings) in the JSON array.
However any empty elements will be retained and any other whitespace characters</p>
<ul>
<li>or carriage returns and/or line feeds in the middle of an element - will be
retained.</li>
</ul>
<p>This is so that the formatting of (multiline) text is retained as much as
possible to ensure the <code v-pre>jsplit</code> is accurate while at the same time any commonly
unwanted &quot;noise&quot; is stripped from the output.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>jsplit</code></li>
<li><code v-pre>list.split</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/2darray.html"><code v-pre>2darray</code> </RouterLink>:
Create a 2D JSON array from multiple input sources</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/map.html"><code v-pre>map</code> </RouterLink>:
Creates a map from two data sources</li>
<li><RouterLink to="/commands/msort.html"><code v-pre>msort</code> </RouterLink>:
Sorts an array - data type agnostic</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
</ul>
</div></template>


