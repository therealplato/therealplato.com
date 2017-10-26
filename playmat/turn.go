func turn() {
beginning:
	{
		untap()
		upkeep()
		draw()
	}
main1:
	{
		if !landThisTurn {
			playLand()
		}
		castSpells()
	}
combat:
	{
		begin()
		declareAttackers()
		declareBlockers()
		combatDamage()
		end()
	}
main2:
	{
		if !landThisTurn {
			playLand()
		}
		castSpells()
	}

end:
	{
		endStep()
		cleanupDmg()
	}
}
