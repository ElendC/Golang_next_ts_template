"use client";

import {useEffect, useState} from "react";

export default function About() {

    const [data, setData] = useState({ title: "Loading...", body: "" });

    useEffect(() => {
        fetch('/api/about')
            .then((res) => res.json())
            .then((data) => setData(data));
    }, []);

    return(
    <div>
        <h1 className="font-bold">About Page</h1>
        <p>This is the about page</p>
        <p>Data from backend under here...</p>
        <p>{data.title}</p>
        <p>{data.body}</p>
    </div>
    );
}