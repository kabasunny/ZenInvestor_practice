// components/StockDataDisplay.tsx
import React from "react";
import { StockData } from "../types/stockTypes";

interface StockDataDisplayProps {
  stockName: string;
  ticker: string;  // ティッカーシンボルを追加
  date: string;
  stockData: StockData | null;
  loading: boolean;
  error: string | null;
}

const StockDataDisplay: React.FC<StockDataDisplayProps> = ({
  stockName,
  ticker,  // ティッカーシンボルを追加
  date,
  stockData,
  loading,
  error,
}) => {
  if (loading) return <p>読み込み中...</p>;
  if (error) return <p>{error}</p>;
  if (!stockData) return <p>データがありません。</p>;

  // ティッカーシンボルによって通貨記号を設定
  const currencySymbol = ticker.endsWith(".T") ? "¥" : "$";

  return (
    <div>
      <h3 className="text-lg font-semibold">銘柄データ</h3>
      <p>銘柄名 : {stockName}</p>
      <p>日付 : {date}</p>
      <p>始値 : {currencySymbol}{stockData.open.toFixed(2)}</p>
      <p>終値 : {currencySymbol}{stockData.close.toFixed(2)}</p>
      <p>高値 : {currencySymbol}{stockData.high.toFixed(2)}</p>
      <p>安値 : {currencySymbol}{stockData.low.toFixed(2)}</p>
      <p>出来高 : {stockData.volume.toLocaleString()}</p>
    </div>
  );
};

export default StockDataDisplay;
