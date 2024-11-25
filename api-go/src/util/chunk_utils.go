// api-go\src\util\chunk_utils.go
package util

// ChunkSymbols は指定されたサイズでシンボルリストをチャンクに分割します
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
