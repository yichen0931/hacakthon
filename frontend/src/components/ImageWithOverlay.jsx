import Image from 'next/image'

function ImageWithOverlay(props) {
    return (
      <div className="relative w-full max-w-md mx-auto">
        {/* Background Image */}
        <Image
          src={props.imgsrc}
          alt="Example"
          className="w-full h-auto object-cover"
        />
        
        {/* Pink Translucent Box */}
        <div className="absolute inset-0 bg-pink-500 opacity-60"></div>
  
        {/* Optional Content on Top */}
        <div className="absolute inset-0 flex items-center justify-center text-white text-4xl font-bold p-10">
          {props.text}
        </div>
      </div>
    );
  }

export default ImageWithOverlay