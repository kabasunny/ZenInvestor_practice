// hooks/useStockData.ts
import { useState, useEffect } from 'react';
import { StockData, StockDataWithDate } from '../types/stockTypes';

const useStockData = (ticker: string, period: string, updateFlag: boolean) => {
  const [stockDataWithDate, setStockDataWithDate] = useState<StockDataWithDate | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
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

        // 日付の配列を取得し、降順にソートして最新の日付を取得
        const dates = Object.keys(data.stock_data).sort((a, b) => (a > b ? 1 : -1));
        const latestDate = dates[dates.length - 1];  // 最新の日付
        const stockData = data.stock_data[latestDate];  // 最新の日付のデータ

        setStockDataWithDate({ date: latestDate, stockData });
      } catch (err) {
        setError("データの取得に失敗しました");
      } finally {
        setLoading(false);
      }
    };

    fetchStockData();
  }, [ticker, period, updateFlag]);

  return { stockDataWithDate, loading, error };
};

export default useStockData;
