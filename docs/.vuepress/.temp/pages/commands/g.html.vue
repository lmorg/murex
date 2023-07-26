<template><div><h1 id="g" tabindex="-1"><a class="header-anchor" href="#g" aria-hidden="true">#</a> <code v-pre>g</code></h1>
<blockquote>
<p>Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Returns a list of files and directories that match a glob pattern.</p>
<p>Output is a JSON list.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    g: pattern -> &lt;stdout>

    [ &lt;stdin> -> ] @g command pattern [ -> &lt;stdout> ]

    !g: pattern -> &lt;stdout>

    &lt;stdin> -> g: pattern -> &lt;stdout>

    &lt;stdin> -> !g: pattern -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Inline globbing:</p>
<pre><code>cat: @{ g: *.txt }
</code></pre>
<p>Writing a JSON array of files to disk:</p>
<pre><code>g: *.txt |&gt; filelist.json
</code></pre>
<p>Writing a list of files to disk:</p>
<pre><code>g: *.txt -&gt; format str |&gt; filelist.txt
</code></pre>
<p>Checking if a file exists:</p>
<pre><code>if { g: somefile.txt } then {
    # file exists
}
</code></pre>
<p>Checking if a file does not exist:</p>
<pre><code>!if { g: somefile.txt } then {
    # file does not exist
}
</code></pre>
<p>Return all files apart from text files:</p>
<pre><code>!g: *.txt
</code></pre>
<p>Filtering a file list based on glob matches:</p>
<pre><code>f: +f -&gt; g: *.md
</code></pre>
<p>Remove any glob matches from a file list:</p>
<pre><code>f: +f -&gt; !g: *.md
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="pattern-reference" tabindex="-1"><a class="header-anchor" href="#pattern-reference" aria-hidden="true">#</a> Pattern Reference</h3>
<ul>
<li><code v-pre>*</code> matches any number of (including zero) characters</li>
<li><code v-pre>?</code> matches any single character</li>
</ul>
<h3 id="inverse-matches" tabindex="-1"><a class="header-anchor" href="#inverse-matches" aria-hidden="true">#</a> Inverse Matches</h3>
<p>If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» g: READ*
[
    "README.md"
]

» !g: *
Error in `!g` (1,1): No data returned.
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="when-used-as-a-method" tabindex="-1"><a class="header-anchor" href="#when-used-as-a-method" aria-hidden="true">#</a> When Used As A Method</h3>
<p><code v-pre>!g</code> first looks for files that match its pattern, then it reads the file list
from STDIN. If STDIN contains contents that are not files then <code v-pre>!g</code> might not
handle those list items correctly. This shouldn't be an issue with <code v-pre>frx</code> in its
normal mode because it is only looking for matches however when used as <code v-pre>!g</code>
any items that are not files will leak through.</p>
<p>This is its designed feature and not a bug. If you wish to remove anything that
also isn't a file then you should first pipe into either <code v-pre>g: *</code>, <code v-pre>rx: .*</code>, or
<code v-pre>f +f</code> and then pipe that into <code v-pre>!g</code>.</p>
<p>The reason for this behavior is to separate this from <code v-pre>!regexp</code> and <code v-pre>!match</code>.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>g</code></li>
<li><code v-pre>!g</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/f.html"><code v-pre>f</code></RouterLink>:
Lists or filters file system objects (eg files)</li>
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
<li><RouterLink to="/commands/rx.html"><code v-pre>rx</code></RouterLink>:
Regexp pattern matching for file system objects (eg <code v-pre>.*\\.txt</code>)</li>
</ul>
</div></template>


