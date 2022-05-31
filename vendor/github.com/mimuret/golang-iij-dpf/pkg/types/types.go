package types

type Boolean int

const (
	Enabled  Boolean = 1
	Disabled Boolean = 0
)

func (c Boolean) String() string {
	booleanToString := map[Boolean]string{
		Enabled:  "Enabled",
		Disabled: "Disabled",
	}
	return booleanToString[c]
}

type State int

const (
	StateBeforeStart State = 1
	StateRunning     State = 2
)

func (c State) String() string {
	stateToString := map[State]string{
		StateBeforeStart: "BeforeStart",
		StateRunning:     "Started",
	}

	return stateToString[c]
}

type Favorite int

const (
	FavoriteHighPriority Favorite = 1
	FavoriteLowPriority  Favorite = 2
)

func (c Favorite) String() string {
	favoriteToString := map[Favorite]string{
		FavoriteHighPriority: "High",
		FavoriteLowPriority:  "Low",
	}
	return favoriteToString[c]
}
