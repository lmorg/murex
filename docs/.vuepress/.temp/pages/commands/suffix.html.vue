<template><div><h1 id="suffix" tabindex="-1"><a class="header-anchor" href="#suffix" aria-hidden="true">#</a> <code v-pre>suffix</code></h1>
<blockquote>
<p>Prefix a string to every item in a list</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Takes a list from STDIN and returns that same list with each element suffixed.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; suffix str -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ja: [Monday..Wednesday] -> suffix foobar
[
    "Mondayfoobar",
    "Tuesdayfoobar",
    "Wednesdayfoobar"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Supported data types can queried via <code v-pre>runtime</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>runtime: --marshallers
runtime: --unmarshallers
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>suffix</code></li>
<li><code v-pre>list.suffix</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/apis/lang.MarshalData.html"><code v-pre>lang.MarshalData()</code> (system API)</RouterLink>:
Converts structured memory into a Murex data-type (eg for stdio)</li>
<li><RouterLink to="/apis/lang.UnmarshalData.html"><code v-pre>lang.UnmarshalData()</code> (system API)</RouterLink>:
Converts a Murex data-type into structured memory</li>
<li><RouterLink to="/commands/left.html"><code v-pre>left</code></RouterLink>:
Left substring every item in a list</li>
<li><RouterLink to="/commands/prefix.html"><code v-pre>prefix</code></RouterLink>:
Prefix a string to every item in a list</li>
<li><RouterLink to="/commands/right.html"><code v-pre>right</code></RouterLink>:
Right substring every item in a list</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


