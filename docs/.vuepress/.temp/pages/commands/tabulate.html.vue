<template><div><h1 id="tabulate" tabindex="-1"><a class="header-anchor" href="#tabulate" aria-hidden="true">#</a> <code v-pre>tabulate</code></h1>
<blockquote>
<p>Table transformation tools</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>tabluate</code> is a swiss army knife for table transforming human readable tables
into machine readable data structure.</p>
<blockquote>
<p>Please note that this builtin is still in active development and the default
behavior will continue to change and evolve. Any features marked with a flag
(see below) will be stable, have numerous tests written against them, and
thus safe to use.</p>
</blockquote>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; tabulate [ flags ] -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--column-wraps</code>
Boolean, used with --map or --key-value to merge trailing lines if the text wraps within the same column</li>
<li><code v-pre>--help</code>
Boolean, displays a list of flags</li>
<li><code v-pre>--joiner</code>
String, used with --map to concatenate any trailing records in a given field</li>
<li><code v-pre>--key-inc-hint</code>
Boolean, used with --map to split any space or equal delimited hints/examples (eg parsing flags)</li>
<li><code v-pre>--key-value</code>
Boolean, discard any records that don't appear key value pairs (auto-enabled when --map used)</li>
<li><code v-pre>--map</code>
Boolean, return JSON map instead of table</li>
<li><code v-pre>--separator</code>
'String, custom regex pattern for spliting fields (default: <code v-pre>(\t|\s[\s]+)+</code>)'</li>
<li><code v-pre>--split-comma</code>
Boolean, split first field and duplicate the line if comma found in first field (eg parsing flags in help pages)</li>
<li><code v-pre>--split-space</code>
Boolean, split first field and duplicate the line if white space found in first field (eg parsing flags in help pages)</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="dynamic-autocompletion" tabindex="-1"><a class="header-anchor" href="#dynamic-autocompletion" aria-hidden="true">#</a> Dynamic Autocompletion</h3>
<p>Because <code v-pre>tabulate</code> is designed to parse human readable tables, it is used a lot
for dynamically turning command like program help output into JSON maps for
<code v-pre>autocomplete</code>'s <strong>DynamicDesc</strong> blocks:</p>
<pre><code>rsync --help -&gt; @[^Options$..--help]re -&gt; tabulate: --map --split-comma --column-wraps --key-inc-hint
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
</ul>
</div></template>


