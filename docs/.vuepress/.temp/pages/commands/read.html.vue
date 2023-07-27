<template><div><h1 id="read" tabindex="-1"><a class="header-anchor" href="#read" aria-hidden="true">#</a> <code v-pre>read</code></h1>
<blockquote>
<p><code v-pre>read</code> a line of input from the user and store as a variable</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>A readline function to allow a line of data inputed from the terminal.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Classic usage:</p>
<pre><code>read: &quot;prompt&quot; var_name

`&lt;stdin&gt;` -&gt; read: var_name
</code></pre>
<p>Script usage:</p>
<pre><code>read: [ --prompt &quot;prompt&quot;         ]
      [ --variable var_name       ]
      [ --default &quot;default value&quot; ]
      [ --datatype data-type      ]
      [ --autocomplete { json }   ]
      [ --mask character          ]
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Classic usage:</strong></p>
<pre><code>read: &quot;What is your name? &quot; name
out: &quot;Hello $name&quot;

out: What is your name? -&gt; read: name
out: &quot;Hello $name&quot;
</code></pre>
<p><strong>Script usage:</strong></p>
<pre><code>read: --prompt &quot;Are you sure? [Y/n]&quot; \
      --variable yn \
      --default Y
</code></pre>
<p><strong>Secrets:</strong></p>
<pre><code>read: --prompt &quot;Password: &quot; --variable pw --mask *
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--autocomplete</code>
Autocompletion suggestions. Can be either a JSON array or a JSON object</li>
<li><code v-pre>--datatype</code>
Murex data-type for the read data (default: str)</li>
<li><code v-pre>--default</code>
If a zero length string is returned but neither <code v-pre>ctrl</code>+<code v-pre>c</code> nor <code v-pre>ctrl</code>+<code v-pre>d</code> were pressed, then the default value defined here will be returned</li>
<li><code v-pre>--mask</code>
Optional password mask, for reading secrets</li>
<li><code v-pre>--prompt</code>
User notification to display</li>
<li><code v-pre>--variable</code>
Variable name to store the read data (default: read)</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="classic-usage" tabindex="-1"><a class="header-anchor" href="#classic-usage" aria-hidden="true">#</a> Classic Usage</h3>
<p>If <code v-pre>read</code> is called as a method then the prompt string is taken from STDIN.
Otherwise the prompt string will be the first parameter. However if no prompt
string is given then <code v-pre>read</code> will not write a prompt.</p>
<p>The last parameter will be the variable name to store the string read by <code v-pre>read</code>.
This variable cannot be prefixed by dollar, <code v-pre>$</code>, otherwise the shell will write
the output of that variable as the last parameter rather than the name of the
variable.</p>
<p>The data type the <code v-pre>read</code> line will be stored as is <code v-pre>str</code> (string). If you
require this to be different then please use <code v-pre>tread</code> (typed read) or call <code v-pre>read</code>
with the <code v-pre>--datatype</code> flag as per the <strong>script usage</strong>.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/cast.html"><code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/tout.html"><code v-pre>tout</code></RouterLink>:
Print a string to the STDOUT and set it's data-type</li>
<li><RouterLink to="/commands/tread.html"><code v-pre>tread</code></RouterLink>:
<code v-pre>read</code> a line of input from the user and store as a user defined <em>typed</em> variable (deprecated)</li>
</ul>
</div></template>


