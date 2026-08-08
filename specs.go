package main

type carStatCore struct {
	sizeWieght dimensions
	engine powerplant
	grip suspension
	aero aerodynamics
	gear drivetrain
	specs performance
}

type performance struct {
	mechTopSpeed struct
	aeroTopSpeed struct
	COG struct
	rollCenter struct
	susFrequency struct
}

type rulesets struct {
	STL carStatCore
}

type checkRules struct {
	myBuild CarStatCore
	myClass rulesets
}

type dimensions struct {
	weight float32
	length float32
	height float32
	width float32
	fTrack float32
	rTrack float32
	weightDistro float32
	fAxleWeight float32
	rAxleWeight float32
}

type powerplant struct {
	displacement float32
	bore float32
	stroke float32
	peakRPM float32
	BHP horsePower
	BTQ torque
	WHP horsePower
	WTQ torque
	aspiration str
}

type horsePower struct {
	pwr float32
	RPM int
}

type torque struct {
	tq float32
	RPM int
}

type suspension struct {
	fStyle str
	rStyle str
	fWheels wheel
	fTires tire
	rWheels wheel
	rTires tire
	ttsa float32
}

type wheel struct {
	width float32
	diameter float32
}

type tire struct {
	width int
	innerDia int
	outterDia float32
}

type aerodynamics struct {
	coDrag float32
	frontalArea float32
	cda float32
	fLift float32
	rLift float32
}


type drivetrain struct {
	gearset ratios
	finalDrive float32
}

type ratios struct {
	first float32
	second float32
	third float32
	fourth float32
	fifth float32
	sixth float32
	seventh float32
	eighth float32
	ninth float32
	tenth float32
}

type brakes struct {
	rotors int
	masterCylDia int
	pedalRatio int
	fPistons calipers
	rPistons calipers
}

type calipers struct {
	pistonCount int
	pistonDiameter int
}

type rotors struct {
	diameter int
	width int
}

func brakeBias() int {
	// axle weights
	// master cyl dia
	// disc diameter
	// piston count
	// pedal angle
}

func speedByGear() struct {
	// ratio, fd, tire dia
}

func aeroTopSpeed() struct {
	// weight, power, cda
}