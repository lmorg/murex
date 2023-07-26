<template><div><h1 id="reserved-variables" tabindex="-1"><a class="header-anchor" href="#reserved-variables" aria-hidden="true">#</a> Reserved Variables</h1>
<blockquote>
<p>Special variables reserved by Murex</p>
</blockquote>
<p>Murex reserves a few special variables which cannot be assigned via <code v-pre>set</code> nor
<code v-pre>let</code>.</p>
<p>The following is a list of reserved variables, their data type, and its usage:</p>
<h2 id="self-json" tabindex="-1"><a class="header-anchor" href="#self-json" aria-hidden="true">#</a> <code v-pre>SELF</code> (json)</h2>
<p>This returns meta information about the running scope.</p>
<p>A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an <code v-pre>if</code> or
<code v-pre>foreach</code> block will be persistent outside of their blocks as long as you're
still inside the same function.</p>
<p>Please see scoping document (link below) for more information on scoping.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» function example { out $SELF }
» example
{
    "Parent": 11357,
    "Scope": 11357,
    "TTY": true,
    "Method": false,
    "Not": false,
    "Background": false,
    "Module": "murex"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h4 id="parent-num" tabindex="-1"><a class="header-anchor" href="#parent-num" aria-hidden="true">#</a> Parent (num)</h4>
<p>This is the function ID of the parent function that created the scope. In
some instances this will be the same value as scope FID. However if in doubt
then please using <strong>Scope</strong> instead.</p>
<h4 id="scope-num" tabindex="-1"><a class="header-anchor" href="#scope-num" aria-hidden="true">#</a> Scope (num)</h4>
<p>The scope value here returns the function ID of the top level function in the
scope.</p>
<h4 id="tty-bool" tabindex="-1"><a class="header-anchor" href="#tty-bool" aria-hidden="true">#</a> TTY (bool)</h4>
<p>A boolean value as to whether STDOUT is a TTY (ie are we printing to the
terminal (TTY) or a pipe?)</p>
<h4 id="method-bool" tabindex="-1"><a class="header-anchor" href="#method-bool" aria-hidden="true">#</a> Method (bool)</h4>
<p>A boolean value to describe whether the current scope is a method (ie being
called mid-way or at the end of a pipeline).</p>
<h4 id="not-bool" tabindex="-1"><a class="header-anchor" href="#not-bool" aria-hidden="true">#</a> Not (bool)</h4>
<p>A boolean value which represents whether the function was called with a bang-
prefix or not.</p>
<h4 id="background-bool" tabindex="-1"><a class="header-anchor" href="#background-bool" aria-hidden="true">#</a> Background (bool)</h4>
<p>A boolean value to identify whether the current scope is running in the
background for foreground.</p>
<h4 id="module-str" tabindex="-1"><a class="header-anchor" href="#module-str" aria-hidden="true">#</a> Module (str)</h4>
<p>This will be the module string for the current scope.</p>
<h3 id="args-json" tabindex="-1"><a class="header-anchor" href="#args-json" aria-hidden="true">#</a> <code v-pre>ARGS</code> (json)</h3>
<p>This returns a JSON array of the command name and parameters within a given
scope.</p>
<p>Unlike <code v-pre>$PARAMS</code>, <code v-pre>$ARGS</code> includes the function name.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» function example { out $ARGS }
» example abc 1 2 3
[
    "example",
    "abc",
    "1",
    "2",
    "3"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="params-json" tabindex="-1"><a class="header-anchor" href="#params-json" aria-hidden="true">#</a> <code v-pre>PARAMS</code> (json)</h3>
<p>This returns a JSON array of the parameters within a given scope.</p>
<p>Unlike <code v-pre>$ARGS</code>, <code v-pre>$PARAMS</code> does not include the function name.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» function example { out $PARAMS }
» example abc 1 2 3
[
    "abc",
    "1",
    "2",
    "3"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="murex-exe-str" tabindex="-1"><a class="header-anchor" href="#murex-exe-str" aria-hidden="true">#</a> <code v-pre>MUREX_EXE</code> (str)</h3>
<p>This is very similar to the <code v-pre>$SHELL</code> environmental variable in that it holds
the full path to the running shell. The reason for defining a reserved variable
is so that the shell path cannot be overridden.</p>
<h3 id="murex-args-json" tabindex="-1"><a class="header-anchor" href="#murex-args-json" aria-hidden="true">#</a> <code v-pre>MUREX_ARGS</code> (json)</h3>
<p>This is TODO: [https://github.com/lmorg/murex/issues/304](Github issue 304)</p>
<h3 id="hostname-str" tabindex="-1"><a class="header-anchor" href="#hostname-str" aria-hidden="true">#</a> <code v-pre>HOSTNAME</code> (str)</h3>
<p>This returns the hostname of the machine Murex is running from.</p>
<h3 id="_0-str" tabindex="-1"><a class="header-anchor" href="#_0-str" aria-hidden="true">#</a> <code v-pre>0</code> (str)</h3>
<p>This returns the name of the executable (like <code v-pre>$ARGS[0]</code>)</p>
<h3 id="_1-2-3-str" tabindex="-1"><a class="header-anchor" href="#_1-2-3-str" aria-hidden="true">#</a> <code v-pre>1</code>, <code v-pre>2</code>, <code v-pre>3</code>... (str)</h3>
<p>This returns parameter <em>n</em> (like <code v-pre>$ARGS[n]</code>). If there is no parameter <em>n</em>
then the variable will not be set thus the upper limit variable is determined
by how many parameters are set. For example if you have 19 parameters passed
then variables <code v-pre>$1</code> through to <code v-pre>$19</code> (inclusive) will all be set.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/array.html">Array (<code v-pre>@</code>) Token</RouterLink>:
Expand values as an array</li>
<li><RouterLink to="/user-guide/bang-prefix.html">Bang Prefix</RouterLink>:
Bang prefixing to reverse default actions</li>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/user-guide/pipeline.html">Pipeline</RouterLink>:
Overview of what a &quot;pipeline&quot; is</li>
<li><RouterLink to="/parser/string.html">String (<code v-pre>$</code>) Token</RouterLink>:
Expand values as a string</li>
<li><RouterLink to="/user-guide/scoping.html">Variable and Config Scoping</RouterLink>:
How scoping works within Murex</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
</ul>
</div></template>


