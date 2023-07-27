<template><div><h1 id="int" tabindex="-1"><a class="header-anchor" href="#int" aria-hidden="true">#</a> <code v-pre>int</code></h1>
<blockquote>
<p>Whole number (primitive)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>An integer is a whole number (eg 1, 2, 3, 4) rather than one with a decimal
point (such as 1.1).</p>
<p>Integers in Murex are sized based on the bit (or word) size of the target
CPU.</p>
<p>A 386, ARMv6 or other 32bit build of Murex would see the range of from
<code v-pre>-2147483648</code> (negative) through <code v-pre>2147483647</code> (positive).</p>
<p>AMD64 or other 64bit built of Murex would see the range from
<code v-pre>-9223372036854775808</code> (negative) through <code v-pre>9223372036854775807</code> (positive).</p>
<blockquote>
<p>Unless you specifically know you only want whole numbers, it is recommended
that you use the default numeric data-type: <code v-pre>num</code>.</p>
</blockquote>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
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


