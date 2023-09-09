import { navbar } from "vuepress-theme-hope";

export default navbar([
  "/",
  {
    text: "Documentation",
    icon: "book",
    children: [
      {
        text: "Shortcuts",
        prefix: "/",
        children: [
          { text: "Install", link: "install.html", icon: "arrow-down" },
          { text: "Getting Started", link: "tour.html", icon: "life-ring" },
          { text: "Rosetta Stone", link: "user-guide/rosetta-stone.html", icon: "language" },
          { text: "User Guide", link: "user-guide/", icon: "book" },
          { text: "Builtins", link: "commands/", icon: "terminal" },
        ],
      },
    ],
  },
  "/changelog/",
  "/blog/",
  { text: "Discuss", link: "https://github.com/lmorg/murex/discussions", icon: "comment" },
  "/contributing",
]);
