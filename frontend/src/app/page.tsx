// app/page.tsx
"use client"; // Enables client-side React code

import { useEffect, useState } from "react";


export default function Page() {
  const [data, setData] = useState({ title: "Loading...", body: "" });

  useEffect(() => {
    fetch('/api/home')
        .then((res) => res.json())
        .then((data) => setData(data));
  }, []);

  return (
      <main>
        <h1>{data.title}</h1>
        <h2>Backend string above and under this header</h2>
        <p>{data.body}</p>
      </main>
  );
}
