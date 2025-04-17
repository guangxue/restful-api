package hongfa

import (
	"fmt"
	"net/http"

	"github.com/guangxue/restful-api/internal/utils"
)

type Item struct {
	ItemName  string `json:"item_name"`
	QtyShelf  int    `json:"qty_shelf"`
	QtyTagged int    `json:"qty_tagged"`
	QtyLoose  int    `json:"qty_loose"`
	TagsLeft  int    `json:"tags_left"`
}

func Hongfa(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hongfa-IMS Api is under development\n\nStay tuned ! Coming Soon !\n")
}

func Items(w http.ResponseWriter, r *http.Request) {
	items := []Item{
		{
			ItemName:  "013810-284BEIGE",
			QtyShelf:  2000,
			QtyTagged: 1000,
			QtyLoose:  500,
			TagsLeft:  500,
		},
		{
			ItemName:  "013810-294MUSTAN",
			QtyShelf:  2000,
			QtyTagged: 1000,
			QtyLoose:  0,
			TagsLeft:  0,
		},
		{
			ItemName:  "013810-318DRKBLU",
			QtyShelf:  2000,
			QtyTagged: 1000,
			QtyLoose:  0,
			TagsLeft:  0,
		},
	}
	utils.JsonResponse(w, items)
}
