# Job Control

> How to manage jobs with Murex

<h2>Table of Contents</h2>

<div id="toc">

- [Overview](#overview)
- [Job Control](#job-control)
  - [Listing executions](#listing-executions)
  - [Background execution](#background-execution)
  - [Foreground execution](#foreground-execution)
  - [Suspension](#suspension)
  - [Termination](#termination)

</div>



## Overview

Job control is a feature that allows users to manage and control multiple executions in a terminal session, it is particularly useful when working with simultaneous jobs as it provides users with the ability to start, stop, pause, resume, and manage the execution.

Murex is very similar to other shells, with the single particularity: builtins are not forked processes like in a traditional POSIX shell but rather virtual threads. This means that you cannot use the typical operating systems command `ps` to list Murex functions.

## Job Control

### Listing executions

This is where `fid-list` (or its alias `jobs`) comes into play. The builtin is used to view all the functions and processes that are managed by the current session.

That includes:

* any aliases within Murex
* public and private Murex functions
* builtins (eg `fid-list` is a builtin command)
* any external processes that were launched from within this shell session
* any background functions or processes of any of the above

```shell
» jobs
PID   State      Background  Process  Parameters
4939  Executing  true        exec     sleep 10000
4996  Executing  true        exec     sleep 10000
5053  Stopped    true        exec     sleep 10000

» fid-list
PID   State      Background  Process  Parameters
4939  Executing  true        exec     sleep 10000
4996  Executing  true        exec     sleep 10000
5053  Stopped    true        exec     sleep 10000
```

Now that we know how to list background jobs, let us review each control operation in more details

### Background execution

Users can easilly start a process in the background with the `bg` builtin. `bg` allows to continue working on other tasks while the process runs independently.

The builtin supports two modes:

1. It can either be run as a marker which executes the function block in the background

```shell
» bg { sleep 5; out "Morning" }
```

2. or it can daemonize stopped job and daemonize it.

```shell
» jobs
PID   State      Background  Process  Parameters
4939  Executing  true        exec     sleep 10000
4996  Executing  true        exec     sleep 10000
5053  Stopped    true        exec     sleep 10000

# Run PID 5053 in the background
# Note that `bg` is context aware, hit TAB to visually select the id
» bg 5053
```

### Foreground execution

Users can bring a background job to the foreground, making it the active task and allowing interaction with it.

```shell
# start 3 background jobs
» bg { sleep 10000; out "Task 1" }
» bg { sleep 10000; out "Task 2" }
» bg { sleep 10000; out "Task 3" }

» jobs
PID   State      Background  Process  Parameters
4939  Executing  true        exec     sleep 10000
4996  Executing  true        exec     sleep 10000
5053  Executing  true        exec     sleep 10000

# bring back one of them to the foreground, it will block on sleep
# Note that `fg` is context aware, hit TAB to visually select the function id
» fg 5053
```

### Suspension

Users can suspend and pause the execution of a running job, which will temporarily halt its progress.

From an interactive session, press `ctrl`+`z` to suspend the currently running job in the foreground.

```shell
» sleep 10000; out "Task 1"
# Hit CTRL Z - terminal should allow new inputs
» jobs
PID   State      Background  Process  Parameters
4944  Executing  true        exec     sleep 10000
# Note how the job has a `paused` state
# from there you can resume execution with either `fg` or `bg`
```

### Termination

Last but not least, users have the option to terminate or halt an execution. This action can be carried out interactively or through built-in functions.

When running an execution in the foreground from an interactive shell, simply press `ctrl`+`c` to terminate the process. This method is straightforward and efficient.

Alternatively, from a scripting perspective, there are two built-in functions that serve the same purpose: `fid-kill`

```shell
» bg { sleep 10000; out "Task 1" }

» jobs
PID   State      Background  Process  Parameters
4944  Executing  true        exec     sleep 10000

» fid-kill 4944
Task 1
```

## See Also

* [Background Process (`bg`)](../commands/bg.md):
  Run processes in the background
* [Display Running Functions (`fid-list`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Display Running Functions (`jobs`)](../commands/fid-list.md):
  Lists all running functions within the current Murex session
* [Execute External Command (`exec`)](../commands/exec.md):
  Runs an executable
* [Execute Shell Function or Builtin (`fexec`)](../commands/fexec.md):
  Execute a command or function, bypassing the usual order of precedence.
* [Foreground Process (`fg`)](../commands/fg.md):
  Sends a background process into the foreground
* [Kill Function (`fid-kill`)](../commands/fid-kill.md):
  Terminate a running Murex function
* [Terminal Hotkeys](../user-guide/terminal-keys.md):
  A list of all the terminal hotkeys and their uses

<hr/>

This document was generated from [gen/user-guide/job-control_doc.yaml](https://github.com/lmorg/murex/blob/master/gen/user-guide/job-control_doc.yaml).