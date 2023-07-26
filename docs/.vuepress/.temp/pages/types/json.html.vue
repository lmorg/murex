<template><div><h1 id="json" tabindex="-1"><a class="header-anchor" href="#json" aria-hidden="true">#</a> <code v-pre>json</code></h1>
<blockquote>
<p>JavaScript Object Notation (JSON)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>JSON is a structured data-type within Murex. It is the standard format for all
structured data within Murex however other formats such as YAML, TOML and CSV
are equally first class citizens.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example JSON document taken from <a href="https://en.wikipedia.org/wiki/JSON" target="_blank" rel="noopener noreferrer">Wikipedia<ExternalLinkIcon/></a></p>
<pre><code>{
  &quot;firstName&quot;: &quot;John&quot;,
  &quot;lastName&quot;: &quot;Smith&quot;,
  &quot;isAlive&quot;: true,
  &quot;age&quot;: 27,
  &quot;address&quot;: {
    &quot;streetAddress&quot;: &quot;21 2nd Street&quot;,
    &quot;city&quot;: &quot;New York&quot;,
    &quot;state&quot;: &quot;NY&quot;,
    &quot;postalCode&quot;: &quot;10021-3100&quot;
  },
  &quot;phoneNumbers&quot;: [
    {
      &quot;type&quot;: &quot;home&quot;,
      &quot;number&quot;: &quot;212 555-1234&quot;
    },
    {
      &quot;type&quot;: &quot;office&quot;,
      &quot;number&quot;: &quot;646 555-4567&quot;
    },
    {
      &quot;type&quot;: &quot;mobile&quot;,
      &quot;number&quot;: &quot;123 456-7890&quot;
    }
  ],
  &quot;children&quot;: [],
  &quot;spouse&quot;: null
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="tips-when-writing-json-inside-for-loops" tabindex="-1"><a class="header-anchor" href="#tips-when-writing-json-inside-for-loops" aria-hidden="true">#</a> Tips when writing JSON inside for loops</h3>
<p>One of the drawbacks (or maybe advantages, depending on your perspective) of
JSON is that parsers generally expect a complete file for processing in that
the JSON specification requires closing tags for every opening tag. This means
it's not always suitable for streaming. For example</p>
<pre><code>» ja [1..3] -&gt; foreach i { out ({ &quot;$i&quot;: $i }) }
{ &quot;1&quot;: 1 }
{ &quot;2&quot;: 2 }
{ &quot;3&quot;: 3 }
</code></pre>
<p><strong>What does this even mean and how can you build a JSON file up sequentially?</strong></p>
<p>One answer if to write the output in a streaming file format and convert back
to JSON</p>
<pre><code>» ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) }
- &quot;1&quot;: 1
- &quot;2&quot;: 2
- &quot;3&quot;: 3

» ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) } -&gt; cast yaml -&gt; format json
[
    {
        &quot;1&quot;: 1
    },
    {
        &quot;2&quot;: 2
    },
    {
        &quot;3&quot;: 3
    }
]
</code></pre>
<p><strong>What if I'm returning an object rather than writing one?</strong></p>
<p>The problem with building JSON structures from existing structures is that you
can quickly end up with invalid JSON due to the specifications strict use of
commas.</p>
<p>For example in the code below, each item block is it's own object and there are
no <code v-pre>[ ... ]</code> encapsulating them to denote it is an array of objects, nor are
the objects terminated by a comma.</p>
<pre><code>» config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar }
{
    &quot;Data-Type&quot;: &quot;bool&quot;,
    &quot;Default&quot;: true,
    &quot;Description&quot;: &quot;Display the interactive shell's hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: true
}
{
    &quot;Data-Type&quot;: &quot;block&quot;,
    &quot;Default&quot;: &quot;{ progress $PID }&quot;,
    &quot;Description&quot;: &quot;Murex function to execute when an `exec` process is stopped&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: &quot;{ progress $PID }&quot;
}
{
    &quot;Data-Type&quot;: &quot;bool&quot;,
    &quot;Default&quot;: true,
    &quot;Description&quot;: &quot;ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: true
}
...
</code></pre>
<p>Luckily JSON also has it's own streaming format: JSON lines (<code v-pre>jsonl</code>). We can
<code v-pre>cast</code> this output as <code v-pre>jsonl</code> then <code v-pre>format</code> it back into valid JSON:</p>
<pre><code>» config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar } -&gt; cast jsonl -&gt; format json
[
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Write shell history (interactive shell) to disk&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    },
    {
        &quot;Data-Type&quot;: &quot;int&quot;,
        &quot;Default&quot;: 4,
        &quot;Description&quot;: &quot;Maximum number of lines with auto-completion suggestions to display&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: &quot;6&quot;
    },
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    },
...
</code></pre>
<h4 id="foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" tabindex="-1"><a class="header-anchor" href="#foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" aria-hidden="true">#</a> <code v-pre>foreach</code> will automatically cast it's output as <code v-pre>jsonl</code> <em>if</em> it's STDIN type is <code v-pre>json</code></h4>
<pre><code>» ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name }
Hello Tom
Hello Dick
Hello Sally

» ja [Tom,Dick,Sally] -&gt; foreach name { out Hello $name } -&gt; debug -&gt; [[ /Data-Type/Murex ]]
jsonl

» ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name } -&gt; format: json
[
    &quot;Hello Tom&quot;,
    &quot;Hello Dick&quot;,
    &quot;Hello Sally&quot;
]
</code></pre>
<h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2>
<ul>
<li><strong>Extension</strong>: <code v-pre>json</code></li>
<li><strong>MIME</strong>: <code v-pre>application/json</code></li>
<li><strong>MIME</strong>: <code v-pre>application/x-json</code></li>
<li><strong>MIME</strong>: <code v-pre>text/json</code></li>
<li><strong>MIME</strong>: <code v-pre>text/x-json</code></li>
</ul>
<h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2>
<ul>
<li><code v-pre>Marshal()</code>
Writes minified JSON when no TTY detected and human readable JSON when stdout is a TTY</li>
<li><code v-pre>ReadArray()</code>
Works with JSON arrays. Maps are converted into arrays</li>
<li><code v-pre>ReadArrayWithType()</code>
Works with JSON arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the JSON type of the element</li>
<li><code v-pre>ReadIndex()</code>
Works against all properties in JSON</li>
<li><code v-pre>ReadMap()</code>
Works with JSON maps</li>
<li><code v-pre>ReadNotIndex()</code>
Works against all properties in JSON</li>
<li><code v-pre>Unmarshal()</code>
Supported</li>
<li><code v-pre>WriteArray()</code>
Works with JSON arrays</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/Marshal.html"><code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/ReadArray.html"><code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadArrayWithType.html"><code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
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
<li><RouterLink to="/types/hcl.html"><code v-pre>hcl</code> </RouterLink>:
HashiCorp Configuration Language (HCL)</li>
<li><RouterLink to="/types/jsonc.html"><code v-pre>jsonc</code> </RouterLink>:
Concatenated JSON</li>
<li><RouterLink to="/types/jsonl.html"><code v-pre>jsonl</code> </RouterLink>:
JSON Lines</li>
<li><RouterLink to="/apis/lang.ArrayTemplate.html"><code v-pre>lang.ArrayTemplate()</code> (template API)</RouterLink>:
Unmarshals a data type into a Go struct and returns the results as an array</li>
<li><RouterLink to="/apis/lang.ArrayWithTypeTemplate.html"><code v-pre>lang.ArrayWithTypeTemplate()</code> (template API)</RouterLink>:
Unmarshals a data type into a Go struct and returns the results as an array with data type included</li>
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


