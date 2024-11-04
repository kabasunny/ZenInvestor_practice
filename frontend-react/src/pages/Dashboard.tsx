import React, { useState } from "react";
import StockForm from "../components/StockForm";
import StockDataDisplay from "../components/StockDataDisplay";
import { TrendingUp, DollarSign, AlertCircle } from "lucide-react";
import { StockData } from "../types/stockTypes";

interface StockDataWithDate {
  date: string;
  stockData: StockData;
}

const Dashboard: React.FC = () => {
  const [stockDataWithDate, setStockDataWithDate] = useState<StockDataWithDate | null>(null);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [ticker, setTicker] = useState<string>("AAPL");
  const [period, setPeriod] = useState<string>("1d");

  const fetchStockData = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(`http://localhost:8086/getStockData?ticker=${ticker}&period=${period}`);
      const data = await response.json();
      const date = Object.keys(data.stockData.stock_data)[0];
      const stockData = data.stockData.stock_data[date];
      setStockDataWithDate({ date, stockData });
      setLoading(false);
    } catch (err) {
      setError("データの取得に失敗しました");
      setLoading(false);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    fetchStockData();
  };

  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ダッシュボード</h1>
      <StockForm ticker={ticker} period={period} setTicker={setTicker} setPeriod={setPeriod} onSubmit={handleSubmit} />
      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> 銘柄指定
          </h2>
          {stockDataWithDate ? (
            <StockDataDisplay
              date={stockDataWithDate.date}
              stockData={stockDataWithDate.stockData}
              loading={loading}
              error={error}
            />
          ) : (
            <StockDataDisplay date="" stockData={null} loading={loading} error={error} />
          )}
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <DollarSign className="mr-2" /> ポートフォリオ概要
          </h2>
          <p>あなたのポートフォリオはバランスが取れています。即座のアクションは不要です。</p>
        </div>
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <AlertCircle className="mr-2" /> Zenアドバイス
          </h2>
          <p>長期的な忍耐が持続可能な成長につながることを忘れないでください。</p>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
