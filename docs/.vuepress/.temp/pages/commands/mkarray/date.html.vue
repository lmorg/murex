<template><div><h1 id="calendar-date-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#calendar-date-ranges-mkarray" aria-hidden="true">#</a> Calendar Date Ranges - mkarray</h1>
<blockquote>
<p>Create arrays of dates</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Unlike bash, Murex also supports date ranges:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: [25-dec-2020..05-jan-2021]
» a: [..25-dec-2020]
» a: [25-dec-2020..]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please refer to <RouterLink to="/commands/a.html">a (mkarray)</RouterLink> for more detailed usage of mkarray.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>a: [start..end] -&gt; `&lt;stdout&gt;`
a: [start..end,start..end] -&gt; `&lt;stdout&gt;`
a: [start..end][start..end] -&gt; `&lt;stdout&gt;`
</code></pre>
<p>All usages also work with <code v-pre>ja</code> and <code v-pre>ta</code> as well, eg:</p>
<pre><code>ja: [start..end] -&gt; `&lt;stdout&gt;`
ta: data-type [start..end] -&gt; `&lt;stdout&gt;`
</code></pre>
<p>You can also inline arrays with the <code v-pre>%[]</code> syntax, eg:</p>
<pre><code>%[start..end]
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>» a: [25-Dec-2020..01-Jan-2021]
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
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="current-date" tabindex="-1"><a class="header-anchor" href="#current-date" aria-hidden="true">#</a> Current Date</h3>
<p>If the start value is missing (eg <code v-pre>[..01-Jan-2020]</code>) then mkarray (<code v-pre>a</code> et al)
will start the range from the current date and count up or down to the end.</p>
<p>If the end value is missing (eg <code v-pre>[01-Jan-2020..]</code>) then mkarray will start at
the start value, as usual, and count up or down to the current date.</p>
<p>For example, if today was 25th December 2020:</p>
<pre><code>» a: [23-December-2020..]
23-December-2020
24-December-2020
25-December-2020

» a: [..23-December-2020]
25-December-2020
24-December-2020
23-December-2020
</code></pre>
<p>This can lead so some fun like countdowns:</p>
<pre><code>» out: &quot;${a: [..01-January-2021] -&gt; len -&gt; =-1} days until the new year!&quot;
7 days until the new year!
</code></pre>
<h3 id="case-sensitivity" tabindex="-1"><a class="header-anchor" href="#case-sensitivity" aria-hidden="true">#</a> Case Sensitivity</h3>
<p>Date ranges are case aware. If the ranges are uppercase then the return will be
uppercase. If the ranges are title case (capital first letter) then the return
will be in title case.</p>
<h4 id="lower-case" tabindex="-1"><a class="header-anchor" href="#lower-case" aria-hidden="true">#</a> lower case</h4>
<pre><code>» a: [01-jan..03-jan]
01-jan
02-jan
03-jan
</code></pre>
<h4 id="title-case" tabindex="-1"><a class="header-anchor" href="#title-case" aria-hidden="true">#</a> Title Case</h4>
<pre><code>» a: [01-Jan..03-Jan]
01-Jan
02-Jan
03-Jan
</code></pre>
<h4 id="upper-case" tabindex="-1"><a class="header-anchor" href="#upper-case" aria-hidden="true">#</a> UPPER CASE</h4>
<pre><code>» a: [01-JAN..03-JAN]
01-JAN
02-JAN
03-JAN
</code></pre>
<h3 id="supported-date-formatting" tabindex="-1"><a class="header-anchor" href="#supported-date-formatting" aria-hidden="true">#</a> Supported Date Formatting</h3>
<p>Below is the source for the supported formatting options for date ranges:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> mkarray

<span class="token keyword">var</span> dateFormat <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	<span class="token comment">// dd mm yy</span>

	<span class="token string">"02-Jan-06"</span><span class="token punctuation">,</span>
	<span class="token string">"02-January-06"</span><span class="token punctuation">,</span>
	<span class="token string">"02-Jan-2006"</span><span class="token punctuation">,</span>
	<span class="token string">"02-January-2006"</span><span class="token punctuation">,</span>

	<span class="token string">"02 Jan 06"</span><span class="token punctuation">,</span>
	<span class="token string">"02 January 06"</span><span class="token punctuation">,</span>
	<span class="token string">"02 Jan 2006"</span><span class="token punctuation">,</span>
	<span class="token string">"02 January 2006"</span><span class="token punctuation">,</span>

	<span class="token string">"02/Jan/06"</span><span class="token punctuation">,</span>
	<span class="token string">"02/January/06"</span><span class="token punctuation">,</span>
	<span class="token string">"02/Jan/2006"</span><span class="token punctuation">,</span>
	<span class="token string">"02/January/2006"</span><span class="token punctuation">,</span>

	<span class="token comment">// mm dd yy</span>

	<span class="token string">"Jan-02-06"</span><span class="token punctuation">,</span>
	<span class="token string">"January-02-06"</span><span class="token punctuation">,</span>
	<span class="token string">"Jan-02-2006"</span><span class="token punctuation">,</span>
	<span class="token string">"January-02-2006"</span><span class="token punctuation">,</span>

	<span class="token string">"Jan 02 06"</span><span class="token punctuation">,</span>
	<span class="token string">"January 02 06"</span><span class="token punctuation">,</span>
	<span class="token string">"Jan 02 2006"</span><span class="token punctuation">,</span>
	<span class="token string">"January 02 2006"</span><span class="token punctuation">,</span>

	<span class="token string">"Jan/02/06"</span><span class="token punctuation">,</span>
	<span class="token string">"January/02/06"</span><span class="token punctuation">,</span>
	<span class="token string">"Jan/02/2006"</span><span class="token punctuation">,</span>
	<span class="token string">"January/02/2006"</span><span class="token punctuation">,</span>

	<span class="token comment">// dd mm</span>

	<span class="token string">"02-Jan"</span><span class="token punctuation">,</span>
	<span class="token string">"02-January"</span><span class="token punctuation">,</span>

	<span class="token string">"02 Jan"</span><span class="token punctuation">,</span>
	<span class="token string">"02 January"</span><span class="token punctuation">,</span>

	<span class="token string">"02/Jan"</span><span class="token punctuation">,</span>
	<span class="token string">"02/January"</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>If you do need any other formatting options not supported there, you can use
<code v-pre>datetime</code> to convert the output of <code v-pre>a</code>. eg:</p>
<pre><code>» a: [01-Jan-2020..03-Jan-2020] -&gt; foreach { -&gt; datetime --in &quot;{go}02-Jan-2006&quot; --out &quot;{py}%A, %d %B&quot;; echo }
Wednesday, 01 January
Thursday, 02 January
Friday, 03 January
</code></pre>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/mkarray/special.html">Special Ranges</RouterLink>:
Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</li>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/range.html"><code v-pre>[</code> (range) </RouterLink>:
Outputs a ranged subset of data from STDIN</li>
<li><RouterLink to="/commands/a.html"><code v-pre>a</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array or list</li>
<li><RouterLink to="/commands/count.html"><code v-pre>count</code></RouterLink>:
Count items in a map, list or array</li>
<li><RouterLink to="/commands/datetime.html"><code v-pre>datetime</code> </RouterLink>:
A date and/or time conversion tool (like <code v-pre>printf</code> but for date and time values)</li>
<li><RouterLink to="/commands/ja.html"><code v-pre>ja</code> (mkarray)</RouterLink>:
A sophisticated yet simply way to build a JSON array</li>
<li><RouterLink to="/commands/mtac.html"><code v-pre>mtac</code></RouterLink>:
Reverse the order of an array</li>
<li><RouterLink to="/commands/ta.html"><code v-pre>ta</code> (mkarray)</RouterLink>:
A sophisticated yet simple way to build an array of a user defined data-type</li>
</ul>
</div></template>


