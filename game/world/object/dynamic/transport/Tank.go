package transport

type Tank struct {
	id string
	x int
	y int
	hp int
	ammo int
	vector int
}

func NewTank(id string, x int, y int) *Tank {
	return &Tank{
		id:   id,
		x:    x,
		y:    y,
		hp:   100,
		ammo: 5,
		vector: 0, //TODO remove
	}
}

func (t *Tank) GetId() string {
	return t.id
}

func (t *Tank) Coordinates() (int, int) {
	return t.x, t.y
}

func (t *Tank) Move(available func(int, int) bool, scale int) {
	nextX, nextY := t.getNextMoveCoordinates()

	isXOutOfRange := nextX > 0 && nextX < scale
	isYOutOfRange := nextY > 0 && nextY < scale

	if isXOutOfRange || isYOutOfRange {
		return
	}

	if !available(nextX, nextY) {
		return
	}

	t.x = nextX
	t.y = nextY
}

func (t *Tank) getNextMoveCoordinates() (int, int) {
	var dx, dy int

	switch t.vector {
	case 0:
		dy = -1
	case 1:
		dx = 1
	case 2:
		dy = 1
	case 3:
		dx = -1
	}

	return t.x + dx, t.y + dy
}

func (t *Tank) GetHP() int {
	return t.hp
}

func (t *Tank) EditHP(v int) {
	t.hp += v
}

func (t *Tank) GetAmmo() int {
	return t.hp
}

func (t *Tank) EditAmmo(v int) {
	t.hp += v
}
