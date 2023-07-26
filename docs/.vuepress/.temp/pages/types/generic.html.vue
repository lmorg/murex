<template><div><h1 id="generic" tabindex="-1"><a class="header-anchor" href="#generic" aria-hidden="true">#</a> <code v-pre>*</code> (generic)</h1>
<blockquote>
<p>generic (primitive)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is the default data type used when STDOUT is returned from any external
executables.</p>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported. Tables columns are aligned</li>
<li><code v-pre>ReadArray()</code>
Treats each new line as a new array element</li>
<li><code v-pre>ReadArrayWithType()</code>
Treats each new line as a new array element, each element is <code v-pre>*</code></li>
<li><code v-pre>ReadIndex()</code>
Indexes treated as table coordinates</li>
<li><code v-pre>ReadMap()</code>
Works against tables such as the output from <code v-pre>ps -fe</code></li>
<li><code v-pre>ReadNotIndex()</code>
Indexes treated as table coordinates</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Writes a new line per array element - tabs are treated as columns</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/Marshal.html"><code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
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
<li><RouterLink to="/types/int.html"><code v-pre>int</code> </RouterLink>:
Whole number (primitive)</li>
<li><RouterLink to="/types/num.html"><code v-pre>num</code> (number)</RouterLink>:
Floating point number (primitive)</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/str.html"><code v-pre>str</code> (string) </RouterLink>:
string (primitive)</li>
</ul>
</div></template>


