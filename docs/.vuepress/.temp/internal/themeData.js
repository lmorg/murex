export const themeData = JSON.parse("{\"encrypt\":{},\"author\":{\"name\":\"Laurence Morgan\",\"url\":\"https://github.com/lmorg\"},\"fullscreen\":false,\"logo\":\"/murex.svg\",\"repo\":\"lmorg/murex\",\"docsDir\":\"docs\",\"editLink\":false,\"displayFooter\":false,\"locales\":{\"/\":{\"lang\":\"en-US\",\"navbarLocales\":{\"langName\":\"English\",\"selectLangAriaLabel\":\"Select language\"},\"metaLocales\":{\"author\":\"Author\",\"date\":\"Writing Date\",\"origin\":\"Original\",\"views\":\"Page views\",\"category\":\"Category\",\"tag\":\"Tag\",\"readingTime\":\"Reading Time\",\"words\":\"Words\",\"toc\":\"On This Page\",\"prev\":\"Prev\",\"next\":\"Next\",\"lastUpdated\":\"Last update\",\"contributors\":\"Contributors\",\"editLink\":\"Edit this page\",\"print\":\"Print\"},\"outlookLocales\":{\"themeColor\":\"Theme Color\",\"darkmode\":\"Theme Mode\",\"fullscreen\":\"Full Screen\"},\"routeLocales\":{\"skipToContent\":\"Skip to main content\",\"notFoundTitle\":\"Page not found\",\"notFoundMsg\":[\"There’s nothing here.\",\"How did we get here?\",\"That’s a Four-Oh-Four.\",\"Looks like we've got some broken links.\"],\"back\":\"Go back\",\"home\":\"Take me home\",\"openInNewWindow\":\"Open in new window\"},\"navbar\":[\"/\",{\"text\":\"Documentations\",\"icon\":\"book\",\"children\":[{\"text\":\"Shortcuts\",\"prefix\":\"/\",\"children\":[\"install\",{\"text\":\"Getting Started\",\"link\":\"tour/\",\"icon\":\"life-ring\"},{\"text\":\"Rosetta Stone\",\"link\":\"rosetta/\",\"icon\":\"language\"},{\"text\":\"User Guide\",\"link\":\"user-guide/\",\"icon\":\"book\"},{\"text\":\"BuiltIns\",\"link\":\"commands/\",\"icon\":\"terminal\"}]}]},\"/changelog\",\"/blog\"],\"sidebar\":{\"/\":[{\"text\":\"Murex\",\"icon\":\"house\",\"children\":[\"/install\",\"/changelog\",\"/blog\",\"/tour\",\"/rosetta\"],\"collapsible\":true},{\"text\":\"User Guide\",\"icon\":\"book\",\"prefix\":\"user-guide/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"Operators And Tokens\",\"icon\":\"equals\",\"prefix\":\"parser/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"Builtins\",\"icon\":\"terminal\",\"prefix\":\"commands/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"Optional Builtins\",\"icon\":\"terminal\",\"prefix\":\"optional/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"Data Types\",\"icon\":\"table\",\"prefix\":\"types/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"Events\",\"icon\":\"bolt\",\"prefix\":\"events/\",\"children\":\"structure\",\"collapsible\":true},{\"text\":\"API Reference\",\"icon\":\"gears\",\"prefix\":\"apis/\",\"children\":\"structure\",\"collapsible\":true}]}}}}")

if (import.meta.webpackHot) {
  import.meta.webpackHot.accept()
  if (__VUE_HMR_RUNTIME__.updateThemeData) {
    __VUE_HMR_RUNTIME__.updateThemeData(themeData)
  }
}

if (import.meta.hot) {
  import.meta.hot.accept(({ themeData }) => {
    __VUE_HMR_RUNTIME__.updateThemeData(themeData)
  })
}