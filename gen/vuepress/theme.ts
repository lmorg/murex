import { hopeTheme } from "vuepress-theme-hope";
import navbar from "./navbar.js";
import sidebar from "./sidebar.js";

export default hopeTheme({
  hostname: "https://murex.rocks",

  author: {
    name: "Laurence Morgan",
    url: "https://github.com/lmorg",
  },

  favicon: "favicon.ico",

  fullscreen: false,

  iconAssets: "fontawesome-with-brands",
  //iconAssets: "public/fontawesome/css/brands.css",

  //logo: "/murex.svg",

  repo: "lmorg/murex",

  docsDir: "docs",

  editLink: false,

  // navbar
  navbar,

  // sidebar
  sidebar,

  // footer: "Empty footer",

  displayFooter: false,

  // encrypt: {
  //   config: {
  //     "/demo/encrypt.html": ["1234"],
  //   },
  // },

  // metaLocales: {
  //   editLink: "Edit this page on GitHub",
  // },

  plugins: {
    // TODO: You should generate and use your own comment service
    // comment: {
    //   provider: "Giscus",
    //   repo: "vuepress-theme-hope/giscus-discussions",
    //   repoId: "R_kgDOG_Pt2A",
    //   category: "Announcements",
    //   categoryId: "DIC_kwDOG_Pt2M4COD69",
    // },

    // All features are enabled for demo, only preserve features you need here
    mdEnhance: {
      align: true,
      attrs: true,
      // chart: true,
      codetabs: true,
      // demo: true,
      // echarts: true,
      // figure: true,
      // flowchart: true,
      //GitHub Flavored Markdown Spec
      //gfm: true,
      imgLazyload: true,
      imgSize: true,
      // include: true,
      // katex: true,
      // mark: true,
      // mermaid: true,
      // playground: {
      //   presets: ["ts", "vue"],
      // },
      // presentation: ["highlight", "math", "search", "notes", "zoom"],
      // stylize: [
      //   {
      //     matcher: "Recommended",
      //     replacer: ({ tag }) => {
      //       if (tag === "em")
      //         return {
      //           tag: "Badge",
      //           attrs: { type: "tip" },
      //           content: "Recommended",
      //         };
      //     },
      //   },
      // ],
      // sub: true,
      // sup: true,
      //tabs: true,
      // vPre: true,
      // vuePlayground: true,
    },
  },
});
