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
#include <iostream>

void showOptions(int mode);

int main() {
    int option;
    std::cout << "[1] Encrypt\n";
    std::cout << "[2] Decrypt\n";
    std::cout << "Select an option:";
    std::cin >> option;

    while (option != 1 && option != 2) {
        std::cout << "Invalid option\n";
        std::cout << "Select an option:";
        std::cin >> option;
    }
    showOptions(option);
    return 0;
}

void showOptions(int mode) {
    int option;
    if (mode == 1) {
        std::cout << "\n[1] Encrypt text from file\n";
        std::cout << "[2] Custom input (single line input)\n";
    }
    if (mode == 2) {
        std::cout << "\n[1] Decrypt text from file\n";
    }
    std::cout << "Select an option:";
    std::cin >> option;

    if (mode == 1) {
        while (option != 1 && option != 2) {
            std::cout << "Invalid option\n";
            std::cout << "Select an option:";
            std::cin >> option;
        }

        std::string path;
        std::cout << "Enter file name or path to save encrypted text:\n";
        std::getline(std::cin.ignore(), path);
        std::ofstream file(path);
        while (!file.is_open()) {
            std::cout << "Error opening file/path\n";
            std::cout << "Enter file name or path to save encrypted text:\n";
            std::getline(std::cin, path);
            file = std::ofstream(path);
        }

        if (option == 1) {
            std::string dat;
            std::cout << "Enter file name to encrypt text from:\n";
            std::getline(std::cin, dat);
            std::ifstream rFileDat(dat);
            while (!rFileDat.is_open()) {
                std::cout << "Error opening file/path\n";
                std::cout << "Enter file name to encrypt text from:\n";
                std::getline(std::cin, dat);
                rFileDat = std::ifstream(dat);
            }

            std::stringstream data;
            std::string line;
            while (std::getline(rFileDat, line)) {
                data << line << "\n";
            }
            rFileDat.close();
            std::string encDat = encryption::encrypt(data.str());
            file << encDat;
            file.close();
            std::cout << "Encrypted text stored in " << path;
        }
        if (option == 2) {
            std::string input;
            std::cout << "Enter input:";
            std::getline(std::cin, input);
            while (input.empty()) {
                std::cout << "Empty\nEnter input:";
                std::getline(std::cin, input);
            }

            std::string encDat = encryption::encrypt(input + "\n");
            file << encDat;
            file.close();
            std::cout << "Encrypted text stored in " << path;
        }
    }
    if (mode == 2) {
        while (option != 1) {
            std::cout << "Invalid option\n";
            std::cout << "Select an option:";
            std::cin >> option;
        }

        std::string dat;
        std::cout << "Enter file name to decrypt text from:\n";
        std::getline(std::cin.ignore(), dat);
        std::ifstream rFileDat(dat);
        while (!rFileDat.is_open()) {
            std::cout << "Error opening file/path\n";
            std::cout << "Enter file name to decrypt text from:\n";
            std::getline(std::cin, dat);
            rFileDat = std::ifstream(dat);
        }

        std::stringstream ss;
        std::string line;
        while (std::getline(rFileDat, line)) {
            ss << line << "\n";
        }
        std::string decDat = encryption::decrypt(ss.str());
        std::cout << "Output:\n"
                  << decDat;
    }
}
