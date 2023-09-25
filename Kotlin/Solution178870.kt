package Kotlin

import java.util.*

data class Point(val startIdx: Int, val endIdx: Int) {

    fun gap(): Int {
        return endIdx - startIdx
    }

    fun toIntArray(): IntArray {
        return intArrayOf(startIdx, endIdx)
    }

}

fun solution(sequence: IntArray, k: Int): IntArray {

    var currentK = -1
    var startIdx = 0
    var endIdx = 0
    val corresponds = ArrayList<Point>();

    while (endIdx < sequence.size) {

        if (startIdx == 0 && endIdx == 0 && currentK == -1) {
            currentK = sequence[endIdx]
            continue
        }

        if (currentK == k) {
            corresponds.add(Point(startIdx, endIdx))
            currentK -= sequence[startIdx]
            startIdx++
            endIdx++
            if (endIdx < sequence.size) {
                currentK += sequence[endIdx]
            }
            continue
        }

        if (currentK < k) {
            endIdx++
            if (endIdx < sequence.size) {
                currentK += sequence[endIdx]
            }
            continue
        } else {
            if (startIdx < endIdx) {
                currentK -= sequence[startIdx]
                startIdx++
            } else if (startIdx == endIdx) {
                startIdx++
                endIdx++
                if (endIdx < sequence.size) {
                    currentK += sequence[endIdx]
                }
            }
            continue
        }
    }

    var optimalPointIdx = -1
    var optimalGap = 1000000
    for ((idx, point) in corresponds.withIndex()) {
        if (point.gap() < optimalGap) {
            optimalGap = point.gap()
            optimalPointIdx = idx
        }
    }

    return corresponds[optimalPointIdx].toIntArray()
}

/**
 * << 프로그래머스 연습문제: 연속된 부분 수열의 합 >>
 * @link /learn/courses/30/lessons/178870?language=kotlin
 */
fun main() {
    val sequenceSample1 = intArrayOf(1, 2, 3, 4, 5)
    val kSample1 = 7
//    solution(sequenceSample1, kSample1)

    val sequenceSample2 = intArrayOf(1, 1, 1, 2, 3, 4, 5)
    val kSample2 = 5
//    solution(sequenceSample2, kSample2)

    val sequenceSample3 = intArrayOf(2, 2, 2, 2, 2)
    val kSample3 = 6
    solution(sequenceSample3, kSample3)
}