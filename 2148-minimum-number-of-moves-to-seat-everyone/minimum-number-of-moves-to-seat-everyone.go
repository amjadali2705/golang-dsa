func minMovesToSeat(seats []int, students []int) int {
    moves := 0.0

    sort.Ints(seats)
    sort.Ints(students)

    for i := 0; i < len(seats); i++ {
        moves += math.Abs(float64(seats[i]) - float64(students[i]))
    }

    return int(moves)
}