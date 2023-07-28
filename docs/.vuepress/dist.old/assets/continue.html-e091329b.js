import{_ as i}from"./plugin-vue_export-helper-c27b6911.js";import{r as a,o as d,c as r,d as e,b as n,w as t,e as c,f as l}from"./app-45f7c304.js";const u={},s=l(`<h1 id="continue" tabindex="-1"><a class="header-anchor" href="#continue" aria-hidden="true">#</a> <code>continue</code></h1><blockquote><p>Terminate process of a block within a caller function</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p><code>continue</code> will terminate execution of a block (eg <code>function</code>, <code>private</code>, <code>foreach</code>, <code>if</code>, etc) right up until the caller function. In iteration loops like <code>foreach</code> and <code>formap</code> this will result in behavior similar to the <code>continue</code> statement in other programming languages.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>continue block-name
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code>%[1..10] -&gt; foreach i {
    if { $i == 5 } then {
        out &quot;continue&quot;
        continue foreach
        out &quot;skip this code&quot;
    }
    out $i
}
</code></pre><p>Running the above code would output:</p><pre><code>» foo
1
2
3
4
continue
6
7
8
9
10
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>continue</code> cannot escape the bounds of its scope (typically the function it is running inside). For example, in the following code we are calling <code>continue bar</code> (which is a different function) inside of the function <code>foo</code>:</p><pre><code>function foo {
    %[1..10] -&gt; foreach i {
        out $i
        if { $i == 5 } then {
            out &quot;exit running function&quot;
            continue bar
            out &quot;ended&quot;
        }
    }
}

function bar {
    foo
}
</code></pre><p>Regardless of whether we run <code>foo</code> or <code>bar</code>, both of those functions will raise the following error:</p><pre><code>Error in \`continue\` (7,17): no block found named \`bar\` within the scope of \`foo\`
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,16),h=e("code",null,"break",-1),f=e("code",null,"exit",-1),m=e("code",null,"foreach",-1),p=e("code",null,"formap",-1),_=e("code",null,"function",-1),b=e("code",null,"if",-1),x=e("code",null,"out",-1),g=e("code",null,"private",-1),k=e("code",null,"return",-1);function w(q,v){const o=a("RouterLink");return d(),r("div",null,[s,e("ul",null,[e("li",null,[n(o,{to:"/commands/break.html"},{default:t(()=>[h]),_:1}),c(": Terminate execution of a block within your processes scope")]),e("li",null,[n(o,{to:"/commands/exit.html"},{default:t(()=>[f]),_:1}),c(": Exit murex")]),e("li",null,[n(o,{to:"/commands/foreach.html"},{default:t(()=>[m]),_:1}),c(": Iterate through an array")]),e("li",null,[n(o,{to:"/commands/formap.html"},{default:t(()=>[p]),_:1}),c(": Iterate through a map or other collection of data")]),e("li",null,[n(o,{to:"/commands/function.html"},{default:t(()=>[_]),_:1}),c(": Define a function block")]),e("li",null,[n(o,{to:"/commands/if.html"},{default:t(()=>[b]),_:1}),c(": Conditional statement to execute different blocks of code depending on the result of the condition")]),e("li",null,[n(o,{to:"/commands/out.html"},{default:t(()=>[x]),_:1}),c(": Print a string to the STDOUT with a trailing new line character")]),e("li",null,[n(o,{to:"/commands/private.html"},{default:t(()=>[g]),_:1}),c(": Define a private function block")]),e("li",null,[n(o,{to:"/commands/return.html"},{default:t(()=>[k]),_:1}),c(": Exits current function scope")])])])}const T=i(u,[["render",w],["__file","continue.html.vue"]]);export{T as default};
