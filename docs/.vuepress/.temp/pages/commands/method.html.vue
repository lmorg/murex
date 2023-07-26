<template><div><h1 id="method" tabindex="-1"><a class="header-anchor" href="#method" aria-hidden="true">#</a> <code v-pre>method</code></h1>
<blockquote>
<p>Define a methods supported data-types</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>method</code> defines what the typical data type would be for a function's STDIN
and STDOUT.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<pre><code>method: define name { json }
</code></pre>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<pre><code>method: define name {
    &quot;Stdin&quot;:  &quot;@Any&quot;,
    &quot;Stdout&quot;: &quot;json&quot;
}
</code></pre>
<h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="type-groups" tabindex="-1"><a class="header-anchor" href="#type-groups" aria-hidden="true">#</a> Type Groups</h3>
<p>You can define a Murex data type or use a type group. The following type
groups are available to use:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> types

<span class="token comment">// These are the different supported type groups</span>
<span class="token keyword">const</span> <span class="token punctuation">(</span>
	Any               <span class="token operator">=</span> <span class="token string">"@Any"</span>
	Text              <span class="token operator">=</span> <span class="token string">"@Text"</span>
	Math              <span class="token operator">=</span> <span class="token string">"@Math"</span>
	Unmarshal         <span class="token operator">=</span> <span class="token string">"@Unmarshal"</span>
	Marshal           <span class="token operator">=</span> <span class="token string">"@Marshal"</span>
	ReadArray         <span class="token operator">=</span> <span class="token string">"@ReadArray"</span>
	ReadArrayWithType <span class="token operator">=</span> <span class="token string">"@ReadArrayWithType"</span>
	WriteArray        <span class="token operator">=</span> <span class="token string">"@WriteArray"</span>
	ReadIndex         <span class="token operator">=</span> <span class="token string">"@ReadIndex"</span>
	ReadNotIndex      <span class="token operator">=</span> <span class="token string">"@ReadNotIndex"</span>
	ReadMap           <span class="token operator">=</span> <span class="token string">"@ReadMap"</span>
<span class="token punctuation">)</span>

<span class="token comment">// GroupText is an array of the data types that make up the `text` type</span>
<span class="token keyword">var</span> GroupText <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	Generic<span class="token punctuation">,</span>
	String<span class="token punctuation">,</span>
	<span class="token string">`generic`</span><span class="token punctuation">,</span>
	<span class="token string">`string`</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token comment">// GroupMath is an array of the data types that make up the `math` type</span>
<span class="token keyword">var</span> GroupMath <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	Number<span class="token punctuation">,</span>
	Integer<span class="token punctuation">,</span>
	Float<span class="token punctuation">,</span>
	Boolean<span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/parser/pipe-arrow.html">Arrow Pipe (<code v-pre>-&gt;</code>) Token</RouterLink>:
Pipes STDOUT from the left hand command to STDIN of the right hand command</li>
<li><RouterLink to="/user-guide/interactive-shell.html">Murex's Interactive Shell</RouterLink>:
What's different about Murex's interactive shell?</li>
<li><RouterLink to="/commands/alias.html"><code v-pre>alias</code></RouterLink>:
Create an alias for a command</li>
<li><RouterLink to="/commands/autocomplete.html"><code v-pre>autocomplete</code></RouterLink>:
Set definitions for tab-completion in the command line</li>
<li><RouterLink to="/commands/function.html"><code v-pre>function</code></RouterLink>:
Define a function block</li>
<li><RouterLink to="/commands/private.html"><code v-pre>private</code></RouterLink>:
Define a private function block</li>
<li><RouterLink to="/commands/runtime.html"><code v-pre>runtime</code></RouterLink>:
Returns runtime information on the internal state of Murex</li>
</ul>
</div></template>


