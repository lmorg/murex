<template><div><h1 id="swivel-datatype" tabindex="-1"><a class="header-anchor" href="#swivel-datatype" aria-hidden="true">#</a> <code v-pre>swivel-datatype</code></h1>
<blockquote>
<p>Converts tabulated data into a map of values for serialised data-types such as JSON and YAML</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>swivel-datatype</code> rotates a table by 90 degrees then exports the output as a
series of maps to be marshalled by a serialised datatype such as JSON or YAML.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; swivel-datatype: data-type -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Lets take the first 5 entries from <code v-pre>ps</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5 -> format: csv
"USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND"
"root","1","0.0","0.1","233996","8736","?","Ss","Feb19","0:02","/sbin/init"
"root","2","0.0","0.0","0","0","?","S","Feb19","0:00","[kthreadd]"
"root","4","0.0","0.0","0","0","?","I&lt;","Feb19","0:00","[kworker/0:0H]"
"root","6","0.0","0.0","0","0","?","I&lt;","Feb19","0:00","[mm_percpu_wq]"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>That data swivelled would look like the following:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5 -> format: csv -> swivel-datatype: yaml
'%CPU':
- "0.0"
- "0.0"
- "0.0"
- "0.0"
'%MEM':
- "0.1"
- "0.0"
- "0.0"
- "0.0"
COMMAND:
- /sbin/init
- '[kthreadd]'
- '[kworker/0:0H]'
- '[mm_percpu_wq]'
PID:
- "1"
- "2"
- "4"
- "6"
RSS:
- "8736"
- "0"
- "0"
- "0"
START:
- Feb19
- Feb19
- Feb19
- Feb19
STAT:
- Ss
- S
- I&lt;
- I&lt;
TIME:
- "0:02"
- "0:00"
- "0:00"
- "0:00"
TTY:
- '?'
- '?'
- '?'
- '?'
USER:
- root
- root
- root
- root
VSZ:
- "233996"
- "0"
- "0"
- "0"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please note that for input data-types whose table doesn't define titles (such as
the <code v-pre>generic</code> datatype), the map keys are defaulted to column numbers:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5 -> swivel-datatype: yaml
"0":
- USER
- root
- root
- root
- root
"1":
- PID
- "1"
- "2"
- "4"
- "6"
"2":
- '%CPU'
- "0.0"
- "0.0"
- "0.0"
- "0.0"
"3":
- '%MEM'
- "0.1"
- "0.0"
- "0.0"
- "0.0"
...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>You can check what output data-types are available via the <code v-pre>runtime</code> command:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>runtime --marshallers
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Marshallers are enabled at compile time from the <code v-pre>builtins/data-types</code> directory.</p>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html">commands/<code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html">commands/<code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/alter.html">commands/<code v-pre>alter</code></RouterLink>:
Change a value within a structured data-type and pass that change along the pipeline without altering the original source input</li>
<li><RouterLink to="/commands/append.html">commands/<code v-pre>append</code></RouterLink>:
Add data to the end of an array</li>
<li><RouterLink to="/commands/cast.html">commands/<code v-pre>cast</code></RouterLink>:
Alters the data type of the previous function without altering it's output</li>
<li><RouterLink to="/commands/format.html">commands/<code v-pre>format</code></RouterLink>:
Reformat one data-type into another data-type</li>
<li><RouterLink to="/commands/prepend.html">commands/<code v-pre>prepend</code> </RouterLink>:
Add data to the start of an array</li>
<li><RouterLink to="/commands/runtime.html">commands/<code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
<li><RouterLink to="/commands/swivel-table.html">commands/<code v-pre>swivel-table</code></RouterLink>:
Rotates a table by 90 degrees</li>
</ul>
</div></template>


