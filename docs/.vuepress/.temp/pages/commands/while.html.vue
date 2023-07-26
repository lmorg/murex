<template><div><h1 id="while" tabindex="-1"><a class="header-anchor" href="#while" aria-hidden="true">#</a> <code v-pre>while</code></h1>
<blockquote>
<p>Loop until condition false</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>while</code> loops until loops until <strong>condition</strong> is false.</p>
<p>Normally the <strong>conditional</strong> and executed code block are 2 separate parameters
however you can call <code v-pre>while</code> with just 1 parameter where the code block acts
as both the conditional and the code to be ran.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Until true</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>while { condition } { code-block } -> &lt;stdout>

while { code-block } -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Until false</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>!while { condition } { code-block } -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><code v-pre>while</code> <strong>$i</strong> is less then <strong>5</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» let i=0; while { =i&lt;5 } { let i=i+1; out $i }
1
2
3
4
5

» let i=0; while { let i=i+1; = i&lt;5; out }
true
true
true
true
false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code v-pre>while</code> <strong>$i</strong> is <em>NOT</em> greater than or equal to <strong>5</strong></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» let i=0; !while { =i>=5 } { let i=i+1; out $i }
1
2
3
4
5

» let i=0; while { let i=i+1; = i>=5; out }
true
true
true
true
false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="meta-values" tabindex="-1"><a class="header-anchor" href="#meta-values" aria-hidden="true">#</a> Meta values</h3>
<p>Meta values are a JSON object stored as the variable <code v-pre>$.</code>. The meta variable
will get overwritten by any other block which invokes meta values. So if you
wish to persist meta values across blocks you will need to reassign <code v-pre>$.</code>, eg</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>%[1..3] -> foreach {
    meta_parent = $.
    %[7..9] -> foreach {
        out "$(meta_parent.i): $.i"
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The following meta values are defined:</p>
<ul>
<li><code v-pre>i</code>: iteration number</li>
</ul>
<h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>while</code></li>
<li><code v-pre>!while</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/err.html"><code v-pre>err</code></RouterLink>:
Print a line to the STDERR</li>
<li><RouterLink to="/commands/for.html"><code v-pre>for</code></RouterLink>:
A more familiar iteration loop to existing developers</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/global.html"><code v-pre>global</code></RouterLink>:
Define a global variable and set it's value</li>
<li><RouterLink to="/commands/let.html"><code v-pre>let</code></RouterLink>:
Evaluate a mathematical function and assign to variable (deprecated)</li>
<li><RouterLink to="/commands/out.html"><code v-pre>out</code></RouterLink>:
Print a string to the STDOUT with a trailing new line character</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


