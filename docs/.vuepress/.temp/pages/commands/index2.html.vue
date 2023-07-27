<template><div><h1 id="index" tabindex="-1"><a class="header-anchor" href="#index" aria-hidden="true">#</a> <code v-pre>[</code> (index)</h1>
<blockquote>
<p>Outputs an element from an array, map or table</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Outputs an element or multiple elements from an array, map or table.</p>
<p>Please note that indexes in Murex are counted from zero.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> [ element ] -> &lt;stdout>
$variable[ element ] -> &lt;stdout>

&lt;stdin> -> ![ element ] -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Return the 2nd (1), 4th (3) and 6th (5) element in an array:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ja [0..9] -> [ 1 3 5 ]
[
    "1",
    "3",
    "5"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the data-type and description of <strong>config shell syntax-highlighting</strong>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config -> [[ /shell/syntax-highlighting ]] -> [ Data-Type Description ]
[
    "bool",
    "Syntax highlighting of murex code when in the interactive shell"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return all elements <em>except</em> for 1 (2nd), 3 (4th) and 5 (6th):</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: [0..9]-> ![ 1 3 5 ]
0
2
4
6
7
8
9
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return all elements except for the data-type and description:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config -> [[ /shell/syntax-highlighting ]] -> ![ Data-Type Description ]
{
    "Default": true,
    "Dynamic": false,
    "Global": true,
    "Value": true
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the top 5 processes from <code v-pre>ps</code>, ordered by memory usage:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps aux -> [PID %MEM COMMAND] -> sort -nrk2 -> [..5]
915961  14.4  /home/lau/dev/go/bin/gopls
916184  4.4   /opt/visual-studio-code/code
108025  2.9   /usr/lib/firefox/firefox
1036    2.4   /usr/lib/baloo_file
915710  1.9   /opt/visual-studio-code/code
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the 1st and 30th row:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps aux -> [*1 *30]
USER    PID     %CPU    %MEM    VSZ     RSS     TTY     STAT    START   TIME    COMMAND
root    37      0.0     0.0     0       0       ?       I&lt;      Dec18   0:00    [kworker/3:0H-events_highpri]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the 1st and 5th column:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» ps aux -> [*A *E] -> head -n5
USER    VSZ
root    168284
root    0
root    0
root    0
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="index-counts-from-zero" tabindex="-1"><a class="header-anchor" href="#index-counts-from-zero" aria-hidden="true">#</a> Index counts from zero</h3>
<p>Indexes in Murex behave like any other computer array in that all arrays
start from zero (<code v-pre>0</code>).</p>
<h3 id="include-vs-exclude" tabindex="-1"><a class="header-anchor" href="#include-vs-exclude" aria-hidden="true">#</a> Include vs exclude</h3>
<p>As demonstrated in the examples above, <code v-pre>[</code> specifies elements to include
where as <code v-pre>![</code> specifies elements to exclude.</p>
<h3 id="don-t-error-upon-missing-elements" tabindex="-1"><a class="header-anchor" href="#don-t-error-upon-missing-elements" aria-hidden="true">#</a> Don't error upon missing elements</h3>
<p>By default, <strong>index</strong> generates an error if an element doesn't exist. However
you can disable this behavior in <code v-pre>config</code></p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» config -> [ foobar ]
Error in `[` ((builtin) 2,11): Key 'foobar' not found

» config set index silent true

» config -> [ foobar ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2>
<ul>
<li><code v-pre>[</code></li>
<li><code v-pre>![</code></li>
<li><code v-pre>index</code></li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/config.html"><code v-pre>config</code></RouterLink>:
Query or define Murex runtime settings</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
</ul>
</div></template>


