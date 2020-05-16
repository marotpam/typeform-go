package typeform_test

import (
	"os"
	"testing"
)

var fakeServer *typeformServer

func TestMain(t *testing.M) {
	fakeServer = newFakeTypeformServer()

	code := t.Run()
	fakeServer.close()

	os.Exit(code)
}
