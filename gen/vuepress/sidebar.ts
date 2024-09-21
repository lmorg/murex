import { sidebar } from "vuepress-theme-hope";
import commandsMenu from "./commands_generated.json"
import parserMenu from "./parser_generated.json"
import variablesMenu from "./variables_generated.json"
import userguideMenu from "./userguide_generated.json"

export default sidebar({
  "/": [
    {
      text: "Murex",
      icon: "house",
      children: [
        "/install/",
        "compatibility/", 
        "/changelog/",
        { text: "Language Tour", link: "tour.html", icon: "plane-departure" }, 
        { text: "Rosetta Stone", link: "user-guide/rosetta-stone.html", icon: "table" },
        { text: "Operators And Tokens", link: "user-guide/operators-and-tokens.html", icon: "hashtag" },
        "/contributing",
        "/blog/",
      ],
      collapsible: false,
    },
    {
      text: "User Guide",
      icon: "book",
      prefix: "/",
      children: userguideMenu,
      collapsible: false,
    },
    {
      text: "Integrations",
      icon: "puzzle-piece",
      prefix: "integrations/",
      children: "structure",
      collapsible: false,
    },
    {
      text: "Operators And Tokens",
      icon: "hashtag",
      prefix: "/",
      children: parserMenu,
      collapsible: false,
    },
    {
      text: "Builtin Commands",
      icon: "cubes",
      prefix: "/",
      children: commandsMenu,
      collapsible: false,
    },
    /*{
      text: "Optional Builtins",
      icon: "cube",
      prefix: "optional/",
      children: "structure",
      collapsible: true,
    },*/
    {
      text: "Special Variables",
      icon: "dollar",
      prefix: "/",
      children: variablesMenu,
      collapsible: false,
    },
    {
      text: "Data Types",
      icon: "file-contract",
      prefix: "types/",
      children: "structure",
      collapsible: false,
    },
    {
      text: "Events",
      icon: "bolt",
      prefix: "events/",
      children: "structure",
      collapsible: false,
    },
    {
      text: "API Reference",
      icon: "gears",
      prefix: "apis/",
      children: "structure",
      collapsible: false,
    },
  ],
});
