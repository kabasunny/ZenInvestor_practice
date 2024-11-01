import React from "react";
import { TrendingUp, DollarSign, AlertCircle } from "lucide-react";

const Dashboard: React.FC = () => {
  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ダッシュボード</h1>
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> 市場概況
          </h2>
          <p>
            市場は前向きな傾向を示しています。冷静さと集中力を保ちましょう。
          </p>
          {/* 市場データのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <DollarSign className="mr-2" /> ポートフォリオ概要
          </h2>
          <p>
            あなたのポートフォリオはバランスが取れています。即座のアクションは不要です。
          </p>
          {/* ユーザーのポートフォリオデータのAPIコール */}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <AlertCircle className="mr-2" /> Zenアドバイス
          </h2>
          <p>
            長期的な忍耐が持続可能な成長につながることを忘れないでください。
          </p>
          {/* パーソナライズされたアドバイスのAPIコール */}
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
