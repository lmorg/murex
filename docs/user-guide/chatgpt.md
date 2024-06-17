# ChatGPT integration

> How to enable ChatGPT hints

## Default settings

This is disabled by default.

## Description

While `man` pages are an invaluable resource, they can sometimes be verbose and
in a language that is hard to follow if you're unfamiliar with either that tool
and/or Linux / UNIX man pages in general.

The rise of AI tools like ChatGPT can offer a lifeline for people who want a
more tailored guide to using command line tools. So Murex has a ChatGPT
integration to support those use cases directly from the command line.

It is disabled by default (see below) however once enabled, you can view
ChatGPT's response in the preview screen, by pressing `[f1]`.

## Enabling ChatGPT preview

To enable, you must add your OpenAI key to your profile, which is typically
`~/.murex_profile` (see [profile](profile.md) for more details):

```
config set preview-exec-api xxxxx
```

(where **xxxxx** is replaced with your OpenAI API key)

That API key must have write access to `/v1/chat/completions`. If you're
uncertain what this means, it will become more clear when you create that API
on OpenAI's website.

> By default, ChatGPT integrations are disabled.
> 
> If enabled then the default settings are to only send the command name to
> ChatGPT, rather than the whole command line.

If you're not working with sensitive command lines, then you might also wish
to enable sending the whole command line rather than just the command name.
See [#working-with-sensitive-data](Working with sensitive data) for more
details. 

### Why config

The reason Murex requires this `config` parameter set rather than reading
directly from an environmental variable (typically `$OPENAI_API_KEY`) is so
that your OpenAI API key doesn't get unintentionally leaked to any other
programs, and also so that you explicitly have to opt into enabling this
feature.

### Working with sensitive data

By default, only the command name is sent to ChatGPT. So if your command line
reads something like `rsync medical-records.csv secret-server:/CSVs/`, the only
data leaked if `preview-exec-api` is enabled would be `rsync`packages.json

However you might want more context specific help based on the entirety of your
command line in situations where you know you aren't going to include sensitive
data within your context.

To enable that context to be included in your ChatGPT prompt, add the following
to your profile:

```
config set openai preview-exec-send-cmdline true
```

> Please note that command lines are still stored in `~/.murex_history`! The
> default of disabling or enabling command line context (ie only sending the
> command name) via **preview-exec-send-cmdline** only applies it to OpenAI.
> And even then, that only happens if **preview-exec-api** is _also_ set, which
> it isn't by default.

## GPT model

The **gpt-3.5-turbo** model is used by default because it is quick and cheap.
However you can change that via **preview-exec-model**. For example

```
config set preview-exec-model gpt-4-turbo
```

You may need to clear the cache to see any changes:

```
runtime --clear-cache
```

## Customising the prompt

The default prompt passed to ChatGPT can be viewed via `config get openai
preview-exec-prompt`.

This can be changed should you want something more specific to your needs. For
example:

```
config set openai preview-exec-prompt "You are a medieval pirate transported into the future and gifted the skills of the ${os} command line. Describe how to use UNIX commands in your own, old English pirate, language"
```

...might produce output like the following:

> Event `chatgpt`:                                                                                                                                                                                                                            ┃
> 
> Arrr matey! To use the powerful `reboot` command in the land of Linux, ye
> must open thy terminal and type in `reboot`. This command will set sail the
> process of restarting thy device, just like hoisting the anchor and setting
> off on a new adventure on the high seas. But beware, for all unsaved
> documents will walk the plank if not properly stashed away afore using this
> command. Fair winds to ye!             

You may need to clear the cache to see any changes:

```
runtime --clear-cache
```

## Caching

A successful ChatGPT response is cached for 50 days. You can clear that cache
with `runtime --clear-cache`.

## Show me the code

This feature is written in Murex's own scripting language, uses the `onPreview`
event, and can be changed, modified or even removed completely from within
Murex itself, eg

```
!event onPreview chatgpt
```

### Source code

if { runtime --event-types -> match onPreview } else {
    return
}

/#
    Example request:

        curl https://api.openai.com/v1/chat/completions \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $OPENAI_API_KEY" \
        -d '{
            "model": "gpt-3.5-turbo",
            "messages": [
            {
                "role": "system",
                "content": "You are a helpful assistant."
            },
            {
                "role": "user",
                "content": "Hello!"
            }
            ]
        }'

    Example response:

        {
            "id": "chatcmpl-123",
            "object": "chat.completion",
            "created": 1677652288,
            "model": "gpt-3.5-turbo-0125",
            "system_fingerprint": "fp_44709d6fcb",
            "choices": [{
                "index": 0,
                "message": {
                    "role": "assistant",
                    "content": "\n\nHello there, how may I assist you today?",
                },
                "logprobs": null,
                "finish_reason": "stop"
            }],
            "usage": {
                "prompt_tokens": 9,
                "completion_tokens": 12,
                "total_tokens": 21
            }
        }
#/

config define openai preview-exec-api %{
    Description: "OpenAI API key with write access to /v1/chat/completions"
    DataType:    str
    Default:     ""
}

config define openai preview-exec-prompt %{
    Description: "How ChatGPT should structure it's response for the [f1] preview of external executables"
    DataType:    str
    Default:     %(You are a command line cheat sheet for command line usage. When asked how to use a command
                line tool, you will provide several sections.
                
                The first section is a description of that command.
                
                The second section is only displayed if additional context of the full command line is
                provided. Then your response should then also include specific and detailed descriptions of
                that command line, and specifically any flags included within that command line. You should
                explain in detail what those flags do, why you might want to use them, and any risks that
                might arise because of their usage.
                
                The third section is commonly used examples of that tool, demonstrating other flags that are
                supported by the tool. You should not repeat any flags used in the context command line and
                there should be several examples.
                
                The forth section will be uncommon examples, ideally still likely to be useful but ultimately
                intended to demonstrate the flexibility of that tool. These examples should be other flags
                that differ from both the common examples and also those used in the command line context.

                Each section should be separated by a title.

                Your output should be formatted as plain text in a way that is easy to read from the
                command line.)
}

config define openai preview-exec-model %{
    Description: "ChatGPT model to use"
    DataType:    str
    Default:     "gpt-3.5-turbo"
}

config define openai preview-exec-send-cmdline %{
    Description: "Should Murex also include the full command line for context specific help?"
    DataType:    bool
    Default:     false
}

event onPreview chatgpt=exec {
    cast str

    $.CacheTTL = 0

    config get openai preview-exec-prompt       -> set prompt
    config get openai preview-exec-api          -> set openai_api
    config get openai preview-exec-model        -> set model
    config get openai preview-exec-send-cmdline -> set cmdline

    !if { $openai_api } then {
        out "This feature requires an OpenAI API key with write access to /v1/chat/completions. eg"
        out "config set openai preview-exec-api xxxxx\n"
        out "Run `murex-docs help user-guide/chatgpt` for more details or visit:"
        out "https://murex.rocks/user-guide/chatgpt"
        return
    }

    <stdin> -> set event

    if { $cmdline } then {
        $context = "For context, my full command line is `$(event.Interrupt.CmdLine)`"
    } else {
        $context = ""
    }

    config set http timeout 120
    config set http headers %{
        api.openai.com: {
            Authorization: "Bearer $(openai_api)"
        }
    }

    trypipe {
        %{
            model:    $model
            messages: [
                {
                    role:    system,
                    content: $prompt
                }
                {
                    role:    user,
                    content: "How do I use `$(event.Interrupt.PreviewItem)` in ${os}? $context"
                }
            ]
        } -> post https://api.openai.com/v1/chat/completions -> [ Body ] -> set json body

        if { $body.error } then {
            $body -> pretty

        } else {
            $.CacheTTL = 60 * 60 * 24 * 30  # 50 days
            $body.choices.0.message.content # print output
        }
    }
}

## See Also

* [Profile Files](../user-guide/profile.md):
  A breakdown of the different files loaded on start up
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses
* [`config`](../commands/config.md):
  Query or define Murex runtime settings
* [`murex-docs`](../commands/murex-docs.md):
  Displays the man pages for Murex builtins
* [`onPreview`](../events/onpreview.md):
  Events triggered by changes in state of the interactive shell
* [events](../user-guide/events.md):
  

<hr/>

This document was generated from [gen/user-guide/chatgpt_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/chatgpt_doc.yaml).