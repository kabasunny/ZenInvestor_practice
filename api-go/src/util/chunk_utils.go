// api-go\src\util\chunk_utils.go
package util

import (
	gsdwd "api-go/src/service/ms_gateway/get_stocks_datalist_with_dates"
	// その他必要なインポート
)

// ChunkSymbols は指定されたサイズでシンボルリストをチャンクに分割
func ChunkSymbols(symbols []string, chunkSize int) [][]string {
	var chunks [][]string                          // チャンクのスライスを初期化
	for i := 0; i < len(symbols); i += chunkSize { // チャンクサイズごとにループ
		end := i + chunkSize    // チャンクの終了インデックスを計算
		if end > len(symbols) { // 最後のチャンクがリストの範囲を超えないように調整
			end = len(symbols)
		}
		chunks = append(chunks, symbols[i:end]) // チャンクをリストに追加
	}
	return chunks // チャンクに分割されたシンボルリストを返す
}

// ChunkData は指定されたサイズでデータリストをチャンクに分割
func ChunkData(data []*gsdwd.StockPrice, chunkSize int) [][]*gsdwd.StockPrice {
	var chunks [][]*gsdwd.StockPrice            // チャンクのスライスを初期化
	for i := 0; i < len(data); i += chunkSize { // チャンクサイズごとにループ
		end := i + chunkSize // チャンクの終了インデックスを計算
		if end > len(data) { // 最後のチャンクがリストの範囲を超えないように調整
			end = len(data)
		}
		chunks = append(chunks, data[i:end]) // チャンクをリストに追加
	}
	return chunks // チャンクに分割されたデータリストを返す
}
