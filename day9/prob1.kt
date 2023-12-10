import java.io.File
import java.io.InputStream


data class Exits (val left: Int, val right: Int)

fun main() {
    // read input
    val inputStream: InputStream = File("input1.txt").inputStream()
    val inputString = inputStream.bufferedReader().use { it.readText() }
    val inputStrings = inputString.split("\n")

    var total = 0
    for (line: String in inputStrings){
        var values = line.split(" ").mapNotNull{it.toIntOrNull()}
        if (values.isNullOrEmpty()){
            continue
        }
        println(values)
        val predictedValue = predictValue(values)
        println(predictedValue)
        total += predictedValue
    }
    // Prepare data structure
    println("Result:  ${total}")
}

fun predictValue(inputList: List<Int>): Int {
    val differenceList = inputList.windowed(2) {it[1] - it[0]} // compute difference list
    println(differenceList)
    if (differenceList.map() {Math.abs(it)}.sum() == 0) {// if the list is all zeros
        return inputList.last()
    }
    return inputList.last() + predictValue(differenceList) // else use recursion
}
