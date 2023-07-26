<template><div><h1 id="prepend" tabindex="-1"><a class="header-anchor" href="#prepend" aria-hidden="true">#</a> <code v-pre>prepend</code></h1>
<blockquote>
<p>Add data to the start of an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>prepend</code> a data to the start of an array.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; prepend: value -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» a: [January..December] -&gt; prepend: 'New Year'
New Year
January
February
March
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
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>prepend</code> and <code v-pre>append</code> are data type aware:</p>
<pre><code>» tout json [1,2,3] -&gt; append 4 5 6 bob
Error in `append` (1,22): cannot convert 'bob' to a floating point number: strconv.ParseFloat: parsing &quot;bob&quot;: invalid syntax
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>prepend</code></li>
<li><code v-pre>list.prepend</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/addheading.html"><code v-pre>addheading</code> </RouterLink>:
Adds headings to a table</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
<li><RouterLink to="/commands/msort.html"><code v-pre>msort</code> </RouterLink>:
Sorts an array - data type agnostic</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
</ul>
</div></template>


