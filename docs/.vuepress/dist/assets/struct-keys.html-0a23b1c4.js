import{_ as n}from"./plugin-vue_export-helper-c27b6911.js";import{r as a,o as r,c as q,d as t,b as u,w as s,e,f as d}from"./app-45f7c304.js";const l={},i=d(`<h1 id="struct-keys" tabindex="-1"><a class="header-anchor" href="#struct-keys" aria-hidden="true">#</a> <code>struct-keys</code></h1><blockquote><p>Outputs all the keys in a structure as a file path</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>struct-keys</code> outputs all of the keys in a structured data-type eg JSON, YAML, TOML, etc.</p><p>The output is a JSON array of the keys with each value being a file path representation of the input structure&#39;s node.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>\`&lt;stdin&gt;\` -&gt; struct-keys [ depth ] -&gt; \`&lt;stdout&gt;\`

\`&lt;stdin&gt;\` -&gt; struct-keys [ flags ] -&gt; \`&lt;stdout&gt;\`
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p>The source for these examples will be defined in the variable <code>$example</code>:</p><pre><code>» set: json example={
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
</code></pre><p>Without any flags set:</p><pre><code>» $example -&gt; struct-keys
[
    &quot;/lastName&quot;,
    &quot;/isAlive&quot;,
    &quot;/age&quot;,
    &quot;/address&quot;,
    &quot;/address/state&quot;,
    &quot;/address/postalCode&quot;,
    &quot;/address/streetAddress&quot;,
    &quot;/address/city&quot;,
    &quot;/phoneNumbers&quot;,
    &quot;/phoneNumbers/0&quot;,
    &quot;/phoneNumbers/0/type&quot;,
    &quot;/phoneNumbers/0/number&quot;,
    &quot;/phoneNumbers/1&quot;,
    &quot;/phoneNumbers/1/number&quot;,
    &quot;/phoneNumbers/1/type&quot;,
    &quot;/phoneNumbers/2&quot;,
    &quot;/phoneNumbers/2/type&quot;,
    &quot;/phoneNumbers/2/number&quot;,
    &quot;/children&quot;,
    &quot;/spouse&quot;,
    &quot;/firstName&quot;
]
</code></pre><p>Defining max depth and changing the separator string:</p><pre><code>» $example -&gt; struct-keys --depth 1 --separator &#39;.&#39;
[
    &quot;.children&quot;,
    &quot;.spouse&quot;,
    &quot;.firstName&quot;,
    &quot;.lastName&quot;,
    &quot;.isAlive&quot;,
    &quot;.age&quot;,
    &quot;.address&quot;,
    &quot;.phoneNumbers&quot;
]
</code></pre><p>An example of a unicode character being used as a separator:</p><pre><code>» $example -&gt; struct-keys --depth 2 --separator ☺
[
    &quot;☺age&quot;,
    &quot;☺address&quot;,
    &quot;☺address☺streetAddress&quot;,
    &quot;☺address☺city&quot;,
    &quot;☺address☺state&quot;,
    &quot;☺address☺postalCode&quot;,
    &quot;☺phoneNumbers&quot;,
    &quot;☺phoneNumbers☺0&quot;,
    &quot;☺phoneNumbers☺1&quot;,
    &quot;☺phoneNumbers☺2&quot;,
    &quot;☺children&quot;,
    &quot;☺spouse&quot;,
    &quot;☺firstName&quot;,
    &quot;☺lastName&quot;,
    &quot;☺isAlive&quot;
]
</code></pre><p>Separator can also be multiple characters:</p><pre><code>» $example -&gt; struct-keys --depth 1 --separator &#39;|||&#39;
[
    &quot;|||firstName&quot;,
    &quot;|||lastName&quot;,
    &quot;|||isAlive&quot;,
    &quot;|||age&quot;,
    &quot;|||address&quot;,
    &quot;|||phoneNumbers&quot;,
    &quot;|||children&quot;,
    &quot;|||spouse&quot;
]
</code></pre><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2><ul><li><code>--depth</code> How far to traverse inside the nested structure</li><li><code>--separator</code> String to use as a separator between fields (defaults to <code>/</code>)</li><li><code>-d</code> Alias for \`--depth</li><li><code>-s</code> Alias for \`--separator</li></ul><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,21),c=t("code",null,"[[",-1),p=t("code",null,"[",-1),h=t("code",null,"formap",-1),m=t("code",null,"set",-1);function f(b,N){const o=a("RouterLink");return r(),q("div",null,[i,t("ul",null,[t("li",null,[u(o,{to:"/commands/element.html"},{default:s(()=>[c,e(" (element)")]),_:1}),e(": Outputs an element from a nested structure")]),t("li",null,[u(o,{to:"/commands/index2.html"},{default:s(()=>[p,e(" (index)")]),_:1}),e(": Outputs an element from an array, map or table")]),t("li",null,[u(o,{to:"/commands/formap.html"},{default:s(()=>[h]),_:1}),e(": Iterate through a map or other collection of data")]),t("li",null,[u(o,{to:"/commands/set.html"},{default:s(()=>[m]),_:1}),e(": Define a local variable and set it's value")])])])}const _=n(l,[["render",f],["__file","struct-keys.html.vue"]]);export{_ as default};
