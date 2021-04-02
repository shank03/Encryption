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

#include "encryption.h"
#include <bitset>
#include <cstring>
#include <sstream>

void encryption::compress(const std::string &input, std::map<char, std::vector<uint64_t>> *map) {
    int len = input.length();
    for (int i = 0; i < len; i++) {
        if (map->find(input[i]) != map->end()) {    // char already exists
            (*map)[input[i]].push_back(i + 1);      // append the position in array
        } else {
            std::vector<uint64_t> pos;
            pos.push_back(i + 1);
            (*map)[input[i]] = pos;    // init map for char [i]
        }
    }
}

uint64_t encryption::binary_to_uint(const std::string &bin) {
    uint64_t out = 0;
    std::istringstream iss(bin);
    iss >> out;
    return out;
}

std::string encryption::encrypt(const std::string &text, const std::string &key) {
    std::stringstream outStr;

    // Key
    uint64_t keyMask = 0;
    for (char i : key) keyMask += keyMask ^ i;

    std::string t_token;
    size_t t_pos, t_index = 0;
    while ((t_pos = text.find('\n', t_index)) != std::string::npos) {
        t_token = text.substr(t_index, t_pos - t_index);

        // Mapping data
        std::map<char, std::vector<uint64_t>> data;
        encryption::compress(t_token, &data);

        uint64_t out = t_token.length() ^ keyMask;
        outStr << LINE_SEPARATOR << "\n"
               << (out) << " \n";

        // Encrypting mapped data
        for (auto &it : data) {
            outStr << ((uint64_t) it.first ^ keyMask) << " ";

            uint64_t ps = 0, c = 0;
            for (uint64_t i : it.second) {
                ps |= i << (8 * c), c++;
                if (c == 4) {
                    ps |= (uint64_t) 4 << 32;
                    outStr << (ps ^ keyMask) << " ";
                    ps = 0, c = 0;
                }
            }
            if (ps != 0) {
                ps |= c << 32;
            }
            outStr << (ps ^ keyMask) << " " << c << " \n";
        }
        t_index = t_pos + 1;
    }
    outStr << LINE_SEPARATOR << "\n";
    return outStr.str();
}

std::string encryption::decrypt(const std::string &encDat, const std::string &key) {
    std::stringstream outStr;

    // Key
    uint64_t keyMask = 0;
    for (char i : key) keyMask += keyMask ^ i;

    bool intoData = false;
    uint64_t datLength = 0;
    char *data = nullptr, c;
    std::string t_token;
    size_t t_pos, t_index = 0;
    while ((t_pos = encDat.find('\n', t_index)) != std::string::npos) {
        t_token = encDat.substr(t_index, t_pos - t_index);
        if (t_token == LINE_SEPARATOR) {
            intoData = true;
            if (data != nullptr && datLength != 0) {
                for (uint64_t i = 0; i < datLength; i++) outStr << data[i];
                outStr << "\n";
            }
            t_index = t_pos + 1;
            continue;
        }
        if (intoData) {
            std::string lenStr = t_token.substr(0, 64);
            datLength = binary_to_uint(lenStr) ^ keyMask;
            data = (char *) malloc(datLength * sizeof(char));
            intoData = false;
        } else {
            size_t p_pos, p_index = 0, count = 0;
            std::string p_token;
            while ((p_pos = t_token.find(' ', p_index)) != std::string::npos) {
                p_token = t_token.substr(p_index, p_pos - p_index);
                if (p_token.empty()) {
                    p_index = p_pos + 1;
                    continue;
                }
                uint64_t o = binary_to_uint(p_token) ^ keyMask;
                if (count == 0) {
                    c = o;
                    count++;
                } else {
                    uint64_t numPos = o >> 32;
                    for (uint64_t p = 0; p < numPos; p++) {
                        uint64_t pos = (o >> (8 * p)) & 0xFF;
                        if (data != nullptr) {
                            data[pos - 1] = c;
                        }
                    }
                }
                p_index = p_pos + 1;
            }
        }
        t_index = t_pos + 1;
    }
    return outStr.str();
}
