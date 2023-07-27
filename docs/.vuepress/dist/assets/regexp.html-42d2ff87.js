import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as d,o as r,c as i,d as e,b as o,w as n,e as t,f as l}from"./app-45f7c304.js";const u={},c=l(`<h1 id="regexp" tabindex="-1"><a class="header-anchor" href="#regexp" aria-hidden="true">#</a> <code>regexp</code></h1><blockquote><p>Regexp tools for arrays / lists of strings</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>regexp</code> provides a few tools for text matching and manipulation against an array or list of strings - thus <code>regexp</code> is Murex data-type aware.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; regexp expression -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><h3 id="find-elements" tabindex="-1"><a class="header-anchor" href="#find-elements" aria-hidden="true">#</a> Find elements</h3><pre><code>» ja: [monday..sunday] -&gt; regexp &#39;f/^([a-z]{3})day/&#39;
[
    &quot;mon&quot;,
    &quot;fri&quot;,
    &quot;sun&quot;
]
</code></pre><p>This returns only 3 days because only 3 days match the expression (where the days have to be 6 characters long) and then it only returns the first 3 characters because those are inside the parenthesis.</p><h3 id="match-elements" tabindex="-1"><a class="header-anchor" href="#match-elements" aria-hidden="true">#</a> Match elements</h3><p>Elements containing</p><pre><code>» ja: [monday..sunday] -&gt; regexp &#39;m/(mon|fri|sun)day/&#39;
[
    &quot;monday&quot;,
    &quot;friday&quot;,
    &quot;sunday&quot;
]
</code></pre><p>Elements excluding</p><pre><code>» ja: [monday..sunday] -&gt; !regexp &#39;m/(mon|fri|sun)day/&#39;
[
    &quot;tuesday&quot;,
    &quot;wednesday&quot;,
    &quot;thursday&quot;,
    &quot;saturday&quot;
]
</code></pre><h3 id="substitute-expression" tabindex="-1"><a class="header-anchor" href="#substitute-expression" aria-hidden="true">#</a> Substitute expression</h3><pre><code>» ja: [monday..sunday] -&gt; regexp &#39;s/day/night/&#39;
[
    &quot;monnight&quot;,
    &quot;tuesnight&quot;,
    &quot;wednesnight&quot;,
    &quot;thursnight&quot;,
    &quot;frinight&quot;,
    &quot;saturnight&quot;,
    &quot;sunnight&quot;
]
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>f</code> output found expressions (doesn&#39;t support bang prefix)</li><li><code>m</code> output elements that match expression (supports bang prefix)</li><li><code>s</code> output all elements - substituting elements that match expression (doesn&#39;t support bang prefix)</li></ul><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>regexp</code> is data-type aware so will work against lists or arrays of whichever Murex data-type is passed to it via STDIN and return the output in the same data-type.</p><h2 id="synonyms" tabindex="-1"><a class="header-anchor" href="#synonyms" aria-hidden="true">#</a> Synonyms</h2><ul><li><code>regexp</code></li><li><code>!regexp</code></li><li><code>list.regex</code></li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,24),h=e("code",null,"2darray",-1),m=e("code",null,"a",-1),p=e("code",null,"append",-1),f=e("code",null,"count",-1),y=e("code",null,"ja",-1),x=e("code",null,"jsplit",-1),_=e("code",null,"map",-1),g=e("code",null,"match",-1),q=e("code",null,"msort",-1),b=e("code",null,"prefix",-1),w=e("code",null,"prepend",-1),S=e("code",null,"pretty",-1),k=e("code",null,"suffix",-1),v=e("code",null,"ta",-1);function N(j,A){const a=d("RouterLink");return r(),i("div",null,[c,e("ul",null,[e("li",null,[o(a,{to:"/commands/2darray.html"},{default:n(()=>[h]),_:1}),t(": Create a 2D JSON array from multiple input sources")]),e("li",null,[o(a,{to:"/commands/a.html"},{default:n(()=>[m,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simple way to build an array or list")]),e("li",null,[o(a,{to:"/commands/append.html"},{default:n(()=>[p]),_:1}),t(": Add data to the end of an array")]),e("li",null,[o(a,{to:"/commands/count.html"},{default:n(()=>[f]),_:1}),t(": Count items in a map, list or array")]),e("li",null,[o(a,{to:"/commands/ja.html"},{default:n(()=>[y,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simply way to build a JSON array")]),e("li",null,[o(a,{to:"/commands/jsplit.html"},{default:n(()=>[x]),_:1}),t(": Splits STDIN into a JSON array based on a regex parameter")]),e("li",null,[o(a,{to:"/commands/map.html"},{default:n(()=>[_]),_:1}),t(": Creates a map from two data sources")]),e("li",null,[o(a,{to:"/commands/match.html"},{default:n(()=>[g]),_:1}),t(": Match an exact value in an array")]),e("li",null,[o(a,{to:"/commands/msort.html"},{default:n(()=>[q]),_:1}),t(": Sorts an array - data type agnostic")]),e("li",null,[o(a,{to:"/commands/prefix.html"},{default:n(()=>[b]),_:1}),t(": Prefix a string to every item in a list")]),e("li",null,[o(a,{to:"/commands/prepend.html"},{default:n(()=>[w]),_:1}),t(": Add data to the start of an array")]),e("li",null,[o(a,{to:"/commands/pretty.html"},{default:n(()=>[S]),_:1}),t(": Prettifies JSON to make it human readable")]),e("li",null,[o(a,{to:"/commands/suffix.html"},{default:n(()=>[k]),_:1}),t(": Prefix a string to every item in a list")]),e("li",null,[o(a,{to:"/commands/ta.html"},{default:n(()=>[v,t(" (mkarray)")]),_:1}),t(": A sophisticated yet simple way to build an array of a user defined data-type")])])])}const E=s(u,[["render",N],["__file","regexp.html.vue"]]);export{E as default};
