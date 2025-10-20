func findCenter(edges [][]int) int {
    u0 := edges[0][0]
    v0 := edges[0][1]

    u1 := edges[1][0]
    v1 := edges[1][1]

    if u0 == u1 || u0 == v1 {
        return u0
    }

    return v0
}