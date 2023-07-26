<template><div><h1 id="readnotindex-type" tabindex="-1"><a class="header-anchor" href="#readnotindex-type" aria-hidden="true">#</a> <code v-pre>ReadNotIndex()</code> (type)</h1>
<blockquote>
<p>Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a function you would write when programming a Murex data-type.</p>
<p>It's called by the index, <code v-pre>![</code>, builtin.</p>
<p>The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for <code v-pre>![</code> (index), thus
allowing it to be data-type agnostic.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Registering your <code v-pre>ReadNotIndex()</code></p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// To avoid data races, this should only happen inside func init()</span>
lang<span class="token punctuation">.</span>ReadNotIndexes<span class="token punctuation">[</span> <span class="token comment">/* your type name */</span> <span class="token punctuation">]</span> <span class="token operator">=</span> <span class="token comment">/* your readIndex func */</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example <code v-pre>ReadIndex()</code> function (the code structure is the same for <code v-pre>ReadIndex</code>
and <code v-pre>ReadNotIndex</code>):</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/lang"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">index</span><span class="token punctuation">(</span>p <span class="token operator">*</span>lang<span class="token punctuation">.</span>Process<span class="token punctuation">,</span> params <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">var</span> jInterface <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span>

	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> p<span class="token punctuation">.</span>Stdin<span class="token punctuation">.</span><span class="token function">ReadAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	err <span class="token operator">=</span> json<span class="token punctuation">.</span><span class="token function">Unmarshal</span><span class="token punctuation">(</span>b<span class="token punctuation">,</span> <span class="token operator">&amp;</span>jInterface<span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>iface <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>iface<span class="token punctuation">,</span> p<span class="token punctuation">.</span>Stdout<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">IndexTemplateObject</span><span class="token punctuation">(</span>p<span class="token punctuation">,</span> params<span class="token punctuation">,</span> <span class="token operator">&amp;</span>jInterface<span class="token punctuation">,</span> marshaller<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>While there is support for a dedicated <code v-pre>ReadNotIndex()</code> for instances of <code v-pre>![</code>,
the template APIs <code v-pre>lang.IndexTemplateObject</code> and <code v-pre>lang.IndexTemplateTable</code> are
both agnostic to the bang prefix.</p>
<h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>*lang.Process</code>: Process's runtime state. Typically expressed as the variable <code v-pre>p</code></li>
<li><code v-pre>[]string</code>: slice of parameters used in <code v-pre>![</code></li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/user-guide/bang-prefix.html">user-guide/Bang Prefix</RouterLink>:
Bang prefixing to reverse default actions</li>
<li><RouterLink to="/apis/ReadArray.html">apis/<code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadArrayWithType.html">apis/<code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
<li><RouterLink to="/apis/ReadIndex.html">apis/<code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/WriteArray.html">apis/<code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/commands/element.html">commands/<code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html">commands/<code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/apis/lang.IndexTemplateObject.html">apis/<code v-pre>lang.IndexTemplateObject()</code> (template API)</RouterLink>:
Returns element(s) from a data structure</li>
<li><RouterLink to="/apis/lang.IndexTemplateTable.html">apis/<code v-pre>lang.IndexTemplateTable()</code> (template API)</RouterLink>:
Returns element(s) from a table</li>
</ul>
</div></template>


