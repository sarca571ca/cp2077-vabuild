package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Basic program that reads a list in a file and trims the contents to allow for it to fit in a new format.
// Specifically for using items-code list on clothing mods for Cyberpunk 2077 and building them into a
// virtual atelier store format.

func main() {
	buildStore()
}

func importFile(fileName string) {
	// Opens the designated file "importFile("file")" to trim and reformat
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	// Creates an output file for the new format to be called later on.
	// This can probably done in memory by putting each line into a slice.
	// Future adjustments: Send directly to a slice.
	items, err := os.Create("items.txt")
	if err != nil {
		fmt.Printf("Failed to create file: %s", err)
	}

	defer items.Close()

	// Uses Fprint to send lines to a new file. Again should send to a slice instead.
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		triml := strings.TrimLeft(line, "Game.AddToInventory(")
		newline := strings.TrimRight(triml, ", 1)")
		fmt.Fprintf(items, "    %s\n", newline)
	}
}

func buildStore() {

	// All the final formating is completed here.
	fmt.Printf("Building Virtual Atelier Store file from item-list.txt\n")

	// Creating variables for the data value name cannot have any spaces or special characters
	// Future adjustments: Beef up the error checks to stop special characters and spaces from
	// being used.
	fmt.Print("\nEnter the data name of the store. Cannot contain spaces or special characters.\n")
	var dsname string
	_, err := fmt.Scanf("%s\n", &dsname)
	if err != nil {
		fmt.Println(err)
	}

	// Creating variable for the Actual in Game name of the store. This can have spaces and special
	// characters but currently doesn't support spaces.
	// Future adjustments: Add in allowances for reading spaces.
	fmt.Print("Enter the in game name of the store.\n")
	sc3 := bufio.NewScanner(os.Stdin)
	sc3.Scan()
	asname := sc3.Text()
	if asname == "" {
		fmt.Println("Field cannot be blank.")
		fmt.Println("Press Enter/Return to retry.")
		_, w := fmt.Scanln()
		if w != nil {
			fmt.Println("Press Enter/Return.")
		}
		buildStore()
	}

	// Creates and names the file that will store the new format. The data name above must match the
	// first part of the filename "dataname-atelier-store.reds".
	fname := fmt.Sprintf("%s-atelier-store.reds", dsname)
	file, err := os.Create(fname)
	if err != nil {
		fmt.Printf("Failed to create file: %s", err)
	}

	defer file.Close()

	// Initial formating string so the Virtual Atelier mod can read the file.
	output := fmt.Sprintf(`@addMethod(gameuiInGameMenuGameController)
protected cb func RegisterTheDreamShopStore(event: ref<VirtualShopRegistration>) -> Bool {
  event.AddStore(
    n"%s",
    "%s",
    [`, dsname, asname)

	// Calling the importFile() function above to parse the specified file.
	// Future adjustments: maybe add tags to call an file of the users choosing
	// as many modders package a file with item codes with their mods also.
	importFile("item-list.txt")

	// Opening the output file from above. Again need to initially store this in a slice.
	f2, err := os.Open("items.txt")
	if err != nil {
		fmt.Printf("Failed to open file: %s", err)
	}

	defer f2.Close()

	// Creates a slice of the "items.txt" output file.
	sc2 := bufio.NewScanner(f2)
	l2 := make([]string, 0)
	for sc2.Scan() {
		l2 = append(l2, "\n", sc2.Text())
	}
	if err := sc2.Err(); err != nil {
		fmt.Printf("Process failed: %s", err)
	}

	// Last section of the formating.
	output2 := fmt.Sprintf(`
    ],
    [],
    r"base/gameplay/gui/world/adverts/naranjita/naranjite_atlas.inkatlas",
    n"CHARACTER",
    []
  );
}`)

	// Trimming the [] created from the slice output.
	l3 := strings.Trim(fmt.Sprint(l2), "[]")

	// Printing final file with formating that the Virtual Atelier shop can read.
	fmt.Fprint(file, "", output, l3, output2)

	// Deletes the "items.txt" that was created earlier.
	os.Remove("items.txt")
}
