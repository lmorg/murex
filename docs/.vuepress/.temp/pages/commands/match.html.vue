<template><div><h1 id="match" tabindex="-1"><a class="header-anchor" href="#match" aria-hidden="true">#</a> <code v-pre>match</code></h1>
<blockquote>
<p>Match an exact value in an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>match</code> takes input from STDIN and returns any array items / lines which
contain an exact match of the parameters supplied.</p>
<p>When multiple parameters are supplied they are concatenated into the search
string and white space delimited. eg all three of the below are the same:</p>
<pre><code>match &quot;a b c&quot;
match a\sb\sc
match a b c
match a    b    c
</code></pre>
<p>If you want to return everything except the search string then use `!match</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Match every occurrence of search string</p>
<pre><code>`&lt;stdin&gt;` -&gt; match search string -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Match everything except search string</p>
<pre><code>`&lt;stdin&gt;` -&gt; !match search string -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Match <strong>Wed</strong></p>
<pre><code>» ja: [Monday..Friday] -&gt; match Wed
[
    &quot;Wednesday&quot;
]
</code></pre>
<p>Match everything except <strong>Wed</strong></p>
<pre><code>» ja: [Monday..Friday] -&gt; !match Wed
[
    &quot;Monday&quot;,
    &quot;Tuesday&quot;,
    &quot;Thursday&quot;,
    &quot;Friday&quot;
]
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>match</code> is data-type aware so will work against lists or arrays of whichever
Murex data-type is passed to it via STDIN and return the output in the
same data-type.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>match</code></li>
<li><code v-pre>!match</code></li>
<li><code v-pre>list.string</code></li>
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
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
<li><RouterLink to="/commands/suffix.html"><code v-pre>suffix</code></RouterLink>:
Prefix a string to every item in a list</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


