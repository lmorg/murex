<template><div><h1 id="regexp" tabindex="-1"><a class="header-anchor" href="#regexp" aria-hidden="true">#</a> <code v-pre>regexp</code></h1>
<blockquote>
<p>Regexp tools for arrays / lists of strings</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>regexp</code> provides a few tools for text matching and manipulation against an
array or list of strings - thus <code v-pre>regexp</code> is Murex data-type aware.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; regexp expression -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<h3 id="find-elements" tabindex="-1"><a class="header-anchor" href="#find-elements" aria-hidden="true">#</a> Find elements</h3>
<pre><code>» ja: [monday..sunday] -&gt; regexp 'f/^([a-z]{3})day/'
[
    &quot;mon&quot;,
    &quot;fri&quot;,
    &quot;sun&quot;
]
</code></pre>
<p>This returns only 3 days because only 3 days match the expression (where
the days have to be 6 characters long) and then it only returns the first 3
characters because those are inside the parenthesis.</p>
<h3 id="match-elements" tabindex="-1"><a class="header-anchor" href="#match-elements" aria-hidden="true">#</a> Match elements</h3>
<p>Elements containing</p>
<pre><code>» ja: [monday..sunday] -&gt; regexp 'm/(mon|fri|sun)day/'
[
    &quot;monday&quot;,
    &quot;friday&quot;,
    &quot;sunday&quot;
]
</code></pre>
<p>Elements excluding</p>
<pre><code>» ja: [monday..sunday] -&gt; !regexp 'm/(mon|fri|sun)day/'
[
    &quot;tuesday&quot;,
    &quot;wednesday&quot;,
    &quot;thursday&quot;,
    &quot;saturday&quot;
]
</code></pre>
<h3 id="substitute-expression" tabindex="-1"><a class="header-anchor" href="#substitute-expression" aria-hidden="true">#</a> Substitute expression</h3>
<pre><code>» ja: [monday..sunday] -&gt; regexp 's/day/night/'
[
    &quot;monnight&quot;,
    &quot;tuesnight&quot;,
    &quot;wednesnight&quot;,
    &quot;thursnight&quot;,
    &quot;frinight&quot;,
    &quot;saturnight&quot;,
    &quot;sunnight&quot;
]
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>f</code>
output found expressions (doesn't support bang prefix)</li>
<li><code v-pre>m</code>
output elements that match expression (supports bang prefix)</li>
<li><code v-pre>s</code>
output all elements - substituting elements that match expression (doesn't support bang prefix)</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>regexp</code> is data-type aware so will work against lists or arrays of whichever
Murex data-type is passed to it via STDIN and return the output in the
same data-type.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>regexp</code></li>
<li><code v-pre>!regexp</code></li>
<li><code v-pre>list.regex</code></li>
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
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
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


