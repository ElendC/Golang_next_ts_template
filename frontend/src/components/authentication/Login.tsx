"use client";
import {FormEvent} from 'react' // For typechecking
import {useRouter} from 'next/navigation' // Redirecting

export default function LoginComponent(){
    const router = useRouter() //Initializes the UseRouter

    async function handleSubmit(event: FormEvent<HTMLFormElement>){
        event.preventDefault()

        const formData = new FormData(event.currentTarget)
        const userName = formData.get('username')
        const password = formData.get('password')

        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({userName, password})
        })

        if (response.ok){
            router.push('/home')
        } else{
            console.log("Wrong login stuff....")
            return //Handle error later
        }
    }

    return (
        <div className="p-20">
            {/*TODO: try: bg-gray-100*/}
            <div className="flex flex-col justify-center items-center bg-whiteBg max-w-96 min-w-56 rounded-lg shadow-lg p-12 pt-10">
                {/*Header container*/}
                <div className="text-foreground text-3xl font-bold p-5">
                    <h2>Login</h2>
                </div>
                <div>
                    <form onSubmit={handleSubmit} className="flex flex-col items-center">
                        <label htmlFor="userName" className="mb-2 w-full font-bold">User Name</label>
                        <input  type="userName" name="userName" id="userName" placeholder="User Name" required
                                className="w-full p-2 mb-5 rounded-md border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-100"
                        />

                        <label htmlFor="password" className="mb-2 w-full font-bold">Password</label>
                        <input id="password" type="password" name="password" placeholder="Password" required
                               className="w-full p-2 mb-5 rounded-md border border-gray-200 focus:outline-none focus:ring-2 focus:ring-blue-100"
                        />

                        <button type="submit"
                        className="w-full p-2 text-white hover:text-green-100 bg-green-600 rounded-md hover:bg-green-700"
                        >Login</button>
                    </form>
                    <div className="text-blue-600 m-4  flex flex-col items-center">
                        <a className="hover:underline" href="/register"  >Register</a>
                    </div>
                </div>
            </div>

            </div>
    )
}