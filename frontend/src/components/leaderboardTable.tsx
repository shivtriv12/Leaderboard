import type { User } from "../types";

interface Props {
  data: User[];
  emptyText?: string;
}

export function LeaderboardTable({
  data,
  emptyText = "No data",
}: Props) {
  return (
    <div>
      <div className="flex px-6 py-3 bg-surface border-b border-border text-xs uppercase text-text-secondary">
        <div className="w-28">Global Rank</div>
        <div className="flex-1">Username</div>
        <div className="w-24 text-right">Rating</div>
      </div>

      {data.length === 0 && (
        <div className="py-12 text-center text-text-secondary">
          {emptyText}
        </div>
      )}

      {data.map((user, i) => (
        <div
          key={`${user.username}`}
          className={`
            flex items-center px-6 py-4
            ${i % 2 === 1 ? "bg-surface/40" : ""}
            hover:bg-surface/70 transition
          `}
        >
          <div
            className={`w-28 font-semibold ${
              user.global_rank <= 3
                ? "text-accent"
                : "text-text"
            }`}
          >
            #{user.global_rank}
          </div>

          <div className="flex-1 truncate">
            {user.username}
          </div>

          <div className="w-24 text-right font-semibold text-accent">
            {user.rating}
          </div>
        </div>
      ))}
    </div>
  );
}
