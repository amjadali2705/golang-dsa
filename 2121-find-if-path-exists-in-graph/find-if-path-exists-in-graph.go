func validPath(n int, edges [][]int, source int, destination int) bool {
    if source == destination {
        return true
    }

    adj := make(map[int][]int)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }
    
    visited := make(map[int]bool)
 
    var dfs func(u int) bool
    dfs = func(u int) bool {
        visited[u] = true
        
        if u == destination {
            return true
        }

        for _, v := range adj[u] {
            if !visited[v] {
                if dfs(v) {
                    return true 
                }
            }
        }
        
        return false
    }

    return dfs(source)
}