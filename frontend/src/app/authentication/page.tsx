"use client";

import React, {useState} from "react";
import Login from "@/components/authentication/Login";

export default function About() {

    //component can be either a string or null
    const [component, setComponent] = useState<string | null>(null);

    function handleComponent(value: string){
        if (component === value){
            setComponent(null);
            return;
        }
        setComponent(value);
    }


    return(
        <div>
            <h1 className=" flex font-black justify-center text-4xl p-2 m-5">Velg påloggingsmetode</h1>

            <ul className="flex flex-col items-center">
                <li onClick={() => handleComponent("FEIDELogin")}
                    className="flex flex-col items-center border border-solid w-1/3 shadow-lg mb-5 cursor-pointer">
                    <img src="/logo/Horisontal_Feide.png" alt="logo"
                         className=" pt-10"
                    />
                    <p className="p-10 select-none"
                    >Feide is the chosen solution of the Department of Education for secure identification in the education sector.</p>
                </li>

                <li onClick={() => handleComponent("userLogin")}
                    className="flex flex-col items-center border border-solid w-1/3 shadow-lg mb-5 cursor-pointer">
                    <button className=" p-2">
                        With User Name
                    </button>
                </li>
            </ul>
            {component === "userLogin" && (
                <div
                    onClick={() => setComponent(null)} // Closes when clicked on the overlay
                    className="absolute inset-0 flex justify-center items-center bg-black bg-opacity-80"
                >
                    <div onClick={(e) => e.stopPropagation()}> {/*Prevents closing when clicking on the form*/}
                        <Login/>
                    </div>
                </div>
            )}
        </div>
    );
}