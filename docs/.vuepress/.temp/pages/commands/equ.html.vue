<template><div><h1 id="arithmetic-evaluation" tabindex="-1"><a class="header-anchor" href="#arithmetic-evaluation" aria-hidden="true">#</a> <code v-pre>=</code> (arithmetic evaluation)</h1>
<blockquote>
<p>Evaluate a mathematical function (deprecated)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>=</code> evaluates a mathematical function and returns it's output</p>
<p><strong>This is a deprecated feature. Please refer to <RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink> instead.</strong></p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> = evaluation -> &lt;stdout>

= evaluation -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>As a method:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» let: age=18
» $age -> = &lt; 21
true

» $age -> = &lt; 21 -> if { out: "Under 21" } else { out: "Over 21" }
Under 21
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>As a function:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» let: age=18
» = age &lt; 21
true

» = age &lt; 21 -> if { out: "Under 21" } else { out: "Over 21" }
Under 21
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Inlining as a function:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» let: age=18
» if { = age &lt; 21 } then { out: "Under 21" } else { out: "Over 21" }
Under 21
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="variables" tabindex="-1"><a class="header-anchor" href="#variables" aria-hidden="true">#</a> Variables</h3>
<p>There are two ways you can use variables with the math functions. Either by
string interpolation like you would normally with any other function, or
directly by name.</p>
<p>String interpolation:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set abc=123
» = $abc==123
true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Directly by name:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set abc=123
» = abc==123
false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>To understand the difference between the two, you must first understand how
string interpolation works; which is where the parser tokenised the parameters
like so</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command line: = $abc==123
token 1: command (name: "=")
token 2: parameter 1, string (content: "")
token 3: parameter 1, variable (name: "abc")
token 4: parameter 1, string (content: "==123")
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Then when the command line gets executed, the parameters are compiled on demand
similarly to this crude pseudo-code</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command: "="
parameters 1: concatenate("", GetValue(abc), "==123")
output: "=" "123==123"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Thus the actual command getting run is literally <code v-pre>123==123</code> due to the variable
being replace <strong>before</strong> the command executes.</p>
<p>Whereas when you call the variable by name it's up to <code v-pre>=</code> or <code v-pre>let</code> to do the
variable substitution.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>command line: = abc==123
token 1: command (name: "=")
token 2: parameter 1, string (content: "abc==123")

command: "="
parameters 1: concatenate("abc==123")
output: "=" "abc==123"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The main advantage (or disadvantage, depending on your perspective) of using
variables this way is that their data-type is preserved.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set str abc=123
» = abc==123
false

» set int abc=123
» = abc==123
true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Unfortunately is one of the biggest areas in Murex where you'd need to be
careful. The simple addition or omission of the dollar prefix, <code v-pre>$</code>, can change
the behavior of <code v-pre>=</code> and <code v-pre>let</code>.</p>
<h3 id="strings" tabindex="-1"><a class="header-anchor" href="#strings" aria-hidden="true">#</a> Strings</h3>
<p>Because the usual Murex tools for encapsulating a string (<code v-pre>&quot;</code>, <code v-pre>'</code> and <code v-pre>()</code>)
are interpreted by the shell language parser, it means we need a new token for
handling strings inside <code v-pre>=</code> and <code v-pre>let</code>. This is where backtick comes to our
rescue.</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set str abc=123
» = abc==`123`
true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please be mindful that if you use string interpolation then you will need to
instruct <code v-pre>=</code> and <code v-pre>let</code> that your field is a string</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set str abc=123
» = `$abc`==`123`
true
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="best-practice-recommendation" tabindex="-1"><a class="header-anchor" href="#best-practice-recommendation" aria-hidden="true">#</a> Best practice recommendation</h3>
<p>As you can see from the sections above, string interpolation offers us some
conveniences when comparing variables of differing data-types, such as a <code v-pre>str</code>
type with a number (eg <code v-pre>num</code> or <code v-pre>int</code>). However it makes for less readable code
when just comparing strings. Thus the recommendation is to avoid using string
interpolation except only where it really makes sense (ie use it sparingly).</p>
<h3 id="non-boolean-logic" tabindex="-1"><a class="header-anchor" href="#non-boolean-logic" aria-hidden="true">#</a> Non-boolean logic</h3>
<p>Thus far the examples given have been focused on comparisons however <code v-pre>=</code> and
<code v-pre>let</code> supports all the usual arithmetic operators:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» = 10+10
20

» = 10/10
1

» = (4 * (3 + 2))
20

» = `foo`+`bar`
foobar
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h3 id="read-more" tabindex="-1"><a class="header-anchor" href="#read-more" aria-hidden="true">#</a> Read more</h3>
<p>Murex uses the <a href="https://github.com/Knetic/govaluate" target="_blank" rel="noopener noreferrer">govaluate package<ExternalLinkIcon/></a>. More information can be found in it's manual.</p>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>=</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/reserved-vars.html">Reserved Variables</RouterLink>:
Special variables reserved by Murex</li>
<li><RouterLink to="/user-guide/scoping.html">Variable and Config Scoping</RouterLink>:
How scoping works within Murex</li>
<li><RouterLink to="/commands/brace-quote.html"><code v-pre>(</code> (brace quote)</RouterLink>:
Write a string to the STDOUT without new line</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/export.html"><code v-pre>export</code></RouterLink>:
Define an environmental variable and set it's value</li>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
<li><RouterLink to="/commands/global.html"><code v-pre>global</code></RouterLink>:
Define a global variable and set it's value</li>
<li><RouterLink to="/commands/global.html"><code v-pre>global</code></RouterLink>:
Define a global variable and set it's value</li>
<li><RouterLink to="/commands/if.html"><code v-pre>if</code></RouterLink>:
Conditional statement to execute different blocks of code depending on the result of the condition</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


