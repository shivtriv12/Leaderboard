export interface User{
    global_rank: number
    username: string
    rating: number
}

export interface Leaderboard{
    data: User[]
    next_cursor: string
}