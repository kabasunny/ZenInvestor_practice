import React from 'react';
import { Zap } from 'lucide-react';

interface HeaderProps {
  onDashboardClick: () => void;
  onMarketInsightsClick: () => void;
  onPortfolioClick: () => void;
  onEducationClick: () => void;
}

const Header: React.FC<HeaderProps> = ({
  onDashboardClick,
  onMarketInsightsClick,
  onPortfolioClick,
  onEducationClick
}) => {
  return (
    <header className="bg-indigo-600 text-white sticky top-0 z-10">
      <div className="container mx-auto px-4 py-4 flex justify-between items-center">
        <div className="flex items-center space-x-2 cursor-pointer" onClick={onDashboardClick}>
          <Zap size={24} />
          <span className="text-xl font-bold">ZenInvestor</span>
        </div>
        <nav>
          <ul className="flex space-x-4">
            <li><button onClick={onDashboardClick} className="hover:text-indigo-200">ダッシュボード</button></li>
            <li><button onClick={onMarketInsightsClick} className="hover:text-indigo-200">市場洞察</button></li>
            <li><button onClick={onPortfolioClick} className="hover:text-indigo-200">ポートフォリオ</button></li>
            <li><button onClick={onEducationClick} className="hover:text-indigo-200">教育</button></li>
          </ul>
        </nav>
      </div>
    </header>
  );
};

export default Header;