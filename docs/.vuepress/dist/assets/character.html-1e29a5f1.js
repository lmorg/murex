import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as l,c as d,d as a,e,b as r,w as n,f as c}from"./app-45f7c304.js";const i={},u=a("h1",{id:"character-arrays-mkarray",tabindex:"-1"},[a("a",{class:"header-anchor",href:"#character-arrays-mkarray","aria-hidden":"true"},"#"),e(" Character arrays - mkarray")],-1),m=a("blockquote",null,[a("p",null,"Making character arrays (a to z)")],-1),h=a("h2",{id:"description",tabindex:"-1"},[a("a",{class:"header-anchor",href:"#description","aria-hidden":"true"},"#"),e(" Description")],-1),_=a("p",null,"You can create arrays from a range of letters (a to z):",-1),p=a("pre",null,[a("code",null,`» a: [a..z]
» a: [z..a]
» a: [A..Z]
» a: [Z..A]
`)],-1),f=a("p",null,"...or any characters within that range.",-1),y=c(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end,start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end][start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>All usages also work with <code>ja</code> and <code>ta</code> as well, eg:</p><pre><code>ja: [start..end] -&gt; \`&lt;stdout&gt;\`
ta: data-type [start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>You can also inline arrays with the <code>%[]</code> syntax, eg:</p><pre><code>%[start..end]
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» a: [a..c]
a
b
c

» a: [c..a]
c
b
a
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,9),g=a("code",null,"[[",-1),k=a("code",null,"[",-1),b=a("code",null,"[",-1),x=a("code",null,"a",-1),w=a("code",null,"count",-1),A=a("code",null,"ja",-1),N=a("code",null,"ta",-1);function C(j,z){const t=s("RouterLink");return l(),d("div",null,[u,m,h,_,p,f,a("p",null,[e("Please refer to "),r(t,{to:"/commands/a.html"},{default:n(()=>[e("a (mkarray)")]),_:1}),e(" for more detailed usage of mkarray.")]),y,a("ul",null,[a("li",null,[r(t,{to:"/mkarray/decimal.html"},{default:n(()=>[e("Decimal Ranges")]),_:1}),e(": Create arrays of decimal integers")]),a("li",null,[r(t,{to:"/mkarray/non-decimal.html"},{default:n(()=>[e("Non-Decimal Ranges")]),_:1}),e(": Create arrays of integers from non-decimal number bases")]),a("li",null,[r(t,{to:"/commands/element.html"},{default:n(()=>[g,e(" (element)")]),_:1}),e(": Outputs an element from a nested structure")]),a("li",null,[r(t,{to:"/commands/index2.html"},{default:n(()=>[k,e(" (index)")]),_:1}),e(": Outputs an element from an array, map or table")]),a("li",null,[r(t,{to:"/commands/range.html"},{default:n(()=>[b,e(" (range) ")]),_:1}),e(": Outputs a ranged subset of data from STDIN")]),a("li",null,[r(t,{to:"/commands/a.html"},{default:n(()=>[x,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simple way to build an array or list")]),a("li",null,[r(t,{to:"/commands/count.html"},{default:n(()=>[w]),_:1}),e(": Count items in a map, list or array")]),a("li",null,[r(t,{to:"/commands/ja.html"},{default:n(()=>[A,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simply way to build a JSON array")]),a("li",null,[r(t,{to:"/commands/ta.html"},{default:n(()=>[N,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const R=o(i,[["render",C],["__file","character.html.vue"]]);export{R as default};
