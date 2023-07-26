<template><div><h1 id="mtac" tabindex="-1"><a class="header-anchor" href="#mtac" aria-hidden="true">#</a> <code v-pre>mtac</code></h1>
<blockquote>
<p>Reverse the order of an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>mtac</code> takes input from STDIN and reverses the order of it.</p>
<p>It's name is derived from a program called <code v-pre>tac</code> - a tool that functions
like <code v-pre>cat</code> but returns the contents in the reverse order. The difference
with the <code v-pre>mtac</code> builtin is that it is data-type aware. So it doesn't just
function as a replacement for <code v-pre>tac</code> but it also works on JSON arrays,
s-expressions, and any other data-type supporting arrays compiled into
Murex.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; mtac -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» ja: [Monday..Friday] -&gt; mtac
[
    &quot;Friday&quot;,
    &quot;Thursday&quot;,
    &quot;Wednesday&quot;,
    &quot;Tuesday&quot;,
    &quot;Monday&quot;
]

# Normal output (without mtac)
» ja: [Monday..Friday]
[
    &quot;Monday&quot;,
    &quot;Tuesday&quot;,
    &quot;Wednesday&quot;,
    &quot;Thursday&quot;,
    &quot;Friday&quot;
]
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Please bare in mind that while Murex is optimised with concurrency and
streaming in mind, it's impossible to reverse an incomplete array. Thus all
all of STDIN must have been read and that file closed before <code v-pre>mtac</code> can
output.</p>
<p>In practical terms you shouldn't notice any difference except for when
STDIN is a long running process or non-standard stream (eg network pipe).</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>mtac</code></li>
<li><code v-pre>list.reverse</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/2darray.html"><code v-pre>2darray</code> </RouterLink>:
Create a 2D JSON array from multiple input sources</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/jsplit.html"><code v-pre>jsplit</code> </RouterLink>:
Splits STDIN into a JSON array based on a regex parameter</li>
<li><RouterLink to="/commands/map.html"><code v-pre>map</code> </RouterLink>:
Creates a map from two data sources</li>
<li><RouterLink to="/commands/msort.html"><code v-pre>msort</code> </RouterLink>:
Sorts an array - data type agnostic</li>
<li><RouterLink to="/commands/prefix.html"><code v-pre>prefix</code></RouterLink>:
Prefix a string to every item in a list</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/suffix.html"><code v-pre>suffix</code></RouterLink>:
Prefix a string to every item in a list</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


