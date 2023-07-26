<template><div><h1 id="alter" tabindex="-1"><a class="header-anchor" href="#alter" aria-hidden="true">#</a> <code v-pre>alter</code></h1>
<blockquote>
<p>Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>alter</code> a value within a structured data-type.</p>
<p>The path separater is defined by the first character in the path. For example
<code v-pre>/path/to/key</code>, <code v-pre>,path,to,key</code>, <code v-pre>|path|to|key</code> and <code v-pre>#path#to#key</code> are all valid
however you should remember to quote or escape any special characters (tokens)
used by the shell (such as pipe, <code v-pre>|</code>, and hash, <code v-pre>#</code>).</p>
<p>The <em>value</em> must always be supplied as JSON however</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>`&lt;stdin&gt;` -&gt; alter: [ -m | --merge | -s | --sum ] /path value -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: -> [ shell ] -> [ prompt ] -> alter: /Value moo
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Value": "moo"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>alter</code> also accepts JSON as a parameter for adding structured data:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config: -> [ shell ] -> [ prompt ] -> alter: /Example { "Foo": "Bar" }
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Example": {
        "Foo": "Bar"
    },
    "Value": "{ out 'murex » ' }"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>However it is also data type aware so if they key you're updating holds a string
(for example) then the JSON data a will be stored as a string:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config: -> [ shell ] -> [ prompt ] -> alter: /Value { "Foo": "Bar" }
{
    "Data-Type": "block",
    "Default": "{ out 'murex » ' }",
    "Description": "Interactive shell prompt.",
    "Value": "{ \"Foo\": \"Bar\" }"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Numbers will also follow the same transparent conversion treatment:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» tout: json { "one": 1, "two": 2 } -> alter: /two "3"
{
    "one": 1,
    "two": 3
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Please note: <code v-pre>alter</code> is not changing the value held inside <code v-pre>config</code> but
instead took the STDOUT from <code v-pre>config</code>, altered a value and then passed that
new complete structure through it's STDOUT.</p>
<p>If you require modifying a structure inside Murex config (such as http
headers) then you can use <code v-pre>config alter</code>. Read the config docs for reference.</p>
</blockquote>
<h3 id="m-merge" tabindex="-1"><a class="header-anchor" href="#m-merge" aria-hidden="true">#</a> -m / --merge</h3>
<p>Thus far all the examples have be changing existing keys. However you can also
alter a structure by appending to an array or a merging two maps together. You
do this with the <code v-pre>--merge</code> (or <code v-pre>-m</code>) flag.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» out: a\nb\nc -> alter: --merge / ([ "d", "e", "f" ])
a
b
c
d
e
f
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="s-sum" tabindex="-1"><a class="header-anchor" href="#s-sum" aria-hidden="true">#</a> -s / --sum</h3>
<p>This behaves similarly to <code v-pre>--merge</code> where structures are blended together.
However where a map exists with two keys the same and the values are numeric,
those values are added together.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» tout json { "a": 1, "b": 2 } -> alter --sum / { "b": 3, "c": 4 }
{
    "a": 1,
    "b": 5,
    "c": 4
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--merge</code>
Merge data structures rather than overwrite</li>
<li><code v-pre>--sum</code>
Sum values in a map, merge items in an array</li>
<li><code v-pre>-m</code>
Alias for `--merge</li>
<li><code v-pre>-s</code>
Alias for `--sum</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="path" tabindex="-1"><a class="header-anchor" href="#path" aria-hidden="true">#</a> Path</h3>
<p>The path parameter can take any character as node separators. The separator is
assigned via the first character in the path. For example</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config -> alter: .shell.prompt.Value moo
config -> alter: >shell>prompt>Value moo
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Just make sure you quote or escape any characters used as shell tokens. eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>config -> alter: '#shell#prompt#Value' moo
config -> alter: ' shell prompt Value' moo
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="supported-data-types" tabindex="-1"><a class="header-anchor" href="#supported-data-types" aria-hidden="true">#</a> Supported data-types</h3>
<p>The <em>value</em> field must always be supplied as JSON however the <em>STDIN</em> struct
can be any data-type supported by murex.</p>
<p>You can check what data-types are available via the <code v-pre>runtime</code> command:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>runtime --marshallers
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Marshallers are enabled at compile time from the <code v-pre>builtins/data-types</code> directory.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/append.html"><code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/prepend.html"><code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


