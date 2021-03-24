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

#include <bitset>
#include <cmath>
#include <cstring>
#include <fstream>
#include <iostream>
#include <map>
#include <sstream>
#include <vector>

namespace encryption {

#define LINE_SEPARATOR "--------------------------------------------------------------------- "

    /**
     * Compresses the characters into hash table
     *
     * @param input the string as input
     * @param map   the address of your initialized hash map
     *              where data will be written.
     *              Format: for each [char], [vector] [i] = position
     */
    void compress(const std::string &input, std::map<char, std::vector<uint64_t>> *map);

    /**
     * Convert 0s and 1s to uint64_t type
     *
     * @param bin the 0s and 1s string input
     * @return uint64_t out
     */
    uint64_t binary_to_uint(const std::string &bin);

    /**
     * Encrypt Data
     *
     * @param text The text to be encrypted. Make sure there is '\n' at end of the string
     * @returns encrypted text (better to store in a file)
     */
    std::string encrypt(const std::string &text);

    /**
     * Decrypt Data
     *
     * @param encDat The encrypted text to be decrypted. Make sure there is '\n' at end of the string
     * @returns encrypted text (better to store in a file)
     */
    std::string decrypt(const std::string &encDat);
}    // namespace encryption
