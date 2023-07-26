<template><div><h1 id="lang-marshaldata-system-api" tabindex="-1"><a class="header-anchor" href="#lang-marshaldata-system-api" aria-hidden="true">#</a> <code v-pre>lang.MarshalData()</code> (system API)</h1>
<blockquote>
<p>Converts structured memory into a Murex data-type (eg for stdio)</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code>b<span class="token punctuation">,</span> err <span class="token operator">:=</span> lang<span class="token punctuation">.</span><span class="token function">MarshalData</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> dataType<span class="token punctuation">,</span> data<span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">func</span> <span class="token function">exampleCommand</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
    data <span class="token operator">:=</span> <span class="token keyword">map</span><span class="token punctuation">[</span><span class="token builtin">string</span><span class="token punctuation">]</span><span class="token builtin">string</span> <span class="token punctuation">{</span>
        <span class="token string">"foo"</span><span class="token punctuation">:</span> <span class="token string">"hello foo"</span><span class="token punctuation">,</span>
        <span class="token string">"bar"</span><span class="token punctuation">:</span> <span class="token string">"hello bar"</span><span class="token punctuation">,</span>
    <span class="token punctuation">}</span>

    dataType <span class="token operator">:=</span> <span class="token string">"json"</span>

    b<span class="token punctuation">,</span> err <span class="token operator">:=</span> lang<span class="token punctuation">.</span><span class="token function">MarshalData</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> dataType<span class="token punctuation">,</span> data<span class="token punctuation">)</span>
    <span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
        <span class="token keyword">return</span> err
    <span class="token punctuation">}</span>

    <span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span>
    <span class="token keyword">return</span> err
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Go source file:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"errors"</span>
<span class="token punctuation">)</span>

<span class="token comment">// MarshalData is a global marshaller which should be called from within murex</span>
<span class="token comment">// builtin commands (etc).</span>
<span class="token comment">// See docs/apis/marshaldata.md for more details</span>
<span class="token keyword">func</span> <span class="token function">MarshalData</span><span class="token punctuation">(</span>p <span class="token operator">*</span>Process<span class="token punctuation">,</span> dataType <span class="token builtin">string</span><span class="token punctuation">,</span> data <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span>b <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> err <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// This is one of the very few maps in Murex which isn't hidden behind a sync</span>
	<span class="token comment">// lock of one description or other. The rational is that even mutexes can</span>
	<span class="token comment">// add a noticeable overhead on the performance of tight loops and I expect</span>
	<span class="token comment">// this function to be called _a lot_ while also only needing to be written</span>
	<span class="token comment">// to via code residing in within builtin types init() function (ie while</span>
	<span class="token comment">// murex is effectively single threaded). So there shouldn't be any data-</span>
	<span class="token comment">// races -- PROVIDING developers strictly follow the pattern of only writing</span>
	<span class="token comment">// to this map within init() func's.</span>
	<span class="token keyword">if</span> Marshallers<span class="token punctuation">[</span>dataType<span class="token punctuation">]</span> <span class="token operator">==</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span><span class="token punctuation">,</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">"I don't know how to marshal `"</span> <span class="token operator">+</span> dataType <span class="token operator">+</span> <span class="token string">"`."</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	b<span class="token punctuation">,</span> err <span class="token operator">=</span> Marshallers<span class="token punctuation">[</span>dataType<span class="token punctuation">]</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> data<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span><span class="token punctuation">,</span> errors<span class="token punctuation">.</span><span class="token function">New</span><span class="token punctuation">(</span><span class="token string">"["</span> <span class="token operator">+</span> dataType <span class="token operator">+</span> <span class="token string">" marshaller] "</span> <span class="token operator">+</span> err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>*lang.Process</code>: Process's runtime state. Typically expressed as the variable <code v-pre>p</code></li>
<li><code v-pre>string</code>: Murex data type</li>
<li><code v-pre>interface{}</code>: data you wish to marshal</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/Marshal.html">apis/<code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/Unmarshal.html">apis/<code v-pre>Unmarshal()</code> (type)</RouterLink>:
Converts a structured file format into structured memory</li>
<li><RouterLink to="/apis/lang.UnmarshalData.html">apis/<code v-pre>lang.UnmarshalData()</code> (system API)</RouterLink>:
Converts a Murex data-type into structured memory</li>
</ul>
</div></template>


