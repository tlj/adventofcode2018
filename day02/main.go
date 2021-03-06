package main

import (
	"fmt"
	"github.com/tlj/adventofcode2018/aoc"
)

func letterCountPerLine(line string) map[int32]int {
	letterCount := make(map[int32]int)

	for _, c := range line {
		letterCount[c]++
	}

	return letterCount
}

func hasCount(letterCount map[int32]int, count int) bool {
	for _, c := range letterCount {
		if c == count {
			return true
		}
	}
	return false
}

func main() {
	part1()
	part2()
}

func part1() {
	lines := aoc.LoadAsLines(input)

	twos := 0
	threes := 0
	for _, line := range lines {
		letterCount := letterCountPerLine(line)
		if hasCount(letterCount, 2) {
			twos++
		}
		if hasCount(letterCount, 3) {
			threes++
		}
	}

	fmt.Printf("Part1 (%d * %d): %d\n", twos, threes, twos * threes)
}

func part2() {
	lines := aoc.LoadAsLines(input)

	var closestString string
	var closestCount int

	for sIdx, source := range lines {
		for tIdx, target := range lines {
			if tIdx <= sIdx { continue }

			sameCount := 0
			var tempString string
			for si := range source {
				if target[si] == source[si] {
					sameCount++
					tempString += string(target[si])
				}
			}
			if sameCount > closestCount {
				closestString = tempString
				closestCount = sameCount
			}
		}
	}

	fmt.Printf("Part2: %s\n", closestString)
}

var input = `lnfgdsywjyleogambzuchirkpx
nnfqdskfjyteogambzuchirkpx
lnfqdvvwjyteofambzuchirkpf
lnfqdsvwjyteogvmbzuthirkpn
ltfqdsvwjyoeogambxuchirkpx
lnfqcsvwjytzogacbzuchirkpx
lpfpdsvwjyteogambyuchirkpx
pnfqdsvwjyteogqmbzuchinkpx
lnfqdsvwjytyopambzpchirkpx
lnfqisswjyteogadbzuchirkpx
lnfqdsuwjcteogambzuchirepx
lnfqdovwjnteigambzuchirkpx
lnfbdsvwjxteogambzuchirkax
lnfqdsawjyteogamxzuchiwkpx
lncqdsvwjoteogambzuchirfpx
lnfadsrwjyteogambzuchirktx
lnfqdsvhjyteorazbzuchirkpx
lwfqdsvwjytdogambzuchirkhx
lnfqdhvwjyteogambzuhairkpx
lnfqdsvwjytlogambzgchyrkpx
lnfqdsvwjyteogamnzjwhirkpx
lnfodsvwjyteogahuzuchirkpx
lnfqdsvtjyteogamvzwchirkpx
lnfqdsvwjzueogambzuxhirkpx
lnfqxsvljytxogambzuchirkpx
lnfqdsvwjyteogambpyvhirkpx
lqzqdsvwjnteogambzuchirkpx
lnfqdsvwjyteogumbzichirapx
lnfqdbvwjytedgaubzuchirkpx
lnfqdsvwpyteogabbhuchirkpx
nnfqdsvwryteogambzuchiwkpx
lrfqdsvwjeteogambzuchhrkpx
lnfqdsvwxyteogamtzucyirkpx
lnfsdsvwjyteogambzulhirknx
lnfqdsvwjyreoyambzmchirkpx
ltfqdsvwjytdogkmbzuchirkpx
lnfqwbvcjyteogambzuchirkpx
lnfqdsvwjyteogamrzfchirmpx
lnfqdsvwjqteogambzucwirkpy
lnfqdslwjyfeogambzuchikkpx
lnfqdsvwjybeogambzuchikjpx
lofqysvwjyteogasbzuchirkpx
lnfqusvwjyteogambzucverkpx
lnfqdsvwjyteogaibzuchfrkzx
lnfqdsvwjyleogabbzuchirkcx
lnfqdsvqjyteogambzuchdqkpx
lnfqdsvwjwtewgambzuciirkpx
lnfqisvwjatwogambzuchirkpx
lnfqdgvwjyteogambzuchizkix
lnfqdsxwjyteogambyuehirkpx
lpffdsvwjyteogamrzuchirkpx
lnirdsvwjyteogambzuchirkbx
lnfqdsxdjyteogazbzuchirkpx
lnfgdgvwyyteogambzuchirkpx
lnfqxsvwjyteogambzmcwirkpx
lnxqjsvwjyteogambzuchirkqx
lnrqdsvwjpteogkmbzuchirkpx
lnfqdstwjyteoeahbzuchirkpx
lnfqdsvwtuteogambzuchixkpx
lwfqvsvwjyteogambzughirkpx
lnkqdsvwjyfeogambzuuhirkpx
lvvqdsvwjyteogambzuchirkpn
jndqdsvwjyteogzmbzuchirkpx
enfqdszwjyteogambcuchirkpx
lnfqdsvwiyteogakbauchirkpx
lnfqdsvwjyteogtmbzxcdirkpx
fnfqdswwjyteogawbzuchirkpx
lnfqdsvwjydejqambzuchirkpx
lnqqdsvwjyteogambzucbdrkpx
lnfqdsvwjyteogadbzuchirxcx
lnfqdslwjyyeogambzulhirkpx
lnfqdsvwjytecgambzucgirkpb
lbmqdsvwjyteogamkzuchirkpx
lbfqdsvrjyteogambzuchirapx
lnfqdsmwjyteogambzucfarkpx
lnfqasvwoyteofambzuchirkpx
bnfudsvwjyteogambzucharkpx
lnfrdsvwjytxogambzuchirkpg
lbfqdsvwjyteagambzucdirkpx
lxfqdsvwjytuogambzucjirkpx
lnfqdsvwjcteogamyzuchiikpx
lnfodsvwjyteognmbzuchirapx
ltfqdsvwjytedgaxbzuchirkpx
lnfqdshwjyteogambzucsilkpx
lnfqdsvwpyteohambzuchitkpx
wnzqdsvwjyteogambiuchirkpx
lnfqdsvwayteogambzhchirkpw
ltfqcsvwjrteogambzuchirkpx
lnfqdsvwaytekgamvzuchirkpx
lnfqdsvwjyteogambzokpirkpx
lnfqysbwjyeeogambzuchirkpx
lnsqdsvwjyteogambzuchikkpd
lrfqdsvwjyteogahbzochirkpx
lnfqdsvwjyreoggmbzuchjrkpx
lxfqdsvwjyteogkmbzuchirkpp
enhqdbvwjyteogambzuchirkpx
jnfqdsvwjyteogamczuuhirkpx
lnfqdsvwuyteogadbzuchirkpw
lnfqdsvjjytergambznchirkpx
lnfqdsvwjyteoglmbzuceijkpx
lwfqdsvwjyteogamieuchirkpx
lnfqdsvwjfaeogambzqchirkpx
lfbqdjvwjyteogambzuchirkpx
lnfqdsvwjxteoaambzuchirkpp
lnfqdsvwjyheogjmbzgchirkpx
lnfqdskwjyteonambzuchiikpx
lnfqdwvwjyteogambxuchirkph
pnfqdsvwdyteogambzuchihkpx
lnoqdsvwjyteogaybznchirkpx
lnfqdsvwfyxefgambzuchirkpx
lnfqdsvwjyteotamxzmchirkpx
lnfqdsvwjyteigwmbzuchivkpx
lnfqdsvwjytekgambcuchirkwx
lnfqdsvwjuteogamrzulhirkpx
lnfqdsvwjyteogambzucczrgpx
wnfqzsvwjyteogambduchirkpx
lnfqdsowjyteogambuuthirkpx
lnfqdsvrjyteogcmbzuclirkpx
knfqdsvwgyteogambzuchorkpx
lnaqdsvwjytuogdmbzuchirkpx
lnfrdsvwjyteogambluchigkpx
lnfqqzvwjyteogambzkchirkpx
lnfqdsvwjyteogamuzuchgrkux
lnfqdsvnjyteogxmbznchirkpx
lnfqdsvwjyteolajbzuchdrkpx
lnfqdsvwjypeoiagbzuchirkpx
lnrqdsvwjyteozamuzuchirkpx
lnfqdsvwjytkogaubzucqirkpx
lnkbdsvwjyteogacbzuchirkpx
unfqdsvwjybeogambwuchirkpx
lnfqfsvzjyteogambzuchiikpx
lnfqdsvgjyreogahbzuchirkpx
lnfqdsewjyteogavbeuchirkpx
lnfqdsvwjdteogambbuchidkpx
lnfqdsvwjythogambzcchirkvx
lnfqdscwjyteorambzuchirgpx
cnfqdzvwjyteogambzushirkpx
lnfgdsgwjytedgambzuchirkpx
lnfqdsvwbyteogimbzuchdrkpx
lnfqdsswjyteogambzuohbrkpx
lnfqdsvwjytqogabbzachirkpx
lnfqdsvwjyteogmmbzucqiukpx
lnfxdsrwjyteogambzuchnrkpx
lnfqnqvwjyteogambzuchiwkpx
lffqisvwjyteogambzulhirkpx
lnfqdsxwjydeogambzucfirkpx
lnfqdsvwjyteogambzucjirkrp
lnfqdsnqjyteogambduchirkpx
fnfqdmvwjyteogamlzuchirkpx
lnfqvsvwjyteooamdzuchirkpx
lnfqdsvcyyteogambzuchickpx
onfqdsvwjyqeogambzuchirqpx
znfqdcvwjyteoaambzuchirkpx
lnfqdsvwjzteogambzuchidklx
lnfqjsvwjyteogjmbzuchirkpv
lnfqdsvwjytgorambzuchirppx
lzfqdsvwpfteogambzuchirkpx
lnfidsfwjyteogapbzuchirkpx
lnfodsvwbyteobambzuchirkpx
lnlqdsvwjytefgambzuchfrkpx
lnkqdsvwjyteogambzkchgrkpx
tnfqdsvwjyteoiamhzuchirkpx
lnfqdsvwjyteogamllschirkpx
lnfqdsvwjmthogamizuchirkpx
lnfqdbvwjyteogafbzuchirkpb
lnfxosvwjyteogahbzuchirkpx
lnmqdsvwjyzeogambzuchirkcx
lnfqdevbjytxogambzuchirkpx
lnfqdsvwjyteogamzzudhipkpx
lnfqdszwjyteoqambzuchirkpp
lffqdsvwjyteogamtouchirkpx
lnfqdsvhjytfogambzucharkpx
hnfqdsvwjyteogembzschirkpx
lnfqdsvwjateogambzuchirmpa
lnfqdsvcjyteogambzocairkpx
lnfqdsvwjyteogamwzmchirkpd
lnfqzsvwjyteogdmbzuyhirkpx
lnfqdsvwjytfyglmbzuchirkpx
lnfndsvwjyteogambzuchirktf
gnfqdnvwjytevgambzuchirkpx
lnfqdsvwjyteoganbpuchorkpx
lnfpdsvwnyteogambzucqirkpx
fnfqdstejyteogambzuchirkpx
lnfqlsvwjyteowambzuchirkmx
lnfqdsvwjyteogmmdzuchtrkpx
lnfqdsvwcyteogaqbzuchirkqx
lnfqdsvwjytlogtmbzuchiwkpx
lnfqdsvwoyteogambzuczirkwx
lnfqdsvwjyteogzybzucdirkpx
lnfqdvvwjyteogumbzuchiukpx
lnfqbwvwjyteogambzuchjrkpx
lnfgdsvwjyteogambzvchirkzx
lnfqdsvwjvtjogambzuchiokpx
lnfedsvwjyteogambzuchivkph
lhfqusvwjytaogambzuchirkpx
lnfqdsvwjyteogacbzuihirkpv
lnfwdsvwjyteogambzucokrkpx
lnfqtsvwjpteognmbzuchirkpx
anfqdswwjyteogambzucairkpx
lnfqdsvwjyteorambzuchirlsx
lnfqdsvwjytgogambzychirkpc
lnfqdhvwjyteogambzachirklx
lnfwdsvwjyteogaobquchirkpx
rnfqdsvwjiteogambzuhhirkpx
lnfqdsuwjyemogambzuchirkpx
hnfqdsvwjyteogambzuchprfpx
anfqssvwjyteogambzumhirkpx
lnfkdsvwjyteogafbzqchirkpx
lnfqdsvwjyteogacqzuchirspx
lnfqdskwjyteggambzuchiakpx
lnnqdsvwjyteooambzuchihkpx
lnlqdsvjjyteogambzuchgrkpx
lnfqdsvwjyteogamszochirkex
lnfqbsvwjyteogambzqchirepx
lnfqdsbwjcteogambzhchirkpx
lnfqdwvzjyteogambzechirkpx
ynfadsvwdyteogambzuchirkpx
tnfqdsvwjytuogambzuohirkpx
lnfqdsvwjyteogambzaohivkpx
mnfqisvwjyteogagbzuchirkpx
lnfqbsvwjyueogambzuchirkhx
ynfqdsvwjyteogdmbzuchinkpx
lnfqdwhwjyteogambzuchirqpx
mnfqdsvwjyteogambzfchkrkpx
lnfqdsnwjyteogambzgchiqkpx
lnfqdsvwjytergambzuchiuklx
lnfqdqvjjyteogamtzuchirkpx
lnfqdscwjyteorambzuchzrgpx
enfqdevwjyteogaabzuchirkpx
gnfqdsvbjyteogambzuchirkph
lnfqdxvwjyteogambzubhixkpx
lnfqdsvwjyteogambojchihkpx
lnfqdsvwjytdogambzuzhilkpx
lnfqdsvwjyteogamezuqhirtpx
tnfhdsvwjyteogambzuvhirkpx
lnfzdsvwjnteogahbzuchirkpx
lnfqdsvwjyteogambzfzhirkvx
lnfqqsvwjyteogambzuchirgpo
lufqpsvwjythogambzuchirkpx
lnfqdsvwjyteogzmbzuchimkix
lnwqdspwjyteogambzcchirkpx
lnfqdsowjyteogambzuchigypx
lnfqdnvvjyteogambzucjirkpx
lnfjdsvwryteogambzuchirkcx
lnfqdsvwbyteogambzuchirfpb
lnfqdsvwjyheogambzxchprkpx
lnfqmsvwjytezgambzuchirlpx
lnaqdsvwjyteogamdzuzhirkpx
lnoqdsvwjytebgambfuchirkpx
lnfqdtvwjytvogambzuchirkpv
`