import React, { useState, useRef, useEffect } from "react";
import Hunter from "../assets/images/hunter.png";
import Techs from "../assets/images/techs.png";
import Reviews from "../assets/images/reviews.png";
import VideoSection from "./VideoSection";

const Hero: React.FC = () => {
    const [isModalOpen, setIsModalOpen] = useState<boolean>(false);
    const modalRef = useRef<HTMLDivElement | null>(null);

    const openModal = (): void => {
        setIsModalOpen(true);
    };

    const closeModal = (): void => {
        setIsModalOpen(false);
    };

    const handleClickOutside = (event: MouseEvent) => {
        if (modalRef.current && !modalRef.current.contains(event.target as Node)) {
            closeModal();
        }
    };

    useEffect(() => {
        if (isModalOpen) {
            document.addEventListener("mousedown", handleClickOutside);
        }
    }, [isModalOpen]);

    return (
        <main className="flex flex-col items-center mt-[50px] text-center">
            {isModalOpen && (
                <div className="fixed inset-0 flex items-center justify-center z-50 bg-black bg-opacity-75">
                    <div ref={modalRef} className="relative bg-white rounded-lg">
                        <button
                            className="absolute -top-5 -right-5 text-black text-2xl text-white"
                            onClick={closeModal}
                        >
                            <i className="fas fa-times-circle text-md text-gray-300 mr-2"></i>
                        </button>
                        <iframe
                            width="560"
                            height="315"
                            src="https://player.vimeo.com/video/1007338212?badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479"
                            title="YouTube video player"
                            allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                            allowFullScreen
                        ></iframe>
                    </div>
                </div>
            )}
            <h1 className="text-4xl font-bold mb-4 text-white">Make Eye Contact</h1>
            <p className="text-xl mb-8 text-white">
                Leverage AI To Connect With Your Audience <br /> And Share Your Ideas To The World
            </p>

            <div className="flex flex-row space-x-4 mb-8 items-center">
                <img className="h-10" src={Reviews} alt="reviews" />
                <button
                    className="border-[1px] border-white bg-white text-black px-8 py-2 h-[42px] w-[170px] hover:border-black hover:bg-black hover:text-white group"
                    onClick={openModal}
                >
                    <i className="fas fa-play text-sm mr-2 group-hover:text-white"></i>
                    View Demo
                </button>
            </div>

            <VideoSection />

            <div className="flex flex-col items-center mt-20">
                <img className="h-12" src={Techs} alt="techs" />
                <a href="">
                    <img className="h-12 mt-12" src={Hunter} alt="hunter" />
                </a>
            </div>

            <div className="flex mt-4"></div>
        </main>
    );
};

export default Hero;
