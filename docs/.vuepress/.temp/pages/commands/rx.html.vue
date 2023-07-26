<template><div><h1 id="rx" tabindex="-1"><a class="header-anchor" href="#rx" aria-hidden="true">#</a> <code v-pre>rx</code></h1>
<blockquote>
<p>Regexp pattern matching for file system objects (eg <code v-pre>.*\\.txt</code>)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Returns a list of files and directories that match a regexp pattern.</p>
<p>Output is a JSON list.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>rx: pattern -&gt; `&lt;stdout&gt;`

!rx: pattern -&gt; `&lt;stdout&gt;`

`&lt;stdin&gt;` -&gt; rx: pattern -&gt; `&lt;stdout&gt;`

`&lt;stdin&gt;` -&gt; !rx: pattern -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Inline regex file matching:</p>
<pre><code>cat: @{ rx: '.*\.txt' }
</code></pre>
<p>Writing a list of files to disk:</p>
<pre><code>rx: '.*\.go' |&gt; filelist.txt
</code></pre>
<p>Checking if files exist:</p>
<pre><code>if { rx: somefiles.* } then {
    # files exist
}
</code></pre>
<p>Checking if files do not exist:</p>
<pre><code>!if { rx: somefiles.* } then {
    # files do not exist
}
</code></pre>
<p>Return all files apart from text files:</p>
<pre><code>!g: '\.txt$'
</code></pre>
<p>Filtering a file list based on regexp matches file:</p>
<pre><code>f: +f -&gt; rx: '.*\.txt'
</code></pre>
<p>Remove any regexp file matches from a file list:</p>
<pre><code>f: +f -&gt; !rx: '.*\.txt'
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="traversing-directories" tabindex="-1"><a class="header-anchor" href="#traversing-directories" aria-hidden="true">#</a> Traversing Directories</h3>
<p>Unlike globbing (<code v-pre>g</code>) which can traverse directories (eg <code v-pre>g: /path/*</code>), <code v-pre>rx</code> is
only designed to match file system objects in the current working directory.</p>
<p><code v-pre>rx</code> uses Go (lang)'s standard regexp engine.</p>
<h3 id="inverse-matches" tabindex="-1"><a class="header-anchor" href="#inverse-matches" aria-hidden="true">#</a> Inverse Matches</h3>
<p>If you want to exclude any matches based on wildcards, rather than include
them, then you can use the bang prefix. eg</p>
<pre><code>» rx: READ*
[
    &quot;README.md&quot;
]

murex-dev» !rx: .*
Error in `!rx` (1,1): No data returned.
</code></pre>
<h3 id="when-used-as-a-method" tabindex="-1"><a class="header-anchor" href="#when-used-as-a-method" aria-hidden="true">#</a> When Used As A Method</h3>
<p><code v-pre>!rx</code> first looks for files that match its pattern, then it reads the file list
from STDIN. If STDIN contains contents that are not files then <code v-pre>!rx</code> might not
handle those list items correctly. This shouldn't be an issue with <code v-pre>rx</code> in its
normal mode because it is only looking for matches however when used as <code v-pre>!rx</code>
any items that are not files will leak through.</p>
<p>This is its designed feature and not a bug. If you wish to remove anything that
also isn't a file then you should first pipe into either <code v-pre>g: *</code>, <code v-pre>rx: .*</code>, or
<code v-pre>f +f</code> and then pipe that into <code v-pre>!rx</code>.</p>
<p>The reason for this behavior is to separate this from <code v-pre>!regexp</code> and <code v-pre>!match</code>.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>rx</code></li>
<li><code v-pre>!rx</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/f.html"><code v-pre>f</code></RouterLink>:
Lists or filters file system objects (eg files)</li>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/commands/match.html"><code v-pre>match</code></RouterLink>:
Match an exact value in an array</li>
<li><RouterLink to="/commands/regexp.html"><code v-pre>regexp</code></RouterLink>:
Regexp tools for arrays / lists of strings</li>
</ul>
</div></template>


