#Output images/vhs-expressive-dark.gif
Output images/vhs-expressive-light.gif
Set Shell murex
#Set Theme "Catppuccin Mocha"
Set Theme "Catppuccin Latte"
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

Type `# Expressions`
Enter

Type `(5 + 5) * 2`
Enter
Sleep 2s

Type `day = datetime( {py}%A )`
Enter
Type `$day`
Enter
Sleep 2s

Type `$day == ${ datetime {py}%A }`
Enter
Sleep 2s


Type `# Parameters`
Enter

Type `echo $PWD[-1] # last element in var`
Enter
Sleep 2s

Type `echo @{ a h[i,o] } # an array of parameters`
Enter
Sleep 2s

Type `echo "The answer to the sum 5 * 5 is" (5 * 5)`
Enter
Sleep 2s


Type `# Variables`
Enter

Type `$PWD # var as a string`
Enter
Sleep 2s

Type `@PWD # var as an array`
Enter
Sleep 2s

Type `$PWD[..3] # first 3 elements (inc root) in var`
Enter
Sleep 2s


Type `# Arrays`

Type `%[ 1, 2, three, four, [alpha, beta], {foo: bar, "first code": "hello world"} ]`
Enter
Sleep 2s

Type `%[1..5]`
Enter
Sleep 2s

Type `%[Monday..Friday]`
Enter
Sleep 2s

Type `%[25-Dec-22..01-Jan-23]`
Enter
Sleep 2s

Type `%[Wednesday..Friday] ~> %[Funday]`
Enter
Sleep 2s


Type `# Maps`
Enter

Type `%{ seasons: [spring..winter], "moon phases": [new moon..waning crescent] }`
Enter
Sleep 2s

Type `%{user: lmorg, sex: male} ~> %{developer: true}`
Enter
Sleep 2s

Sleep 3s