package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/me-hem/PL-LevelDB/leveldb"
)

func main() {
	fmt.Printf("Testing with %d records... (PL)\n", 1000000)
	var writeDuration, readDuration time.Duration
	for i := 0; i < 5; i++ {
		err := os.RemoveAll("./test_dldb")
		if err != nil {
			fmt.Println("")
		} else {
			fmt.Println("DB successfully deleted")
		}
		wr, rd := testLevelDB(1000000)
		writeDuration += wr
		readDuration += rd

	}
	fmt.Printf("Write duration for %d records: %v\n", 1000000, writeDuration/5)
	fmt.Printf("Read duration for %d records: %v\n", 1000000, readDuration/5)

	stats()
}

func stats() {
	db, err := leveldb.OpenFile("test_dldb", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	stats, err := db.GetProperty("leveldb.stats")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Stats:")
	fmt.Println(stats)
}

func testLevelDB(count int) (time.Duration, time.Duration) {
	db, err := leveldb.OpenFile("test_dldb", nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Write performance test
	start := time.Now()
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err = db.Put([]byte(key), []byte(value), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	writeDuration := time.Since(start)

	// Read performance test
	start = time.Now()
	for i := count - 1; i > -1; i-- {
		key := fmt.Sprintf("key%d", i)
		_, err = db.Get([]byte(key), nil)
		if err != nil {
			log.Fatal(err)
		}
	}
	readDuration := time.Since(start)

	// Cleanup
	db.Close()

	return writeDuration, readDuration
}
