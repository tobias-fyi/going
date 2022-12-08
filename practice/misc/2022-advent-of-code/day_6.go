package main

// === ===
// === 2022 Advent of Code, Day 6 === //
// === ===

/*
--- Day 6: Tuning Trouble ---
The preparations are finally complete; you and the Elves leave camp on foot and begin to make your way toward the star fruit grove.
As you move through the dense undergrowth, one of the Elves gives you a handheld device. He says that it has many fancy features, but the most important one to set up right now is the communication system.
However, because he's heard you have significant experience dealing with signal-based systems, he convinced the other Elves that it would be okay to give you their one malfunctioning device - surely you'll have no problem fixing it.
As if inspired by comedic timing, the device emits a few colorful sparks.
To be able to communicate with the Elves, the device needs to lock on to their signal. The signal is a series of seemingly-random characters that the device receives one at a time.
To fix the communication system, you need to add a subroutine to the device that detects a start-of-packet marker in the datastream. In the protocol being used by the Elves, the start of a packet is indicated by a sequence of four characters that are all different.
The device will send your subroutine a datastream buffer (your puzzle input); your subroutine needs to identify the first position where the four most recently received characters were all different. Specifically, it needs to report the number of characters from the beginning of the buffer to the end of the first such four-character marker.
For example, suppose you receive the following datastream buffer:

		mjqjpqmgbljsphdztnvjfqwrcgsmlb

After the first three characters (mjq) have been received, there haven't been enough characters received yet to find the marker. The first time a marker could occur is after the fourth character is received, making the most recent four characters mjqj. Because j is repeated, this isn't a marker.
The first time a marker appears is after the seventh character arrives. Once it does, the last four characters received are jpqm, which are all different. In this case, your subroutine should report the value 7, because the first start-of-packet marker is complete after 7 characters have been processed.
Here are a few more examples:

* `bvwbjplbgvbhsrlpgdmjqwftvncz`: first marker after character 5
* `nppdvjthqldpwncqszvftbrmjlhg`: first marker after character 6
* `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: first marker after character 10
* `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`: first marker after character 11

How many characters need to be processed before the first start-of-packet marker is detected?
*/

// type Queue interface {
// 	Enqueue(v string) error
// 	Dequeue() (string, error)
// 	UniqueValues() bool
// }

// type PacketQueue struct {
// 	buffer []string
// 	maxLen int
// 	count  int
// }

// func newPacketQueue(maxLen int) PacketQueue {
// 	return PacketQueue{
// 		buffer: []string{},
// 		maxLen: maxLen,
// 		count:  0,
// 	}
// }

// func (p *PacketQueue) Enqueue(v string) error {
// 	if len(p.buffer) == p.maxLen {
// 		return fmt.Errorf("queue full: %s not added", v)
// 	}
// 	p.buffer = append(p.buffer, v)
// 	log.Printf("Enqueued %s: %v\n", v, p.buffer)
// 	p.count++
// 	return nil
// }

// func (p *PacketQueue) Dequeue() (string, error) {
// 	if len(p.buffer) == 0 {
// 		return "", fmt.Errorf("queue is empty: cannot dequeue")
// 	}
// 	v := p.buffer[0]
// 	log.Printf("Dequeued %s: %v\n", v, p.buffer)
// 	p.buffer[0] = ""
// 	p.buffer = p.buffer[1:]
// 	return v, nil
// }

// func (p *PacketQueue) UniqueValues() bool {
// 	keys := make(map[string]bool)
// 	list := []string{}
// 	for _, s := range p.buffer {
// 		if _, value := keys[s]; !value {
// 			keys[s] = true
// 			list = append(list, s)
// 		}
// 	}
// 	return len(list) == len(p.buffer)
// }

// func main() {
// 	signal, err := readFileToString("input6.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(signal)

// 	ml := 4
// 	pq := newPacketQueue(ml)

// 	for i, s := range signal {
// 		if len(pq.buffer) == ml {
// 			if pq.UniqueValues() {
// 				log.Println("Index:", i)
// 				log.Println("Count:", pq.count)
// 				break
// 			}

// 			if _, err := pq.Dequeue(); err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		if err := pq.Enqueue(string(s)); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func readFileToString(file string) (string, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return "", fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	bytes, err := io.ReadAll(f)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading file: %w", err)
// 	}

// 	return string(bytes), nil
// }

/*
--- Part Two ---
Your device's communication system is correctly detecting packets, but still isn't working. It looks like it also needs to look for messages.
A start-of-message marker is just like a start-of-packet marker, except it consists of 14 distinct characters rather than 4.
Here are the first positions of start-of-message markers for all of the above examples:

- `mjqjpqmgbljsphdztnvjfqwrcgsmlb`: first marker after character 19
- `bvwbjplbgvbhsrlpgdmjqwftvncz`: first marker after character 23
- `nppdvjthqldpwncqszvftbrmjlhg`: first marker after character 23
- `nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg`: first marker after character 29
- `zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw`: first marker after character 26

How many characters need to be processed before the first start-of-message marker is detected?
*/

// type Queue interface {
// 	Enqueue(v string) error
// 	Dequeue() (string, error)
// 	UniqueValues() bool
// }

// type PacketQueue struct {
// 	buffer []string
// 	maxLen int
// 	count  int
// }

// func newPacketQueue(maxLen int) PacketQueue {
// 	return PacketQueue{
// 		buffer: []string{},
// 		maxLen: maxLen,
// 		count:  0,
// 	}
// }

// func (p *PacketQueue) Enqueue(v string) error {
// 	if len(p.buffer) == p.maxLen {
// 		return fmt.Errorf("queue full: %s not added", v)
// 	}
// 	p.buffer = append(p.buffer, v)
// 	log.Printf("Enqueued %s: %v\n", v, p.buffer)
// 	p.count++
// 	return nil
// }

// func (p *PacketQueue) Dequeue() (string, error) {
// 	if len(p.buffer) == 0 {
// 		return "", fmt.Errorf("queue is empty: cannot dequeue")
// 	}
// 	v := p.buffer[0]
// 	log.Printf("Dequeued %s: %v\n", v, p.buffer)
// 	p.buffer[0] = ""
// 	p.buffer = p.buffer[1:]
// 	return v, nil
// }

// func (p *PacketQueue) UniqueValues() bool {
// 	keys := make(map[string]bool)
// 	list := []string{}
// 	for _, s := range p.buffer {
// 		if _, value := keys[s]; !value {
// 			keys[s] = true
// 			list = append(list, s)
// 		}
// 	}
// 	return len(list) == len(p.buffer)
// }

// func main() {
// 	signal, err := readFileToString("input6.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println(signal)

// 	ml := 14
// 	pq := newPacketQueue(ml)

// 	for i, s := range signal {
// 		if len(pq.buffer) == ml {
// 			if pq.UniqueValues() {
// 				log.Println("Index:", i)
// 				log.Println("Count:", pq.count)
// 				break
// 			}

// 			if _, err := pq.Dequeue(); err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		if err := pq.Enqueue(string(s)); err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// }

// func readFileToString(file string) (string, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return "", fmt.Errorf("error opening file: %w", err)
// 	}
// 	defer f.Close()

// 	bytes, err := io.ReadAll(f)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading file: %w", err)
// 	}

// 	return string(bytes), nil
// }
