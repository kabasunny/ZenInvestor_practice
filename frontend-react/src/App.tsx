// App.tsx (修正後)
import React, { useRef, useEffect } from "react";
import Header from "./components/Header";
import Footer from "./components/Footer";
import Dashboard from "./pages/Dashboard";
import MarketInsights from "./pages/MarketInsights";
import Portfolio from "./pages/Portfolio";
import Education from "./pages/Education";
import LoginForm from "./components/LoginForm";
import { useAuth } from "./context/AuthContext"; // useAuth をインポート

const App: React.FC = () => {
  const { isLoggedIn, login, logout } = useAuth(); // useAuth()を使ってログイン状態を取得 index.tsx内でAppをAuthProviderでラップしているため使用可能
  const dashboardRef = useRef<HTMLDivElement>(null);
  const marketInsightsRef = useRef<HTMLDivElement>(null);
  const portfolioRef = useRef<HTMLDivElement>(null);
  const educationRef = useRef<HTMLDivElement>(null);

  const scrollToRef = (ref: React.RefObject<HTMLDivElement>) => {
    ref.current?.scrollIntoView({ behavior: "smooth" });
  };

  useEffect(() => {
    console.log("isLoggedIn が更新されました:", isLoggedIn);
  }, [isLoggedIn]);

  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      {isLoggedIn ? ( // ログイン済みの場合
        <>
          <Header
            onDashboardClick={() => scrollToRef(dashboardRef)}
            onMarketInsightsClick={() => scrollToRef(marketInsightsRef)}
            onPortfolioClick={() => scrollToRef(portfolioRef)}
            onEducationClick={() => scrollToRef(educationRef)}
            onLogout={logout}
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
        // 未ログインの場合
        <div className="flex items-center justify-center min-h-screen">
          <LoginForm onLogin={login} /> {/* onLogin プロパティを渡す */}
        </div>
      )}
    </div>
  );
};

export default App;
