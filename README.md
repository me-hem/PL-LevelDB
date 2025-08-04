# PL-LevelDB

A modified implementation of the [LevelDB key/value database](http://code.google.com/p/leveldb) with **Power Law Distribution** compaction table sizing strategy in the [Go programming language](http://golang.org).

## Overview

PL-LevelDB introduces a new compaction strategy that uses power law distribution for determining table sizes across different levels, replacing the traditional exponential growth model. This modification aims to improve read performance.

## Key Features

- **Power Law Compaction Strategy**: Modified table sizing using mathematical power law distribution
- **Improved Performance**: Enhanced throughput and reduced latency for large datasets
- **Backward Compatibility**: Maintains full compatibility with original LevelDB API
- **Configurable Growth Rate**: Tunable exponent parameter for different workload characteristics
- **Comprehensive Benchmarking**: Included performance comparison tools

## Project Structure

```
PL-LevelDB/
├── leveldb/                   # Core LevelDB implementation
│   └── opt/
│       └── options.go         # Modified compaction options with power law distribution
├── test/                      # Performance testing suite
│   └── main.go               # Benchmark tests for different workload sizes
├── LICENSE                    # Project license
├── LevelDB_Stats.txt         # Performance benchmark results
├── README.md                 # Project documentation
├── go.mod                    # Go module dependencies
└── go.sum                    # Go module checksums
```

## Key Modifications

### Power Law Compaction Strategy

The core modifications are in `/leveldb/opt/options.go` where two critical functions implement power law distribution:

#### 1. [GetCompactionTableSize()](https://github.com/me-hem/PL-LevelDB/blob/master/leveldb/opt/options.go#L446): Controls the size of individual SSTable files at each level.


#### 2. [GetCompactionTotalSize()](https://github.com/me-hem/PL-LevelDB/blob/master/leveldb/opt/options.go#L470): Controls the total size limit for each level.



## Installation & Usage

### Requirements
- Go 1.5 or newer

### Installation
```bash
go get github.com/me-hem/PL-LevelDB/leveldb
```

### Basic Usage

```go
package main

import (
    "github.com/me-hem/PL-LevelDB/leveldb"
)

func main() {
    // Open database with power law compaction
    db, err := leveldb.OpenFile("path/to/db", nil)
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Write operations
    err = db.Put([]byte("key"), []byte("value"), nil)
    if err != nil {
        panic(err)
    }
    
    // Read operations
    data, err := db.Get([]byte("key"), nil)
    if err != nil {
        panic(err)
    }
    
    // Delete operations
    err = db.Delete([]byte("key"), nil)
    if err != nil {
        panic(err)
    }
}
```

### Advanced Configuration

```go
import "github.com/me-hem/PL-LevelDB/leveldb/opt"

// Custom options with modified power law parameters
options := &opt.Options{
    CompactionTableSize: 4 * opt.MiB,           // Base table size
    CompactionTotalSize: 20 * opt.MiB,          // Base total size
    // Power law exponent is hardcoded to 1.5 for optimal performance
}

db, err := leveldb.OpenFile("path/to/db", options)
```

## Performance Benchmarks

### Performance Improvements

Power Law implementation shows:
- **2.5% higher throughput** on average
- **2% lower average latency**
- **Better level distribution** reducing deep level formation
- **Improved storage efficiency** with more balanced file sizes

## 🧪 Testing Framework

### Running Benchmarks

```bash
cd test
go run main.go
```

The test suite provides:
- **Configurable workloads**: 100K, 200K, 500K, 1M operations
- **Write performance measurement**: Insertion throughput and latency
- **Read performance measurement**: Query throughput and latency  
- **Database statistics**: Level distribution and compaction metrics
- **Multiple iterations**: Averaged results for statistical accuracy

### Custom Benchmarking

```go
// Modify test/main.go for custom tests
func main() {
    recordCount := 1000000  // Adjust record count
    iterations := 5         // Number of test runs
    
    var totalWrite, totalRead time.Duration
    for i := 0; i < iterations; i++ {
        write, read := testLevelDB(recordCount)
        totalWrite += write
        totalRead += read
    }
    
    fmt.Printf("Avg Write: %v\n", totalWrite/time.Duration(iterations))
    fmt.Printf("Avg Read: %v\n", totalRead/time.Duration(iterations))
}
```

## Technical Analysis

### Level Distribution Comparison

**Traditional LevelDB (GP)**:
```
Level 0: 1 table (3.74MB)
Level 1: 50 tables (98.76MB)  
Level 2: 5 tables (10.06MB)
```

**Power Law LevelDB**:
```
Level 0: 1 table (3.74MB)
Level 2: 4 tables (24.39MB)
Level 3: 6 tables (68.36MB)
Level 4: 1 table (16.09MB)
```

The power law distribution creates more balanced level sizes, reducing the formation of extremely large levels.

### Storage Efficiency

- **Reduced write amplification**: Better level size distribution
- **Improved compaction patterns**: Less frequent deep-level compactions
- **Optimized disk usage**: More uniform file size distribution

## 📚 Documentation

For detailed API documentation, refer to the [original LevelDB documentation](http://godoc.org/github.com/syndtr/goleveldb).


### Migration from Standard LevelDB

```go
// Simply replace the import
// import "github.com/syndtr/goleveldb/leveldb"
import "github.com/me-hem/PL-LevelDB/leveldb"

// No code changes required - drop-in replacement
db, err := leveldb.OpenFile("existing_db", nil)
```

## 📄 License

This project maintains the same BSD-style license as the original LevelDB implementation.
