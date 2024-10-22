package leveldb

import (
	"testing"

	"github.com/me-hem/PL-LevelDB/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
