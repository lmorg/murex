import { defineUserConfig } from "vuepress";
import { searchProPlugin } from "vuepress-plugin-search-pro";
import theme from "./theme.js";

const environment = process.env.NODE_ENV;

const config = {
  // this must be replace with the context path in PROD
  base: "/murex-docs/",
  lang: "en-US",
  title: "",
  description: "Murex, a typed, content aware shell from the 2020s",
  head: [
    ["link", { rel: "preconnect", href: "https://fonts.googleapis.com" }],
    [
      "link",
      { rel: "preconnect", href: "https://fonts.gstatic.com", crossorigin: "" },
    ],
    [
      "link",
      {
        href: "https://fonts.googleapis.com/css2?family=Quicksand:wght@400;500;700&display=swap",
        rel: "stylesheet",
      },
    ],
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
        // {
        //   //@ts-ignore
        //   getter: (page) => page.frontmatter.tag,
        //   formatter: "Tag: $content",
        // },
      ],
    }),
  ],
};

//if (environment === "DEV") {
  config.base = "/";
//}
//@ts-ignore
export default defineUserConfig(config);
