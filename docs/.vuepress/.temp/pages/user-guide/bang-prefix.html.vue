<template><div><h1 id="bang-prefix" tabindex="-1"><a class="header-anchor" href="#bang-prefix" aria-hidden="true">#</a> Bang Prefix</h1>
<blockquote>
<p>Bang prefixing to reverse default actions</p>
</blockquote>
<p>Some builtins support a bang prefix, <code v-pre>!</code>, which provides a shorthand negative
action to default behavior. For example, <code v-pre>set</code> defines a variable where as
<code v-pre>!set</code> will undefine a variable.</p>
<p>Sometimes the shortcut will be logical, like a <strong>not</strong> operator, as is the case
with <code v-pre>and</code> where typically each result has to equal <strong>true</strong> normally or
<strong>false</strong> if used in <code v-pre>!and</code>.</p>
<p>Sometimes the shortcut will be more philosophical, such as with <code v-pre>config</code> where
normal operations is to query or set configuration but <code v-pre>!config</code> resets the
configuration to defaults (thus operating the same as <code v-pre>config default</code>).</p>
<p>Please read the respecting commands doc for details on whether it supports a
bang prefix and what the behavior of that prefix is.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/commands/not.html"><code v-pre>!</code> (not)</RouterLink>:
Reads the STDIN and exit number from previous process and not's it's condition</li>
<li><RouterLink to="/commands/and.html"><code v-pre>and</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether multiple conditions are met</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/or.html"><code v-pre>or</code></RouterLink>:
Returns <code v-pre>true</code> or <code v-pre>false</code> depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


