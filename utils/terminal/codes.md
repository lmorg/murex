# Murex Shell Integration — OSC 7, 133 and 633 Reference

This reference is intended for implementing shell integration in Murex, covering the established protocols and VS Code's experimental extensions.

It includes **all documented OSC 7, OSC 133 and OSC 633 operations, plus the undocumented OSC 633 operations identified in VS Code's source code**.

### Escape sequence notation

All examples use the following notation:

| Symbol | Bytes | Description |
|---|---|---|
| `ESC` | `\x1b` | Escape character |
| `BEL` | `\x07` | Bell character |
| `ST` | `\x1b\x5c` | String terminator |
| `OSC` | `ESC ]` | Operating System Command introducer |

For example:

```text
OSC 133 ; A ST
```

Can be emitted in Go using:

```go
fmt.Print("\x1b]133;A\x1b\\")
```

Or using BEL:

```go
fmt.Print("\x1b]133;A\x07")
```

Both terminators are commonly supported.

## 1. OSC 133 — Semantic prompt and command tracking

**Origin:** FinalTerm / iTerm2  
**Status:** Established, widely implemented convention  
**Purpose:** Communicate prompt boundaries and command execution state.

| Code | Description | Example | Status |
|---|---|---|---|
| `133;A` | Prompt begins. Emitted immediately before displaying the shell prompt. | `OSC 133;A ST` | Established |
| `133;B` | Prompt ends. Shell is ready to accept command input. | `OSC 133;B ST` | Established |
| `133;C` | Command execution begins. Emitted after the user submits a command, before its output. | `OSC 133;C ST` | Established |
| `133;D` | Command finishes. No exit status supplied. Also used when no command was executed. | `OSC 133;D ST` | Established |
| `133;D;<exit>` | Command finishes with an exit status. | `OSC 133;D;0 ST` | Established |

**Example lifecycle:**

```text
OSC 133;A ST
murex> 
OSC 133;B ST
echo "hello"
OSC 133;C ST
hello
OSC 133;D;0 ST
OSC 133;A ST
murex>
OSC 133;B ST
```

The terminal can use these markers to identify commands and their output, navigate between commands, and display exit-status indicators.

Source: [VS Code's FinalTerm compatibility documentation](https://code.visualstudio.com/docs/terminal/shell-integration).

## 2. OSC 633 — VS Code shell integration

**Origin:** Microsoft VS Code  
**Status:** Mix of documented and experimental operations  
**Purpose:** Command tracking, command metadata, shell properties, environment reporting and terminal marks.

### 2.1 Documented operations

| Code | Description | Example | Status |
|---|---|---|---|
| `633;A` | Prompt begins. | `OSC 633;A ST` | Documented |
| `633;B` | Prompt ends; input begins. | `OSC 633;B ST` | Documented |
| `633;C` | Command execution begins. | `OSC 633;C ST` | Documented |
| `633;D` | Command finishes without an exit code. | `OSC 633;D ST` | Documented |
| `633;D;<exit>` | Command finishes with exit status. | `OSC 633;D;1 ST` | Documented |
| `633;E;<command>` | Reports the exact command line. | `OSC 633;E;git status ST` | Documented |
| `633;E;<command>;<nonce>` | Reports the command line with a verification nonce. | `OSC 633;E;git status;abc123 ST` | Documented |
| `633;P;Cwd=<path>` | Reports the current working directory. | `OSC 633;P;Cwd=/home/user ST` | Documented |
| `633;P;IsWindows=<bool>` | Indicates whether the shell uses a Windows terminal backend. | `OSC 633;P;IsWindows=False ST` | Documented |
| `633;P;HasRichCommandDetection=<bool>` | Indicates that command boundaries and command-line reporting are reliable. | `OSC 633;P;HasRichCommandDetection=True ST` | Documented |

The expected rich-command sequence is:

```text
A → B → E → C → D
```

The `E` marker should immediately precede `C` when rich command detection is enabled.

### 2.2 Additional and experimental operations

These operations are present in VS Code's implementation but are not part of its stable, publicly documented shell integration contract.

| Code | Description | Example | Status |
|---|---|---|---|
| `633;F` | Continuation prompt begins, such as the secondary prompt for multiline commands. | `OSC 633;F ST` | Experimental |
| `633;G` | Continuation prompt ends. | `OSC 633;G ST` | Experimental |
| `633;H` | Right-hand prompt begins. | `OSC 633;H ST` | Experimental |
| `633;I` | Right-hand prompt ends. | `OSC 633;I ST` | Experimental |
| `633;P;ContinuationPrompt=<text>` | Reports the secondary prompt string. | `OSC 633;P;ContinuationPrompt=> ST` | Implementation-specific |
| `633;P;Prompt=<text>` | Reports the prompt text; used by VS Code's PowerShell integration. | `OSC 633;P;Prompt=PS> ST` | Implementation-specific |
| `633;SetMark` | Creates a navigation mark in the terminal buffer. | `OSC 633;SetMark ST` | Experimental |
| `633;SetMark;Id=<id>` | Creates a mark with an identifier. | `OSC 633;SetMark;Id=build-1 ST` | Experimental |
| `633;SetMark;Hidden` | Creates a hidden mark. | `OSC 633;SetMark;Hidden ST` | Experimental |
| `633;SetMark;Id=<id>;Hidden` | Creates an identified, hidden mark. | `OSC 633;SetMark;Id=build-1;Hidden ST` | Experimental |
| `633;EnvJson;<json>;<nonce>` | Reports environment variables as a JSON object. | `OSC 633;EnvJson;{"FOO":"bar"};abc123 ST` | Experimental |
| `633;EnvSingleStart;<clear>` | Begins a batch of environment variable updates. | `OSC 633;EnvSingleStart;1 ST` | Experimental |
| `633;EnvSingleEntry;<key>;<value>` | Adds or updates an environment variable in the pending environment snapshot. | `OSC 633;EnvSingleEntry;EDITOR;vim ST` | Experimental |
| `633;EnvSingleDelete;<key>;<value>` | Removes an environment variable from the pending snapshot; the value is part of the current protocol format. | `OSC 633;EnvSingleDelete;EDITOR;vim ST` | Experimental |
| `633;EnvSingleEnd` | Completes a batch of environment updates. | `OSC 633;EnvSingleEnd ST` | Experimental |

The `EnvSingleStart`, `EnvSingleEntry`, `EnvSingleDelete` and `EnvSingleEnd` operations also support optional nonce parameters.

**Important:** Microsoft explicitly warns against using the experimental sequences in third-party shell integration scripts because their behaviour and format may change.

Source: [VS Code shellIntegrationAddon.ts](https://github.com/microsoft/vscode/blob/main/src/vs/platform/terminal/common/xterm/shellIntegrationAddon.ts).

## 3. OSC 7 — Current working directory

**Origin:** Shell integration conventions, supported by iTerm2 and other terminals  
**Status:** Established convention  
**Purpose:** Report the current working directory using a file URI.

Unlike OSC 133 and OSC 633, OSC 7 has only one principal operation.

| Code | Description | Example | Status |
|---|---|---|---|
| `7;file://<host>/<path>` | Reports the shell's current working directory and hostname. | `OSC 7;file://localhost/home/user ST` | Established |
| `7;file://<host>/<encoded-path>` | Same operation, with URI percent-encoding for special characters. | `OSC 7;file://localhost/home/user/my%20project ST` | Established |

For example, in Go:

```go
import (
    "fmt"
    "net/url"
    "os"
)

func reportCwd() error {
    cwd, err := os.Getwd()
    if err != nil {
        return err
    }

    hostname, err := os.Hostname()
    if err != nil {
        return err
    }

    uri := url.URL{
        Scheme: "file",
        Host:   hostname,
        Path:   cwd,
    }

    fmt.Printf("\x1b]7;%s\x07", uri.String())
    return nil
}
```

OSC 7 reports the directory; it does not change the shell's working directory.

Source: [iTerm2 escape sequence documentation](https://iterm2.com/documentation-escape-codes.html).

## 4. Additional OSC 133 extensions worth investigating

There is a newer development worth considering for Murex.

Some terminal implementations have extended OSC 133 with additional parameters to distinguish different types of prompts.

These are **implementation-specific extensions**, rather than universally supported parts of the original FinalTerm protocol.

| Code | Description | Example | Status |
|---|---|---|---|
| `133;A;k=i` | Primary interactive prompt begins. | `OSC 133;A;k=i ST` | Extension |
| `133;A;k=s` | Secondary or continuation prompt begins. | `OSC 133;A;k=s ST` | Extension |
| `133;A;k=r` | Right-hand prompt begins. | `OSC 133;A;k=r ST` | Extension |
| `133;A;k=c` | Additional continuation-related prompt kind. | `OSC 133;A;k=c ST` | Extension |
| `133;P;k=r` | Property-style marker associated with a right-hand prompt. | `OSC 133;P;k=r ST` | Extension |
| `133;A;aid=<id>` | Associates a prompt with a shell integration session identifier. | `OSC 133;A;aid=123 ST` | Extension |
| `133;D;aid=<id>` | Finishes a command associated with a particular integration session. | `OSC 133;D;aid=123 ST` | Extension |

These extensions are particularly interesting for nested shells, SSH sessions and terminals that support multiple prompt types.

The `k=` extensions are used by projects such as Reedline, while recent iTerm2 development has introduced support for prompt kinds and nested session identifiers. Their interoperability should be tested before relying on them. [Docs.rs](https://docs.rs/reedline/latest/src/reedline/terminal_extensions/semantic_prompt.rs.html?utm_source=chatgpt.com)

## 5. OSC 633 environment reporting — example lifecycle

Because Murex already has environment variable management, this is potentially useful functionality to standardise.

A complete environment snapshot using the experimental OSC 633 operations might look like:

```text
OSC 633;EnvSingleStart;1;abc123 ST
OSC 633;EnvSingleEntry;EDITOR;vim;abc123 ST
OSC 633;EnvSingleEntry;LANG;en_GB.UTF-8;abc123 ST
OSC 633;EnvSingleEntry;SHELL;/usr/local/bin/murex;abc123 ST
OSC 633;EnvSingleEnd;abc123 ST
```

Here:

- `1` requests clearing the previous cached environment at the start of the batch.
- `EnvSingleEntry` provides individual variables.
- `EnvSingleEnd` completes the update.
- `abc123` represents the nonce supplied by the terminal integration.

Subsequent updates can report only changed variables.

Alternatively, the entire environment can be reported as JSON:

```text
OSC 633;EnvJson;{"EDITOR":"vim","LANG":"en_GB.UTF-8"};abc123 ST
```

In both cases, values must be escaped according to OSC 633's encoding rules.

**Security consideration:** Avoid automatically reporting sensitive environment variables such as API keys, credentials and authentication tokens. Prefer an explicit allowlist of variables the terminal actually needs.

The current VS Code implementation contains these operations, but their formats are not guaranteed to remain stable. [GitHub](https://github.com/microsoft/vscode/blob/main/src/vs/platform/terminal/common/xterm/shellIntegrationAddon.ts?utm_source=chatgpt.com)

## 6. OSC 633 string escaping

For the agent implementing OSC 633, the following encoding rules are important.

| Character | Encoding |
|---|---|
| `;` | `\x3b` |
| Newline | `\x0a` |
| Tab | `\x09` |
| Backslash | `\\` |
| ASCII control characters | `\xHH` |
| Space | `\x20` |

The last row is important: VS Code's specification requires escaping characters with codes **0x20 and below**, including spaces.

For example, the command:

```text
echo "hello; world"
```

Would be encoded as:

```text
OSC 633;E;echo\x20"hello\x3b\x20world" ST
```

The escaping rules are defined in [VS Code's shell integration documentation](https://code.visualstudio.com/docs/terminal/shell-integration).

## 7. Implementation recommendations for Murex

For handing this work to another agent, I'd suggest the following priorities.

| Priority | Feature | Protocol | Recommendation |
|---|---|---|---|
| 1 | Current working directory | OSC 7 | Retain existing implementation |
| 1 | Prompt begins | OSC 133;A | Implement |
| 1 | Prompt ends | OSC 133;B | Implement |
| 1 | Command begins | OSC 133;C | Implement |
| 1 | Command finishes | OSC 133;D | Implement |
| 2 | Exact command line | OSC 633;E | Implement for VS Code |
| 2 | Rich command detection | OSC 633;P | Implement for VS Code |
| 2 | Continuation prompts | OSC 133 `k=s` / OSC 633;F/G | Investigate compatibility |
| 2 | Right-hand prompts | OSC 133 `k=r` / OSC 633;H/I | Investigate compatibility |
| 3 | Environment reporting | OSC 633;EnvJson | Experimental; consider an allowlist |
| 3 | Incremental environment updates | OSC 633;EnvSingle* | Experimental |
| 3 | Terminal marks | OSC 633;SetMark | Experimental |
| 3 | Nested shell identification | OSC 133 `aid=` | Investigate compatibility |

### Important implementation considerations

**Avoid emitting both OSC 133 and OSC 633 lifecycle markers indiscriminately.** VS Code prioritises its own OSC 633 sequences when both are present. Prefer OSC 133 for general terminal compatibility and OSC 633 for VS Code-specific features.

**Keep terminal detection separate from protocol support.** `$TERM_PROGRAM=vscode` is a useful indication that VS Code-specific integration is appropriate, but it is not a general-purpose capability negotiation mechanism.

**Emit sequences at the correct execution boundaries.** Since Murex controls its own parser, prompt renderer and execution lifecycle, the integration can be implemented directly rather than through shell hooks.

**Do not assume OSC 133 means the terminal is ready for arbitrary input.** It identifies shell prompt and command boundaries, not whether another foreground application is waiting for input.

**Treat experimental sequences as optional.** Murex can continue using its existing mechanisms for features where OSC 633 provides only an unstable equivalent.

### Primary references for the implementing agent

The most useful sources are:

- [VS Code — Shell integration specification](https://code.visualstudio.com/docs/terminal/shell-integration)
- [VS Code — Shell integration implementation](https://github.com/microsoft/vscode/blob/main/src/vs/platform/terminal/common/xterm/shellIntegrationAddon.ts)
- [VS Code — Bash integration script](https://github.com/microsoft/vscode/blob/main/src/vs/workbench/contrib/terminal/common/scripts/shellIntegration-bash.sh)
- [VS Code — Zsh integration script](https://github.com/microsoft/vscode/blob/main/src/vs/workbench/contrib/terminal/common/scripts/shellIntegration-rc.zsh)
- [iTerm2 — FinalTerm and OSC 133 specification](https://iterm2.com/documentation-escape-codes.html)
- [Reedline — OSC 133 prompt-kind extensions](https://docs.rs/reedline/latest/src/reedline/terminal_extensions/semantic_prompt.rs.html)

**Overall recommendation:** Standardise Murex's prompt and execution tracking on OSC 133, retain OSC 7 for working-directory reporting, and selectively adopt OSC 633 for richer VS Code integration. Treat the undocumented extensions as candidates for interoperability testing rather than stable replacements for Murex's existing functionality.