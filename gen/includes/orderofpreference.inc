There is an order of preference for which commands are looked up:
1. Aliases - defined via `alias`. All aliases are global
2. _murex_ functions - defined via `function`. All functions are global
3. private functions - defined via `private`. Private's cannot be global and
   are scoped only to the module or source that defined them. For example, You
   cannot call a private function from the interactive command line
4. variables (dollar prefixed) - declared via `set` or `let`
5. auto-globbing prefix: `@g`
6. murex builtins
7. external executable files 