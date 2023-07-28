import{_ as t}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as s,c as o,d as e,b as a,w as l,e as n,f as r}from"./app-45f7c304.js";const c={},u=r(`<h1 id="while" tabindex="-1"><a class="header-anchor" href="#while" aria-hidden="true">#</a> <code>while</code></h1><blockquote><p>Loop until condition false</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>while</code> loops until loops until <strong>condition</strong> is false.</p><p>Normally the <strong>conditional</strong> and executed code block are 2 separate parameters however you can call <code>while</code> with just 1 parameter where the code block acts as both the conditional and the code to be ran.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><p>Until true</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>while { condition } { code-block } -&gt; &lt;stdout&gt;

while { code-block } -&gt; &lt;stdout&gt;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Until false</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>!while { condition } { code-block } -&gt; &lt;stdout&gt;
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p><code>while</code> <strong>$i</strong> is less then <strong>5</strong></p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» let i=0; while { =i&lt;5 } { let i=i+1; out $i }
1
2
3
4
5

» let i=0; while { let i=i+1; = i&lt;5; out }
true
true
true
true
false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p><code>while</code> <strong>$i</strong> is <em>NOT</em> greater than or equal to <strong>5</strong></p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» let i=0; !while { =i&gt;=5 } { let i=i+1; out $i }
1
2
3
4
5

» let i=0; while { let i=i+1; = i&gt;=5; out }
true
true
true
true
false
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="meta-values" tabindex="-1"><a class="header-anchor" href="#meta-values" aria-hidden="true">#</a> Meta values</h3><p>Meta values are a JSON object stored as the variable <code>$.</code>. The meta variable will get overwritten by any other block which invokes meta values. So if you wish to persist meta values across blocks you will need to reassign <code>$.</code>, eg</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>%[1..3] -&gt; foreach {
    meta_parent = $.
    %[7..9] -&gt; foreach {
        out &quot;$(meta_parent.i): $.i&quot;
    }
}
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The following meta values are defined:</p><ul><li><code>i</code>: iteration number</li></ul><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>while</code></li><li><code>!while</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,24),h=e("code",null,"err",-1),v=e("code",null,"for",-1),m=e("code",null,"foreach",-1),b=e("code",null,"formap",-1),g=e("code",null,"global",-1),p=e("code",null,"let",-1),f=e("code",null,"out",-1),_=e("code",null,"set",-1);function x(w,k){const i=d("RouterLink");return s(),o("div",null,[u,e("ul",null,[e("li",null,[a(i,{to:"/commands/err.html"},{default:l(()=>[h]),_:1}),n(": Print a line to the STDERR")]),e("li",null,[a(i,{to:"/commands/for.html"},{default:l(()=>[v]),_:1}),n(": A more familiar iteration loop to existing developers")]),e("li",null,[a(i,{to:"/commands/foreach.html"},{default:l(()=>[m]),_:1}),n(": Iterate through an array")]),e("li",null,[a(i,{to:"/commands/formap.html"},{default:l(()=>[b]),_:1}),n(": Iterate through a map or other collection of data")]),e("li",null,[a(i,{to:"/commands/global.html"},{default:l(()=>[g]),_:1}),n(": Define a global variable and set it's value")]),e("li",null,[a(i,{to:"/commands/let.html"},{default:l(()=>[p]),_:1}),n(": Evaluate a mathematical function and assign to variable (deprecated)")]),e("li",null,[a(i,{to:"/commands/out.html"},{default:l(()=>[f]),_:1}),n(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[a(i,{to:"/commands/set.html"},{default:l(()=>[_]),_:1}),n(": Define a local variable and set it's value")])])])}const N=t(c,[["render",x],["__file","while.html.vue"]]);export{N as default};
