import{_ as p}from"./plugin-vue_export-helper-c27b6911.js";import{r as c,o as i,c as l,d as s,b as t,w as e,e as n,f as o}from"./app-45f7c304.js";const u={},r=o(`<h1 id="lang-arraywithtypetemplate-template-api-api-reference" tabindex="-1"><a class="header-anchor" href="#lang-arraywithtypetemplate-template-api-api-reference" aria-hidden="true">#</a> <code>lang.ArrayWithTypeTemplate()</code> (template API) - API Reference</h1><blockquote><p>Unmarshals a data type into a Go struct and returns the results as an array with data type included</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>This is a template API you can use for your custom data types to wrap around an existing Go marshaller and return a Murex array which is consistent with other structures such as nested JSON or YAML documents.</p><p>It should only be called from <code>ReadArrayWithType()</code> functions.</p><p>Because <code>lang.ArrayTemplateWithType()</code> relies on a marshaller, it means any types that rely on this API are not going to be stream-able.</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>Example calling <code>lang.ArrayTemplate()</code> function:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>

	<span class="token string">&quot;github.com/lmorg/murex/lang&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/lang/stdio&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/utils/json&quot;</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readArray</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token comment">// Create a marshaller function to pass to ArrayTemplate</span>
	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> read<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">ArrayTemplate</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> marshaller<span class="token punctuation">,</span> json<span class="token punctuation">.</span>Unmarshal<span class="token punctuation">,</span> read<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="api-source" tabindex="-1"><a class="header-anchor" href="#api-source" aria-hidden="true">#</a> API Source:</h3><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">&quot;context&quot;</span>

	<span class="token string">&quot;github.com/lmorg/murex/lang/stdio&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/lang/types&quot;</span>
	<span class="token string">&quot;github.com/lmorg/murex/utils&quot;</span>
<span class="token punctuation">)</span>

<span class="token comment">// ArrayWithTypeTemplate is a template function for reading arrays from marshalled data</span>
<span class="token keyword">func</span> <span class="token function">ArrayWithTypeTemplate</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> dataType <span class="token builtin">string</span><span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> unmarshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span><span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
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
	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceInterface</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> dataType<span class="token punctuation">,</span> marshal<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceString</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">float64</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceFloat</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceInt</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span><span class="token function">string</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">rune</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span><span class="token function">string</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">bool</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceBool</span><span class="token punctuation">(</span>ctx<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token comment">/*case map[string]string:
		return readArrayWithTypeByMapStrStr(v, callback)

	case map[string]interface{}:
		return readArrayWithTypeByMapStrIface(marshal, v, callback)

	case map[interface{}]string:
		return readArrayWithTypeByMapIfaceStr(v, callback)

	case map[interface{}]interface{}:
		return readArrayWithTypeByMapIfaceIface(marshal, v, callback)
	*/</span>
	<span class="token keyword">default</span><span class="token punctuation">:</span>
		jBytes<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span>
		<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>

			<span class="token keyword">return</span> err
		<span class="token punctuation">}</span>

		<span class="token function">callback</span><span class="token punctuation">(</span>jBytes<span class="token punctuation">,</span> dataType<span class="token punctuation">)</span>

		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span>v <span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceInt</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">int</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Integer<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceFloat</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">float64</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Number<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceBool</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">bool</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Boolean<span class="token punctuation">)</span>

		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceString</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceInterface</span><span class="token punctuation">(</span>ctx context<span class="token punctuation">.</span>Context<span class="token punctuation">,</span> dataType <span class="token builtin">string</span><span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token keyword">select</span> <span class="token punctuation">{</span>
		<span class="token keyword">case</span> <span class="token operator">&lt;-</span>ctx<span class="token punctuation">.</span><span class="token function">Done</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">:</span>
			<span class="token keyword">return</span> <span class="token boolean">nil</span>

		<span class="token keyword">default</span><span class="token punctuation">:</span>
			<span class="token keyword">switch</span> v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>

			<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>

			<span class="token keyword">case</span> <span class="token builtin">float64</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">float64</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Number<span class="token punctuation">)</span>

			<span class="token keyword">case</span> <span class="token builtin">int</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">int</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Integer<span class="token punctuation">)</span>

			<span class="token keyword">case</span> <span class="token builtin">bool</span><span class="token punctuation">:</span>
				<span class="token keyword">if</span> v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">bool</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
					<span class="token function">callback</span><span class="token punctuation">(</span><span class="token boolean">true</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Boolean<span class="token punctuation">)</span>
				<span class="token punctuation">}</span> <span class="token keyword">else</span> <span class="token punctuation">{</span>
					<span class="token function">callback</span><span class="token punctuation">(</span><span class="token boolean">false</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Boolean<span class="token punctuation">)</span>
				<span class="token punctuation">}</span>

			<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span><span class="token function">string</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>

			<span class="token keyword">case</span> <span class="token boolean">nil</span><span class="token punctuation">:</span>
				<span class="token function">callback</span><span class="token punctuation">(</span><span class="token boolean">nil</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>Null<span class="token punctuation">)</span>

			<span class="token keyword">default</span><span class="token punctuation">:</span>
				jBytes<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
				<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
					<span class="token keyword">return</span> err
				<span class="token punctuation">}</span>
				<span class="token function">callback</span><span class="token punctuation">(</span>jBytes<span class="token punctuation">,</span> dataType<span class="token punctuation">)</span>
			<span class="token punctuation">}</span>
		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token comment">/*func readArrayWithTypeByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case &lt;-ctx.Done():
			return nil
		default:
		bKey := []byte(fmt.Sprint(key) + &quot;: &quot;)
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}}

	return nil
}

func readArrayWithTypeByMapStrStr(v map[string]string, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case &lt;-ctx.Done():
			return nil
		default:
		callback([]byte(key + &quot;: &quot; + val))
	}}

	return nil
}

func readArrayWithTypeByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case &lt;-ctx.Done():
			return nil
		default:
		bKey := []byte(key + &quot;: &quot;)
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}}

	return nil
}

func readArrayWithTypeByMapIfaceStr(v map[interface{}]string, callback func([]byte, string)) error {
	for key, val := range v {
	select {
		case &lt;-ctx.Done():
			return nil
		default:
		callback([]byte(fmt.Sprint(key) + &quot;: &quot; + val))
	}}

	return nil
}
*/</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2><ol><li><code>func(interface{}) ([]byte, error)</code>: data type&#39;s marshaller</li><li><code>func([]byte, interface{}) error</code>: data type&#39;s unmarshaller</li><li><code>stdio.Io</code>: stream to read from (eg STDIN)</li><li><code>func(interface{}, string)</code>: callback function to write each array element, with data type</li></ol><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),k=s("code",null,"ReadArray()",-1),d=s("code",null,"ReadArrayWithType()",-1),v=s("code",null,"ReadIndex()",-1),b=s("code",null,"[",-1),m=s("code",null,"ReadMap()",-1),y=s("code",null,"ReadNotIndex()",-1),f=s("code",null,"![",-1),h=s("code",null,"WriteArray()",-1),w=s("code",null,"lang.IndexTemplateObject()",-1),g=s("code",null,"lang.IndexTemplateTable()",-1);function x(T,A){const a=c("RouterLink");return i(),l("div",null,[r,s("ul",null,[s("li",null,[t(a,{to:"/apis/ReadArray.html"},{default:e(()=>[n("apis/"),k,n(" (type)")]),_:1}),n(": Read from a data type one array element at a time")]),s("li",null,[t(a,{to:"/apis/ReadArrayWithType.html"},{default:e(()=>[n("apis/"),d,n(" (type)")]),_:1}),n(": Read from a data type one array element at a time and return the elements contents and data type")]),s("li",null,[t(a,{to:"/apis/ReadIndex.html"},{default:e(()=>[n("apis/"),v,n(" (type)")]),_:1}),n(": Data type handler for the index, "),b,n(", builtin")]),s("li",null,[t(a,{to:"/apis/ReadMap.html"},{default:e(()=>[n("apis/"),m,n(" (type)")]),_:1}),n(": Treat data type as a key/value structure and read its contents")]),s("li",null,[t(a,{to:"/apis/ReadNotIndex.html"},{default:e(()=>[n("apis/"),y,n(" (type)")]),_:1}),n(": Data type handler for the bang-prefixed index, "),f,n(", builtin")]),s("li",null,[t(a,{to:"/apis/WriteArray.html"},{default:e(()=>[n("apis/"),h,n(" (type)")]),_:1}),n(": Write a data type, one array element at a time")]),s("li",null,[t(a,{to:"/apis/lang.IndexTemplateObject.html"},{default:e(()=>[n("apis/"),w,n(" (template API)")]),_:1}),n(": Returns element(s) from a data structure")]),s("li",null,[t(a,{to:"/apis/lang.IndexTemplateTable.html"},{default:e(()=>[n("apis/"),g,n(" (template API)")]),_:1}),n(": Returns element(s) from a table")])])])}const I=p(u,[["render",x],["__file","lang.ArrayWithTypeTemplate.html.vue"]]);export{I as default};
