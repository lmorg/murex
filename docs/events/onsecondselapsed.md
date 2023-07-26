# `onSecondsElapsed` - events

> Events triggered by time intervals

## Description

`onSecondsElapsed` events are triggered every _n_ seconds.

## Usage

```
event: onSecondsElapsed name=seconds { code block }

!event: onSecondsElapsed name
```

## Valid Interrupts

* `<seconds>`
    Duration in seconds. eg `60` would be 60 seconds / 1 minute

## Examples

```
event: onSecondsElapsed example=60 {
    out: "60 seconds has passed"
}
```

## Detail

### Payload

The following payload is passed to the function via STDIN:

```
{
    "Name": "",
    "Interrupt": 0
}
```

#### Name

This is the name you specified when defining the event.

#### Interrupt

This is the duration you defined the event to wait for.

## See Also

* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`event`](../commands/event.md):
  Event driven programming for shell scripts