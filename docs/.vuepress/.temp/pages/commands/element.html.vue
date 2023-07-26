<template><div><h1 id="element" tabindex="-1"><a class="header-anchor" href="#element" aria-hidden="true">#</a> <code v-pre>[[</code> (element)</h1>
<blockquote>
<p>Outputs an element from a nested structure</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Outputs an element from an array, map or table. Unlike <strong>index</strong> (<code v-pre>[</code>),
<strong>element</strong> takes a path parameter which means it can work inside nested
structures without pipelining multiple commands together. However this
comes with the drawback that you can only return one element.</p>
<p><strong>Element</strong> (<code v-pre>[[</code>) also doesn't support the bang prefix (unlike) <strong>index</strong>.</p>
<p>Please note that indexes in Murex are counted from zero.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; [[ element ]] -&gt; `&lt;stdout&gt;`

$variable[[ element ]] -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Return the 2nd element in an array</p>
<pre><code>» ja [0..9] -&gt; [[ /1 ]]
[
    &quot;1&quot;,
]
</code></pre>
<p>Return the data-type and description of <strong>config shell syntax-highlighting</strong></p>
<pre><code>» config -&gt; [[ /shell/syntax-highlighting/Data-Type ]]
bool
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="element-counts-from-zero" tabindex="-1"><a class="header-anchor" href="#element-counts-from-zero" aria-hidden="true">#</a> Element counts from zero</h3>
<p>Indexes in Murex behave like any other computer array in that all arrays
start from zero (<code v-pre>0</code>).</p>
<h3 id="alternative-path-separators" tabindex="-1"><a class="header-anchor" href="#alternative-path-separators" aria-hidden="true">#</a> Alternative path separators</h3>
<p><strong>Element</strong> uses the first character in the path as the separator. So the
following are all valid parameters:</p>
<pre><code>» config -&gt; [[ ,shell,syntax-highlighting,Data-Type ]]
bool

» config -&gt; [[ &gt;shell&gt;syntax-highlighting&gt;Data-Type ]]
bool

» config -&gt; [[ \|shell\|syntax-highlighting\|Data-Type ]]
bool

» config -&gt; [[ &gt;shell&gt;syntax-highlighting&gt;Data-Type ]]
bool
</code></pre>
<p>However there are a few of caveats:</p>
<ol>
<li>
<p>Currently <strong>element</strong> does not support unicode separators. All separators
must be 1 byte characters. This limitation is highlighted as a bug, albeit
a low priority one. If this limitation does directly affect you then raise
an issue on GitHub to get the priority bumped up.</p>
</li>
<li>
<p>Any shell tokens (eg pipe <code v-pre>|</code>, <code v-pre>;</code>, <code v-pre>}</code>, etc) will need to be escaped. For
readability reasons it is recommended not to use such characters even
though it is technically possible to.</p>
<pre><code> # Would fail because the semi-colon is an unescaped / unquoted shell token
 config -&gt; [[ ;shell-syntax-highlighting;Data-Type ]]
</code></pre>
</li>
<li>
<p>Please also make sure you don't use a character that is also used inside
key names because keys <em>cannot</em> be escaped. For example both of the
following would fail:</p>
<pre><code> # Would fail because 'syntax-highlighting' and 'Data-Type' both also contain
 # the separator character
 config -&gt; [[ -shell-syntax-highlighting-Data-Type ]]

 # Would fail because you cannot escape key names (escaping happens at the
 # shell parser level rather than command parameter level)
 config -&gt; [[ -shell-syntax\-highlighting-Data\-Type ]]
</code></pre>
</li>
</ol>
<h3 id="quoting-parameters" tabindex="-1"><a class="header-anchor" href="#quoting-parameters" aria-hidden="true">#</a> Quoting parameters</h3>
<p>In Murex, everything is a function. Thus even <code v-pre>[[</code> is a function name and
the closing <code v-pre>]]</code> is actually a last parameter. This means the recommended way
to quote <strong>element</strong> parameters is to quote specific key names or the entire
path:</p>
<pre><code>» config -&gt; [[ /shell/&quot;syntax-highlighting&quot;/Data-Type ]]
bool

» config -&gt; [[ &quot;|shell|syntax-highlighting|Data-Type&quot; ]]
bool
</code></pre>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>[[</code></li>
<li><code v-pre>element</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
</ul>
</div></template>


