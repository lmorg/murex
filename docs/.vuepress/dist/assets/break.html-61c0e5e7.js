import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as c,o as r,c as d,d as e,b as t,w as n,e as a,f as l}from"./app-45f7c304.js";const u={},s=l(`<h1 id="break" tabindex="-1"><a class="header-anchor" href="#break" aria-hidden="true">#</a> <code>break</code></h1><blockquote><p>Terminate execution of a block within your processes scope</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>break</code> will terminate execution of a block (eg <code>function</code>, <code>private</code>, <code>foreach</code>, <code>if</code>, etc).</p><p><code>break</code> requires a parameter and that parameter is the name of the caller block you wish to break out of. If it is a <code>function</code> or <code>private</code>, then it will be the name of that function or private. If it is an <code>if</code> or <code>foreach</code> loop, then it will be <code>if</code> or <code>foreach</code> (respectively).</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>break block-name
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><p><strong>Exiting an iteration block:</strong></p><pre><code>function foo {
    %[1..10] -&gt; foreach i {
        out $i
        if { $i == 5 } then {
            out &quot;exit running function&quot;
            break foo
            out &quot;ended&quot;
        }
    }
}
</code></pre><p>Running the above code would output:</p><pre><code>» foo
1
2
3
4
5
exit running function
</code></pre><p><strong>Exiting a function:</strong></p><p><code>break</code> can be considered to exhibit the behavior of <em>return</em> (from other languages) too</p><pre><code>function example {
    if { $USER == &quot;root&quot; } then {
        err &quot;Don&#39;t run this as root&quot;
        break example
    }

    # ... do something ...
}
</code></pre><p>Though in this particular use case it is recommended that you use <code>return</code> instead, the above code does illustrate how <code>break</code> behaves.</p><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>break</code> cannot escape the bounds of its scope (typically the function it is running inside). For example, in the following code we are calling <code>break bar</code> (which is a different function) inside of the function <code>foo</code>:</p><pre><code>function foo {
    %[1..10] -&gt; foreach i {
        out $i
        if { $i == 5 } then {
            out &quot;exit running function&quot;
            break bar
            out &quot;ended&quot;
        }
    }
}

function bar {
    foo
}
</code></pre><p>Regardless of whether we run <code>foo</code> or <code>bar</code>, both of those functions will raise the following error:</p><pre><code>Error in \`break\` (7,17): no block found named \`bar\` within the scope of \`foo\`
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,22),h=e("code",null,"continue",-1),f=e("code",null,"exit",-1),p=e("code",null,"foreach",-1),m=e("code",null,"formap",-1),b=e("code",null,"function",-1),_=e("code",null,"if",-1),k=e("code",null,"out",-1),x=e("code",null,"private",-1),g=e("code",null,"return",-1);function w(q,v){const o=c("RouterLink");return r(),d("div",null,[s,e("ul",null,[e("li",null,[t(o,{to:"/commands/continue.html"},{default:n(()=>[h]),_:1}),a(": Terminate process of a block within a caller function")]),e("li",null,[t(o,{to:"/commands/exit.html"},{default:n(()=>[f]),_:1}),a(": Exit murex")]),e("li",null,[t(o,{to:"/commands/foreach.html"},{default:n(()=>[p]),_:1}),a(": Iterate through an array")]),e("li",null,[t(o,{to:"/commands/formap.html"},{default:n(()=>[m]),_:1}),a(": Iterate through a map or other collection of data")]),e("li",null,[t(o,{to:"/commands/function.html"},{default:n(()=>[b]),_:1}),a(": Define a function block")]),e("li",null,[t(o,{to:"/commands/if.html"},{default:n(()=>[_]),_:1}),a(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[t(o,{to:"/commands/out.html"},{default:n(()=>[k]),_:1}),a(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[t(o,{to:"/commands/private.html"},{default:n(()=>[x]),_:1}),a(": Define a private function block")]),e("li",null,[t(o,{to:"/commands/return.html"},{default:n(()=>[g]),_:1}),a(": Exits current function scope")])])])}const D=i(u,[["render",w],["__file","break.html.vue"]]);export{D as default};
