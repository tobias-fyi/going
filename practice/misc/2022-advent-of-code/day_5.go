package main

// === ===
// === 2022 Advent of Code, Day 5 === //
// === ===

/*
--- Day 5: Supply Stacks ---
The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.
The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.
The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.
They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

				[D]
		[N] [C]
		[Z] [M] [P]
		1   2   3

		move 1 from 2 to 1
		move 3 from 1 to 3
		move 2 from 2 to 1
		move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.
Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

		[D]
		[N] [C]
		[Z] [M] [P]
		1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

						[Z]
						[N]
				[C] [D]
				[M] [P]
		1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

						[Z]
						[N]
		[M]     [D]
		[C]     [P]
		1   2   3

 Finally, one crate is moved from stack 1 to stack 2:

						[Z]
						[N]
						[D]
		[C] [M] [P]
		1   2   3

The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.
After the rearrangement procedure completes, what crate ends up on top of each stack?
*/

// === Part 1 === //

// type Stack struct {
// 	data []string
// }

// func (s *Stack) pop() string {
// 	if len(s.data) == 0 {
// 		return ""
// 	}
// 	l := len(s.data)
// 	val := s.data[l-1]
// 	s.data[l-1] = ""
// 	s.data = s.data[:l-1]
// 	return val
// }

// func (s *Stack) push(val string) {
// 	s.data = append(s.data, val)
// }

// type CargoArea struct {
// 	stacks []*Stack
// }

// func (c *CargoArea) newEmptyStack() int {
// 	c.stacks = append(c.stacks, &Stack{})
// 	return len(c.stacks)
// }

// func (c *CargoArea) addStackFromSlice(sl []string) {
// 	c.stacks = append(c.stacks, &Stack{data: sl})
// }

// func (c *CargoArea) executeInstruction(instr []int) error {
// 	if instr[1] > len(c.stacks) || instr[2] > len(c.stacks) {
// 		return fmt.Errorf("executeInstruction: invalid instruction: stack number does not exist")
// 	}

// 	for i := 0; i < instr[0]; i++ {
// 		val := c.stacks[instr[1]-1].pop()
// 		c.stacks[instr[2]-1].push(val)
// 	}

// 	return nil
// }

// func (c *CargoArea) ReturnTopCrates() string {
// 	b := strings.Builder{}

// 	for _, st := range c.stacks {
// 		b.WriteString(st.data[len(st.data)-1])
// 	}

// 	return b.String()
// }

// func (c *CargoArea) isEmpty() bool {
// 	for _, s := range c.stacks {
// 		if len(s.data) > 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (c CargoArea) String() string {
// 	if c.isEmpty() {
// 		return ""
// 	}
// 	stackStr := []string{""}
// 	for _, st := range c.stacks {
// 		stackStr = append(stackStr, fmt.Sprint(st.data))
// 	}
// 	return fmt.Sprint(strings.Join(stackStr, "\n"))
// }

// func main() {
// 	cargo, linesInstr, err := createStacksAndInstructions("input5.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Stack before:", cargo)
// 	// log.Println("Instructions:", linesInstr)

// 	for _, instr := range linesInstr {
// 		err := cargo.executeInstruction(instr)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	log.Println("Stack after:", cargo)

// 	log.Println(cargo.ReturnTopCrates())
// }

// func createStacksAndInstructions(file string) (CargoArea, [][]int, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	linesStack := [][]string{}
// 	linesInstr := [][]int{}

// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() { // Extract crates and stack numbers
// 		if strings.TrimSpace(sc.Text()) == "" {
// 			// Indicates break between stack and directions; move onto next scan block
// 			break
// 		}
// 		chnks := chunks(sc.Text(), 4)
// 		linesStack = append(linesStack, chnks)
// 	}
// 	if err := sc.Err(); err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("scan file error: %w", err)
// 	}
// 	for sc.Scan() { // Extract numbers from instructions
// 		instrInt := make([]int, 3)
// 		instrSplit := strings.Split(sc.Text(), " ")
// 		for i := 1; i < len(instrInt)+1; i++ {
// 			instrInt[i-1], err = strconv.Atoi(instrSplit[i*2-1])
// 			if err != nil {
// 				return CargoArea{}, nil, fmt.Errorf("error converting to int: %w", err)
// 			}
// 		}
// 		linesInstr = append(linesInstr, instrInt)
// 	}
// 	if err := sc.Err(); err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("scan file error: %w", err)
// 	}

// 	// Parse stack lines into cargo area
// 	cargo := CargoArea{}
// 	for il := len(linesStack) - 2; il >= 0; il-- {
// 		for is, val := range linesStack[il] {
// 			// v := strings.Trim(strings.TrimSpace(val), "[ ]")
// 			v := strings.Trim(val, "[ ]")
// 			log.Println(v)
// 			if is+1 > len(cargo.stacks) {
// 				cargo.newEmptyStack()
// 			}
// 			if v != "" {
// 				cargo.stacks[is].push(v)
// 			}
// 		}
// 	}

// 	return cargo, linesInstr, nil
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

/*
--- Part Two ---
As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.
Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.
The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.
Again considering the example above, the crates begin in the same configuration:

				[D]
		[N] [C]
		[Z] [M] [P]
		1   2   3

Moving a single crate from stack 2 to stack 1 behaves the same as before:

		[D]
		[N] [C]
		[Z] [M] [P]
		1   2   3

However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

						[D]
						[N]
				[C] [Z]
				[M] [P]
		1   2   3

Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

						[D]
						[N]
		[C]     [Z]
		[M]     [P]
		1   2   3

Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

						[D]
						[N]
						[Z]
		[M] [C] [P]
		1   2   3

In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.
Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?
*/

// type Stack struct {
// 	crates []string
// }

// func (s *Stack) pop(n int) []string {
// 	if len(s.crates) == 0 {
// 		return []string{}
// 	}
// 	l := len(s.crates)
// 	if n > l {
// 		return s.crates
// 	}
// 	crates := s.crates[l-n:]
// 	s.crates = s.crates[:l-n]
// 	return crates
// }

// func (s *Stack) push(crates []string) {
// 	s.crates = append(s.crates, crates...)
// }

// type CargoArea struct {
// 	stacks []*Stack
// }

// func (c *CargoArea) newEmptyStack() int {
// 	c.stacks = append(c.stacks, &Stack{})
// 	return len(c.stacks)
// }

// func (c *CargoArea) addStackFromSlice(sl []string) {
// 	c.stacks = append(c.stacks, &Stack{crates: sl})
// }

// func (c *CargoArea) executeInstruction(instr []int) error {
// 	if instr[1] > len(c.stacks) || instr[2] > len(c.stacks) {
// 		return fmt.Errorf("executeInstruction: invalid instruction: stack number does not exist")
// 	}

// 	crates := c.stacks[instr[1]-1].pop(instr[0])
// 	c.stacks[instr[2]-1].push(crates)

// 	return nil
// }

// func (c *CargoArea) ReturnTopCrates() string {
// 	b := strings.Builder{}

// 	for _, st := range c.stacks {
// 		b.WriteString(st.crates[len(st.crates)-1])
// 	}

// 	return b.String()
// }

// func (c *CargoArea) isEmpty() bool {
// 	for _, s := range c.stacks {
// 		if len(s.crates) > 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func (c CargoArea) String() string {
// 	if c.isEmpty() {
// 		return ""
// 	}
// 	stackStr := []string{""}
// 	for _, st := range c.stacks {
// 		stackStr = append(stackStr, fmt.Sprint(st.crates))
// 	}
// 	return fmt.Sprint(strings.Join(stackStr, "\n"))
// }

// func main() {
// 	cargo, linesInstr, err := createStacksAndInstructions("input5.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Stack before:", cargo)
// 	// log.Println("Instructions:", linesInstr)

// 	for _, instr := range linesInstr {
// 		err := cargo.executeInstruction(instr)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	log.Println("Stack after:", cargo)

// 	log.Println(cargo.ReturnTopCrates())
// }

// func createStacksAndInstructions(file string) (CargoArea, [][]int, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	linesStack := [][]string{}
// 	linesInstr := [][]int{}

// 	sc := bufio.NewScanner(f)
// 	for sc.Scan() { // Extract crates and stack numbers
// 		if strings.TrimSpace(sc.Text()) == "" {
// 			// Indicates break between stack and directions; move onto next scan block
// 			break
// 		}
// 		chnks := chunks(sc.Text(), 4)
// 		linesStack = append(linesStack, chnks)
// 	}
// 	if err := sc.Err(); err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("scan file error: %w", err)
// 	}
// 	for sc.Scan() { // Extract numbers from instructions
// 		instrInt := make([]int, 3)
// 		instrSplit := strings.Split(sc.Text(), " ")
// 		for i := 1; i < len(instrInt)+1; i++ {
// 			instrInt[i-1], err = strconv.Atoi(instrSplit[i*2-1])
// 			if err != nil {
// 				return CargoArea{}, nil, fmt.Errorf("error converting to int: %w", err)
// 			}
// 		}
// 		linesInstr = append(linesInstr, instrInt)
// 	}
// 	if err := sc.Err(); err != nil {
// 		return CargoArea{}, nil, fmt.Errorf("scan file error: %w", err)
// 	}

// 	// Parse stack lines into cargo area
// 	cargo := CargoArea{}
// 	for il := len(linesStack) - 2; il >= 0; il-- {
// 		for is, val := range linesStack[il] {
// 			// v := strings.Trim(strings.TrimSpace(val), "[ ]")
// 			v := strings.Trim(val, "[ ]")
// 			if is+1 > len(cargo.stacks) {
// 				cargo.newEmptyStack()
// 			}
// 			if v != "" {
// 				cargo.stacks[is].push([]string{v})
// 			}
// 		}
// 	}

// 	return cargo, linesInstr, nil
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
