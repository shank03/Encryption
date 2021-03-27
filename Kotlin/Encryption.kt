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

import kotlin.math.pow

@ExperimentalUnsignedTypes
object Encryption {

    private const val LINE_SEPARATOR = "--------------------------------------------------------------------- "

    private fun compress(input: String): HashMap<Char, ArrayList<ULong>> {
        val map = HashMap<Char, ArrayList<ULong>>()
        for (i in input.indices) {
            if (map.containsKey(input[i])) {
                map[input[i]]?.add((i + 1).toULong())
            } else {
                val pos = ArrayList<ULong>().also { it.add((i + 1).toULong()) }
                map[input[i]] = pos
            }
        }
        return map
    }

    private fun binaryToULong(bin: String): ULong {
        var out: ULong = 0U
        for (i in bin.indices) {
            if (bin[bin.length - 1 - i] == '1') out += 2.0.pow(i.toDouble()).toULong()
        }
        return out
    }

    private fun formatULong(i: ULong): String = String.format("%64s", i.toString(2)).replace(" ", "0")

    fun encrypt(text: String, key: String): String {
        var out = ""

        var keyMask: ULong = 0U
        for (c: Char in key) {
            keyMask += c.toInt().toULong()
        }

        val tLines = text.split("\n")
        for (line: String in tLines) {
            val data = compress(line)
            val info = line.length.toULong().xor(keyMask)
            out += LINE_SEPARATOR + "\n" + formatULong(info) + " \n"

            for ((_key, value) in data) {
                out += when (_key) {
                    '\n', '\r' -> "nl "
                    '\t' -> "tb "
                    ' ' -> "sp "
                    else -> "$_key "
                }

                var ps: ULong = 0U
                var c: ULong = 0U
                for (pos: ULong in value) {
                    ps = ps.or(pos.shl(8 * c.toInt()))
                    c++
                    if (c.toUInt() == 4U) {
                        ps = ps.or(c.shl(32))
                        out += formatULong(ps.xor(keyMask)) + " "
                        ps = 0U
                        c = 0U
                    }
                }
                if (ps.toUInt() != 0U) {
                    ps = ps.or(c.shl(32))
                }
                out += formatULong(ps.xor(keyMask)) + " \n"
            }
        }
        return out + LINE_SEPARATOR + "\n"
    }

    fun decrypt(input: String, key: String): String {
        var out = ""

        var keyMask: ULong = 0U
        for (c: Char in key) {
            keyMask += c.toInt().toULong()
        }

        val tLines = input.split("\n")
        var intoData = false
        var datLen = 0
        var data: CharArray? = null
        var char = '-'
        for (line: String in tLines) {
            if (line == LINE_SEPARATOR) {
                intoData = true
                if (data != null && datLen != 0) {
                    for (c: Char in data) out += c
                    out += "\n"
                }
                continue
            }
            if (intoData) {
                val len = line.trim().replace("\n", "")
                datLen = binaryToULong(len).xor(keyMask).toInt()
                data = CharArray(datLen)
                intoData = false
            } else {
                var count = 0
                val tTokens = line.replace("\n", "").split(" ")
                for (_token: String in tTokens) {
                    val token = _token.trim()
                    if (token == "") continue
                    if (count == 0) {
                        if (token.length == 1) char = token[0]
                        if (token.length == 2) {
                            if (token == "nl") char = '\n'
                            if (token == "sp") char = ' '
                            if (token == "tb") char = '\t'
                        }
                        count++
                    } else {
                        if (token.length == 64) {
                            val o = binaryToULong(token).xor(keyMask)
                            val numPos = o.shr(32)
                            var i = 0
                            while (i < numPos.toInt()) {
                                val pos = o.shr(8 * i).and(0xFFU)
                                if (data != null) data[pos.toInt() - 1] = char
                                i++
                            }
                        }
                    }
                }
            }
        }
        return out
    }
}