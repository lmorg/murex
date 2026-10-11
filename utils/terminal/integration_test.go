package terminal

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/lmorg/murex/debug"
	"github.com/lmorg/murex/test/count"
)

func captureOSC(t *testing.T, callback func()) string {
	t.Helper()
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	previous := os.Stdout
	os.Stdout = writer
	defer func() {
		os.Stdout = previous
		reader.Close()
		writer.Close()
	}()
	callback()
	writer.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	return string(output)
}

func TestTtyphoonOutputBlocks(t *testing.T) {
	for _, application := range []string{"ttyphoon", "vscode", "other"} {
		t.Run(application, func(t *testing.T) {
			count.Tests(t, 1)
			t.Setenv("TERM_PROGRAM", "vscode")
			t.Setenv("MXTTY", "false")
			if application == "ttyphoon" {
				t.Setenv("MXTTY", "true")
			} else if application == "other" {
				t.Setenv("TERM_PROGRAM", "other")
			}
			output := captureOSC(t, func() {
				ttyphoonOutputBlockBegin("echo \"hello\"\n", map[string]any{"EDITOR": "vim"})
				ttyphoonOutputBlockEnd(0)
				ttyphoonOutputBlockEnd(12)
			})
			expected := ""
			if application == "ttyphoon" {
				expected = seqApc + `begin;output-block;{"CmdLine":"echo \"hello\"\n","EnvVars":{"EDITOR":"vim"}}` + seqST +
					seqApc + `end;output-block;{"ExitNum":0}` + seqST +
					seqApc + `end;output-block;{"ExitNum":12}` + seqST
			}
			if output != expected {
				t.Fatalf("got %q, want %q", output, expected)
			}
		})
	}
}

func TestLifecycle(t *testing.T) {
	for _, application := range []string{"vscode", "other", "ttyphoon"} {
		t.Run(application, func(t *testing.T) {
			count.Tests(t, 1)
			t.Setenv("TERM_PROGRAM", application)
			t.Setenv("MXTTY", "false")
			code := "133"
			if application == "vscode" {
				code = "633"
			}
			if application == "ttyphoon" {
				t.Setenv("MXTTY", "true")
				t.Setenv("TERM_PROGRAM", "vscode")
			}
			output := captureOSC(t, func() {
				PromptBegin()
				PromptEnd()
				CommandBegin("")
				CommandFinished(-1)
				CommandFinished(0)
				CommandFinished(12)
			})
			var expected string
			for _, marker := range []string{"A", "B", "C", "D", "D;0", "D;12"} {
				if code == "133" && (marker == "A" || strings.HasPrefix(marker, "D")) {
					marker += ";aid=" + shellSessionID()
				}
				expected += "\x1b]" + code + ";" + marker + "\x1b\\"
			}
			if output != expected {
				t.Fatalf("got %q, want %q", output, expected)
			}
		})
	}
}

func TestSessionID(t *testing.T) {
	count.Tests(t, 1)
	identifier := shellSessionID()
	decoded, err := hex.DecodeString(identifier)
	if err != nil || len(decoded) != 16 || shellSessionID() != identifier {
		t.Fatal("session identifier must be stable and contain 16 random bytes")
	}
}

func TestNonceInitialization(t *testing.T) {
	if os.Getenv("MUREX_TEST_NONCE_INIT") == "1" {
		if vscodeNonce != "test-nonce" {
			t.Fatal("nonce was not captured during package initialization")
		}
		if _, present := os.LookupEnv("VSCODE_NONCE"); present {
			t.Fatal("nonce remains in the child process environment")
		}
		return
	}
	count.Tests(t, 1)
	command := exec.Command(os.Args[0], "-test.run=^TestNonceInitialization$")
	for _, entry := range os.Environ() {
		if !strings.HasPrefix(entry, "VSCODE_NONCE=") && !strings.HasPrefix(entry, "MUREX_TEST_NONCE_INIT=") {
			command.Env = append(command.Env, entry)
		}
	}
	command.Env = append(command.Env, "VSCODE_NONCE=test-nonce", "MUREX_TEST_NONCE_INIT=1")
	if output, err := command.CombinedOutput(); err != nil {
		t.Fatalf("nonce initialization failed: %v\n%s", err, output)
	}
}

func TestMissingNonce(t *testing.T) {
	count.Tests(t, 1)
	previousNonce := vscodeNonce
	vscodeNonce = ""
	t.Cleanup(func() { vscodeNonce = previousNonce })
	t.Setenv("TERM_PROGRAM", "vscode")
	t.Setenv("MXTTY", "false")
	output := captureOSC(t, func() {
		vscodeCommandLine("ls")
		vscodeEnvironment(map[string]string{})
		EnvironmentStart(false)
		EnvironmentEntry("KEY", "value")
		EnvironmentDelete("KEY", "value")
		EnvironmentEnd()
		EnvironmentEntry("VSCODE_NONCE", "secret")
		EnvironmentDelete("VSCODE_NONCE", "secret")
	})
	expected := ""
	for _, fields := range []string{"E;ls", "EnvJson;{}", "EnvSingleStart;0", "EnvSingleEntry;KEY;value", "EnvSingleDelete;KEY;value", "EnvSingleEnd"} {
		expected += "\x1b]633;" + fields + "\x1b\\"
	}
	if output != expected {
		t.Fatalf("got %q, want %q", output, expected)
	}
}

func TestVSCodeFeatures(t *testing.T) {
	previousNonce := vscodeNonce
	vscodeNonce = "nonce"
	t.Cleanup(func() { vscodeNonce = previousNonce })
	t.Setenv("TERM_PROGRAM", "vscode")
	t.Setenv("MXTTY", "false")
	testCases := []struct {
		name string
		call func()
		want string
	}{
		{"command", func() { vscodeCommandLine("echo ;\\\n\t\x07\x1b\x7f") }, `E;echo\x20\x3b\\\x0a\x09\x07\x1b\x7f;nonce`},
		{"command short", func() { vscodeCommandLine("ls") }, "E;ls;nonce"},
		{"secondary begin", func() { PromptBegin(PromptOptions{Kind: PromptSecondary}) }, "F"},
		{"secondary end", func() { PromptEnd(PromptSecondary) }, "G"},
		{"right begin", func() { PromptBegin(PromptOptions{Kind: PromptRight}) }, "H"},
		{"right end", func() { PromptEnd(PromptRight) }, "I"},
		{"continuation begin", func() { PromptBegin(PromptOptions{Kind: PromptContinuation}) }, "F"},
		{"continuation end", func() { PromptEnd(PromptContinuation) }, "G"},
		{"windows", func() { WindowsBackend(false) }, "P;IsWindows=False"},
		{"rich detection", func() { RichCommandDetection(true) }, "P;HasRichCommandDetection=True"},
		{"continuation text", func() { ContinuationPrompt("> ") }, `P;ContinuationPrompt=>\x20`},
		{"prompt text", func() { Prompt("murex> ") }, `P;Prompt=murex>\x20`},
		{"mark", func() { SetMark("", false) }, "SetMark"},
		{"named mark", func() { SetMark("build;1", false) }, `SetMark;Id=build\x3b1`},
		{"hidden mark", func() { SetMark("", true) }, "SetMark;Hidden"},
		{"named hidden mark", func() { SetMark("build-1", true) }, "SetMark;Id=build-1;Hidden"},
		{"environment", func() { vscodeEnvironment(map[string]string{"EDITOR": "vim", "VSCODE_NONCE": "secret"}) }, `EnvJson;{"EDITOR":"vim"};nonce`},
		{"environment start", func() { EnvironmentStart(true) }, "EnvSingleStart;1;nonce"},
		{"environment incremental", func() { EnvironmentStart(false) }, "EnvSingleStart;0;nonce"},
		{"environment entry", func() { EnvironmentEntry("KEY", "a;b") }, `EnvSingleEntry;KEY;a\x3bb;nonce`},
		{"environment delete", func() { EnvironmentDelete("KEY", "a;b") }, `EnvSingleDelete;KEY;a\x3bb;nonce`},
		{"environment end", func() { EnvironmentEnd() }, "EnvSingleEnd;nonce"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			count.Tests(t, 1)
			if output := captureOSC(t, testCase.call); output != "\x1b]633;"+testCase.want+"\x1b\\" {
				t.Fatalf("unexpected output: %q", output)
			}
		})
	}
	t.Run("unsupported right property", func(t *testing.T) {
		count.Tests(t, 1)
		if output := captureOSC(t, RightPromptProperty); output != "" {
			t.Fatalf("unexpected output: %q", output)
		}
	})
	t.Setenv("TERM_PROGRAM", "other")
	count.Tests(t, 1)
	if output := captureOSC(t, func() {
		for _, testCase := range testCases {
			if !strings.Contains(testCase.name, "begin") && !strings.Contains(testCase.name, "end") {
				testCase.call()
			}
		}
		EnvironmentEnd()
	}); output != "" {
		t.Fatalf("unsupported features emitted %q", output)
	}
}

func TestPromptExtensions(t *testing.T) {
	count.Tests(t, 2)
	t.Setenv("TERM_PROGRAM", "other")
	t.Setenv("MXTTY", "false")
	output := captureOSC(t, func() {
		for _, kind := range []PromptKind{PromptPrimary, PromptSecondary, PromptRight, PromptContinuation} {
			PromptBegin(PromptOptions{Kind: kind})
		}
		RightPromptProperty()
		CommandFinished(-1)
	})
	expected := ""
	for _, kind := range []string{"i", "s", "r", "c"} {
		expected += "\x1b]133;A;k=" + kind + ";aid=" + shellSessionID() + "\x1b\\"
	}
	expected += "\x1b]133;P;k=r\x1b\\\x1b]133;D;aid=" + shellSessionID() + "\x1b\\"
	if output != expected {
		t.Fatalf("got %q, want %q", output, expected)
	}
	if output := captureOSC(t, func() {
		PromptBegin(PromptOptions{Kind: "invalid"})
		PromptEnd("invalid")
	}); output != "" {
		t.Fatalf("invalid fields emitted %q", output)
	}
}

func TestWorkingDirectory(t *testing.T) {
	count.Tests(t, 2)
	t.Setenv("MXTTY", "false")
	directory := t.TempDir()
	t.Setenv("TERM_PROGRAM", "vscode")
	if output := captureOSC(t, func() { WorkingDirectory(directory) }); output != "\x1b]633;P;Cwd="+escapeOSC(directory)+"\x1b\\" {
		t.Fatalf("unexpected directory output: %q", output)
	}
	t.Setenv("TERM_PROGRAM", "other")
	directory += "/space ;\x07"
	hostname, err := os.Hostname()
	if err != nil {
		t.Fatal(err)
	}
	uri := url.URL{Scheme: "file", Host: hostname, Path: directory}
	if output := captureOSC(t, func() { WorkingDirectory(directory) }); output != "\x1b]7;"+uri.String()+"\x1b\\" {
		t.Fatalf("unexpected directory output: %q", output)
	}
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("test output failure")
}

func TestWriteErrorLogged(t *testing.T) {
	count.Tests(t, 1)
	previousOutput, previousFlags, previousPrefix := log.Writer(), log.Flags(), log.Prefix()
	previousEnabled := debug.Enabled
	defer func() {
		log.SetOutput(previousOutput)
		log.SetFlags(previousFlags)
		log.SetPrefix(previousPrefix)
		debug.Enabled = previousEnabled
	}()
	var output bytes.Buffer
	log.SetOutput(&output)
	log.SetFlags(0)
	log.SetPrefix("")
	debug.Enabled = true
	writeOSC(failingWriter{}, 133, "A")
	if !strings.Contains(output.String(), "test output failure") {
		t.Fatalf("write error not logged: %q", output.String())
	}
}
