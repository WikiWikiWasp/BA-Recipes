package main

import (
	"fmt"
	"local/recipes/config"
)

// cuisines := make([]string, 0)
// ing := make([]string, 0)

//Recipe holds the needed info for a full recipe
type Recipe struct {
	name        string
	cuisine     string
	servings    int
	time        []int
	description string
	ingredients map[string]string
	seasoning   []string
	steps       map[int]map[string][]string
}

func main() {

	// cuisines := make([]string, 0)
	// cuisines = []string{
	// 	"chinese",
	// 	"japanese",
	// 	"thai",
	// 	"korean",
	// 	"indian",
	// 	"middle eastern",
	// 	"french",
	// 	"american",
	// 	"italian",
	// 	"mexican",
	// 	"spanish",
	// 	"greek",
	// 	"Mediterranean",
	// }

	client := config.ConnectToMongo()
	coll := config.CreateCollection(client)
	fmt.Println("Collection `", coll.Name(), "` created!")

	// r1 := Recipe{
	// 	name:        "Smoky Chickpea & Pepper Shakshuka\nwith Feta & Mint",
	// 	cuisine:     "Middle Eastern",
	// 	servings:    2,
	// 	time:        []int{25, 35},
	// 	description: "This shakshuka - a dish of eggs baked in a rustic sauce of tomato, peppers,\nand onion - gets extra heartiness from plump chickpeas.\nHarissa paste gives the sauce its characteristacally bold, smoky flavor,\nwhile the classic Middle Eastern spices of za'atar give our side of toasted\npita wedges an irresistibly herbaceous lift.",
	// 	ingredients: map[string]string{
	// 		"egg":                               "2",
	// 		"whole peeled san marzano tomatoes": "14 oz",
	// 		"red onion":                         "1",
	// 		"garlic":                            "2 cloves",
	// 		"red harissa paste":                 "1 tbsp",
	// 		"sliced roasted red peppers":        "1 oz",
	// 		"chickpeas":                         "7.75 oz",
	// 		"pocketless pita":                   "2",
	// 		"sweet peppers":                     "4 oz",
	// 		"mint":                              "1 bunch",
	// 		"feta cheese":                       "1.5 oz",
	// 		"za'atar seasoning":                 "1 tbsp",
	// 	},
	// 	seasoning: []string{"sumac", "marjoram", "oregano", "thyme", "sesame seeds", "salt"},
	// 	steps: map[int]map[string][]string{
	// 		1: {
	// 			"Prepare the ingredients": []string{
	// 				"Wash, dry the fresh produce.",
	// 				"Halve, peel, and medium dice the onion.",
	// 				"Peel, roughly chop 2 cloves of garlic.",
	// 				"Cut off, discard the stems of the sweet peppers; remove cores, thinly slice into rings",
	// 				"Drain, rinses chickpeas.",
	// 				"Place tomatoes in a bowl; gently break apart with hands.",
	// 				"Roughly chop roasted peppers.",
	// 				"Pick mint leaves off stems",
	// 			}},
	// 		2: map[string][]string{
	// 			"Make the za'atar oil": []string{
	// 				"In a medium pan, heat 2 tablespoon of olive oil on med-high until hot.",
	// 				"Add half the za'atar. Cook, stirring frequently, 1 to 2 minutes, or until fragrant.",
	// 				"Transfer to a bowl.",
	// 				"Rinse and wipe out the pan.",
	// 			}},
	// 		3: map[string][]string{
	// 			"Cook the vegetables": []string{
	// 				"In the same pan, heat 1 teaspoon of olive oil on med-high until hot.",
	// 				"Add the diced onion and chopped garlic; season with salt and pepper.",
	// 				"Cook, stirring occasionally, 3-4 mins, or until lightly browned and slightly softened.",
	// 				"Add the sliced sweet peppers.",
	// 				"Cook, stirring occasionally, 1-2 mins, or until slightly softened.",
	// 			}},
	// 		4: map[string][]string{
	// 			"Start the shakshuka": []string{
	// 				"Add the harissa paste and chickpeas to the pan.",
	// 				"Cook, stirring frequently, 30sec-1min, or until thoroughly combined.",
	// 				"Add the tomatoes and chopped roasted peppers; season with S&P.",
	// 				"Cook, stirring occasionally, 2-3mins, or until slightly thickened.",
	// 				"Turn off the heat.",
	// 			}},
	// 		5: map[string][]string{
	// 			"Make the pita wedges": []string{
	// 				"Head a separate dry, medium pan on med-high until hot.",
	// 				"Working one at a time, add the pitas. Cook 1-2mins per side, or until heated through and pliable.",
	// 				"Carefully transfer to a work surface and brush or drizzle with the za'atar oil.",
	// 				"Cut into equal-sized wedges.",
	// 			}},
	// 		6: map[string][]string{
	// 			"Finish the shakshuka & serve your dish": []string{
	// 				"Using a spoon, create 2 shallow wells in the center of the shakshuka.",
	// 				"Carefully crack an egg into each well. Cover the pan with foil and cook 4-5mins, or until the egg whites are set and the yolks are cooked to your desired degree of doneness.",
	// 				"Turn off the heat.",
	// 				"Serve the finished shakshuka garnished with the:\n\tcheese (crumbling before adding)\n\tmint leaves (tearing just before adding)\n\tand remaining za'atar.",
	// 				"Server with the pita wedges on the side. Enjoy!",
	// 			}},
	// 	},
	// }

	// printRecipe(r1)
}

func printRecipe(rec Recipe) {

	var colors = map[string]string{
		"HEADER":    "\033[95m",
		"OKBLUE":    "\033[94m",
		"OKGREEN":   "\033[92m",
		"WARNING":   "\033[93m",
		"FAIL":      "\033[91m",
		"ENDC":      "\033[0m",
		"BOLD":      "\033[1m",
		"UNDERLINE": "\033[4m",
	}

	fmt.Printf("%s", colors["OKBLUE"])
	fmt.Printf("\n---------------------------------------------\n")
	fmt.Printf("%s", rec.name)
	fmt.Printf("\n---------------------------------------------\n")
	fmt.Printf("%s", colors["ENDC"])

	fmt.Printf("\n%sTIME%s: %d-%d minutes", colors["HEADER"], colors["ENDC"], rec.time[0], rec.time[1])
	fmt.Printf("\n%sSERVINGS%s: %d", colors["HEADER"], colors["ENDC"], rec.servings)

	fmt.Printf("\n\n%s%s%s\n\n", colors["OKGREEN"], rec.description, colors["ENDC"])

	fmt.Printf("%sINGREDIENTS%s:\n", colors["HEADER"], colors["ENDC"])
	for i, v := range rec.ingredients {
		fmt.Printf("%s%10s %s%s\n", colors["WARNING"], v, colors["ENDC"], i)
	}

	fmt.Printf("\n%sSTEPS%s:\n", colors["HEADER"], colors["ENDC"])

	for i := 1; i <= len(rec.steps); i++ {
		for k, v := range rec.steps {
			if k == i {
				fmt.Printf("%s%d. ", colors["WARNING"], i)
				for step := range v {
					fmt.Printf("%s%s\n", step, colors["ENDC"])
					for substep := range rec.steps[k][step] {
						fmt.Printf("   - %s\n", rec.steps[k][step][substep])
					}
				}
			}
		}
	}
}
