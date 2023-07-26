<template><div><h1 id="count" tabindex="-1"><a class="header-anchor" href="#count" aria-hidden="true">#</a> <code v-pre>count</code></h1>
<blockquote>
<p>Count items in a map, list or array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; count: [ --duplications | --unique | --total ] -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Count number of items in a map, list or array:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» tout: json (["a", "b", "c"]) -> count
3
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--duplications</code>
Output a JSON map of items and the number of their occurrences in a list or array</li>
<li><code v-pre>--total</code>
Read an array, list or map from STDIN and output the length for that array (default behaviour)</li>
<li><code v-pre>--unique</code>
Print the number of unique elements in a list or array</li>
<li><code v-pre>-d</code>
Alias for `--duplications</li>
<li><code v-pre>-t</code>
Alias for `--total</li>
<li><code v-pre>-u</code>
Alias for `--unique</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="modes" tabindex="-1"><a class="header-anchor" href="#modes" aria-hidden="true">#</a> Modes</h3>
<p>If no flags are set, <code v-pre>count</code> will default to using <code v-pre>--total</code>.</p>
<h4 id="total-total-t" tabindex="-1"><a class="header-anchor" href="#total-total-t" aria-hidden="true">#</a> Total: <code v-pre>--total</code> / <code v-pre>-t</code></h4>
<p>This will read an array, list or map from STDIN and output the length for
that array.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a [25-Dec-2020..05-Jan-2021] -> count
12
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>This also replaces the older <code v-pre>len</code> method.</p>
</blockquote>
<p>Please note that this returns the length of the <em>array</em> rather than string.
For example <code v-pre>out &quot;foobar&quot; -&gt; count</code> would return <code v-pre>1</code> because an array in the
<code v-pre>str</code> data type would be new line separated (eg <code v-pre>out &quot;foo\nbar&quot; -&gt; count</code>
would return <code v-pre>2</code>). If you need to count characters in a string and are
running POSIX (eg Linux / BSD / OSX) then it is recommended to use <code v-pre>wc</code>
instead. But be mindful that <code v-pre>wc</code> will also count new line characters.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: "foobar" -> count
1

» out: "foo\nbar" -> count
2

» out: "foobar" -> wc: -c
7

» out: "foo\nbar" -> wc: -c
8

» printf: "foobar" -> wc: -c
6
# (printf does not print a trailing new line)
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="duplications-duplications-d" tabindex="-1"><a class="header-anchor" href="#duplications-duplications-d" aria-hidden="true">#</a> Duplications: <code v-pre>--duplications</code> / <code v-pre>-d</code></h4>
<p>This returns a JSON map of items and the number of their occurrences in a list
or array.</p>
<p>For example in the quote below, only the word &quot;the&quot; is repeated so that entry
will have a value of <code v-pre>2</code> while ever other entry has a value of <code v-pre>1</code> because they
only appear once in the quote.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: "the quick brown fox jumped over the lazy dog" -> jsplit: \s -> count: --duplications
{
    "brown": 1,
    "dog": 1,
    "fox": 1,
    "jumped": 1,
    "lazy": 1,
    "over": 1,
    "quick": 1,
    "the": 2
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="unique-unique-u" tabindex="-1"><a class="header-anchor" href="#unique-unique-u" aria-hidden="true">#</a> Unique: <code v-pre>--unique</code> / <code v-pre>-u</code></h4>
<p>Returns the number of unique elements in a list or array.</p>
<p>For example in the quote below, only the word &quot;the&quot; is repeated, thus the
unique count should be one less than the total count:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out "the quick brown fox jumped over the lazy dog" -> jsplit \s -> count --unique
8
» out "the quick brown fox jumped over the lazy dog" -> jsplit \s -> count --total
9
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>count</code></li>
<li><code v-pre>len</code></li>
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
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/jsplit.html"><code v-pre>jsplit</code> </RouterLink>:
Splits STDIN into a JSON array based on a regex parameter</li>
<li><RouterLink to="/commands/jsplit.html"><code v-pre>jsplit</code> </RouterLink>:
Splits STDIN into a JSON array based on a regex parameter</li>
<li><RouterLink to="/commands/map.html"><code v-pre>map</code> </RouterLink>:
Creates a map from two data sources</li>
<li><RouterLink to="/commands/msort.html"><code v-pre>msort</code> </RouterLink>:
Sorts an array - data type agnostic</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


