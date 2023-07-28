import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as s,c as l,d as e,e as a,b as n,w as r,f as i}from"./app-45f7c304.js";const c={},u=e("h1",{id:"decimal-ranges-mkarray",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#decimal-ranges-mkarray","aria-hidden":"true"},"#"),a(" Decimal Ranges - mkarray")],-1),h=e("blockquote",null,[e("p",null,"Create arrays of decimal integers")],-1),m=e("h2",{id:"description",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#description","aria-hidden":"true"},"#"),a(" Description")],-1),p=e("p",null,[a("This document describes how to create arrays of decimals using mkarray ("),e("code",null,"a"),a(" et al).")],-1),_=i(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end,start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end][start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>All usages also work with <code>ja</code> and <code>ta</code> as well, eg:</p><pre><code>ja: [start..end] -&gt; \`&lt;stdout&gt;\`
ta: data-type [start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>You can also inline arrays with the <code>%[]</code> syntax, eg:</p><pre><code>%[start..end]
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» a: [1..3]
1
2
3

» a: [3..1]
3
2
1

» a: [01..03]
01
02
03
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="floating-point-numbers" tabindex="-1"><a class="header-anchor" href="#floating-point-numbers" aria-hidden="true">#</a> Floating Point Numbers</h3><p>If you do need a range of fixed floating point numbers generated then you can do so by merging two decimal integer ranges together. For example</p><pre><code>» a [0..5].[0..9]
0.0
0.1
0.2
0.3
0.4
0.5
0.6
0.7
0.8
0.9
1.0
1.1
1.2
1.3
...
4.8
4.9
5.0
5.1
5.2
5.3
5.4
5.5
5.6
5.7
5.8
5.9
</code></pre><h3 id="everything-is-a-string" tabindex="-1"><a class="header-anchor" href="#everything-is-a-string" aria-hidden="true">#</a> Everything Is A String</h3><p>Please note that all arrays are created as strings. Even when using typed arrays such as JSON (<code>ja</code>).</p><pre><code>» ja [0..5]
[
    &quot;0&quot;,
    &quot;1&quot;,
    &quot;2&quot;,
    &quot;3&quot;,
    &quot;4&quot;,
    &quot;5&quot;
]
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,16),g=e("code",null,"[[",-1),f=e("code",null,"[",-1),y=e("code",null,"[",-1),b=e("code",null,"a",-1),x=e("code",null,"count",-1),k=e("code",null,"ja",-1),q=e("code",null,"ta",-1);function w(N,v){const t=d("RouterLink");return s(),l("div",null,[u,h,m,p,e("p",null,[a("Please refer to "),n(t,{to:"/commands/a.html"},{default:r(()=>[a("a (mkarray)")]),_:1}),a(" for more detailed usage of mkarray.")]),_,e("ul",null,[e("li",null,[n(t,{to:"/mkarray/character.html"},{default:r(()=>[a("Character arrays")]),_:1}),a(": Making character arrays (a to z)")]),e("li",null,[n(t,{to:"/mkarray/non-decimal.html"},{default:r(()=>[a("Non-Decimal Ranges")]),_:1}),a(": Create arrays of integers from non-decimal number bases")]),e("li",null,[n(t,{to:"/commands/element.html"},{default:r(()=>[g,a(" (element)")]),_:1}),a(": Outputs an element from a nested structure")]),e("li",null,[n(t,{to:"/commands/index2.html"},{default:r(()=>[f,a(" (index)")]),_:1}),a(": Outputs an element from an array, map or table")]),e("li",null,[n(t,{to:"/commands/range.html"},{default:r(()=>[y,a(" (range) ")]),_:1}),a(": Outputs a ranged subset of data from STDIN")]),e("li",null,[n(t,{to:"/commands/a.html"},{default:r(()=>[b,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array or list")]),e("li",null,[n(t,{to:"/commands/count.html"},{default:r(()=>[x]),_:1}),a(": Count items in a map, list or array")]),e("li",null,[n(t,{to:"/commands/ja.html"},{default:r(()=>[k,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[n(t,{to:"/commands/ta.html"},{default:r(()=>[q,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const C=o(c,[["render",w],["__file","decimal.html.vue"]]);export{C as default};
