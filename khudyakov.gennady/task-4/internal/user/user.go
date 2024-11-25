package user

type User struct {
	Email string
}

const (
	UniqueUsers = 10
)

var emails = []string{
	"eldon_avalos73@gmail.com",
	"lilian_grimes99@gmail.com",
	"justin_meadows11@gmail.com",
	"curtis_buddy61@gmail.com",
	"allison-francis19@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
	"barr_natasha94@gmail.com",
	"madeleine_cline6@gmail.com",
	"tiffanie_mckay74@gmail.com",
	"hollis-javier66@gmail.com",
	"danette-boone1@gmail.com",
}

func GenerateUsers(out chan<- User) {
	for _, email := range emails {
		out <- User{Email: email}
	}
	close(out)
}
