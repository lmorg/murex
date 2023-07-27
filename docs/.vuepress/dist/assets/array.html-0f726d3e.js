import{_ as r}from"./plugin-vue_export-helper-c27b6911.js";import{r as l,o as s,c as i,d as o,b as a,w as n,e,f as d}from"./app-45f7c304.js";const u={},c=d(`<h1 id="array-token-parser-reference" tabindex="-1"><a class="header-anchor" href="#array-token-parser-reference" aria-hidden="true">#</a> Array (<code>@</code>) Token - Parser Reference</h1><blockquote><p>Expand values as an array</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>The array token is used to tell Murex to expand the string as multiple parameters (an array) rather than as a single parameter string.</p><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p><strong>ASCII variable names:</strong></p><pre><code>» $example = &quot;foobar&quot;
» out $example
foobar
</code></pre><p><strong>Unicode variable names:</strong></p><p>Variable names can be non-ASCII however they have to be surrounded by parenthesis. eg</p><pre><code>» $(比如) = &quot;举手之劳就可以使办公室更加环保，比如，使用再生纸。&quot;
» out $(比如)
举手之劳就可以使办公室更加环保，比如，使用再生纸。
</code></pre><p><strong>Infixing inside text:</strong></p><p>Sometimes you need to denote the end of a variable and have text follow on.</p><pre><code>» $partial_word = &quot;orl&quot;
» out &quot;Hello w$(partial_word)d!&quot;
Hello world!
</code></pre><p><strong>Variables are tokens:</strong></p><p>Please note the new line (<code>\\n</code>) character. This is not split using <code>$</code>:</p><pre><code>» $example = &quot;foo\\nbar&quot;
</code></pre><p>Output as a string:</p><pre><code>» out $example
foo
bar
</code></pre><p>Output as an array:</p><pre><code>» out @example
foo bar
</code></pre><p>The string and array tokens also works for subshells:</p><pre><code>» out \${ %[Mon..Fri] }
[&quot;Mon&quot;,&quot;Tue&quot;,&quot;Wed&quot;,&quot;Thu&quot;,&quot;Fri&quot;]

» out @{ %[Mon..Fri] }
Mon Tue Wed Thu Fri
</code></pre><blockquote><p><code>out</code> will take an array and output each element, space delimited. Exactly the same how <code>echo</code> would in Bash.</p></blockquote><p><strong>Variable as a command:</strong></p><p>If a variable is used as a commend then Murex will just print the content of that variable.</p><pre><code>» $example = &quot;Hello World!&quot;

» $example
Hello World!
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p>Since arrays are expanded over multiple parameters, you cannot expand an array inside quoted strings like you can with a string variable:</p><pre><code>» out: &quot;foo \${ ja: [1..5] } bar&quot;
foo [&quot;1&quot;,&quot;2&quot;,&quot;3&quot;,&quot;4&quot;,&quot;5&quot;] bar

» out: &quot;foo @{ ja: [1..5] } bar&quot;
foo  1 2 3 4 5  bar

» %(\${ ja: [1..5] })
[&quot;1&quot;,&quot;2&quot;,&quot;3&quot;,&quot;4&quot;,&quot;5&quot;]

» %(@{ ja: [1..5] })
@{ ja: [1..5] }
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,30),p=o("code",null,"%(",-1),h=o("code",null,")",-1),m=o("code",null,'"',-1),q=o("code",null,"'",-1),_=o("code",null,"$",-1),b=o("code",null,"~",-1),f=o("code",null,"(",-1),x=o("code",null,"ja",-1),g=o("code",null,"out",-1),y=o("code",null,"set",-1);function k(v,w){const t=l("RouterLink");return s(),i("div",null,[c,o("ul",null,[o("li",null,[a(t,{to:"/parser/brace-quote.html"},{default:n(()=>[e("Brace Quote ("),p,e(", "),h,e(") Tokens")]),_:1}),e(": Initiates or terminates a string (variables expanded)")]),o("li",null,[a(t,{to:"/parser/double-quote.html"},{default:n(()=>[e("Double Quote ("),m,e(") Token")]),_:1}),e(": Initiates or terminates a string (variables expanded)")]),o("li",null,[a(t,{to:"/parser/single-quote.html"},{default:n(()=>[e("Single Quote ("),q,e(") Token")]),_:1}),e(": Initiates or terminates a string (variables not expanded)")]),o("li",null,[a(t,{to:"/parser/string.html"},{default:n(()=>[e("String ("),_,e(") Token")]),_:1}),e(": Expand values as a string")]),o("li",null,[a(t,{to:"/parser/tilde.html"},{default:n(()=>[e("Tilde ("),b,e(") Token")]),_:1}),e(": Home directory path variable")]),o("li",null,[a(t,{to:"/commands/brace-quote.html"},{default:n(()=>[f,e(" (brace quote)")]),_:1}),e(": Write a string to the STDOUT without new line")]),o("li",null,[a(t,{to:"/commands/ja.html"},{default:n(()=>[x,e(" (mkarray)")]),_:1}),e(": A sophisticated yet simply way to build a JSON array")]),o("li",null,[a(t,{to:"/commands/out.html"},{default:n(()=>[g]),_:1}),e(": Print a string to the STDOUT with a trailing new line character")]),o("li",null,[a(t,{to:"/commands/set.html"},{default:n(()=>[y]),_:1}),e(": Define a local variable and set it's value")])])])}const S=r(u,[["render",k],["__file","array.html.vue"]]);export{S as default};
