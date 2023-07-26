import{_ as p}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as u,c,d as n,e as s,b as t,w as e,f as o}from"./app-45f7c304.js";const l={},r=o(`<h1 id="special-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#special-ranges-mkarray" aria-hidden="true">#</a> Special Ranges - mkarray</h1><blockquote><p>Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Unlike bash, Murex also supports some special ranges:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» a: [mon..sun]
» a: [monday..sunday]
» a: [jan..dec]
» a: [january..december]
» a: [spring..winter]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5),d=o(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end,start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end][start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>All usages also work with <code>ja</code> and <code>ta</code> as well, eg:</p><pre><code>ja: [start..end] -&gt; \`&lt;stdout&gt;\`
ta: data-type [start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>You can also inline arrays with the <code>%[]</code> syntax, eg:</p><pre><code>%[start..end]
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» a: [summer..winter]
summer
autumn
winter
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="case-sensitivity" tabindex="-1"><a class="header-anchor" href="#case-sensitivity" aria-hidden="true">#</a> Case Sensitivity</h3><p>Special ranges are case aware. If the ranges are uppercase then the return will be uppercase. If the ranges are title case (capital first letter) then the return will be in title case.</p><h4 id="lower-case" tabindex="-1"><a class="header-anchor" href="#lower-case" aria-hidden="true">#</a> lower case</h4><pre><code>» a: [monday..wednesday]
monday
tuesday
wednesday
</code></pre><h4 id="title-case" tabindex="-1"><a class="header-anchor" href="#title-case" aria-hidden="true">#</a> Title Case</h4><pre><code>» a: [Monday..Wednesday]
Monday
Tuesday
Wednesday
</code></pre><h4 id="upper-case" tabindex="-1"><a class="header-anchor" href="#upper-case" aria-hidden="true">#</a> UPPER CASE</h4><pre><code>» a: [MONDAY..WEDNESDAY]
MONDAY
TUESDAY
WEDNESDAY
</code></pre><h3 id="looping-vs-negative-ranges" tabindex="-1"><a class="header-anchor" href="#looping-vs-negative-ranges" aria-hidden="true">#</a> Looping vs Negative Ranges</h3><p>Where the special ranges differ from a regular range is they cannot cannot down. eg <code>a: [3..1]</code> would output</p><pre><code>» a: [3..1]
3
2
1
</code></pre><p>however a negative range in special ranges will cycle through to the end of the range and then loop back from the start:</p><pre><code>» a: [Thursday..Wednesday]
Thursday
Friday
Saturday
Sunday
Monday
Tuesday
Wednesday
</code></pre><p>This decision was made because generally with ranges of this type, you would more often prefer to cycle through values rather than iterate backwards through the list.</p><p>If you did want to reverse then pipe the output into another tool:</p><pre><code>» a: [Monday..Friday] -&gt; mtac
Friday
Thursday
Wednesday
Tuesday
Monday
</code></pre><p>There are other UNIX tools which aren&#39;t data type aware but would work in this specific scenario:</p><ul><li><p><code>tac</code> (Linux),</p></li><li><p><code>tail -r</code> (BSD / OS X)</p></li><li><p><code>perl -e &quot;print reverse &lt;&gt;&quot;</code> (Multi-platform but requires Perl installed)</p></li></ul><h3 id="supported-dictionary-terms" tabindex="-1"><a class="header-anchor" href="#supported-dictionary-terms" aria-hidden="true">#</a> Supported Dictionary Terms</h3><p>Below is the source for the supported dictionary terms:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> mkarray

<span class="token keyword">var</span> mapRanges <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	rangeWeekdayLong<span class="token punctuation">,</span>
	rangeWeekdayShort<span class="token punctuation">,</span>
	rangeMonthLong<span class="token punctuation">,</span>
	rangeMonthShort<span class="token punctuation">,</span>
	rangeSeason<span class="token punctuation">,</span>
	rangeMoon<span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeWeekdayLong <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;monday&quot;</span><span class="token punctuation">:</span>    <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;tuesday&quot;</span><span class="token punctuation">:</span>   <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;wednesday&quot;</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;thursday&quot;</span><span class="token punctuation">:</span>  <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">&quot;friday&quot;</span><span class="token punctuation">:</span>    <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">&quot;saturday&quot;</span><span class="token punctuation">:</span>  <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">&quot;sunday&quot;</span><span class="token punctuation">:</span>    <span class="token number">7</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeWeekdayShort <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;mon&quot;</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;tue&quot;</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;wed&quot;</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;thu&quot;</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">&quot;fri&quot;</span><span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">&quot;sat&quot;</span><span class="token punctuation">:</span> <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">&quot;sun&quot;</span><span class="token punctuation">:</span> <span class="token number">7</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMonthLong <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;january&quot;</span><span class="token punctuation">:</span>   <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;february&quot;</span><span class="token punctuation">:</span>  <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;march&quot;</span><span class="token punctuation">:</span>     <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;april&quot;</span><span class="token punctuation">:</span>     <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">&quot;may&quot;</span><span class="token punctuation">:</span>       <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">&quot;june&quot;</span><span class="token punctuation">:</span>      <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">&quot;july&quot;</span><span class="token punctuation">:</span>      <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">&quot;august&quot;</span><span class="token punctuation">:</span>    <span class="token number">8</span><span class="token punctuation">,</span>
	<span class="token string">&quot;september&quot;</span><span class="token punctuation">:</span> <span class="token number">9</span><span class="token punctuation">,</span>
	<span class="token string">&quot;october&quot;</span><span class="token punctuation">:</span>   <span class="token number">10</span><span class="token punctuation">,</span>
	<span class="token string">&quot;november&quot;</span><span class="token punctuation">:</span>  <span class="token number">11</span><span class="token punctuation">,</span>
	<span class="token string">&quot;december&quot;</span><span class="token punctuation">:</span>  <span class="token number">12</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMonthShort <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;jan&quot;</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;feb&quot;</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;mar&quot;</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;apr&quot;</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">&quot;may&quot;</span><span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">&quot;jun&quot;</span><span class="token punctuation">:</span> <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">&quot;jul&quot;</span><span class="token punctuation">:</span> <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">&quot;aug&quot;</span><span class="token punctuation">:</span> <span class="token number">8</span><span class="token punctuation">,</span>
	<span class="token string">&quot;sep&quot;</span><span class="token punctuation">:</span> <span class="token number">9</span><span class="token punctuation">,</span>
	<span class="token string">&quot;oct&quot;</span><span class="token punctuation">:</span> <span class="token number">10</span><span class="token punctuation">,</span>
	<span class="token string">&quot;nov&quot;</span><span class="token punctuation">:</span> <span class="token number">11</span><span class="token punctuation">,</span>
	<span class="token string">&quot;dec&quot;</span><span class="token punctuation">:</span> <span class="token number">12</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeSeason <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;spring&quot;</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;summer&quot;</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;autumn&quot;</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;winter&quot;</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMoon <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">&quot;new moon&quot;</span><span class="token punctuation">:</span>        <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">&quot;waxing crescent&quot;</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">&quot;first quarter&quot;</span><span class="token punctuation">:</span>   <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">&quot;waxing gibbous&quot;</span><span class="token punctuation">:</span>  <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">&quot;full moon&quot;</span><span class="token punctuation">:</span>       <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">&quot;waning gibbous&quot;</span><span class="token punctuation">:</span>  <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">&quot;third quarter&quot;</span><span class="token punctuation">:</span>   <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">&quot;waning crescent&quot;</span><span class="token punctuation">:</span> <span class="token number">8</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,31),k=n("code",null,"[[",-1),m=n("code",null,"[",-1),v=n("code",null,"[",-1),b=n("code",null,"a",-1),h=n("code",null,"count",-1),g=n("code",null,"datetime",-1),q=n("code",null,"printf",-1),y=n("code",null,"ja",-1),f=n("code",null,"mtac",-1),w=n("code",null,"ta",-1);function _(x,S){const a=i("RouterLink");return u(),c("div",null,[r,n("p",null,[s("Please refer to "),t(a,{to:"/commands/a.html"},{default:e(()=>[s("a (mkarray)")]),_:1}),s(" for more detailed usage of mkarray.")]),d,n("ul",null,[n("li",null,[t(a,{to:"/mkarray/date.html"},{default:e(()=>[s("Calendar Date Ranges")]),_:1}),s(": Create arrays of dates")]),n("li",null,[t(a,{to:"/commands/element.html"},{default:e(()=>[k,s(" (element)")]),_:1}),s(": Outputs an element from a nested structure")]),n("li",null,[t(a,{to:"/commands/index2.html"},{default:e(()=>[m,s(" (index)")]),_:1}),s(": Outputs an element from an array, map or table")]),n("li",null,[t(a,{to:"/commands/range.html"},{default:e(()=>[v,s(" (range) ")]),_:1}),s(": Outputs a ranged subset of data from STDIN")]),n("li",null,[t(a,{to:"/commands/a.html"},{default:e(()=>[b,s(" (mkarray)")]),_:1}),s(": A sophisticated yet simple way to build an array or list")]),n("li",null,[t(a,{to:"/commands/count.html"},{default:e(()=>[h]),_:1}),s(": Count items in a map, list or array")]),n("li",null,[t(a,{to:"/commands/datetime.html"},{default:e(()=>[g]),_:1}),s(": A date and/or time conversion tool (like "),q,s(" but for date and time values)")]),n("li",null,[t(a,{to:"/commands/ja.html"},{default:e(()=>[y,s(" (mkarray)")]),_:1}),s(": A sophisticated yet simply way to build a JSON array")]),n("li",null,[t(a,{to:"/commands/mtac.html"},{default:e(()=>[f]),_:1}),s(": Reverse the order of an array")]),n("li",null,[t(a,{to:"/commands/ta.html"},{default:e(()=>[w,s(" (mkarray)")]),_:1}),s(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const T=p(l,[["render",_],["__file","special.html.vue"]]);export{T as default};
