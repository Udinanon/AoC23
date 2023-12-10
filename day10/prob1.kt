import java.io.File
import java.io.InputStream


data class Node (var neighbors: MutableList<Pair<Int, Int>> = mutableListOf(), 
                 val x: Int = -1, val y: Int = -1)


data class Neighborhood(var N: Pair<Int, Int> = Pair(-1, -1),
                        var S: Pair<Int, Int> = Pair(-1, -1),
                        var W: Pair<Int, Int> = Pair(-1, -1),
                        var E: Pair<Int, Int> = Pair(-1, -1))

fun main() {
    // read input
    val inputStream: InputStream = File("input1.txt").inputStream()
    val inputString = inputStream.bufferedReader().use { it.readText() }
    val inputStrings = inputString.split("\n")
    val yMax = inputStrings.count()
    val xMax = inputStrings[0].count()
    var network: Array<Array<Node>> = Array (inputStrings.count()) { 
        Array (inputStrings[0].count()) {
            Node()
        }
    }
    var x = 0
    var y = 0
    var startPos: Pair<Int, Int> = Pair(-1, -1)

    for (line: String in inputStrings){

        for (step: Char in line.toList()){
            val newNode:Node = Node(x = x, y = y)
            val neighborhood: Neighborhood = Neighborhood()
            if (y > 0)      { neighborhood.N = Pair(x, y-1) }
            if (y < yMax-1) { neighborhood.S = Pair(x, y+1) }
            if (x > 0)      { neighborhood.W = Pair(x-1, y) }
            if (x < xMax-1) { neighborhood.E = Pair(x+1, y) }
            when (step){
                '|' -> newNode.neighbors = mutableListOf(neighborhood.N, neighborhood.S)
                '-' -> newNode.neighbors = mutableListOf(neighborhood.W, neighborhood.E) 
                'L' -> newNode.neighbors = mutableListOf(neighborhood.N, neighborhood.E)
                'J' -> newNode.neighbors = mutableListOf(neighborhood.N, neighborhood.W)
                '7' -> newNode.neighbors = mutableListOf(neighborhood.S, neighborhood.W)
                'F' -> newNode.neighbors = mutableListOf(neighborhood.E, neighborhood.S)
                '.' -> {}
                'S' -> {
                    startPos = Pair(x, y)
                }
                else -> {}
            }
            network[y][x] = newNode
            x++   
        }
        x = 0
        y++
    }
    //  Correctly identify the Starting Node stype
    val startNode = network[startPos.second][startPos.first]
    println(startPos)
    println(network[startPos.second-1][startPos.first].neighbors)
    println(network[startPos.second][startPos.first-1].neighbors)
    if (startPos in network[startPos.second-1][startPos.first].neighbors){
        startNode.neighbors.add(Pair(startPos.first, startPos.second-1))
    }
    if (startPos in network[startPos.second][startPos.first-1].neighbors){
        startNode.neighbors.add(Pair(startPos.first-1, startPos.second))
    }
    if (startPos in network[startPos.second+1][startPos.first].neighbors){
        startNode.neighbors.add(Pair(startPos.first, startPos.second+1))
    }
    if (startPos in network[startPos.second][startPos.first+1].neighbors){
        startNode.neighbors.add(Pair(startPos.first+1, startPos.second))
    }
    println(startNode)

    var done = false
    var currNode = startNode
    var prevNode = startNode
    var newNode: Node
    var counter = 0
    while (!done){
        counter++
        newNode = updatePos(currNode, prevNode, network)
        if (newNode == startNode){
            done = true
        }
        prevNode = currNode
        currNode = newNode
    }

    
    // Prepare data structure
    println("Result:  ${counter/2}")
}

fun updatePos(currNode: Node, prevNode: Node, network: Array<Array<Node>>): Node{
    //println("CURR: ${currNode}")
    //println("PREV: ${prevNode}")
    if (currNode.neighbors[0] == Pair(prevNode.x, prevNode.y)) {
        val (newNodeX, newNodeY) = currNode.neighbors[1]
        return network[newNodeY][newNodeX]
    }
    val (newNodeX, newNodeY) = currNode.neighbors[0]
    return network[newNodeY][newNodeX]

}