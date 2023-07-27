<template><div><h1 id="variable-and-config-scoping" tabindex="-1"><a class="header-anchor" href="#variable-and-config-scoping" aria-hidden="true">#</a> Variable and Config Scoping</h1>
<blockquote>
<p>How scoping works within Murex</p>
</blockquote>
<p>A 'scope' in Murex is a collection of code blocks to which variables and
config are persistent within. In Murex, a variable declared inside an <code v-pre>if</code> or
<code v-pre>foreach</code> block will be persistent outside of their blocks as long as you're
still inside the same function.</p>
<p>For example lets start with the following function that sets a variable called
<strong>foo</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>function example {
    if { true } then { set foo=bar }
    out $foo
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>In here the value is getting set inside an <code v-pre>if</code> block but its value is is
retrieved outside of that block. <code v-pre>out</code> and <code v-pre>set</code> have different parents but
the same scoping.</p>
<p>Then lets set <strong>foo</strong> outside of that function and see what happens:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set foo=oof
» $foo
oof

» example
bar

» $foo
oof
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Despite setting a variable named <strong>foo</strong>, the value inside <strong>example</strong> does not
overwrite the value outside <strong>example</strong> because they occupy different scoping.</p>
<h2 id="what-instantiates-a-new-scope" tabindex="-1"><a class="header-anchor" href="#what-instantiates-a-new-scope" aria-hidden="true">#</a> What Instantiates A New Scope?</h2>
<p>A new scope is instantiated by anything which resembles a function. This would
be code inside events, dynamic autocompletes, open agents, any code blocks
defined in <code v-pre>config</code>, as well as public and private functions too.</p>
<p>Code inside an <code v-pre>if</code>, <code v-pre>switch</code>, <code v-pre>foreach</code> and <code v-pre>source</code> do not create a new
scope. Subshells also do not create a new scoping either.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/event.html"><code v-pre>event</code></RouterLink>:
Event driven programming for shell scripts</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/openagent.html"><code v-pre>openagent</code></RouterLink>:
Creates a handler function for `open</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/source.html"><code v-pre>source</code> </RouterLink>:
Import Murex code from another file of code block</li>
<li><RouterLink to="/commands/switch.html"><code v-pre>switch</code></RouterLink>:
Blocks of cascading conditionals</li>
</ul>
</div></template>


