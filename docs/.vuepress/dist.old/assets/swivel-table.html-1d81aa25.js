import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as l,c as d,d as o,b as u,w as a,e as t,f as s}from"./app-45f7c304.js";const q={},r=s(`<h1 id="swivel-table" tabindex="-1"><a class="header-anchor" href="#swivel-table" aria-hidden="true">#</a> <code>swivel-table</code></h1><blockquote><p>Rotates a table by 90 degrees</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>swivel-table</code> rotates a table by 90 degrees so the <em>x</em> axis becomes the <em>y</em>.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>&lt;stdin&gt; -&gt; swivel-table -&gt; &lt;stdout&gt;
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5
USER       PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root         1  0.0  0.1 233996  8736 ?        Ss   Feb19   0:02 /sbin/init
root         2  0.0  0.0      0     0 ?        S    Feb19   0:00 [kthreadd]
root         4  0.0  0.0      0     0 ?        I&lt;   Feb19   0:00 [kworker/0:0H]
root         6  0.0  0.0      0     0 ?        I&lt;   Feb19   0:00 [mm_percpu_wq]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>That data swivelled would look like the following:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5 -&gt; swivel-table
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
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please note that column one is numbered because by default Murex couldn&#39;t guess whether the first line of generic output is a title or data. However if we format that as a CSV, which by default does have a title row (configurable via <code>config</code>), then you would see titles as column one:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5 -&gt; format: csv
&quot;USER&quot;,&quot;PID&quot;,&quot;%CPU&quot;,&quot;%MEM&quot;,&quot;VSZ&quot;,&quot;RSS&quot;,&quot;TTY&quot;,&quot;STAT&quot;,&quot;START&quot;,&quot;TIME&quot;,&quot;COMMAND&quot;
&quot;root&quot;,&quot;1&quot;,&quot;0.0&quot;,&quot;0.1&quot;,&quot;233996&quot;,&quot;8736&quot;,&quot;?&quot;,&quot;Ss&quot;,&quot;Feb19&quot;,&quot;0:02&quot;,&quot;/sbin/init&quot;
&quot;root&quot;,&quot;2&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;S&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[kthreadd]&quot;
&quot;root&quot;,&quot;4&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;I&lt;&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[kworker/0:0H]&quot;
&quot;root&quot;,&quot;6&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;I&lt;&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[mm_percpu_wq]&quot;

» ps: aux -&gt; head: -n5 -&gt; format: csv -&gt; swivel-table
&quot;USER&quot;,&quot;root&quot;,&quot;root&quot;,&quot;root&quot;,&quot;root&quot;
&quot;PID&quot;,&quot;1&quot;,&quot;2&quot;,&quot;4&quot;,&quot;6&quot;
&quot;%CPU&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0.0&quot;
&quot;%MEM&quot;,&quot;0.1&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0.0&quot;
&quot;VSZ&quot;,&quot;233996&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;
&quot;RSS&quot;,&quot;8736&quot;,&quot;0&quot;,&quot;0&quot;,&quot;0&quot;
&quot;TTY&quot;,&quot;?&quot;,&quot;?&quot;,&quot;?&quot;,&quot;?&quot;
&quot;STAT&quot;,&quot;Ss&quot;,&quot;S&quot;,&quot;I&lt;&quot;,&quot;I&lt;&quot;
&quot;START&quot;,&quot;Feb19&quot;,&quot;Feb19&quot;,&quot;Feb19&quot;,&quot;Feb19&quot;
&quot;TIME&quot;,&quot;0:02&quot;,&quot;0:00&quot;,&quot;0:00&quot;,&quot;0:00&quot;
&quot;COMMAND&quot;,&quot;/sbin/init&quot;,&quot;[kthreadd]&quot;,&quot;[kworker/0:0H]&quot;,&quot;[mm_percpu_wq]&quot;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,13),c=o("code",null,"[[",-1),m=o("code",null,"[",-1),v=o("code",null,"alter",-1),b=o("code",null,"append",-1),h=o("code",null,"cast",-1),p=o("code",null,"format",-1),_=o("code",null,"prepend",-1),f=o("code",null,"swivel-datatype",-1);function g(S,x){const e=i("RouterLink");return l(),d("div",null,[r,o("ul",null,[o("li",null,[u(e,{to:"/commands/element.html"},{default:a(()=>[t("commands/"),c,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),o("li",null,[u(e,{to:"/commands/index2.html"},{default:a(()=>[t("commands/"),m,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),o("li",null,[u(e,{to:"/commands/alter.html"},{default:a(()=>[t("commands/"),v]),_:1}),t(": Change a value within a structured data-type and pass that change along the pipeline without altering the original source input")]),o("li",null,[u(e,{to:"/commands/append.html"},{default:a(()=>[t("commands/"),b]),_:1}),t(": Add data to the end of an array")]),o("li",null,[u(e,{to:"/commands/cast.html"},{default:a(()=>[t("commands/"),h]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),o("li",null,[u(e,{to:"/commands/format.html"},{default:a(()=>[t("commands/"),p]),_:1}),t(": Reformat one data-type into another data-type")]),o("li",null,[u(e,{to:"/commands/prepend.html"},{default:a(()=>[t("commands/"),_]),_:1}),t(": Add data to the start of an array")]),o("li",null,[u(e,{to:"/commands/swivel-datatype.html"},{default:a(()=>[t("commands/"),f]),_:1}),t(": Converts tabulated data into a map of values for serialised data-types such as JSON and YAML")])])])}const M=n(q,[["render",g],["__file","swivel-table.html.vue"]]);export{M as default};
