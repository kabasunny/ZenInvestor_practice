// types/rankingTypes.ts

export interface RankingData {
    ranking: number;
    ticker: string;
    date: string;
    avg_turnover: number;  // 修正後のフィールド名
    name: string;
    latest_close: number;
  }
  