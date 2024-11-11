// hooks/useStockData.ts
import { useState, useEffect } from 'react';
import { StockData, StockDataWithDate } from '../types/stockTypes';

const useStockData = (ticker: string, period: string) => {
  const [stockDataWithDate, setStockDataWithDate] = useState<StockDataWithDate | null>(null);
  const [loading, setLoading] = useState<boolean>(true); // 初回レンダリング時にデータを取得するため true に設定
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
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
      } catch (err) {
        setError("データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchStockData();
  }, [ticker, period]);

  return { stockDataWithDate, loading, error };
};

export default useStockData;