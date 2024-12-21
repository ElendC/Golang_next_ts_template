// app/page.tsx
"use client"; // Enables client-side React code

import React, { useEffect, useState } from "react";
import Login from "@/components/authentication/Login";


export default function Page() {
    // Data is initialized with title and body from the parameters
    // setData is the function that changes the data with new values
    const [data, setData] = useState({ title: "Loading...", body: "" });

    // useEffect: Runs code inside the effect after the component/page renders.
    // Empty array [] at the end, ensures effect only runs once.
    useEffect(() => {
        fetch('/api/', {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json',
            },
        })
            .then((res) => {
                if (!res.ok) {
                    throw new Error('Response from backend was not ok'); // TODO: Remove the error string in production
                }
                return res.json();
            })
            .then((data) => setData(data))
            .catch((error) => console.error('Error: ', error));
    }, []);


    return (
        <main className="p-1">
            <h1>{data.title}</h1>
            <h2>Backend string above and under this header</h2>
            <p>{data.body}</p>
        </main>
    );
}
