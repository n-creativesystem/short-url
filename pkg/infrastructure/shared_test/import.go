package sharedtest

import "context"

var (
	Ctx context.Context = context.Background()
)

func Import() {
	/*
		NOTE
		var _ = Describe("xxx", func() {})
		が各ライブラリ側のテストで読み取られ処理されるので
		各ライブラリはこのメソッドを呼び出すと共通のテストが実施される
	*/
}
