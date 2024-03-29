package keeper

import (
	"fmt"
	"github.com/pajicf/stanalizator/internal/services/emailjs"
)

type Keeper struct {
	cfg Config
}

func NewKeeper(cfg Config) Keeper {
	return Keeper{
		cfg,
	}
}

func (k *Keeper) Run() {
	fmt.Println("Running the keeper, yay!")
	fmt.Printf("Build config test: %s \n", k.cfg.EmailJSServiceID)
	fmt.Printf("Args config test: %s \n", k.cfg.ToMail)
	fmt.Printf("Geo Config item 0, type: %d & type api value %s \n", k.cfg.GeoConfig[0].Type, k.cfg.GeoConfig[0].Type.ApiValue())

	_ = emailjs.NewEmailJS(k.cfg.EmailJSUserID, k.cfg.EmailJSServiceID)
}
