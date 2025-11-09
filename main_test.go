package main

import "testing"

func TestHello(t *testing.T) {
    got := Hello()
    want := "hi zhangjiong 4"
    if got != want {
        t.Fatalf("Hello() = %q, want %q", got, want)
    }
}

