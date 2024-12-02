import { useState, useEffect } from 'react';

const useModalStockChart = (ticker: string, period: string = '1mo') => {
  const [chartData, setChartData] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchStockChartData = async () => {
      setLoading(true);
      setError(null);
      try {
        // ティッカーシンボルの末尾に".T"を付加
        const tickerWithSuffix = ticker.endsWith('.T') ? ticker : `${ticker}.T`;
        console.log(`Fetching chart data for ${tickerWithSuffix} with period ${period}`);
        
        const response = await fetch("http://localhost:8086/getStockChart", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            ticker: tickerWithSuffix,
            period: period,
            indicators: [],
            includeVolume: false,
          }),
        });

        if (!response.ok) {
          const errorData = await response.json();
          throw new Error(errorData.error || "サーバーエラー");
        }
        const data = await response.json();
        console.log("Fetched data:", data);
        setChartData(data.chart_data);
      } catch (err: any) {
        setError(err.message || "データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchStockChartData();
  }, [ticker, period]);

  return { chartData, loading, error };
};

export default useModalStockChart;
