<template><div><h1 id="murex-shell-docs" tabindex="-1"><a class="header-anchor" href="#murex-shell-docs" aria-hidden="true">#</a> Murex Shell Docs</h1>
<h2 id="api-reference-lang-arraytemplatewithtype-template-api" tabindex="-1"><a class="header-anchor" href="#api-reference-lang-arraytemplatewithtype-template-api" aria-hidden="true">#</a> API Reference: <code v-pre>lang.ArrayTemplateWithType()</code> (template API)</h2>
<blockquote>
<p>Unmarshals a data type into a Go struct and returns the results as an array with data type included</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p>This is a template API you can use for your custom data types to wrap around an
existing Go marshaller and return a Murex array which is consistent with
other structures such as nested JSON or YAML documents.</p>
<p>It should only be called from <code v-pre>ReadArrayWithType()</code> functions.</p>
<p>Because <code v-pre>lang.ArrayTemplateWithType()</code> relies on a marshaller, it means any types that
rely on this API are not going to be stream-able.</p>
<h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>Example calling <code v-pre>lang.ArrayTemplate()</code> function:</p>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> json

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/lang"</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/utils/json"</span>
<span class="token punctuation">)</span>

<span class="token keyword">func</span> <span class="token function">readArray</span><span class="token punctuation">(</span>read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token comment">// Create a marshaller function to pass to ArrayTemplate</span>
	marshaller <span class="token operator">:=</span> <span class="token keyword">func</span><span class="token punctuation">(</span>v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> json<span class="token punctuation">.</span><span class="token function">Marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> read<span class="token punctuation">.</span><span class="token function">IsTTY</span><span class="token punctuation">(</span><span class="token punctuation">)</span><span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> lang<span class="token punctuation">.</span><span class="token function">ArrayTemplate</span><span class="token punctuation">(</span>marshaller<span class="token punctuation">,</span> json<span class="token punctuation">.</span>Unmarshal<span class="token punctuation">,</span> read<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2>
<h3 id="api-source" tabindex="-1"><a class="header-anchor" href="#api-source" aria-hidden="true">#</a> API Source:</h3>
<div class="language-go line-numbers-mode" data-ext="go"><pre v-pre class="language-go"><code><span class="token keyword">package</span> lang

<span class="token keyword">import</span> <span class="token punctuation">(</span>
	<span class="token string">"github.com/lmorg/murex/lang/stdio"</span>
	<span class="token string">"github.com/lmorg/murex/lang/types"</span>
<span class="token punctuation">)</span>

<span class="token comment">// ArrayWithTypeTemplate is a template function for reading arrays from marshalled data</span>
<span class="token keyword">func</span> <span class="token function">ArrayWithTypeTemplate</span><span class="token punctuation">(</span>dataType <span class="token builtin">string</span><span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> unmarshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token builtin">error</span><span class="token punctuation">,</span> read stdio<span class="token punctuation">.</span>Io<span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	b<span class="token punctuation">,</span> err <span class="token operator">:=</span> read<span class="token punctuation">.</span><span class="token function">ReadAll</span><span class="token punctuation">(</span><span class="token punctuation">)</span>
	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">var</span> v <span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span>
	err <span class="token operator">=</span> <span class="token function">unmarshal</span><span class="token punctuation">(</span>b<span class="token punctuation">,</span> <span class="token operator">&amp;</span>v<span class="token punctuation">)</span>

	<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> err
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> v <span class="token operator">:=</span> v<span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceString</span><span class="token punctuation">(</span>v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">:</span>
		<span class="token keyword">return</span> <span class="token function">readArrayWithTypeBySliceInterface</span><span class="token punctuation">(</span>dataType<span class="token punctuation">,</span> marshal<span class="token punctuation">,</span> v<span class="token punctuation">,</span> callback<span class="token punctuation">)</span>

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

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeByString</span><span class="token punctuation">(</span>v <span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceString</span><span class="token punctuation">(</span>v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
		<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token keyword">func</span> <span class="token function">readArrayWithTypeBySliceInterface</span><span class="token punctuation">(</span>dataType <span class="token builtin">string</span><span class="token punctuation">,</span> marshal <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">)</span> <span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">error</span><span class="token punctuation">)</span><span class="token punctuation">,</span> v <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token keyword">interface</span><span class="token punctuation">{</span><span class="token punctuation">}</span><span class="token punctuation">,</span> callback <span class="token keyword">func</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">,</span> <span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span> <span class="token builtin">error</span> <span class="token punctuation">{</span>
	<span class="token keyword">if</span> <span class="token function">len</span><span class="token punctuation">(</span>v<span class="token punctuation">)</span> <span class="token operator">==</span> <span class="token number">0</span> <span class="token punctuation">{</span>
		<span class="token keyword">return</span> <span class="token boolean">nil</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">switch</span> v<span class="token punctuation">[</span><span class="token number">0</span><span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token keyword">type</span><span class="token punctuation">)</span> <span class="token punctuation">{</span>
	<span class="token keyword">case</span> <span class="token builtin">string</span><span class="token punctuation">:</span>
		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
			<span class="token function">callback</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token function">byte</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token builtin">string</span><span class="token punctuation">)</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">case</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">:</span>
		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">.</span><span class="token punctuation">(</span><span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">byte</span><span class="token punctuation">)</span><span class="token punctuation">,</span> types<span class="token punctuation">.</span>String<span class="token punctuation">)</span>
		<span class="token punctuation">}</span>

	<span class="token keyword">default</span><span class="token punctuation">:</span>
		<span class="token keyword">for</span> i <span class="token operator">:=</span> <span class="token keyword">range</span> v <span class="token punctuation">{</span>

			jBytes<span class="token punctuation">,</span> err <span class="token operator">:=</span> <span class="token function">marshal</span><span class="token punctuation">(</span>v<span class="token punctuation">[</span>i<span class="token punctuation">]</span><span class="token punctuation">)</span>
			<span class="token keyword">if</span> err <span class="token operator">!=</span> <span class="token boolean">nil</span> <span class="token punctuation">{</span>
				<span class="token keyword">return</span> err
			<span class="token punctuation">}</span>
			<span class="token function">callback</span><span class="token punctuation">(</span>jBytes<span class="token punctuation">,</span> dataType<span class="token punctuation">)</span>

		<span class="token punctuation">}</span>
	<span class="token punctuation">}</span>

	<span class="token keyword">return</span> <span class="token boolean">nil</span>
<span class="token punctuation">}</span>

<span class="token comment">/*func readArrayWithTypeByMapIfaceIface(marshal func(interface{}) ([]byte, error), v map[interface{}]interface{}, callback func([]byte, string)) error {
	for key, val := range v {

		bKey := []byte(fmt.Sprint(key) + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayWithTypeByMapStrStr(v map[string]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(key + ": " + val))
	}

	return nil
}

func readArrayWithTypeByMapStrIface(marshal func(interface{}) ([]byte, error), v map[string]interface{}, callback func([]byte, string)) error {
	for key, val := range v {

		bKey := []byte(key + ": ")
		b, err := marshal(val)
		if err != nil {
			return err
		}
		callback(append(bKey, b...))
	}

	return nil
}

func readArrayWithTypeByMapIfaceStr(v map[interface{}]string, callback func([]byte, string)) error {
	for key, val := range v {

		callback([]byte(fmt.Sprint(key) + ": " + val))
	}

	return nil
}
*/</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="parameters" tabindex="-1"><a class="header-anchor" href="#parameters" aria-hidden="true">#</a> Parameters</h2>
<ol>
<li><code v-pre>func(interface{}) ([]byte, error)</code>: data type's marshaller</li>
<li><code v-pre>func([]byte, interface{}) error</code>: data type's unmarshaller</li>
<li><code v-pre>stdio.Io</code>: stream to read from (eg STDIN)</li>
<li><code v-pre>func([]byte, string)</code>: callback function to write each array element</li>
</ol>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/apis/ReadArray.html">apis/<code v-pre>ReadArray()</code> (type)</RouterLink>:
Read from a data type one array element at a time</li>
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


