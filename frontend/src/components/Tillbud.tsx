"use client";
import {FormEvent} from 'react' // For typechecking
import {useRouter} from 'next/navigation' // Redirecting

export default function Tilbud(){


    return (
        <div
            className="relative w-4/5 h-fit min-w-56
            rounded-lg
            border border-gray-300 bg-white
            shadow-xl
            overflow-y-auto
            ">
            <h1
                className="bg-gray-100
                    text-3xl
                    border rounded-lg
                    p-5
                ">Student Tilbud. Kupponger og mer</h1>

            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>


            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>


            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>
            <h2 className="p-5 border border-gray-300 rounded-md hover:bg-gray-100 transition">Tilbud 1</h2>

            {/*TODO: Add logic to hide component*/}
            <button
                className="absolute top-2 right-2
                text-gray-500 hover:text-gray-700
                hover:bg-gray-200
                rounded-full w-8 h-8 flex items-center justify-center
                transition duration-150 ease-in-out">
                X
            </button>


        </div>
    )
}