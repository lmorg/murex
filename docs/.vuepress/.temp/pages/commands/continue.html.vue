<template><div><h1 id="continue" tabindex="-1"><a class="header-anchor" href="#continue" aria-hidden="true">#</a> <code v-pre>continue</code></h1>
<blockquote>
<p>Terminate process of a block within a caller function</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>continue</code> will terminate execution of a block (eg <code v-pre>function</code>, <code v-pre>private</code>,
<code v-pre>foreach</code>, <code v-pre>if</code>, etc) right up until the caller function. In iteration loops
like <code v-pre>foreach</code> and <code v-pre>formap</code> this will result in behavior similar to the
<code v-pre>continue</code> statement in other programming languages.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>continue block-name
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[1..10] -> foreach i {
    if { $i == 5 } then {
        out "continue"
        continue foreach
        out "skip this code"
    }
    out $i
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Running the above code would output:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» foo
1
2
3
4
continue
6
7
8
9
10
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>continue</code> cannot escape the bounds of its scope (typically the function it is
running inside). For example, in the following code we are calling <code v-pre>continue bar</code> (which is a different function) inside of the function <code v-pre>foo</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function foo {
    %[1..10] -> foreach i {
        out $i
        if { $i == 5 } then {
            out "exit running function"
            continue bar
            out "ended"
        }
    }
}

function bar {
    foo
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Regardless of whether we run <code v-pre>foo</code> or <code v-pre>bar</code>, both of those functions will
raise the following error:</p>
<pre><code>Error in `continue` (7,17): no block found named `bar` within the scope of `foo`
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
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


