package caesar

import (
	"bytes"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

var customTests string

func TestMain(m *testing.M) {
	flag.StringVar(&customTests, "custom", "", "input file for custom test cases")
	flag.Parse()
	os.Exit(m.Run())
}

func TestCustom(t *testing.T) {
	if customTests == "" {
		t.Skip("skipping custom tests")
	}
	f, err := os.Open(customTests)
	if err != nil {
		t.Fatalf("could not open custom test input: %v", err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatalf("could not read custom test input: %v", err)
	}
	buf := bytes.NewBuffer(data)
	for {
		// read input
		inp, err := buf.ReadBytes(',')
		if err == io.EOF {
			return
		} else if err != nil {
			t.Fatalf("read error: %v", err)
		}

		// read expected output
		exp, err := buf.ReadBytes('\n')
		if err == io.EOF {
			return
		} else if err != nil {
			t.Fatalf("read error: %v", err)
		}

		if len(inp) < 2 || len(exp) < 2 {
			t.Fatalf("malformed input")
		}

		inString := string(inp[:len(inp)-1])
		expectString := string(exp[:len(exp)-1])

		t.Logf("encoding custom input %q, expecting %q", inString, expectString)

		result := Encode(inString)
		if result != expectString {
			t.Logf("input: %s", inString)
			t.Logf("output: %s", result)
			t.Logf("expected: %s", expectString)
			t.Fail()
		}
	}
	Encode(string(data))
}
