package codes

import (
	"strconv"
	"strings"
)

type DaitchMokotoffMap struct {
	MaxOutputLength int
	phoneticMap     map[string][3]int

	lastUsage struct {
		keyBuffer []string
		key       string
	}
}

func GetDaitchMokotoffMap() *DaitchMokotoffMap {
	var dm = &DaitchMokotoffMap{}
	dm.MaxOutputLength = 6
	dm.phoneticMap = map[string][3]int{
		"A":             {0, -1, -1},
		"A-I":           {0, 1, -1},
		"A-J":           {0, 1, -1},
		"A-Y":           {0, 1, -1},
		"A-U":           {0, 7, -1},
		"B":             {7, 7, 7},
		"C":             {5, 5, 5},
		"C-1":           {4, 4, 4},
		"C-Z":           {4, 4, 4},
		"C-Z-S":         {4, 4, 4},
		"C-S":           {4, 4, 4},
		"C-S-Z":         {4, 4, 4},
		"C-K":           {5, 5, 5},
		"C-K-1":         {45, 45, 45},
		"C-H":           {5, 5, 5},
		"C-H-1":         {4, 4, 4},
		"C-H-S":         {5, 54, 54},
		"D":             {3, 3, 3},
		"D-T":           {3, 3, 3},
		"D-Z":           {4, 4, 4},
		"D-Z-H":         {4, 4, 4},
		"D-Z-S":         {4, 4, 4},
		"D-S":           {4, 4, 4},
		"D-S-H":         {4, 4, 4},
		"D-S-Z":         {4, 4, 4},
		"D-R":           {0, 0, 0},
		"D-R-S":         {4, 4, 4},
		"D-R-Z":         {4, 4, 4},
		"E":             {0, -1, -1},
		"E-I":           {0, 1, -1},
		"E-J":           {0, 1, -1},
		"E-Y":           {0, 1, -1},
		"E-U":           {1, 1, -1},
		"F":             {7, 7, 7},
		"F-B":           {7, 7, 7},
		"G":             {5, 5, 5},
		"H":             {5, 5, -1},
		"I":             {0, -1, -1},
		"I-A":           {0, -1, -1},
		"J":             {4, 4, 4},
		"K":             {5, 5, 5},
		"K-H":           {5, 5, 5},
		"K-S":           {5, 54, 54},
		"L":             {8, 8, 8},
		"M":             {6, 6, 6},
		"M-N":           {66, 66, 66},
		"N":             {6, 6, 6},
		"N-M":           {66, 66, 66},
		"O":             {0, -1, -1},
		"O-I":           {0, 1, -1},
		"O-J":           {0, 1, -1},
		"O-Y":           {0, 1, -1},
		"P":             {7, 7, 7},
		"P-F":           {7, 7, 7},
		"P-H":           {7, 7, 7},
		"Q":             {5, 5, 5},
		"R":             {9, 9, 9},
		"R-Z":           {94, 94, 94},
		"R-Z-1":         {94, 94, 94},
		"R-S":           {94, 94, 94},
		"R-S-1":         {94, 94, 94},
		"S":             {4, 4, 4},
		"S-Z":           {4, 4, 4},
		"S-Z-T":         {2, 43, 43},
		"S-Z-C-Z":       {2, 4, 4},
		"S-Z-C-S":       {2, 4, 4},
		"S-Z-D":         {2, 43, 43},
		"S-D":           {2, 43, 43},
		"S-T":           {2, 43, 43},
		"S-T-R-Z":       {2, 4, 4},
		"S-T-R-S":       {2, 4, 4},
		"S-T-C-H":       {2, 4, 4},
		"S-T-S-H":       {2, 4, 4},
		"S-T-S-C-H":     {2, 4, 4},
		"S-C":           {2, 4, 4},
		"S-C-H":         {4, 4, 4},
		"S-C-H-T":       {2, 43, 43},
		"S-C-H-T-S-C-H": {2, 4, 4},
		"S-C-H-T-S-H":   {2, 4, 4},
		"S-C-H-T-C-H":   {2, 4, 4},
		"S-C-H-D":       {2, 43, 43},
		"S-H":           {4, 4, 4},
		"S-H-T":         {2, 43, 43},
		"S-H-T-C-H":     {2, 4, 4},
		"S-H-T-S-H":     {2, 4, 4},
		"S-H-C-H":       {2, 4, 4},
		"S-H-D":         {2, 43, 43},
		"T":             {3, 3, 3},
		"T-C":           {4, 4, 4},
		"T-C-H":         {4, 4, 4},
		"T-Z":           {4, 4, 4},
		"T-Z-S":         {4, 4, 4},
		"T-S":           {4, 4, 4},
		"T-S-Z":         {4, 4, 4},
		"T-S-H":         {4, 4, 4},
		"T-S-C-H":       {4, 4, 4},
		"T-T-S":         {4, 4, 4},
		"T-T-S-Z":       {4, 4, 4},
		"T-T-C-H":       {4, 4, 4},
		"T-T-Z":         {4, 4, 4},
		"T-H":           {3, 3, 3},
		"T-R-Z":         {4, 4, 4},
		"T-R-S":         {4, 4, 4},
		"U":             {0, -1, -1},
		"U-E":           {0, -1, -1},
		"U-I":           {0, 1, -1},
		"U-J":           {0, 1, -1},
		"U-Y":           {0, 1, -1},
		"V":             {7, 7, 7},
		"W":             {7, 7, 7},
		"X":             {5, 54, 54},
		"Y":             {1, -1, -1},
		"Z":             {4, 4, 4},
		"Z-D":           {2, 43, 43},
		"Z-D-Z":         {2, 4, 4},
		"Z-D-Z-H":       {2, 4, 4},
		"Z-H":           {4, 4, 4},
		"Z-H-D":         {4, 4, 4},
		"Z-H-D-Z-H":     {2, 4, 4},
		"Z-S":           {4, 4, 4},
		"Z-S-H":         {4, 4, 4},
		"Z-S-C-H":       {4, 4, 4},
	}

	return dm
}

func (dm *DaitchMokotoffMap) MapHasKey(posibleKey string) bool {
	_, ok := dm.phoneticMap[posibleKey]
	if ok {
		return true
	}

	return false
}

func (dm *DaitchMokotoffMap) GetItem(key string) ([3]int, error) {
	if dm.MapHasKey(key) {
		return dm.phoneticMap[key], nil
	}

	return [3]int{}, nil
}

func (dm *DaitchMokotoffMap) Run(str string, isCyrillic bool) string {

	length := len(str)
	str = strings.ToUpper(str)
	outputCode := ""
	prevCode := -2

	var i int

	for i < length {
		dm.clearKeyBuffer()
		dm.insertKeyBuffer(string(str[i]))
		j := 1
		for k := 1; k < 7; k++ {
			nextK := i + k
			if nextK >= length {
				break
			}

			dm.insertKeyBuffer(string(str[nextK]))
			if !dm.MapHasKey(dm.getKey()) {
				dm.removeLastInKeyBuffer()
				break
			}

			j = k + 1
		}

		code := -1
		itemKey := 0
		codeMap, _ := dm.GetItem(dm.getKey())
		if i == 0 {
			code = codeMap[itemKey]
		} else if i+j > length || codeMap[0] != 0 {
			itemKey = 2
		} else {
			itemKey = 1
		}

		if itemKey != 0 {
			codeMap, _ = dm.GetItem(dm.getKey())
			if isCyrillic {
				if dm.MapHasKey(dm.getSecondKey()) {
					codeMap, _ = dm.GetItem(dm.getSecondKey())
					code = codeMap[itemKey]
				} else {
					code = codeMap[itemKey]
				}
			} else {
				code = codeMap[itemKey]
			}
		}

		if code != -1 && code != prevCode {
			outputCode += strconv.Itoa(code)
		}
		prevCode = code
		i += j
	}

	outputLen := len([]rune(outputCode))
	if outputLen > dm.MaxOutputLength {
		outputLen = dm.MaxOutputLength
	}

	return padRight(outputCode[0:outputLen], dm.MaxOutputLength)
}

func (p *DaitchMokotoffMap) getKey() string {
	return strings.Join(p.lastUsage.keyBuffer, "-")
}

func (p *DaitchMokotoffMap) getSecondKey() string {
	return p.getKey() + "-1"
}

func (p *DaitchMokotoffMap) insertKeyBuffer(k string) {
	p.lastUsage.keyBuffer = append(p.lastUsage.keyBuffer, k)
}

func (p *DaitchMokotoffMap) clearKeyBuffer() {
	p.lastUsage.keyBuffer = p.lastUsage.keyBuffer[:0]
}

func (p *DaitchMokotoffMap) getLastInKeyBuffer() string {
	return p.lastUsage.keyBuffer[len(p.lastUsage.keyBuffer)-1]
}

func (p *DaitchMokotoffMap) removeLastInKeyBuffer() {
	lastKey := len(p.lastUsage.keyBuffer) - 1
	p.lastUsage.keyBuffer = p.lastUsage.keyBuffer[:lastKey]
}

func padRight(str string, length int) string {
	for len(str) < length {
		str = str + "0"
	}

	return str
}
