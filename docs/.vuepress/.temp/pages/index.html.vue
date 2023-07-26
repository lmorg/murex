<template><div><h2 id="👁‍🗨-screenshots" tabindex="-1"><a class="header-anchor" href="#👁‍🗨-screenshots" aria-hidden="true">#</a> 👁‍🗨 Screenshots</h2>
<!-- markdownlint-disable -->
<div class="image-preview">
  <img src="/murex-kill-autocomplete.png" />
  <img src="/murex-open-foreach.png" />
  <img src="/murex-spellchecker.png" />
</div>
<!-- markdownlint-restore -->
<h2 id="📦-easy-to-install" tabindex="-1"><a class="header-anchor" href="#📦-easy-to-install" aria-hidden="true">#</a> 📦 Easy to Install</h2>
<p>Install <code v-pre>Murex</code> from your favorite package manager</p>
<CodeTabs id="13" :data='[{"id":"Brew"},{"id":"AUR Package"},{"id":"PKGBuild file"},{"id":"MacPorts"}]' tab-id="shell">
<template #title0="{ value, isActive }">Brew</template>
<template #title1="{ value, isActive }">AUR Package</template>
<template #title2="{ value, isActive }">PKGBuild file</template>
<template #title3="{ value, isActive }">MacPorts</template>
<template #tab0="{ value, isActive }">
<div class="language-bash line-numbers-mode" data-ext="sh"><pre v-pre class="language-bash"><code>brew <span class="token function">install</span> murex
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div></template>
<template #tab1="{ value, isActive }">
<div class="language-bash line-numbers-mode" data-ext="sh"><pre v-pre class="language-bash"><code>https://aur.archlinux.org/packages/murex
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div></template>
<template #tab2="{ value, isActive }">
<div class="language-bash line-numbers-mode" data-ext="sh"><pre v-pre class="language-bash"><code>https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h<span class="token operator">=</span>murex
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div></template>
<template #tab3="{ value, isActive }">
<div class="language-bash line-numbers-mode" data-ext="sh"><pre v-pre class="language-bash"><code><span class="token function">sudo</span> port <span class="token function">install</span> murex
</code></pre><div class="line-numbers" aria-hidden="true"><div class="line-number"></div></div></div></template>
</CodeTabs>
<p>More options are available in the <RouterLink to="/install/">INSTALL</RouterLink> document</p>
<h2 id="🛟-getting-started" tabindex="-1"><a class="header-anchor" href="#🛟-getting-started" aria-hidden="true">#</a> 🛟 Getting Started</h2>
<p>Take your your first steps with <code v-pre>Murex</code> by following our <RouterLink to="/tour/">Language Tutorial</RouterLink></p>
</div></template>



<style>
  .image-preview {
    display: flex;
    justify-content: space-evenly;
    align-items: center;
    flex-wrap: wrap;
  }

  .image-preview > img {
     box-sizing: border-box;
     width: 33.3% !important;
     padding: 9px;
     border-radius: 16px;
  }

  @media (max-width: 719px){
    .image-preview > img {
      width: 50% !important;
    }
  }

  @media (max-width: 419px){
    .image-preview > img {
      width: 100% !important;
    }
  }
</style>