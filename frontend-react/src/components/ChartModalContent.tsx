// frontend-react/src/components/ChartModalContent.tsx

import React from "react";
import useModalStockChart from "../hooks/useModalStockChart";

interface ChartModalContentProps {
  rank: number | null;
  ticker: string | null;
  stockName: string | null;
}

const ChartModalContent: React.FC<ChartModalContentProps> = ({
  rank,
  ticker,
  stockName,
}) => {
  if (!ticker) return null;

  const { chartData, loading, error } = useModalStockChart(ticker);

  if (loading) return <p>チャートを読み込み中...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div style={{ width: "100%", height: "400px" }}>
      <h3 className="text-xl mb-2">
        {rank} 位 <br />
        {stockName}
      </h3>
      {chartData ? (
        <img
          src={`data:image/png;base64,${chartData}`}
          alt="チャート"
          style={{ width: "100%", height: "100%" }}
        />
      ) : (
        <p>データがありません</p>
      )}
    </div>
  );
};

export default ChartModalContent;
