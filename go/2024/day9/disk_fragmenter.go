package day9

func partOne(lines []string) int {
	disk, _ := parseDiskMap(lines[0])
	res := 0

	left := 0
	right := len(disk) - 1

	for left < right {
		for disk[left] != nil {
			left += 1
		}
		for disk[right] == nil {
			right -= 1
		}
		if left >= right {
			break
		}
		disk[left] = disk[right]
		disk[right] = nil
		left += 1
		right -= 1
	}

	for i, u := range disk {
		if u == nil {
			break
		}

		res += i * *u
	}

	return res
}

func partTwo(lines []string) int {
	disk, files := parseDiskMap(lines[0])
	res := 0

	fileIdx := len(files) - 1

	for fileIdx >= 0 {
		leftBeg := 0
		leftEnd := 0
		for files[fileIdx][0] > leftBeg {
			for disk[leftBeg] != nil {
				leftBeg += 1
			}
			leftEnd = leftBeg
			for disk[leftEnd] == nil {
				leftEnd += 1
			}
			emptySpaces := leftEnd - leftBeg

			if files[fileIdx][1] > emptySpaces {
				leftBeg = leftEnd
				continue
			}

			if files[fileIdx][0] <= leftBeg {
				break
			}

			i := 0
			fi := fileIdx
			for i < files[fileIdx][1] {
				disk[leftBeg+i] = &fi
				disk[files[fileIdx][0]+i] = nil

				i += 1
			}
			break
		}

		fileIdx -= 1
	}

	for i, u := range disk {
		if u == nil {
			continue
		}

		res += i * *u
	}

	return res
}

const ZeroCharCode = 48

func parseDiskMap(diskMap string) ([]*int, [][]int) {
	disk := make([]*int, 0)
	mode := 0
	fileId := 0
	files := make([][]int, 0)

	for _, c := range diskMap {
		n := c - ZeroCharCode
		fi := fileId
		if mode == 0 {
			files = append(files, []int{len(disk), int(n)})
		}

		for n > 0 {
			if mode == 0 {
				disk = append(disk, &fi)
			} else {
				disk = append(disk, nil)
			}
			n -= 1
		}

		if mode == 0 {
			fileId += 1
			mode = 1
		} else {
			mode = 0
		}
	}
	return disk, files
}
