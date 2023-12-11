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
    var network: Array<Array<Node>> = Array (yMax) { 
        Array (xMax) {
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
    var solution: MutableList<Node> = mutableListOf()
    while (!done){
        counter++
        newNode = updatePos(currNode, prevNode, network)
        if (newNode == startNode){
            done = true
        }
        solution.add(newNode)
        prevNode = currNode
        currNode = newNode
    }
    // Prepare data structure
    // Now we need to:
    // compute which side is in and which is out

    val leftMostStep = solution.minBy { it.x}
    val indexLeftMostStep = solution.indexOf(leftMostStep)
    val successiveStep = solution[indexLeftMostStep + 1]
    println(successiveStep)
    println(leftMostStep)
    // these two should be a t the same X position and only change on Y, as
    val diff = Pair(successiveStep.x - leftMostStep.x, successiveStep.y - leftMostStep.y)
    println(diff)

    enum class Side{LEFT, RIGHT}
    if (diff.second > 0){
        // Then we go down so the inside is left
        val insideSide = Side.LEFT
    } else if (diff.second < 0 ) {
        // Then we go up so the inside is right
        val insideSide = Side.RIGHT
    } else {println("UNEXPECTED")}


    var insideBorder: Array<Pair<Int, Int>> = Array (solution.count()){
        Pair(0, 0)
    }

    insideBorder = computeInside(solution, insideSide)

    // create a 2D matrix to keep track of visited and not counted points
    var insideMatrix: Array<Array<Int>> = Array() (yMax) { 
        Array (xMax) {
            0
        }
    }
    // Easier access to border points
    var solutionMatrix: Array<Array<Int>> = Array() (yMax) { 
        Array (xMax) {
            0
        }
    }
    for (step:Node in solution){
        solutionMatrix[step.y][step.x] = 1
    }

    // follow the loop and run a BFs on each of the inside nodes
    for (step: Node in solution){
        val leftNode = getLeftNode(step)
        BFS(leftNode, visitedMatrix, solutionMatrix)
    }
    // using the matrix to keep track of which ones are inside


    // when we're done we just sum the inside marrtix and return
    return visitedMatrix.sum()
    // just need to actuallly do it lmao

}

fun computeInside(solution: Array<Node>, insideSide: Side): Array<Pair<Int, Int>> {
    for (step:Node in solution){
        val indexStep = solution.indexOf(step)
        val successiveStep = solution[(indexStep + 1 ) % solution.count()]
        println(successiveStep)
        println(leftMostStep)
        // these two should be a t the same X position and only change on Y, as
        val diff = Pair(successiveStep.x - leftMostStep.x, successiveStep.y - leftMostStep.y)
        // compute what is inside 
        

    }    
    
} 

fun BFS(currNode: Node, visitedMatrix: Array<Array<Int>>){}

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