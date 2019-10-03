package main

import "testing"

func TestSum(t *testing.T) {
    got := sum(5,5)

    want := 10

    if got != want {
      t.Errorf("Erro na soma", got, want)
    }
   
}