import { defineUserConfig } from "vuepress";
import { searchProPlugin } from "vuepress-plugin-search-pro";
import theme from "./theme.js";

const environment = process.env.NODE_ENV;
const commitHash = process.env.COMMITHASHSHORT;

var domain = "murex.rocks"
if (environment === "DEV") {
  domain = "dev.murex.rocks"
}

const config = {
  // this must be replace with the context path in PROD
  base: "/",
  lang: "en-US",
  title: "Murex",
  description: "Murex, a typed, content aware shell from the 2020s",
  head: [
    ["link", { rel: "preconnect", href: "https://fonts.googleapis.com" }],
    [
      "link",
      { rel: "preconnect", href: "https://fonts.gstatic.com", crossorigin: "" },
    ],
    /*[
      "link",
      {
        href: "https://fonts.googleapis.com/css2?family=Quicksand:wght@400;500;700&display=swap",
        rel: "stylesheet",
      },
    ],*/
    [
      "link",
      {
        href: "https://fonts.googleapis.com/css?family=Lato:wght@300|Jura:wght@700|Source+Code+Pro&display=swap",
        rel: "stylesheet",
      },
    ],
    [
      "link",
      {
        href: "https://fonts.googleapis.com/css2?family=Libre+Baskerville:ital,wght@0,400;0,700;1,400&display=swap",
        rel: "stylesheet",
      },
    ],
    [
      "link",
      {
        href: "https://fonts.googleapis.com/css2?family=JetBrains+Mono:ital,wght@0,100..800;1,100..800&display=swap",
        rel: "stylesheet",
      },
    ],
    [
      "link",
      {
        href: "/favicon.ico?v="+commitHash,
        rel: "icon",
        type: "image/png",
      },
    ],
    [
      "link",
      {
        href: "/favicon-16x16.png?v="+commitHash,
        rel: "icon",
        type: "image/png",
      },
    ],
    [
      "link",
      {
        href: "/favicon-32x32.png?v="+commitHash,
        rel: "icon",
        type: "image/png",
      },
    ],


    [ "meta", { property:"og:type",  content: "article" } ],
    [ "meta", { property:"og:image", content: "https://"+domain+"/og-murex-v3.png" } ],

  ],
  theme,
  plugins: [
    searchProPlugin({
      // index all contents
      indexContent: true,
      // add supports for category and tags
      customFields: [
        {
          //@ts-ignore
          getter: (page) => page.frontmatter.category,
          formatter: "Category: $content",
        },
        {
          //@ts-ignore
          getter: (page) => page.frontmatter.tag,
          formatter: "Tag: $content",
        },
      ],
    }),
  ],
};

if (environment === "DEV") {
  config.base = "/";
}

//@ts-ignore
export default defineUserConfig(config);
