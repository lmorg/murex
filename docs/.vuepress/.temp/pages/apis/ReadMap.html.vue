<template><div><h1 id="readmap-type" tabindex="-1"><a class="header-anchor" href="#readmap-type" aria-hidden="true">#</a> <code v-pre>ReadMap()</code> (type)</h1>
<blockquote>
<p>Treat data type as a key/value structure and read its contents</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a function you would write when programming a Murex data-type.</p>
<p>It's called by builtins to allow them to read data structures one key/value
pair at a time.</p>
<p>The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Registering your <code v-pre>ReadMap()</code></p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// To avoid confusion, this should only happen inside func init()</span>
stdio<span class="token punctuation">.</span><span class="token function">RegisterReadMap</span><span class="token punctuation">(</span><span class="token comment">/* your type name */</span><span class="token punctuation">,</span> <span class="token comment">/* your readMap func */</span><span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example <code v-pre>ReadMap()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/config"</span>
	<span class="token string">"github.com/lmorg/murex/lang"</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/lang/types"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readMap</span><span class="token punctuation">(</span>read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> <span class="token boolean">_</span> <span class="token operator">*</span>config<span class="token punctuation">.</span>Config<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token operator">*</span>stdio<span class="token punctuation">.</span>Map<span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token comment">// Create a marshaller function to pass to ArrayWithTypeTemplate</span>
	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> read<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">MapTemplate</span><span class="token punctuation">(</span>types<span class="token punctuation">.</span>Json<span class="token punctuation">,</span> marshaller<span class="token punctuation">,</span> json<span class="token punctuation">.</span>Unmarshal<span class="token punctuation">,</span> read<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>There isn't (yet) a template read function for types to call. However that
might follow in a future release of Murex.</p>
<h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>stdio.Io</code>: stream to read from (eg STDIN)</li>
<li><code v-pre>*config.Config</code>: scoped config (eg your data type might have configurable parsing rules)</li>
<li><code v-pre>func(key, value string, last bool)</code>: callback function: key and value of map plus boolean which is true if last element in row (eg reading from tables rather than key/values)</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/ReadArray.html">apis/<code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
<li><RouterLink to="/apis/ReadArrayWithType.html">apis/<code v-pre>ReadArrayWithType()</code> (type)</RouterLink>:
Read from a data type one array element at a time and return the elements contents and data type</li>
<li><RouterLink to="/apis/ReadIndex.html">apis/<code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/ReadNotIndex.html">apis/<code v-pre>ReadNotIndex()</code> (type)</RouterLink>:
Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</li>
<li><RouterLink to="/apis/WriteArray.html">apis/<code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
</ul>
</div></template>


