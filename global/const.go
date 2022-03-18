package global

const (
	dburi       = "mongodb+srv://root:Megamind@cluster0.2qauw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbname      = "blog-application"
	performance = 100
)

var (
	jWTSecret = []byte("blogSecret")
)
