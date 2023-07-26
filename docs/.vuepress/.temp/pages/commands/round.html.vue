<template><div><h1 id="round" tabindex="-1"><a class="header-anchor" href="#round" aria-hidden="true">#</a> <code v-pre>round</code></h1>
<blockquote>
<p>Round a number by a user defined precision</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>round</code> supports a few different levels of precision:</p>
<h3 id="nearest-decimal-place" tabindex="-1"><a class="header-anchor" href="#nearest-decimal-place" aria-hidden="true">#</a> Nearest decimal place</h3>
<p>Syntax: <code v-pre>0.12345</code> (any numbers can be used)</p>
<p>If a decimal place is supplied then <code v-pre>round</code> will round your number to however
many decimal places you specify. It doesn't matter what digits you include in
your precision value as the only thing which is used to drive the result is the
position of the decimal point. Thus a precision value of <code v-pre>0.000</code> would perform
the same rounding as <code v-pre>9.999</code>.</p>
<p>Decimal places are always rounded to the nearest. <code v-pre>--down</code> and <code v-pre>--up</code> flags are
not supported.</p>
<h3 id="nearest-integer" tabindex="-1"><a class="header-anchor" href="#nearest-integer" aria-hidden="true">#</a> Nearest integer</h3>
<p>Syntax: either <code v-pre>0</code> or <code v-pre>1</code> (either behaves the same)</p>
<p>This will round your value to the nearest whole number. For example <code v-pre>3.33</code>
would be rounded to <code v-pre>3</code>.</p>
<p>If <code v-pre>--down</code> flag is supplied then the remainder is dropped. For example <code v-pre>9.99</code>
would then be rounded to <code v-pre>9</code> instead of <code v-pre>10</code>.</p>
<p>If <code v-pre>--up</code> flag is is supplied then the the input value would always be rounded
up to the nearest whole number. For example <code v-pre>3.33</code> would be rounded to <code v-pre>4</code>
instead of <code v-pre>3</code>.</p>
<h3 id="nearest-multiple" tabindex="-1"><a class="header-anchor" href="#nearest-multiple" aria-hidden="true">#</a> Nearest Multiple</h3>
<p>Syntax: <code v-pre>50</code> (any integer greater than `1)</p>
<p>This will round your input value to the nearest multiple of your precision.</p>
<p>Like with <strong>nearest integer</strong> (see above), <code v-pre>--down</code> and <code v-pre>--up</code> will specify to
round whether to always round down or up rather than returning the nearest
match in either direction.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>round value precision -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p><strong>Rounding to the nearest multiple of <code v-pre>20</code>:</strong></p>
<pre><code>» round 15 20
20
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--down</code>
Rounds down to the nearest multiple (not supported when precision is to decimal places)</li>
<li><code v-pre>--up</code>
Rounds up to the nearest multiple (not supported when precision is to decimal places)</li>
<li><code v-pre>-d</code>
shorthand for `--down</li>
<li><code v-pre>-u</code>
shorthand for `--up</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/expr.html"><code v-pre>expr</code></RouterLink>:
Expressions: mathematical, string comparisons, logical operators</li>
</ul>
</div></template>


