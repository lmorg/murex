import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as s,o as d,c as l,d as e,b as r,w as n,e as a,f as i}from"./app-45f7c304.js";const c={},u=i(`<h1 id="a-mkarray" tabindex="-1"><a class="header-anchor" href="#a-mkarray" aria-hidden="true">#</a> <code>a</code> (mkarray)</h1><blockquote><p>A sophisticated yet simple way to build an array or list</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Pronounced &quot;make array&quot;, like <code>mkdir</code> (etc), Murex has a pretty sophisticated builtin for generating arrays. Think like bash&#39;s <code>{1..9}</code> syntax:</p><pre><code>a: [1..9]
</code></pre><p>Except Murex also supports other sets of ranges like dates, days of the week, and alternative number bases.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
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
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="advanced-array-syntax" tabindex="-1"><a class="header-anchor" href="#advanced-array-syntax" aria-hidden="true">#</a> Advanced Array Syntax</h3><p>The syntax for <code>a</code> is a comma separated list of parameters with expansions stored in square brackets. You can have an expansion embedded inside a parameter or as it&#39;s own parameter. Expansions can also have multiple parameters.</p><pre><code>» a: 01,02,03,05,06,07
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
</code></pre><p>You can also have multiple expansion blocks in a single parameter:</p><pre><code>» a: a[1..3]b[5..7]
a1b5
a1b6
a1b7
a2b5
a2b6
a2b7
a3b5
a3b6
a3b7
</code></pre><p><code>a</code> will cycle through each iteration of the last expansion, moving itself backwards through the string; behaving like an normal counter.</p><h3 id="creating-json-arrays-with-ja" tabindex="-1"><a class="header-anchor" href="#creating-json-arrays-with-ja" aria-hidden="true">#</a> Creating JSON arrays with <code>ja</code></h3><p>As you can see from the previous examples, <code>a</code> returns the array as a list of strings. This is so you can stream excessively long arrays, for example every IPv4 address: <code>a: [0..254].[0..254].[0..254].[0..254]</code> (this kind of array expansion would hang bash).</p><p>However if you needed a JSON string then you can use all the same syntax as <code>a</code> but forgo the streaming capability:</p><pre><code>» ja: [Monday..Sunday]
[
    &quot;Monday&quot;,
    &quot;Tuesday&quot;,
    &quot;Wednesday&quot;,
    &quot;Thursday&quot;,
    &quot;Friday&quot;,
    &quot;Saturday&quot;,
    &quot;Sunday&quot;
]
</code></pre><p>This is particularly useful if you are adding formatting that might break under <code>a</code>&#39;s formatting (which uses the <code>str</code> data type).</p><h3 id="smart-arrays" tabindex="-1"><a class="header-anchor" href="#smart-arrays" aria-hidden="true">#</a> Smart arrays</h3><p>Murex supports a number of different formats that can be used to generate arrays. For more details on these please refer to the documents for each format</p>`,28),h=e("h2",{id:"see-also",tabindex:"-1"},[e("a",{class:"header-anchor",href:"#see-also","aria-hidden":"true"},"#"),a(" See Also")],-1),m=e("code",null,"%[]",-1),p=e("code",null,"[[",-1),y=e("code",null,"[",-1),f=e("code",null,"[",-1),g=e("code",null,"count",-1),b=e("code",null,"ja",-1),_=e("code",null,"mtac",-1),x=e("code",null,"str",-1),k=e("code",null,"ta",-1);function w(q,v){const t=s("RouterLink");return d(),l("div",null,[u,e("ul",null,[e("li",null,[r(t,{to:"/mkarray/date.html"},{default:n(()=>[a("Calendar Date Ranges")]),_:1}),a(": Create arrays of dates")]),e("li",null,[r(t,{to:"/mkarray/character.html"},{default:n(()=>[a("Character arrays")]),_:1}),a(": Making character arrays (a to z)")]),e("li",null,[r(t,{to:"/mkarray/decimal.html"},{default:n(()=>[a("Decimal Ranges")]),_:1}),a(": Create arrays of decimal integers")]),e("li",null,[r(t,{to:"/mkarray/non-decimal.html"},{default:n(()=>[a("Non-Decimal Ranges")]),_:1}),a(": Create arrays of integers from non-decimal number bases")]),e("li",null,[r(t,{to:"/mkarray/special.html"},{default:n(()=>[a("Special Ranges")]),_:1}),a(": Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)")])]),h,e("ul",null,[e("li",null,[r(t,{to:"/parser/create-array.html"},{default:n(()=>[a("Create array ("),m,a(") constructor")]),_:1}),a(": Quickly generate arrays")]),e("li",null,[r(t,{to:"/commands/element.html"},{default:n(()=>[p,a(" (element)")]),_:1}),a(": Outputs an element from a nested structure")]),e("li",null,[r(t,{to:"/commands/index2.html"},{default:n(()=>[y,a(" (index)")]),_:1}),a(": Outputs an element from an array, map or table")]),e("li",null,[r(t,{to:"/commands/range.html"},{default:n(()=>[f,a(" (range) ")]),_:1}),a(": Outputs a ranged subset of data from STDIN")]),e("li",null,[r(t,{to:"/commands/count.html"},{default:n(()=>[g]),_:1}),a(": Count items in a map, list or array")]),e("li",null,[r(t,{to:"/commands/ja.html"},{default:n(()=>[b,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[r(t,{to:"/commands/mtac.html"},{default:n(()=>[_]),_:1}),a(": Reverse the order of an array")]),e("li",null,[r(t,{to:"/types/str.html"},{default:n(()=>[x,a(" (string) ")]),_:1}),a(": string (primitive)")]),e("li",null,[r(t,{to:"/commands/ta.html"},{default:n(()=>[k,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const j=o(c,[["render",w],["__file","a.html.vue"]]);export{j as default};
