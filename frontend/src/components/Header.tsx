import React from "react"
import Logo from "../assets/images/logo.png"

const Header: React.FC = () => {
    return (
        <nav className="bg-black text-white p-4 flex justify-between items-center navbar-padding shadow-lg shadow-gray-800/70">
            <div className="flex items-center logo-margin">
                <img src={Logo} alt="gAIze logo" className="h-8 w-8 mr-2" />
                <span className="text-xl font-bold">gAIze</span>
            </div>
            <div className="space-x-6 navbar-content">
                <a href="#" className="hover:underline">Features</a>
                <a href="#" className="hover:underline">Solutions</a>
                <a href="#" className="hover:underline">Resources</a>
                <a href="#" className="hover:underline">For Teams</a>
                <a href="#" className="hover:underline">Contact</a>
                <a href="#" className="hover:underline">For Developers</a>
            </div>
        <button className="bg-white/30 font-medium text-sm border-white  text-white  w-32 h-[42px] px-8 py-2  rounded-full login-margin  border-[1px] hover:bg-white hover:text-black hover:border-[2px]" > LOGIN </button>
        </nav>
    )
}

export default Header