<template><div><h1 id="mxjson" tabindex="-1"><a class="header-anchor" href="#mxjson" aria-hidden="true">#</a> mxjson</h1>
<blockquote>
<p>Murex-flavoured JSON (deprecated)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<blockquote>
<p>This format has been deprecated in favour of <code v-pre>%{}</code> constructors.</p>
</blockquote>
<p>mxjson is an extension to JSON designed to integrate more seamlessly when
use as a configuration file. Thus mxjson supports comments and Murex code
blocks embedded into the JSON schema.</p>
<blockquote>
<p>mxjson is a format that is pre-parsed into a valid JSON format.</p>
</blockquote>
<p>mxjson isn't a Murex data-type in that you cannot marshal
and unmarshal mxjson files. Currently it is a format that is only supported
by a small subset of Murex builtins (eg <code v-pre>config</code> and <code v-pre>autocomplete</code>) where
config might embed Murex code blocks.</p>
<p><strong>mxjson features the following enhancements:</strong></p>
<h3 id="line-comments" tabindex="-1"><a class="header-anchor" href="#line-comments" aria-hidden="true">#</a> Line Comments</h3>
<p>Line comments are prefixed with a 'hash', <code v-pre>#</code>, just like with regular Murex
code.</p>
<h3 id="block-quotation" tabindex="-1"><a class="header-anchor" href="#block-quotation" aria-hidden="true">#</a> Block Quotation</h3>
<p>Code blocks are quoted with <code v-pre>(</code>, <code v-pre>)</code>. For example, below &quot;ExampleFunction&quot;
uses the <code v-pre>({ block quote })</code> method.</p>
<pre><code>{
    &quot;ExampleFunction&quot;: ({
        out: &quot;This is an example Murex function&quot;
        if { =1==2 } then {
            err: &quot;The laws of the universe are broken&quot;
        }
    })
}
</code></pre>
<p>Any block quoted by this method will be converted to the following valid JSON:</p>
<pre><code>{
    &quot;ExampleFunction&quot;: &quot;\n    out: \&quot;This is an example Murex function\&quot;\n    if { =1==2 } then {\n        err: \&quot;The laws of the universe are broken\&quot;\n    }&quot;
}
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/brace-quote.html">Brace Quote (<code v-pre>%(</code>, <code v-pre>)</code>) Tokens</RouterLink>:
Initiates or terminates a string (variables expanded)</li>
<li><RouterLink to="/user-guide/code-block.html">Code Block Parsing</RouterLink>:
Overview of how code blocks are parsed</li>
<li><RouterLink to="/parser/create-array.html">Create array (<code v-pre>%[]</code>) constructor</RouterLink>:
Quickly generate arrays</li>
<li><RouterLink to="/parser/curly-brace.html">Curly Brace (<code v-pre>{</code>, <code v-pre>}</code>) Tokens</RouterLink>:
Initiates or terminates a code block</li>
<li><RouterLink to="/apis/Marshal.html"><code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/Unmarshal.html"><code v-pre>Unmarshal()</code> (type)</RouterLink>:
Converts a structured file format into structured memory</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/types/hcl.html"><code v-pre>hcl</code> </RouterLink>:
HashiCorp Configuration Language (HCL)</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/types/jsonc.html"><code v-pre>jsonc</code> </RouterLink>:
Concatenated JSON</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
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
<li><RouterLink to="/parser/create-object.html"><code v-pre>create-object</code></RouterLink>:</li>
</ul>
</div></template>


