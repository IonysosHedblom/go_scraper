package entity

type ScraperRequest struct {
	Ingredients []string
}

type RecipeDetailsResponse struct {
	Context              string               `json:"@context"`
	Type                 string               `json:"@type"`
	Name                 string               `json:"name"`
	ImageUrl             string               `json:"image"`
	RecipeUrl            string               `json:"url"`
	Description          string               `json:"description"`
	DatePublished        string               `json:"datePublished"`
	DateModified         string               `json:"dateModified"`
	Author               Author               `json:"author"`
	Rating               Rating               `json:"aggregateRating"`
	TotalTime            string               `json:"totalTime"`
	CookingMethod        string               `json:"cookingMethod"`
	RecipeCategory       string               `json:"recipeCategory"`
	RecipeCuisine        string               `json:"recipeCuisine"`
	RecipeYield          string               `json:"recipeYield"`
	Nutrition            Nutrition            `json:"nutrition"`
	RecipeIngredient     []string             `json:"recipeIngredient"`
	RecipeInstructions   []Instruction        `json:"recipeInstructions"`
	InteractionStatistic InteractionStatistic `json:"interactionStatistic"`
	CommentCount         int                  `json:"commentCount"`
}

type Instruction struct {
	Type string `json:"@type"`
	Text string `json:"text"`
}

type InteractionStatistic struct {
	Type                 string `json:"@type"`
	InteractionType      string `json:"interactionType"`
	UserInteractionCount int    `json:"userInteractionCount"`
}

type Nutrition struct {
	Type                string `json:"@type"`
	ServingSize         string `json:"servingSize"`
	Calories            string `json:"calories"`
	FatContent          string `json:"fatContent"`
	CarboHydrateContent string `json:"carbohydrateContent"`
	ProteinContent      string `json:"proteinContent"`
}

type Rating struct {
	Type        string  `json:"@type"`
	RatingValue float64 `json:"ratingValue"`
	ReviewCount int     `json:"reviewCount"`
}

type Author struct {
	Type string `json:"@type"`
	Name string `json:"name"`
}
