import { useState } from "react"

function ImageSlider({slides, sliderHeight, sliderWidth}) {
    const [currentIndex, setCurrentIndex] = useState(0);

    const sliderStyle = {
        height: sliderHeight,
        width: sliderWidth,
    }
    const slideStyle = {
        backgroundImage: `url(${slides[currentIndex].url})`
    }

    function goToPrevious() {
        const isFirstSlide = currentIndex === 0;
        const newIndex = isFirstSlide ? slides.length - 1 : currentIndex - 1;
        setCurrentIndex(newIndex);
    }
    function goToNext() {
        const isLastSlide = currentIndex === slides.length - 1;
        const newIndex = isLastSlide ? 0 : currentIndex + 1;
        setCurrentIndex(newIndex);
    }
    return (
        <div className="h-full relative" style={sliderStyle}>
            <div className="w-full h-full bg-center bg-cover rounded-lg" style={slideStyle}></div>
            <div className="absolute top-1/2 left-5 text-3xl text-white cursor-pointer -translate-y-1/2" onClick={goToPrevious}>❰</div>
            <div className="absolute top-1/2 right-5 text-3xl text-white cursor-pointer -translate-y-1/2" onClick={goToNext}>❱</div>
            <div className="absolute w-full bottom shadow-inner rounded-lg bg-slate-300 dark:bg-gray-700 py-0.5 px-0.5">
                <div className="text-center text-lg text-black dark:text-white">{slides[currentIndex].description}</div>
            </div>
        </div>
    )
}

export default ImageSlider