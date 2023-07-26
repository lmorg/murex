<template><div><h1 id="f" tabindex="-1"><a class="header-anchor" href="#f" aria-hidden="true">#</a> <code v-pre>f</code></h1>
<blockquote>
<p>Lists or filters file system objects (eg files)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>f</code> lists or filters lists of file system objects, such as files, directories,
etc. <code v-pre>f</code> provides a quick way to output file system items that confirm to
specific criteria.</p>
<p>You define this criteria by using +flags (eg <code v-pre>+s</code> for all symlinks) and
optionally then restricting that criteria with -flags (eg <code v-pre>-x</code> to remove all
executable items). All flags supported as <code v-pre>+</code> are also supported as a <code v-pre>-</code>.</p>
<p>By default <code v-pre>f</code> will return no results. You need to include +flags.</p>
<p>Output is a JSON array as this format preserves whitespace in file names.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>f: options -&gt; `&lt;stdout&gt;`

`&lt;stdin&gt;` -&gt; f: options -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Return only directories:</p>
<pre><code>f: +d
</code></pre>
<p>Return file and directories but exclude symlinks:</p>
<pre><code>f: +fd -s
</code></pre>
<p>Filter out files in a list (eg created by <code v-pre>g</code>) using conditions set by <code v-pre>f</code>:</p>
<pre><code>g: '*.go' -&gt; f: +f

rx: '\.(txt|md)' -&gt; f: +fw
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>+</code>
include files (pair this with any other flag apart from <code v-pre>-</code>)</li>
<li><code v-pre>-</code>
exclude files (pair this with any other flag apart from <code v-pre>+</code>)</li>
<li><code v-pre>?</code>
deprecated -- use <code v-pre>i</code> instead</li>
<li><code v-pre>D</code>
regular directories</li>
<li><code v-pre>E</code>
other read permissions</li>
<li><code v-pre>F</code>
regular files (exc symlinks, devices, sockets, named pipes, etc)</li>
<li><code v-pre>Q</code>
other write permissions</li>
<li><code v-pre>R</code>
user read permissions</li>
<li><code v-pre>S</code>
sockets</li>
<li><code v-pre>W</code>
user write permissions</li>
<li><code v-pre>X</code>
user execute permissions</li>
<li><code v-pre>Z</code>
other execute permissions</li>
<li><code v-pre>b</code>
block devices</li>
<li><code v-pre>c</code>
character devices</li>
<li><code v-pre>d</code>
all directories (inc symlinks)</li>
<li><code v-pre>e</code>
group read permissions</li>
<li><code v-pre>f</code>
all files (inc symlinks, devices, sockets, name pipes, etc)</li>
<li><code v-pre>i</code>
irregular files (nothing else is known about these files)</li>
<li><code v-pre>l</code>
symlinks</li>
<li><code v-pre>p</code>
POSIX named pipes (FIFO)</li>
<li><code v-pre>q</code>
group write permissions</li>
<li><code v-pre>r</code>
read permissions (user, group, or other)</li>
<li><code v-pre>s</code>
symlinks</li>
<li><code v-pre>w</code>
write permissions (user, group, or other)</li>
<li><code v-pre>x</code>
execute permissions (user, group, or other)</li>
<li><code v-pre>z</code>
group execute permissions</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>+</code> flags are always matched first. Then the <code v-pre>-</code> flags are used to filter out
any matches from the <code v-pre>+</code> flags.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/g.html"><code v-pre>g</code></RouterLink>:
Glob pattern matching for file system objects (eg <code v-pre>*.txt</code>)</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/commands/rx.html"><code v-pre>rx</code></RouterLink>:
Regexp pattern matching for file system objects (eg <code v-pre>.*\\.txt</code>)</li>
</ul>
</div></template>


