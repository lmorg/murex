<template><div><h1 id="return" tabindex="-1"><a class="header-anchor" href="#return" aria-hidden="true">#</a> <code v-pre>return</code></h1>
<blockquote>
<p>Exits current function scope</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>return</code> will terminate execution of a block at the scope level (eg <code v-pre>function</code>,
<code v-pre>private</code>, etc)</p>
<p>Conceptually it is the same as <code v-pre>break</code> except it doesn't require the scope name
as a parameter and you can specify the exit number rather than defaulting to 0.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>return [ exit-number ]
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Setting an exit number:</strong></p>
<pre><code>function example {
    out foo
    return 13
    out bar
}
example
exitnum
</code></pre>
<p>Running the above code would output:</p>
<pre><code>foo
13
</code></pre>
<p><strong>Returning withing an exit number:</strong></p>
<p>If we were to run the same code as above but with <code v-pre>return</code> written without any
parameters (ie instead of <code v-pre>return 13</code> it would be just <code v-pre>return</code>), then you
would see the following output:</p>
<pre><code>foo
0
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Any process that has been initialised within a <code v-pre>return</code>ed scope will have their
exit number updated to the value specified in <code v-pre>return</code> (or <code v-pre>0</code> if no parameter
was passed).</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/continue.html"><code v-pre>continue</code></RouterLink>:
Terminate process of a block within a caller function</li>
<li><RouterLink to="/commands/exit.html"><code v-pre>exit</code></RouterLink>:
Exit murex</li>
<li><RouterLink to="/commands/exitnum.html"><code v-pre>exitnum</code></RouterLink>:
Output the exit number of the previous process</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
</ul>
</div></template>


