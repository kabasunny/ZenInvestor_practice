import React, { useState, useEffect, useRef } from 'react';
import Header from './components/Header';
import Footer from './components/Footer';
import Dashboard from './pages/Dashboard';
import MarketInsights from './pages/MarketInsights';
import Portfolio from './pages/Portfolio';
import Education from './pages/Education';
import LoginForm from './components/LoginForm';

// assetsフォルダから画像をインポート
import img1 from './assets/ZenInvestor_8.jpg';
import img2 from './assets/ZenInvestor_22.jpg';
import img3 from './assets/ZenInvestor_26.jpg';
import img4 from './assets/ZenInvestor_28.jpg';
import img5 from './assets/ZenInvestor_33.jpg';
import img10 from './assets/ZenInvestor_70.jpg';
import img11 from './assets/ZenInvestor_73.jpg';
import img12 from './assets/ZenInvestor_84.jpg';

const images = [img1, img2, img3, img4, img5, img10, img11, img12];

const App: React.FC = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [backgroundImage, setBackgroundImage] = useState<string>('');

  const dashboardRef = useRef<HTMLDivElement>(null);
  const marketInsightsRef = useRef<HTMLDivElement>(null);
  const portfolioRef = useRef<HTMLDivElement>(null);
  const educationRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const randomImage = images[Math.floor(Math.random() * images.length)];
    console.log("Selected background image: ", randomImage); // デバッグ用メッセージ
    setBackgroundImage(randomImage);
  }, []);

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
  };

  const handleLogout = () => {
    setIsLoggedIn(false);
  };

  return (
    <div className="flex flex-col min-h-screen bg-indigo-100">
      {isLoggedIn ? (
        <>
          <Header
            onDashboardClick={() => scrollToRef(dashboardRef)}
            onMarketInsightsClick={() => scrollToRef(marketInsightsRef)}
            onPortfolioClick={() => scrollToRef(portfolioRef)}
            onEducationClick={() => scrollToRef(educationRef)}
            onLogout={handleLogout}
            isLoggedIn={isLoggedIn}
          />
          <main className="flex-grow container mx-auto px-4 py-8">
            <div ref={dashboardRef}>
              <Dashboard />
            </div>
            <div ref={marketInsightsRef}>
              <MarketInsights />
            </div>
            <div ref={portfolioRef}>
              <Portfolio />
            </div>
            <div ref={educationRef}>
              <Education />
            </div>
          </main>
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
