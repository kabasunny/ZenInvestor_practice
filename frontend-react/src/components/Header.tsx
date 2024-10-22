import React from "react";
import { Zap } from "lucide-react";
import { useAuth } from "../context/AuthContext"; // useAuthをインポート

interface HeaderProps {
  onDashboardClick: () => void;
  onMarketInsightsClick: () => void;
  onPortfolioClick: () => void;
  onEducationClick: () => void;
  onLogout: () => void;
  isLoggedIn: boolean;
}

const Header: React.FC<HeaderProps> = ({
  onDashboardClick,
  onMarketInsightsClick,
  onPortfolioClick,
  onEducationClick,
  onLogout,
  isLoggedIn,
}) => {
  return (
    <header className="bg-indigo-500 text-white sticky top-0 z-10">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div
          className="flex items-center space-x-2 cursor-pointer"
          onClick={onDashboardClick}
        >
          <Zap size={24} />
          <span className="text-xl font-bold">ZenInvestor</span>
        </div>
        <nav>
          <ul className="flex space-x-4">
            {isLoggedIn && (
              <>
<<<<<<< HEAD
                <li><button onClick={onDashboardClick} className="hover:text-indigo-200">Dashboard</button></li>
                <li><button onClick={onMarketInsightsClick} className="hover:text-indigo-200">MarketInsights</button></li>
                <li><button onClick={onPortfolioClick} className="hover:text-indigo-200">Portfolio</button></li>
                <li><button onClick={onEducationClick} className="hover:text-indigo-200">Education</button></li>
=======
                <li>
                  <button
                    onClick={onDashboardClick}
                    className="hover:text-indigo-200"
                  >
                    ダッシュボード
                  </button>
                </li>
                <li>
                  <button
                    onClick={onMarketInsightsClick}
                    className="hover:text-indigo-200"
                  >
                    市場洞察
                  </button>
                </li>
                <li>
                  <button
                    onClick={onPortfolioClick}
                    className="hover:text-indigo-200"
                  >
                    ポートフォリオ
                  </button>
                </li>
                <li>
                  <button
                    onClick={onEducationClick}
                    className="hover:text-indigo-200"
                  >
                    教育
                  </button>
                </li>
>>>>>>> c22ddb56ec4640cf7d6a03e8fc452cd83b596c91
              </>
            )}
          </ul>
        </nav>
        {isLoggedIn && (
<<<<<<< HEAD
          <button onClick={onLogoutClick} className="ml-4 bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded">
            Logout
=======
          <button
            onClick={onLogout} // logout関数を使用
            className="ml-4 bg-red-600 hover:bg-red-800 text-white font-bold py-2 px-4 rounded"
          >
            ログアウト
>>>>>>> c22ddb56ec4640cf7d6a03e8fc452cd83b596c91
          </button>
        )}
      </div>
    </header>
  );
};

export default Header;
