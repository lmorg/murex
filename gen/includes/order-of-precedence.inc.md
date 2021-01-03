There is an order of precedence for which commands are looked up:
1. `test` and `pipe` functions because they alter the behavior of the compiler
2. Aliases - defined via `alias`. All aliases are global
3. _murex_ functions - defined via `function`. All functions are global
4. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function from the interactive command line
5. variables (dollar prefixed) - declared via `set` or `let`
6. auto-globbing prefix: `@g`
7. murex builtins
8. external executable files