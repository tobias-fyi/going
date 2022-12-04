package main

// === ===
// === 2022 Advent of Code, Day 2 === //
// === ===

/*
--- Day 3: Rucksack Reorganization ---
One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey.
Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need
to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly
one of the two compartments. The Elf that did the packing failed to follow this rule for exactly
one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but
they need your help finding the errors. Every item type is identified by a single lowercase or
uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack
always has the same number of items in each of its two compartments, so the first half of the
characters represent items in the first compartment, while the second half of the characters
represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

		vJrwpWtwJgWrhcsFMMfFFhFp
		jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
		PmmdzqPrVvPwwTWBwg
		wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
		ttgJtRGJQctTZtZT
		CrZsJsPPZsGzwwsLwLmpwMDw

* The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment
contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp.
The only item type that appears in both compartments is lowercase p.
* The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item
type that appears in both compartments is uppercase L.
* The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is
uppercase P.
* The fourth rucksack's compartments only share item type v.
* The fifth rucksack's compartments only share item type t.
* The sixth rucksack's compartments only share item type s.

To help prioritize item rearrangement, every item type can be converted to a priority:

* Lowercase item types a through z have priorities 1 through 26.
* Uppercase item types A through Z have priorities 27 through 52.

In the above example, the priority of the item type that appears in both compartments of each
rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?
*/

// === Part 1 === //

// func main() {
// 	score, err := scanFileAndCalcSum("input3.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(score)
// }

// func scanFileAndCalcSum(file string) (int, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return 0, fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	var sum int
// 	prio := makePrioMap("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() {
// 		l := sc.Text()
// 		// Split line in half
// 		halves := chunks(l, len(l)/2)
// 		log.Println(halves)
// 		// For each of the items in the first, check if that one exists in second
// 	out:
// 		for _, itemOne := range halves[0] {
// 			for _, itemTwo := range halves[1] {
// 				if itemOne == itemTwo {
// 					// When match is found, add that priority to sum
// 					log.Println(string(itemOne))
// 					log.Println(prio[string(itemOne)])
// 					sum += prio[string(itemOne)]
// 					break out
// 				}
// 			}
// 		}
// 	}
// 	if err := sc.Err(); err != nil {
// 		return 0, fmt.Errorf("scan file error: %w", err)
// 	}

// 	return sum, nil
// }

// func makePrioMap(s string) map[string]int {
// 	p := map[string]int{}
// 	for i, c := range s {
// 		p[string(c)] = i + 1
// 	}
// 	return p
// }

// func chunks(s string, chunkSize int) []string {
// 	if len(s) == 0 {
// 		return nil
// 	}
// 	if chunkSize >= len(s) {
// 		return []string{s}
// 	}
// 	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
// 	currentLen := 0
// 	currentStart := 0
// 	for i := range s {
// 		if currentLen == chunkSize {
// 			chunks = append(chunks, s[currentStart:i])
// 			currentLen = 0
// 			currentStart = i
// 		}
// 		currentLen++
// 	}
// 	chunks = append(chunks, s[currentStart:])
// 	return chunks
// }

// === Part 2 === //

/*
--- Part Two ---
As you finish identifying the misplaced items, the Elves come to you with another issue.

For safety, the Elves are divided into groups of three. Every Elf carries a badge that identifies
their group. For efficiency, within each group of three Elves, the badge is the only item type
carried by all three Elves. That is, if a group's badge is item type B, then all three Elves will
have item type B somewhere in their rucksack, and at most two of the Elves will be carrying any
other item type.

The problem is that someone forgot to put this year's updated authenticity sticker on the badges.
All of the badges need to be pulled out of the rucksacks so the new authenticity stickers can be
attached.

Additionally, nobody wrote down which item type corresponds to each group's badges. The only way to
tell which item type is the right one is by finding the one item type that is common between all
three Elves in each group.

Every set of three lines in your list corresponds to a single group, but each group can have a
different badge item type. So, in the above example, the first group's rucksacks are the first
three lines:

		vJrwpWtwJgWrhcsFMMfFFhFp
		jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
		PmmdzqPrVvPwwTWBwg

And the second group's rucksacks are the next three lines:

		wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
		ttgJtRGJQctTZtZT
		CrZsJsPPZsGzwwsLwLmpwMDw

In the first group, the only item type that appears in all three rucksacks is lowercase r; this
must be their badges. In the second group, their badge item type must be Z.

Priorities for these items must still be found to organize the sticker attachment efforts: here,
they are 18 (r) for the first group and 52 (Z) for the second group. The sum of these is 70.

Find the item type that corresponds to the badges of each three-Elf group. What is the sum of the
priorities of those item types?
*/

// func main() {
// 	score, err := calcSum("input3.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(score)
// }

// func calcSum(file string) (int, error) {
// 	var sum int
// 	prio := makePrioMap("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 	lines, err := scanFileLinesIntoSlice(file)
// 	if err != nil {
// 		return 0, err
// 	}

// 	for i := 0; i < len(lines); i += 3 {
// 		group := lines[i : i+3]
// 	out:
// 		for _, item1 := range group[0] {
// 			for _, item2 := range group[1] {
// 				for _, item3 := range group[2] {
// 					if item1 == item2 && item2 == item3 {
// 						i1 := string(item1)
// 						// log.Println(i1)
// 						// log.Println(prio[i1])
// 						sum += prio[i1]
// 						break out
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return sum, nil
// }

// func scanFileLinesIntoSlice(file string) ([]string, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	lines := []string{}

// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() {
// 		lines = append(lines, sc.Text())
// 	}
// 	if err := sc.Err(); err != nil {
// 		return nil, fmt.Errorf("scan file error: %w", err)
// 	}

// 	return lines, nil
// }

// func makePrioMap(s string) map[string]int {
// 	p := map[string]int{}
// 	for i, c := range s {
// 		p[string(c)] = i + 1
// 	}
// 	return p
// }

// func chunks(s string, chunkSize int) []string {
// 	if len(s) == 0 {
// 		return nil
// 	}
// 	if chunkSize >= len(s) {
// 		return []string{s}
// 	}
// 	var chunks []string = make([]string, 0, (len(s)-1)/chunkSize+1)
// 	currentLen := 0
// 	currentStart := 0
// 	for i := range s {
// 		if currentLen == chunkSize {
// 			chunks = append(chunks, s[currentStart:i])
// 			currentLen = 0
// 			currentStart = i
// 		}
// 		currentLen++
// 	}
// 	chunks = append(chunks, s[currentStart:])
// 	return chunks
// }
