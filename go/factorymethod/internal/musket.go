package gun

type musket struct {
	gun
}

func newMusket() Gun {
	return &musket{
		gun: gun{
			name: "Musket",
			power: 1,
		},
	}
}
