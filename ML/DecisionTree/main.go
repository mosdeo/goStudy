package main

import (
	"fmt"
	"math"
	"strconv"
	//rand"
	// "time"
)

var columnsName = []string{"编号", "长相", "性格", "年龄", "学历", "住址", "小木是否选择"}

var samples = [][]string{
	//{编号 长相 性格 年龄 学历 住址 小木是否选择}
	{"1", "白", "温柔", "26", "硕士", "城市", "否"},
	{"2", "黑", "平淡", "24", "高中", "城市", "否"},
	{"3", "白", "超凶", "20", "本科", "城市", "是"},
	{"4", "黑", "温柔", "25", "硕士", "农村", "否"},
	{"5", "白", "平淡", "30", "本科", "城市", "否"},
	{"6", "黑", "超凶", "28", "博士", "农村", "是"},
}

func Transformer(samples [][]string) [6][7]float64 {
	encodedSamples := [6][7]float64{{}}

	//文字特徵映射表
	lookup := map[string]float64{
		"白": 0, "黑": 1,
		"温柔": 0, "平淡": 1, "超凶": 2,
		"高中": 0, "本科": 1, "硕士": 2,
		"城市": 0, "农村": 1,
		"否": 0, "是": 1,
	}

	//特徵與標籤轉換
	for i, sample := range samples {
		for j, feature := range sample {
			if 3 == j {
				//年齡分開轉換
				if s, err := strconv.ParseFloat(samples[i][3], 64); err == nil {
					encodedSamples[i][3] = s
				}
			} else {
				//其他查表轉換
				encodedSamples[i][j] = lookup[feature]
			}
		}
	}

	return encodedSamples
}

func Entropy(samples [][]string) []float64 {

	H := make([]float64, len(samples[0]))

	//計算第i個特徵的熵
	for i := 1; i < len(samples[0])-1; i++ {
		featureCount := map[string][]int{}

		for _, s := range samples {
			// fmt.Printf("s[%d]=%v\n", i, s[i])
			//計算每一種標籤的數量
			if _, ok := featureCount[s[i]]; !ok {
				featureCount[s[i]] = make([]int, 2)
			}
			featureCount[s[i]][0]++

			//計算這種標籤被選中的次數
			if "是" == s[len(samples[0])-1] {
				featureCount[s[i]][1]++
			}
		}

		//計算每一種標籤的機率、第i個特徵的熵
		for k, v := range featureCount {
			theFeatureSelectedRate := (float64)(v[1]) / (float64)(v[0])
			// fmt.Printf("%v, theFeatureSelectedRate=%v\n", k, theFeatureSelectedRate)
			samplesNum := len(samples)

			p := float64(v[0]) / float64(samplesNum)

			logTerm1, logTerm2 := 0.0, 0.0
			if 0 == theFeatureSelectedRate || 1 == theFeatureSelectedRate {
				logTerm1, logTerm2 = 0, 0
			} else {
				logTerm1 = theFeatureSelectedRate * math.Log2(theFeatureSelectedRate)
				logTerm2 = (1 - theFeatureSelectedRate) * math.Log2(1-theFeatureSelectedRate)
			}

			H[i] -= p * (logTerm1 + logTerm2)
			fmt.Printf("k:%v => H[%v] -= (%v/%v) * ( %v + %v )\n", k, i, v[0], samplesNum, logTerm1, logTerm2)
		}
	}

	return H[1 : len(H)-1]
}

// func InformationGain(s, a float64) float64{

// }

var model = []int{0, 0, 0}

func main() {
	fmt.Println("原始樣本:")
	for _, s := range samples {
		fmt.Println(s)
	}

	fmt.Println("編碼後樣本:")
	for _, ts := range Transformer(samples) {
		fmt.Println(ts)
	}

	entropy := Entropy(samples)

	fmt.Println("Entropy for each feature:")
	fmt.Printf("H(i): %v\n", entropy)
	fmt.Printf("      %v\n", columnsName[1:len(columnsName)-1])

	sortedEntropyIndex := GetSortedIndex(entropy)

	fmt.Println(sortedEntropyIndex)
}

// 取得每一個位置的大小順序index
func GetSortedIndex(nums []float64) []int {

	sortedIndex := make([]int, len(nums))
	for i := range sortedIndex {
		sortedIndex[i] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			// fmt.Printf("i=%d, j=%d\n", i, j)
			// fmt.Println(nums[sortedIndex[0]], nums[sortedIndex[1]], nums[sortedIndex[2]], nums[sortedIndex[3]], nums[sortedIndex[4]])
			if nums[sortedIndex[i]] < nums[sortedIndex[j]] {
				sortedIndex[i], sortedIndex[j] = sortedIndex[j], sortedIndex[i]
			}
		}
	}

	// fmt.Println(nums[sortedIndex[0]], nums[sortedIndex[1]], nums[sortedIndex[2]], nums[sortedIndex[3]], nums[sortedIndex[4]])

	return sortedIndex
}
