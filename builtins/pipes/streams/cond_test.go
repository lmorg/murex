package streams

import (
    "errors"
    "io"
    "testing"
    "time"
)

// Test that Read blocks until data arrives, then returns it
func TestReadBlocksUntilWrite(t *testing.T) {
    stream := NewStdin()
    // Ensure reader doesn't see EOF immediately
    stream.Open()
    defer stream.Close()

    done := make(chan struct{})
    buf := make([]byte, 5)

    go func() {
        n, err := stream.Read(buf)
        if err != nil {
            t.Errorf("unexpected read error: %v", err)
        }
        if n != 5 || string(buf[:n]) != "hello" {
            t.Errorf("unexpected read: n=%d buf=%q", n, string(buf[:n]))
        }
        close(done)
    }()

    // Give goroutine time to block
    time.Sleep(50 * time.Millisecond)
    if len(stream.buffer) != 0 {
        t.Fatalf("expected empty buffer before write")
    }

    if _, err := stream.Write([]byte("hello")); err != nil {
        t.Fatalf("unexpected write error: %v", err)
    }

    select {
    case <-done:
    case <-time.After(2 * time.Second):
        t.Fatal("read did not unblock after write")
    }
}

// Test that Write blocks when buffer is at max and unblocks after a read frees space
func TestWriteBlocksUntilSpace(t *testing.T) {
    stream := NewStdin()
    stream.Open()
    defer stream.Close()

    stream.mutex.Lock()
    stream.max = 2
    stream.mutex.Unlock()

    // Fill buffer to max
    if _, err := stream.Write([]byte{1, 2}); err != nil {
        t.Fatalf("unexpected initial write error: %v", err)
    }

    started := make(chan struct{})
    done := make(chan struct{})
    go func() {
        close(started)
        if _, err := stream.Write([]byte{3}); err != nil {
            t.Errorf("unexpected write error: %v", err)
        }
        close(done)
    }()

    <-started
    // Give writer time to block
    time.Sleep(50 * time.Millisecond)

    // Now read one byte to free space
    b := make([]byte, 1)
    n, err := stream.Read(b)
    if err != nil || n != 1 {
        t.Fatalf("unexpected read result n=%d err=%v", n, err)
    }

    select {
    case <-done:
    case <-time.After(2 * time.Second):
        t.Fatal("write did not unblock after space freed")
    }
}

// Test that Close unblocks a reader with EOF
func TestCloseUnblocksReaderWithEOF(t *testing.T) {
    stream := NewStdin()
    stream.Open()

    errCh := make(chan error, 1)
    go func() {
        buf := make([]byte, 1)
        _, err := stream.Read(buf)
        errCh <- err
    }()

    // Let reader block, then close last dependent
    time.Sleep(50 * time.Millisecond)
    stream.Close()

    select {
    case err := <-errCh:
        if !errors.Is(err, io.EOF) {
            t.Fatalf("expected EOF, got %v", err)
        }
    case <-time.After(2 * time.Second):
        t.Fatal("reader did not unblock on Close")
    }
}

