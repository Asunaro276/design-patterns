package gun

type musket struct {
	gun
}

func newMusket() IGun {
	return &musket{
		gun: gun{
			name: "Musket",
			power: 1,
		},
	}
}
