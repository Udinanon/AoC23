import java.io.File
import java.io.InputStream


data class Exits (val left: Int, val right: Int)
data class Node (val value: Int, val name: String, val exits: Exits, val isEnd: Boolean)
fun main() {
    // read input
    val inputStream: InputStream = File("input1.txt").inputStream()
    val inputString = inputStream.bufferedReader().use { it.readText() }
    val inputStrings = inputString.split("\n")
    // Prepare data structure
    var nodes = mutableMapOf<Int, Node>()
    val startPoints = mutableListOf<Node?>()
    
    // Parse inputs 
    for (line: String in inputStrings.drop(2)){
        val splitLine = line.split("=").map{it.trim().trim('(',')')}
        val entryInt = nodeToInt(splitLine[0])
        val exits = splitLine[1].split(",").map{ nodeToInt(it.trim())}
        val nodeExits = Exits(left = exits[0], right = exits[1])
        val isEnd = splitLine[0][2] == 'Z'
        if (isEnd){
            println(splitLine[0])
        }
        val newNode = Node(value = entryInt, name = splitLine[0], exits = nodeExits, isEnd = isEnd)
        if (splitLine[0][2] == 'A'){
            println(newNode)
            startPoints.add(newNode)
        }
        nodes[entryInt] = newNode
    }
    // move in the graph
    var curr_pos = startPoints
    var nSteps = 0
    var nEnds = 0
    var totalEnds = startPoints.count()
    var done = false
    while (!done) {
        for (char in inputStrings[0]){
            nSteps++
            nEnds = 0
            for (index in curr_pos.indices){
                val pos = curr_pos[index]
                val currNode: Node? = nodes[pos!!.value]
                if (char == 'L'){
                    curr_pos[index] = nodes[currNode!!.exits.left]
                } else {
                    curr_pos[index] = nodes[currNode!!.exits.right]
                }
                if (curr_pos[index]!!.isEnd){
                    nEnds++
                }
            }
            println(nSteps)
            if (curr_pos.count() != totalEnds){
                println("ERROR IN NUMBER OF POSITIONS")
            }
            if (nEnds == totalEnds) {
                println(curr_pos)
                done = true
            }
            //if (nEnds != 0) {println(nEnds)}
        }
    }
    println("Result:  ${nSteps}")
}

fun nodeToInt(node: String): Int{
    return node.toList().map{it - 'A'} // get letters to list of numbers from A
                        .mapIndexed {index, number -> number * Math.pow(26.toDouble(), index.toDouble())} // convert from base26
                        .sum().toInt() //sum together and return
}