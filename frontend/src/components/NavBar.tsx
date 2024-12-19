
export default function NavBar(){
    return(
        /*Whole Navbar outside*/
        <nav className="
        fixed
        top-0
        h-14
        w-screen
        flex
        {/*bg-gray-800*/}
        bg-navBarBg
        items-center
        rounded-lg
        shadow-sm
        ">
            {/*Elements inside navbar*/}
            <div className="flex justify-between w-full h-full text-navBarText">
                {/*nav links Left side*/}
                <ul className="flex items-center">
                    <li className="pl-3 hover:text-navBarHover cursor-pointer h-4/5 min-w-12">
                        <img className="h-full drop-shadow-2xl"
                            src="/UiS_Logo.svg" alt="Logo"/>
                    </li>
                </ul>
                {/*nav links Right side*/}
                <ul className="flex flex-wrap items-center overflow-hidden  max-h-14">
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Search</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Home</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">About</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Something</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Something</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer pr-3">Login</i></li>
                </ul>
            </div>
        </nav>
    )
}