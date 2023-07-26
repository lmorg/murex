<template><div><h1 id="formap" tabindex="-1"><a class="header-anchor" href="#formap" aria-hidden="true">#</a> <code v-pre>formap</code></h1>
<blockquote>
<p>Iterate through a map or other collection of data</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>formap</code> is a generic tool for iterating through a map, table or other
sequences of data similarly like a <code v-pre>foreach</code>. In fact <code v-pre>formap</code> can even be
used on array too.</p>
<p>Unlike <code v-pre>foreach</code>, <code v-pre>formap</code>'s default output is <code v-pre>str</code>, so each new line will be
treated as a list item. This behaviour will differ if any additional flags are
used with <code v-pre>foreach</code>, such as <code v-pre>--jmap</code>.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p><code v-pre>formap</code> writes a list:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    &lt;stdin> -> foreach variable { code-block } -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p><code v-pre>formap</code> writes to a buffered JSON map:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>    &lt;stdin> -> formap --jmap key value { code-block (map key) } { code-block (map value) } -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>First of all lets assume the following dataset:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>set json people={
    "Tom": {
        "Age": 32,
        "Gender": "Male"
    },
    "Dick": {
        "Age": 43,
        "Gender": "Male"
    },
    "Sally": {
        "Age": 54,
        "Gender": "Female"
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>We can create human output from this:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $people -> formap key value { out "$key is $value[Age] years old" }
Sally is 54 years old
Tom is 32 years old
Dick is 43 years old
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><blockquote>
<p>Please note that maps are intentionally unsorted so you cannot guarantee the
order of the output produced even if the input has been superficially set in
a specific order.</p>
</blockquote>
<p>With <code v-pre>--jmap</code> we can turn that structure into a new structure:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $people -> formap --jmap key value { $key } { $value[Age] }
{
    "Dick": "43",
    "Sally": "54",
    "Tom": "32"
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--jmap</code>
Write a <code v-pre>json</code> map to STDOUT instead of an array</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p><code v-pre>formap</code> can also work against arrays and tables as well. However <code v-pre>foreach</code> is
a much better tool for ordered lists and tables can look a little funky when
when there are more than 2 columns. In those instances you're better off using
<code v-pre>[</code> (index) to specify columns and then <code v-pre>tabulate</code> for any data transformation.</p>
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
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/break.html"><code v-pre>break</code></RouterLink>:
Terminate execution of a block within your processes scope</li>
<li><RouterLink to="/commands/for.html"><code v-pre>for</code></RouterLink>:
A more familiar iteration loop to existing developers</li>
<li><RouterLink to="/commands/foreach.html"><code v-pre>foreach</code></RouterLink>:
Iterate through an array</li>
<li><RouterLink to="/types/json.html"><code v-pre>json</code> </RouterLink>:
JavaScript Object Notation (JSON)</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
<li><RouterLink to="/commands/tabulate.html"><code v-pre>tabulate</code></RouterLink>:
Table transformation tools</li>
<li><RouterLink to="/commands/while.html"><code v-pre>while</code></RouterLink>:
Loop until condition false</li>
</ul>
</div></template>


