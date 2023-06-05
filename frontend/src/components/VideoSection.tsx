import React, { useState, ChangeEvent, MouseEvent } from "react";
import axios from "axios";
import Loader from "./Loader";

const VideoSection: React.FC = () => {
    const [videoFile, setVideoFile] = useState<File | null>(null);
    const [previewURL, setPreviewURL] = useState<string>("");
    const [resultURL, setResultURL] = useState<string>("");
    const [inputFileName, setInputFileName] = useState<string>("");
    const [isLoading, setIsLoading] = useState<boolean>(false);

    const handleFileChange = (event: ChangeEvent<HTMLInputElement>):void => {
        const file = event.target.files?.[0];
        if (file && file.type === "video/mp4") {
        setVideoFile(file);
        setPreviewURL(URL.createObjectURL(file));
        setInputFileName(file.name);
        } else {
        alert('Please upload a valid MP4 video file.');
        }
    };

    const handleReset = ():void => {
        setVideoFile(null);
        setPreviewURL("");
        setInputFileName("");
        setResultURL("");
        setIsLoading(false);
    };

    const handleGenerate = async (): Promise<void> => {
 
        //Uploading
        if (!videoFile) {
            alert("Please upload a video file first.");
            return;
        }
        setIsLoading(true);
        const formData = new FormData();
        formData.append("file", videoFile);

        try {
        await axios.post(
            "http://localhost:8080/data",
            formData,
            {
            headers: { "Content-Type": "multipart/form-data" },
            }
        );
        } catch (error) {
        console.error("Error uploading file:", error);
        }

        //Processing
        if (!videoFile) {
        alert("Please upload a video file first.");
        return;
        }

        try {
        const response = await axios.get(
            `http://localhost:8080/result?file=${videoFile.name}`,
            {
            responseType: "blob",
            }
        );
        const url = URL.createObjectURL(response.data);
        setResultURL(url);
        } catch (error) {
        console.error("Error fetching result:", error);
        }
    };

    return (
        <div className="">
            <button className="border-[1px] font-medium text-sm border-white text-white/80 w-32 h-[42px] px-8 py-2 rounded-full mb-8 hover:border-[2px] hover:text-white" onClick={handleReset}>RESET</button>
            <button className="border-[1px] font-medium text-sm border-white bg-white/30 text-white  w-32 h-[42px] px-4 py-2 ml-4 rounded-full mb-8 hover:bg-white hover:text-black hover:font-semibold hover:border-[1px]" onClick={handleGenerate}>GENERATE</button>
            <div className="flex space-x-8 mb-4 content-boxes ">
                <div>
                    <div className="border-[1px] border-white rounded-t-lg shadow-lg max-w-[400px] max-h-[225px] flex flex-col items-center justify-center relative content-box group" style={{ width: "400px", height: "225px" }}>
                        {!previewURL ? (
                        <>
                            <input type="file" accept="video/mp4" className="absolute inset-0 opacity-0 cursor-pointer" onChange={handleFileChange} />
                            <i className="fas fa-upload text-4xl mb-4 text-white  group-hover:text-gray-200/60"></i>
                            <h2 className="text-xl font-bold mb-2 text-white group-hover:text-gray-200/60">UPLOAD</h2>
                            <p className="text-white group-hover:text-gray-200/60">Drag or drop your video (*.mp4)</p>
                        </>
                        ) : (
                        previewURL && <video controls src={previewURL} />
                        )}
                    </div>

                    <div className="border-[1px]  border-dashed border-gray-300  px-4 rounded-b-md shadow-lg max-w-[400px] max-h-8 flex flex-col items-start justify-center relative content-box" style={{ width: "400px", height: "225px" }}>
                        <div className="flex flex=row">
                            <p className="text-gray-300/80 text-sm">File:///i:/ </p>
                            <p className="text-gray-300/80 text-sm ml-0.2">
                                {inputFileName ? inputFileName : ""}
                            </p>
                        </div>
                    </div>
                </div>

                <div>
                    <div className="border-[1px]  border-white rounded-t-lg  shadow-lg max-w-[400px] max-h-[225px] flex flex-col items-center justify-center relative content-box" style={{ width: "400px", height: "225px" }} >
                        {!resultURL ? (
                        !isLoading ? (
                            <>
                            <i className="fas fa-video text-4xl mb-4 text-white"></i>
                            <p className="text-white">
                                Your processed video will appear here!
                            </p>
                            </>
                        ) : (
                            <Loader />
                        )
                        ) : (
                        resultURL && <video controls src={resultURL} />
                        )}
                    </div>
                    <div className="border-[1px]  border-dashed border-gray-300 px-4 rounded-b-md shadow-lg max-w-[400px] max-h-8 flex flex-col items-start justify-center relative content-box" style={{ width: "400px", height: "225px" }} >
                        <div className="flex flex=row">
                            <p className="text-gray-300/80 text-sm">File:///o:/ </p>
                            <p className="text-gray-300/80 text-sm ml-0.2">
                                {resultURL && resultURL.match(/\/([a-f0-9]{8}-[a-f0-9]{4})/i)?.[1] + ".mp4"}
                            </p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default VideoSection;
