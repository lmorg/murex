<template><div><h1 id="special-ranges-mkarray" tabindex="-1"><a class="header-anchor" href="#special-ranges-mkarray" aria-hidden="true">#</a> Special Ranges - mkarray</h1>
<blockquote>
<p>Create arrays from ranges of dictionary terms (eg weekdays, months, seasons, etc)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>Unlike bash, Murex also supports some special ranges:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» a: [mon..sun]
» a: [monday..sunday]
» a: [jan..dec]
» a: [january..december]
» a: [spring..winter]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Please refer to <RouterLink to="/commands/a.html">a (mkarray)</RouterLink> for more detailed usage of mkarray.</p>
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
<pre><code>» a: [summer..winter]
summer
autumn
winter
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="case-sensitivity" tabindex="-1"><a class="header-anchor" href="#case-sensitivity" aria-hidden="true">#</a> Case Sensitivity</h3>
<p>Special ranges are case aware. If the ranges are uppercase then the return will
be uppercase. If the ranges are title case (capital first letter) then the
return will be in title case.</p>
<h4 id="lower-case" tabindex="-1"><a class="header-anchor" href="#lower-case" aria-hidden="true">#</a> lower case</h4>
<pre><code>» a: [monday..wednesday]
monday
tuesday
wednesday
</code></pre>
<h4 id="title-case" tabindex="-1"><a class="header-anchor" href="#title-case" aria-hidden="true">#</a> Title Case</h4>
<pre><code>» a: [Monday..Wednesday]
Monday
Tuesday
Wednesday
</code></pre>
<h4 id="upper-case" tabindex="-1"><a class="header-anchor" href="#upper-case" aria-hidden="true">#</a> UPPER CASE</h4>
<pre><code>» a: [MONDAY..WEDNESDAY]
MONDAY
TUESDAY
WEDNESDAY
</code></pre>
<h3 id="looping-vs-negative-ranges" tabindex="-1"><a class="header-anchor" href="#looping-vs-negative-ranges" aria-hidden="true">#</a> Looping vs Negative Ranges</h3>
<p>Where the special ranges differ from a regular range is they cannot
cannot down. eg <code v-pre>a: [3..1]</code> would output</p>
<pre><code>» a: [3..1]
3
2
1
</code></pre>
<p>however a negative range in special ranges will cycle through to the end
of the range and then loop back from the start:</p>
<pre><code>» a: [Thursday..Wednesday]
Thursday
Friday
Saturday
Sunday
Monday
Tuesday
Wednesday
</code></pre>
<p>This decision was made because generally with ranges of this type, you
would more often prefer to cycle through values rather than iterate
backwards through the list.</p>
<p>If you did want to reverse then pipe the output into another tool:</p>
<pre><code>» a: [Monday..Friday] -&gt; mtac
Friday
Thursday
Wednesday
Tuesday
Monday
</code></pre>
<p>There are other UNIX tools which aren't data type aware but would work in
this specific scenario:</p>
<ul>
<li>
<p><code v-pre>tac</code> (Linux),</p>
</li>
<li>
<p><code v-pre>tail -r</code> (BSD / OS X)</p>
</li>
<li>
<p><code v-pre>perl -e &quot;print reverse &lt;&gt;&quot;</code> (Multi-platform but requires Perl installed)</p>
</li>
</ul>
<h3 id="supported-dictionary-terms" tabindex="-1"><a class="header-anchor" href="#supported-dictionary-terms" aria-hidden="true">#</a> Supported Dictionary Terms</h3>
<p>Below is the source for the supported dictionary terms:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> mkarray

<span class="token keyword">var</span> mapRanges <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	rangeWeekdayLong<span class="token punctuation">,</span>
	rangeWeekdayShort<span class="token punctuation">,</span>
	rangeMonthLong<span class="token punctuation">,</span>
	rangeMonthShort<span class="token punctuation">,</span>
	rangeSeason<span class="token punctuation">,</span>
	rangeMoon<span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeWeekdayLong <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"monday"</span><span class="token punctuation">:</span>    <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"tuesday"</span><span class="token punctuation">:</span>   <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"wednesday"</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"thursday"</span><span class="token punctuation">:</span>  <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">"friday"</span><span class="token punctuation">:</span>    <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">"saturday"</span><span class="token punctuation">:</span>  <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">"sunday"</span><span class="token punctuation">:</span>    <span class="token number">7</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeWeekdayShort <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"mon"</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"tue"</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"wed"</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"thu"</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">"fri"</span><span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">"sat"</span><span class="token punctuation">:</span> <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">"sun"</span><span class="token punctuation">:</span> <span class="token number">7</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMonthLong <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"january"</span><span class="token punctuation">:</span>   <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"february"</span><span class="token punctuation">:</span>  <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"march"</span><span class="token punctuation">:</span>     <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"april"</span><span class="token punctuation">:</span>     <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">"may"</span><span class="token punctuation">:</span>       <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">"june"</span><span class="token punctuation">:</span>      <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">"july"</span><span class="token punctuation">:</span>      <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">"august"</span><span class="token punctuation">:</span>    <span class="token number">8</span><span class="token punctuation">,</span>
	<span class="token string">"september"</span><span class="token punctuation">:</span> <span class="token number">9</span><span class="token punctuation">,</span>
	<span class="token string">"october"</span><span class="token punctuation">:</span>   <span class="token number">10</span><span class="token punctuation">,</span>
	<span class="token string">"november"</span><span class="token punctuation">:</span>  <span class="token number">11</span><span class="token punctuation">,</span>
	<span class="token string">"december"</span><span class="token punctuation">:</span>  <span class="token number">12</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMonthShort <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"jan"</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"feb"</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"mar"</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"apr"</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">"may"</span><span class="token punctuation">:</span> <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">"jun"</span><span class="token punctuation">:</span> <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">"jul"</span><span class="token punctuation">:</span> <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">"aug"</span><span class="token punctuation">:</span> <span class="token number">8</span><span class="token punctuation">,</span>
	<span class="token string">"sep"</span><span class="token punctuation">:</span> <span class="token number">9</span><span class="token punctuation">,</span>
	<span class="token string">"oct"</span><span class="token punctuation">:</span> <span class="token number">10</span><span class="token punctuation">,</span>
	<span class="token string">"nov"</span><span class="token punctuation">:</span> <span class="token number">11</span><span class="token punctuation">,</span>
	<span class="token string">"dec"</span><span class="token punctuation">:</span> <span class="token number">12</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeSeason <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"spring"</span><span class="token punctuation">:</span> <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"summer"</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"autumn"</span><span class="token punctuation">:</span> <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"winter"</span><span class="token punctuation">:</span> <span class="token number">4</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token keyword">var</span> rangeMoon <span class="token operator">=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">{</span>
	<span class="token string">"new moon"</span><span class="token punctuation">:</span>        <span class="token number">1</span><span class="token punctuation">,</span>
	<span class="token string">"waxing crescent"</span><span class="token punctuation">:</span> <span class="token number">2</span><span class="token punctuation">,</span>
	<span class="token string">"first quarter"</span><span class="token punctuation">:</span>   <span class="token number">3</span><span class="token punctuation">,</span>
	<span class="token string">"waxing gibbous"</span><span class="token punctuation">:</span>  <span class="token number">4</span><span class="token punctuation">,</span>
	<span class="token string">"full moon"</span><span class="token punctuation">:</span>       <span class="token number">5</span><span class="token punctuation">,</span>
	<span class="token string">"waning gibbous"</span><span class="token punctuation">:</span>  <span class="token number">6</span><span class="token punctuation">,</span>
	<span class="token string">"third quarter"</span><span class="token punctuation">:</span>   <span class="token number">7</span><span class="token punctuation">,</span>
	<span class="token string">"waning crescent"</span><span class="token punctuation">:</span> <span class="token number">8</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/mkarray/date.html">Calendar Date Ranges</RouterLink>:
Create arrays of dates</li>
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


