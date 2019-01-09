package main

func bfs(Node root) {
	if root == nil {
		return
	}
	visit(root)
	root.visited = true
	for root.Left && root.Right {
		if n.visited == false {
			search(n)
		}
	}
}
