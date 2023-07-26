<template><div><h1 id="struct-keys" tabindex="-1"><a class="header-anchor" href="#struct-keys" aria-hidden="true">#</a> <code v-pre>struct-keys</code></h1>
<blockquote>
<p>Outputs all the keys in a structure as a file path</p>
</blockquote>
<h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2>
<p><code v-pre>struct-keys</code> outputs all of the keys in a structured data-type eg JSON, YAML,
TOML, etc.</p>
<p>The output is a JSON array of the keys with each value being a file path
representation of the input structure's node.</p>
<h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>&lt;stdin> -> struct-keys [ depth ] -> &lt;stdout>

&lt;stdin> -> struct-keys [ flags ] -> &lt;stdout>
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2>
<p>The source for these examples will be defined in the variable <code v-pre>$example</code>:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» set: json example={
      "firstName": "John",
      "lastName": "Smith",
      "isAlive": true,
      "age": 27,
      "address": {
          "streetAddress": "21 2nd Street",
          "city": "New York",
          "state": "NY",
          "postalCode": "10021-3100"
      },
      "phoneNumbers": [
          {
              "type": "home",
              "number": "212 555-1234"
          },
          {
              "type": "office",
              "number": "646 555-4567"
          },
          {
              "type": "mobile",
              "number": "123 456-7890"
          }
      ],
      "children": [],
      "spouse": null
  }
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Without any flags set:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $example -> struct-keys
[
    "/lastName",
    "/isAlive",
    "/age",
    "/address",
    "/address/state",
    "/address/postalCode",
    "/address/streetAddress",
    "/address/city",
    "/phoneNumbers",
    "/phoneNumbers/0",
    "/phoneNumbers/0/type",
    "/phoneNumbers/0/number",
    "/phoneNumbers/1",
    "/phoneNumbers/1/number",
    "/phoneNumbers/1/type",
    "/phoneNumbers/2",
    "/phoneNumbers/2/type",
    "/phoneNumbers/2/number",
    "/children",
    "/spouse",
    "/firstName"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Defining max depth and changing the separator string:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $example -> struct-keys --depth 1 --separator '.'
[
    ".children",
    ".spouse",
    ".firstName",
    ".lastName",
    ".isAlive",
    ".age",
    ".address",
    ".phoneNumbers"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>An example of a unicode character being used as a separator:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $example -> struct-keys --depth 2 --separator ☺
[
    "☺age",
    "☺address",
    "☺address☺streetAddress",
    "☺address☺city",
    "☺address☺state",
    "☺address☺postalCode",
    "☺phoneNumbers",
    "☺phoneNumbers☺0",
    "☺phoneNumbers☺1",
    "☺phoneNumbers☺2",
    "☺children",
    "☺spouse",
    "☺firstName",
    "☺lastName",
    "☺isAlive"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><p>Separator can also be multiple characters:</p>
<div class="language-text line-numbers-mode" data-ext="text"><pre v-pre class="language-text"><code>» $example -> struct-keys --depth 1 --separator '|||'
[
    "|||firstName",
    "|||lastName",
    "|||isAlive",
    "|||age",
    "|||address",
    "|||phoneNumbers",
    "|||children",
    "|||spouse"
]
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div><div class="line-number"></div></div></div><h2 id="flags" tabindex="-1"><a class="header-anchor" href="#flags" aria-hidden="true">#</a> Flags</h2>
<ul>
<li><code v-pre>--depth</code>
How far to traverse inside the nested structure</li>
<li><code v-pre>--separator</code>
String to use as a separator between fields (defaults to <code v-pre>/</code>)</li>
<li><code v-pre>-d</code>
Alias for `--depth</li>
<li><code v-pre>-s</code>
Alias for `--separator</li>
</ul>
<h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>
<ul>
<li><RouterLink to="/commands/element.html"><code v-pre>[[</code> (element)</RouterLink>:
Outputs an element from a nested structure</li>
<li><RouterLink to="/commands/index2.html"><code v-pre>[</code> (index)</RouterLink>:
Outputs an element from an array, map or table</li>
<li><RouterLink to="/commands/formap.html"><code v-pre>formap</code></RouterLink>:
Iterate through a map or other collection of data</li>
<li><RouterLink to="/commands/set.html"><code v-pre>set</code></RouterLink>:
Define a local variable and set it's value</li>
</ul>
</div></template>


