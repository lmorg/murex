Output images/vhs-lambda-github-issues-dark.gif
Set Shell murex
Set Theme "Catppuccin Mocha"
#Set Theme "Catppuccin Latte"
Set Height 600
Set Width 1000
Set FontSize 14
Set LineHeight 1.3
Set WindowBar Colorful
Set BorderRadius 0
Set Margin 2
Set MarginFill "#888aaa"
Set Padding 10
Env USER John-Doe
Env HOME "/Users/John-Doe/"
Sleep 1s

Type `# Lets open a JSON document from Github's API and filter only issues created by "lmorg"`
Enter
Type `open https://api.github.com/repos/lmorg/murex/issues | [{ $.v.user.login != "lmorg" }] | set issues`
Sleep 1s
Enter
Sleep 2s
Type `# The content is JSON but it's not very readable...`
Enter
Type `$issues`
Sleep 1s
Enter
Sleep 1s
Type `# ...so lets reformat it into something readable for humans`
Enter
Type `$issues | foreach issue { out "$issue.number: $issue.title [$issue.state]" }`
Sleep 1s
Enter
Sleep 5s
