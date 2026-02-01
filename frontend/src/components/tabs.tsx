export type Tab = "Leaderboard" | "Search"

interface TabProps{
    value: Tab
    onChange: (tab:Tab)=>void
}

export function Tabs({value,onChange}:TabProps){
    return <div className="w-full rounded-2xl bg-surface p-1 flex">
        <button
            onClick={() => onChange("Leaderboard")}
            className={`flex-1 rounded-xl py-3 text-sm font-semibold transition-all duration-200 ease-out
            ${
                value === "Leaderboard"
                ? "bg-accent text-background shadow-sm"
                : "text-text-secondary hover:text-text"
            }`}
        >
            Leaderboard
        </button>

        <button
            onClick={() => onChange("Search")}
            className={`flex-1 rounded-xl py-3 text-sm font-semibold transition-all duration-200 ease-out
            ${
                value === "Search"
                ? "bg-accent text-background shadow-sm"
                : "text-text-secondary hover:text-text"
            }`}
        >
            Search
        </button>
    </div>
}