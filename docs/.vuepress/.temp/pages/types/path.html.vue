<template><div><h1 id="murex-shell-docs" tabindex="-1"><a class="header-anchor" href="#murex-shell-docs" aria-hidden="true">#</a> Murex Shell Docs</h1>
<h2 id="data-type-reference-path-string" tabindex="-1"><a class="header-anchor" href="#data-type-reference-path-string" aria-hidden="true">#</a> Data-Type Reference: <code v-pre>path</code> (string)</h2>
<blockquote>
<p>path data type</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This type is modelled closely on generic but is more tailored for textual
(non-tabulated) data.</p>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>ReadArray()</code>
Treats each new directory as a new array element</li>
<li><code v-pre>ReadArrayWithType()</code>
Treats each directory as a new array element, each array element is <code v-pre>str</code></li>
<li><code v-pre>ReadIndex()</code>
Indexes treated as a path separated list</li>
<li><code v-pre>ReadMap()</code>
Treats each new directory as a numbered map element</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Writes a new path, each array element as a directory</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/types/str.html"><code v-pre>str</code> (string) </RouterLink>:
string (primitive)</li>
</ul>
</div></template>


