### Order of execution

Interrupts are run in alphabetical order. So an event named "alfa" would run
before an event named "zulu". If you are writing multiple events and the order
of execution matters, then you can prefix the names with a number, eg `10_jump`

### Namespacing

This event is namespaced as `$(NAME).$(OPERATION)`.

For example, if an event in `onPrompt` was defined as `example=eof` then its
namespace would be `example.eof` and thus a subsequent event with the same name
but different operation, eg `example=abort`, would not overwrite the former
event defined against the interrupt `eof`.

The reason for this namespacing is because you might legitimately want the same
name for different operations (eg a smart prompt that has elements triggered
from different interrupts).