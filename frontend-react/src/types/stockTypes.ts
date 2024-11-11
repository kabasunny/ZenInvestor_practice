// types/stockTypes.ts
export interface StockData {
    open: number;
    close: number;
    high: number;
    low: number;
    volume: number;
}

export interface StockDataResponse {
    stockData: {
        stock_data: {
            [date: string]: StockData;
        };
    };
}

export interface StockDataWithDate {
    date: string;
    stockData: StockData;
  }