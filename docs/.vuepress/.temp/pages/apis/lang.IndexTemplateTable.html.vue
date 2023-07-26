<template><div><h1 id="lang-indextemplatetable-template-api" tabindex="-1"><a class="header-anchor" href="#lang-indextemplatetable-template-api" aria-hidden="true">#</a> <code v-pre>lang.IndexTemplateTable()</code> (template API)</h1>
<blockquote>
<p>Returns element(s) from a table</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a template API you can use for your custom data types.</p>
<p>It should only be called from <code v-pre>ReadIndex()</code> and <code v-pre>ReadNotIndex()</code> functions.</p>
<p>This function ensures consistency with the index, <code v-pre>[</code>, builtin when used with
different Murex data types. Thus making indexing a data type agnostic
capability.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example calling <code v-pre>lang.IndexTemplateTable()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> generic

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"bytes"</span>
	<span class="token string">"strings"</span>

	<span class="token string">"github.com/lmorg/murex/lang"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">index</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	cRecords <span class="token operator">:=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">chan</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> <span class="token number">1</span><span class="token punctuation">)</span>

	<span class="token keyword">go</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdin<span class="token punctuation">.</span><span class="token function">ReadLine</span><span class="token punctuation">(</span><span class="token keyword">func</span><span class="token punctuation">(</span>b <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			cRecords <span class="token operator">&lt;-</span> rxWhitespace<span class="token punctuation">.</span><span class="token function">Split</span><span class="token punctuation">(</span><span class="token function">string</span><span class="token punctuation">(</span>bytes<span class="token punctuation">.</span><span class="token function">TrimSpace</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> <span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span><span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
			p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
		<span class="token function">close</span><span class="token punctuation">(</span>cRecords<span class="token punctuation">)</span>
	<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>s <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>b <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		b <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>strings<span class="token punctuation">.</span><span class="token function">Join</span><span class="token punctuation">(</span>s<span class="token punctuation">,</span> <span class="token string">"\t"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token keyword">return</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">IndexTemplateTable</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> params<span class="token punctuation">,</span> cRecords<span class="token punctuation">,</span> marshaller<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="api-source" tabindex="-1"><a class="header-anchor" href="#api-source" aria-hidden="true">#</a> API Source:</h3>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"errors"</span>
	<span class="token string">"fmt"</span>
	<span class="token string">"regexp"</span>
	<span class="token string">"strconv"</span>
	<span class="token string">"strings"</span>

	<span class="token string">"github.com/lmorg/murex/utils"</span>
<span class="token punctuation">)</span>

<span class="token keyword">const</span> <span class="token punctuation">(</span>
	byRowNumber <span class="token operator">=</span> <span class="token boolean">iota</span> <span class="token operator">+</span> <span class="token number">1</span>
	byColumnNumber
	byColumnName

	maxReportedUnmatched <span class="token operator">=</span> <span class="token number">5</span>
<span class="token punctuation">)</span>

<span class="token keyword">var</span> <span class="token punctuation">(</span>
	rxColumnPrefixOld <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">`^:[0-9]+$`</span><span class="token punctuation">)</span>
	rxRowSuffixOld    <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">`^[0-9]+:$`</span><span class="token punctuation">)</span>
	rxColumnPrefixNew <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">`^\*[a-zA-Z]$`</span><span class="token punctuation">)</span>
	rxRowSuffixNew    <span class="token operator">=</span> regexp<span class="token punctuation">.</span><span class="token function">MustCompile</span><span class="token punctuation">(</span><span class="token string">`^\*[0-9]+$`</span><span class="token punctuation">)</span>
	errMixAndMatch    <span class="token operator">=</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">"you cannot mix and match matching modes"</span><span class="token punctuation">)</span>
<span class="token punctuation">)</span>

<span class="token comment">// IndexTemplateTable is a handy standard indexer you can use in your custom data types for tabulated / streamed data.</span>
<span class="token comment">// The point of this is to minimize code rewriting and standardising the behavior of the indexer.</span>
<span class="token keyword">func</span> <span class="token function">IndexTemplateTable</span><span class="token punctuation">(</span>p <span class="token operator">*</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> cRecords <span class="token keyword">chan</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> marshaller <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">if</span> p<span class="token punctuation">.</span>IsNot <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token function">ittNot</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> params<span class="token punctuation">,</span> cRecords<span class="token punctuation">,</span> marshaller<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token keyword">return</span> <span class="token function">ittIndex</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> params<span class="token punctuation">,</span> cRecords<span class="token punctuation">,</span> marshaller<span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">charToIndex</span><span class="token punctuation">(</span>b <span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">int</span> <span class="token punctuation">{</span>
	<span class="token keyword">if</span> b <span class="token operator">></span> <span class="token number">96</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token function">int</span><span class="token punctuation">(</span>b <span class="token operator">-</span> <span class="token number">97</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token keyword">return</span> <span class="token function">int</span><span class="token punctuation">(</span>b <span class="token operator">-</span> <span class="token number">65</span><span class="token punctuation">)</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">ittIndex</span><span class="token punctuation">(</span>p <span class="token operator">*</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> cRecords <span class="token keyword">chan</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> marshaller <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">var</span> <span class="token punctuation">(</span>
		mode           <span class="token builtin">int</span>
		matchStr       <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
		matchInt       <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span>
		unmatched      <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
		unmatchedCount <span class="token builtin">int</span>
	<span class="token punctuation">)</span>

	<span class="token keyword">defer</span> <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>unmatched<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
			p<span class="token punctuation">.</span>ExitNum <span class="token operator">=</span> <span class="token number">1</span>
			<span class="token keyword">if</span> unmatchedCount <span class="token operator">></span> maxReportedUnmatched <span class="token punctuation">{</span>
				unmatched <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>unmatched<span class="token punctuation">,</span> fmt<span class="token punctuation">.</span><span class="token function">Sprintf</span><span class="token punctuation">(</span><span class="token string">"...plus %d more"</span><span class="token punctuation">,</span> unmatchedCount<span class="token operator">-</span>maxReportedUnmatched<span class="token punctuation">)</span><span class="token punctuation">)</span>
			<span class="token punctuation">}</span>
			err <span class="token operator">=</span> fmt<span class="token punctuation">.</span><span class="token function">Errorf</span><span class="token punctuation">(</span><span class="token string">"some records did not contain all of the requested fields:%s%s"</span><span class="token punctuation">,</span>
				utils<span class="token punctuation">.</span>NewLineString<span class="token punctuation">,</span>
				strings<span class="token punctuation">.</span><span class="token function">Join</span><span class="token punctuation">(</span>unmatched<span class="token punctuation">,</span> utils<span class="token punctuation">.</span>NewLineString<span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

	errUnmatched <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>recs <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		unmatchedCount<span class="token operator">++</span>
		<span class="token keyword">if</span> unmatchedCount <span class="token operator">></span> maxReportedUnmatched <span class="token punctuation">{</span>
			<span class="token keyword">return</span>
		<span class="token punctuation">}</span>
		unmatched <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>unmatched<span class="token punctuation">,</span> strings<span class="token punctuation">.</span><span class="token function">Join</span><span class="token punctuation">(</span>recs<span class="token punctuation">,</span> <span class="token string">"\t"</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> params <span class="token punctuation">{</span>
		<span class="token keyword">switch</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> rxRowSuffixOld<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byRowNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byRowNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token punctuation">:</span><span class="token function">len</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>matchInt<span class="token punctuation">,</span> num<span class="token punctuation">)</span>

		<span class="token keyword">case</span> rxRowSuffixNew<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byRowNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byRowNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>matchInt<span class="token punctuation">,</span> num<span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">)</span> <span class="token comment">// Don't count from zero</span>

		<span class="token keyword">case</span> rxColumnPrefixOld<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byColumnNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>matchInt<span class="token punctuation">,</span> num<span class="token punctuation">)</span>

		<span class="token keyword">case</span> rxColumnPrefixNew<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byColumnNumber
			num <span class="token operator">:=</span> <span class="token function">charToIndex</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>matchInt<span class="token punctuation">,</span> num<span class="token punctuation">)</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnName <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			matchStr <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>matchStr<span class="token punctuation">,</span> params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
			mode <span class="token operator">=</span> byColumnName

		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> mode <span class="token punctuation">{</span>
	<span class="token keyword">case</span> byRowNumber<span class="token punctuation">:</span>
		<span class="token keyword">var</span> <span class="token punctuation">(</span>
			ordered <span class="token operator">=</span> <span class="token boolean">true</span>
			last    <span class="token builtin">int</span>
			max     <span class="token builtin">int</span>
		<span class="token punctuation">)</span>
		<span class="token comment">// check order</span>
		<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> matchInt <span class="token punctuation">{</span>
			<span class="token keyword">if</span> i <span class="token operator">&lt;</span> last <span class="token punctuation">{</span>
				ordered <span class="token operator">=</span> <span class="token boolean">false</span>
			<span class="token punctuation">}</span>
			<span class="token keyword">if</span> i <span class="token operator">></span> max <span class="token punctuation">{</span>
				max <span class="token operator">=</span> i
			<span class="token punctuation">}</span>
			last <span class="token operator">=</span> i
		<span class="token punctuation">}</span>

		<span class="token keyword">if</span> ordered <span class="token punctuation">{</span>
			<span class="token comment">// ordered matching - for this we can just read in the records we want sequentially. Low memory overhead</span>
			<span class="token keyword">var</span> i <span class="token builtin">int</span>
			<span class="token keyword">for</span> <span class="token punctuation">{</span>
				recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
				<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
					<span class="token keyword">return</span> <span class="token boolean">nil</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">if</span> i <span class="token operator">==</span> matchInt<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>
					<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
						p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
					<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>matchInt<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">1</span> <span class="token punctuation">{</span>
						matchInt<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token operator">-</span><span class="token number">1</span>
						<span class="token keyword">return</span> <span class="token boolean">nil</span>
					<span class="token punctuation">}</span>
					matchInt <span class="token operator">=</span> matchInt<span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token punctuation">]</span>
				<span class="token punctuation">}</span>
				i<span class="token operator">++</span>
			<span class="token punctuation">}</span>

		<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
			<span class="token comment">// unordered matching - for this we load the entire data set into memory - up until the maximum value</span>
			<span class="token keyword">var</span> <span class="token punctuation">(</span>
				i     <span class="token builtin">int</span>
				lines <span class="token operator">=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> max<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">)</span>
			<span class="token punctuation">)</span>
			<span class="token keyword">for</span> <span class="token punctuation">{</span>
				recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
				<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
					<span class="token keyword">break</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">if</span> i <span class="token operator">&lt;=</span> max <span class="token punctuation">{</span>
					lines<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> recs
				<span class="token punctuation">}</span>
				i<span class="token operator">++</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> j <span class="token operator">:=</span> <span class="token keyword">range</span> matchInt <span class="token punctuation">{</span>
				<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>lines<span class="token punctuation">[</span>j<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
					p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">return</span> <span class="token boolean">nil</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">case</span> byColumnNumber<span class="token punctuation">:</span>
		<span class="token keyword">for</span> <span class="token punctuation">{</span>
			recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
			<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
				<span class="token keyword">return</span> <span class="token boolean">nil</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">var</span> line <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
			<span class="token keyword">for</span> <span class="token boolean">_</span><span class="token punctuation">,</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> matchInt <span class="token punctuation">{</span>
				<span class="token keyword">if</span> i <span class="token operator">&lt;</span> <span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span> <span class="token punctuation">{</span>
					line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
					<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token operator">||</span> <span class="token punctuation">(</span><span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">1</span> <span class="token operator">&amp;&amp;</span> recs<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span> <span class="token operator">==</span> <span class="token string">""</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
						<span class="token keyword">continue</span>
					<span class="token punctuation">}</span>
					<span class="token function">errUnmatched</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
			<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
				<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
					p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">case</span> byColumnName<span class="token punctuation">:</span>
		<span class="token keyword">var</span> <span class="token punctuation">(</span>
			lineNum  <span class="token builtin">int</span>
			headings <span class="token operator">=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">)</span>
		<span class="token punctuation">)</span>

		<span class="token keyword">for</span> <span class="token punctuation">{</span>
			<span class="token keyword">var</span> line <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
			recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
			<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
				<span class="token keyword">return</span> <span class="token boolean">nil</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">if</span> lineNum <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> recs <span class="token punctuation">{</span>
					headings<span class="token punctuation">[</span>recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span> <span class="token operator">=</span> i <span class="token operator">+</span> <span class="token number">1</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> matchStr <span class="token punctuation">{</span>
					<span class="token keyword">if</span> headings<span class="token punctuation">[</span>matchStr<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
						line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> matchStr<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
					<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
						p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>

			<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> matchStr <span class="token punctuation">{</span>
					col <span class="token operator">:=</span> headings<span class="token punctuation">[</span>matchStr<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span>
					<span class="token keyword">if</span> col <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> col <span class="token operator">&lt;</span> <span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span><span class="token operator">+</span><span class="token number">1</span> <span class="token punctuation">{</span>
						line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> recs<span class="token punctuation">[</span>col<span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
						<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token operator">||</span> <span class="token punctuation">(</span><span class="token function">len</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">1</span> <span class="token operator">&amp;&amp;</span> recs<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span> <span class="token operator">==</span> <span class="token string">""</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
							<span class="token keyword">continue</span>
						<span class="token punctuation">}</span>
						<span class="token function">errUnmatched</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
					<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
						p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
			lineNum<span class="token operator">++</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">"you haven't selected any rows / columns"</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">ittNot</span><span class="token punctuation">(</span>p <span class="token operator">*</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> cRecords <span class="token keyword">chan</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> marshaller <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">var</span> <span class="token punctuation">(</span>
		mode     <span class="token builtin">int</span>
		matchStr <span class="token operator">=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">bool</span><span class="token punctuation">)</span>
		matchInt <span class="token operator">=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">int</span><span class="token punctuation">]</span><span class="token builtin">bool</span><span class="token punctuation">)</span>
	<span class="token punctuation">)</span>

	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> params <span class="token punctuation">{</span>
		<span class="token keyword">switch</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> rxRowSuffixOld<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byRowNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byRowNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token punctuation">:</span><span class="token function">len</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token operator">-</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt<span class="token punctuation">[</span>num<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token boolean">true</span>

		<span class="token keyword">case</span> rxRowSuffixNew<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byRowNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byRowNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt<span class="token punctuation">[</span>num<span class="token operator">+</span><span class="token number">1</span><span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token boolean">true</span> <span class="token comment">// Don't count from zero</span>

		<span class="token keyword">case</span> rxColumnPrefixOld<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byColumnNumber
			num<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">:=</span> strconv<span class="token punctuation">.</span><span class="token function">Atoi</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">:</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt<span class="token punctuation">[</span>num<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token boolean">true</span>

		<span class="token keyword">case</span> rxColumnPrefixNew<span class="token punctuation">.</span><span class="token function">MatchString</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnNumber <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			mode <span class="token operator">=</span> byColumnNumber
			num <span class="token operator">:=</span> <span class="token function">charToIndex</span><span class="token punctuation">(</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">[</span><span class="token number">1</span><span class="token punctuation">]</span><span class="token punctuation">)</span>
			matchInt<span class="token punctuation">[</span>num<span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token boolean">true</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token keyword">if</span> mode <span class="token operator">!=</span> <span class="token number">0</span> <span class="token operator">&amp;&amp;</span> mode <span class="token operator">!=</span> byColumnName <span class="token punctuation">{</span>
				<span class="token keyword">return</span> errMixAndMatch
			<span class="token punctuation">}</span>
			matchStr<span class="token punctuation">[</span>params<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token boolean">true</span>
			mode <span class="token operator">=</span> byColumnName

		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> mode <span class="token punctuation">{</span>
	<span class="token keyword">case</span> byRowNumber<span class="token punctuation">:</span>
		i <span class="token operator">:=</span> <span class="token operator">-</span><span class="token number">1</span>
		<span class="token keyword">for</span> <span class="token punctuation">{</span>
			recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
			<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
				<span class="token keyword">return</span> <span class="token boolean">nil</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">if</span> <span class="token operator">!</span>matchInt<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token punctuation">{</span>
				<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>recs<span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
					p<span class="token punctuation">.</span>Stderr<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
			i<span class="token operator">++</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">case</span> byColumnNumber<span class="token punctuation">:</span>
		<span class="token keyword">for</span> <span class="token punctuation">{</span>
			recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
			<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
				<span class="token keyword">return</span> <span class="token boolean">nil</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">var</span> line <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
			<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> recs <span class="token punctuation">{</span>
				<span class="token keyword">if</span> <span class="token operator">!</span>matchInt<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token punctuation">{</span>
					line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
			<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
				p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
			<span class="token punctuation">}</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">case</span> byColumnName<span class="token punctuation">:</span>
		<span class="token keyword">var</span> <span class="token punctuation">(</span>
			lineNum  <span class="token builtin">int</span>
			headings <span class="token operator">=</span> <span class="token function">make</span><span class="token punctuation">(</span><span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">int</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span>
		<span class="token punctuation">)</span>

		<span class="token keyword">for</span> <span class="token punctuation">{</span>
			<span class="token keyword">var</span> line <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
			recs<span class="token punctuation">,</span> ok <span class="token operator">:=</span> <span class="token operator">&lt;-</span>cRecords
			<span class="token keyword">if</span> <span class="token operator">!</span>ok <span class="token punctuation">{</span>
				<span class="token keyword">return</span> <span class="token boolean">nil</span>
			<span class="token punctuation">}</span>

			<span class="token keyword">if</span> lineNum <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> recs <span class="token punctuation">{</span>
					headings<span class="token punctuation">[</span>i<span class="token punctuation">]</span> <span class="token operator">=</span> recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span>
					<span class="token keyword">if</span> <span class="token operator">!</span>matchStr<span class="token punctuation">[</span>headings<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>
						line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>
				<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
					p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>

			<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
				<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> recs <span class="token punctuation">{</span>
					<span class="token keyword">if</span> <span class="token operator">!</span>matchStr<span class="token punctuation">[</span>headings<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">]</span> <span class="token punctuation">{</span>
						line <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>line<span class="token punctuation">,</span> recs<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
					<span class="token punctuation">}</span>
				<span class="token punctuation">}</span>

				<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span> <span class="token operator">!=</span> <span class="token number">0</span> <span class="token punctuation">{</span>
					p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token function">marshaller</span><span class="token punctuation">(</span>line<span class="token punctuation">)</span><span class="token punctuation">)</span>
				<span class="token punctuation">}</span>
			<span class="token punctuation">}</span>
			lineNum<span class="token operator">++</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">"you haven't selected any rows / columns"</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>*lang.Process</code>: Process's runtime state. Typically expressed as the variable <code v-pre>p</code></li>
<li><code v-pre>[]string</code>: slice of parameters used in <code v-pre>[</code> / <code v-pre>![</code></li>
<li><code v-pre>chan []string</code>: a channel for rows (each element in the slice is a column within the row). This allows tables to be stream-able</li>
<li><code v-pre>func(interface{}) ([]byte, error)</code>: data type marshaller function</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/ReadArray.html">apis/<code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadArrayWithType.html">apis/<code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
<li><RouterLink to="/apis/ReadIndex.html">apis/<code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/ReadMap.html">apis/<code v-pre>ReadMap()</code> (type)</RouterLink>:
Treat data type as a key/value structure and read its contents</li>
<li><RouterLink to="/apis/ReadNotIndex.html">apis/<code v-pre>ReadNotIndex()</code> (type)</RouterLink>:
Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</li>
<li><RouterLink to="/apis/WriteArray.html">apis/<code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/commands/index2.html">commands/<code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/apis/lang.IndexTemplateObject.html">apis/<code v-pre>lang.IndexTemplateObject()</code> (template API)</RouterLink>:
Returns element(s) from a data structure</li>
</ul>
</div></template>


