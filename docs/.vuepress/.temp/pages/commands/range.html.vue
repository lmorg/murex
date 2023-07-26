<template><div><h1 id="range" tabindex="-1"><a class="header-anchor" href="#range" aria-hidden="true">#</a> <code v-pre>[</code> (range)</h1>
<blockquote>
<p>Outputs a ranged subset of data from STDIN</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This will read from STDIN and output a subset of data in a defined range.</p>
<p>The range can be defined as a number of different range types - such as the
content of the array or it's index / row number. You can also omit either
the start or the end of the search criteria to cover all items before or
after the remaining search criteria.</p>
<p><strong>Please note that <code v-pre>@[</code> syntax has been deprecated in favour of <code v-pre>[</code> syntax
instead</strong></p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; [start..end]flags -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Range over all months after March:</strong></p>
<pre><code>» a: [January..December] -&gt; [March..]se
April
May
June
July
August
September
October
November
December
</code></pre>
<p><strong>Range from the 6th to the 10th month:</strong></p>
<p>By default, ranges start from one, <code v-pre>1</code></p>
<pre><code>» a: [January..December] -&gt; [5..9]
May
June
July
August
September
</code></pre>
<p><strong>Return the first 3 months:</strong></p>
<p>This usage is similar to <code v-pre>head -n3</code></p>
<pre><code>» a: [January..December] -&gt; [..3]
October
November
December
</code></pre>
<p><strong>Return the last 3 months:</strong></p>
<p>This usage is similar to <code v-pre>tail -n3</code></p>
<pre><code>» a: [January..December] -&gt; [-3..]
October
November
December
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>8</code>
handles backspace characters (char 8) instead of treating it like a printable character</li>
<li><code v-pre>b</code>
removes blank (empty) lines from source</li>
<li><code v-pre>e</code>
exclude the start and end search criteria from the range</li>
<li><code v-pre>n</code>
numeric offset (indexed from 0)</li>
<li><code v-pre>r</code>
regexp match</li>
<li><code v-pre>s</code>
exact string match</li>
<li><code v-pre>t</code>
trims whitespace from source</li>
</ul>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>@[</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
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
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
</ul>
</div></template>


