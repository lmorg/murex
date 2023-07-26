<template><div><h1 id="commonlog" tabindex="-1"><a class="header-anchor" href="#commonlog" aria-hidden="true">#</a> <code v-pre>commonlog</code></h1>
<blockquote>
<p>Apache httpd &quot;common&quot; log format</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Apache httpd supports a few different log formats. This Murex type is for
parsing the &quot;common&quot; log format.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>The code here is very rudimentary. If you have large log files or need more complex
data querying then this data-type is probably not the right tool. Maybe try one of
the following:</p>
<ul>
<li><a href="https://github.com/lmorg/firesword" target="_blank" rel="noopener noreferrer">Firesword<ExternalLinkIcon/></a> - for command line analysis</li>
<li><a href="https://github.com/lmorg/plasmasword" target="_blank" rel="noopener noreferrer">Plasmasword<ExternalLinkIcon/></a> - exports fields to an sqlite3 or mysql database</li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported though no unmarshalling is currently supported</li>
<li><code v-pre>ReadArray()</code>
Supported. Each line is considered an index (like with <code v-pre>str</code> data type)</li>
<li><code v-pre>ReadArrayWithType()</code>
Supported. Each line is considered an index with <code v-pre>commonlog</code> data type</li>
<li><code v-pre>ReadIndex()</code>
Entire log file is read and then the indexes are derived from there</li>
<li><code v-pre>ReadMap()</code>
Not supported, currently a work in progress</li>
<li><code v-pre>ReadNotIndex()</code>
Entire log file is read and then the indexes are derived from there</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
<li><RouterLink to="/apis/Marshal.html"><code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/ReadArray.html"><code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadIndex.html"><code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/ReadMap.html"><code v-pre>ReadMap()</code> (type)</RouterLink>:
Treat data type as a key/value structure and read its contents</li>
<li><RouterLink to="/apis/ReadNotIndex.html"><code v-pre>ReadNotIndex()</code> (type)</RouterLink>:
Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</li>
<li><RouterLink to="/apis/Unmarshal.html"><code v-pre>Unmarshal()</code> (type)</RouterLink>:
Converts a structured file format into structured memory</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/str.html"><code v-pre>str</code> (string) </RouterLink>:
string (primitive)</li>
</ul>
</div></template>


