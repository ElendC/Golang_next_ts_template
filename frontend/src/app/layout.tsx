import React from "react";
import "@/styles/globals.css";
import NavBar from "@/components/NavBar";

export default function RootLayout({
                                       children,
                                   }: {
    children: React.ReactNode
}) {
    return (
        <html lang="en">
        <head>
            <title>Make a SEO friendly title here...</title>{/*UIS et eller et annet*/}
            <meta charSet="UTF-8"/>
        </head>
        <body>
        <div className="flex">
            <NavBar/>
        </div>
        <h1 className="font-black">hello</h1>
        {children}</body>
        </html>
    )
}