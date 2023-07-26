<template><div><h1 id="jsonl" tabindex="-1"><a class="header-anchor" href="#jsonl" aria-hidden="true">#</a> <code v-pre>jsonl</code></h1>
<blockquote>
<p>JSON Lines</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>The following description is taken from <a href="http://jsonlines.org/" target="_blank" rel="noopener noreferrer">jsonlines.org<ExternalLinkIcon/></a>:</p>
<blockquote>
<p>JSON Lines is a convenient format for storing structured data that may be
processed one record at a time. It works well with unix-style text
processing tools and shell pipelines. It's a great format for log files.
It's also a flexible format for passing messages between cooperating
processes.</p>
</blockquote>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example JSON lines documents taken from <a href="http://jsonlines.org/examples/" target="_blank" rel="noopener noreferrer">jsonlines.org<ExternalLinkIcon/></a></p>
<h3 id="tabulated-data" tabindex="-1"><a class="header-anchor" href="#tabulated-data" aria-hidden="true">#</a> Tabulated data</h3>
<pre><code>[&quot;Name&quot;, &quot;Session&quot;, &quot;Score&quot;, &quot;Completed&quot;]
[&quot;Gilbert&quot;, &quot;2013&quot;, 24, true]
[&quot;Alexa&quot;, &quot;2013&quot;, 29, true]
[&quot;May&quot;, &quot;2012B&quot;, 14, false]
[&quot;Deloise&quot;, &quot;2012A&quot;, 19, true]
</code></pre>
<p>This format is equatable to <code v-pre>generic</code> and <code v-pre>csv</code>.</p>
<h3 id="nested-objects" tabindex="-1"><a class="header-anchor" href="#nested-objects" aria-hidden="true">#</a> Nested objects</h3>
<pre><code>{&quot;name&quot;: &quot;Gilbert&quot;, &quot;wins&quot;: [[&quot;straight&quot;, &quot;7♣&quot;], [&quot;one pair&quot;, &quot;10♥&quot;]]}
{&quot;name&quot;: &quot;Alexa&quot;, &quot;wins&quot;: [[&quot;two pair&quot;, &quot;4♠&quot;], [&quot;two pair&quot;, &quot;9♠&quot;]]}
{&quot;name&quot;: &quot;May&quot;, &quot;wins&quot;: []}
{&quot;name&quot;: &quot;Deloise&quot;, &quot;wins&quot;: [[&quot;three of a kind&quot;, &quot;5♣&quot;]]}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="concatenated-json" tabindex="-1"><a class="header-anchor" href="#concatenated-json" aria-hidden="true">#</a> Concatenated JSON</h3>
<p>Technically the <code v-pre>jsonl</code> Unmarshal() method supports Concatenated JSON, as
described on <a href="https://en.wikipedia.org/wiki/JSON_streaming#Concatenated_JSON" target="_blank" rel="noopener noreferrer">Wikipedia<ExternalLinkIcon/></a>:</p>
<blockquote>
<p>Concatenated JSON streaming allows the sender to simply write each JSON
object into the stream with no delimiters. It relies on the receiver using
a parser that can recognize and emit each JSON object as the terminating
character is parsed. Concatenated JSON isn't a new format, it's simply a
name for streaming multiple JSON objects without any delimiters.</p>
<p>The advantage of this format is that it can handle JSON objects that have
been formatted with embedded newline characters, e.g., pretty-printed for
human readability. For example, these two inputs are both valid and produce
the same output:</p>
<h4 id="single-line-concatenated-json" tabindex="-1"><a class="header-anchor" href="#single-line-concatenated-json" aria-hidden="true">#</a> Single line concatenated JSON</h4>
<pre><code>{&quot;some&quot;:&quot;thing\n&quot;}{&quot;may&quot;:{&quot;include&quot;:&quot;nested&quot;,&quot;objects&quot;:[&quot;and&quot;,&quot;arrays&quot;]}}
</code></pre>
<h4 id="multi-line-concatenated-json" tabindex="-1"><a class="header-anchor" href="#multi-line-concatenated-json" aria-hidden="true">#</a> Multi-line concatenated JSON</h4>
<pre><code>{
  &quot;some&quot;: &quot;thing\n&quot;
}
{
  &quot;may&quot;: {
    &quot;include&quot;: &quot;nested&quot;,
    &quot;objects&quot;: [
      &quot;and&quot;,
      &quot;arrays&quot;
    ]
  }
}
</code></pre>
</blockquote>
<p>...however in Murex's case, only single line concatenated JSON files
(example 1) are supported; and that is only supported to cover some edge
cases when writing JSON lines and a new line character isn't included. The
primary example might be when generating JSON lines from inside a <code v-pre>for</code> loop.</p>
<p>This is resolved in the new data-type parser <code v-pre>jsonc</code> (Concatenated JSON). See
line below.</p>
<h3 id="more-information" tabindex="-1"><a class="header-anchor" href="#more-information" aria-hidden="true">#</a> More information</h3>
<p>This format is sometimes also referred to as LDJSON and NDJSON, as described
on <a href="https://en.wikipedia.org/wiki/JSON_streaming#Line-delimited_JSON" target="_blank" rel="noopener noreferrer">Wikipedia<ExternalLinkIcon/></a>.</p>
<p>Murex's <RouterLink to="/types/json.html"><code v-pre>json</code> data-type document</RouterLink> also describes some use
cases for JSON lines.</p>
<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2>
<ul>
<li><strong>Extension</strong>: <code v-pre>json-lines</code></li>
<li><strong>Extension</strong>: <code v-pre>jsonl</code></li>
<li><strong>Extension</strong>: <code v-pre>jsonlines</code></li>
<li><strong>Extension</strong>: <code v-pre>ldjson</code></li>
<li><strong>Extension</strong>: <code v-pre>murex_history</code></li>
<li><strong>Extension</strong>: <code v-pre>ndjson</code></li>
<li><strong>MIME</strong>: <code v-pre>application/json-lines</code></li>
<li><strong>MIME</strong>: <code v-pre>application/jsonl</code></li>
<li><strong>MIME</strong>: <code v-pre>application/jsonlines</code></li>
<li><strong>MIME</strong>: <code v-pre>application/ldjson</code></li>
<li><strong>MIME</strong>: <code v-pre>application/ndjson</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-json-lines</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-jsonl</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-jsonlines</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-ldjson</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-ndjson</code></li>
<li><strong>MIME</strong>: <code v-pre>text/json-lines</code></li>
<li><strong>MIME</strong>: <code v-pre>text/jsonl</code></li>
<li><strong>MIME</strong>: <code v-pre>text/jsonlines</code></li>
<li><strong>MIME</strong>: <code v-pre>text/ldjson</code></li>
<li><strong>MIME</strong>: <code v-pre>text/ndjson</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-json-lines</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-jsonl</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-jsonlines</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-ldjson</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-ndjson</code></li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>ReadArray()</code>
Works with JSON arrays. Maps are converted into arrays</li>
<li><code v-pre>ReadArrayWithType()</code>
Works with JSON arrays. Maps are converted into arrays. Element data type is <code v-pre>json</code></li>
<li><code v-pre>ReadIndex()</code>
Works against all properties in JSON</li>
<li><code v-pre>ReadMap()</code>
Not currently supported.</li>
<li><code v-pre>ReadNotIndex()</code>
Works against all properties in JSON</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Supported</li>
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
<li><RouterLink to="/apis/WriteArray.html"><code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/types/csv.html"><code v-pre>csv</code> </RouterLink>:
CSV files (and other character delimited tables)</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/hcl.html"><code v-pre>hcl</code> </RouterLink>:
HashiCorp Configuration Language (HCL)</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/types/jsonc.html"><code v-pre>jsonc</code> </RouterLink>:
Concatenated JSON</li>
<li><RouterLink to="/commands/open.html"><code v-pre>open</code></RouterLink>:
Open a file with a preferred handler</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/types/toml.html"><code v-pre>toml</code> </RouterLink>:
Tom's Obvious, Minimal Language (TOML)</li>
<li><RouterLink to="/types/yaml.html"><code v-pre>yaml</code> </RouterLink>:
YAML Ain't Markup Language (YAML)</li>
<li><RouterLink to="/types/mxjson.html">mxjson</RouterLink>:
Murex-flavoured JSON (deprecated)</li>
</ul>
</div></template>


