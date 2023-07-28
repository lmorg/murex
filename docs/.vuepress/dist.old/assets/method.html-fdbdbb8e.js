import{_ as o}from"./plugin-vue_export-helper-c27b6911.js";import{r as i,o as l,c as r,d as n,b as s,w as t,e as a,f as p}from"./app-45f7c304.js";const d={},c=p(`<h1 id="method" tabindex="-1"><a class="header-anchor" href="#method" aria-hidden="true">#</a> <code>method</code></h1><blockquote><p>Define a methods supported data-types</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>method</code> defines what the typical data type would be for a function&#39;s STDIN and STDOUT.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>method: define name { json }
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>method: define name {
    &quot;Stdin&quot;:  &quot;@Any&quot;,
    &quot;Stdout&quot;: &quot;json&quot;
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="type-groups" tabindex="-1"><a class="header-anchor" href="#type-groups" aria-hidden="true">#</a> Type Groups</h3><p>You can define a Murex data type or use a type group. The following type groups are available to use:</p><div class="language-go line-numbers-mode" data-ext="go"><pre class="language-go"><code><span class="token keyword">package</span> types

<span class="token comment">// These are the different supported type groups</span>
<span class="token keyword">const</span> <span class="token punctuation">(</span>
	Any               <span class="token operator">=</span> <span class="token string">&quot;@Any&quot;</span>
	Text              <span class="token operator">=</span> <span class="token string">&quot;@Text&quot;</span>
	Math              <span class="token operator">=</span> <span class="token string">&quot;@Math&quot;</span>
	Unmarshal         <span class="token operator">=</span> <span class="token string">&quot;@Unmarshal&quot;</span>
	Marshal           <span class="token operator">=</span> <span class="token string">&quot;@Marshal&quot;</span>
	ReadArray         <span class="token operator">=</span> <span class="token string">&quot;@ReadArray&quot;</span>
	ReadArrayWithType <span class="token operator">=</span> <span class="token string">&quot;@ReadArrayWithType&quot;</span>
	WriteArray        <span class="token operator">=</span> <span class="token string">&quot;@WriteArray&quot;</span>
	ReadIndex         <span class="token operator">=</span> <span class="token string">&quot;@ReadIndex&quot;</span>
	ReadNotIndex      <span class="token operator">=</span> <span class="token string">&quot;@ReadNotIndex&quot;</span>
	ReadMap           <span class="token operator">=</span> <span class="token string">&quot;@ReadMap&quot;</span>
<span class="token punctuation">)</span>

<span class="token comment">// GroupText is an array of the data types that make up the \`text\` type</span>
<span class="token keyword">var</span> GroupText <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	Generic<span class="token punctuation">,</span>
	String<span class="token punctuation">,</span>
	<span class="token string">\`generic\`</span><span class="token punctuation">,</span>
	<span class="token string">\`string\`</span><span class="token punctuation">,</span>
<span class="token punctuation">}</span>

<span class="token comment">// GroupMath is an array of the data types that make up the \`math\` type</span>
<span class="token keyword">var</span> GroupMath <span class="token operator">=</span> <span class="token punctuation">[</span><span class="token punctuation">]</span><span class="token builtin">string</span><span class="token punctuation">{</span>
	Number<span class="token punctuation">,</span>
	Integer<span class="token punctuation">,</span>
	Float<span class="token punctuation">,</span>
	Boolean<span class="token punctuation">,</span>
<span class="token punctuation">}</span>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,13),u=n("code",null,"->",-1),m=n("code",null,"alias",-1),h=n("code",null,"autocomplete",-1),v=n("code",null,"function",-1),k=n("code",null,"private",-1),b=n("code",null,"runtime",-1);function f(g,_){const e=i("RouterLink");return l(),r("div",null,[c,n("ul",null,[n("li",null,[s(e,{to:"/parser/pipe-arrow.html"},{default:t(()=>[a("Arrow Pipe ("),u,a(") Token")]),_:1}),a(": Pipes STDOUT from the left hand command to STDIN of the right hand command")]),n("li",null,[s(e,{to:"/user-guide/interactive-shell.html"},{default:t(()=>[a("Murex's Interactive Shell")]),_:1}),a(": What's different about Murex's interactive shell?")]),n("li",null,[s(e,{to:"/commands/alias.html"},{default:t(()=>[m]),_:1}),a(": Create an alias for a command")]),n("li",null,[s(e,{to:"/commands/autocomplete.html"},{default:t(()=>[h]),_:1}),a(": Set definitions for tab-completion in the command line")]),n("li",null,[s(e,{to:"/commands/function.html"},{default:t(()=>[v]),_:1}),a(": Define a function block")]),n("li",null,[s(e,{to:"/commands/private.html"},{default:t(()=>[k]),_:1}),a(": Define a private function block")]),n("li",null,[s(e,{to:"/commands/runtime.html"},{default:t(()=>[b]),_:1}),a(": Returns runtime information on the internal state of Murex")])])])}const x=o(d,[["render",f],["__file","method.html.vue"]]);export{x as default};
