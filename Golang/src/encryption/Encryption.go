/*
 * Copyright (c) 2021, Shashank Verma <shashank.verma2002@gmail.com>(shank03)
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 */

package main

import (
    "strconv"
    "strings"
)

const lineSeparator = "--------------------------------------------------------------------- "

func compress(input string, m *map[rune][]uint64) {
	for i := 0; i < len(input); i++ {
		_, found := (*m)[rune(input[i])]
		if found {
			(*m)[rune(input[i])] = append((*m)[rune(input[i])], []uint64{uint64(i + 1)}...)
		} else {
			var pos []uint64
			pos = append(pos, []uint64{uint64(i + 1)}...)
			(*m)[rune(input[i])] = pos
		}
	}
}

func binaryToUint(bin string) uint64 {
	out, err := strconv.ParseUint(bin, 10, 64);
	if err == nil {
	    return out
    } else {
        return 0
    }
}

func encrypt(text string, key string) string {
	var outStr string

	var keyMask uint64 = 0
	for i := 0; i < len(key); i++ {
		keyMask += uint64(key[i])
	}

	for _, line := range strings.Split(text, "\n") {
		if line == "" {
			continue
		}
		dataMap := make(map[rune][]uint64)
		compress(line, &dataMap)

		var out = uint64(len(line)) ^ keyMask
		outStr += lineSeparator + "\n" + strconv.FormatUint(out, 10) + " \n"

		for key, pos := range dataMap {
		    outStr += strconv.FormatUint(uint64(key) ^ keyMask, 10) + " "

			var ps, c uint64 = 0, 0
			for i := 0; i < len(pos); i++ {
				ps |= pos[i] << (8 * c)
				c++
				if c == 4 {
					ps |= 4 << 32
					outStr += strconv.FormatUint(ps^keyMask, 10) + " "
					c, ps = 0, 0
				}
			}
			if ps != 0 {
				ps |= c << 32
			}
			outStr += strconv.FormatUint(ps^keyMask, 10) + " \n"
		}
	}
	return outStr + lineSeparator + "\n"
}

func decrypt(input string, key string) string {
	var outStr string

	var keyMask uint64 = 0
	for i := 0; i < len(key); i++ {
		keyMask += uint64(key[i])
	}

	var intoData = false
	var data []rune = nil
	var char uint64
	for _, line := range strings.Split(input, "\n") {
		if line == lineSeparator {
			intoData = true
			if data != nil {
				for i := range data {
					outStr += string(data[i])
				}
			}
			continue
		}
		if intoData {
			datLength := binaryToUint(strings.Split(line, " ")[0]) ^ keyMask
			data = nil
			data = make([]rune, datLength)
			intoData = false
		} else {
			tokens := strings.Split(line, " ")
			for t, tok := range tokens {
				if tok == "" {
					continue
				}
                o := binaryToUint(tok) ^ keyMask
                if t == 0 {
					char = o
				} else {
				    numPos := o >> 32
				    for p := uint64(0); p < numPos; p++ {
				        pos := (o >> (8 * p)) & 0xFF
				        if data != nil {
				            data[pos-1] = rune(char)
				        }
				    }
				}
			}
		}
	}
	return outStr
}
