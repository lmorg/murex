<template><div><h1 id="datetime" tabindex="-1"><a class="header-anchor" href="#datetime" aria-hidden="true">#</a> <code v-pre>datetime</code></h1>
<blockquote>
<p>A date and/or time conversion tool (like <code v-pre>printf</code> but for date and time values)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>While <code v-pre>date</code> is a standard UNIX tool, it's syntax can vary from Linux to macOS.
<code v-pre>datetype</code> aims to be a cross platform alternative while also offering a range
of saner syntax options too.</p>
<p>The syntax for <code v-pre>datetime</code> is modelled from date and time libraries from various
popular programming languages.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Pass date/time value as a parameter:</p>
<pre><code>datetime --in &quot;format&quot; --out &quot;format&quot; --value &quot;date/time&quot; -&gt; `&lt;stdout&gt;`
</code></pre>
<p>Read date/time value from STDIN:</p>
<pre><code>`&lt;stdin&gt;` -&gt; datetime --in &quot;format&quot; --out &quot;format&quot; -&gt; `&lt;stdout&gt;`
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Output current date and time:</p>
<pre><code>» datetime: --in &quot;{now}&quot; --out &quot;{go}01/02/06 15:04:05&quot;
12/08/21 22:32:30
</code></pre>
<p>Convert STDIN into epoch:</p>
<pre><code>» echo &quot;12/08/21 22:32:30&quot; -&gt; datetime: --in &quot;{go}01/02/06 15:04:05&quot; --out &quot;{unix}&quot;
1639002750
</code></pre>
<p>Convert value passed as a command line argument:</p>
<pre><code>» datetime: --value &quot;12/08/21 22:32:30&quot; --in &quot;{go}01/02/06 15:04:05&quot; --out &quot;{unix}&quot;
1639002750
</code></pre>
<h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--in</code>
Defines the date/time string is formatted in `--value</li>
<li><code v-pre>--out</code>
Defined how the date/time string should be formatted in STDOUT</li>
<li><code v-pre>--value</code>
Date/time value to convert (if omitted and the input format is not set to <code v-pre>{now}</code> then date/time is read from STDIN)</li>
</ul>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="date-time-format-code-parsers" tabindex="-1"><a class="header-anchor" href="#date-time-format-code-parsers" aria-hidden="true">#</a> Date Time Format Code Parsers</h3>
<p><code v-pre>datetime</code> supports a number of parsers, defined in curly braces.</p>
<h4 id="py-python-strftime-strptime-format-codes" tabindex="-1"><a class="header-anchor" href="#py-python-strftime-strptime-format-codes" aria-hidden="true">#</a> <code v-pre>{py}</code>: Python strftime / strptime format codes</h4>
<p>Murex doesn't support all the language and locale features of Python, instead
defaulting to English. However enough support is there to cover most use cases.</p>
<p>Documentation regarding these format codes can be found on <a href="https://docs.python.org/3/library/datetime.html#strftime-and-strptime-behavior" target="_blank" rel="noopener noreferrer">docs.python.org<ExternalLinkIcon/></a>.</p>
<h4 id="go-go-lang-time-time-format-codes" tabindex="-1"><a class="header-anchor" href="#go-go-lang-time-time-format-codes" aria-hidden="true">#</a> <code v-pre>{go}</code>: Go (lang) time.Time format codes</h4>
<p>Murex has full support for Go's date/time parsing.</p>
<p>Documentation regarding these format codes can be found on <a href="https://pkg.go.dev/time#pkg-constants" target="_blank" rel="noopener noreferrer">pkg.go.dev<ExternalLinkIcon/></a>.</p>
<h4 id="now-current-date-and-time" tabindex="-1"><a class="header-anchor" href="#now-current-date-and-time" aria-hidden="true">#</a> <code v-pre>{now}</code>: Current date and time</h4>
<p>This is only supported as an input. When it is used <code v-pre>--value</code> flag is not
required.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
</ul>
</div></template>


