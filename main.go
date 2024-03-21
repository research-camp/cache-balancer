package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/go-faker/faker/v4"
)

// const variables for managing the balancing factors
const matchValue = 100000
const loadFactor = 10

func match(a string, b string) int {
	count := 0
	limit := int(math.Min(float64(len(a)), float64(len(b))))

	for i := 0; i < limit; i++ {
		if a[i] == b[i] {
			count = count + matchValue
		}
	}

	return count
}

func stringToBin(s string) string {
	binString := ""

	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}

	return binString
}

type MockFile struct {
	Name string `faker:"word"`
	Size int    `faker:"boundary_start=5, boundary_end=1000"`
}

func (f MockFile) BinaryName() string {
	return stringToBin(f.Name)
}

type Cache struct {
	IP    string `faker:"ipv4"`
	Local []*MockFile
}

func (c *Cache) ToString() string {
	size := 0

	for _, file := range c.Local {
		size += file.Size
	}

	return fmt.Sprintf("[%s] files: %d, sizes: %d", c.IP, len(c.Local), size)
}

func (c *Cache) Upload(file *MockFile) {
	c.Local = append(c.Local, file)
}

func (c *Cache) Factor() int {
	size := 0

	for _, file := range c.Local {
		size += file.Size
	}

	if size == 0 {
		return 1
	}

	avg := 4 * size / len(c.Local) // 4 = log2(count(caches))
	delta := size * len(c.Local) / avg

	delta = int(math.Max(float64(delta/loadFactor), 1))

	return delta
}

func (c *Cache) Value(input string) int {
	return match(stringToBin(c.IP), input)
}

func (c *Cache) Download(input string) bool {
	for _, item := range c.Local {
		if item.Name == input {
			return true
		}
	}

	return false
}

func generateFiles(number int) []*MockFile {
	var list []*MockFile

	for i := 0; i < number; i++ {
		tmp := new(MockFile)
		_ = faker.FakeData(&tmp)
		list = append(list, tmp)
	}

	return list
}

func generateCaches(number int) []*Cache {
	var list []*Cache

	for i := 0; i < number; i++ {
		tmp := &Cache{}
		_ = faker.FakeData(&tmp)

		tmp.Local = make([]*MockFile, 0)

		list = append(list, tmp)
	}

	return list
}

func sortFunction(bname string, caches []*Cache) func(int, int) bool {
	return func(i, j int) bool {
		return caches[j].Value(bname)*caches[i].Factor() < caches[i].Value(bname)*caches[j].Factor()
	}
}

func main() {
	// creating inputs
	inputs := generateFiles(100)
	caches := generateCaches(10)

	fmt.Printf("Load factor: %d\nMatch value: %d\n\n", loadFactor, matchValue)

	// upload files
	for _, upload := range inputs {
		sort.Slice(caches, sortFunction(upload.BinaryName(), caches))
		caches[0].Upload(upload)
	}

	fmt.Println("======== caches =========")

	for _, cache := range caches {
		fmt.Println(cache.ToString())
	}

	fmt.Println("======== ------ ========")

	success := 0

	// download files
	for _, upload := range inputs {
		sort.Slice(caches, func(i, j int) bool {
			tmp := upload.BinaryName()
			return caches[i].Value(tmp) > caches[j].Value(tmp)
		})

		if caches[0].Download(upload.Name) {
			success++
		}
	}

	fmt.Printf("\nResolved %d outof %d (%d%%)\n", success, len(inputs), 100*success/len(inputs))
}
