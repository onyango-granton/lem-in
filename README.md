# Lem-in

## Overview
Lem-in is a pathfinding algorithm project from **School 42**, where the goal is to efficiently move ants from a **start room** to an **end room** through a network of interconnected rooms and tunnels. The challenge is to optimize the movement to achieve the minimum number of steps.

## Features
- Parses a graph-based input representing rooms, tunnels, and ants.
- Implements **BFS (Breadth-First Search)** and **max-flow algorithms** for pathfinding.
- Distributes ants across multiple paths for efficiency.
- Outputs the movement of ants in the correct format.

## Installation
Clone the repository:
```sh
 git clone https://github.com/onyango-granton/lem-in.git
 cd lem-in
```

## Usage
Run the program and provide a valid map as input:
```sh
 go run main.go < maps/example.map
```

## Input Format
- The first line contains the number of ants.
- Rooms are defined by `name x y` (e.g., `A 2 3`).
- Links are defined by `room1-room2` (e.g., `A-B`).
- Special rooms are marked as `##start` and `##end`.

### Example Input
```
10
##start
A 2 3
B 3 4
C 4 5
##end
D 5 6
A-B
B-C
C-D
```

## Algorithm
1. **Parse the input** into a graph structure.
2. **Find the shortest paths** using BFS.
3. **Use flow algorithms** to determine the best paths for multiple ants.
4. **Output the ant movements** in the required format.

## Author
**Granton Onyango**  
GitHub: [onyango-granton](https://github.com/onyango-granton)
LinkedIn: [Granton Onyango](https://www.linkedin.com/in/granton-onyango-298ba6213/)

## License
This project is licensed under the MIT License.

