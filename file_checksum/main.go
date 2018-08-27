package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// file checksum
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	checksum := h.Sum(nil)
	fmt.Println("file.txt checksum", hex.EncodeToString(checksum))

	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// file.txt
	// head checksum
	head := make([]byte, 10)

	_, err = f.Read(head)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(head))
	hb := sha256.Sum256(head)
	fHead10Checksum := hex.EncodeToString(hb[:])
	fmt.Println("[checksum] head 10 bytes", fHead10Checksum)

	_, err = f.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// tail checksum
	tail := make([]byte, 10)
	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.ReadAt(tail, fi.Size()-int64(len(tail)))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(tail))
	tb := sha256.Sum256(tail)
	fmt.Println("[checksum] tail 10 bytes", hex.EncodeToString(tb[:]))
	fmt.Println()

	// file2.txt
	f2, err := os.Open("file2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	// file checksum
	h = sha256.New()
	if _, err := io.Copy(h, f2); err != nil {
		log.Fatal(err)
	}

	checksum = h.Sum(nil)
	fmt.Println("file2.txt checksum", hex.EncodeToString(checksum))

	_, err = f2.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// head checksum
	head = make([]byte, 10)

	_, err = f2.Read(head)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(head))
	hb = sha256.Sum256(head)
	f2Head10Checksum := hex.EncodeToString(hb[:])
	fmt.Println("[checksum] head 10 bytes", f2Head10Checksum)

	_, err = f2.Seek(0, 0)
	if err != nil {
		log.Fatal(err)
	}

	// tail checksum
	tail = make([]byte, 10)
	fi2, err := f2.Stat()
	if err != nil {
		log.Fatal(err)
	}

	_, err = f2.ReadAt(tail, fi2.Size()-int64(len(tail)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(tail))
	tb = sha256.Sum256(tail)
	fmt.Println("[checksum] tail 10 bytes", hex.EncodeToString(tb[:]))
	fmt.Println()

	// compare head checksum
	if fHead10Checksum == f2Head10Checksum {
		fmt.Println("file.txt head == file2.txt head")
	}
}
