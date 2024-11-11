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
  const [stockDataWithDate, setStockDataWithDate] =
  useState<StockDataWithDate | null>(null);
  const [chartData, setChartData] = useState<string | null>(null); // Base64エンコードされた画像データを格納

  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);
  const [ticker, setTicker] = useState<string>("AAPL");
  const [period, setPeriod] = useState<string>("1d");
  // indicatorsの状態を追加
  const [indicators, setIndicators] = useState<any[]>([
    { type: 'SMA', params: { window_size: '20' } }
  ]);

  const fetchStockData = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch(
        `http://localhost:8086/getStockData?ticker=${ticker}&period=${period}`
      );
      const data = await response.json();
      const date = Object.keys(data.stock_data)[0];
      const stockData = data.stock_data[date];
      setStockDataWithDate({ date, stockData });
      console.log("chartData:", data.stock_data);
      setLoading(false);
    } catch (err) {
      setError("データの取得に失敗しました");
      setLoading(false);
    }
  };

  const fetchStocChartkData = async () => {
    setLoading(true);
    setError(null);
    try {
      const response = await fetch("http://localhost:8086/getStockChart", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          ticker: ticker,
          period: period,
          indicators: indicators,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || "サーバーエラー");
      }
      const data = await response.json();
      // console.log("chartData:", data.chart_data);
      setChartData(data.chart_data);
    } catch (err: any) {
      console.error("Error fetching stock data:", err);
      setError(err.message || "データの取得に失敗しました");

    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    fetchStocChartkData();
    fetchStockData();
  };

  return (
    <div className="space-y-6 mb-32">
      <h1 className="text-3xl font-bold">ダッシュボード</h1>
      <div className="grid grid-cols-1 gap-4"> {/* 親要素を1列グリッドに設定 */}
        <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
          <h2 className="text-xl font-semibold mb-4 flex items-center">
            <TrendingUp className="mr-2" /> 指定銘柄チャート表示
          </h2>
          <StockForm
            ticker={ticker}
            period={period}
            indicators={indicators}
            setTicker={setTicker}
            setPeriod={setPeriod}
            setIndicators={setIndicators}
            onSubmit={handleSubmit}
          />
          {loading && <p>Loading...</p>}
          {error && <p className="text-red-500">{error}</p>}
          {stockDataWithDate ? (
            <StockDataDisplay
              date={stockDataWithDate.date}
              stockData={stockDataWithDate.stockData}
              loading={loading}
              error={error}
            />
          ) : (
            <StockDataDisplay
              date=""
              stockData={null}
              loading={loading}
              error={error}
            />
          )}
          {chartData && (
            <img src={`data:image/png;base64,${chartData}`} alt="stock chart" />
          )}
        </div>

        {/* 二つの要素を横並びに表示 */}
        <div className="grid grid-cols-2 gap-4"> {/* ここでグリッドを2列に設定 */}
          <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
            <h2 className="text-xl font-semibold mb-4 flex items-center">
              <DollarSign className="mr-2" /> ポートフォリオ概要
            </h2>
            <p>
              あなたのポートフォリオはバランスが取れています。即座のアクションは不要です。
            </p>
          </div>
          <div className="bg-teal-50 p-6 rounded-lg shadow-lg">
            <h2 className="text-xl font-semibold mb-4 flex items-center">
              <AlertCircle className="mr-2" /> Zenアドバイス
            </h2>
            <p>
              長期的な忍耐が持続可能な成長につながることを忘れないでください。
            </p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Dashboard;
