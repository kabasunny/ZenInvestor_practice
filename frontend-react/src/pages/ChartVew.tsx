// pages/Dashboard.tsx
import React, { useState, useEffect } from "react";
import ChartForm from "../components/ChartForm";
import StockDataDisplay from "../components/StockDataDisplay";
import PortfolioOverview from "../components/PortfolioOverview";
import ZenAdvice from "../components/ZenAdvice";
import { TrendingUp } from "lucide-react";
import useStockData from "../hooks/useStockData";
import useStockChart from "../hooks/useStockChart";

const Dashboard: React.FC = () => {
  const [ticker, setTicker] = useState<string>("6752.T");
  const [period, setPeriod] = useState<string>("1y");
  const [indicators, setIndicators] = useState<any[]>([
    { type: 'SMA', params: { window_size: '20' } }
  ]);

  const [includeVolume, setIncludeVolume] = useState<boolean>(false); // 出来高要否の状態を管理

  const [updateFlag, setUpdateFlag] = useState(false);
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    console.log('handleSubmit function called');
    // フラグを更新して、カスタムフックを再実行させる
    setUpdateFlag(prevFlag => !prevFlag);
  };

  useEffect(() => {
    // includeVolumeが変更された時に自動的に更新
    setUpdateFlag(prevFlag => !prevFlag);
  }, [includeVolume]);

  const {
    stockDataWithDate,
    stockName,
    loading: stockLoading,
    error: stockError
  } = useStockData(ticker, period, updateFlag);

  const {
    chartData,
    loading: chartLoading,
    error: chartError
  } = useStockChart(ticker, period, indicators, updateFlag, includeVolume);

  const loading = stockLoading || chartLoading;
  const error = stockError || chartError;  

  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ダッシュボード</h1>
      <div className="grid grid-cols-1 gap-4">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> カスタムチャート表示
          </h2>

          <div className="flex space-x-4">
            {/* StockForm コンポーネント */}
            <ChartForm
              ticker={ticker}
              period={period}
              indicators={indicators}
              setTicker={setTicker}
              setPeriod={setPeriod}
              setIndicators={setIndicators}
              includeVolume={includeVolume} // includeVolumeを追加
              setIncludeVolume={setIncludeVolume} // setIncludeVolumeを追加
              onSubmit={handleSubmit}
            />

            {/* StockDataDisplay コンポーネント */}
            <div className="flex-grow">
              {loading && <p>Loading...</p>}
              {error && <p className="text-red-500">{error}</p>}
              {stockDataWithDate && (
                <StockDataDisplay
                  stockName={stockName || "銘柄名がありません"} // デフォルト値を設定銘柄名を渡す
                  ticker={ticker}
                  date={stockDataWithDate.date}
                  stockData={stockDataWithDate.stockData}
                  loading={loading}
                  error={error}
                />
              )}
            </div>
          </div>

          {chartData && (
            <img
              src={`data:image/png;base64,${chartData}`}
              alt="stock chart"
              className="mt-4"
            />
          )}
        </div>

        <div className="grid grid-cols-2 gap-4">
          <PortfolioOverview />
          <ZenAdvice />
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
