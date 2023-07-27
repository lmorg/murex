<template><div><h1 id="data-type-reference" tabindex="-1"><a class="header-anchor" href="#data-type-reference" aria-hidden="true">#</a> Data-Type Reference</h1>
<p>This section is a glossary of data-types which Murex is natively aware.</p>
<p>Most of the time you will not need to worry about typing in Murex as the
shell is designed around productivity as opposed to strictness despite
generally following a strictly typed design.</p>
<p>Read the <a href="/tour">Language Tour</a> for more detail on this topic.</p>
<h2 id="definitions" tabindex="-1"><a class="header-anchor" href="#definitions" aria-hidden="true">#</a> Definitions</h2>
<p>For clarity, it is worth explaining a couple of terms:</p>
<ol>
<li>
<p>&quot;Data-types&quot; in Murex are a description of the format of data. This
means that while any stdio stream in UNIX will by &quot;bytes&quot;, Murex might
label that data as being a JSON string or CSV file (for example) which
means any builtins that parse that stdio stream, for example to return
the first 8 items, would need to parse those types differently. Thus a
&quot;data-type&quot; in Murex is actually more than just a description of a data
structure; it is a series of APIs to marshall and unmarshall data from
complex file formats. This enables you to use the same command line tools
to query any type of output.</p>
</li>
<li>
<p>&quot;Primitive&quot; data-types refer to types that are the required by Murex
to function. These will be <code v-pre>int</code>, <code v-pre>float</code> / <code v-pre>number</code>, <code v-pre>bool</code>, <code v-pre>string</code>,
<code v-pre>generic</code>, and <code v-pre>null</code>.</p>
</li>
</ol>
<h2 id="feature-sets" tabindex="-1"><a class="header-anchor" href="#feature-sets" aria-hidden="true">#</a> Feature Sets</h2>
<p>Since not all data formats are equal (for example the TOML file format
doesn't support naked arrays where as JSON does), you may find some
features missing in some data-types which are present in others. If in
doubt then refer to the manual here or check the API manual for more
details about specific hooks.</p>
</div></template>


