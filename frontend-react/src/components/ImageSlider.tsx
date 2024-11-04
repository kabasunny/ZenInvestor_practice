import React from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css'; // スタイルのインポート
import { Autoplay } from 'swiper/modules'; // モジュールのインポート

interface ImageSliderProps {
  images: string[];
  direction: 'left-to-right' | 'right-to-left';
}

const ImageSlider: React.FC<ImageSliderProps> = ({ images, direction }) => {
  return (
    <Swiper
      modules={[Autoplay]}
      spaceBetween={10}
      slidesPerView={'auto'}
      loop={true}
      autoplay={{
        delay: 0,
        disableOnInteraction: false,
      }}
      speed={6000}
      allowTouchMove={false}
      direction="horizontal"
      style={{ direction: direction === 'left-to-right' ? 'ltr' : 'rtl' }}
    >
      {images.map((image, index) => (
        <SwiperSlide key={index} style={{ width: 'auto' }}>
          <img src={image} alt={`Image ${index}`} className="h-60" />
        </SwiperSlide>
      ))}
    </Swiper>
  );
};

export default ImageSlider;


// import React from 'react';

// interface ImageSliderProps {
//   images: string[];
//   direction: 'left-to-right' | 'right-to-left';
// }

// const ImageSlider: React.FC<ImageSliderProps> = ({ images, direction }) => {
//   if (!images || images.length === 0) {
//     return null; // 画像がない場合は何も表示しない
//   }

//   const sliderClass = direction === 'left-to-right' ? 'animate-slide-left' : 'animate-slide-right';

//   return (
//     <div className="overflow-hidden">
//       <div className={`flex ${sliderClass}`}>
//         {images.map((image, index) => (
//           <img key={index} src={image} alt={`Image ${index}`} className="w-auto h-60" />
//         ))}
//       </div>
//     </div>
//   );
// };

// export default ImageSlider;