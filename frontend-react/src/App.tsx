// frontend-react\src\App.tsx
import React, { useState, useEffect, useRef } from 'react';
// import axios from 'axios';
import Header from './components/Header';
import Footer from './components/Footer';
import ChartView from './pages/ChartView';
import MarketInsights from './pages/MarketInsights';
import Portfolio from './pages/Portfolio';
import Education from './pages/Education';
import LoginForm from './components/LoginForm';
import { getRandomImage, getMotoSlideImages, getHumanSlideImages, getMoneySlideImages, getCarSlideImages, getSnowboradSlideImages, getAnimeSlideImages } from './components/ImageImporter';
import ImageSlider from './components/ImageSlider';

const App: React.FC = () => {
  // セッションストレージからログイン状態を取得
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(sessionStorage.getItem('isLoggedIn') === 'true');
  
  // ログイン前のバックイメージ
  const [backgroundImage, setBackgroundImage] = useState<string>('');

  // スライド用のイメージ 
  const [motoSlideImages, setMotoSlideImages] = useState<string[]>([]);
  const [humanSlideImages, setHumanSlideImages] = useState<string[]>([]);
  const [moneySlideImages, setMoneySlideImages] = useState<string[]>([]);
  const [carSlideImages, setCarSlideImages] = useState<string[]>([]);
  const [SnowboradSlideImages, setSnowboradSlideImages] = useState<string[]>([]);
  const [animeSlideImages, setAnimeSlideImages] = useState<string[]>([]);

  const dashboardRef = useRef<HTMLDivElement>(null);
  const marketInsightsRef = useRef<HTMLDivElement>(null);
  const portfolioRef = useRef<HTMLDivElement>(null);
  const educationRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const randomImage = getRandomImage();
    console.log("Selected background image: ", randomImage); // デバッグ用メッセージ
    setBackgroundImage(randomImage);
  }, []);

  // スライド用画像データの取得
  // useEffect(() => {
  //   const fetchImages = async () => {
  //     try {
  //       const response = await axios.get<string[]>('/api/images'); // Web API実装後、エンドポイントに置き換える
  //       setImageData(response.data);
  //     } catch (error) {
  //       console.error('スライダー用の画像データの取得に失敗しました:', error);
  //     }
  //   };
  //   // ログインしていない場合に画像データを取得
  //   if (!isLoggedIn) {
  //     fetchImages();
  //   }
  // }, [isLoggedIn]);

  // API実装までダミーでスライド用画像データの取得
  useEffect(() => {
    console.log('Fetching slider images');
    // Web APIの代わりにローカルの画像を使用
    setMotoSlideImages(getMotoSlideImages());
    setHumanSlideImages(getHumanSlideImages());
    setMoneySlideImages(getMoneySlideImages());
    setCarSlideImages(getCarSlideImages());
    setSnowboradSlideImages(getSnowboradSlideImages());
    setAnimeSlideImages(getAnimeSlideImages());
  }, [isLoggedIn]);  // isLoggedInの依存関係を追加

  const scrollToRef = (ref: React.RefObject<HTMLDivElement>) => {
    const offset = 80;
    const element = ref.current;
    if (element) {
      const topPosition = element.getBoundingClientRect().top + window.pageYOffset - offset;
      window.scrollTo({ top: topPosition, behavior: "smooth" });
    }
  };

  const handleLogin = () => {
    setIsLoggedIn(true);
    sessionStorage.setItem('isLoggedIn', 'true'); // セッションストレージにログイン状態を保存
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
    sessionStorage.setItem('isLoggedIn', 'false'); // セッションストレージからログイン状態を削除
  };

  return (
    <div className="flex flex-col min-h-screen bg-indigo-100">
      {isLoggedIn ? (
        <>
          <Header
            onChartViewClick={() => scrollToRef(dashboardRef)}
            onMarketInsightsClick={() => scrollToRef(marketInsightsRef)}
            onPortfolioClick={() => scrollToRef(portfolioRef)}
            onEducationClick={() => scrollToRef(educationRef)}
            onLogout={handleLogout}
            isLoggedIn={isLoggedIn}
          />

          {/* 右から左へのスライダーを配置 */}
          <ImageSlider images={moneySlideImages} direction="left-to-right" />

          <main className="flex-grow container mx-auto px-4 py-8">

            {/* 左から右へのスライダーを配置 */}
            <ImageSlider images={humanSlideImages} direction="right-to-left" /><br />

            <div ref={dashboardRef}>
              <ChartView />
            </div>

            {/* 右から左へのスライダーを配置 */}
            <ImageSlider images={motoSlideImages} direction="left-to-right" /><br />

            <div ref={marketInsightsRef}>
              <MarketInsights />
            </div>

            {/* 左から右へのスライダーを配置 */}
            <ImageSlider images={carSlideImages} direction="right-to-left" /><br />

            <div ref={portfolioRef}>
              <Portfolio />
            </div>
            {/* 右から左へのスライダーを配置 */}
            <ImageSlider images={animeSlideImages} direction="left-to-right" /><br />

            <div ref={educationRef}>
              <Education />
            </div>

            {/* 左から右へのスライダーを配置 */}
            <ImageSlider images={SnowboradSlideImages} direction="right-to-left" />

          </main>

          {/* 右から左へのスライダーを配置 */}
          <ImageSlider images={moneySlideImages} direction="right-to-left" />

          <Footer />
        </>
      ) : (
        <div
          className="flex items-center justify-center min-h-screen bg-cover bg-center"
          style={{ backgroundImage: `url(${backgroundImage})` }}
        >
          <div className="bg-white bg-opacity-75 p-8 rounded-lg shadow-lg">
            <LoginForm onLogin={handleLogin} />
          </div>
        </div>
      )}
    </div>
  );
};

export default App;
