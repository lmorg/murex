import { navbar } from "vuepress-theme-hope";

export default navbar([
  {
    text: "Documentation",
    icon: "book",
    children: [
      {
        //text: "Documentation",
        prefix: "/",
        children: [
          { text: "User Guide", link: "user-guide/", icon: "book" },
          { text: "Integrations", link: "integrations/", icon: "puzzle-piece" },
          { text: "Operators And Tokens", link: "parser/", icon: "hashtag" },
          { text: "Builtins Commands", link: "commands/", icon: "cubes" },
          { text: "Special Variables", link: "variables/", icon: "dollar" },
          { text: "Data Types", link: "types/", icon: "file-contract" },
          { text: "Events", link: "events/", icon: "bolt" },
          { text: "Blog", link: "blog/", icon: "comment" },
          "changelog/",
        ],
      },
    ],
  },
  { text: "Install", link: "install.html", icon: "download" },
  { text: "Tour", link: "tour.html", icon: "plane-departure" },
  { text: "Cheat Sheet", link: "user-guide/rosetta-stone.html", icon: "table" },
  { text: "Discuss", link: "https://github.com/lmorg/murex/discussions", icon: "comment" },
  { text: "Contribute", link: "/contributing", icon: "handshake-angle" },
  //"/contributing",
]);
