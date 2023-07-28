import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as i,d as a,e,b as n,w as r,f as o}from"./app-45f7c304.js";const c={},h=o(`<h1 id="non-decimal-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#non-decimal-ranges-mkarray" aria-hidden="true">#</a> Non-Decimal Ranges - mkarray</h1><blockquote><p>Create arrays of integers from non-decimal number bases</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>When making arrays you can specify ranges of an alternative number base by using an <code>x</code> or <code>.</code> in the end range:</p><pre><code>a: [00..ffx16]
a: [00..ff.16]
</code></pre><p>All number bases from 2 (binary) to 36 (0-9 plus a-z) are supported. Please note that the start and end range are written in the target base while the base identifier is written in decimal: <code>[hex..hex.dec]</code></p><p>Also note that the additional zeros denotes padding (ie the results will start at <code>00</code>, <code>01</code>, etc rather than <code>0</code>, <code>1</code>...)</p>`,7),u=o(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end,start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end][start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>All usages also work with <code>ja</code> and <code>ta</code> as well, eg:</p><pre><code>ja: [start..end] -&gt; \`&lt;stdout&gt;\`
ta: data-type [start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>You can also inline arrays with the <code>%[]</code> syntax, eg:</p><pre><code>%[start..end]
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» a: [08..10x16]
08
09
0a
0b
0c
0d
0e
0f
10

» a: [10..08x16]
10
f
e
d
c
b
a
9
8
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="floating-point-numbers" tabindex="-1"><a class="header-anchor" href="#floating-point-numbers" aria-hidden="true">#</a> Floating Point Numbers</h3><p>If you do need a range of fixed floating point numbers generated then you can do so by merging two decimal integer ranges together. For example</p><pre><code>» a: [05..10x8].[0..7]
05.0
05.1
05.2
05.3
05.4
05.5
05.6
05.7
06.0
06.1
06.2
...
07.5
07.6
07.7
10.0
10.1
10.2
10.3
10.4
10.5
10.6
10.7
</code></pre><h3 id="everything-is-a-string" tabindex="-1"><a class="header-anchor" href="#everything-is-a-string" aria-hidden="true">#</a> Everything Is A String</h3><p>Please note that all arrays are created as strings. Even when using typed arrays such as JSON (<code>ja</code>).</p><pre><code>» ja [0..5]
[
    &quot;0&quot;,
    &quot;1&quot;,
    &quot;2&quot;,
    &quot;3&quot;,
    &quot;4&quot;,
    &quot;5&quot;
]
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,16),m=a("code",null,"[[",-1),p=a("code",null,"[",-1),f=a("code",null,"[",-1),g=a("code",null,"a",-1),y=a("code",null,"count",-1),_=a("code",null,"ja",-1),b=a("code",null,"ta",-1);function x(k,w){const t=s("RouterLink");return l(),i("div",null,[h,a("p",null,[e("Please refer to "),n(t,{to:"/commands/a.html"},{default:r(()=>[e("a (mkarray)")]),_:1}),e(" for more detailed usage of mkarray.")]),u,a("ul",null,[a("li",null,[n(t,{to:"/mkarray/character.html"},{default:r(()=>[e("Character arrays")]),_:1}),e(": Making character arrays (a to z)")]),a("li",null,[n(t,{to:"/mkarray/decimal.html"},{default:r(()=>[e("Decimal Ranges")]),_:1}),e(": Create arrays of decimal integers")]),a("li",null,[n(t,{to:"/commands/element.html"},{default:r(()=>[m,e(" (element)")]),_:1}),e(": Outputs an element from a nested structure")]),a("li",null,[n(t,{to:"/commands/index2.html"},{default:r(()=>[p,e(" (index)")]),_:1}),e(": Outputs an element from an array, map or table")]),a("li",null,[n(t,{to:"/commands/range.html"},{default:r(()=>[f,e(" (range) ")]),_:1}),e(": Outputs a ranged subset of data from STDIN")]),a("li",null,[n(t,{to:"/commands/a.html"},{default:r(()=>[g,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simple way to build an array or list")]),a("li",null,[n(t,{to:"/commands/count.html"},{default:r(()=>[y]),_:1}),e(": Count items in a map, list or array")]),a("li",null,[n(t,{to:"/commands/ja.html"},{default:r(()=>[_,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simply way to build a JSON array")]),a("li",null,[n(t,{to:"/commands/ta.html"},{default:r(()=>[b,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const v=d(c,[["render",x],["__file","non-decimal.html.vue"]]);export{v as default};
