<template><div><h1 id="writearray-type" tabindex="-1"><a class="header-anchor" href="#writearray-type" aria-hidden="true">#</a> <code v-pre>WriteArray()</code> (type)</h1>
<blockquote>
<p>Write a data type, one array element at a time</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a function you would write when programming a Murex data-type.</p>
<p>It's called by builtins to allow them to write data structures one array
element at a time.</p>
<p>The purpose of this function is to allow builtins to support sequential writes
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.</p>
<h3 id="a-collection-of-functions" tabindex="-1"><a class="header-anchor" href="#a-collection-of-functions" aria-hidden="true">#</a> A Collection of Functions</h3>
<p><code v-pre>WriteArray()</code> should return a <code v-pre>struct</code> that satisfies the following
<code v-pre>interface{}</code>:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> stdio

<span class="token comment">// ArrayWriter is a simple interface types can adopt for buffered writes of formatted arrays in structured types (eg JSON)</span>
<span class="token keyword">type</span> ArrayWriter <span class="token keyword">interface</span> <span class="token punctuation">{</span>
	<span class="token function">Write</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">error</span>
	<span class="token function">WriteString</span><span class="token punctuation">(</span><span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span>
	<span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">error</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Registering your <code v-pre>WriteArray()</code></p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// To avoid confusion, this should only happen inside func init()</span>
stdio<span class="token punctuation">.</span><span class="token function">RegisterWriteArray</span><span class="token punctuation">(</span><span class="token comment">/* your type name */</span><span class="token punctuation">,</span> <span class="token comment">/* your writeArray func */</span><span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example <code v-pre>WriteArray()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> <span class="token builtin">string</span>

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
<span class="token punctuation">)</span>

<span class="token keyword">type</span> arrayWriter <span class="token keyword">struct</span> <span class="token punctuation">{</span>
	writer stdio<span class="token punctuation">.</span>Io
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">newArrayWriter</span><span class="token punctuation">(</span>writer stdio<span class="token punctuation">.</span>Io<span class="token punctuation">)</span> <span class="token punctuation">(</span>stdio<span class="token punctuation">.</span>ArrayWriter<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	w <span class="token operator">:=</span> <span class="token operator">&amp;</span>arrayWriter<span class="token punctuation">{</span>writer<span class="token punctuation">:</span> writer<span class="token punctuation">}</span>
	<span class="token keyword">return</span> w<span class="token punctuation">,</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">Write</span><span class="token punctuation">(</span>b <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">:=</span> w<span class="token punctuation">.</span>writer<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span>
	<span class="token keyword">return</span> err
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">WriteString</span><span class="token punctuation">(</span>s <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">:=</span> w<span class="token punctuation">.</span>writer<span class="token punctuation">.</span><span class="token function">Writeln</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>s<span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span> err
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span> <span class="token keyword">return</span> <span class="token boolean">nil</span> <span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>Since not all data types will be stream-able (for example <code v-pre>json</code>), some types
may need to cache the array and then to write it once the array writer has been
closed.</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">type</span> arrayWriter <span class="token keyword">struct</span> <span class="token punctuation">{</span>
	array  <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span>
	writer stdio<span class="token punctuation">.</span>Io
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">newArrayWriter</span><span class="token punctuation">(</span>writer stdio<span class="token punctuation">.</span>Io<span class="token punctuation">)</span> <span class="token punctuation">(</span>stdio<span class="token punctuation">.</span>ArrayWriter<span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	w <span class="token operator">:=</span> <span class="token operator">&amp;</span>arrayWriter<span class="token punctuation">{</span>writer<span class="token punctuation">:</span> writer<span class="token punctuation">}</span>
	<span class="token keyword">return</span> w<span class="token punctuation">,</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">Write</span><span class="token punctuation">(</span>b <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	w<span class="token punctuation">.</span>array <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>w<span class="token punctuation">.</span>array<span class="token punctuation">,</span> <span class="token function">string</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">WriteString</span><span class="token punctuation">(</span>s <span class="token builtin">string</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	w<span class="token punctuation">.</span>array <span class="token operator">=</span> <span class="token function">append</span><span class="token punctuation">(</span>w<span class="token punctuation">.</span>array<span class="token punctuation">,</span> s<span class="token punctuation">)</span>
	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token punctuation">(</span>w <span class="token operator">*</span>arrayWriter<span class="token punctuation">)</span> <span class="token function">Close</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>w<span class="token punctuation">.</span>array<span class="token punctuation">,</span> w<span class="token punctuation">.</span>writer<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token boolean">_</span><span class="token punctuation">,</span> err <span class="token operator">=</span> w<span class="token punctuation">.</span>writer<span class="token punctuation">.</span><span class="token function">Write</span><span class="token punctuation">(</span>b<span class="token punctuation">)</span>
	<span class="token keyword">return</span> err
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
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
</ul>
</div></template>


