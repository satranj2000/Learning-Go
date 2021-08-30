package maps

func MapsBasics() {

	// initialize a map using make. Key is string and value is int
	districtSchools := make(map[string]int)
	districtSchools["Bangalore Urban"] = 1000
	districtSchools["Bangalore Rural"] = 800
	districtSchools["Mysuru"] = 810
	districtSchools["Shivamogga"] = 200

	println(districtSchools)

	//
}
