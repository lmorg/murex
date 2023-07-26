<template><div><h1 id="bexists" tabindex="-1"><a class="header-anchor" href="#bexists" aria-hidden="true">#</a> <code v-pre>bexists</code></h1>
<blockquote>
<p>Check which builtins exist</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>bexists</code> takes an array of parameters and returns which commands have been
compiled into Murex. The 'b' in <code v-pre>bexists</code> stands for 'builtins'</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>bexists command... -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» bexists: qr gzip runtime config
{
    "Installed": [
        "runtime",
        "config"
    ],
    "Missing": [
        "qr",
        "gzip"
    ]
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>This builtin dates back to the start of Murex when all of the builtins were
considered optional. This was intended to be a way for scripts to determine
which builtins were compiled. Since then <code v-pre>runtime</code> has absorbed and centralized
a number of similar commands which have since been deprecated. The same fate
might also happen to <code v-pre>bexists</code> however it is in use by a few modules and for
that reason alone it has been spared from the axe.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/modules.html">Modules and Packages</RouterLink>:
An introduction to Murex modules and packages</li>
<li><RouterLink to="/commands/fexec.html"><code v-pre>fexec</code> </RouterLink>:
Execute a command or function, bypassing the usual order of precedence.</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


