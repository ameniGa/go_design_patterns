package flyweight

import "time"

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

func getTeamFactory(teamID int) Team {
	switch teamID {
	case TEAM_A:
		return Team{ID: 2, Name: "team A"}
	default:
		return Team{ID: 3, Name: "team B"}
	}
}

const (
	TEAM_A = iota
	TEAM_B
)

type teamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func NewTeamFlyweightFactory() teamFlyweightFactory {
	return teamFlyweightFactory{createdTeams: make(map[int]*Team)}
}

func (t *teamFlyweightFactory) GetTeam(teamID int) *Team {
	createdTeam, ok := t.createdTeams[teamID]
	if ok {
		return createdTeam
	}

	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team

	return &team
}

func (t *teamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
