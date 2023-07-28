import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as r,d as e,b as a,w as t,e as n,f as o}from"./app-45f7c304.js";const u={},c=o(`<h1 id="index" tabindex="-1"><a class="header-anchor" href="#index" aria-hidden="true">#</a> <code>[</code> (index)</h1><blockquote><p>Outputs an element from an array, map or table</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Outputs an element or multiple elements from an array, map or table.</p><p>Please note that indexes in Murex are counted from zero.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>&lt;stdin&gt; -&gt; [ element ] -&gt; &lt;stdout&gt;
$variable[ element ] -&gt; &lt;stdout&gt;

&lt;stdin&gt; -&gt; ![ element ] -&gt; &lt;stdout&gt;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Return the 2nd (1), 4th (3) and 6th (5) element in an array:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ja [0..9] -&gt; [ 1 3 5 ]
[
    &quot;1&quot;,
    &quot;3&quot;,
    &quot;5&quot;
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the data-type and description of <strong>config shell syntax-highlighting</strong>:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» config -&gt; [[ /shell/syntax-highlighting ]] -&gt; [ Data-Type Description ]
[
    &quot;bool&quot;,
    &quot;Syntax highlighting of murex code when in the interactive shell&quot;
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return all elements <em>except</em> for 1 (2nd), 3 (4th) and 5 (6th):</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» a: [0..9]-&gt; ![ 1 3 5 ]
0
2
4
6
7
8
9
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return all elements except for the data-type and description:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» config -&gt; [[ /shell/syntax-highlighting ]] -&gt; ![ Data-Type Description ]
{
    &quot;Default&quot;: true,
    &quot;Dynamic&quot;: false,
    &quot;Global&quot;: true,
    &quot;Value&quot;: true
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the top 5 processes from <code>ps</code>, ordered by memory usage:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps aux -&gt; [PID %MEM COMMAND] -&gt; sort -nrk2 -&gt; [..5]
915961  14.4  /home/lau/dev/go/bin/gopls
916184  4.4   /opt/visual-studio-code/code
108025  2.9   /usr/lib/firefox/firefox
1036    2.4   /usr/lib/baloo_file
915710  1.9   /opt/visual-studio-code/code
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the 1st and 30th row:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps aux -&gt; [*1 *30]
USER    PID     %CPU    %MEM    VSZ     RSS     TTY     STAT    START   TIME    COMMAND
root    37      0.0     0.0     0       0       ?       I&lt;      Dec18   0:00    [kworker/3:0H-events_highpri]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Return the 1st and 5th column:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps aux -&gt; [*A *E] -&gt; head -n5
USER    VSZ
root    168284
root    0
root    0
root    0
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="index-counts-from-zero" tabindex="-1"><a class="header-anchor" href="#index-counts-from-zero" aria-hidden="true">#</a> Index counts from zero</h3><p>Indexes in Murex behave like any other computer array in that all arrays start from zero (<code>0</code>).</p><h3 id="include-vs-exclude" tabindex="-1"><a class="header-anchor" href="#include-vs-exclude" aria-hidden="true">#</a> Include vs exclude</h3><p>As demonstrated in the examples above, <code>[</code> specifies elements to include where as <code>![</code> specifies elements to exclude.</p><h3 id="don-t-error-upon-missing-elements" tabindex="-1"><a class="header-anchor" href="#don-t-error-upon-missing-elements" aria-hidden="true">#</a> Don&#39;t error upon missing elements</h3><p>By default, <strong>index</strong> generates an error if an element doesn&#39;t exist. However you can disable this behavior in <code>config</code></p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» config -&gt; [ foobar ]
Error in \`[\` ((builtin) 2,11): Key &#39;foobar&#39; not found

» config set index silent true

» config -&gt; [ foobar ]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>[</code></li><li><code>![</code></li><li><code>index</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,33),m=e("code",null,"[[",-1),v=e("code",null,"[",-1),h=e("code",null,"a",-1),b=e("code",null,"config",-1),g=e("code",null,"count",-1),p=e("code",null,"ja",-1),x=e("code",null,"mtac",-1);function f(_,y){const i=s("RouterLink");return l(),r("div",null,[c,e("ul",null,[e("li",null,[a(i,{to:"/commands/element.html"},{default:t(()=>[m,n(" (element)")]),_:1}),n(": Outputs an element from a nested structure")]),e("li",null,[a(i,{to:"/commands/range.html"},{default:t(()=>[v,n(" (range) ")]),_:1}),n(": Outputs a ranged subset of data from STDIN")]),e("li",null,[a(i,{to:"/commands/a.html"},{default:t(()=>[h,n(" (mkarray)")]),_:1}),n(": A sophisticated yet simple way to build an array or list")]),e("li",null,[a(i,{to:"/commands/config.html"},{default:t(()=>[b]),_:1}),n(": Query or define Murex runtime settings")]),e("li",null,[a(i,{to:"/commands/count.html"},{default:t(()=>[g]),_:1}),n(": Count items in a map, list or array")]),e("li",null,[a(i,{to:"/commands/ja.html"},{default:t(()=>[p,n(" (mkarray)")]),_:1}),n(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[a(i,{to:"/commands/mtac.html"},{default:t(()=>[x]),_:1}),n(": Reverse the order of an array")])])])}const R=d(u,[["render",f],["__file","index2.html.vue"]]);export{R as default};
