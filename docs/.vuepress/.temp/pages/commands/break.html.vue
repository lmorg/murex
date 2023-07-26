<template><div><h1 id="break" tabindex="-1"><a class="header-anchor" href="#break" aria-hidden="true">#</a> <code v-pre>break</code></h1>
<blockquote>
<p>Terminate execution of a block within your processes scope</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>break</code> will terminate execution of a block (eg <code v-pre>function</code>, <code v-pre>private</code>,
<code v-pre>foreach</code>, <code v-pre>if</code>, etc).</p>
<p><code v-pre>break</code> requires a parameter and that parameter is the name of the caller
block you wish to break out of. If it is a <code v-pre>function</code> or <code v-pre>private</code>, then it
will be the name of that function or private. If it is an <code v-pre>if</code> or <code v-pre>foreach</code>
loop, then it will be <code v-pre>if</code> or <code v-pre>foreach</code> (respectively).</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>break block-name
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Exiting an iteration block:</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            break foo
            out "ended"
        }
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Running the above code would output:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» foo
1
2
3
4
5
exit running function
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><strong>Exiting a function:</strong></p>
<p><code v-pre>break</code> can be considered to exhibit the behavior of <em>return</em> (from other
languages) too</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function example {
    if { $USER == "root" } then {
        err "Don't run this as root"
        break example
    }

    # ... do something ...
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Though in this particular use case it is recommended that you use <code v-pre>return</code>
instead, the above code does illustrate how <code v-pre>break</code> behaves.</p>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>break</code> cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling <code v-pre>break bar</code> (which is a different function) inside of the function <code v-pre>foo</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            break bar
            out "ended"
        }
    }
}

function bar {
    foo
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Regardless of whether we run <code v-pre>foo</code> or <code v-pre>bar</code>, both of those functions will
raise the following error:</p>
<pre><code>Error in `break` (7,17): no block found named `bar` within the scope of `foo`
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/continue.html"><code v-pre>continue</code></RouterLink>:
Terminate process of a block within a caller function</li>
<li><RouterLink to="/commands/exit.html"><code v-pre>exit</code></RouterLink>:
Exit murex</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/return.html"><code v-pre>return</code></RouterLink>:
Exits current function scope</li>
</ul>
</div></template>


