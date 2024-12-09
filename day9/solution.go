package day9

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger = log.Default()

func Part1() {
	file, _ := os.Open("day9/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	disk := make([]int, 0)
	isFile := true
	fileNum := 0
	dataCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range strings.Split(line, "") {
			length, _ := strconv.Atoi(ch)
			if isFile {
				for l := 0; l < length; l++ {
					disk = append(disk, fileNum)
					dataCount++
				}
			} else {
				for l := 0; l < length; l++ {
					disk = append(disk, -1)
				}
			}
			isFile = !isFile
			if isFile {
				fileNum++
			}
		}

	}

	newDisk := make([]int, dataCount)
	fromBack := len(disk)
	for i := 0; i < dataCount; i++ {
		if disk[i] != -1 {
			newDisk[i] = disk[i]
		} else {
			for {
				fromBack--
				if disk[fromBack] != -1 {
					newDisk[i] = disk[fromBack]
					break
				}
			}

		}
	}
	sum := 0
	for i, v := range newDisk {
		sum = sum + i*v
	}
	logger.Printf("Day 9, part 1: %d", sum)
}

type block struct {
	length int
	fileId int
	next   *block
	prev   *block
}

func Part2() {
	file, _ := os.Open("day9/input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	isFile := true
	fileNum := 0
	var diskStart *block
	var currBlock *block
	var lastBlock *block
	for scanner.Scan() {
		line := scanner.Text()
		for _, ch := range strings.Split(line, "") {
			length, _ := strconv.Atoi(ch)
			lastBlock = currBlock
			if isFile {
				currBlock = &block{length: length, fileId: fileNum, next: nil, prev: lastBlock}
			} else {
				currBlock = &block{length: length, fileId: -1, next: nil, prev: lastBlock}
			}
			if lastBlock != nil {
				lastBlock.next = currBlock
			}
			if diskStart == nil {
				diskStart = currBlock
			}
			isFile = !isFile
			if isFile {
				fileNum++
			}
		}
	}

	compact(diskStart)
	maxFileId := fileNum
	for fileId := maxFileId; fileId > 0; fileId-- {
		file := findFile(fileId, diskStart)
		freeSpace := findBlockIndexWithSpace(file, diskStart)
		if freeSpace != nil {
			beforeFreeSpace := freeSpace.prev
			afterFreeSpace := freeSpace.next

			fileCopy := &block{
				length: file.length,
				fileId: file.fileId,
				next:   nil,
				prev:   beforeFreeSpace,
			}
			beforeFreeSpace.next = fileCopy

			leftoverSpace := &block{fileId: -1, next: afterFreeSpace, prev: fileCopy, length: freeSpace.length - fileCopy.length}
			afterFreeSpace.prev = leftoverSpace

			fileCopy.next = leftoverSpace

			freeSpace.next = nil
			freeSpace.prev = nil

			file.fileId = -1 //free space
			compact(diskStart)
		}
	}
	disk := toArray(diskStart)

	sum := 0
	for i, v := range disk {
		if v > -1 {
			sum = sum + i*v
		}
	}
	logger.Printf("Day 9, part 2: %d", sum)

}

func toArray(start *block) []int {
	disk := make([]int, 0)
	curr := start
	for curr != nil {
		for l := 0; l < curr.length; l++ {
			disk = append(disk, curr.fileId)
		}
		curr = curr.next
	}

	return disk
}

func compact(diskStart *block) {
	curr := diskStart
	for {
		if curr.next.fileId == curr.fileId {
			curr.length += curr.next.length
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
		if curr.next == nil {
			break
		}
	}
}

func findFile(fileId int, diskStart *block) *block {
	curr := diskStart
	for {
		if curr == nil {
			return nil
		}
		if curr.fileId == fileId {
			return curr
		}
		curr = curr.next
	}
}

func findBlockIndexWithSpace(file *block, diskStart *block) *block {
	curr := diskStart
	for {
		if curr == nil {
			return nil
		}
		if curr.fileId == file.fileId {
			return nil
		}
		if curr.fileId == -1 && curr.length >= file.length {
			return curr
		}
		curr = curr.next
	}
}
