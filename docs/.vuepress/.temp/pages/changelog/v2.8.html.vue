<template><div><h1 id="what-s-new-in-murex-v2-8-change-log" tabindex="-1"><a class="header-anchor" href="#what-s-new-in-murex-v2-8-change-log" aria-hidden="true">#</a> What's new in murex v2.8 - Change Log</h1>
<p>This release comes with a number of experimental but stable features that might eventually become standard practice. The features are there to use if you with but adjacent from the older code so there is zero risk in updating to this version.</p>
<p>Breaking Changes:</p>
<ul>
<li>None</li>
</ul>
<p>Features:</p>
<ul>
<li>
<p>New data-type <code v-pre>jsonc</code> (Concatenated JSON). This is an experimental parser for reading multiple JSON documents in a single file or stream. The expectation is that it will eventually replace <code v-pre>jsonl</code> (jsonlines) and possibly even the default <code v-pre>json</code> data-type</p>
</li>
<li>
<p>Autocomplete can now split on <code v-pre>=</code>. However this is an experimental feature</p>
</li>
<li>
<p>New <code v-pre>autocomplete</code> directive: <code v-pre>Goto</code>. This is an experimental way to add a little more structure to autocomplete definitions without resorting to using dynamic code everywhere</p>
</li>
<li>
<p><code v-pre>autocomplete</code> directive <code v-pre>FlagValues</code> now supports two values: <code v-pre>&quot;*&quot;</code> and <code v-pre>&quot;&quot;</code>. These define default properties regardless of the flag used</p>
</li>
</ul>
<p>Minor Changes:</p>
<ul>
<li>
<p>Unmatched records in <strong>index</strong> (<code v-pre>[</code>) when parsing a table now produces the same configurable behavior (erroring by default but which can be disabled) as when using <strong>index</strong> against an object</p>
</li>
<li>
<p>Autocompletion suggestions for <code v-pre>yarn</code></p>
</li>
<li>
<p>Data-type field in <code v-pre>function</code> is now optional</p>
</li>
<li>
<p>Lots of improvements to the docs</p>
</li>
</ul>
<p>Bug Fixes:</p>
<ul>
<li><code v-pre>$MyVar[[/path/to/element]]</code> now behaves the same when used infixed as it did when used as a function</li>
</ul>
<hr>
<p>Published: 23.05.2022 at 22:23</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/types/jsonc.html"><code v-pre>jsonc</code> </RouterLink>:
Concatenated JSON</li>
</ul>
</div></template>


