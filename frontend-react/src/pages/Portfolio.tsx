import React from 'react';
import { PieChart, TrendingUp, AlertTriangle } from 'lucide-react';

const Portfolio: React.FC = () => {
  return (
    <div className="space-y-6 mb-16">
      <h1 className="text-3xl font-bold">あなたのポートフォリオ</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="bg-white p-6 rounded-lg shadow-md col-span-2">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <PieChart className="mr-2" /> 資産配分
          </h2>
          <p>現在の資産配分がここに表示されます。</p>
          {/* ユーザーのポートフォリオデータのAPIコール */}
        </div>
        <div className="bg-white p-6 rounded-lg shadow-md">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> パフォーマンス
          </h2>
          <p>ポートフォリオのパフォーマンス指標がここに表示されます。</p>
          {/* ポートフォリオパフォーマンスデータのAPIコール */}
        </div>
        <div className="bg-white p-6 rounded-lg shadow-md col-span-full">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <AlertTriangle className="mr-2" /> Zenレコメンデーション
          </h2>
          <p>感情を考慮したパーソナライズされた投資推奨がここに表示されます。</p>
          {/* パーソナライズされた推奨のAPIコール */}
        </div>
      </div>
    </div>
  );
};

export default Portfolio;