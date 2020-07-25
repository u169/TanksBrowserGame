package transport

type Tank struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	HP     int `json:"hp"`
	Ammo   int `json:"ammo"`
	Vector int `json:"vector"`
	shoot  bool
}

func NewTank(x int, y int) *Tank {
	return &Tank{
		X:      x,
		Y:      y,
		HP:     100,
		Ammo:   5,
		Vector: 0,
		shoot:  false,
	}
}

func (t *Tank) Coordinates() (int, int) {
	return t.X, t.Y
}

func (t *Tank) Move(isBusied func(int, int) bool, scale int) {
	nextX, nextY := t.getNextMoveCoordinates()

	isXOutOfRange := (nextX < 0) || (nextX >= scale)
	isYOutOfRange := (nextY < 0) || (nextY >= scale)

	if isXOutOfRange || isYOutOfRange || isBusied(nextX, nextY) {
		return
	}

	t.X = nextX
	t.Y = nextY
}

func (t *Tank) getNextMoveCoordinates() (int, int) {
	var dx, dy int

	switch t.Vector {
	case 0:
		dy = -1
	case 1:
		dx = 1
	case 2:
		dy = 1
	case 3:
		dx = -1
	}

	return t.X + dx, t.Y + dy
}

func (t *Tank) Rotate(vector int) {
	t.Vector = vector
}

func (t *Tank) SetShoot(b bool) {
	t.shoot = b
}

func (t *Tank) GetShoot() bool {
	return t.shoot
}

func (t *Tank) GetHP() int {
	return t.HP
}

func (t *Tank) EditHP(v int) {
	t.HP += v
}

func (t *Tank) GetAmmo() int {
	return t.HP
}

func (t *Tank) EditAmmo(v int) {
	t.HP += v
}
