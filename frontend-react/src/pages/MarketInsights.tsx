// frontend-react/src/pages/MarketInsights.tsx

import React from "react";
import { BarChart, Globe, Newspaper } from "lucide-react";
import RankingDisplay from "../components/RankingDisplay"; // 追加

const MarketInsights: React.FC = () => {
  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">市場動向</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg col-span-full">
          {" "}
          {/* ここで全域幅に設定 */}
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <BarChart className="mr-2" /> 売買代金ランキング (5日平均)
          </h2>
          <RankingDisplay /> {/* ここで呼び出し */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <Globe className="mr-2" /> グローバル経済指標
          </h2>
          <p>市場に影響を与える主要な経済指標がここに表示されます。</p>
          {/* 経済指標データのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <Newspaper className="mr-2" /> 最新の市場ニュース
          </h2>
          <p>
            最近のニュース記事と市場への潜在的な影響がここにリストされます。
          </p>
          {/* 市場ニュースデータのAPIコール */}
        </div>
      </div>
    </div>
  );
};

export default MarketInsights;
