package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type episode struct {
	nettohen   int
	production int
	broadcast  int
	viz        int
	name       string
	rjname     string
	jpname     string
	date       time.Time
}

var episodes []episode

func (e *episode) String() string {
	ret := ""
	if e.nettohen > 0 {
		ret += fmt.Sprintf("Nettohen Episode %d, Broadcast Episode %d, Viz Episode %d, Production Episode %d\n",
			e.nettohen, e.broadcast, e.viz, e.production)
		ret += fmt.Sprintf("English title: %s\n", e.name)
		ret += fmt.Sprintf("Japanese title: %s (%s)\n", e.jpname, e.rjname)
		ret += fmt.Sprintf("First aired %s", jpDate(e.date))
	} else {
		ret += fmt.Sprintf("Broadcast Episode %d, Production Episode %d\n",
			e.broadcast, e.production)
		ret += fmt.Sprintf("English title: %s\n", e.name)
		ret += fmt.Sprintf("Japanese title: %s (%s)\n", e.jpname, e.rjname)
		ret += fmt.Sprintf("First aired %s", jpDate(e.date))
	}
	return ret
}

func getDate(date string) (time.Time, error) {
	// Jan 2 15:04:05 2006 MST
	return time.Parse("January 2, 2006", date)
}

func jpDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func ogEpisode(broad, prod int, name, rjname, jpname, date string) episode {
	ret := episode{
		nettohen:   -1,
		production: prod,
		broadcast:  broad,
		viz:        broad,
		name:       name,
		rjname:     rjname,
		jpname:     jpname,
	}

	pdate, err := getDate(date)
	if err != nil {
		panic(err)
	}

	ret.date = pdate

	return ret
}

func nhEpisode(nh int, order, name, rjname, jpname, date string) episode {
	var prod, broad, viz int
	orders := strings.Split(order, "/")
	switch len(orders) {
	case 1:
		// Single number given - broadcast, viz, and production are the same.
		num, err := strconv.Atoi(orders[0])
		if err != nil {
			panic(err)
		}
		viz, prod, broad = num, num, num
	case 2:
		var err error
		// Two numbers given - broadcast and production disagree; viz is the same as production
		broad, err = strconv.Atoi(orders[0])
		if err != nil {
			panic(err)
		}
		prod, err = strconv.Atoi(orders[1])
		if err != nil {
			panic(err)
		}
		viz = prod
	case 3:
		var err error
		// Three numbers given (broadcast, viz, prod)
		broad, err = strconv.Atoi(orders[0])
		if err != nil {
			panic(err)
		}
		viz, err = strconv.Atoi(orders[1])
		if err != nil {
			panic(err)
		}
		prod, err = strconv.Atoi(orders[2])
		if err != nil {
			panic(err)
		}
	}
	ret := episode{
		nettohen:   nh,
		production: prod,
		broadcast:  broad,
		viz:        viz,
		name:       name,
		rjname:     rjname,
		jpname:     jpname,
	}

	pdate, err := getDate(date)
	if err != nil {
		panic(err)
	}

	ret.date = pdate

	return ret
}

func init() {
	// adapted from https://ranma.fandom.com/wiki/List_of_Ranma_%C2%BD_episodes
	// Using keyboard macros makes this ezpz, lol (thanks Emacs)
	episodes = make([]episode, 0, 161)
	episodes = append(episodes, ogEpisode(1, 1, "Here's Ranma", "Chūgoku kara Kita Aitsu! Chotto Hen!!", "中国からきたあいつ!ちょっとヘン!!", "April 15, 1989"))
	episodes = append(episodes, ogEpisode(2, 2, "School is No Place for Horsing Around", "Asobi Janai no Yo Gakkō wa", "遊びじゃないのよ学校は", "April 22, 1989"))
	episodes = append(episodes, ogEpisode(3, 3, "A Sudden Storm of Love", "Ikinari Ai no Arashi Chotto Matte Yo", "いきなり愛の嵐ちょっと待ってョ", "April 29, 1989"))
	episodes = append(episodes, ogEpisode(4, 4, "Ranma and... Ranma? If It's Not One Thing, It's Another", "Ranma to Ranma? Gokai ga Tomaranai", "乱馬とらんま?誤解がとまらない", "May 6, 1989"))
	episodes = append(episodes, ogEpisode(5, 5, "Love Me to the Bone! The Compound Fracture of Akane's Heart", "Kotsu Made Aishite? Akane Koi no Fukuzatsu Kossetsu", "骨まで愛して?あかね恋の複雑骨折", "May 13, 1989"))
	episodes = append(episodes, ogEpisode(6, 6, "Akane's Lost Love... These Things Happen, You Know", "Akane no Shitsuren Datte Shōganai Janai", "あかねの失恋だってしょうがないじゃない", "May 20, 1989"))
	episodes = append(episodes, ogEpisode(7, 7, "Enter Ryoga! The Eternal \"Lost Boy\"", "Tōjō! Eien no Mayoigo - Ryōga", "登場!永遠の迷い子·良牙", "May 27, 1989"))
	episodes = append(episodes, ogEpisode(8, 8, "School is a Battlefield! Ranma vs. Ryoga", "Gakkō wa Senjō da! Taiketsu Ranma Buiesu Ryōga", "学校は戦場だ!対決 乱馬VS良牙", "June 3, 1989"))
	episodes = append(episodes, ogEpisode(9, 9, "True Confessions! A Girl's Hair is Her Life!", "Otome Hakusho - Kami wa Onna no Inochi Nano", "乙女白書·髪は女のいのちなの", "June 17, 1989"))
	episodes = append(episodes, ogEpisode(10, 10, "P-P-P-Chan! He's Good For Nothin'", "Pi-pi-P-chan Roku Namonjanē", "ピーピーPちゃん ろくなもんじゃねェ", "July 1, 1989"))
	episodes = append(episodes, ogEpisode(11, 11, "Ranma Meets Love Head-On! Enter the Delinquent Juvenile Gymnast!", "Ranma wo Gekiai! Shintaisō no Sukeban Tōjō", "乱馬を激愛!新体操のスケバン登場", "July 15, 1989"))
	episodes = append(episodes, ogEpisode(12, 12, "A Woman's Love is War! The Martial Arts Rhythmic Gymnastics Challenge!", "Onna no Koi wa Sensō yo! Kakutō Shintaisō de Iza Shōbu", "女の恋は戦争よ!格闘新体操でいざ勝負", "July 22, 1989"))
	episodes = append(episodes, ogEpisode(13, 13, "A Tear in a Girl-Delinquent's Eye? The End of the Martial Arts Rhythmic Gymnastics Challenge!", "Sukeban no Me ni Namida? Rūru Muyō no Kakutō Shintaisō Kecchaku", "スケバンの目に涙?ルール無用の格闘新体操決着", "July 27, 1989"))
	episodes = append(episodes, ogEpisode(14, 17, "Pelvic Fortune-Telling? Ranma is the No. One Bride in Japan", "Kotsuban Uranai! Ranma wa Nippon-ichi no Oyomesan", "骨盤占い!らんまは日本一のお嫁さん", "August 19, 1989"))
	episodes = append(episodes, ogEpisode(15, 18, "Enter Shampoo, the Gung-Ho Girl! I Put My Life in Your Hands", "Gekiretsu Shōjo Shanpū Tōjō! Watashi Inochi Azukemasu", "激烈少女シャンプー登場!ワタシ命あずけます", "August 26, 1989"))
	episodes = append(episodes, ogEpisode(16, 19, "Shampoo's Revenge! The Shiatsu Technique That Steals Heart and Soul", "Shanpū no Hangeki! Hissatsu Shiatsu Kobushi wa Mukuro mo Kokoro mo Ubau", "シャンプーの反撃!必殺指圧拳は身も心も奪う", "September 2, 1989"))
	episodes = append(episodes, ogEpisode(17, 20, "I Love You, Ranma! Please Don't Say Goodbye", "Ranma Daisuki! Sayonara wa Iwanaide!!", "乱馬大好き!さよならはいわないで!!", "September 9, 1989"))
	episodes = append(episodes, ogEpisode(18, 21, "I Am a Man! Ranma's Going Back to China!?", "Ore wa Otoko da! Ranma Chūgoku e Kaeru?", "オレは男だ!らんま中国へ帰る?", "September 16, 1989"))
	episodes = append(episodes, nhEpisode(1, "19/22", "Clash of the Delivery Girls! The Martial Arts Takeout Race", "Gekitotsu! Demae Kakutō Rēsu", "激突!出前格闘レース", "October 20, 1989"))
	episodes = append(episodes, nhEpisode(2, "20/23", "You Really Do Hate Cats!", "Yappari Neko ga Kirai?", "やっぱり猫が嫌い?", "November 3, 1989"))
	episodes = append(episodes, nhEpisode(3, "21/24", "This Ol' Gal's the Leader of the Amazon Tribe!", "Watashi ga Joketsuzoku no Obaba!", "私が女傑族のおばば!", "November 10, 1989"))
	episodes = append(episodes, nhEpisode(4, "22/25", "Behold! The \"Chestnuts Roasting on an Open Fire\" Technique", "Deta! Hissatsu Tenshin Amaguriken", "出た!必殺天津甘栗拳!!", "November 17, 1989"))
	episodes = append(episodes, nhEpisode(5, "23/26", "Enter Mousse! The Fist of the White Swan", "Hakushō no Otoko Mūsu Tōjō!", "白鳥拳の男ムース登場!", "November 24, 1989"))
	episodes = append(episodes, nhEpisode(6, "24/27", "Cool Runnings! The Race of the Snowmen", "Bakusō! Yukidaruma Hakobi Rēsu", "爆走!雪だるま運びレース", "December 1, 1989"))
	episodes = append(episodes, nhEpisode(7, "25/19/14", "The Abduction of P-Chan", "Sarawareta P-chan", "さらわれたPちゃん!", "December 8, 1989"))
	episodes = append(episodes, nhEpisode(8, "26/20/15", "Close Call! The Dance of Death... On Ice!", "Kikiippatsu! Shiryō no Bonodori", "危機一髪!死霊の盆踊り", "December 15, 1989"))
	episodes = append(episodes, nhEpisode(9, "27/21/16", "P-Chan Explodes! The Icy Fountain of Love!", "P-chan Bakuhatsu! Ai no Mizubashira", "Pちゃん爆発!愛の水柱", "December 22, 1989"))
	episodes = append(episodes, nhEpisode(10, "28/29", "Ranma Trains on Mt. Terror", "Ranma Kyōfu no Yama Gomori", "乱馬恐怖の山ごもり", "January 12, 1990"))
	episodes = append(episodes, nhEpisode(11, "29/30", "The Breaking Point!? Ryoga's Great Revenge", "Bakusai Tenketsu to wa? Ryōga Daigyakushū", "爆砕点穴とは?良牙大逆襲", "January 19, 1990"))
	episodes = append(episodes, nhEpisode(12, "30/28", "Danger at the Tendo Dojo!", "Ayoushi! Tendō Dōjō", "危うし!天道道場", "January 26, 1990"))
	episodes = append(episodes, nhEpisode(13, "31", "The Abduction of Akane!", "Sarawareta Akane!", "さらわれたあかね!", "February 2, 1990"))
	episodes = append(episodes, nhEpisode(14, "32", "Ranma vs. Mousse! To Lose Is To Win", "Taiketsu Mūsu! Makeru ga Kachi", "対決ムース!負けるが勝ち", "February 9, 1990"))
	episodes = append(episodes, nhEpisode(15, "33", "Enter Happosai, the Lustful Lecher!", "Kyūkyoku no Ero Yōkai Happōsai", "究極のエロ妖怪八宝斉", "February 16, 1990"))
	episodes = append(episodes, nhEpisode(16, "34", "Assault on the Girls' Locker Room", "Joshi Kōishitsu wo Osoe?", "女子更衣室を襲え?", "February 23, 1990"))
	episodes = append(episodes, nhEpisode(17, "35", "Kuno's House of Gadgets! Guests Check In, But They Don't Check Out", "Oni mo Nigedasu Karakuri Yashiki", "鬼も逃げだすカラクリ屋敷", "March 2, 1990"))
	episodes = append(episodes, nhEpisode(18, "36", "Goodbye Girl-Type", "Kore de Onna to Osaraba?", "これで女とおさらば?", "March 9, 1990"))
	episodes = append(episodes, nhEpisode(19, "37", "It's a Fine Line Between Pleasure and Pain", "Ai to Nikushimi no Okurimono", "愛と憎しみの贈物", "March 16, 1990"))
	episodes = append(episodes, nhEpisode(20, "38", "S.O.S.! The Wrath of Happosai", "SOS Ero Yōkai Happōsai", "SOSエロ妖怪八宝斉", "March 23, 1990"))
	episodes = append(episodes, nhEpisode(21, "39", "Kissing is Such Sweet Sorrow! The Taking of Akane's Lips", "Akane no Kuchibiru wo Ubae", "あかねの口びるを奪え", "April 6, 1990"))
	episodes = append(episodes, nhEpisode(22, "40", "Bathhouse Battle! We're in Some Hot Water Now", "Ii Yu da na? Sentou de Sentou", "いい湯だな?銭湯で戦闘", "April 13, 1990"))
	episodes = append(episodes, nhEpisode(23, "41", "Ranma Gains Yet Another Suitor", "Mata Hitori Ranma wo Aishita Yatsu", "また一人乱馬を愛したヤツ", "April 20, 1990"))
	episodes = append(episodes, nhEpisode(24, "42", "Ryoga & Akane: 2-Gether, 4-Ever", "Netsuai? Ryōga to Akane", "熱愛?良牙とあかね", "April 27, 1990"))
	episodes = append(episodes, nhEpisode(25, "43", "Sneeze Me, Squeeze Me, Please Me! Shampoo's Recipe For Disaster", "Kushami Ippatsu Aishite Naito", "くしゃみ一発愛してナイト", "May 4, 1990"))
	episodes = append(episodes, nhEpisode(26, "44", "Rub-a-Dub-Dub! There's a Pervert in the Tub", "Maboroshi no Happōdaikarin wo Sagase", "幻の八宝大華輪を探せ", "May 11, 1990"))
	episodes = append(episodes, nhEpisode(27, "45", "I Love You! My Dear, Dear Ukyo", "Daisuki! Watashi no Ucchan", "大好き!私のうっちゃん", "May 18, 1990"))
	episodes = append(episodes, nhEpisode(28, "46", "The Witch Who Loved Me: A Japanese Ghost Story", "Majo ga Aishita Shitagi Dorobō", "魔女が愛した下着ドロボー", "May 25, 1990"))
	episodes = append(episodes, nhEpisode(29, "47", "Transform! Akane the Super-Duper Girl", "Henshin! Mukimuki-man Akane", "変身!ムキムキマンあかね", "June 1, 1990"))
	episodes = append(episodes, nhEpisode(30, "48", "The Killer From Jusenkyo", "Jusenkyō kara Kita Koroshiya", "呪泉郷から来た殺し屋", "June 8, 1990"))
	episodes = append(episodes, nhEpisode(31, "49", "Am I... Pretty? Ranma's Declaration of Womanhood", "Watashi Kirei? Ranma Onna Sengen", "私ってきれい?乱馬女宣言", "June 15, 1990"))
	episodes = append(episodes, nhEpisode(32, "50", "Final Facedown! Happosai vs. The Invisible Man", "Taiketsu! Happōsai Buiesu Tōmeiningen", "対決!八宝斉VS透明人間", "June 22, 1990"))
	episodes = append(episodes, nhEpisode(33, "51", "Les Misérables of the Kuno Estate", "Kunōke no Re Miseraburu", "九能家のレ·ミゼラブル", "June 29, 1990"))
	episodes = append(episodes, nhEpisode(34, "52", "Ghost Story! Ranma and the Magic Sword", "Kaidan! Ranma to Mashō no Ken", "怪談!乱馬と魔性の剣", "July 6, 1990"))
	episodes = append(episodes, nhEpisode(35, "53", "All It Takes is One! The Kiss of Love is the Kiss of Death", "Hitotsubu Korori - Zetsurin Hore Gusuri", "一粒コロリ·絶倫ホレ薬", "July 13, 1990"))
	episodes = append(episodes, nhEpisode(36, "54", "The Ultimate Team-up!? The Ryoga/Mousse Alliance", "Shijō Saikyō? Ryōga to Mūsu Dōmei", "史上最強?良牙とムース同盟", "July 20, 1990"))
	episodes = append(episodes, nhEpisode(37, "55", "Back to the Happosai!", "Bakku Tu Za Happōsai", "バック·トゥ·ザ·八宝斉", "July 27, 1990"))
	episodes = append(episodes, nhEpisode(38, "56", "Kodachi the Black Rose! The Beeline to True Love", "Kurobara no Kodachi! Jun'ai Icchokusen", "黒バラの小太刀!純愛一直線", "August 3, 1990"))
	episodes = append(episodes, nhEpisode(39, "57", "The Last Days of Happosai...?", "Happōsai Saigo no Hi?", "八宝斉 最期の日?", "August 10, 1990"))
	episodes = append(episodes, nhEpisode(40, "58", "Two, Too Violent Girls: Ling-Ling & Lung-Lung", "Abarenbō Musume Rinrin Ranran", "暴れん坊娘リンリンランラン", "August 17, 1990"))
	episodes = append(episodes, nhEpisode(41, "59", "Ranma and the Evil Within", "Ranma wo Osō Kyōfu no Tatari", "乱馬を襲う恐怖のタタリ", "August 24, 1990"))
	episodes = append(episodes, nhEpisode(42, "60", "Enter Ken and His Copycat Kerchief", "Toujou! Monomane Kakutōgi", "登場!ものまね格闘技", "August 31, 1990"))
	episodes = append(episodes, nhEpisode(43, "61", "Ryoga's Miracle Cure! Hand Over That Soap", "Ryōga no Taishitsu Kaizen Sekken!", "良牙の体質改善セッケン!", "September 7, 1990"))
	episodes = append(episodes, nhEpisode(44, "62", "Fight! The Anything-Goes Obstacle Course Race", "Kakutō! Shōgaibutsu Rēsu", "格闘!障害物レース", "September 14, 1990"))
	episodes = append(episodes, nhEpisode(45, "63/64", "Ranma Goes Back to Jusenkyo at Last", "Ranma, Tsuini Jusenkyō e Iku", "乱馬, ついに呪泉郷へ行く", "September 21, 1990"))
	episodes = append(episodes, nhEpisode(46, "64/65", "The Return of the Hawaiʻian Headmaster from Hell", "Kaettekita Hentai Kōchō", "帰ってきた変態校長", "October 5, 1990"))
	episodes = append(episodes, nhEpisode(47, "65/66", "Enter Kuno, the Night-Prowling Knight", "Tōjō! Shijō Saikyō ni Kunō", "登場!史上最強の九能", "October 12, 1990"))
	episodes = append(episodes, nhEpisode(48, "66/67", "Ranma Gets Weak!", "Ranma ga Yowaku Nacchatta!", "乱馬が弱くなっちゃった!", "October 19, 1990"))
	episodes = append(episodes, nhEpisode(49, "67/68", "Eureka! The Desperate Move of Desperation", "Kansei! Tondemonai Hissatsuwaza", "完成!とんでもない必殺技", "October 26, 1990"))
	episodes = append(episodes, nhEpisode(50, "68/69", "Showdown! Can Ranma Make a Comeback?", "Kessen! Ranma Fukkatsu Naru ka?", "決戦!乱馬復活なるか?", "November 2, 1990"))
	episodes = append(episodes, nhEpisode(51, "69/63", "Ukyo's Skirt! The Great Girly-Girl Gambit", "Ukyō no Sukāto Daisakusen", "右京のスカート大作戦!", "November 9, 1990"))
	episodes = append(episodes, nhEpisode(52, "70", "Here Comes Ranma's Mom!", "Ranma no Mama ga Yattekita!", "乱馬のママがやってきた!", "November 16, 1990"))
	episodes = append(episodes, nhEpisode(53, "71", "From Ryoga with Love", "Ryōga, Ai to Kunō wo Koete", "良牙, 愛と苦悩を越えて", "November 23, 1990"))
	episodes = append(episodes, nhEpisode(54, "72", "My Fiancé, the Cat", "Fianse wa Bakeneko", "フィアンセは化け猫", "November 30, 1990"))
	episodes = append(episodes, nhEpisode(55, "73", "Blow, Wind! To Be Young is to Go Gung-Ho", "Fukeyo Kaze! Seishun wa Nekketsuda", "吹けよ風!青春は熱血だ", "December 7, 1990"))
	episodes = append(episodes, nhEpisode(56, "74", "A Formidable New Disciple Appears", "Osorubeki Shindeshi Arawaru", "恐るべき新弟子現わる", "December 14, 1990"))
	episodes = append(episodes, nhEpisode(57, "75", "Step Outside!", "Omote ni Deyagare!", "おもてに出やがれ!", "December 21, 1990"))
	episodes = append(episodes, nhEpisode(58, "76", "Ryoga's \"Tendo Dojo Houseguest\" Diary", "Ryōga no Tendō Dōjō Isōrō Nikki", "良牙の天道道場居候日記", "January 11, 1991"))
	episodes = append(episodes, nhEpisode(59, "77", "Happosai's Happy Heart!", "Happōsai no Koi!", "八宝斉の恋!", "January 18, 1991"))
	episodes = append(episodes, nhEpisode(60, "78", "Extra, Extra! Kuno & Nabiki: Read All About It!", "Kunō Bōzen! Koi no Daiyogen", "九能ボー然!恋の大予言", "January 25, 1991"))
	episodes = append(episodes, nhEpisode(61, "79", "Ryoga the Strong... Too Strong", "Tsuyoku Narisugita Ryōga", "強くなりすぎた良牙", "February 1, 1991"))
	episodes = append(episodes, nhEpisode(62, "80", "Close Call! P-chan's Secret", "Ayaushi! P-chan no Himitsu", "あやうし!Pちゃんの秘密", "February 8, 1991"))
	episodes = append(episodes, nhEpisode(63, "81", "The Egg-Catcher Man", "Tamago wo Tsukamu Otoko", "たまごをつかむ男", "February 15, 1991"))
	episodes = append(episodes, nhEpisode(64, "82", "Ranma and Kuno's... First Kiss", "Ranma to Kunō no Hatsu Kisu?!", "らんまと九能の初キッス?!", "February 22, 1991"))
	episodes = append(episodes, nhEpisode(65, "83", "Shampoo's Red Thread of Dread!", "Shanpū no Akai Ito", "シャンプーの赤い糸", "March 1, 1991"))
	episodes = append(episodes, nhEpisode(66, "84", "Mousse Goes Home to the Country!", "Mūsu Kokyō ni Kaeru", "ムース故郷に帰る", "March 8, 1991"))
	episodes = append(episodes, nhEpisode(67, "85", "The Dumbest Bet in History!", "Shijō Saite no Kake", "史上サイテーの賭け", "March 15, 1991"))
	episodes = append(episodes, nhEpisode(68, "86", "Kuno Becomes a Marianne!", "Mariannu ni Natta Kunō", "マリアンヌになった九能", "March 22, 1991"))
	episodes = append(episodes, nhEpisode(69, "87", "Ranma, You Are Such A Jerk!", "Ranma Nanka Daikirai!", "乱馬なんか大キライ!", "March 29, 1991"))
	episodes = append(episodes, nhEpisode(70, "88/90", "Gimme That Pigtail", "Sono Osage Moratta!", "そのおさげもらったぁ!", "April 5, 1991"))
	episodes = append(episodes, nhEpisode(71, "89/88", "When a Guy's Pride and Joy is Gone", "Otoko no Yabō ga Tsukiru Toki...", "男の野望が尽きる時...", "April 12, 1991"))
	episodes = append(episodes, nhEpisode(72, "90/89", "Ling-Ling & Lung-Lung Strike Back!", "Rinrin Ranran no Gyakushū", "リンリン·ランランの逆襲", "April 19, 1991"))
	episodes = append(episodes, nhEpisode(73, "91", "Ryoga's Proposal", "Ryōga no Puropōsu", "良牙のプロポーズ", "April 26, 1991"))
	episodes = append(episodes, nhEpisode(74, "92", "Genma Takes a Walk", "Genma, Iede Suru", "玄馬, 家出する", "May 3, 1991"))
	episodes = append(episodes, nhEpisode(75, "93", "The Gentle Art of Martial Tea Ceremony", "Kore ga Kakutō Sadō de Omasu", "これが格闘茶道でおます", "May 10, 1991"))
	episodes = append(episodes, nhEpisode(76, "94", "And the Challenger is... A Girl?!", "Dōjō Yaburi wa Onna no Ko?", "道場破りは女の子?", "May 17, 1991"))
	episodes = append(episodes, nhEpisode(77, "95", "Hot Springs Battle Royale!", "Zekkyō! Onsen Batoru", "絶叫!温泉バトル", "May 24, 1991"))
	episodes = append(episodes, nhEpisode(78, "96", "Me is Kuno's Daddy, Me is", "Mī ga Kunō no Dadi Desu", "ミーが九能のダディです", "May 31, 1991"))
	episodes = append(episodes, nhEpisode(79, "97", "The Matriarch Takes a Stand", "Kakutō Sadō Iemoto Tatsu!", "格闘茶道·家元立つ!", "June 7, 1991"))
	episodes = append(episodes, nhEpisode(80, "98", "A Leotard is a Girl's Burden", "Reotādo wa Otome no Noroi", "レオタードは乙女の呪い", "June 14, 1991"))
	episodes = append(episodes, nhEpisode(81, "99", "The Mixed-Bath Horror!", "Kyōfu no Kon'yoku Onsen", "恐怖の混浴温泉", "June 21, 1991"))
	episodes = append(episodes, nhEpisode(82, "100", "The Frogman's Curse!", "Kaeru no Urami Harashimasu", "カエルのうらみはらします", "June 28, 1991"))
	episodes = append(episodes, nhEpisode(83, "101", "Revenge! Raging Okonomiyaki...!", "Gyakushū! Ikari no Okonomiyaki", "逆襲!怒りのお好み焼き", "July 5, 1991"))
	episodes = append(episodes, nhEpisode(84, "102", "Ranma the Lady-Killer", "Nanpa ni Natta Ranma", "ナンパになった乱馬", "July 12, 1991"))
	episodes = append(episodes, nhEpisode(85, "103", "Shogi Showdown", "Kakutō Shogi wa Inochi Gake", "格闘将棋は命懸け", "July 19, 1991"))
	episodes = append(episodes, nhEpisode(86, "104", "Sasuke's \"Mission: Improbable\"", "Sasuke no Supai Daisakusen", "佐助のスパイ大作戦", "July 26, 1991"))
	episodes = append(episodes, nhEpisode(87, "105", "Bonjour, Furinkan!", "Bonjūru de Gozaimasu", "ボンジュールでございます", "August 2, 1991"))
	episodes = append(episodes, nhEpisode(88, "106", "Dinner at Ringside!", "Dinā wa Ringu no Uede", "ディナーはリングの上で", "August 9, 1991"))
	episodes = append(episodes, nhEpisode(89, "107", "Swimming with Psychos", "Akane, Namida no Suiei Daitokkun", "あかね, 涙の水泳大特訓", "August 16, 1991"))
	episodes = append(episodes, nhEpisode(90, "108", "Ryoga, Run Into the Sunset", "Ryōga! Yūhi ni Mukatte Hashire", "良牙!夕日に向かって走れ", "August 23, 1991"))
	episodes = append(episodes, nhEpisode(91, "109", "Into the Darkness", "Yume no Naka e", "夢の中へ", "August 30, 1991"))
	episodes = append(episodes, nhEpisode(92, "110", "Nabiki, Ranma's New Fiancée!", "Ranma wa Nabiki no Iinazuke?", "乱馬はなびきの許婚?", "September 6, 1991"))
	episodes = append(episodes, nhEpisode(93, "111", "Case of the Missing Takoyaki!", "Tendo-ke Kieta Takoyaki no Nazo", "天道家消えたたこ焼きの謎", "September 13, 1991"))
	episodes = append(episodes, nhEpisode(94, "112", "Ranma Versus Shadow Ranma!", "Taiketsu! Ranma Buiesu Kage Ranma", "対決!乱馬VS影乱馬", "September 20, 1991"))
	episodes = append(episodes, nhEpisode(95, "113", "Dear Daddy... Love, Kodachi!", "Kodachi no Mai Raburī Papa", "小太刀のマイラブリーパパ", "September 27, 1991"))
	episodes = append(episodes, nhEpisode(96, "114", "Enter Gosunkugi, The New Rival!?", "Kyōteki? Gosunkugi-kun Tōjō", "強敵?五寸釘くん登場", "October 4, 1991"))
	episodes = append(episodes, nhEpisode(97, "115", "Ranma's Calligraphy Challenge", "Ranma wa Hetakuso? Kakutō Shodō", "乱馬はヘタクソ?格闘書道", "October 11, 1991"))
	episodes = append(episodes, nhEpisode(98, "116", "The Secret Don of Furinkan High", "Furinkan Kōkō, Kage no Don Tōjō", "風林館高校, 影のドン登場", "October 18, 1991"))
	episodes = append(episodes, nhEpisode(99, "117", "Back to the Way We Were... Please!", "Higan! Futsū no Otoko ni Modoritai", "悲願!普通の男に戻りたい", "October 25, 1991"))
	episodes = append(episodes, nhEpisode(100, "118", "Ryoga Inherits the Saotome School?", "Saotome Ryū no Atotsugi wa Ryōga?", "早乙女流の跡継ぎは良牙?", "November 1, 1991"))
	episodes = append(episodes, nhEpisode(101, "119", "Tendo Family Goes to the Amusement Park!", "Tendō-ke, Yūenchi e Iku", "天道家, 遊園地へ行く", "November 15, 1991"))
	episodes = append(episodes, nhEpisode(102, "120", "The Case of the Furinkan Stalker!", "Furinkan Kōkō: Toorima Jiken", "風林館高校·通り魔事件", "November 29, 1991"))
	episodes = append(episodes, nhEpisode(103, "121/122", "The Demon from Jusenkyo, Part I", "Jusenkyō Kara Kita Akuma - Zenpen", "呪泉郷から来た悪魔 前編", "December 6, 1991"))
	episodes = append(episodes, nhEpisode(104, "122/123", "The Demon from Jusenkyo, Part II", "Jusenkyō Kara Kita Akuma - Kōhen", "呪泉郷から来た悪魔 後編", "December 13, 1991"))
	episodes = append(episodes, nhEpisode(105, "123/125", "A Xmas Without Ranma", "Ranma ga Inai Xmas", "乱馬がいないXmas", "December 20, 1991"))
	episodes = append(episodes, nhEpisode(106, "124/126", "A Cold Day in Furinkan", "Yukinko Fuyu Monogatari", "雪ん子冬物語", "January 10, 1992"))
	episodes = append(episodes, nhEpisode(107, "125/128", "Curse of the Scribbled Panda", "Rakugaki Panda no Noroi", "らくがきパンダの呪い", "January 17, 1992"))
	episodes = append(episodes, nhEpisode(108, "126/121", "The Date-Monster of Watermelon Island", "Suikatō no Kōsaiki", "スイカ島の交際鬼", "January 24, 1992"))
	episodes = append(episodes, nhEpisode(109, "127/129", "Legend of the Lucky Panda!", "Shiawase no Panda Densetsu", "幸せのパンダ伝説", "January 31, 1992"))
	episodes = append(episodes, nhEpisode(110, "128/131", "Ukyo's Secret Sauce, Part 1", "Ranma to Ukyo ga Sōsu Sōai?", "乱馬と右京がソース相愛?", "February 7, 1992"))
	episodes = append(episodes, nhEpisode(111, "129/132", "Ukyo's Secret Sauce, Part 2", "Itsuwari Fūfu yo Eien ni...", "偽り夫婦よ永遠に...", "February 14, 1992"))
	episodes = append(episodes, nhEpisode(112, "130/124", "The Missing Matriarch of Martial Arts Tea!", "Kakutō Sadō! Sarawareta Iemoto", "格闘茶道!さらわれた家元", "February 21, 1992"))
	episodes = append(episodes, nhEpisode(113, "131/127", "Akane Goes to the Hospital!", "Taihen! Akane ga Nyūin Shita", "大変!あかねが入院した", "February 28, 1992"))
	episodes = append(episodes, nhEpisode(114, "132/130", "Mystery of the Marauding Octopus Pot!", "Nazono Abare Takotsubo Arawareru?!", "謎の暴れタコツボ現る?!", "March 6, 1992"))
	episodes = append(episodes, nhEpisode(115, "133/134", "Gosunkugi's Paper Dolls of Love", "Gosunkugi! Ah Koi no Kaminingyō", "五寸釘!あぁ恋の紙人形", "March 13, 1992"))
	episodes = append(episodes, nhEpisode(116, "134/135", "Akane's Unfathomable Heart", "Akane no Kokoro ga Wakaranai", "あかねの心がわからない", "March 20, 1992"))
	episodes = append(episodes, nhEpisode(117, "135/133", "A Teenage Ghost Story", "Tsuiseki! Temari Uta no Nazo", "追跡!手まり唄の謎", "March 27, 1992"))
	episodes = append(episodes, nhEpisode(118, "136", "Master and Student... Forever!?", "Mou Anata kara Hanarenai", "もうあなたから離れない", "April 3, 1992"))
	episodes = append(episodes, nhEpisode(119, "137", "Tatewaki Kuno, Substitute Principal", "Kunō Tatewaki, Dairi Kōchō wo Meizu", "九能帯刀, 代理校長を命ず", "April 10, 1992"))
	episodes = append(episodes, nhEpisode(120, "138", "Ranma's Greatest Challenge!", "Ranma, Tsukiyo ni Hoeru", "乱馬, 月夜に吠える", "April 17, 1992"))
	episodes = append(episodes, nhEpisode(121, "139", "Nihao! Jusenkyo Guide", "Nihao! Jusenkyō no Gaido-san", "你好(ニーハオ)!呪泉郷のガイドさん", "April 24, 1992"))
	episodes = append(episodes, nhEpisode(122, "140", "Pick-a-Peck o' Happosai", "Meiwaku! Rokunin no Happōsai", "迷惑!六人の八宝斉", "May 1, 1992"))
	episodes = append(episodes, nhEpisode(123, "141", "From the Depth of Despair, Part I", "Kibun Shidai no Hissatsuwaza - Zen", "気分しだいの必殺技(前)", "May 8, 1992"))
	episodes = append(episodes, nhEpisode(124, "142", "From the Depth of Despair, Part II", "Kibun Shidai no Hissatsuwaza - Kō", "気分しだいの必殺技(後)", "May 15, 1992"))
	episodes = append(episodes, nhEpisode(125, "143", "Shampoo's Curséd Kiss", "Shanpū Toraware no Kissu", "シャンプー囚われのキッス", "May 22, 1992"))
	episodes = append(episodes, nhEpisode(126, "144", "Run Away with Me, Ranma", "Boku to Kakeochi Shite kudasai", "ボクと駆け落ちして下さい", "May 29, 1992"))
	episodes = append(episodes, nhEpisode(127, "145", "Let's Go to the Mushroom Temple", "Kinoko Dera e Ikō", "キノコ寺へ行こう", "June 5, 1992"))
	episodes = append(episodes, nhEpisode(128, "146", "The Cradle from Hell", "Hissatsu! Jigoku no Yurikago", "必殺!地獄のゆりかご", "June 12, 1992"))
	episodes = append(episodes, nhEpisode(129, "147", "Madame St. Paul's Cry for Help", "Aoi Kyōfu ni Bonjūru", "青い恐怖にボンジュール", "June 19, 1992"))
	episodes = append(episodes, nhEpisode(130, "148", "Meet You in the Milky Way", "Orihime wa Nagareboshi ni Notte", "織姫は流れ星に乗って", "June 26, 1992"))
	episodes = append(episodes, nhEpisode(131, "149", "Wretched Rice Cakes Of Love", "Hitotsu Meshimase Koi no Sakuramochi", "一つ召しませ恋の桜餅", "July 3, 1992"))
	episodes = append(episodes, nhEpisode(132, "150", "The Horrible Happo Mold-Burst!", "Dekita! Happō Dai Kabin", "できた!八宝大カビン", "July 10, 1992"))
	episodes = append(episodes, nhEpisode(133, "151", "The Kuno Sibling Scandal", "Kunō Kyōdai Sukyandaru no Arashi", "九能兄妹スキャンダルの嵐", "July 17, 1992"))
	episodes = append(episodes, nhEpisode(134, "152", "Battle for the Golden Tea Set", "Ougon no Chaki, Gojōnotō no Kessen", "黄金の茶器, 五重塔の決戦", "July 24, 1992"))
	episodes = append(episodes, nhEpisode(135, "153", "Gosunkugi's Summer Affair!", "Gosunkugi Hikaru, Hito Natsu no Koi", "五寸釘光, ひと夏の恋", "July 31, 1992"))
	episodes = append(episodes, nhEpisode(136, "154/155", "Bring It On! Love as a Cheerleader - Part 1", "Ai no Kakutō Chiagāru - Zen", "愛の格闘チアガール(前)", "August 7, 1992"))
	episodes = append(episodes, nhEpisode(137, "155/156", "Bring It On! Love as a Cheerleader - Part 2", "Ai no Kakutō Chiagāru - Kō", "愛の格闘チアガール(後)", "August 14, 1992"))
	episodes = append(episodes, nhEpisode(138, "156/154", "The Battle for Miss Beachside", "Kettei! Misu Bīchisaido", "決定!ミス·ビーチサイド", "August 21, 1992"))
	episodes = append(episodes, nhEpisode(139, "157", "The Musical Instruments of Destruction", "Bakuretsu! Haipā Tsuzumi", "爆裂!ハイパーツヅミ", "August 28, 1992"))
	episodes = append(episodes, nhEpisode(140, "158", "A Ninja's Dog is Black and White", "Shinobi no Inu wa Shiro to Kuro", "忍の犬は白と黒", "September 4, 1992"))
	episodes = append(episodes, nhEpisode(141, "159", "The Tendo Dragon Legend", "Tendō-ke: Ryūjin Densetsu", "天道家·龍神伝説", "September 11, 1992"))
	episodes = append(episodes, nhEpisode(142, "160", "Boy Meets Mom Part 1", "Ranma, Mītsu Mazā", "乱馬, ミーツ·マザー", "September 18, 1992"))
	episodes = append(episodes, nhEpisode(143, "161", "Boy Meets Mom Part 2 Someday, Somehow...", "Itsu no Hi ka, Kitto...", "いつの日か, きっと...", "September 25, 1992"))
}

func findEpisode(matcher func(episode) bool) (*episode, error) {
	for _, epi := range episodes {
		if matcher(epi) {
			return &epi, nil
		}
	}
	return nil, errors.New("cannot find matching episode")
}

func usage() {
	fmt.Printf("%s: Ranma ½ episode search utility\n\nUsage: %s <COMMAND> [ARG]...\n", os.Args[0], os.Args[0])
	fmt.Println("\nCommands are:")
	fmt.Println("\tnettohen\t(Alias \"nh\") Find episode by Nettohen number.")
	fmt.Println("\tbroadcast\t(Alias \"bc\") Find episode by broadcast order.")
	fmt.Println("\tproduction\t(Alias \"prod\") Find episode by production order.")
	fmt.Println("\tviz\tFind episode by Viz home release order.")
	fmt.Println("\tname\tFind episode by English name (fuzzy find)")
	fmt.Println("\trjname\tFind episode by Japanese (romaji) name (fuzzy find)")
	fmt.Println("\tepisodes\t List episodes as tab-separated data")
	fmt.Println("\thelp\t(Alias \"usage\") Display this message.")
}

func requiresArgs() {
	if len(os.Args) < 2 {
		fmt.Println("This command requires at least one argument")
		os.Exit(-1)
	}
}

func fuzzy(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r
	}
	if r > 0x0100 && r < 0x0170 {
		if r == 0x0101 {
			return 'a'
		} else if r == 0x0113 {
			return 'e'
		} else if r == 0x012B {
			return 'i'
		} else if r == 0x014D {
			return 'o'
		} else if r == 0x016B {
			return 'u'
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(-1)
	}

	switch os.Args[1] {
	case "help", "usage":
		usage()
	case "rjname":
		requiresArgs()
		arg := strings.Join(os.Args[2:], " ")
		fuzz := strings.Map(fuzzy, strings.ToLower(arg))

		epi, err := findEpisode(func(ep episode) bool {
			nfuzz := strings.Map(fuzzy, strings.ToLower(ep.rjname))
			return nfuzz == fuzz
		})

		if err != nil {
			fmt.Printf("Can't find episode romaji name \"%s\"\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	case "name":
		requiresArgs()
		arg := strings.Join(os.Args[2:], " ")
		fuzz := strings.Map(fuzzy, strings.ToLower(arg))

		epi, err := findEpisode(func(ep episode) bool {
			nfuzz := strings.Map(fuzzy, strings.ToLower(ep.name))
			return nfuzz == fuzz
		})

		if err != nil {
			fmt.Printf("Can't find episode name \"%s\"\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	case "prod", "production":
		requiresArgs()

		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Bad argument: %v\n", err)
			os.Exit(-1)
		}

		epi, err := findEpisode(func(ep episode) bool {
			return ep.production == arg
		})

		if err != nil {
			fmt.Printf("Can't find production episode %d\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	case "viz":
		requiresArgs()

		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Bad argument: %v\n", err)
			os.Exit(-1)
		}

		epi, err := findEpisode(func(ep episode) bool {
			return ep.viz == arg
		})

		if err != nil {
			fmt.Printf("Can't find Viz episode %d\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	case "bc", "broadcast":
		requiresArgs()

		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Bad argument: %v\n", err)
			os.Exit(-1)
		}

		epi, err := findEpisode(func(ep episode) bool {
			return ep.broadcast == arg
		})

		if err != nil {
			fmt.Printf("Can't find broadcast episode %d\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	case "episodes":
		fmt.Println("Nettohen No.\tBroadcast No.\tViz No.\tProduction No.\tEN Title\tJP Title (romaji)\tJP Title\tBroadcast Date (YYYY-MM-DD)")
		for _, epi := range episodes {
			fmt.Printf("%d\t%d\t%d\t%d\t%s\t%s\t%s\t%s\n",
				epi.nettohen,
				epi.broadcast,
				epi.viz,
				epi.production,
				epi.name,
				epi.rjname,
				epi.jpname,
				jpDate(epi.date))
		}
	case "nh", "nettohen":
		requiresArgs()

		arg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Bad argument: %v\n", err)
			os.Exit(-1)
		}

		epi, err := findEpisode(func(ep episode) bool {
			return ep.nettohen == arg
		})

		if err != nil {
			fmt.Printf("Can't find Nettohen episode %d\n", arg)
			os.Exit(-1)
		}

		fmt.Println(epi)
	default:
		fmt.Printf("Unknown command %s\n", os.Args[1])
		os.Exit(-1)
	}
}
