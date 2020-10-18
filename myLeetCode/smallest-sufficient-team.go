package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type TestCase struct {
	Req_skills []string
	People     [][]string
	Answer     []int
}

func R(nums []int) {
	fmt.Println(nums)
	if 1 != len(nums) {
		for skip_i := range nums {
			new_nums := make([]int, len(nums))
			copy(new_nums, nums)
			new_nums = append(new_nums[:skip_i], new_nums[skip_i+1:]...)
			R(new_nums)
		}
	}
}

func main() {
	//nums := []int{0, 1, 2}
	//R(nums)

	var testCases = []TestCase{
		TestCase{
			Req_skills: []string{"java", "nodejs", "reactjs"},
			People:     [][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
			Answer:     []int{0, 2},
		},
		TestCase{
			Req_skills: []string{"cdkpfwkhlfbps", "hnvepiymrmb", "cqrdrqty", "pxivftxovnpf", "uefdllzzmvpaicyl", "idsyvyl"},
			People:     [][]string{{""}, {"hnvepiymrmb"}, {"uefdllzzmvpaicyl"}, {""}, {"hnvepiymrmb", "cqrdrqty"}, {"pxivftxovnpf"}, {"hnvepiymrmb", "pxivftxovnpf"}, {"hnvepiymrmb"}, {"cdkpfwkhlfbps"}, {"idsyvyl"}, {}, {"cdkpfwkhlfbps", "uefdllzzmvpaicyl"}, {"cdkpfwkhlfbps", "uefdllzzmvpaicyl"}, {"pxivftxovnpf", "uefdllzzmvpaicyl"}, {""}, {"cqrdrqty"}, {""}, {"cqrdrqty", "pxivftxovnpf", "idsyvyl"}, {"hnvepiymrmb", "idsyvyl"}, {""}},
			Answer:     []int{12, 17, 18},
		},
		TestCase{
			Req_skills: []string{"mmcmnwacnhhdd", "vza", "mrxyc"},
			People:     [][]string{{"mmcmnwacnhhdd"}, {}, {}, {"vza", "mrxyc"}},
			Answer:     []int{0, 3},
		},
		TestCase{
			Req_skills: []string{"hdbxcuzyzhliwv", "uvwlzkmzgis", "sdi", "bztg", "ylopoifzkacuwp", "dzsgleocfpl"},
			People:     [][]string{{"hdbxcuzyzhliwv", "dzsgleocfpl"}, {"hdbxcuzyzhliwv", "sdi", "ylopoifzkacuwp", "dzsgleocfpl"}, {"bztg", "ylopoifzkacuwp"}, {"bztg", "dzsgleocfpl"}, {"hdbxcuzyzhliwv", "bztg"}, {"dzsgleocfpl"}, {"uvwlzkmzgis"}, {"dzsgleocfpl"}, {"hdbxcuzyzhliwv"}, {}, {"dzsgleocfpl"}, {"hdbxcuzyzhliwv"}, {}, {"hdbxcuzyzhliwv", "ylopoifzkacuwp"}, {"sdi"}, {"bztg", "dzsgleocfpl"}, {"hdbxcuzyzhliwv", "uvwlzkmzgis", "sdi", "bztg", "ylopoifzkacuwp"}, {"hdbxcuzyzhliwv", "sdi"}, {"hdbxcuzyzhliwv", "ylopoifzkacuwp"}, {"sdi", "bztg", "ylopoifzkacuwp", "dzsgleocfpl"}, {"dzsgleocfpl"}, {"sdi", "ylopoifzkacuwp"}, {"hdbxcuzyzhliwv", "uvwlzkmzgis", "sdi"}, {}, {}, {"ylopoifzkacuwp"}, {}, {"sdi", "bztg"}, {"bztg", "dzsgleocfpl"}, {"sdi", "bztg"}},
		},
		TestCase{
			Req_skills: []string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
			People:     [][]string{{"algorithms", "math", "java"}, {"algorithms", "math", "reactjs"}, {"java", "csharp", "aws"}, {"reactjs", "csharp"}, {"csharp", "math"}, {"aws", "java"}},
			Answer:     []int{1, 2},
		},
		TestCase{
			Req_skills: []string{"wmycibrjxh","wicacrldwneag","ndutqtjuzu","pgo","gxsskiz","rbrymc","erpevpmu","jboexi","vpfdcjwngzuf","w"},
			People:     [][]string{{},{"ndutqtjuzu","pgo"},{"ndutqtjuzu"},{"pgo","rbrymc"},{"wicacrldwneag","ndutqtjuzu"},{},{"wicacrldwneag","rbrymc","erpevpmu"},{"w"},{"wmycibrjxh","wicacrldwneag","pgo","w"},{},{"w"},{"gxsskiz","erpevpmu","vpfdcjwngzuf"},{"wicacrldwneag"},{"vpfdcjwngzuf"},{"wmycibrjxh","erpevpmu"},{"ndutqtjuzu","pgo"},{"ndutqtjuzu","pgo"},{"wmycibrjxh","erpevpmu","jboexi"},{"wmycibrjxh","wicacrldwneag","jboexi"},{"wmycibrjxh","wicacrldwneag","rbrymc"},{"wicacrldwneag"},{"erpevpmu","vpfdcjwngzuf"},{"wmycibrjxh"},{"jboexi","w"},{"erpevpmu","jboexi","w"},{"w"},{"erpevpmu","jboexi"},{"jboexi"},{"wicacrldwneag"},{},{"jboexi","vpfdcjwngzuf"},{"wmycibrjxh","jboexi"},{"wicacrldwneag"},{},{"pgo"},{"wicacrldwneag"},{},{"wmycibrjxh","vpfdcjwngzuf"},{"wmycibrjxh"},{"pgo","vpfdcjwngzuf","w"},{"wicacrldwneag","jboexi"},{"wicacrldwneag","erpevpmu","vpfdcjwngzuf"},{"wicacrldwneag"},{"wmycibrjxh","pgo","erpevpmu","vpfdcjwngzuf"},{"w"},{"vpfdcjwngzuf","w"},{"wmycibrjxh","erpevpmu"},{"wicacrldwneag","pgo","jboexi"},{"wmycibrjxh","erpevpmu","vpfdcjwngzuf"},{"w"},{},{},{},{"pgo","jboexi"},{"wicacrldwneag"},{"wicacrldwneag","erpevpmu","jboexi"},{"wmycibrjxh","pgo"},{"wmycibrjxh","wicacrldwneag","gxsskiz"},{"erpevpmu"},{"pgo","rbrymc","erpevpmu","w"}},
			Answer:     []int{},
		},
	}

	for _, testCase := range testCases {
		fmt.Println(smallestSufficientTeam(testCase.Req_skills, testCase.People))
	}
}

// 以上不進 LeetCode

var savedComputeTimes = 0
var spentComputeTimes = 0

type People struct {
	Uid         int
	MatchIndex  []int
	NumOfMatchs int
}

type candidates []People

func (c candidates) Len() int {
	return len(c)
}
func (c candidates) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c candidates) Less(i, j int) bool {
	return c[i].NumOfMatchs < c[j].NumOfMatchs
}

var tableComputed map[string]([]int)

func smallestSufficientTeam(req_skills []string, people [][]string) []int {
	tableComputed = make(map[string][]int)

	//建立對候選人技能對應列清單
	var myCandidates candidates
	for i, p := range people {
		theCandidate := MatchSkills(req_skills, p)
		if 0 < theCandidate.NumOfMatchs {
			theCandidate.Uid = i
			myCandidates = append(myCandidates, theCandidate)
		}
	}

	//剔除能力可以被其他人覆蓋的候選人
	for i := 0; i < len(myCandidates); i++ {
		//fmt.Printf("len(myCandidates)=%v, i=%v\n",len(myCandidates),i)
		for j := i; j < len(myCandidates); j++ {
			//fmt.Printf("len(myCandidates)=%v, j=%v\n",len(myCandidates),j)
			OrResult := Or(myCandidates[i].MatchIndex, myCandidates[j].MatchIndex)
			i_EqRes := Equal(OrResult, myCandidates[i].MatchIndex)
			j_EqRes := Equal(OrResult, myCandidates[j].MatchIndex)

			//剔除不能改變結果的那一邊
			if i_EqRes && !j_EqRes {
				// fmt.Println("")
				// fmt.Printf("myCandidates[%2d].MatchIndex = %v\n", i, myCandidates[i].MatchIndex)
				// fmt.Printf("myCandidates[%2d].MatchIndex = %v\n", j, myCandidates[j].MatchIndex)
				// fmt.Println("OrResult                   = ", OrResult)
				// fmt.Printf("Delete myCandidates[%2d]\n", j)
				myCandidates = append(myCandidates[:j], myCandidates[j+1:]...)
				j--
			} else if j_EqRes && !i_EqRes {
				// fmt.Println("")
				// fmt.Printf("myCandidates[%2d].MatchIndex = %v\n", i, myCandidates[i].MatchIndex)
				// fmt.Printf("myCandidates[%2d].MatchIndex = %v\n", j, myCandidates[j].MatchIndex)
				// fmt.Println("OrResult                   = ", OrResult)
				// fmt.Printf("Delete myCandidates[%2d]\n", i)
				myCandidates = append(myCandidates[:i], myCandidates[i+1:]...)
				//i--
			}
			//fmt.Println("")
		}
	}

	Recursive(myCandidates, len(req_skills))

	samllestKey := "---------------------------------------------------------------"
	for k, v := range tableComputed {
		// fmt.Println(k, v)
		if IsNonZero(v) {
			if len(samllestKey) > len(k) {
				samllestKey = k
			}
		}
	}

	fmt.Println("samllestKey=", samllestKey)
	fmt.Println("tableComputed[samllestKey]=", tableComputed[samllestKey])
	fmt.Println("savedComputeTimes=", savedComputeTimes)
	fmt.Println("spentComputeTimes=", spentComputeTimes)
	return IntStringToIntSlice(samllestKey)
}

func Recursive(myCandidates candidates, len_req_skills int) []int {
	//如果遞迴到只剩下一個
	if 1 == len(myCandidates) {
		tableComputed[strconv.Itoa(myCandidates[0].Uid)] = myCandidates[0].MatchIndex
		return myCandidates[0].MatchIndex
	}

	//生成ID
	var temp []int
	for _, c := range myCandidates {
		temp = append(temp, c.Uid)
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(temp)))
	sort.Sort(sort.IntSlice(temp))
	strID := IntSliceToString(temp)
	//fmt.Println("ID=", strID)

	if v, ok := tableComputed[strID]; ok {
		return v
	} else {
		for skip_i, c := range myCandidates {
			//生成子集ID
			splited_strID := strings.Split(strID, "-")
			splited_strID = append(splited_strID[:skip_i], splited_strID[skip_i+1:]...)
			str_subset_ID := strings.Join(splited_strID, "-")

			if _, ok := tableComputed[str_subset_ID]; !ok {
				spentComputeTimes++

				//安全複製子集
				subset_myCandidates := make([]People, len(myCandidates))
				copy(subset_myCandidates, myCandidates)
				subset_myCandidates = append(subset_myCandidates[:skip_i], subset_myCandidates[skip_i+1:]...)

				//計算子集
				//為子集建立key, value
				tableComputed[str_subset_ID] = Recursive(subset_myCandidates, len_req_skills)
			} else {
				savedComputeTimes++
			}

			if 0 == skip_i {
				tableComputed[strID] = Or(c.MatchIndex, tableComputed[str_subset_ID])
			}
		}
	}

	return tableComputed[strID]
}

func MatchSkills(req_skills []string, peopleHaveSkills []string) People {

	p := People{MatchIndex: make([]int, len(req_skills))}

	for i, req_skill := range req_skills {
		found := false
		for _, peopleHaveSkill := range peopleHaveSkills {
			if req_skill == peopleHaveSkill {
				p.MatchIndex[i] = 1
				p.NumOfMatchs++
				found = true
				break
			}
		}

		if !found {
			p.MatchIndex[i] = 0
		}
	}

	return p
}

func IsSubslice(mainSlice []string, subSlice []string) bool {
	if len(mainSlice) > len(subSlice) {
		return false
	}
	for _, e := range mainSlice {
		if !contains(subSlice, e) {
			return false
		}
	}
	return true
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func IsNonZero(nums []int) bool {
	for _, num := range nums {
		if 0 == num {
			return false
		}
	}
	return true
}

func Or(a, b []int) []int {
	if len(a) != len(b) {
		return []int{-1}
	}

	out := make([]int, len(a))

	for i, _ := range out {
		out[i] = a[i] | b[i]
	}
	return out
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func IntSliceToString(nums []int) string {
	var outStr string
	for i, num := range nums {
		if 0 == i {
			outStr = strconv.Itoa(num)
		} else {
			outStr = outStr + "-" + strconv.Itoa(num)
		}
	}
	return outStr
}

func IntStringToIntSlice(strOfInt string) []int {
	var intSlice []int
	_strOfInt := strings.Split(strOfInt, "-")

	for _, s := range _strOfInt {
		num, _ := strconv.Atoi(s)
		intSlice = append(intSlice, num)
	}
	return intSlice
}
