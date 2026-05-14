package main

import (
	"fmt"
	"strings"
)

func similarPairs(words []string) int {
	res := 0
	cnt := make(map[int]int)
	for _, word := range words {
		state := 0
		for _, c := range word {
			state |= 1 << (c - 'a')
		}
		res += cnt[state]
		cnt[state]++
	}
	return res
}

func main() {
	//fmt.Println(similarPairs([]string{"aba", "aabb", "abcd", "bac", "aabc"}))
	//fmt.Println(similarPairs([]string{"aabb", "ab", "ba"}))
	//fmt.Println(similarPairs([]string{"nba", "cba", "dba"}))

	instance := `128
         682
         735
         738
         742
         744
         747
         751
         754
         756
         759
         763
         765
         768
         772
         775
         776
         778
         779
         780
         790
         791
         792
         795
         806
         807
         808
         810
         821
         822
         823
         824
         828
         833
         834
         835
         836
         840
         844
         845
         846
         847
         857
         858
         860
         861
         862
         875
         876
         877
         878
         882
         887
         888
         889
         890
         901
         902
         903
         904
         905
         914
         915
         916
         917
         931
         932
         933
         934
         935
         944
         945
         946
         947
         948
        1013
        1014
        1017
        1021
        1305
        1306
        1307
        1308
        1310
        1379
        1381
        1382
        1383
        1384
        1393
        1394
        1395
        1396
        1465
        1466
        1467
        1468
        1469
        1481
        1482`

	instance = strings.Replace(instance, " ", "", -1)
	instanceArr := strings.Split(instance, "\n")

	fmt.Println(strings.Join(instanceArr, ","))

}
