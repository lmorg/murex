import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as d,c,d as e,e as t,b as o,w as n,f as s}from"./app-45f7c304.js";const u={},h=s('<h1 id="toml-data-type-reference" tabindex="-1"><a class="header-anchor" href="#toml-data-type-reference" aria-hidden="true">#</a> <code>toml</code> - Data-Type Reference</h1><blockquote><p>Tom&#39;s Obvious, Minimal Language (TOML)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>TOML support within Murex is pretty mature however it is not considered a primitive. Which means, while it is a recommended builtin which you should expect in most deployments of Murex, it&#39;s still an optional package and thus may not be present in some edge cases. This is because it relies on external source packages for the shell to compile.</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>',5),p={href:"https://en.wikipedia.org/wiki/TOML",target:"_blank",rel:"noopener noreferrer"},m=s(`<pre><code># This is a TOML document.

title = &quot;TOML Example&quot;

[owner]
name = &quot;Tom Preston-Werner&quot;
dob = 1979-05-27T07:32:00-08:00 # First class dates

[database]
server = &quot;192.168.1.1&quot;
ports = [ 8001, 8001, 8002 ]
connection_max = 5000
enabled = true

[servers]

  # Indentation (tabs and/or spaces) is allowed but not required
  [servers.alpha]
  ip = &quot;10.0.0.1&quot;
  dc = &quot;eqdc10&quot;

  [servers.beta]
  ip = &quot;10.0.0.2&quot;
  dc = &quot;eqdc10&quot;

[clients]
data = [ [&quot;gamma&quot;, &quot;delta&quot;], [1, 2] ]

# Line breaks are OK when inside arrays
hosts = [
  &quot;alpha&quot;,
  &quot;omega&quot;
]
</code></pre><h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2><ul><li><strong>Extension</strong>: <code>toml</code></li><li><strong>MIME</strong>: <code>application/toml</code></li><li><strong>MIME</strong>: <code>application/x-toml</code></li><li><strong>MIME</strong>: <code>text/toml</code></li><li><strong>MIME</strong>: <code>text/x-toml</code></li></ul><h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2><ul><li><code>Marshal()</code> Supported</li><li><code>ReadArray()</code> Hook supported albeit TOML doesn&#39;t support naked arrays</li><li><code>ReadArrayWithType()</code> Hook supported albeit TOML doesn&#39;t support naked arrays</li><li><code>ReadIndex()</code> Works against all properties in TOML</li><li><code>ReadNotIndex()</code> Works against all properties in TOML</li><li><code>Unmarshal()</code> Supported</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,6),_=e("code",null,"Marshal()",-1),f=e("code",null,"ReadArray()",-1),y=e("code",null,"ReadIndex()",-1),x=e("code",null,"[",-1),M=e("code",null,"ReadMap()",-1),b=e("code",null,"ReadNotIndex()",-1),k=e("code",null,"![",-1),q=e("code",null,"Unmarshal()",-1),g=e("code",null,"WriteArray()",-1),L=e("code",null,"[[",-1),T=e("code",null,"[",-1),O=e("code",null,"cast",-1),R=e("code",null,"format",-1),v=e("code",null,"json",-1),w=e("code",null,"jsonl",-1),I=e("code",null,"open",-1),A=e("code",null,"runtime",-1),E=e("code",null,"yaml",-1);function N(W,S){const i=l("ExternalLinkIcon"),a=l("RouterLink");return d(),c("div",null,[h,e("p",null,[t("Example TOML document taken from "),e("a",p,[t("Wikipedia"),o(i)])]),m,e("ul",null,[e("li",null,[o(a,{to:"/apis/Marshal.html"},{default:n(()=>[_,t(" (type)")]),_:1}),t(": Converts structured memory into a structured file format (eg for stdio)")]),e("li",null,[o(a,{to:"/apis/ReadArray.html"},{default:n(()=>[f,t(" (type)")]),_:1}),t(": Read from a data type one array element at a time")]),e("li",null,[o(a,{to:"/apis/ReadIndex.html"},{default:n(()=>[y,t(" (type)")]),_:1}),t(": Data type handler for the index, "),x,t(", builtin")]),e("li",null,[o(a,{to:"/apis/ReadMap.html"},{default:n(()=>[M,t(" (type)")]),_:1}),t(": Treat data type as a key/value structure and read its contents")]),e("li",null,[o(a,{to:"/apis/ReadNotIndex.html"},{default:n(()=>[b,t(" (type)")]),_:1}),t(": Data type handler for the bang-prefixed index, "),k,t(", builtin")]),e("li",null,[o(a,{to:"/apis/Unmarshal.html"},{default:n(()=>[q,t(" (type)")]),_:1}),t(": Converts a structured file format into structured memory")]),e("li",null,[o(a,{to:"/apis/WriteArray.html"},{default:n(()=>[g,t(" (type)")]),_:1}),t(": Write a data type, one array element at a time")]),e("li",null,[o(a,{to:"/commands/element.html"},{default:n(()=>[L,t(" (element)")]),_:1}),t(": Outputs an element from a nested structure")]),e("li",null,[o(a,{to:"/commands/index2.html"},{default:n(()=>[T,t(" (index)")]),_:1}),t(": Outputs an element from an array, map or table")]),e("li",null,[o(a,{to:"/commands/cast.html"},{default:n(()=>[O]),_:1}),t(": Alters the data type of the previous function without altering it's output")]),e("li",null,[o(a,{to:"/commands/format.html"},{default:n(()=>[R]),_:1}),t(": Reformat one data-type into another data-type")]),e("li",null,[o(a,{to:"/types/json.html"},{default:n(()=>[v]),_:1}),t(": JavaScript Object Notation (JSON)")]),e("li",null,[o(a,{to:"/types/jsonl.html"},{default:n(()=>[w]),_:1}),t(": JSON Lines")]),e("li",null,[o(a,{to:"/commands/open.html"},{default:n(()=>[I]),_:1}),t(": Open a file with a preferred handler")]),e("li",null,[o(a,{to:"/commands/runtime.html"},{default:n(()=>[A]),_:1}),t(": Returns runtime information on the internal state of Murex")]),e("li",null,[o(a,{to:"/types/yaml.html"},{default:n(()=>[E]),_:1}),t(": YAML Ain't Markup Language (YAML)")])])])}const C=r(u,[["render",N],["__file","toml.html.vue"]]);export{C as default};
