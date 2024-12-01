import React from "react";
import { Zap } from "lucide-react";
import { useAuth } from "../context/AuthContext"; // useAuthをインポート

interface HeaderProps {
  onChartViewClick: () => void;
  onMarketInsightsClick: () => void;
  onPortfolioClick: () => void;
  onEducationClick: () => void;
  onLogout: () => void;
  isLoggedIn: boolean;
}

const Header: React.FC<HeaderProps> = ({
  onChartViewClick,
  onMarketInsightsClick,
  onPortfolioClick,
  onEducationClick,
  onLogout,
  isLoggedIn,
}) => {
  const scrollToTop = () => {
    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });
  };

  return (
    <header className="bg-indigo-500 text-white sticky top-0 z-10">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div
          className="flex items-center space-x-2 cursor-pointer"
          onClick={scrollToTop} // ここを修正
        >
          <Zap size={24} />
          <span className="text-xl font-bold">ZenInvestor</span>
        </div>
        <nav>
          <ul className="flex space-x-4">
            {isLoggedIn && (
              <>
                <li><button onClick={onChartViewClick} className="hover:text-indigo-200">Dashboard</button></li>
                <li><button onClick={onMarketInsightsClick} className="hover:text-indigo-200">MarketInsights</button></li>
                <li><button onClick={onPortfolioClick} className="hover:text-indigo-200">Portfolio</button></li>
                <li><button onClick={onEducationClick} className="hover:text-indigo-200">Education</button></li>
              </>
            )}
          </ul>
        </nav>
        {isLoggedIn && (
          <button onClick={onLogout} className="ml-4 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
            Logout
          </button>
        )}
      </div>
    </header>
  );
};

export default Header;
