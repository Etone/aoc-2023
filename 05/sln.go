package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type seed = int
type location = int

var seedSoil map[intRange]intRange
var soilFertilizer map[intRange]intRange
var fertilizerWater map[intRange]intRange
var waterLight map[intRange]intRange
var lightTemperatur map[intRange]intRange
var temperaturHumidity map[intRange]intRange
var humidityLocation map[intRange]intRange

var seedsindividual []seed

var seedRange []intRange

var minLocationPart2 int

var wg sync.WaitGroup

type intRange struct {
	start, end int
}

func readFileToBlocks(path string) (blocks []string) {
	inputByte, _ := os.ReadFile(path)
	input := string(inputByte)
	blocks = strings.Split(input, "\n\n")
	return
}

func parseSeedsIdividual(seedLine string) {
	seedsText := strings.Split(strings.TrimPrefix(seedLine, "seeds: "), " ")

	for _, seedText := range seedsText {
		seed, _ := strconv.Atoi(seedText)
		seedsindividual = append(seedsindividual, seed)
	}
}

func parseSeedsRange(seedLine string) {
	seedsText := strings.Split(strings.TrimPrefix(seedLine, "seeds: "), " ")
	for len(seedsText) > 0 {
		pair := seedsText[:2]
		seedRangeStart, _ := strconv.Atoi(pair[0])
		seedRangeSteps, _ := strconv.Atoi(pair[1])
		seedsText = seedsText[2:]
		seedRange = append(seedRange, intRange{start: seedRangeStart, end: seedRangeStart + seedRangeSteps})
	}
}

func parseMap(mapText string) (parsed map[intRange]intRange) {
	parsed = make(map[intRange]intRange)
	entries := strings.Split(mapText, "\n")[1:]
	for _, entry := range entries {
		values := strings.Split(entry, " ")
		destinationStart, _ := strconv.Atoi(values[0])
		sourceStart, _ := strconv.Atoi(values[1])
		steps, _ := strconv.Atoi(values[2])

		parsed[intRange{start: sourceStart, end: sourceStart + steps}] = intRange{start: destinationStart, end: destinationStart + steps}

	}
	return
}

func initializeMaps(blocks []string) {
	minLocationPart2 = math.MaxInt32

	seedSoil = parseMap(blocks[1])
	soilFertilizer = parseMap(blocks[2])
	fertilizerWater = parseMap(blocks[3])
	waterLight = parseMap(blocks[4])
	lightTemperatur = parseMap(blocks[5])
	temperaturHumidity = parseMap(blocks[6])
	humidityLocation = parseMap(blocks[7])
}

func getWithFallback(lookup map[intRange]intRange, key int) int {
	keys := make([]intRange, 0, len(lookup))
	for ir := range lookup {
		keys = append(keys, ir)
	}
	for _, mapKey := range keys {
		if key >= mapKey.start && key <= mapKey.end {
			difference := key - mapKey.start
			return lookup[mapKey].start + difference
		}
	}
	return key
}

func getLocationForSeed(seed seed) (location location) {
	soil := getWithFallback(seedSoil, seed)
	fertilizer := getWithFallback(soilFertilizer, soil)
	water := getWithFallback(fertilizerWater, fertilizer)
	light := getWithFallback(waterLight, water)
	temperatur := getWithFallback(lightTemperatur, light)
	humidity := getWithFallback(temperaturHumidity, temperatur)
	location = getWithFallback(humidityLocation, humidity)

	return
}

func getLocationForSeedAsync(seed chan intRange) {

	r := <- seed

	for i := r.start; i<= r.end; i++ {
		soil := getWithFallback(seedSoil, i)
		fertilizer := getWithFallback(soilFertilizer, soil)
		water := getWithFallback(fertilizerWater, fertilizer)
		light := getWithFallback(waterLight, water)
		temperatur := getWithFallback(lightTemperatur, light)
		humidity := getWithFallback(temperaturHumidity, temperatur)
		location := getWithFallback(humidityLocation, humidity)

		minLocationPart2 = min(minLocationPart2, location)
	}

	defer wg.Done()
}

func main() {
	part1()
	part2()
}

func part1() {
	blocks := readFileToBlocks("input_01.txt")
	parseSeedsIdividual(blocks[0])
	initializeMaps(blocks)

	seedLocations := []int{}

	for i := 0; i < len(seedsindividual); i++ {
		seedLocation := getLocationForSeed(seedsindividual[i])
		seedLocations = append(seedLocations, seedLocation)
	}
	fmt.Printf("Part 1: %d\n", slices.Min(seedLocations))
}

func part2() {
	blocks := readFileToBlocks("input_01.txt")
	parseSeedsRange(blocks[0])
	initializeMaps(blocks)

	for _, sr := range seedRange {
		channel := make(chan intRange)
		wg.Add(1)
		go getLocationForSeedAsync(channel)
		channel <- sr
	}

	wg.Wait()
	fmt.Printf("Part 2: %d\n", minLocationPart2)
}
