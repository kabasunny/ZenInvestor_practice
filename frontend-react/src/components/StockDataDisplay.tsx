import React from "react";
import { StockData } from "../types/stockTypes";

interface StockDataDisplayProps {
  date: string;
  stockData: StockData | null;
  loading: boolean;
  error: string | null;
}

const StockDataDisplay: React.FC<StockDataDisplayProps> = ({
  date,
  stockData,
  loading,
  error,
}) => {
  if (loading) return <p>読み込み中...</p>;
  if (error) return <p>{error}</p>;
  if (!stockData) return <p>データがありません。</p>;

  return (
    <div>
      <h3 className="text-lg font-semibold">銘柄データ</h3>
      <p>日付: - {date}</p>
      <p>始値: ${stockData.open.toFixed(2)}</p>
      <p>終値: ${stockData.close.toFixed(2)}</p>
      <p>高値: ${stockData.high.toFixed(2)}</p>
      <p>安値: ${stockData.low.toFixed(2)}</p>
      <p>出来高: {stockData.volume.toLocaleString()}</p>
    </div>
  );
};

export default StockDataDisplay;
