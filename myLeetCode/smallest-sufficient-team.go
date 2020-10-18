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
			Req_skills: []string{"wmycibrjxh", "wicacrldwneag", "ndutqtjuzu", "pgo", "gxsskiz", "rbrymc", "erpevpmu", "jboexi", "vpfdcjwngzuf", "w"},
			People:     [][]string{{}, {"ndutqtjuzu", "pgo"}, {"ndutqtjuzu"}, {"pgo", "rbrymc"}, {"wicacrldwneag", "ndutqtjuzu"}, {}, {"wicacrldwneag", "rbrymc", "erpevpmu"}, {"w"}, {"wmycibrjxh", "wicacrldwneag", "pgo", "w"}, {}, {"w"}, {"gxsskiz", "erpevpmu", "vpfdcjwngzuf"}, {"wicacrldwneag"}, {"vpfdcjwngzuf"}, {"wmycibrjxh", "erpevpmu"}, {"ndutqtjuzu", "pgo"}, {"ndutqtjuzu", "pgo"}, {"wmycibrjxh", "erpevpmu", "jboexi"}, {"wmycibrjxh", "wicacrldwneag", "jboexi"}, {"wmycibrjxh", "wicacrldwneag", "rbrymc"}, {"wicacrldwneag"}, {"erpevpmu", "vpfdcjwngzuf"}, {"wmycibrjxh"}, {"jboexi", "w"}, {"erpevpmu", "jboexi", "w"}, {"w"}, {"erpevpmu", "jboexi"}, {"jboexi"}, {"wicacrldwneag"}, {}, {"jboexi", "vpfdcjwngzuf"}, {"wmycibrjxh", "jboexi"}, {"wicacrldwneag"}, {}, {"pgo"}, {"wicacrldwneag"}, {}, {"wmycibrjxh", "vpfdcjwngzuf"}, {"wmycibrjxh"}, {"pgo", "vpfdcjwngzuf", "w"}, {"wicacrldwneag", "jboexi"}, {"wicacrldwneag", "erpevpmu", "vpfdcjwngzuf"}, {"wicacrldwneag"}, {"wmycibrjxh", "pgo", "erpevpmu", "vpfdcjwngzuf"}, {"w"}, {"vpfdcjwngzuf", "w"}, {"wmycibrjxh", "erpevpmu"}, {"wicacrldwneag", "pgo", "jboexi"}, {"wmycibrjxh", "erpevpmu", "vpfdcjwngzuf"}, {"w"}, {}, {}, {}, {"pgo", "jboexi"}, {"wicacrldwneag"}, {"wicacrldwneag", "erpevpmu", "jboexi"}, {"wmycibrjxh", "pgo"}, {"wmycibrjxh", "wicacrldwneag", "gxsskiz"}, {"erpevpmu"}, {"pgo", "rbrymc", "erpevpmu", "w"}},
			Answer:     []int{},
		},
		TestCase{
			Req_skills: []string{"hkyodbbhr", "p", "biflxurxdvb", "x", "qq", "yhiwcn"},
			People:     [][]string{{"yhiwcn"}, {}, {}, {}, {"biflxurxdvb", "yhiwcn"}, {"hkyodbbhr"}, {"hkyodbbhr", "p"}, {"hkyodbbhr"}, {}, {"yhiwcn"}, {"hkyodbbhr", "qq"}, {"qq"}, {"hkyodbbhr"}, {"yhiwcn"}, {}, {"biflxurxdvb"}, {}, {"hkyodbbhr"}, {"hkyodbbhr", "yhiwcn"}, {"yhiwcn"}, {"hkyodbbhr"}, {"hkyodbbhr", "p"}, {}, {}, {"hkyodbbhr"}, {"biflxurxdvb"}, {"qq", "yhiwcn"}, {"hkyodbbhr", "yhiwcn"}, {"hkyodbbhr"}, {}, {}, {"hkyodbbhr"}, {}, {"yhiwcn"}, {}, {"hkyodbbhr"}, {"yhiwcn"}, {"yhiwcn"}, {}, {}, {"hkyodbbhr", "yhiwcn"}, {"yhiwcn"}, {"yhiwcn"}, {}, {}, {}, {"yhiwcn"}, {}, {"yhiwcn"}, {"x"}, {"hkyodbbhr"}, {}, {}, {"yhiwcn"}, {}, {"biflxurxdvb"}, {}, {}, {"hkyodbbhr", "biflxurxdvb", "yhiwcn"}, {}},
			Answer:     []int{21, 26, 49, 58},
		},
		TestCase{
			Req_skills: []string{"hfkbcrslcdjq", "jmhobexvmmlyyzk", "fjubadocdwaygs", "peaqbonzgl", "brgjopmm", "x", "mf", "pcfpppaxsxtpixd", "ccwfthnjt", "xtadkauiqwravo", "zezdb", "a", "rahimgtlopffbwdg", "ulqocaijhezwfr", "zshbwqdhx", "hyxnrujrqykzhizm"},
			People:     [][]string{{"peaqbonzgl", "xtadkauiqwravo"}, {"peaqbonzgl", "pcfpppaxsxtpixd", "zshbwqdhx"}, {"x", "a"}, {"a"}, {"jmhobexvmmlyyzk", "fjubadocdwaygs", "xtadkauiqwravo", "zshbwqdhx"}, {"fjubadocdwaygs", "x", "zshbwqdhx"}, {"x", "xtadkauiqwravo"}, {"x", "hyxnrujrqykzhizm"}, {"peaqbonzgl", "x", "pcfpppaxsxtpixd", "a"}, {"peaqbonzgl", "pcfpppaxsxtpixd"}, {"a"}, {"hyxnrujrqykzhizm"}, {"jmhobexvmmlyyzk"}, {"hfkbcrslcdjq", "xtadkauiqwravo", "a", "zshbwqdhx"}, {"peaqbonzgl", "mf", "a", "rahimgtlopffbwdg", "zshbwqdhx"}, {"xtadkauiqwravo"}, {"fjubadocdwaygs"}, {"x", "a", "ulqocaijhezwfr", "zshbwqdhx"}, {"peaqbonzgl"}, {"pcfpppaxsxtpixd", "ulqocaijhezwfr", "hyxnrujrqykzhizm"}, {"a", "ulqocaijhezwfr", "hyxnrujrqykzhizm"}, {"a", "rahimgtlopffbwdg"}, {"zshbwqdhx"}, {"fjubadocdwaygs", "peaqbonzgl", "brgjopmm", "x"}, {"hyxnrujrqykzhizm"}, {"jmhobexvmmlyyzk", "a", "ulqocaijhezwfr"}, {"peaqbonzgl", "x", "a", "ulqocaijhezwfr", "zshbwqdhx"}, {"mf", "pcfpppaxsxtpixd"}, {"fjubadocdwaygs", "ulqocaijhezwfr"}, {"fjubadocdwaygs", "x", "a"}, {"zezdb", "hyxnrujrqykzhizm"}, {"ccwfthnjt", "a"}, {"fjubadocdwaygs", "zezdb", "a"}, {}, {"peaqbonzgl", "ccwfthnjt", "hyxnrujrqykzhizm"}, {"xtadkauiqwravo", "hyxnrujrqykzhizm"}, {"peaqbonzgl", "a"}, {"x", "a", "hyxnrujrqykzhizm"}, {"zshbwqdhx"}, {}, {"fjubadocdwaygs", "mf", "pcfpppaxsxtpixd", "zshbwqdhx"}, {"pcfpppaxsxtpixd", "a", "zshbwqdhx"}, {"peaqbonzgl"}, {"peaqbonzgl", "x", "ulqocaijhezwfr"}, {"ulqocaijhezwfr"}, {"x"}, {"fjubadocdwaygs", "peaqbonzgl"}, {"fjubadocdwaygs", "xtadkauiqwravo"}, {"pcfpppaxsxtpixd", "zshbwqdhx"}, {"peaqbonzgl", "brgjopmm", "pcfpppaxsxtpixd", "a"}, {"fjubadocdwaygs", "x", "mf", "ulqocaijhezwfr"}, {"jmhobexvmmlyyzk", "brgjopmm", "rahimgtlopffbwdg", "hyxnrujrqykzhizm"}, {"x", "ccwfthnjt", "hyxnrujrqykzhizm"}, {"hyxnrujrqykzhizm"}, {"peaqbonzgl", "x", "xtadkauiqwravo", "ulqocaijhezwfr", "hyxnrujrqykzhizm"}, {"brgjopmm", "ulqocaijhezwfr", "zshbwqdhx"}, {"peaqbonzgl", "pcfpppaxsxtpixd"}, {"fjubadocdwaygs", "x", "a", "zshbwqdhx"}, {"fjubadocdwaygs", "peaqbonzgl", "x"}, {"ccwfthnjt"}},
		},
		TestCase{
			Req_skills: []string{"zp","jpphhnhwpw","pscleb","arn","acrsxqvus","fseqih","fpqbjbbxglivyonn","cjozlkyodt","mvtwffgkhjrtibto","kumdvfwsvrht","i","s","ucr","oo","yqkqkhhhwngyjrg","odiwidzqw"},
			People:     [][]string{{"acrsxqvus"},{"zp"},{},{"fpqbjbbxglivyonn"},{},{},{"kumdvfwsvrht"},{},{"oo"},{},{"fseqih"},{},{"arn"},{},{},{"yqkqkhhhwngyjrg"},{},{},{},{"kumdvfwsvrht"},{"s"},{},{},{"zp","ucr"},{},{"pscleb"},{},{},{},{},{},{},{},{"jpphhnhwpw"},{},{},{},{"oo"},{},{"i"},{"pscleb"},{},{},{},{},{},{},{"i"},{},{"mvtwffgkhjrtibto","odiwidzqw"},{},{"cjozlkyodt","odiwidzqw"},{"arn"},{},{},{"acrsxqvus"},{},{},{},{"ucr"}},
		},
	}

	for _, testCase := range testCases {
		fmt.Printf("len(testCase.Req_skills)=%d, len(testCase.People)=%d\n", len(testCase.Req_skills), len(testCase.People))
		fmt.Println(smallestSufficientTeam(testCase.Req_skills, testCase.People))
		fmt.Println("Answer: ", testCase.Answer)
	}
}

// 以上不進 LeetCode
var spentComputeTimes = 0
var savedComputeTimes = 0

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

	//剔除能力可以被其他人覆蓋的候選人，確保留下的候選人具有直觀上的不可取代性
	for i := 0; i < len(myCandidates); i++ {
		for j := i; j < len(myCandidates); j++ {
			OrResult := Or(myCandidates[i].MatchIndex, myCandidates[j].MatchIndex)
			i_EqRes := Equal(OrResult, myCandidates[i].MatchIndex)
			j_EqRes := Equal(OrResult, myCandidates[j].MatchIndex)

			//剔除不能改變結果的那一邊
			if i_EqRes != j_EqRes {
				if i_EqRes {
					myCandidates = append(myCandidates[:j], myCandidates[j+1:]...)
				}
				if j_EqRes {
					myCandidates = append(myCandidates[:i], myCandidates[i+1:]...)
				}
				j, i = -1, 0 // 交叉對比的矩陣長寬同時改變，所以要回到原點再比較
			}
		}
	}
	fmt.Printf("Cutted len(myCandidates)=%d\n", len(myCandidates))

	//將候選人按照命中機率排序，高的優先
	sort.Sort(sort.Reverse(myCandidates))

	samllestKey := IntSliceToString(DFS_Match(myCandidates, len(req_skills)))

	fmt.Println("samllestKey=", samllestKey)
	fmt.Println("tableComputed[samllestKey]=", tableComputed[samllestKey])
	fmt.Println("spentComputeTimes=", spentComputeTimes)
	fmt.Println("savedComputeTimes=", savedComputeTimes)
	return IntStringToIntSlice(samllestKey)
}

func DFS_Match(myCandidates candidates, depth int) []int {

	// 只要還沒到最底層，就去看更深一層
	if 1 == depth{
		for _, c := range myCandidates {
			//存入計算結果，滿足解答就 return
			tableComputed[strconv.Itoa(c.Uid)] = c.MatchIndex
			//fmt.Println("strKey=",strconv.Itoa(c.Uid))
			if IsNonZero(c.MatchIndex) {
				return []int{c.Uid}
			}
		}
		return nil //代表這一層沒有答案出現
	} else {
		result := DFS_Match(myCandidates, depth-1)
		// // 更深的層有結果了，此層免算
		if nil != result {
			return result
		}
	}
	
	//更深的層都沒有答案，所以要把這一層算完
	for _, c := range myCandidates {
		for k, v := range tableComputed {
			
			prevKey, currKey := IntStringToIntSlice(k), c.Uid
			
			//是否為這層要計算的長度？
			if len(prevKey)+1 != depth{
				continue
			}

			//檢驗是否有重複key？
			if containsInt(prevKey, currKey) {
				savedComputeTimes++
				continue //key 重複了，此圈免算
			}

			//生成合併排序key
			strKey := func(prevKey []int, currKey int) string {
				prevKey = append(prevKey, currKey)
				sort.Sort(sort.IntSlice(prevKey))
				return IntSliceToString(prevKey)
			}(prevKey, currKey)

			//檢查此key有沒有算過？
			if _, ok := tableComputed[strKey]; ok {
				savedComputeTimes++
				continue //算過，此圈免算
			}

			//fmt.Printf("len(strKey)=%d, strkey=%v\n",len(IntStringToIntSlice(strKey)), strKey)
			
			//存入計算結果，滿足解答就 return
			spentComputeTimes++
			tableComputed[strKey] = Or(v, c.MatchIndex)
			if IsNonZero(tableComputed[strKey]) {
				return IntStringToIntSlice(strKey)
			}
		}
	}

	return nil //代表這一層沒有答案出現
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

func containsInt(s []int, e int) bool {
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
	var outStrSlice []string
	for _, num := range nums {
		outStrSlice = append(outStrSlice, strconv.Itoa(num))
	}
	return strings.Join(outStrSlice, "-")
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
