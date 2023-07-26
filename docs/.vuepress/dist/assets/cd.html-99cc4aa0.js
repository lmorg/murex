import{_ as d}from"./plugin-vue_export-helper-c27b6911.js";import{r as c,o as n,c as i,d as e,b as r,w as t,e as a,f as s}from"./app-45f7c304.js";const h={},l=s(`<h1 id="cd" tabindex="-1"><a class="header-anchor" href="#cd" aria-hidden="true">#</a> <code>cd</code></h1><blockquote><p>Change (working) directory</p></blockquote><h2 id="description" tabindex="-1"><a class="header-anchor" href="#description" aria-hidden="true">#</a> Description</h2><p>Changes current working directory.</p><h2 id="usage" tabindex="-1"><a class="header-anchor" href="#usage" aria-hidden="true">#</a> Usage</h2><pre><code>cd path
</code></pre><h2 id="examples" tabindex="-1"><a class="header-anchor" href="#examples" aria-hidden="true">#</a> Examples</h2><pre><code># Home directory
» cd: ~

# Absolute path
» cd: /etc/

# Relative path
» cd: Documents
» cd: ./Documents
</code></pre><h2 id="detail" tabindex="-1"><a class="header-anchor" href="#detail" aria-hidden="true">#</a> Detail</h2><p><code>cd</code> updates an environmental variable, <code>$PWDHIST</code> with an array of paths. You can then use that to change to a previous directory</p><pre><code># View the working directory history
» $PWDHIST

# Change to a previous directory
» cd $PWDHIST[0]
</code></pre><h3 id="auto-cd" tabindex="-1"><a class="header-anchor" href="#auto-cd" aria-hidden="true">#</a> auto-cd</h3><p>Some people prefer to omit <code>cd</code> and just write the path, with their shell automatically changing to that directory if the &quot;command&quot; is just a directory. In Murex you can enable this behaviour by turning on &quot;auto-cd&quot;:</p><pre><code>config: set shell auto-cd true
</code></pre><h2 id="see-also" tabindex="-1"><a class="header-anchor" href="#see-also" aria-hidden="true">#</a> See Also</h2>`,15),u=e("code",null,"source",-1);function p(m,f){const o=c("RouterLink");return n(),i("div",null,[l,e("ul",null,[e("li",null,[r(o,{to:"/user-guide/reserved-vars.html"},{default:t(()=>[a("Reserved Variables")]),_:1}),a(": Special variables reserved by Murex")]),e("li",null,[r(o,{to:"/commands/source.html"},{default:t(()=>[u]),_:1}),a(": Import Murex code from another file of code block")])])])}const x=d(h,[["render",p],["__file","cd.html.vue"]]);export{x as default};
