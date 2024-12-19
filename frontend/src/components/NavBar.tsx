
export default function NavBar(){
    return(
        /*Whole Navbar outside*/
        <div className="
        fixed
        top-0
        h-14
        w-screen
        flex
        bg-gray-800
        items-center
        rounded-2xl
        ">
            {/*Elements inside navbar*/}
            <div className="flex justify-between w-full">
                {/*nav links Left side*/}
                <div className="flex">
                    <i className="m-2">Home</i>
                    <i className="m-2">About</i>
                    <i className="m-2">Contact</i>
                </div>
                {/*nav links Right side*/}
                <div className="flex">
                    <i className="m-2">Login</i>
                    <i className="m-2">Logout</i>
                    <i className="m-2">Search</i>
                </div>
            </div>
        </div>
    )
}