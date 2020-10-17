package main

import (
	"fmt"
	"strconv"
	"sort"
)

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
		theCandidate.Uid = i
		myCandidates = append(myCandidates, theCandidate)
	}

	SufficientExam(myCandidates, len(req_skills))

	//窮舉檢查
	var keySet []int
	for k, _ := range tableComputed {
		intkey, _ := strconv.Atoi(k)
		keySet = append(keySet, intkey)
	}
	return keySet

	samllestKey := "---------------------------------------------------------------"
	for k, v := range tableComputed {
		if IsNonZero(v) {
			if len(samllestKey) > len(k) {
				samllestKey = k
			}
		}
	}

	var outputNums []int
	for _, s := range samllestKey {
		outputNums = append(outputNums, int(s)-48)
	}

	return outputNums
}

func main() {
	req_skills := []string{"java", "nodejs", "reactjs"}
	people := [][]string{{"java"}, {"nodejs"}, {"reactjs"}, {"nodejs", "reactjs"}}
	fmt.Print(smallestSufficientTeam(req_skills, people))
}

func SufficientExam(myCandidates candidates, len_req_skills int) {

	//如果遞迴到只剩下一個
	if 1 == len(myCandidates) {
		tableComputed[strconv.Itoa(myCandidates[0].Uid)] = myCandidates[0].MatchIndex
		return
	}

	//建立候選人清單單一index
	current_candidates_idx := []int{}
	for _, c := range myCandidates {
		current_candidates_idx = append(current_candidates_idx, c.Uid)
	}
	sort.Reverse(sort.IntSlice(current_candidates_idx))
	str_current_candidates_idx := ""
	for i := range current_candidates_idx {
		str_current_candidates_idx += strconv.Itoa(i)
	}

	//檢查這組合是否被計算過
	if _, ok := tableComputed[str_current_candidates_idx]; ok {

	} else {
		//找尋有被計算過的各種-1子集
		for i := 0; i < len(str_current_candidates_idx); i++ {
			//產生-1子集 key
			str_subset_candidates_idx := str_current_candidates_idx[:i] + str_current_candidates_idx[i+1:]

			if subset_v, ok := tableComputed[str_subset_candidates_idx]; ok {
				// 有現成子集
				tableComputed[str_current_candidates_idx] = Or(myCandidates[i].MatchIndex, subset_v)
			} else {
				//無現成子集
				subset_myCandidates := append(myCandidates[:i], myCandidates[i+1:]...)
				SufficientExam(subset_myCandidates, len_req_skills)
				tableComputed[str_current_candidates_idx] = Or(myCandidates[i].MatchIndex, tableComputed[str_subset_candidates_idx])
			}
		}
	}
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
