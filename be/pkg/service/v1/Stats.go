package v1

import "strconv"

type Stats struct {
	StraightFlushes int
	FourOfAKinds int
	FullHouses int
	Flushes int
	Straights int
	ThreeOfAKinds int
	TwoPairs int
	Pairs int
	HighCards int
	Player1Wins int
	Player2Wins int
	TotalRounds int
}

func (s *Stats) StraightFlush() {
	s.StraightFlushes++
}

func (s *Stats) FourOfAKind() {
	s.FourOfAKinds++
}

func (s *Stats) FullHouse() {
	s.FullHouses++
}

func (s *Stats) Flush() {
	s.Flushes++
}

func (s *Stats) Straight() {
	s.Straights++
}
func (s *Stats) ThreeOfAKind() {
	s.ThreeOfAKinds++
}

func (s *Stats) TwoPair() {
	s.TwoPairs++
}

func (s *Stats) Pair() {
	s.Pairs++
}

func (s *Stats) HighCard() {
	s.HighCards++
}

func (s *Stats) Player1Win() {
	s.Player1Wins++
}

func (s *Stats) Player2Win() {
	s.Player2Wins++
}

func (s *Stats) RoundInc() {
	s.TotalRounds++
}

func (s *Stats) GetData() string {
	return "SUCCESS!\n----------------------------------\n" +
	"Statistics\n" +
	"----------------------------------\n" + s.DataForDB() + s.MessageForDB()
}

func (s *Stats) DataForDB() string {
	return "StraightFlushes: " + strconv.Itoa(s.Straights) + "\n" +
		"FourOfAKinds: " + strconv.Itoa(s.FourOfAKinds) + "\n" +
		"FullHouses: " + strconv.Itoa(s.FullHouses) + "\n" +
		"Flushes: " + strconv.Itoa(s.Flushes) + "\n" +
		"Straights: " + strconv.Itoa(s.Straights) + "\n" +
		"ThreeOfAKinds: " + strconv.Itoa(s.ThreeOfAKinds) + "\n" +
		"TwoPairs: " + strconv.Itoa(s.TwoPairs) + "\n" +
		"Pairs: " + strconv.Itoa(s.Pairs) + "\n" +
		"HighCards: " + strconv.Itoa(s.HighCards) + "\n" +
		"Player1Wins: " + strconv.Itoa(s.Player1Wins) + "\n" +
		"Player2Wins: " + strconv.Itoa(s.Player2Wins) + "\n" +
		"TotalRounds: " + strconv.Itoa(s.TotalRounds) + "\n"
}

func (s *Stats) MessageForDB() string {
	return "--------------------------------------------------------------\n" +
		"To view data in db, you can access it with the following CLI: \n" +
		"mysql -h bgtusc51sgcqzmjlcisq-mysql.services.clever-cloud.com -P 3306 -u uj8arcdthnwjt1sj -p bgtusc51sgcqzmjlcisq\n" +
		"The password is 'sLvaummcX1e7viu4kiLo'\n"+
		"Then type 'use bgtusc51sgcqzmjlcisq' \n" +
		"Then 'select * from malta_be'.\n" +
		"--------------------------------------------------------------\n"
}

func (s *Stats) Reset() {
	s.StraightFlushes = 0
	s.FourOfAKinds = 0
	s.FullHouses = 0
	s.Flushes = 0
	s.Straights = 0
	s.ThreeOfAKinds = 0
	s.TwoPairs = 0
	s.Pairs = 0
	s.HighCards = 0
	s.Player1Wins = 0
	s.Player2Wins = 0
	s.TotalRounds = 0
}



