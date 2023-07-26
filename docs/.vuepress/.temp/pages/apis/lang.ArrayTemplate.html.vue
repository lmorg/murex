<template><div><h1 id="lang-arraytemplate-template-api" tabindex="-1"><a class="header-anchor" href="#lang-arraytemplate-template-api" aria-hidden="true">#</a> <code v-pre>lang.ArrayTemplate()</code> (template API)</h1>
<blockquote>
<p>Unmarshals a data type into a Go struct and returns the results as an array</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a template API you can use for your custom data types to wrap around an
existing Go marshaller and return a Murex array which is consistent with
other structures such as nested JSON or YAML documents.</p>
<p>It should only be called from <code v-pre>ReadArray()</code> functions.</p>
<p>Because <code v-pre>lang.ArrayTemplate()</code> relies on a marshaller, it means any types that
rely on this API are not going to be stream-able.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example calling <code v-pre>lang.ArrayTemplate()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"context"</span>

	<span class="token string">"github.com/lmorg/murex/lang"</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readArray</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token comment">// Create a marshaller function to pass to ArrayTemplate</span>
	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> read<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">ArrayTemplate</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> marshaller<span class="token punctuation">,</span> json<span class="token punctuation">.</span>Unmarshal<span class="token punctuation">,</span> read<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="api-source" tabindex="-1"><a class="header-anchor" href="#api-source" aria-hidden="true">#</a> API Source:</h3>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"context"</span>
	<span class="token string">"fmt"</span>

	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/utils"</span>
<span class="token punctuation">)</span>

<span class="token comment">// ArrayTemplate is a template function for reading arrays from marshalled data</span>
<span class="token keyword">func</span> <span class="token function">ArrayTemplate</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> unmarshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span><span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> read<span class="token punctuation">.</span><span class="token function">ReadAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>utils<span class="token punctuation">.</span><span class="token function">CrLfTrim</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">var</span> v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span>
	err <span class="token operator">=</span> <span class="token function">unmarshal</span><span class="token punctuation">(</span>b<span class="token punctuation">,</span> <span class="token operator">&amp;</span>v<span class="token punctuation">)</span>

	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> v <span class="token operator">:=</span> v<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayByString</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayBySliceString</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayBySliceInterface</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> marshal<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayByMapStrStr</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayByMapStrIface</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> marshal<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayByMapIfaceStr</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayByMapIfaceIface</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> marshal<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		jBytes<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>

			<span class="token keyword">return</span> err
		<span class="token punctuation">}</span>

		<span class="token function">callback</span><span class="token punctuation">(</span>jBytes<span class="token punctuation">)</span>

		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayByString</span><span class="token punctuation">(</span>v <span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token punctuation">)</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayBySliceString</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayBySliceInterface</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token keyword">switch</span> v <span class="token operator">:=</span> v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
			<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token punctuation">)</span>

			<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span>

			<span class="token keyword">default</span><span class="token punctuation">:</span>
				jBytes<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span>
				<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
					<span class="token keyword">return</span> err
				<span class="token punctuation">}</span>
				<span class="token function">callback</span><span class="token punctuation">(</span>jBytes<span class="token punctuation">)</span>
			<span class="token punctuation">}</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayByMapIfaceIface</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> v <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> key<span class="token punctuation">,</span> val <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			bKey <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>fmt<span class="token punctuation">.</span><span class="token function">Sprint</span><span class="token punctuation">(</span>key<span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token string">": "</span><span class="token punctuation">)</span>
			b<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>val<span class="token punctuation">)</span>
			<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
				<span class="token keyword">return</span> err
			<span class="token punctuation">}</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token function">append</span><span class="token punctuation">(</span>bKey<span class="token punctuation">,</span> b<span class="token operator">...</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayByMapStrStr</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> key<span class="token punctuation">,</span> val <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>key <span class="token operator">+</span> <span class="token string">": "</span> <span class="token operator">+</span> val<span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayByMapStrIface</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> v <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> key<span class="token punctuation">,</span> val <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			bKey <span class="token operator">:=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>key <span class="token operator">+</span> <span class="token string">": "</span><span class="token punctuation">)</span>
			b<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>val<span class="token punctuation">)</span>
			<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
				<span class="token keyword">return</span> err
			<span class="token punctuation">}</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token function">append</span><span class="token punctuation">(</span>bKey<span class="token punctuation">,</span> b<span class="token operator">...</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayByMapIfaceStr</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> key<span class="token punctuation">,</span> val <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>fmt<span class="token punctuation">.</span><span class="token function">Sprint</span><span class="token punctuation">(</span>key<span class="token punctuation">)</span> <span class="token operator">+</span> <span class="token string">": "</span> <span class="token operator">+</span> val<span class="token punctuation">)</span><span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>func(interface{}) ([]byte, error)</code>: data type's marshaller</li>
<li><code v-pre>func([]byte, interface{}) error</code>: data type's unmarshaller</li>
<li><code v-pre>stdio.Io</code>: stream to read from (eg STDIN)</li>
<li><code v-pre>func([]byte)</code>: callback function to write each array element</li>
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
<li><RouterLink to="/apis/lang.IndexTemplateObject.html">apis/<code v-pre>lang.IndexTemplateObject()</code> (template API)</RouterLink>:
Returns element(s) from a data structure</li>
<li><RouterLink to="/apis/lang.IndexTemplateTable.html">apis/<code v-pre>lang.IndexTemplateTable()</code> (template API)</RouterLink>:
Returns element(s) from a table</li>
</ul>
</div></template>


