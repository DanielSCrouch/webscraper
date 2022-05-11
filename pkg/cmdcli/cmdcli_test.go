package cmdcli

import (
	"log"
	"os"
	"testing"
)

// TestCLI - simple check to ensure registered cmd's func is called
// with correct arguments
func TestCLI(t *testing.T) {

	// Append custom args to imitate cmd-line call
	os.Args = append(os.Args, "test", "foo", "bar")

	// Define simple test function (collects args)
	testRespArgs := []string{}
	testFunc := func(args []string) {
		testRespArgs = append(testRespArgs, args...)
	}

	// Call CLI and validate response
	cli := NewCLI("vidsy", *log.Default())
	cli.RegisterCmd("test", testFunc)

	err := cli.run()
	if err != nil {
		t.Fatal(err)
	}

	if len(testRespArgs) != 2 {
		t.Fatalf("received %d arguments expected 2", len(testRespArgs))
	}

	if testRespArgs[0] != "foo" || testRespArgs[1] != "bar" {
		t.Fatalf("returned arguments do not match expected")
	}
}
