import { navbar } from "vuepress-theme-hope";

export default navbar([
  "/",
  {
    text: "Documentations",
    icon: "book",
    children: [
      {
        text: "Shortcuts",
        prefix: "/",
        children: [
          { text: "Install", link: "install.html", icon: "terminal" },
          { text: "Getting Started", link: "tour.html", icon: "life-ring" },
          { text: "Rosetta Stone", link: "user-guide/rosetta-stone.html", icon: "language" },
          { text: "User Guide", link: "user-guide/", icon: "book" },
          { text: "Commands", link: "commands/", icon: "terminal" },
        ],
      },
    ],
  },
  "/changelog/",
  "/blog/",
  { text: "Discuss", link: "https://github.com/lmorg/murex/discussions", icon: "comment" },
  "/contributing",
]);
