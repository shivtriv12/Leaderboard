import axios from "axios";
import type { Leaderboard } from "../types";

const baseApiURL = import.meta.env.VITE_BASE_API_URL
const apiURLLeaderboard = baseApiURL+"/api/leaderboard"
const apiURLSearch = baseApiURL+"/api/search"

export const getLeaderboard = async (limit: number=50, cursor: string="")=>{
    const response = await axios.get<Leaderboard>(`${apiURLLeaderboard}?cursor=${cursor}&limit=${limit}`)
    return response.data
}

export const searchUsers = async (limit: number=50, cursor: string="", query: string="")=>{
    const response = await axios.get<Leaderboard>(`${apiURLSearch}?cursor=${cursor}&limit=${limit}&q=${query}`)
    return response.data
}