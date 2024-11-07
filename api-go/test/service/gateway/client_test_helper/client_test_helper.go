package client_test_helper

import (
	"github.com/joho/godotenv"
)

func LoadTestEnv() {
	godotenv.Load("../../../../.env") //テストではパスを指定しないとうまく読み取らない
	// 上記でgrpcクライアントのポートを読み込む必要がある
}
