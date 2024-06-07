package component

func SnowData() AccountList {
	var accounts AccountList = [Nmax]Account{
		{
			Username:       "john_doe",
			YouTubeChannel: "JohnDoeVlogs",
			Email:          "john.doe@example.com",
			Status:         "silver",
			Subscribers:    15000,
			Balance:        1200,
			Videos: VideoList{
				{Title: "My First Vlog", Duration: 10, ViewCount: 15000, PublishDate: Date{Day: 1, Month: 1, Year: 2020}},
				{Title: "A Day in the Life", Duration: 15, ViewCount: 18000, PublishDate: Date{Day: 15, Month: 2, Year: 2020}},
			},
		},
		{
			Username:       "tech_guru",
			YouTubeChannel: "TechGuruReviews",
			Email:          "tech.guru@example.com",
			Status:         "Active",
			Subscribers:    25000,
			Balance:        2300,
			Videos: VideoList{
				{Title: "Latest Tech Review", Duration: 20, ViewCount: 22000, PublishDate: Date{Day: 10, Month: 3, Year: 2021}},
				{Title: "Unboxing the Future", Duration: 12, ViewCount: 25000, PublishDate: Date{Day: 5, Month: 5, Year: 2021}},
			},
		},
		{
			Username:       "cooking_queen",
			YouTubeChannel: "CookingWithQueen",
			Email:          "cooking.queen@example.com",
			Status:         "Active",
			Subscribers:    30000,
			Balance:        3000,
			Videos: VideoList{
				{Title: "Easy Pasta Recipe", Duration: 8, ViewCount: 31000, PublishDate: Date{Day: 20, Month: 4, Year: 2021}},
				{Title: "Baking 101", Duration: 14, ViewCount: 28000, PublishDate: Date{Day: 18, Month: 6, Year: 2021}},
			},
		},
		{
			Username:       "travel_buddy",
			YouTubeChannel: "TravelWithBuddy",
			Email:          "travel.buddy@example.com",
			Status:         "Active",
			Subscribers:    18000,
			Balance:        1500,
			Videos: VideoList{
				{Title: "Exploring Bali", Duration: 18, ViewCount: 17000, PublishDate: Date{Day: 5, Month: 8, Year: 2020}},
				{Title: "Top 10 Travel Tips", Duration: 20, ViewCount: 16000, PublishDate: Date{Day: 12, Month: 9, Year: 2020}},
			},
		},
		{
			Username:       "gamer_pro",
			YouTubeChannel: "ProGamerPlay",
			Email:          "gamer.pro@example.com",
			Status:         "Active",
			Subscribers:    22000,
			Balance:        2000,
			Videos: VideoList{
				{Title: "Epic Game Highlights", Duration: 25, ViewCount: 21000, PublishDate: Date{Day: 22, Month: 11, Year: 2021}},
				{Title: "Game Review: XYZ", Duration: 30, ViewCount: 22000, PublishDate: Date{Day: 2, Month: 1, Year: 2022}},
			},
		},
		{
			Username:       "fitness_freak",
			YouTubeChannel: "FitnessFreakDaily",
			Email:          "fitness.freak@example.com",
			Status:         "Active",
			Subscribers:    28000,
			Balance:        2700,
			Videos: VideoList{
				{Title: "Morning Workout Routine", Duration: 15, ViewCount: 30000, PublishDate: Date{Day: 8, Month: 3, Year: 2021}},
				{Title: "Healthy Eating Tips", Duration: 12, ViewCount: 29000, PublishDate: Date{Day: 10, Month: 5, Year: 2021}},
			},
		},
		{
			Username:       "music_lover",
			YouTubeChannel: "MusicLoverChannel",
			Email:          "music.lover@example.com",
			Status:         "Active",
			Subscribers:    32000,
			Balance:        3500,
			Videos: VideoList{
				{Title: "Top 10 Songs of 2021", Duration: 10, ViewCount: 34000, PublishDate: Date{Day: 20, Month: 7, Year: 2021}},
				{Title: "Music Review: Album ABC", Duration: 18, ViewCount: 33000, PublishDate: Date{Day: 25, Month: 9, Year: 2021}},
			},
		},
		{
			Username:       "science_geek",
			YouTubeChannel: "ScienceGeekExplains",
			Email:          "science.geek@example.com",
			Status:         "Active",
			Subscribers:    26000,
			Balance:        2400,
			Videos: VideoList{
				{Title: "Physics Explained", Duration: 22, ViewCount: 25000, PublishDate: Date{Day: 1, Month: 11, Year: 2020}},
				{Title: "Chemistry Basics", Duration: 15, ViewCount: 26000, PublishDate: Date{Day: 15, Month: 12, Year: 2020}},
			},
		},
		{
			Username:       "fashionista",
			YouTubeChannel: "FashionistaStyle",
			Email:          "fashionista@example.com",
			Status:         "Active",
			Subscribers:    24000,
			Balance:        2200,
			Videos: VideoList{
				{Title: "Summer Fashion Trends", Duration: 12, ViewCount: 23000, PublishDate: Date{Day: 2, Month: 6, Year: 2021}},
				{Title: "Winter Wardrobe Essentials", Duration: 14, ViewCount: 24000, PublishDate: Date{Day: 20, Month: 12, Year: 2021}},
			},
		},
		{
			Username:       "movie_buff",
			YouTubeChannel: "MovieBuffReviews",
			Email:          "movie.buff@example.com",
			Status:         "Active",
			Subscribers:    20000,
			Balance:        1800,
			Videos: VideoList{
				{Title: "Movie Review: XYZ", Duration: 20, ViewCount: 19000, PublishDate: Date{Day: 12, Month: 10, Year: 2021}},
				{Title: "Top 10 Movies of 2021", Duration: 18, ViewCount: 20000, PublishDate: Date{Day: 5, Month: 1, Year: 2022}},
			},
		},
	}
	return accounts
}
