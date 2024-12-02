// frontend-react/src/hooks/useRankingData.ts

import { useState, useEffect } from 'react';
import axios from 'axios';
import { RankingData } from '../types/rankingTypes';  // インポート

const useRankingData = () => {
  const [data, setData] = useState<RankingData[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await axios.get<RankingData[]>('http://localhost:8086/RankingByRange?startRank=1&endRank=100');
        setData(response.data);
        setLoading(false);
      } catch (err) {
        setError('データの取得に失敗しました');
        setLoading(false);
      }
    };
    fetchData();
  }, []);

  return { data, loading, error };
};

export default useRankingData;
