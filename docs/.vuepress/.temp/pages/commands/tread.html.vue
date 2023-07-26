<template><div><h1 id="tread" tabindex="-1"><a class="header-anchor" href="#tread" aria-hidden="true">#</a> <code v-pre>tread</code></h1>
<blockquote>
<p><code v-pre>read</code> a line of input from the user and store as a user defined <em>typed</em> variable (deprecated)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>A readline function to allow a line of data inputted from the terminal and then
store that as a typed variable.</p>
<p><strong>This builtin is now deprecated. Please use <code v-pre>read --datatype ...</code> instead</strong></p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>tread: data-type &quot;prompt&quot; var_name

&lt;stdin&gt; -&gt; tread: data-type var_name
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>tread: qs &quot;Please paste a URL: &quot; url
out: &quot;The query string values included were:&quot;
$url -&gt; format json

out: Please paste a URL: -&gt; tread: qs url
out: &quot;The query string values included were:&quot;
$url -&gt; format json
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>If <code v-pre>tread</code> is called as a method then the prompt string is taken from STDIN.
Otherwise the prompt string will be the first parameter. However if no prompt
string is given then <code v-pre>tread</code> will not write a prompt.</p>
<p>The last parameter will be the variable name to store the string read by <code v-pre>tread</code>.
This variable cannot be prefixed by dollar, <code v-pre>$</code>, otherwise the shell will write
the output of that variable as the last parameter rather than the name of the
variable.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/format.html"><code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/pretty.html"><code v-pre>pretty</code></RouterLink>:
Prettifies JSON to make it human readable</li>
<li><RouterLink to="/commands/read.html"><code v-pre>read</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a variable</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
</ul>
</div></template>


