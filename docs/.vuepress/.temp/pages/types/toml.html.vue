<template><div><h1 id="toml" tabindex="-1"><a class="header-anchor" href="#toml" aria-hidden="true">#</a> <code v-pre>toml</code></h1>
<blockquote>
<p>Tom's Obvious, Minimal Language (TOML)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>TOML support within Murex is pretty mature however it is not considered a
primitive. Which means, while it is a recommended builtin which you should
expect in most deployments of Murex, it's still an optional package and
thus may not be present in some edge cases. This is because it relies on
external source packages for the shell to compile.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example TOML document taken from <a href="https://en.wikipedia.org/wiki/TOML" target="_blank" rel="noopener noreferrer">Wikipedia<ExternalLinkIcon/></a></p>
<pre><code># This is a TOML document.

title = &quot;TOML Example&quot;

[owner]
name = &quot;Tom Preston-Werner&quot;
dob = 1979-05-27T07:32:00-08:00 # First class dates

[database]
server = &quot;192.168.1.1&quot;
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # Indentation (tabs and/or spaces) is allowed but not required
  [servers.alpha]
  ip = &quot;10.0.0.1&quot;
  dc = &quot;eqdc10&quot;

  [servers.beta]
  ip = &quot;10.0.0.2&quot;
  dc = &quot;eqdc10&quot;

[clients]
data = [ [&quot;gamma&quot;, &quot;delta&quot;], [1, 2] ]

# Line breaks are OK when inside arrays
hosts = [
  &quot;alpha&quot;,
  &quot;omega&quot;
]
</code></pre>
<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2>
<ul>
<li><strong>Extension</strong>: <code v-pre>toml</code></li>
<li><strong>MIME</strong>: <code v-pre>application/toml</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-toml</code></li>
<li><strong>MIME</strong>: <code v-pre>text/toml</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-toml</code></li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>ReadArray()</code>
Hook supported albeit TOML doesn't support naked arrays</li>
<li><code v-pre>ReadArrayWithType()</code>
Hook supported albeit TOML doesn't support naked arrays</li>
<li><code v-pre>ReadIndex()</code>
Works against all properties in TOML</li>
<li><code v-pre>ReadNotIndex()</code>
Works against all properties in TOML</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
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
<li><RouterLink to="/apis/WriteArray.html"><code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/yaml.html"><code v-pre>yaml</code> </RouterLink>:
YAML Ain't Markup Language (YAML)</li>
</ul>
</div></template>


