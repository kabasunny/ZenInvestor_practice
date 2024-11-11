// components/PortfolioOverview.tsx
import React from 'react';
import { DollarSign } from 'lucide-react';

const PortfolioOverview: React.FC = () => (
  <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
    <h2 className="text-xl font-semibold mb-4 flex items-center">
      <DollarSign className="mr-2" /> ポートフォリオ概要
    </h2>
    <p>
      あなたのポートフォリオはバランスが取れています。即座のアクションは不要です。
    </p>
  </div>
);

export default PortfolioOverview;