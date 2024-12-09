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
    <div style={{ width: "100%", height: "600px" }}>
      <h3 className="text-xl mb-2">
        ★直近5日の平均売買代金★
        <br /> {rank} 位 : {stockName} ( 直近1ヶ月の価格推移 )
      </h3>
      {chartData ? (
        <img
          src={`data:image/png;base64,${chartData}`}
          alt="チャート"
          style={{ width: "100%", height: "90%" }}
        />
      ) : (
        <p>データがありません</p>
      )}
    </div>
  );
};

export default ChartModalContent;
