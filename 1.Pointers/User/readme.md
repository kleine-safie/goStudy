
```
go build -gcflags "-m -l" study
go build -gcflags "-m=2 -l" study
```
https://github.com/golang/go/blob/master/src/cmd/compile/internal/escape/escape.go


(1) pointers to stack objects cannot be stored in the heap
(2) pointers to a stack object cannot outlive that object (
   e.g., because the declaring function returned and destroyed the object's stack frame, or its space is reused across loop iterations for logically distinct variables).


よくあるパターン (例外もあり)：
- pointerを返す関数を作るときに、stack上の変数を返すと、stackが解放されてしまうので、heapに逃げる。
- channelでpointerを送ると、heapに逃げる。
- mapやsliceが関数のパラメータのときは、heapに逃げる。
- 関数のinterfaceパラメータのメソッドを呼ぶと、heapに逃げる(dynamic dispatch)
