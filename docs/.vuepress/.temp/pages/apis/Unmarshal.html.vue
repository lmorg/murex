<template><div><h1 id="unmarshal-type" tabindex="-1"><a class="header-anchor" href="#unmarshal-type" aria-hidden="true">#</a> <code v-pre>Unmarshal()</code> (type)</h1>
<blockquote>
<p>Converts a structured file format into structured memory</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a function you would write when programming a Murex data-type.
The unmarshal function takes in a byte slice and returns a Go (golang)
<code v-pre>type</code> or <code v-pre>struct</code> or an error.</p>
<p>This unmarshaller is then registered to Murex inside an <code v-pre>init()</code> function
and Murex builtins can use that unmarshaller via the <code v-pre>UnmarshalData()</code>
API.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Registering <code v-pre>Unmarshal()</code> (for writing builtin data-types)</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// To avoid data races, this should only happen inside func init()</span>
lang<span class="token punctuation">.</span>Unmarshallers<span class="token punctuation">[</span> <span class="token comment">/* your type name */</span> <span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token comment">/* your readIndex func */</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><p>Using an existing unmarshaller (eg inside a builtin command)</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// See documentation on lang.UnmarshalData for more details</span>
v<span class="token punctuation">,</span> err <span class="token operator">:=</span> lang<span class="token punctuation">.</span><span class="token function">UnmarshalData</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">,</span> dataType <span class="token builtin">string</span><span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Defining a marshaller for a murex data-type</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> example

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"encoding/json"</span>

	<span class="token string">"github.com/lmorg/murex/lang"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">init</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// Register data-type</span>
	lang<span class="token punctuation">.</span>Unmarshallers<span class="token punctuation">[</span><span class="token string">"example"</span><span class="token punctuation">]</span> <span class="token operator">=</span> unmarshal
<span class="token punctuation">}</span>

<span class="token comment">// Describe unmarshaller</span>
<span class="token keyword">func</span> <span class="token function">unmarshal</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token comment">// Read data from STDIN. Because JSON expects closing tokens, we should</span>
	<span class="token comment">// read the entire stream before unmarshalling it. For formats like CSV or</span>
	<span class="token comment">// jsonlines which are more line based, we might want to read STDIN line by</span>
	<span class="token comment">// line. However given there is just one data return, you still effectively</span>
	<span class="token comment">// head to read the entire file before returning the structure. There are</span>
	<span class="token comment">// other APIs for iterative returns for streaming data - more akin to the</span>
	<span class="token comment">// traditional way UNIX pipes would work.</span>
	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdin<span class="token punctuation">.</span><span class="token function">ReadAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span><span class="token punctuation">,</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">var</span> v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span>
	err <span class="token operator">=</span> json<span class="token punctuation">.</span><span class="token function">Unmarshal</span><span class="token punctuation">(</span>b<span class="token punctuation">,</span> <span class="token operator">&amp;</span>v<span class="token punctuation">)</span>

	<span class="token comment">// Return the Go data structure or error</span>
	<span class="token keyword">return</span> v<span class="token punctuation">,</span> err
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>*lang.Process</code>: Process's runtime state. Typically expressed as the variable <code v-pre>p</code></li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/Marshal.html">apis/<code v-pre>Marshal()</code> (type)</RouterLink>:
Converts structured memory into a structured file format (eg for stdio)</li>
<li><RouterLink to="/apis/lang.MarshalData.html">apis/<code v-pre>lang.MarshalData()</code> (system API)</RouterLink>:
Converts structured memory into a Murex data-type (eg for stdio)</li>
<li><RouterLink to="/apis/lang.UnmarshalData.html">apis/<code v-pre>lang.UnmarshalData()</code> (system API)</RouterLink>:
Converts a Murex data-type into structured memory</li>
</ul>
</div></template>


