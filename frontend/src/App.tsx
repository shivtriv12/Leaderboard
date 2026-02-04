import { useEffect, useState, useCallback, useRef } from "react";
import "./App.css";
import { Tabs, type Tab } from "./components/tabs";
import type { User } from "./types";
import { getLeaderboard, searchUsers } from "./services/api";
import { SearchBar } from "./components/searchBar";
import { LeaderboardTable } from "./components/leaderboardTable";

const LIMIT = 50;

function App() {
  const [tab, setTab] = useState<Tab>("Leaderboard");
  const [data, setData] = useState<User[]>([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const searchQueryRef = useRef<string>("");
  const loadingRef = useRef(false);
  const hasMoreRef = useRef(true);
  const cursorRef = useRef("");

  const setLoadingState = (value: boolean) => {
    loadingRef.current = value;
    setLoading(value);
  };

  const resetPagination = (hasMore:boolean = true) => {
    cursorRef.current = "";
    hasMoreRef.current = hasMore;
  };

  const fetchLeaderboard = useCallback(async (reset = false) => {
    if (loadingRef.current || (!hasMoreRef.current && !reset)) return;

    setLoadingState(true);
    setError(null);

    try {
      const resp = await getLeaderboard(LIMIT, reset ? "" : cursorRef.current);

      setData((prev) =>
        reset ? (resp.data ?? []) : [...prev, ...(resp.data ?? [])]
      );
      cursorRef.current = resp.next_cursor ?? "";
      hasMoreRef.current = Boolean(resp.next_cursor);
    } catch (err) {
      setError("Failed to load leaderboard. Please try again.");
      console.error("Leaderboard fetch error:", err);
    } finally {
      setLoadingState(false);
    }
  }, []);

  const fetchSearch = useCallback(async (query: string, reset = false) => {
    if (loadingRef.current || (!hasMoreRef.current && !reset)) return;

    setLoadingState(true);
    setError(null);

    try {
      const resp = await searchUsers(LIMIT, reset ? "" : cursorRef.current, query);
      setData((prev) =>
        reset ? (resp.data ?? []) : [...prev, ...(resp.data ?? [])]
      );
      cursorRef.current = resp.next_cursor ?? "";
      hasMoreRef.current = Boolean(resp.next_cursor);
    } catch (err) {
      setError("Search failed. Please try again.");
      console.error("Search fetch error:", err);
    } finally {
      setLoadingState(false);
    }
  }, []);

  const handleSearch = useCallback((query: string) => {
    if (!query.trim()) {
      setData([]);
      resetPagination(false);
      return;
    }

    searchQueryRef.current = query;
    resetPagination(true);
    setData([]);
    fetchSearch(query, true);
  }, [fetchSearch]);

  useEffect(() => {
    setError(null);
    if (tab === "Leaderboard") {
      setData([]);
      resetPagination(true);
      searchQueryRef.current = "";
      fetchLeaderboard(true);
    } else {
      setData([]);
      resetPagination(false);
      searchQueryRef.current = "";
    }
  }, [tab, fetchLeaderboard]);

  const onScroll = (e: React.UIEvent<HTMLDivElement>) => {
    const { scrollHeight, scrollTop, clientHeight } = e.currentTarget;
    const bottom = scrollHeight - scrollTop <= clientHeight + 200;

    if (bottom && !loadingRef.current && hasMoreRef.current) {
      if (tab === "Leaderboard") {
        fetchLeaderboard();
      } else if (tab === "Search" && searchQueryRef.current) {
        fetchSearch(searchQueryRef.current);
      }
    }
  };

  return (
    <div className="h-screen bg-background text-text flex flex-col">
      <div className="px-6 pt-6">
        <Tabs value={tab} onChange={setTab} />
      </div>

      {tab === "Search" && (
        <SearchBar onSearch={handleSearch} />
      )}

      {tab === "Leaderboard" && (
        <div className="flex items-center gap-3 px-6 mt-4">
          <button
            onClick={() => fetchLeaderboard(true)}
            className="px-4 py-2 rounded-lg bg-surface hover:bg-surface/70 transition cursor-pointer"
          >
            Refresh
          </button>

          <div className="relative group">
            <button className="px-3 py-2 rounded-lg bg-surface cursor-pointer">
              ℹ
            </button>
            <div className="absolute left-0 top-12 w-80 rounded-lg bg-surface p-4 text-sm text-text-secondary opacity-0 group-hover:opacity-100 transition pointer-events-none shadow-xl border border-border z-50">
              <p className="mb-3">
                Scores are updated in the backend every 10 seconds for ~1000
                users. There are 10,000 users total (user_1 → user_10000).
                Refresh to see rating changes.
              </p>
              <p className="text-yellow-500/80 text-xs font-medium bg-yellow-500/5 p-2 rounded">
                Note: Deployed on Render Free Tier. Initial load may take up to 1
                minute due to server cold start.
              </p>
            </div>
          </div>
        </div>
      )}

      {error && (
        <div className="px-6 mt-4 py-3 bg-red-500/20 text-red-400 rounded-lg mx-6">
          {error}
        </div>
      )}

      <div
        className="flex-1 overflow-y-auto mt-4"
        onScroll={onScroll}
      >
        <LeaderboardTable
          data={data}
          emptyText={
            tab === "Search"
              ? "Search for a username"
              : "No rankings available"
          }
        />

        {loading && (
          <div className="py-6 text-center text-text-secondary">
            Loading...
          </div>
        )}
      </div>
    </div>
  );
}

export default App;