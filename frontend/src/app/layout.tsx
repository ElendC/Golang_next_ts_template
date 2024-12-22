import React from "react";
import "@/styles/globals.css";
import NavBar from "@/components/NavBar";
import Login from "@/components/authentication/Login";
import Tilbud from "@/components/Tillbud";


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
{/*        <div className="absolute inset-0 flex justify-center  mt-28">
            <Tilbud></Tilbud>
        </div>*/}
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>

        {children}


        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>
        <br/><br/><br/><br/><br/><br/><br/><br/><br/>

        </body>
        </html>
    )
}