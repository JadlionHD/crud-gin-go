package seeds

import (
	"os"

	"github.com/JadlionHD/crud-gin-go/internal/database/migration"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) bool {
	if _, err := os.Stat("test.db"); err == nil {
		return false
	}

	posts := []*migration.PostDB{
		{
			UserID: 1,
			Title:  "His mother had always taught him",
			Body:   "His mother had always taught him not to ever think of himself as better than others. He'd tried to live by this motto. He never looked down on those who were less fortunate or who had less money than him. But the stupidity of the group of people he was talking to made him change his mind.",
			Tags:   []string{"history", "american", "crime"},
			Reactions: &migration.PostReactionDB{
				Likes:    293,
				Dislikes: 5,
			},
			Views: 523,
		},
		{
			UserID: 1,
			Title:  "The old town clock struck midnight",
			Body:   "The old town clock struck midnight, and in the eerie silence that followed, the whispers of the past began to stir once again. People had long forgotten the mysteries buried in the town’s history, but something was about to bring them to the surface.",
			Tags:   []string{"mystery", "history", "supernatural"},
			Reactions: &migration.PostReactionDB{
				Likes:    432,
				Dislikes: 20,
			},
			Views: 674,
		},
		{
			UserID: 3,
			Title:  "The last train left without me",
			Body:   "As the train doors closed and the train slowly disappeared into the night, I realized that I was completely alone in the dark. I thought about all the things I could have done differently, all the chances I missed, and the empty silence that awaited me now.",
			Tags:   []string{"sadness", "regret", "travel"},
			Reactions: &migration.PostReactionDB{
				Likes:    103,
				Dislikes: 10,
			},
			Views: 235,
		},
		{
			UserID: 2,
			Title:  "The secret beneath the lake",
			Body:   "Everyone knew the legends about the lake, but no one ever took them seriously. It wasn’t until I decided to dive beneath the water's surface that I discovered the truth. There was something down there, something that had been hidden for centuries.",
			Tags:   []string{"adventure", "mystery", "legends"},
			Reactions: &migration.PostReactionDB{
				Likes:    322,
				Dislikes: 9,
			},
			Views: 675,
		},
		{
			UserID: 4,
			Title:  "A letter from the past",
			Body:   "The old letter had been tucked away in a dusty attic for years, its edges yellowed and fragile. But when I finally opened it, I found a message that would change everything I thought I knew about my family and their history.",
			Tags:   []string{"aw", "a"},
			Reactions: &migration.PostReactionDB{
				Likes:    320,
				Dislikes: 5,
			},
			Views: 800,
		},
	}

	db.Create(posts)

	return true
}
