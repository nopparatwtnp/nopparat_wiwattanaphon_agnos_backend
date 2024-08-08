package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCalculateSteps(t *testing.T) {
    assert.Equal(t, 3, calculateSteps("aA1"))
    assert.Equal(t, 0, calculateSteps("1445D1cd"))
}
