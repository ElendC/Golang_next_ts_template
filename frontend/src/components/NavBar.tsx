import Link from "next/link";

export default function NavBar(){
    return(
        /*Whole Navbar outside*/
        <nav className="
        fixed
        z-50
        top-0
        h-14
        w-screen
        flex
        {/*bg-gray-800*/}
        bg-whiteBg
        items-center
        rounded-lg
        shadow-sm
        pr-3
        ">
            {/*Elements inside navbar*/}
            <div className="flex justify-between w-full h-full text-navBarText">
                {/*nav links Left side*/}
                <ul className="flex items-center">
                    <li className="pl-3 hover:text-navBarHover  h-4/5 min-w-12">
                        <Link href="/">
                            <img className="h-full drop-shadow-2xl cursor-pointer"
                                 src="/icons/UiS_Logo.svg" alt="Logo"/>
                        </Link>
                    </li>
                </ul>
                {/*nav links Right side*/}
                <ul className="flex flex-wrap items-center overflow-hidden  max-h-14">
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Search</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Stud. Org</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Nyheter</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">SIS</i></li>
                    <li><i className="m-2 hover:text-navBarHover cursor-pointer">Tilbud</i></li>
                    <li>
                        <Link href={"/authentication"}>
                            <img className="h-full drop-shadow-2xl m-3 hover:text-navBarHover cursor-pointer"
                                 src="/icons/log-in.svg" alt="Logo"/>
                        </Link>
                    </li>
                </ul>
            </div>
        </nav>
    )
}