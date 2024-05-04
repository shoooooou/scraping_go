package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
	// "github.com/gocolly/colly/v2"
)

type articleInfo struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func saveArticlesJson(fName string, a []articleInfo) {
	// Create json file
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	// Dump json to the standard output
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(a)
	if err != nil {
		log.Fatal(err)
	}

	// Struct to json
	b, _ := json.MarshalIndent(a, "", "  ")
	fmt.Println(string(b))
	// fmt.Println(p)
}

func main() {
	// Target URL
	url := "https://tabelog.com/rstLst/?vs=1&sa=&sk=%25E5%2596%25AB%25E8%258C%25B6%25E5%25BA%2597&lid=hd_search1&vac_net=&svd=20240504&svt=1900&svps=2&hfc=1&sunday=&LstRange=&LstReserve=&ChkCard=&ChkCardType=&ChkRoom=&ChkSemiRoom=&ChkRoomType=&ChkCharter=&ChkCharterType=&LstSmoking=&ChkBunen=&ChkParking=&ChkVegKodawari=&ChkFishKodawari=&ChkHealthy=&ChkVegetarianMenu=&ChkSake=&ChkSakeKodawari=&ChkShochu=&ChkShochuKodawari=&ChkWine=&ChkWineKodawari=&ChkCocktail=&ChkCocktailKodawari=&ChkNomihoudai=&ChkOver180minNomihoudai=&ChkNomihoudaiOnly=&ChkTabehoudai=&ChkFineView=&ChkNightView=&ChkOceanView=&ChkHotel=&ChkKakurega=&ChkHouse=&ChkStylish=&ChkRelax=&ChkWideSeat=&ChkCoupleSeat=&ChkCounter=&ChkSofa=&ChkZashiki=&ChkHorikotatsu=&ChkTerrace=&ChkKaraoke=&ChkDarts=&ChkLive=&ChkSports=&ChkOver150minParty=&ChkCelebrate=&ChkBirthdayPrivilege=&ChkCarryOnDrink=&ChkSommelier=&ChkPet=&ChkTakeout=&ChkDelivery=&ChkHappyHour=&ChkPremiumCoupon=&ChkOnlineBooking=&freecall=&ChkNewOpen=&award_prize=&chk_hyakumeiten_genres=&ChkTpointGive=&ChkTpointUse=&ChkEnglishMenu=&ChkMorningMenu=&ChkAllergyLabeling=&ChkCalorieLabeling=&ChkSweetsTabehoudai=&ChkEMoneyPayment=&ChkKoutsuuIcPayment=&ChkRakutenEdyPayment=&ChkNanacoPayment=&ChkWaonPayment=&ChkIdPayment=&ChkQuicPayPayment=&ChkQrcodePayment=&ChkQrcodePaymentType=&ChkKids=&ChkBabies=&ChkPreschoolChild=&ChkElementarySchoolStudent=&ChkKidsMenu=&ChkBabycar=&ChkTachiNomi=&ChkProjector=&ChkPowerSupply=&ChkWheelchair=&ChkFreeWifi=&ChkPaidWifi=&ChkSeatOnly=&ChkCoordinatorReward=&LstCos=&LstCosT=&RdoCosTp=2&LstSitu=&LstRev=&ChkCoupon=&Cat=SC&LstCat=SC10&LstCatD=SC1002&cat_sk=%E5%96%AB%E8%8C%B6%E5%BA%97"

	// Instantiate default collector
	c := colly.NewCollector()

	i := 0
	// Extract li class="new-entry-item"
	c.OnHTML(".list-rst__rst-data", func(e *colly.HTMLElement) {
		i++
		shopName := e.ChildText(".list-rst__rst-name-target")
		address := e.ChildText(".list-rst__area-genre")
		fmt.Println("id:", i)
		fmt.Println("shopName:", shopName)
		fmt.Println("address:", address)
	})

	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})

	// After making a request extract status code
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("StatusCode:", r.StatusCode)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("error:", r.StatusCode, err)
	})

	// Start scraping on https://XXX
	c.Visit(url)

	// Wait until threads are finished
	c.Wait()

	// Save as JSON format
	// saveArticlesJson("articles.json", articles)
}
