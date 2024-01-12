package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := NewTeamFlyweightFactory()
	teamA1 := factory.GetTeam(TEAM_A)
	assert.NotNil(t, teamA1)

	teamA2 := factory.GetTeam(TEAM_A)
	assert.NotNil(t, teamA1)

	assert.Equal(t, teamA1, teamA2)

	assert.Equal(t, 1, factory.GetNumberOfObjects())
}
