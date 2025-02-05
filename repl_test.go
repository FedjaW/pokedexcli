package main

import ("testing")

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    } {
        {
            input: " hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "HELLo World",
            expected: []string{"hello", "world"},
        },
        {
            input: "goisnice",
            expected: []string{"goisnice"},
        },
        {
            input: "go is nice is NICE",
            expected: []string{"go", "is", "nice", "is", "nice"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("actual: '%v', but expected: '%v'", word, expectedWord)
                t.Fatal()
            }
        }
    }
}
