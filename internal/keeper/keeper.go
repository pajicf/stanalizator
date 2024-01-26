package keeper

import "fmt"

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
}
