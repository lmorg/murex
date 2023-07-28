import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r,o as d,c as l,d as n,e as a,b as s,w as t,f as o}from"./app-45f7c304.js";const c={},u=o(`<h1 id="calendar-date-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#calendar-date-ranges-mkarray" aria-hidden="true">#</a> Calendar Date Ranges - mkarray</h1><blockquote><p>Create arrays of dates</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Unlike bash, Murex also supports date ranges:</p><div class="language-text line-numbers-mode" data-ext="text"><pre class="language-text"><code>» a: [25-dec-2020..05-jan-2021]
» a: [..25-dec-2020]
» a: [25-dec-2020..]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div>`,5),p=o(`<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>a: [start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end,start..end] -&gt; \`&lt;stdout&gt;\`
a: [start..end][start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>All usages also work with <code>ja</code> and <code>ta</code> as well, eg:</p><pre><code>ja: [start..end] -&gt; \`&lt;stdout&gt;\`
ta: data-type [start..end] -&gt; \`&lt;stdout&gt;\`
</code></pre><p>You can also inline arrays with the <code>%[]</code> syntax, eg:</p><pre><code>%[start..end]
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>» a: [25-Dec-2020..01-Jan-2021]
25-Dec-2020
26-Dec-2020
27-Dec-2020
28-Dec-2020
29-Dec-2020
30-Dec-2020
31-Dec-2020
01-Jan-2021

» a: [31-Dec..25-Dec]
31-Dec
30-Dec
29-Dec
28-Dec
27-Dec
26-Dec
25-Dec
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="current-date" tabindex="-1"><a class="header-anchor" href="#current-date" aria-hidden="true">#</a> Current Date</h3><p>If the start value is missing (eg <code>[..01-Jan-2020]</code>) then mkarray (<code>a</code> et al) will start the range from the current date and count up or down to the end.</p><p>If the end value is missing (eg <code>[01-Jan-2020..]</code>) then mkarray will start at the start value, as usual, and count up or down to the current date.</p><p>For example, if today was 25th December 2020:</p><pre><code>» a: [23-December-2020..]
23-December-2020
24-December-2020
25-December-2020

» a: [..23-December-2020]
25-December-2020
24-December-2020
23-December-2020
</code></pre><p>This can lead so some fun like countdowns:</p><pre><code>» out: &quot;\${a: [..01-January-2021] -&gt; len -&gt; =-1} days until the new year!&quot;
7 days until the new year!
</code></pre><h3 id="case-sensitivity" tabindex="-1"><a class="header-anchor" href="#case-sensitivity" aria-hidden="true">#</a> Case Sensitivity</h3><p>Date ranges are case aware. If the ranges are uppercase then the return will be uppercase. If the ranges are title case (capital first letter) then the return will be in title case.</p><h4 id="lower-case" tabindex="-1"><a class="header-anchor" href="#lower-case" aria-hidden="true">#</a> lower case</h4><pre><code>» a: [01-jan..03-jan]
01-jan
02-jan
03-jan
</code></pre><h4 id="title-case" tabindex="-1"><a class="header-anchor" href="#title-case" aria-hidden="true">#</a> Title Case</h4><pre><code>» a: [01-Jan..03-Jan]
01-Jan
02-Jan
03-Jan
</code></pre><h4 id="upper-case" tabindex="-1"><a class="header-anchor" href="#upper-case" aria-hidden="true">#</a> UPPER CASE</h4><pre><code>» a: [01-JAN..03-JAN]
01-JAN
02-JAN
03-JAN
</code></pre><h3 id="supported-date-formatting" tabindex="-1"><a class="header-anchor" href="#supported-date-formatting" aria-hidden="true">#</a> Supported Date Formatting</h3><p>Below is the source for the supported formatting options for date ranges:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> mkarray

<span class="token keyword">var</span> dateFormat <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	<span class="token comment">// dd mm yy</span>

	<span class="token string">&quot;02-Jan-06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02-January-06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02-Jan-2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02-January-2006&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;02 Jan 06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02 January 06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02 Jan 2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02 January 2006&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;02/Jan/06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02/January/06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02/Jan/2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02/January/2006&quot;</span><span class="token punctuation">,</span>

	<span class="token comment">// mm dd yy</span>

	<span class="token string">&quot;Jan-02-06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January-02-06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;Jan-02-2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January-02-2006&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;Jan 02 06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January 02 06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;Jan 02 2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January 02 2006&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;Jan/02/06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January/02/06&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;Jan/02/2006&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;January/02/2006&quot;</span><span class="token punctuation">,</span>

	<span class="token comment">// dd mm</span>

	<span class="token string">&quot;02-Jan&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02-January&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;02 Jan&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02 January&quot;</span><span class="token punctuation">,</span>

	<span class="token string">&quot;02/Jan&quot;</span><span class="token punctuation">,</span>
	<span class="token string">&quot;02/January&quot;</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>If you do need any other formatting options not supported there, you can use <code>datetime</code> to convert the output of <code>a</code>. eg:</p><pre><code>» a: [01-Jan-2020..03-Jan-2020] -&gt; foreach { -&gt; datetime --in &quot;{go}02-Jan-2006&quot; --out &quot;{py}%A, %d %B&quot;; echo }
Wednesday, 01 January
Thursday, 02 January
Friday, 03 January
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,30),m=n("code",null,"[[",-1),h=n("code",null,"[",-1),v=n("code",null,"[",-1),k=n("code",null,"a",-1),g=n("code",null,"count",-1),b=n("code",null,"datetime",-1),y=n("code",null,"printf",-1),q=n("code",null,"ja",-1),f=n("code",null,"mtac",-1),J=n("code",null,"ta",-1);function _(D,x){const e=r("RouterLink");return d(),l("div",null,[u,n("p",null,[a("Please refer to "),s(e,{to:"/commands/a.html"},{default:t(()=>[a("a (mkarray)")]),_:1}),a(" for more detailed usage of mkarray.")]),p,n("ul",null,[n("li",null,[s(e,{to:"/mkarray/special.html"},{default:t(()=>[a("Special Ranges")]),_:1}),a(": Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)")]),n("li",null,[s(e,{to:"/commands/element.html"},{default:t(()=>[m,a(" (element)")]),_:1}),a(": Outputs an element from a nested structure")]),n("li",null,[s(e,{to:"/commands/index2.html"},{default:t(()=>[h,a(" (index)")]),_:1}),a(": Outputs an element from an array, map or table")]),n("li",null,[s(e,{to:"/commands/range.html"},{default:t(()=>[v,a(" (range) ")]),_:1}),a(": Outputs a ranged subset of data from STDIN")]),n("li",null,[s(e,{to:"/commands/a.html"},{default:t(()=>[k,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array or list")]),n("li",null,[s(e,{to:"/commands/count.html"},{default:t(()=>[g]),_:1}),a(": Count items in a map, list or array")]),n("li",null,[s(e,{to:"/commands/datetime.html"},{default:t(()=>[b]),_:1}),a(": A date and/or time conversion tool (like "),y,a(" but for date and time values)")]),n("li",null,[s(e,{to:"/commands/ja.html"},{default:t(()=>[q,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simply way to build a JSON array")]),n("li",null,[s(e,{to:"/commands/mtac.html"},{default:t(()=>[f]),_:1}),a(": Reverse the order of an array")]),n("li",null,[s(e,{to:"/commands/ta.html"},{default:t(()=>[J,a(" (mkarray)")]),_:1}),a(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const N=i(c,[["render",_],["__file","date.html.vue"]]);export{N as default};
