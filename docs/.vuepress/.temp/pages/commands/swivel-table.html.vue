<template><div><h1 id="swivel-table" tabindex="-1"><a class="header-anchor" href="#swivel-table" aria-hidden="true">#</a> <code v-pre>swivel-table</code></h1>
<blockquote>
<p>Rotates a table by 90 degrees</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>swivel-table</code> rotates a table by 90 degrees so the <em>x</em> axis becomes the <em>y</em>.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>&lt;stdin&gt; -&gt; swivel-table -&gt; &lt;stdout&gt;
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.1 233996  8736 ?        Ss   Feb19   0:02 /sbin/init
root         2  0.0  0.0      0     0 ?        S    Feb19   0:00 [kthreadd]
root         4  0.0  0.0      0     0 ?        I&lt;   Feb19   0:00 [kworker/0:0H]
root         6  0.0  0.0      0     0 ?        I&lt;   Feb19   0:00 [mm_percpu_wq]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>That data swivelled would look like the following:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5 -> swivel-table
0       USER    root    root    root    root
1       PID     1       2       4       6
2       %CPU    0.0     0.0     0.0     0.0
3       %MEM    0.1     0.0     0.0     0.0
4       VSZ     233996  0       0       0
5       RSS     8736    0       0       0
6       TTY     ?       ?       ?       ?
7       STAT    Ss      S       I&lt;      I&lt;
8       START   Feb19   Feb19   Feb19   Feb19
9       TIME    0:02    0:00    0:00    0:00
10      COMMAND /sbin/init      [kthreadd]      [kworker/0:0H]  [mm_percpu_wq]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please note that column one is numbered because by default Murex couldn't
guess whether the first line of generic output is a title or data. However if we
format that as a CSV, which by default does have a title row (configurable via
<code v-pre>config</code>), then you would see titles as column one:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps: aux -> head: -n5 -> format: csv
"USER","PID","%CPU","%MEM","VSZ","RSS","TTY","STAT","START","TIME","COMMAND"
"root","1","0.0","0.1","233996","8736","?","Ss","Feb19","0:02","/sbin/init"
"root","2","0.0","0.0","0","0","?","S","Feb19","0:00","[kthreadd]"
"root","4","0.0","0.0","0","0","?","I&lt;","Feb19","0:00","[kworker/0:0H]"
"root","6","0.0","0.0","0","0","?","I&lt;","Feb19","0:00","[mm_percpu_wq]"

» ps: aux -> head: -n5 -> format: csv -> swivel-table
"USER","root","root","root","root"
"PID","1","2","4","6"
"%CPU","0.0","0.0","0.0","0.0"
"%MEM","0.1","0.0","0.0","0.0"
"VSZ","233996","0","0","0"
"RSS","8736","0","0","0"
"TTY","?","?","?","?"
"STAT","Ss","S","I&lt;","I&lt;"
"START","Feb19","Feb19","Feb19","Feb19"
"TIME","0:02","0:00","0:00","0:00"
"COMMAND","/sbin/init","[kthreadd]","[kworker/0:0H]","[mm_percpu_wq]"
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
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
<li><RouterLink to="/commands/swivel-datatype.html">commands/<code v-pre>swivel-datatype</code></RouterLink>:
Converts tabulated data into a map of values for serialised data-types such as JSON and YAML</li>
</ul>
</div></template>


