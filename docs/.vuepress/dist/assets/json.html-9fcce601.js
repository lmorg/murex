import{_ as s}from"./plugin-vue_export-helper-c27b6911.js";import{r as u,o as r,c as d,d as t,e,b as o,w as n,f as l}from"./app-45f7c304.js";const c={},h=l('<h1 id="json-data-type-reference" tabindex="-1"><a class="header-anchor" href="#json-data-type-reference" aria-hidden="true">#</a> <code>json</code> - Data-Type Reference</h1><blockquote><p>JavaScript Object Notation (JSON)</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>JSON is a structured data-type within Murex. It is the standard format for all structured data within Murex however other formats such as YAML, TOML and CSV are equally first class citizens.</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>',5),q={href:"https://en.wikipedia.org/wiki/JSON",target:"_blank",rel:"noopener noreferrer"},p=l(`<pre><code>{
  &quot;firstName&quot;: &quot;John&quot;,
  &quot;lastName&quot;: &quot;Smith&quot;,
  &quot;isAlive&quot;: true,
  &quot;age&quot;: 27,
  &quot;address&quot;: {
    &quot;streetAddress&quot;: &quot;21 2nd Street&quot;,
    &quot;city&quot;: &quot;New York&quot;,
    &quot;state&quot;: &quot;NY&quot;,
    &quot;postalCode&quot;: &quot;10021-3100&quot;
  },
  &quot;phoneNumbers&quot;: [
    {
      &quot;type&quot;: &quot;home&quot;,
      &quot;number&quot;: &quot;212 555-1234&quot;
    },
    {
      &quot;type&quot;: &quot;office&quot;,
      &quot;number&quot;: &quot;646 555-4567&quot;
    },
    {
      &quot;type&quot;: &quot;mobile&quot;,
      &quot;number&quot;: &quot;123 456-7890&quot;
    }
  ],
  &quot;children&quot;: [],
  &quot;spouse&quot;: null
}
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><h3 id="tips-when-writing-json-inside-for-loops" tabindex="-1"><a class="header-anchor" href="#tips-when-writing-json-inside-for-loops" aria-hidden="true">#</a> Tips when writing JSON inside for loops</h3><p>One of the drawbacks (or maybe advantages, depending on your perspective) of JSON is that parsers generally expect a complete file for processing in that the JSON specification requires closing tags for every opening tag. This means it&#39;s not always suitable for streaming. For example</p><pre><code>» ja [1..3] -&gt; foreach i { out ({ &quot;$i&quot;: $i }) }
{ &quot;1&quot;: 1 }
{ &quot;2&quot;: 2 }
{ &quot;3&quot;: 3 }
</code></pre><p><strong>What does this even mean and how can you build a JSON file up sequentially?</strong></p><p>One answer if to write the output in a streaming file format and convert back to JSON</p><pre><code>» ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) }
- &quot;1&quot;: 1
- &quot;2&quot;: 2
- &quot;3&quot;: 3

» ja [1..3] -&gt; foreach i { out (- &quot;$i&quot;: $i) } -&gt; cast yaml -&gt; format json
[
    {
        &quot;1&quot;: 1
    },
    {
        &quot;2&quot;: 2
    },
    {
        &quot;3&quot;: 3
    }
]
</code></pre><p><strong>What if I&#39;m returning an object rather than writing one?</strong></p><p>The problem with building JSON structures from existing structures is that you can quickly end up with invalid JSON due to the specifications strict use of commas.</p><p>For example in the code below, each item block is it&#39;s own object and there are no <code>[ ... ]</code> encapsulating them to denote it is an array of objects, nor are the objects terminated by a comma.</p><pre><code>» config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar }
{
    &quot;Data-Type&quot;: &quot;bool&quot;,
    &quot;Default&quot;: true,
    &quot;Description&quot;: &quot;Display the interactive shell&#39;s hint text helper. Please note, even when this is disabled, it will still appear when used for regexp searches and other readline-specific functions&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: true
}
{
    &quot;Data-Type&quot;: &quot;block&quot;,
    &quot;Default&quot;: &quot;{ progress $PID }&quot;,
    &quot;Description&quot;: &quot;Murex function to execute when an \`exec\` process is stopped&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: &quot;{ progress $PID }&quot;
}
{
    &quot;Data-Type&quot;: &quot;bool&quot;,
    &quot;Default&quot;: true,
    &quot;Description&quot;: &quot;ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, {SGR} variables, etc&quot;,
    &quot;Dynamic&quot;: false,
    &quot;Foo&quot;: &quot;Bar&quot;,
    &quot;Global&quot;: true,
    &quot;Value&quot;: true
}
...
</code></pre><p>Luckily JSON also has it&#39;s own streaming format: JSON lines (<code>jsonl</code>). We can <code>cast</code> this output as <code>jsonl</code> then <code>format</code> it back into valid JSON:</p><pre><code>» config -&gt; [ shell ] -&gt; formap k v { $v -&gt; alter /Foo Bar } -&gt; cast jsonl -&gt; format json
[
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Write shell history (interactive shell) to disk&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    },
    {
        &quot;Data-Type&quot;: &quot;int&quot;,
        &quot;Default&quot;: 4,
        &quot;Description&quot;: &quot;Maximum number of lines with auto-completion suggestions to display&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: &quot;6&quot;
    },
    {
        &quot;Data-Type&quot;: &quot;bool&quot;,
        &quot;Default&quot;: true,
        &quot;Description&quot;: &quot;Display some status information about the stop process when ctrl+z is pressed (conceptually similar to ctrl+t / SIGINFO on some BSDs)&quot;,
        &quot;Dynamic&quot;: false,
        &quot;Foo&quot;: &quot;Bar&quot;,
        &quot;Global&quot;: true,
        &quot;Value&quot;: true
    },
...
</code></pre><h4 id="foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" tabindex="-1"><a class="header-anchor" href="#foreach-will-automatically-cast-it-s-output-as-jsonl-if-it-s-stdin-type-is-json" aria-hidden="true">#</a> <code>foreach</code> will automatically cast it&#39;s output as <code>jsonl</code> <em>if</em> it&#39;s STDIN type is <code>json</code></h4><pre><code>» ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name }
Hello Tom
Hello Dick
Hello Sally

» ja [Tom,Dick,Sally] -&gt; foreach name { out Hello $name } -&gt; debug -&gt; [[ /Data-Type/Murex ]]
jsonl

» ja: [Tom,Dick,Sally] -&gt; foreach: name { out Hello $name } -&gt; format: json
[
    &quot;Hello Tom&quot;,
    &quot;Hello Dick&quot;,
    &quot;Hello Sally&quot;
]
</code></pre><h2 id="default-associations" tabindex="-1"><a class="header-anchor" href="#default-associations" aria-hidden="true">#</a> Default Associations</h2><ul><li><strong>Extension</strong>: <code>json</code></li><li><strong>MIME</strong>: <code>application/json</code></li><li><strong>MIME</strong>: <code>application/x-json</code></li><li><strong>MIME</strong>: <code>text/json</code></li><li><strong>MIME</strong>: <code>text/x-json</code></li></ul><h2 id="supported-hooks" tabindex="-1"><a class="header-anchor" href="#supported-hooks" aria-hidden="true">#</a> Supported Hooks</h2><ul><li><code>Marshal()</code> Writes minified JSON when no TTY detected and human readable JSON when stdout is a TTY</li><li><code>ReadArray()</code> Works with JSON arrays. Maps are converted into arrays</li><li><code>ReadArrayWithType()</code> Works with JSON arrays. Maps are converted into arrays. Elements data-type in Murex mirrors the JSON type of the element</li><li><code>ReadIndex()</code> Works against all properties in JSON</li><li><code>ReadMap()</code> Works with JSON maps</li><li><code>ReadNotIndex()</code> Works against all properties in JSON</li><li><code>Unmarshal()</code> Supported</li><li><code>WriteArray()</code> Works with JSON arrays</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,21),m=t("code",null,"Marshal()",-1),f=t("code",null,"ReadArray()",-1),y=t("code",null,"ReadArrayWithType()",-1),_=t("code",null,"ReadIndex()",-1),g=t("code",null,"[",-1),b=t("code",null,"ReadMap()",-1),x=t("code",null,"ReadNotIndex()",-1),S=t("code",null,"![",-1),w=t("code",null,"Unmarshal()",-1),N=t("code",null,"WriteArray()",-1),j=t("code",null,"[[",-1),D=t("code",null,"[",-1),k=t("code",null,"cast",-1),O=t("code",null,"format",-1),T=t("code",null,"hcl",-1),M=t("code",null,"jsonc",-1),v=t("code",null,"jsonl",-1),J=t("code",null,"lang.ArrayTemplate()",-1),A=t("code",null,"lang.ArrayWithTypeTemplate()",-1),R=t("code",null,"open",-1),I=t("code",null,"pretty",-1),W=t("code",null,"runtime",-1),L=t("code",null,"toml",-1),$=t("code",null,"yaml",-1);function B(H,E){const i=u("ExternalLinkIcon"),a=u("RouterLink");return r(),d("div",null,[h,t("p",null,[e("Example JSON document taken from "),t("a",q,[e("Wikipedia"),o(i)])]),p,t("ul",null,[t("li",null,[o(a,{to:"/apis/Marshal.html"},{default:n(()=>[m,e(" (type)")]),_:1}),e(": Converts structured memory into a structured file format (eg for stdio)")]),t("li",null,[o(a,{to:"/apis/ReadArray.html"},{default:n(()=>[f,e(" (type)")]),_:1}),e(": Read from a data type one array element at a time")]),t("li",null,[o(a,{to:"/apis/ReadArrayWithType.html"},{default:n(()=>[y,e(" (type)")]),_:1}),e(": Read from a data type one array element at a time and return the elements contents and data type")]),t("li",null,[o(a,{to:"/apis/ReadIndex.html"},{default:n(()=>[_,e(" (type)")]),_:1}),e(": Data type handler for the index, "),g,e(", builtin")]),t("li",null,[o(a,{to:"/apis/ReadMap.html"},{default:n(()=>[b,e(" (type)")]),_:1}),e(": Treat data type as a key/value structure and read its contents")]),t("li",null,[o(a,{to:"/apis/ReadNotIndex.html"},{default:n(()=>[x,e(" (type)")]),_:1}),e(": Data type handler for the bang-prefixed index, "),S,e(", builtin")]),t("li",null,[o(a,{to:"/apis/Unmarshal.html"},{default:n(()=>[w,e(" (type)")]),_:1}),e(": Converts a structured file format into structured memory")]),t("li",null,[o(a,{to:"/apis/WriteArray.html"},{default:n(()=>[N,e(" (type)")]),_:1}),e(": Write a data type, one array element at a time")]),t("li",null,[o(a,{to:"/commands/element.html"},{default:n(()=>[j,e(" (element)")]),_:1}),e(": Outputs an element from a nested structure")]),t("li",null,[o(a,{to:"/commands/index2.html"},{default:n(()=>[D,e(" (index)")]),_:1}),e(": Outputs an element from an array, map or table")]),t("li",null,[o(a,{to:"/commands/cast.html"},{default:n(()=>[k]),_:1}),e(": Alters the data type of the previous function without altering it's output")]),t("li",null,[o(a,{to:"/commands/format.html"},{default:n(()=>[O]),_:1}),e(": Reformat one data-type into another data-type")]),t("li",null,[o(a,{to:"/types/hcl.html"},{default:n(()=>[T]),_:1}),e(": HashiCorp Configuration Language (HCL)")]),t("li",null,[o(a,{to:"/types/jsonc.html"},{default:n(()=>[M]),_:1}),e(": Concatenated JSON")]),t("li",null,[o(a,{to:"/types/jsonl.html"},{default:n(()=>[v]),_:1}),e(": JSON Lines")]),t("li",null,[o(a,{to:"/apis/lang.ArrayTemplate.html"},{default:n(()=>[J,e(" (template API)")]),_:1}),e(": Unmarshals a data type into a Go struct and returns the results as an array")]),t("li",null,[o(a,{to:"/apis/lang.ArrayWithTypeTemplate.html"},{default:n(()=>[A,e(" (template API)")]),_:1}),e(": Unmarshals a data type into a Go struct and returns the results as an array with data type included")]),t("li",null,[o(a,{to:"/commands/open.html"},{default:n(()=>[R]),_:1}),e(": Open a file with a preferred handler")]),t("li",null,[o(a,{to:"/commands/pretty.html"},{default:n(()=>[I]),_:1}),e(": Prettifies JSON to make it human readable")]),t("li",null,[o(a,{to:"/commands/runtime.html"},{default:n(()=>[W]),_:1}),e(": Returns runtime information on the internal state of Murex")]),t("li",null,[o(a,{to:"/types/toml.html"},{default:n(()=>[L]),_:1}),e(": Tom's Obvious, Minimal Language (TOML)")]),t("li",null,[o(a,{to:"/types/yaml.html"},{default:n(()=>[$]),_:1}),e(": YAML Ain't Markup Language (YAML)")]),t("li",null,[o(a,{to:"/types/mxjson.html"},{default:n(()=>[e("mxjson")]),_:1}),e(": Murex-flavoured JSON (deprecated)")])])])}const C=s(c,[["render",B],["__file","json.html.vue"]]);export{C as default};
