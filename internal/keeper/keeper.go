package keeper

type Keeper struct {
	cfg *Config
}

func NewKeeper(cfg *Config) Keeper {
	return Keeper{
		cfg: cfg,
	}
}

func (k *Keeper) Run() {
	println("Running the keeper, yay!")
}
