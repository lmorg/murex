<template><div><h1 id="language-guide-arrays-and-maps" tabindex="-1"><a class="header-anchor" href="#language-guide-arrays-and-maps" aria-hidden="true">#</a> Language Guide: Arrays And Maps</h1>
<h2 id="working-with-structured-data" tabindex="-1"><a class="header-anchor" href="#working-with-structured-data" aria-hidden="true">#</a> Working with structured data</h2>
<p>Firstly this shell doesn't have support for arrays as a native data type
however since Murex is aware of the structure of various data formats
it is possible to use these formats to maintain complex structured data
natively within Murex. For example a <code v-pre>days.json</code> file might look like</p>
<pre><code>[
        &quot;monday&quot;,
        &quot;tuesday&quot;,
        &quot;wednesday&quot;,
        &quot;thursday&quot;,
        &quot;friday&quot;,
        &quot;saturday&quot;,
        &quot;sunday&quot;
]
</code></pre>
<p>...which can be queried directly within Murex via a variety of builtins.</p>
<p>To iterate through the array and print each element and print the value:</p>
<pre><code>» open: days.json -&gt; foreach: day { $day }

monday
tuesday
wednesday
thursday
friday
saturday
sunday
</code></pre>
<p>To iterate through the map or array and print each index and its value:</p>
<pre><code>» open: days.json -&gt; formap: key value { echo: &quot;$key: $value&quot; }

0: &quot;monday&quot;
1: &quot;tuesday&quot;
2: &quot;wednesday&quot;
3: &quot;thursday&quot;
4: &quot;friday&quot;
5: &quot;saturday&quot;
6: &quot;sunday&quot;
</code></pre>
<p>To return a specific element within an array or map you can query it
directly by its key using the <code v-pre>index</code> builtin:</p>
<pre><code>» open: days.json -&gt; [ 0 ]

monday
</code></pre>
<p>Or multiple elements in the data set:</p>
<pre><code>» open: days.json -&gt; [ 0 2 5 6 ]

[&quot;monday&quot;,&quot;wednesday&quot;,&quot;saturday&quot;,&quot;sunday&quot;]
</code></pre>
<p>The <code v-pre>index</code> builtin returned the values in JSON format because the input
format was JSON. If the input format was a CSV then it would return the
selected columns of that CSV. Or if it's just a new line separated list
of strings then it would return a the rows in the list.</p>
<h2 id="the-array-builtin" tabindex="-1"><a class="header-anchor" href="#the-array-builtin" aria-hidden="true">#</a> The <code v-pre>array</code> builtin</h2>
<p>Murex has a pretty sophisticated builtin for generating arrays. Think
like bash's <code v-pre>{1..9}</code> syntax:</p>
<pre><code>a: [1..9]
</code></pre>
<p>You can also specify an alternative number base by using an <code v-pre>x</code> or <code v-pre>.</code>
in the end range:</p>
<pre><code>a: [00..ffx16]
a: [00..ff.16]
</code></pre>
<p>All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported.
Please note that the start and end range are written in the target base
while the base identifier is written in decimal: <code v-pre>[hex..hex.dec]</code></p>
<p>Also note that the additional zeros denotes padding (ie the results will
start at <code v-pre>00</code>, <code v-pre>01</code>, etc rather than <code v-pre>0</code>, <code v-pre>1</code>...</p>
<h3 id="character-arrays" tabindex="-1"><a class="header-anchor" href="#character-arrays" aria-hidden="true">#</a> Character arrays</h3>
<p>You can select a range of letters (a to z):</p>
<pre><code>a: [a..z]
a: [z..a]
a: [A..Z]
a: [Z..A]
</code></pre>
<p>...or any characters within that range.</p>
<h3 id="special-ranges" tabindex="-1"><a class="header-anchor" href="#special-ranges" aria-hidden="true">#</a> Special ranges</h3>
<p>Unlike bash, Murex also supports some special ranges:</p>
<pre><code>a: [mon..sun]
a: [monday..sunday]
a: [jan..dec]
a: [janurary..december]
a: [spring..winter]
</code></pre>
<p>It is also case aware. If the ranges are uppercase then the return will
be uppercase. If the ranges are title case (capital first letter) then
the return will be in title case:</p>
<pre><code>» a: [Monday..Sunday]

Monday
Tuesday
Wednesday
Thursday
Friday
Saturday
Sunday
</code></pre>
<p>Where the special ranges differ from a regular range is they cannot
cannot down. eg <code v-pre>a: [3..1]</code> would output</p>
<pre><code>3
2
1
</code></pre>
<p>however a negative range in special ranges will cycle through to the end
of the range and then loop back from the start:</p>
<pre><code>» a: [Thursday..Wednesday]

Thursday
Friday
Saturday
Sunday
Monday
Tuesday
Wednesday
</code></pre>
<p>This decision was made because generally with ranges of this type, you
would more often prefer to cycle through values rather than iterate
backwards through the list.</p>
<p>If you did want to reverse then just pipe the output into another UNIX
tool:</p>
<pre><code>» a: [Monday..Friday] -&gt; tac                         # Linux
» a: [Monday..Friday] -&gt; tail -r                     # BSD / OS X
» a: [Monday..Friday] -&gt; perl -e &quot;print reverse &lt;&gt;&quot;  # Multiplaform

Friday
Thurday
Wednesday
Tuesday
Monday
</code></pre>
<p>(I may build a reverse builtin to standardise this and make Murex more
accessible to Windows users)</p>
<h3 id="advanced-array-syntax" tabindex="-1"><a class="header-anchor" href="#advanced-array-syntax" aria-hidden="true">#</a> Advanced <code v-pre>array</code> syntax</h3>
<p>The syntax for <code v-pre>array</code> is a comma separated list of parameters with
expansions stored in square brackets. You can have an expansion embedded
inside a parameter or as it's own parameter. Expansions can also have
multiple parameters.</p>
<pre><code>» a: 01,02,03,05,06,07

01
02
03
05
06
07

» a: 0[1..3],0[5..7]

01
02
03
05
06
07

» a: 0[1..3,5..7]

01
02
03
05
06
07

» a: b[o,i]b

bob
bib
</code></pre>
<p>You can also have multiple expansion blocks in a single parameter:</p>
<pre><code>» a: a[1..3]b[5..7]

a1b5
a1b6
a1b7
a2b5
a2b6
a2b7
a3b5
a3b6
a3b7
</code></pre>
<p><code v-pre>array</code> will cycle through each iteration of the last expansion, moving
itself backwards through the string; behaving like an normal counter:</p>
<pre><code>» ja: [0..2][0..9] -&gt; format: str &quot;,&quot;

00,01,02,03,04,05,06,07,08,09,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29
</code></pre>
<p>(<code v-pre>format</code> used here for readability)</p>
<h3 id="creating-json-arrays-with-ja" tabindex="-1"><a class="header-anchor" href="#creating-json-arrays-with-ja" aria-hidden="true">#</a> Creating JSON arrays with <code v-pre>ja</code></h3>
<p>As you can see from the previous examples, <code v-pre>a</code> returns the array as a
list of strings. This is so you can stream excessively long arrays, for
example every IPv4 address: <code v-pre>a: [0..254].[0..254].[0..254].[0..254]</code>
(this kind of array expansion would hang bash).</p>
<p>However if you needed a JSON string then you can use all the same syntax
as <code v-pre>a</code> but forgo the streaming capability:</p>
<pre><code>» ja: [Monday..Sunday]

[
        &quot;Monday&quot;,
        &quot;Tuesday&quot;,
        &quot;Wednesday&quot;,
        &quot;Thursday&quot;,
        &quot;Friday&quot;,
        &quot;Saturday&quot;,
        &quot;Sunday&quot;
]
</code></pre>
</div></template>


