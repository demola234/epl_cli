packa

func (m model) View() string {
    // Sort the teams slice based on points, descending
    sort.SliceStable(m.teams, func(i, j int) bool {
        return m.teams[i].points > m.teams[j].points
    })

    s := "Premier League Table\n\n"
    for i, team := range m.teams {
        s += fmt.Sprintf("%d. %s, %d pts\n", i+1, team.name, team.points)
    }
    return s
}