// hooks/useStockChart.ts
import { useState, useEffect } from 'react';

const useStockChart = (ticker: string, period: string, indicators: any[]) => {
  const [chartData, setChartData] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true); // 初回レンダリング時にデータを取得するため true に設定
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchStockChartData = async () => {
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
        setChartData(data.chart_data);
      } catch (err: any) {
        setError(err.message || "データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchStockChartData();
  }, [ticker, period, indicators]);

  return { chartData, loading, error };
};

export default useStockChart;