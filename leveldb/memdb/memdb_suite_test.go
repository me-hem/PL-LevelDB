package memdb

import (
	"testing"

	"github.com/me-hem/PL-LevelDB/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
