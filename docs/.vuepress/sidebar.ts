import { sidebar } from "vuepress-theme-hope";

export default sidebar({
  "/": [
    {
      text: "Murex",
      icon: "house",
      children: ["/install", "/changelog", "/blog", "/tour", "/rosetta"],
      collapsible: true,
    },
    {
      text: "User Guide",
      icon: "book",
      prefix: "user-guide/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "Operators And Tokens",
      icon: "equals",
      prefix: "parser/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "Builtins",
      icon: "terminal",
      prefix: "commands/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "Optional Builtins",
      icon: "terminal",
      prefix: "optional/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "Data Types",
      icon: "table",
      prefix: "types/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "Events",
      icon: "bolt",
      prefix: "events/",
      children: "structure",
      collapsible: true,
    },
    {
      text: "API Reference",
      icon: "gears",
      prefix: "apis/",
      children: "structure",
      collapsible: true,
    },
  ],
});
