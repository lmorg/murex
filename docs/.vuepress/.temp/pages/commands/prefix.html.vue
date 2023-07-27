<template><div><h1 id="prefix" tabindex="-1"><a class="header-anchor" href="#prefix" aria-hidden="true">#</a> <code v-pre>prefix</code></h1>
<blockquote>
<p>Prefix a string to every item in a list</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Takes a list from STDIN and returns that same list with each element prefixed.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; prefix str -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» ja: [Monday..Wednesday] -&gt; prefix foobar
[
    &quot;foobarMonday&quot;,
    &quot;foobarTuesday&quot;,
    &quot;foobarWednesday&quot;
]
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Supported data types can queried via <code v-pre>runtime</code></p>
<pre><code>runtime: --marshallers
runtime: --unmarshallers
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>prefix</code></li>
<li><code v-pre>list.prefix</code></li>
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
<li><RouterLink to="/commands/right.html"><code v-pre>right</code></RouterLink>:
Right substring every item in a list</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/suffix.html"><code v-pre>suffix</code></RouterLink>:
Prefix a string to every item in a list</li>
</ul>
</div></template>


