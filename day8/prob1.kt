import java.io.File
import java.io.InputStream


data class Exits (val left: Int, val right: Int)

fun main() {
    // read input
    val inputStream: InputStream = File("input1.txt").inputStream()
    val inputString = inputStream.bufferedReader().use { it.readText() }
    val inputStrings = inputString.split("\n")
    // Prepare data structure
    var nodes = mutableMapOf<Int, Exits>()
    val start_point = nodeToInt("AAA")
    val end_point = nodeToInt("ZZZ")
    // Parse inputs 
    for (line: String in inputStrings.drop(2)){
        val split_line = line.split("=").map{it.trim().trim('(',')')}
        val entry_int = nodeToInt(split_line[0])
        val exits = split_line[1].split(",").map{ nodeToInt(it.trim())}
        nodes[entry_int] = Exits(left = exits[0], right = exits[1])
    }
    // move in the graph
    var curr_pos = start_point
    var nSteps = 0
    var done = false
    while (!done) {
        for (char in inputStrings[0]){
            nSteps++
            val curr_node = nodes[curr_pos]
            if (char == 'L'){
                curr_pos = curr_node!!.left
            } else {
                curr_pos = curr_node!!.right
            }
            if (curr_pos == end_point) {
                done = true
            }
        }
    }
    println("Result:  ${nSteps}")
}

fun nodeToInt(node: String): Int{
    return node.toList().map{it - 'A'} // get letters to list of numbers from A
                        .mapIndexed {index, number -> number * Math.pow(26.toDouble(), index.toDouble())} // convert from base26
                        .sum().toInt() //sum together and return
}