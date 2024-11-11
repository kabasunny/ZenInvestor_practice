// pages/Dashboard.tsx
import React, { useState } from "react";
import StockForm from "../components/StockForm";
import StockDataDisplay from "../components/StockDataDisplay";
import PortfolioOverview from "../components/PortfolioOverview";
import ZenAdvice from "../components/ZenAdvice";
import { TrendingUp } from "lucide-react";
import useStockData from "../hooks/useStockData";
import useStockChart from "../hooks/useStockChart";

const Dashboard: React.FC = () => {
  const [ticker, setTicker] = useState<string>("AAPL");
  const [period, setPeriod] = useState<string>("1y");
  const [indicators, setIndicators] = useState<any[]>([
    { type: 'SMA', params: { window_size: '20' } }
  ]);

  // カスタムフックを使用してデータを取得
  const {
    stockDataWithDate,
    loading: stockLoading,
    error: stockError
  } = useStockData(ticker, period);

  const {
    chartData,
    loading: chartLoading,
    error: chartError
  } = useStockChart(ticker, period, indicators);

  // ローディングとエラーステートを統合
  const loading = stockLoading || chartLoading;
  const error = stockError || chartError;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // フォームの送信時に状態が更新されるため、カスタムフックが再呼び出しされデータが再取得されます
  };

  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ダッシュボード</h1>
      <div className="grid grid-cols-1 gap-4">
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> 指定銘柄チャート表示
          </h2>

          <div className="flex space-x-4">
            {/* StockForm コンポーネント */}
            <StockForm
              ticker={ticker}
              period={period}
              indicators={indicators}
              setTicker={setTicker}
              setPeriod={setPeriod}
              setIndicators={setIndicators}
              onSubmit={handleSubmit}
            />

            {/* StockDataDisplay コンポーネント */}
            <div className="flex-1">
              {loading && <p>Loading...</p>}
              {error && <p className="text-red-500">{error}</p>}
              {stockDataWithDate && (
                <StockDataDisplay
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
