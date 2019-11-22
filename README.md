# LKY's Golang Study

<img align="right" src="go_learn.png">

### ToDo

- 了解 new 與 make 有哪些不同？
- 了解 func 與 method 有哪些不同？在 Go 裡面好像分得很清楚。
- 了解關鍵字 fallthrough

### Next

- or-channel
- Channel
  - sync.Mutx（應該可以跳過，再看看）
  - close, range
  - Channel 消耗到空的時候取值不如預期中得到 false，這個現象還要花時間搞清楚
    - chan return ok 應該是代表 close 與否，待驗證。
- 內建 Testing 模式
- ~~了解 https method 必需要 goroutine 的一些初步簡單用法，希望能順便加入未來一週的練習中。~~

### 2019.11.22(Fri)

- 跟 TG 群討論測試前一天  LeetCode in Concurrency 更好的解法
  - 「Print Zero-Even-Odd」：能不能用 unbuffered chan 解？可能是無解。
    - 理由：若使用 unbuffered chan，收與送的 goroutine 兩邊都用 select 會導致兩邊都不走 case，只走 default，結果等於 deadlock。
  - 「Print in order」
    - 群友提供更好解法
      1. 所有 goroutine 都是依照一個鍊執行，所以不用設置 syncWatiGroup，多用一個 unbuffered chan recevier 卡住 main goroutine，等待最後一個完成就好。
      2. interface{} 本質上是一個 struct + 2個pointer，所以對於「只在乎收發、不在乎內容」的 channel 傳遞來說，直接使用 chan struct{} 可能會比官方慣用的 interface{} 更輕量。

### 2019.11.21(Thu)

- 嘗試做兩題 LeetCode in Concurrency
  - 要求是任意順序呼叫多個函式，但是要以相同順序執行。
  - 要求三個 goroutine 以 1 2 1 3 ...(repeat) 的順序輸出到特定次數。
    - 這一題有循環交接棒，但是在程式結尾交棒(send to channel)的時候，消費者 goroutine 已經結束造成 deadlock 的問題，最後是把 chan 改成 len=1 的 buffered chan 就解決了 XD 因為這樣就不用面對消費者消失的問題。但這樣也讓我懷疑 unbuffered chan 是否有不可取代性？是不是都可以用 len=1 的 buffered chan 取代？

### 2019.11.20(Wed)

- 感覺目前有點碰壁，先從頭閱讀[官方文件 Goroutine 章節](https://golang.org/doc/effective_go.html#goroutines)，但還是有些地方...有看沒有懂。
  - [Channels of channels](https://golang.org/doc/effective_go.html#chan_of_chan)這部分很難，看不懂在說什麼，程式範例也不完整，有沒宣告的變數不知道怎麼跑。
- 自己評估目前缺乏遇到 Concurrency 問題的經驗，所以看到進階的 goroutine 才會看不懂在講什麼，目前可能只有 pipeline 我比較看得懂。
- LeetCode 上面有 Concurrency 的挑戰，我覺得這是一個培養問題經驗很好的切入點，但點進去看發現還沒有 Go 語言的。

### 2019.11.19(Tue)

- 嘗試 goroutine leak
- Channel 關閉(close)的原則
  - 不要在消費者端關閉
  - 不要在尚存多並行生產者的時候關閉
  - 應該只在「唯一且最後」的生產者 goroutine 中關閉
  - [參考資料：如何优雅的关闭Go Channel【译】](https://zhuanlan.zhihu.com/p/32529039)
- 寫一個計算圓周率的程式，比較單一與多 goroutine 計算所需時間，發現 1 VS 4 最多也只能把時間縮短到 48%

### 2019.11.18(Mon)

- 了解 channel 用法中的「特定約束 VS 詞法約束」
- for-select 優先順序

### 2019.11.16(Sat)

- 抓取空的 channel 並不一定會 deadlock，要 「channel 沒有 close」且「佔據 main」 才會。放在 goroutine 裡面怎麼卡都可以，main 一結束就會自然釋放。
- 了解 Select Case 的行為細節：
  - Closed Channel > default > Open Channel，如果多個 Closed Channel 就隨機分配。
- 嘗試 goroutine 的 closure trap、race codition、佔用記憶體測量。
- 了解 Go 語言從傳統 lock-unlock 的 critical section 控制，過渡到 goroutine-channel 的模型演變，中間有一個「M:N 調度器」的實作。
  - thread 比 goroutine 更輕量嗎？

### 2019.11.15(Fri)

- Q1:為什麼放了這麼多東西進去 chan 變數，還是 len=0？這樣我測量不到未被消費的元素數量啊。
- Q2:是不是使用到 Channel 型別參數的函式，就會被當成 goroutine？我的函式呼叫了 chan 型別，不加上 go 修飾詞就會跑出「all goroutines are asleep - deadlock!」這個錯誤。
- A1&A2:可能是沒有搞懂 buffered 跟 unbuffered channel 的差異：
  - [Golang buffered/unbuffer channel and pipeline](https://www.evanlin.com/til-buffered-channel/)
  - [用五分鐘了解什麼是 unbuffered vs buffered channel](https://blog.wu-boy.com/2019/04/understand-unbuffered-vs-buffered-channel-in-five-minutes/)
  - 弄懂了 Channel 在 buffered 跟 unbuffered 的差異
    - Unbuffered Channel 只能存放一個元素所以 size = 0（為什麼不能是 1 ?）
      - TG 群問答：
        - 孝玠 蔡：「unbuffered channel 不能存放噢」
        - 我：「如果不能存放的話，那在「放」跟「取」中間的時間差，這一個元素去哪裡了？」
        - 孝玠 蔡：「沒有放進去噢，取跟放是同時發生的，那個時間差，其實放的程式碼是卡住的」
    - Buffered Channel 才有比較大的彈性，不會被一吞接一吐的順序限制卡死。
    - 另外發現1：Buffered 還是有可能會發生 deadlock 的情況：channel 放不下放不完 & 消費者尚未啟動。
    - 另外發現2：沒有平行化的情況下，對 channel 都不用取，光是放就會 deadlock。
- Channel、Select、defer 混搭使用，還是會有非預期現象，需要繼續搞懂。

### 2019.11.14(Thu)

- 練習 Interface 轉換與組合
  - 轉換：Go 對 interface 採用「鴨子型別」：interface cast 只看方法、不看名稱，只要「新參考方法集合 ∈ 舊參考方法集合」就可以轉換。
  - 組合：介面的內嵌介面，會與介面本身產生同等約束效果，因此要連帶將內嵌介面一併實作才行。
- 練習 Goroutine
  - sync.WaitGroup
  - Channel，執行結果有點不如預期，應該還有什麼沒搞懂。

### 2019.11.13(Wed)

- 練習 Interface 各種花式用法
  - 這一篇是我覺得把 Interface in golang 解釋得比較清楚的 https://yami.io/golang-interface/ 
  - Q1: struct 實現 interface 的型別只能是 func(t T)Foo() 而不是 func(t *T)Foo() 的話，那要怎麼在實現 interface 的方法裡面改變 struct 的自身狀態？
    - 相關討論：
      - [再议go语言的value receiver和pointer receiver](https://www.jianshu.com/p/d1a9bbd0ae36)
      - [Golang method with pointer receiver [duplicate]
  ](https://stackoverflow.com/questions/33936081/golang-method-with-pointer-receiver)
  - A1: 其實 interface method 可以有 pointer receiver，但是呼叫的 struct 本身也要是 pointer type，不可以是 value type（如果 struct 已經宣告成 value type，那就對 instance 取址再用）。這似乎是 Go 語言設計的防呆防錯安全機制，看「[再议go语言的value receiver和pointer receiver](https://www.jianshu.com/p/d1a9bbd0ae36)」這篇發現的。
  - 心得：
    - Go 語言中的 interface 已經取代其他 OO 語言中所有繼承的功能。
    - type interface{} 相當於 C# 或 Java 中的 type Object，也就是所有型別的頂層型別。

### 2019.11.12(Tue)

- 練習 Struct 基本用法

### 2019.11.11(Mon)

- 練習「godoc 註解即文件」
- 練習 go package 管理，順便看如何解決 LeetCode 的 code 放在同一個資料夾底下後出現「found packages ...」的錯誤。
- 心得：
  - 要放在 $GOPATH 或 $GOROOT 底下才有用。要不然放在專案內的 src 底下只是能跟著 main 執行，什麼 go install, go doc 都沒用。
- 練習 map 基本用法
- 練習 defer, panic, recover 基本用法，總之就是 golang 的例外處理體系。
- 刷過一題 LeetCode

---

### 2019.11.9(Fri)

- 預計到台北參加例行會議

### 2019.11.8(Thu)

- ~~了解「godoc 註解即文件」~~ 來不及看
- 了解 LeetCode 使用方法
  - 初次體驗 LeetCode，的確對熟悉新語言有幫助
  - 看別人的解答會看到還不熟悉的語法，有助於拓展語法認識
  - 整天就只刷了兩題。
    - 順便了解 struct 用法
    - 有些題目懂 map 會有比較漂亮的解法
  - Runtime 的部分，同樣程式碼也會經常變來變去，不知道平台有什麼問題？
  - Submission Detail 那邊，不知道怎麼看其他語言的提交結果統計？是要付費版嗎？

### 2019.11.7(Wed)

- 完成 Anonymous Function 練習
- Closure
  - 重現「閉包陷阱」
  - 對閉包(Closure)定義仍不是相當清楚，只是大概知道現象類似 C++ static local variable。
- 了解外部 import 包使用方法
- LINQ in Golang
  - LINQ 是在 C# 裡相當好用的東西，可以用簡潔語法避免掉很多又臭又長的迴圈、判斷式，不過 Golang 原生沒有提供類似的好用工具。Github 上有這樣的 Libary，是 Google 員工開發的第三方套件，要來好好的玩一下，順便看作者是如何實作出 Golang 沒有的泛型。
  - 目前的實現還是不夠完美，可能有點疊床架屋，沒有各種可以直接吐出來的 interface、ToSlice()要傳一個 slice 進去而不是直接吐出來。這裡不知道還沒有讓我提出 Pull Request 的機會？
  - 使用過程發現自己對 interface in Golang 還不了解所以用起來卡，所以現在用算是有點過早。

### 2019.11.6(Tue)

- 今天了解 slice 只是單純的 reference，底層依然是 array，所以可以多個 slice 映射到同一個 array 上的任意區段。但不同的 array 就一定是不同的 instance。
- 了解 定義型態（defined type）與 底層型態（underlying type）
- function: 函式參考、不定長度參數、自定義函式型與別名、callback

### 2019.11.5(Mon)

  - 今天一早 VSCode 給我推薦安裝了這些包，我要來看他們到底是什麼

  ```bash
    gocode
    gopkgs
    go-outline
    go-symbols
    guru
    gorename
    gotests
    gomodifytags
    impl
    fillstruct
    goplay
    godoctor
    dlv
    gocode-gomod
    godef
    goreturns
    golint
  ```
  
  ~~結果發現，只要 *.go 存擋，就會自動幫我 gofmt。~~ 不是 gofmt，因為會沒用到的 import 會被刪除，但我測試一下，發現 gofmt 並不會做這件事。

  - 了解 gofmt
    - 類似於 Python 的 pep
    - 輸入 $gofmt 會 hang 住
    - 正確的是 $go fmt *.go 會自動修正 *.go 的格式
    - 參數 -n -x 沒作用，還沒搞懂問題在哪裡？目前只能透過 git 來觀察 gofmt 到底改了哪些東西
    - 發現自己不小心加入很多結尾分號，但不影響 go 執行
    - 這真的是強迫症的懶人好工具
  - [理解Go語言的nil](https://www.jianshu.com/p/dd80f6be7969)
    - 先跳過二元樹、map、function、channel
  
  還沒找到 runtime 才決定型別的 slice 產生方法，不能以 TypeOf 的 return 當作型別用

```Go
s3 := make(reflect.TypeOf(s1), len(s1))
//error: reflect.TypeOf(s1) is not a type
```
目前找到的答案看起來很複雜 https://stackoverflow.com/questions/39363798/how-to-create-a-slice-of-variable-type-in-go
    

---

### 2019.11.1(Fri)

  - 編譯語言居然有當直譯用的模式，go run hello.go 就可以用了，太神奇了。這樣不論是在學習或使用都方便很多。
  - Go 語法與 Python 相似的地方很多，寫起來非常像。
  - Go 特別的地方
      - import 了沒用到的包、或宣告了沒用到的變數，都會無法通過編譯，其他語言都只是 warning。
      - 型別是宣告在變數的後面，這是第一次看到。
      - 對括號位置有強制性要求（這比 Python 只用縮排劃出 scope 可讀性好多了）
      - 對「:=」這個神秘的運算子還不太了解，希望找到書能講的清楚一些。
