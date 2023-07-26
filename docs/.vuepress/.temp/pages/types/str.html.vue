<template><div><h1 id="str-string" tabindex="-1"><a class="header-anchor" href="#str-string" aria-hidden="true">#</a> <code v-pre>str</code> (string)</h1>
<blockquote>
<p>string (primitive)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This type is modelled closely on generic but is more tailored for textual
(non-tabulated) data.</p>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>ReadArray()</code>
Treats each new line as a new array element</li>
<li><code v-pre>ReadArrayWithType()</code>
Treats each new line as a new array element, each array element is <code v-pre>str</code></li>
<li><code v-pre>ReadIndex()</code>
Indexes treated as a new line separated list</li>
<li><code v-pre>ReadMap()</code>
Treats each new line as a numbered map element</li>
<li><code v-pre>ReadNotIndex()</code>
Indexes treated as a new line separated list</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Writes a new line per array element</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/generic.html"><code v-pre>*</code> (generic) </RouterLink>:
generic (primitive)</li>
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
</ul>
</div></template>


