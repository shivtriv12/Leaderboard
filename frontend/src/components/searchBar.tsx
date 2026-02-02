import { useEffect, useState } from "react";

interface Props {
  onSearch: (query: string) => void;
}

export function SearchBar({ onSearch }: Props) {
  const [value, setValue] = useState("");

  useEffect(() => {
    const t = setTimeout(() => {
      onSearch(value);
    }, 400);

    return () => clearTimeout(t);
  }, [value, onSearch]);

  return (
    <div className="px-6 mt-4">
      <input
        value={value}
        onChange={(e) => setValue(e.target.value)}
        placeholder="Search username..."
        className="w-full rounded-xl bg-surface px-4 py-3 text-text outline-none border border-border"
      />
    </div>
  );
}