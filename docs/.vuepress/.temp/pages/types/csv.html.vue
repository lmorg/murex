<template><div><h1 id="csv" tabindex="-1"><a class="header-anchor" href="#csv" aria-hidden="true">#</a> <code v-pre>csv</code></h1>
<blockquote>
<p>CSV files (and other character delimited tables)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This data type can be used for not only CSV files but also TSV (tab separated)
or any other exotic characters used as a delimiter.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>The CSV parser is configurable via <code v-pre>config</code> (see link below for docs on how to
use <code v-pre>config</code>)</p>
<pre><code>» config: -&gt; [csv]
{
    &quot;comment&quot;: {
        &quot;Data-Type&quot;: &quot;str&quot;,
        &quot;Default&quot;: &quot;#&quot;,
        &quot;Description&quot;: &quot;The prefix token for comments in a CSV table.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: &quot;#&quot;
    },
    &quot;separator&quot;: {
        &quot;Data-Type&quot;: &quot;str&quot;,
        &quot;Default&quot;: &quot;,&quot;,
        &quot;Description&quot;: &quot;The delimiter for records in a CSV file.&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Global&quot;: false,
        &quot;Value&quot;: &quot;,&quot;
    }
}
</code></pre>
<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2>
<ul>
<li><strong>Extension</strong>: <code v-pre>csv</code></li>
<li><strong>MIME</strong>: <code v-pre>application/csv</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-csv</code></li>
<li><strong>MIME</strong>: <code v-pre>text/csv</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-csv</code></li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Supported</li>
<li><code v-pre>ReadArray()</code>
Treats each new line as a new array element</li>
<li><code v-pre>ReadArrayWithType()</code>
Treats each new line as a new array element, each element is mini <code v-pre>csv</code> file</li>
<li><code v-pre>ReadIndex()</code>
Indexes treated as table coordinates</li>
<li><code v-pre>ReadMap()</code>
Works against tables such as the output from <code v-pre>ps -fe</code></li>
<li><code v-pre>ReadNotIndex()</code>
Indexes treated as table coordinates</li>
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
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/int.html"><code v-pre>int</code> </RouterLink>:
Whole number (primitive)</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/optional/select.html"><code v-pre>select</code> </RouterLink>:
Inlining SQL into shell pipelines</li>
<li><RouterLink to="/types/str.html"><code v-pre>str</code> (string) </RouterLink>:
string (primitive)</li>
</ul>
</div></template>


