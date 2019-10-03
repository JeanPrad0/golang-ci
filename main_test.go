package main

import "testing"

func TestSum(t *testing.T) {
    got := sum(5,5)

    if got != 10 {
       return  t.Errorf("Erro na soma", got)
    }
    return nil
}