import React, { useState, useEffect } from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import 'swiper/css'; // スタイルのインポート
import { Autoplay } from 'swiper/modules'; // モジュールのインポート

interface ImageSliderProps {
  images: string[];
  direction: 'left-to-right' | 'right-to-left';
}

const ImageSlider: React.FC<ImageSliderProps> = ({ images, direction }) => {
  const [imagesLoaded, setImagesLoaded] = useState<boolean>(false);

  useEffect(() => {
    let loadedCount = 0;
    images.forEach((src) => {
      const img = new Image();
      img.src = src;
      img.onload = () => {
        loadedCount += 1;
        if (loadedCount === images.length) {
          setImagesLoaded(true);
        }
      };
    });
  }, [images]);

  if (!imagesLoaded) {
    return null; // 画像が読み込まれるまで何も表示しない
  }

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
