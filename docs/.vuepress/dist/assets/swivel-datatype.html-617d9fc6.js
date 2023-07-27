import{_ as u}from"./plugin-vue_export-helper-c27b6911.js";import{r as a,o as d,c as s,d as t,b as i,w as o,e,f as l}from"./app-45f7c304.js";const r={},c=l(`<h1 id="swivel-datatype" tabindex="-1"><a class="header-anchor" href="#swivel-datatype" aria-hidden="true">#</a> <code>swivel-datatype</code></h1><blockquote><p>Converts tabulated data into a map of values for serialised data-types such as JSON and YAML</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>swivel-datatype</code> rotates a table by 90 degrees then exports the output as a series of maps to be marshalled by a serialised datatype such as JSON or YAML.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>&lt;stdin&gt; -&gt; swivel-datatype: data-type -&gt; &lt;stdout&gt;
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Lets take the first 5 entries from <code>ps</code>:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5 -&gt; format: csv
&quot;USER&quot;,&quot;PID&quot;,&quot;%CPU&quot;,&quot;%MEM&quot;,&quot;VSZ&quot;,&quot;RSS&quot;,&quot;TTY&quot;,&quot;STAT&quot;,&quot;START&quot;,&quot;TIME&quot;,&quot;COMMAND&quot;
&quot;root&quot;,&quot;1&quot;,&quot;0.0&quot;,&quot;0.1&quot;,&quot;233996&quot;,&quot;8736&quot;,&quot;?&quot;,&quot;Ss&quot;,&quot;Feb19&quot;,&quot;0:02&quot;,&quot;/sbin/init&quot;
&quot;root&quot;,&quot;2&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;S&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[kthreadd]&quot;
&quot;root&quot;,&quot;4&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;I&lt;&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[kworker/0:0H]&quot;
&quot;root&quot;,&quot;6&quot;,&quot;0.0&quot;,&quot;0.0&quot;,&quot;0&quot;,&quot;0&quot;,&quot;?&quot;,&quot;I&lt;&quot;,&quot;Feb19&quot;,&quot;0:00&quot;,&quot;[mm_percpu_wq]&quot;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>That data swivelled would look like the following:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5 -&gt; format: csv -&gt; swivel-datatype: yaml
&#39;%CPU&#39;:
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
&#39;%MEM&#39;:
- &quot;0.1&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
COMMAND:
- /sbin/init
- &#39;[kthreadd]&#39;
- &#39;[kworker/0:0H]&#39;
- &#39;[mm_percpu_wq]&#39;
PID:
- &quot;1&quot;
- &quot;2&quot;
- &quot;4&quot;
- &quot;6&quot;
RSS:
- &quot;8736&quot;
- &quot;0&quot;
- &quot;0&quot;
- &quot;0&quot;
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
- &quot;0:02&quot;
- &quot;0:00&quot;
- &quot;0:00&quot;
- &quot;0:00&quot;
TTY:
- &#39;?&#39;
- &#39;?&#39;
- &#39;?&#39;
- &#39;?&#39;
USER:
- root
- root
- root
- root
VSZ:
- &quot;233996&quot;
- &quot;0&quot;
- &quot;0&quot;
- &quot;0&quot;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please note that for input data-types whose table doesn&#39;t define titles (such as the <code>generic</code> datatype), the map keys are defaulted to column numbers:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» ps: aux -&gt; head: -n5 -&gt; swivel-datatype: yaml
&quot;0&quot;:
- USER
- root
- root
- root
- root
&quot;1&quot;:
- PID
- &quot;1&quot;
- &quot;2&quot;
- &quot;4&quot;
- &quot;6&quot;
&quot;2&quot;:
- &#39;%CPU&#39;
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
&quot;3&quot;:
- &#39;%MEM&#39;
- &quot;0.1&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
- &quot;0.0&quot;
...
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>You can check what output data-types are available via the <code>runtime</code> command:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>runtime --marshallers
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><p>Marshallers are enabled at compile time from the <code>builtins/data-types</code> directory.</p><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,18),v=t("code",null,"[[",-1),m=t("code",null,"[",-1),q=t("code",null,"alter",-1),b=t("code",null,"append",-1),h=t("code",null,"cast",-1),p=t("code",null,"format",-1),f=t("code",null,"prepend",-1),_=t("code",null,"runtime",-1),g=t("code",null,"swivel-table",-1);function x(y,w){const n=a("RouterLink");return d(),s("div",null,[c,t("ul",null,[t("li",null,[i(n,{to:"/commands/element.html"},{default:o(()=>[e("commands/"),v,e(" (element)")]),_:1}),e(": Outputs an element from a nested structure")]),t("li",null,[i(n,{to:"/commands/index2.html"},{default:o(()=>[e("commands/"),m,e(" (index)")]),_:1}),e(": Outputs an element from an array, map or table")]),t("li",null,[i(n,{to:"/commands/alter.html"},{default:o(()=>[e("commands/"),q]),_:1}),e(": Change a value within a structured data-type and pass that change along the pipeline without altering the original source input")]),t("li",null,[i(n,{to:"/commands/append.html"},{default:o(()=>[e("commands/"),b]),_:1}),e(": Add data to the end of an array")]),t("li",null,[i(n,{to:"/commands/cast.html"},{default:o(()=>[e("commands/"),h]),_:1}),e(": Alters the data type of the previous function without altering it's output")]),t("li",null,[i(n,{to:"/commands/format.html"},{default:o(()=>[e("commands/"),p]),_:1}),e(": Reformat one data-type into another data-type")]),t("li",null,[i(n,{to:"/commands/prepend.html"},{default:o(()=>[e("commands/"),f]),_:1}),e(": Add data to the start of an array")]),t("li",null,[i(n,{to:"/commands/runtime.html"},{default:o(()=>[e("commands/"),_]),_:1}),e(": Returns runtime information on the internal state of Murex")]),t("li",null,[i(n,{to:"/commands/swivel-table.html"},{default:o(()=>[e("commands/"),g]),_:1}),e(": Rotates a table by 90 degrees")])])])}const M=u(r,[["render",x],["__file","swivel-datatype.html.vue"]]);export{M as default};
