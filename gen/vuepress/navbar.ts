import { navbar } from "vuepress-theme-hope";

export default navbar([
  {
    text: "Documentation",
    icon: "book",
    children: [
      {
        text: "Documentation",
        prefix: "/",
        children: [
          "user-guide/",
          "integrations/",
          "parser/",
          "commands/",
          "variables/",
          "types/",
          "events/",
          "blog/",
          "changelog/",
        ],
      },
    ],
  },
  { text: "Install", link: "install.html", icon: "download" },
  { text: "Language Tour", link: "tour.html", icon: "plane-departure" },
  { text: "Cheat Sheet", link: "user-guide/rosetta-stone.html", icon: "table" },
  { text: "Discuss", link: "https://github.com/lmorg/murex/discussions", icon: "comment" },
  "/contributing",
]);
