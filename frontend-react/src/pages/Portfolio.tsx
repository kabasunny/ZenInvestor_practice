import React from 'react';
import { PieChart, TrendingUp, AlertTriangle } from 'lucide-react';
import LossCutSimulator from '../components/LossCutSimulator';

const Portfolio: React.FC = () => {
  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ポートフォリオ</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg col-span-2">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <PieChart className="mr-2" /> 資産配分
          </h2>
          <p>現在の資産配分がここに表示されます。</p>
          {/* ユーザーのポートフォリオデータのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> パフォーマンス
          </h2>
          <p>ポートフォリオのパフォーマンス指標がここに表示されます。</p>
          {/* ポートフォリオパフォーマンスデータのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg col-span-full">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <AlertTriangle className="mr-2" /> ロスカットシミュレーター
          </h2>
          <LossCutSimulator />
        </div>
      </div>
    </div>
  );
};

export default Portfolio;
