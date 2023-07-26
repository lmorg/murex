<template><div><h1 id="readarraywithtype-type" tabindex="-1"><a class="header-anchor" href="#readarraywithtype-type" aria-hidden="true">#</a> <code v-pre>ReadArrayWithType()</code> (type)</h1>
<blockquote>
<p>Read from a data type one array element at a time and return the elements contents and data type</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a function you would write when programming a Murex data-type.</p>
<p>It's called by builtins to allow them to read data structures one array element
at a time.</p>
<p>The purpose of this function is to allow builtins to support sequential reads
(where possible) and also create a standard interface for builtins, thus
allowing them to be data-type agnostic.</p>
<p>This differs from ReadArray() because it also returns the data type.</p>
<p>There is a good chance ReadArray() might get deprecated in the medium to long
term.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<p>Registering your <code v-pre>ReadArrayWithType()</code></p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token comment">// To avoid confusion, this should only happen inside func init()</span>
stdio<span class="token punctuation">.</span><span class="token function">RegisterReadArrayWithType</span><span class="token punctuation">(</span><span class="token comment">/* your type name */</span><span class="token punctuation">,</span> <span class="token comment">/* your readArray func */</span><span class="token punctuation">)</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example <code v-pre>ReadArrayWithType()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> <span class="token builtin">string</span>

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"bufio"</span>
	<span class="token string">"context"</span>
	<span class="token string">"fmt"</span>
	<span class="token string">"strings"</span>

	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/lang/types"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithType</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	scanner <span class="token operator">:=</span> bufio<span class="token punctuation">.</span><span class="token function">NewScanner</span><span class="token punctuation">(</span>read<span class="token punctuation">)</span>
	<span class="token keyword">for</span> scanner<span class="token punctuation">.</span><span class="token function">Scan</span><span class="token punctuation">(</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> scanner<span class="token punctuation">.</span><span class="token function">Err</span><span class="token punctuation">(</span><span class="token punctuation">)</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>strings<span class="token punctuation">.</span><span class="token function">TrimSpace</span><span class="token punctuation">(</span>scanner<span class="token punctuation">.</span><span class="token function">Text</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	err <span class="token operator">:=</span> scanner<span class="token punctuation">.</span><span class="token function">Err</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> fmt<span class="token punctuation">.</span><span class="token function">Errorf</span><span class="token punctuation">(</span><span class="token string">"error while reading a %s array: %s"</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">,</span> err<span class="token punctuation">.</span><span class="token function">Error</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>
	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<p>If your data type is not a stream-able array, it is then recommended that
you pass your array to <code v-pre>lang.ArrayWithTypeTemplate()</code> which is a handler to
convert Go structures into Murex arrays. This also makes writing <code v-pre>ReadArray()</code>
handlers easier since you can just pass <code v-pre>lang.ArrayTemplate()</code> your marshaller.
For example:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"context"</span>

	<span class="token string">"github.com/lmorg/murex/lang"</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/lang/types"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithType</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token comment">// Create a marshaller function to pass to ArrayWithTypeTemplate</span>
	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> read<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">ArrayWithTypeTemplate</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> types<span class="token punctuation">.</span>Json<span class="token punctuation">,</span> marshaller<span class="token punctuation">,</span> json<span class="token punctuation">.</span>Unmarshal<span class="token punctuation">,</span> read<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>The downside of this is that you're then unmarshalling the entire file, which
could be slow on large files and also breaks the streaming nature of UNIX
pipelines.</p>
<h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>stdio.Io</code>: stream to read from (eg STDIN)</li>
<li><code v-pre>func(interface{}, string)</code>: callback function. Each callback will be the value in its native Go data type (eg string, int, float64, bool) for an array element</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/ReadIndex.html">apis/<code v-pre>ReadIndex()</code> (type)</RouterLink>:
Data type handler for the index, <code v-pre>[</code>, builtin</li>
<li><RouterLink to="/apis/ReadMap.html">apis/<code v-pre>ReadMap()</code> (type)</RouterLink>:
Treat data type as a key/value structure and read its contents</li>
<li><RouterLink to="/apis/ReadNotIndex.html">apis/<code v-pre>ReadNotIndex()</code> (type)</RouterLink>:
Data type handler for the bang-prefixed index, <code v-pre>![</code>, builtin</li>
<li><RouterLink to="/apis/WriteArray.html">apis/<code v-pre>WriteArray()</code> (type)</RouterLink>:
Write a data type, one array element at a time</li>
<li><RouterLink to="/apis/lang.ArrayTemplate.html">apis/<code v-pre>lang.ArrayTemplate()</code> (template API)</RouterLink>:
Unmarshals a data type into a Go struct and returns the results as an array</li>
<li><RouterLink to="/apis/lang.ArrayWithTypeTemplate.html">apis/<code v-pre>lang.ArrayWithTypeTemplate()</code> (template API)</RouterLink>:
Unmarshals a data type into a Go struct and returns the results as an array with data type included</li>
</ul>
</div></template>


