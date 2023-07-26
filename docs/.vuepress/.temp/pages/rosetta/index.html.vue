<template><div><h1 id="rosetta-stone" tabindex="-1"><a class="header-anchor" href="#rosetta-stone" aria-hidden="true">#</a> Rosetta Stone</h1>
<blockquote>
<p>A tabulated list of Bashism's and their equivalent Murex syntax</p>
</blockquote>
<p>Below is a reference table of common Bash code and how it could be written in
Murex.</p>
<p>It is also recommended that you read the language <a href="/tour">tour</a>
if you want to learn more about shell scripting in Murex.</p>
<table>
<thead>
<tr>
<th>Description</th>
<th>Bash</th>
<th>Murex</th>
</tr>
</thead>
<tbody>
<tr>
<td><RouterLink to="/commands/out.html">Write to STDOUT</RouterLink></td>
<td><code v-pre>echo &quot;Hello Bash&quot;</code></td>
<td><code v-pre>out &quot;Hello Murex&quot;</code> <br/><br/><code v-pre>echo &quot;Hello Murex&quot;</code> <a href="#footnotes">[1]</a></td>
</tr>
<tr>
<td><RouterLink to="/commands/err.html">Write to STDERR</RouterLink></td>
<td><code v-pre>echo &quot;Hello Bash&quot; &gt;2</code></td>
<td><code v-pre>err &quot;Hello Murex&quot;</code></td>
</tr>
<tr>
<td>Write to file (truncate)</td>
<td><code v-pre>echo &quot;Hello Bash&quot; &gt; hello.txt</code></td>
<td><code v-pre>echo &quot;Hello Murex&quot; |&gt; hello.txt</code></td>
</tr>
<tr>
<td>Write to file (append)</td>
<td><code v-pre>echo &quot;Hello Bash&quot; &gt;&gt; hello.txt</code></td>
<td><code v-pre>echo &quot;Hello Murex&quot; &gt;&gt; hello.txt</code></td>
</tr>
<tr>
<td><RouterLink to="/parser/pipe-arrow.html">Pipe commands</RouterLink></td>
<td><code v-pre>echo &quot;Hello Bash | grep Bash</code></td>
<td><code v-pre>echo &quot;Hello Murex | grep Murex</code> <br/><br/> <code v-pre>out &quot;Hello Murex&quot; -&gt; regexp m/Murex/</code></td>
</tr>
<tr>
<td><RouterLink to="/parser/pipe-err.html">Redirect errors to STDOUT</RouterLink></td>
<td><code v-pre>curl murex.rocks 2&gt;&amp;1 | less</code></td>
<td><code v-pre>curl murex.rocks ? less</code> <br/><br/> <code v-pre>curl &lt;!out&gt; murex.rocks | less</code></td>
</tr>
<tr>
<td>Redirect output to STDERR</td>
<td><code v-pre>uname -a &gt;&amp;2</code></td>
<td><code v-pre>uname &lt;err&gt; -a</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Quoting strings</strong></td>
<td></td>
<td></td>
</tr>
<tr>
<td><RouterLink to="/parser/double-quote.html">Infixing</RouterLink></td>
<td><code v-pre>echo &quot;Hello $SHELL&quot;</code></td>
<td><code v-pre>out &quot;Hello $SHELL&quot;</code></td>
</tr>
<tr>
<td><RouterLink to="/parser/single-quote.html">String literals</RouterLink></td>
<td><code v-pre>echo 'Hello' $SHELL</code></td>
<td><code v-pre>out 'Hello' $SHELL</code></td>
</tr>
<tr>
<td><RouterLink to="/parser/brace-quote.html">Nesting quotes</RouterLink></td>
<td><code v-pre>echo 'Hello \'Bob\''</code></td>
<td><code v-pre>out %(Hello 'Bob')</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Process management</strong></td>
<td></td>
<td></td>
</tr>
<tr>
<td><RouterLink to="/commands/exitnum.html">Exit number</RouterLink></td>
<td><code v-pre>$?</code></td>
<td><code v-pre>exitnum</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/bg.html">Background jobs</RouterLink></td>
<td><code v-pre>command &amp;</code></td>
<td><code v-pre>bg { command }</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/fid-list.html">Job control</RouterLink></td>
<td><code v-pre>ps</code>,<br/><code v-pre>jobs</code>,<br/><code v-pre>bg pid</code>,<br/><code v-pre>fg pid</code></td>
<td><code v-pre>fid-list</code>,<br/><code v-pre>jobs</code>,<br/><code v-pre>bg fid</code>,<br/><code v-pre>fg fid</code></td>
</tr>
<tr>
<td>Happy paths</td>
<td><code v-pre>command &amp;&amp; command</code></td>
<td><code v-pre>command &amp;&amp; command</code> <br/><br/> <code v-pre>try {command; command}</code></td>
</tr>
<tr>
<td>Unhappy paths</td>
<td><code v-pre>command || command</code></td>
<td><code v-pre>command || command</code> <br/><br/> <code v-pre>try {command}; catch {command}</code></td>
</tr>
<tr>
<td>Pipe fail</td>
<td><code v-pre>set -o pipefail</code></td>
<td><code v-pre>runmode trypipe module</code> <br/><br/> <code v-pre>runmode trypipe function</code> <br/><br/> <code v-pre>trypipe { commands }</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Comments</strong></td>
<td></td>
<td></td>
</tr>
<tr>
<td>Single line</td>
<td><code v-pre># comment</code></td>
<td><code v-pre># comment</code></td>
</tr>
<tr>
<td>Multiple lines</td>
<td><code v-pre>:&lt;&lt;EOC</code><br/><code v-pre>line 1</code><br/><code v-pre>line 2</code><br/><code v-pre>EOC</code></td>
<td><code v-pre>/#</code><br/><code v-pre>line 1</code><br/><code v-pre>line 2</code><br/><code v-pre>#/</code></td>
</tr>
<tr>
<td>Mid-line</td>
<td>n/a</td>
<td>eg <code v-pre>out foo/#comment#/bar</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>File pattern matching</strong><br/>(also known as &quot;wildcards&quot;)</td>
<td></td>
<td></td>
</tr>
<tr>
<td><RouterLink to="/commands/g.html">Globbing</RouterLink></td>
<td>eg <code v-pre>ls *.txt</code></td>
<td>eg <code v-pre>ls *.txt</code> (in the interactive terminal) <br/><br/> <code v-pre>g pattern</code> <br/><br/> eg <code v-pre>ls @{g *.txt}</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/rx.html">Regexp</RouterLink></td>
<td>n/a</td>
<td><code v-pre>rx pattern</code> <br/><br/> eg <code v-pre>ls @{rx '*\\.txt'}</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/f.html">File type matching</RouterLink></td>
<td>n/a</td>
<td><code v-pre>f flags</code> <br/><br/> eg <code v-pre>f +s</code> (only return symlinks)</td>
</tr>
<tr>
<td>Chaining</td>
<td>n/a</td>
<td>eg <code v-pre>f +f | g *.txt | !g murex.*</code> <br/> (returns only files with the extension &quot;txt&quot; that aren't called &quot;murex&quot;)</td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Expressions</strong></td>
<td></td>
<td></td>
</tr>
<tr>
<td>Assignment</td>
<td><code v-pre>foobar = $((1 + 2 * 3))</code></td>
<td><code v-pre>foobar = 1 + 2 * 3</code> <a href="#footnotes">[2]</a></td>
</tr>
<tr>
<td>Comparison, string</td>
<td><code v-pre>[ &quot;$(command)&quot; == &quot;value&quot; ]</code></td>
<td><code v-pre>${command} == &quot;value&quot;</code> <a href="#footnotes">[2]</a> <a href="#footnotes">[5]</a></td>
</tr>
<tr>
<td>Comparison, numeric</td>
<td><code v-pre>[ $integer -eq 5 ]</code></td>
<td><code v-pre>$number == 5</code> <a href="#footnotes">[2]</a></td>
</tr>
<tr>
<td>Arithmetic</td>
<td><code v-pre>echo $(( 1+2*3 ))</code></td>
<td><code v-pre>1 + 2 * 3</code> <a href="#footnotes">[2]</a> <br/><br/> <code v-pre>out ${1+2*3}</code> <a href="#footnotes">[2]</a> <a href="#footnotes">[5]</a></td>
</tr>
<tr>
<td>Supported data types</td>
<td>1. String,<br/>2. Integer<br/>(all variables are strings)</td>
<td>1. String,<br/>2. Integer,<br/>3. Float (default number type),<br/>4. Boolean<br/>5. Array,<br/>6. Object,<br/>7. Null<br/>(all variables can be treated as strings and/or their primitive)</td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Variables</strong><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><RouterLink to="/commands/set.html">Assign a local variable</RouterLink></td>
<td><code v-pre>local foo=&quot;bar&quot;</code></td>
<td><code v-pre>$foo = &quot;bar&quot;</code> <a href="#footnotes">[2]</a> <a href="#footnotes">[6]</a><br/><br/><code v-pre>out &quot;bar&quot; | set foo</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/global.html">Assign a global variable</RouterLink></td>
<td><code v-pre>foo=&quot;bar&quot;</code></td>
<td><code v-pre>$GLOBAL.foo = &quot;bar&quot;</code> <a href="#footnotes">[6]</a><br/><br/><code v-pre>out &quot;bar&quot; | global foo</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/export.html">Assign an environmental variable</RouterLink></td>
<td><code v-pre>export foo=&quot;bar&quot;</code></td>
<td><code v-pre>export foo = &quot;bar&quot;</code> <a href="#footnotes">[1]</a> <a href="#footnotes">[2]</a> <a href="#footnotes">[3]</a><br/><br/><code v-pre>$ENV.foo = &quot;bar&quot;</code> <a href="#footnotes">[6]</a><br/><br/><code v-pre>out &quot;bar&quot; | export foo</code> <a href="#footnotes">[3]</a></td>
</tr>
<tr>
<td><RouterLink to="/parser/string.html">Printing a variable</RouterLink></td>
<td><code v-pre>echo &quot;$foobar&quot;</code></td>
<td><code v-pre>out $foobar</code> <a href="#footnotes">[5]</a><br/><br/><code v-pre>$foobar</code> <br/><br/> (variables don't need to be quoted in Murex)</td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Arrays</strong><br/>(eg arrays, lists)</td>
<td></td>
<td></td>
</tr>
<tr>
<td>Creating an array</td>
<td><code v-pre>array_name=(value1 value2 value3)</code></td>
<td><code v-pre>%[value1 value2 value3]</code> <br/><br/><code v-pre>%[value1, value2, value3]</code> <br/><br/> eg <code v-pre>array_name = %[1, 2, 3]</code>, <br/> eg <code v-pre>%[hello world] | foreach { ... }</code></td>
</tr>
<tr>
<td>Accessing an array element</td>
<td><code v-pre>${array_name[0]}</code></td>
<td><code v-pre>$array_name[0]</code> (immutable) <br/><br/><code v-pre>$array_name.0</code> (mutable) <a href="#footnotes">[5]</a> <br/><br/> <code v-pre>array | [0]</code></td>
</tr>
<tr>
<td>Printing multiple elements</td>
<td><code v-pre>echo ${array_name[1]} ${array_name[0]}</code></td>
<td><code v-pre>@array_name[1 0]</code> <br/><br/> <code v-pre>array | [1 0]</code></td>
</tr>
<tr>
<td>Printing a range of elements</td>
<td>n/a</td>
<td><code v-pre>@array_name[1..3]</code> <br/><br/><code v-pre>array | [1..3]</code></td>
</tr>
<tr>
<td><RouterLink to="/parser/array.html">Printing all elements</RouterLink></td>
<td><code v-pre>echo ${array_name[*]}</code></td>
<td><code v-pre>@array_name</code></td>
</tr>
<tr>
<td><RouterLink to="/commands/foreach.html">Iterating through an array</RouterLink></td>
<td><code v-pre>for item in array; do;</code><br/>    <code v-pre>$item</code><br/><code v-pre>done;</code></td>
<td><code v-pre>array | foreach item { $item }</code> <br/><br/> eg <code v-pre>%[Tom Richard Sally] | foreach name { out &quot;Hello $name&quot; }</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Objects</strong><br/>(eg JSON objects, maps, hashes, dictionaries)</td>
<td></td>
<td></td>
</tr>
<tr>
<td>Creating an object</td>
<td>n/a</td>
<td><code v-pre>%{ key: value, array: [1, 2, 3] }</code> <a href="#footnotes">[2]</a> <br/><br/> eg <code v-pre>object_name = %{ key: val, arr: [1,3,3] }</code> <br/> eg <code v-pre>%{ a:1, b:2, c:3 } | formap { ... }</code></td>
</tr>
<tr>
<td>Accessing an element</td>
<td>n/a</td>
<td><code v-pre>$object_name[key]</code> (immutable) <br/><br/> <code v-pre>$object_name.key</code> <a href="#footnotes">[5]</a> (mutable) <br/><br/> <code v-pre>object | [key]</code></td>
</tr>
<tr>
<td>Printing multiple elements</td>
<td>n/a</td>
<td><code v-pre>$object_name[key1 key2]</code> <br/><br/> <code v-pre>object | [key1 key2]</code></td>
</tr>
<tr>
<td>Accessing a nested element</td>
<td>n/a</td>
<td><code v-pre>$object_name[[.path.to.element]]</code> (immutable) <a href="#footnotes">[4]</a><br/><br/> <code v-pre>$object_name.path.to.element</code> (mutable)<br/><br/> <code v-pre>object | [[.path.to.element]]</code> <a href="#footnotes">[4]</a><br/><br/></td>
</tr>
<tr>
<td><RouterLink to="/commands/formap.html">Iterating through an map</RouterLink></td>
<td>n/a</td>
<td><code v-pre>object | formap key value { $key; $value }</code> <br/><br/> eg <code v-pre>%{Bob: {age: 10}, Richard: {age: 20}, Sally: {age: 30} } | formap name person { out &quot;$name is $person[age] years old&quot; }</code></td>
</tr>
<tr>
<td><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td><strong>Sub-shells</strong><br/></td>
<td></td>
<td></td>
</tr>
<tr>
<td>Sub-shell, string</td>
<td><code v-pre>&quot;$(commands)&quot;</code> <br/><br/> eg <code v-pre>&quot;echo $(echo &quot;Hello world&quot;)&quot;</code></td>
<td><code v-pre>${commands}</code> <a href="#footnotes">[5]</a> <br/><br/> eg <code v-pre>out ${out Hello world}</code></td>
</tr>
<tr>
<td>Sub-shell, arrays</td>
<td><code v-pre>$(commands)</code> <br/><br/> eg <code v-pre>$(echo 1 2 3)</code></td>
<td><code v-pre>@{commands}</code> <a href="#footnotes">[5]</a> <br/><br/> eg <code v-pre>out @{ %[1,2,3] }</code></td>
</tr>
</tbody>
</table>
<h3 id="footnotes" tabindex="-1"><a class="header-anchor" href="#footnotes" aria-hidden="true">#</a> Footnotes</h3>
<ol>
<li>Supported for compatibility with traditional shells like Bash.</li>
<li>Unlike Bash, whitespace (or the absence of) is optional.</li>
<li>Environmental variables can only be stored as a string. This is a limitation of current operating systems.</li>
<li>Path separator can be any 1 byte wide character, eg <code v-pre>/</code>. The path separator is defined by the first character in a path.</li>
<li>Murex uses <code v-pre>${}</code> for subshells and <code v-pre>$()</code> for variables, the reverse of what Bash and others use. The reason for this difference is because <code v-pre>{}</code> always denotes a code block and <code v-pre>()</code> denotes strings. So <code v-pre>${foobar}</code> makes more sense as a subshell executing the command <code v-pre>foobar</code>, while <code v-pre>$(foobar)</code> makes more sense as the variable <code v-pre>$foobar</code>.</li>
<li>When assigning a variable where the right hand side is an expression, eg <code v-pre>$foo = &quot;bar&quot;</code>, the dollar prefix is optional. The <code v-pre>set</code>, <code v-pre>global</code> and <code v-pre>export</code> keywords are considered deprecated.</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/logical-and.html">And (<code v-pre>&amp;&amp;</code>) Logical Operator</RouterLink>:
Continues next operation if previous operation passes</li>
<li><RouterLink to="/parser/pipe-append.html">Append Pipe (<code v-pre>&gt;&gt;</code>) Token</RouterLink>:
Redirects STDOUT to a file and append its contents</li>
<li><RouterLink to="/parser/array.html">Array (<code v-pre>@</code>) Token</RouterLink>:
Expand values as an array</li>
<li><RouterLink to="/user-guide/namedpipes.html">Murex Named Pipes</RouterLink>:
A detailed breakdown of named pipes in Murex</li>
<li><RouterLink to="/parser/logical-or.html">Or (<code v-pre>||</code>) Logical Operator</RouterLink>:
Continues next operation only if previous operation fails</li>
<li><RouterLink to="/parser/string.html">String (<code v-pre>$</code>) Token</RouterLink>:
Expand values as a string</li>
<li><RouterLink to="/user-guide/terminal-keys.html">Terminal Hotkeys</RouterLink>:
A list of all the terminal hotkeys and their uses</li>
<li><RouterLink to="/commands/greater-than-greater-than.html"><code v-pre>&gt;&gt;</code> (append file)</RouterLink>:
Writes STDIN to disk - appending contents if file already exists</li>
<li><RouterLink to="/commands/greater-than.html"><code v-pre>&gt;</code> (truncate file)</RouterLink>:
Writes STDIN to disk - overwriting contents if file already exists</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/runmode.html"><code v-pre>runmode</code></RouterLink>:
Alter the scheduler's behaviour at higher scoping level</li>
<li><RouterLink to="/commands/try.html"><code v-pre>try</code></RouterLink>:
Handles errors inside a block of code</li>
<li><RouterLink to="/commands/trypipe.html"><code v-pre>trypipe</code></RouterLink>:
Checks state of each function in a pipeline and exits block on error</li>
</ul>
</div></template>


