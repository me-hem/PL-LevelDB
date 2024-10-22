package iterator_test

import (
	"testing"

	"github.com/me-hem/PL-LevelDB/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
