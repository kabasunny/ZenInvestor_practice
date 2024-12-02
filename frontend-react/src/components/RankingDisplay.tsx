// frontend-react/src/components/RankingDisplay.tsx

import React, { useState } from "react";
import useRankingData from "../hooks/useRankingData";
import Modal from "./Modal";

const RankingDisplay: React.FC = () => {
  const { data, loading, error } = useRankingData();
  const [isModalOpen, setModalOpen] = useState(false);
  const [selectedStock, setSelectedStock] = useState<string | null>(null);

  const openModal = (ticker: string) => {
    setSelectedStock(ticker);
    setModalOpen(true);
  };

  const closeModal = () => {
    setSelectedStock(null);
    setModalOpen(false);
  };

  if (loading) return <p>読み込み中...</p>;
  if (error) return <p>{error}</p>;

  return (
    <div>
      <h3 className="text-lg font-bold mb-2">売買代金ランキング</h3>
      <table className="min-w-full bg-white border-collapse">
        <thead>
          <tr>
            <th
              className="border px-4 py-2 text-center"
              style={{ width: "10%" }}
            >
              順位
            </th>
            <th
              className="border px-4 py-2 text-center"
              style={{ width: "15%" }}
            >
              銘柄コード
            </th>
            <th
              className="border px-4 py-2 text-center"
              style={{ width: "35%", wordWrap: "break-word" }}
            >
              銘柄名
            </th>
            <th
              className="border px-4 py-2 text-right"
              style={{ width: "20%" }}
            >
              売買代金(5日平均)
            </th>
            <th
              className="border px-4 py-2 text-right"
              style={{ width: "20%" }}
            >
              終値
            </th>
          </tr>
        </thead>
        <tbody>
          {data.map((item) => (
            <tr key={item.ranking}>
              <td className="border px-4 py-2 text-center">{item.ranking}</td>
              <td
                className="border px-4 py-2 text-center cursor-pointer text-blue-500 underline"
                onClick={() => openModal(item.ticker)}
              >
                {item.ticker ?? "ティッカーなし"}
              </td>
              <td
                className="border px-4 py-2 text-center"
                style={{ wordWrap: "break-word" }}
              >
                {item.name ?? "名前がありません"}
              </td>
              <td className="border px-4 py-2 text-right">
                {item.avg_turnover?.toLocaleString() ?? "データなし"}
              </td>
              <td className="border px-4 py-2 text-right">
                {item.latest_close?.toLocaleString() ?? "データなし"}円
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      {/* モーダル表示 */}
      <Modal isOpen={isModalOpen} onClose={closeModal}>
        <h2 className="text-2xl mb-4">チャート</h2>
        <p>ここにチャートを表示</p>
        <p>選択された銘柄コード: {selectedStock}</p>
      </Modal>
    </div>
  );
};

export default RankingDisplay;